# gardener-operator memory analysis (runtime cluster) – first pass

Date: 2026-11-12. Cluster: `dev-setup/kubeconfigs/runtime` (local KinD setup).

## Measurements

| Process                                                         | RSS     | Go heap alloc | Go heap in use | Go sys  | Goroutines |
| --------------------------------------------------------------- | ------- | ------------- | -------------- | ------- | ---------- |
| `cmd/dummy` (watches Secret, ConfigMap, Pod; run locally, ~25s) | ~37 MB  | ~15 MB        | ~19 MB         | n/a     | n/a        |
| `gardener-operator` (in-cluster, pod age ~28m)                  | ~212 MB | ~70 MB        | ~106 MB        | ~367 MB | 703        |

`kubectl top` reports ~160Mi for the operator pod (working set, differs from RSS). The operator deployment has **no resource requests/limits** and no `GOMEMLIMIT`, so the Go GC is not bounded.

Cluster size (JSON size of full lists, a rough upper bound for the cache size of the typed objects):

- Secrets: 375 objects, ~3.6 MB
- ConfigMaps: 116 objects, ~2.2 MB
- Pods: 120 objects, ~2.5 MB

So the dummy baseline (all three kinds, cluster-wide) costs only ~15 MB heap / ~37 MB RSS. Caching Secrets/ConfigMaps/Pods does **not** explain the operator's footprint (~175 MB RSS more).

## Observations

1. The operator uses the default manager cache (`cmd/gardener-operator/app/app.go`): no `Cache` options (no `ByObject`, label/field selectors, namespace restriction, or `DisableFor`). Every kind read via `mgr.GetClient()` (Get/List) or watched creates a cluster-wide informer holding full objects.
2. Explicit watches in `pkg/operator/controller`: `Garden`, `Extension`, `Secret` (garden controller, token-requestor), `Secret` + `ConfigMap` (extension controller). Secrets/ConfigMaps are watched cluster-wide although the operator mainly cares about the `garden` namespace (and some predicates filter only after the object is cached).
3. Implicit informers via the cached client (Get/List of typed objects in care/health checks, garden reconciler, networkpolicy, etc.): `ManagedResource`, `Namespace`, `Secret`, `ConfigMap`, `Extension`/`DNSRecord`/`BackupBucket`/`BackupEntry`, `Prometheus`, and many more used by the component deployers. These are probably much of the unaccounted memory, but not verified yet.
4. 703 goroutines (several MB of stacks: ~5 MB) – each informer, watch, workqueue, and the virtual-garden client map adds some.
5. Go sys (367 MB) far above RSS (212 MB) and heap in use (106 MB) vs alloc (70 MB): lots of GC headroom/fragmentation; with no memory limit the GC lets the heap grow (GOGC=100). Reconciles that render charts/manifests/image vectors create short-lived garbage that inflates the peak.
6. The operator also loads a very large scheme (`operatorclient.RuntimeScheme`: Gardener, extensions, resources, druid, monitoring, VPA, etc.) and the image vector at startup; this is a static cost that the dummy does not have. Its size is not measured yet.
7. Operator is more active than the dummy: 866 PUT, 191 GET, 156 PATCH requests so far; a Garden reconcile runs periodically (syncPeriod 1h) and on changes.

## Conclusion (preliminary)

The baseline for "a controller watching many objects" is small (~15 MB heap here), so the operator's ~100 MB+ heap is not caused by Secrets/ConfigMaps/Pods caches alone. Likely contributors, to be confirmed: (a) many additional cluster-wide informers from the cached client, (b) transient allocations during Garden/Extension reconciliation, (c) static costs (scheme, image vector, charts), (d) unbounded GC.

## Next steps

1. Enable profiling (`debugging.enableProfiling: true` in operator config), port-forward `:2752`, and take `heap` and `allocs` profiles (`go tool pprof -sample_index=inuse_space`). Compare right after startup vs. after a Garden reconcile.
2. List informers actually started (e.g. log `cache.GetInformer` kinds or use the `rest_client`/watch metrics) to quantify per-kind cost.
3. Extend `cmd/dummy` incrementally: add the operator's scheme, then ManagedResource/Namespace/Extension informers, to see the per-step cost.
4. Experiment with cache restrictions (namespace-scoped / label selectors for Secrets and ConfigMaps, `DisableFor` + `APIReader` for rarely used kinds, transform functions to strip `managedFields`) and with `GOMEMLIMIT`/resource limits.
5. Use a larger/more realistic cluster; here the cluster is small, so cache cost will grow with real data (many shoots => many Secrets/ManagedResources).

Caveat: measurements are single snapshots on a small dev cluster; dummy ran locally, operator in-cluster.

## Update: dummy extended with all discovered watches

`cmd/dummy` now uses `operatorclient.RuntimeScheme` / `VirtualScheme` and watches the types the operator logs as event sources (`Starting EventSource`):

- Runtime cluster: Secret, ConfigMap, Pod (baseline), Garden, Extension (operator), ManagedResource, BackupBucket, BackupEntry, DNSRecord, Namespace, Endpoints, NetworkPolicy, VerticalPodAutoscaler.
- Virtual garden (second `cluster.Cluster`, `VIRTUAL_KUBECONFIG`), via `WatchesRawSource`: ControllerInstallation, Gardenlet.

Run: `KUBECONFIG=dev-setup/kubeconfigs/runtime/kubeconfig VIRTUAL_KUBECONFIG=dev-setup/kubeconfigs/virtual-garden/kubeconfig go run ./cmd/dummy`

| Process | RSS | Heap alloc | Heap in use | Sys | Goroutines |
|---|---|---|---|---|---|
| dummy v1 (Secret, ConfigMap, Pod) | ~37 MB | ~15 MB | ~19 MB | – | – |
| dummy v2, t=20s | ~80 MB | ~27 MB | ~33 MB | ~51 MB | 211 |
| dummy v2, t=60s | ~44 MB | ~27 MB | ~34 MB | ~51 MB | 211 |
| gardener-operator | ~212 MB | ~70 MB | ~106 MB | ~367 MB | 703 |

(RSS spike at startup is the initial list/decode garbage, released later.)

Takeaways:
- Adding the full scheme and all watched kinds roughly doubles the heap (15 → 27 MB), but that still only explains ~40% of the operator's live heap (70 MB) and ~20% of its RSS.
- The remaining gap (~45 MB heap, ~170 MB RSS) comes from things the dummy does not do: implicit informers from cached Get/List calls in reconcilers and health checks, reconcile-time allocations (charts, rendering, image vector), webhook server/cert management, the virtual-garden client map, and an unbounded GC (no GOMEMLIMIT).
- Not yet covered: kinds the operator reads through the cached client without an explicit watch (unknown set; needs a heap profile or informer listing).

## Update 2: heap profile of gardener-operator

Profiling was enabled via `config.debugging.enableProfiling: 'true'` in `setValueTemplates` of `dev-setup/skaffold-operator.yaml` (dev-setup only, chart default unchanged) and redeployed. Profile taken ~1-3 min after start via the API server proxy: `kubectl get --raw /api/v1/namespaces/garden/pods/<pod>:2752/proxy/debug/pprof/heap`.

Metrics at that time: RSS ~273 MB, heap alloc ~121 MB, heap in use ~143 MB (higher than before the restart, because the initial Garden/Extension reconciliations just ran).

### Live heap (`inuse_space`, 73.7 MB in profile)

1. **CustomResourceDefinition informer: ~21 MB (28.5%).** `JSONSchemaProps*.UnmarshalJSON` under the informer watch decoder. The runtime cluster has 96 CRDs = ~19 MB of JSON (big OpenAPI schemas), all cached in full. Cause: `kubernetesutils.WaitUntilCRDManifestsReady` (used by `pkg/component/crddeployer` and the resource-manager component) does `client.Get` on `CustomResourceDefinition` with the manager's cached client, which implicitly starts a cluster-wide CRD informer. This one was invisible in the "Starting EventSource" list.
2. Secrets (~3 MB) and ConfigMaps (~3.5 MB) informers: ~6-7 MB combined, plus ObjectMeta etc. Matches the dummy.
3. Rest: REST mapper/discovery, scheme, JSON type caches (`typeFields` ~5 MB), zap, etc. (static costs).

### Allocations (`alloc_space`, 1450 MB total since start) – explains the high RSS/Sys

1. **`brotli` ring buffers: ~533 MB (37% of all allocated bytes).** All from `managedresources.(*Registry).SerializedObjects` / `AddAllAndSerialize` (compressing the ManagedResource secret data), called from `component.(*deploy).Deploy` in the Garden and Extension reconcilers (`controllerregistration`, `admission`, ...). Each compression allocates a fresh multi-MB brotli encoder buffer that is not reused/pooled. This is short-lived garbage but, with no `GOMEMLIMIT`, it drives the heap target up and explains `go_memstats_sys_bytes` (~367 MB) and RSS far above live heap.
2. **`JSONSchemaProps.DeepCopy`: ~135 MB.** Every cached `Get` of a CRD deep-copies the whole schema (cache returns copies), repeated by the 1s polling in `WaitUntilCRDManifestsReady` for all CRDs. Same root cause as live-heap item 1.
3. YAML (`yaml.v2/v4` parser/emitter, ~150 MB) and `json.Marshal` (~105 MB) – manifest rendering and serialization.

### Conclusion

The ~100 MB gap to the dummy is mostly (a) the unrestricted cluster-wide CRD informer (~21 MB live, plus large copy churn), and (b) transient allocation churn from brotli compression and YAML/JSON rendering in the reconcilers, inflated by an unbounded GC.

### Suggested fixes to evaluate

- Read CRDs through the uncached API reader (or `Cache.DisableFor: CustomResourceDefinition`, or a metadata-only/`PartialObjectMetadata` informer); `WaitUntilCRDManifestsReady` only needs status conditions.
- Reuse brotli writers (`sync.Pool` + `Reset`) or lower the compression buffer/quality in `pkg/utils/managedresources` (check `brotli.WriterOptions` `LGWin`) to avoid ~multi-MB allocations per call.
- Set `GOMEMLIMIT` (e.g. via memory limit/requests in the operator chart).
- Restrict Secret/ConfigMap caches by namespace/label or strip managedFields with a transform.
- Verify the effect by re-running the profile after each change.

Note: the `make gardener-up` run completed `operator-up` and `garden-up`; `seed-up` failed in my sandbox on missing docker buildx (I had to override `DOCKER_CONFIG`), unrelated to the operator.

## Update 3: CRD optimization

Changes:
- `cmd/gardener-operator/app/app.go`: manager `Client.Cache.DisableFor` now includes `CustomResourceDefinition`, so no CRD informer is started and CRD reads go to the API server.
- `pkg/utils/kubernetes/crd.go`: new `WaitUntilCRDManifestsReadyWithWatch(ctx, client.WithWatch, names...)`. Per CRD it lists with a `metadata.name` field selector, then watches from the returned resourceVersion (re-watching if the server closes the watch) instead of polling every second. Needs an uncached `client.WithWatch`. It is **not wired in yet** (`crddeployer` and the resource-manager component still use the polling `WaitUntilCRDManifestsReady`, because their clients are plain `client.Client`); unit tests added in `crd_test.go`.

Measurement after redeploy (profiling on, ~3 min after start, comparable to Update 2):

| | before | after |
|---|---|---|
| Live heap in profile (`inuse_space`) | 73.7 MB | 49.7 MB |
| Heap in use (metrics) | ~143 MB | ~71 MB |
| Heap alloc (metrics) | ~121 MB | ~47 MB |
| Go sys | ~367 MB | ~307 MB |
| RSS | ~273 MB | ~194 MB |
| `JSONSchemaProps.DeepCopy` (alloc_space) | ~135 MB | ~46 MB |
| brotli ring buffers (alloc_space) | ~533 MB | ~542 MB (unchanged) |

The CRD informer (~21 MB live) is gone. Single snapshots, so there is some noise. Remaining big items: brotli allocation churn in `managedresources.Registry` and YAML/JSON rendering; the remaining CRD DeepCopy comes from the `Get` in `Deploy` and the polling reads, which wiring in the watch-based wait should reduce further.

## Update 4: watch-based CRD wait wired into the operator

- `crddeployer.Wait` now uses `WaitUntilCRDManifestsReadyWithWatch` if its client implements `client.WithWatch`, otherwise it polls as before (gardenlet is unchanged for now; the polling path still works there).
- The Garden reconciler got a `CRDClient client.WithWatch` field, created in `AddToManager` via `client.NewWithWatch` (uncached, mgr scheme/REST mapper), and passes it to all 10 CRD deployers in `components.go`. As a side effect, `Deploy`/`Destroy` for CRDs use this uncached client too.
- Unit tests in `crddeployer` register a `metadata.name` index for the fake client (needed for field selectors); unit tests for `pkg/component/...`, `pkg/gardenlet/controller/seed`, `pkg/gardenlet/operation/botanist` and `pkg/operator` pass. Integration (envtest) tests were not run.

Result after redeploy (Garden stays `Succeeded`, no CRD-related errors in the operator log):

| | Update 2 (baseline) | Update 3 | Update 4 |
|---|---|---|---|
| Live heap (`inuse_space`) | 73.7 MB | 49.7 MB | 43.3 MB |
| Heap alloc (metrics) | ~121 MB | ~47 MB | ~66 MB |
| Heap in use (metrics) | ~143 MB | ~71 MB | ~83 MB |
| RSS | ~273 MB | ~194 MB | ~166 MB |
| `JSONSchemaProps.DeepCopy` (alloc) | ~135 MB | ~46 MB | ~40 MB |

Single snapshots; the heap numbers fluctuate with GC timing, so only the profile's live heap and RSS are a reasonably stable comparison. The remaining `DeepCopy` comes from `GetAndCreateOrMergePatch` in `Deploy` (it deep-copies the object for building the patch), not from the wait. Brotli compression churn is the next big item.

## Update 5: same CRD change for gardenlet

- `kubernetesutils.NewUncachedWatchClient(restConfig, client)` helper (uncached `client.WithWatch` reusing scheme and REST mapper).
- Seed reconciler (`pkg/gardenlet/controller/seed/seed`): new `CRDClient` field, created in `AddToManager` from the manager config; used for all CRD deployers in `components.go` (falls back to the seed client if unset).
- Shoot flow (`pkg/gardenlet/operation/botanist/flow.go`, self-hosted shoots only): `ReconcileCustomResourceDefinitionsTaskGroup` lazily creates one such client (`sync.OnceValues`) and passes it to the CRD deployers.
- Not changed: the gardenlet seed manager still has no `DisableFor` for CRDs, so other cached CRD reads (if any) would still start an informer. Not verified, see below.
- Unit tests for the changed packages pass. Gardenlet redeployed with `make seed-up`: seed `local` is `Ready`/`Reconcile Succeeded`, no error logs.
- Measurement: no before/after comparison, because gardenlet profiling is disabled and the old pod was replaced before I sampled it. Current snapshot (new pod, ~2 min): RSS ~195 MB, heap alloc ~91 MB, heap in use ~100 MB, 886 goroutines.

## Update 6: gardenlet before/after (properly profiled)

Method: gardenlet profiling enabled temporarily (`debugging.enableProfiling: true` in `dev-setup/gardenlet/base/gardenlet.yaml`, since reverted), profile via API server proxy on the metrics port 2729, taken ~4 min after pod start. "Before" = gardenlet without the watch-client wiring (seed client cached, CRD wait polls); "after" = with the wiring from Update 5. Both used the same operator build and same cluster.

| | before | after |
|---|---|---|
| Live heap in profile (`inuse_space`) | 68.7 MB | 55.1 MB |
| CRD data in live heap (`JSONSchemaProps` unmarshal) | 17.5 MB (25%), under the informer's `StreamWatcher`/reflector (37 traces) | 10 MB, no informer/reflector, only transient in-flight `client.List` / `DecodeCRD` during a reconcile (0 reflector traces) |
| `JSONSchemaProps.DeepCopy` (alloc_space) | ~5 MB | ~0.5 MB |
| Total allocated (alloc_space) | 1151 MB | 1067 MB |
| Heap alloc (metrics) | ~69 MB | ~82 MB |
| Heap in use (metrics) | ~104 MB | ~97 MB |
| Go sys (metrics) | ~269 MB | ~193 MB |
| RSS (metrics) | ~246 MB | ~228 MB |
| Goroutines | 895 | 886 |

Takeaways:
- The CRD informer in the gardenlet is gone (no reflector/watch decoding of CRDs any more); the live heap in the profile drops by ~14 MB (-20%) and Go sys by ~75 MB.
- RSS improves less (~-18 MB) and heap alloc/in-use from the metrics are noisy (single snapshots, GC timing, in-flight reconcile), so treat those as roughly equal. The remaining ~10 MB of CRD data in the "after" profile is the transient result of reading/decoding CRDs during a seed reconcile, not a cache.
- The gardenlet still shows the same non-CRD costs as the operator (brotli ring buffers in `managedresources.Registry`, YAML/JSON rendering); the next optimization target.
- Seed `local` stayed `Ready`/`Reconcile Succeeded`; 0 error logs in the new gardenlet.
