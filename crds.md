# Optimized CRD Handling in gardener-operator and gardenlet

## Problem

Heap profiling of `gardener-operator` (local KinD runtime cluster, 96 CRDs, ~19 MB of JSON) showed that the `CustomResourceDefinition` informer was the largest single consumer of live heap:

- ~21 MB (28.5% of the 73.7 MB live heap in the profile) was CRD data (`JSONSchemaProps.UnmarshalJSON` under the informer watch decoder).
- `JSONSchemaProps.DeepCopy` accounted for ~135 MB of allocations, because every cached `Get` deep-copies the full schema, and the 1s polling in `WaitUntilCRDManifestsReady` repeats this for all CRDs.

Root cause: `kubernetesutils.WaitUntilCRDManifestsReady` (used by `pkg/component/crddeployer` and the resource-manager component) calls `client.Get` on CRDs through the manager's cached client. This implicitly starts a cluster-wide CRD informer that caches all CRDs in full. It did not show up in the "Starting EventSource" log because it is an implicit informer.

The same pattern exists in gardenlet (seed reconciler and self-hosted shoot flow).

## Changes

1. **`cmd/gardener-operator/app/app.go`**: manager `Client.Cache.DisableFor` now includes `CustomResourceDefinition`. No CRD informer is started; reads go to the API server.
2. **`pkg/utils/kubernetes/crd.go`**: new `WaitUntilCRDManifestsReadyWithWatch(ctx, client.WithWatch, names...)`. Per CRD, it lists with a `metadata.name` field selector and then watches from the returned resourceVersion (re-watching when the server closes the watch) instead of polling every second. It requires an uncached `client.WithWatch`. Unit tests are in `crd_test.go`.
3. **`pkg/component/crddeployer`**: `Wait` uses the watch-based variant if its client implements `client.WithWatch`, otherwise it falls back to polling. Unit tests register a `metadata.name` index for the fake client (needed for field selectors).
4. **Operator Garden reconciler**: new `CRDClient client.WithWatch` field, created in `AddToManager` via `client.NewWithWatch` (uncached, manager scheme/REST mapper) and passed to all 10 CRD deployers in `components.go`. Side effect: `Deploy`/`Destroy` of CRDs also use the uncached client.
5. **gardenlet**:
   - New helper `kubernetesutils.NewUncachedWatchClient(restConfig, client)`.
   - Seed reconciler (`pkg/gardenlet/controller/seed/seed`): new `CRDClient` field created in `AddToManager`, used for all CRD deployers in `components.go` (falls back to the seed client if unset).
   - Shoot flow (`pkg/gardenlet/operation/botanist/flow.go`, self-hosted shoots only): `ReconcileCustomResourceDefinitionsTaskGroup` lazily creates one such client (`sync.OnceValues`) and passes it to the CRD deployers.

## Results

Measured on a local dev cluster, ~3-4 min after pod start, with profiling enabled. Single snapshots: the live heap in the profile and RSS are the most reliable numbers; heap alloc/in-use from metrics fluctuate with GC timing.

### gardener-operator

|                                    | Before  | CRD informer removed | + watch-based wait |
| ---------------------------------- | ------- | -------------------- | ------------------ |
| Live heap (`inuse_space`)          | 73.7 MB | 49.7 MB              | 43.3 MB            |
| RSS                                | ~273 MB | ~194 MB              | ~166 MB            |
| `JSONSchemaProps.DeepCopy` (alloc) | ~135 MB | ~46 MB               | ~40 MB             |

Live heap is down ~41% and RSS ~39%. The remaining `DeepCopy` comes from `GetAndCreateOrMergePatch` in `Deploy`, not from the wait.

### gardenlet

|                                    | Before                       | After                                |
| ---------------------------------- | ---------------------------- | ------------------------------------ |
| Live heap (`inuse_space`)          | 68.7 MB                      | 55.1 MB                              |
| CRD data in live heap              | 17.5 MB (informer/reflector) | 10 MB (transient only, no reflector) |
| `JSONSchemaProps.DeepCopy` (alloc) | ~5 MB                        | ~0.5 MB                              |
| Go sys                             | ~269 MB                      | ~193 MB                              |
| RSS                                | ~246 MB                      | ~228 MB                              |

Live heap is down ~20% and Go sys by ~75 MB. The RSS improvement is smaller and the metrics-based heap numbers are noisy. The remaining ~10 MB of CRD data is transient decoding during a seed reconcile, not a cache.

## Verification

- Garden stays `Succeeded`; seed `local` stays `Ready` / `Reconcile Succeeded`; no CRD-related errors in the operator or gardenlet logs.
- Unit tests pass for `pkg/component/...`, `pkg/gardenlet/controller/seed`, `pkg/gardenlet/operation/botanist`, `pkg/operator`.
- **Integration (envtest) tests were not run.**

## Caveats and open points

- Measurements are single snapshots on a small dev cluster; benefits should grow with more CRDs/real clusters.
- The gardenlet seed manager still has no `DisableFor` for CRDs. Other cached CRD reads, if any, would still start an informer (not verified).
- Not part of this change but identified as next targets:
  - brotli ring buffer allocations in `managedresources.Registry` (~533 MB of ~1450 MB allocated, unchanged by this PR). Consider pooling writers or lowering the window size.
  - YAML/JSON rendering churn.
  - Missing `GOMEMLIMIT` / resource limits on the operator.
  - Restricting Secret/ConfigMap caches.
