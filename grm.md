# gardener-resource-manager: target cache memory findings

Scope: memory used by the target-cluster cache of `gardener-resource-manager` (GRM), with focus on the shoot GRM
(up to ~120 instances per seed, so per-instance savings multiply).

Status: derived from code reading and the local setup (runtime cluster, managed seed shoot `shoot--garden--managedseed`).
**No heap profile of a GRM has been taken yet**; the numbers below are `kubectl top` working sets, not live heap.

## Background

- The health controller (`pkg/resourcemanager/controller/health/health`) watches **every GroupVersionKind** that appears in
  `status.resources` of a ManagedResource, via `source.Kind(targetCluster.GetCache(), ...)` (`add.go`, `ensureWatchForGVK`).
  This is one informer per kind, covering all objects of that kind visible to the target cache.
- If the kind is registered in the target scheme (`pkg/resourcemanager/client/scheme.go`), the informer is **typed and caches full
  objects**. Only unregistered kinds fall back to `PartialObjectMetadata` (`newObjectForHealthCheck`).
- `CheckHealth` (`health/utils/health_checker.go`) has dedicated checks only for: CustomResourceDefinition (v1, v1beta1), DaemonSet,
  Deployment, Job, Pod, ReplicaSet, ReplicationController, Service, StatefulSet, Prometheus, Alertmanager,
  VerticalPodAutoscaler, Certificate, Issuer. For every other kind, the controller only cares whether the object exists.
- Shoot and virtual-garden GRMs use a separate target cache restricted to `TargetNamespaces` (shoot: `kube-system`,
  `kubernetes-dashboard`, `kube-node-lease`); this restriction does **not** apply to cluster-scoped kinds.
- Source and target caches are separate in these GRMs, so changes to the target cache do not affect the source side that
  reads the ManagedResource Secrets.

Observed ManagedResource contents of the managed seed shoot (kinds x count, objects referenced): ClusterRoleBinding 40,
CustomResourceDefinition 26, ClusterRole 25, ConfigMap 20, ServiceAccount 18, Secret 17, Service 15, VPA 14, Deployment 12,
RoleBinding 12, NetworkPolicy 10, PrometheusRule 10, ...
Working set of its GRM pods: ~153Mi (one replica) vs ~45Mi (other replica). Informers only run in the leader, so the gap
is probably cache cost. Not verified.

## Findings, ordered by priority

### 1. Use metadata-only health watches for kinds without a dedicated health check (high value, low risk)

- Kinds like Secret, ConfigMap, ServiceAccount, ClusterRole(Binding), Role(Binding), NetworkPolicy, PriorityClass, PDB,
  HPA, StorageClass, APIService, webhook configurations, Namespace, and Istio/monitoring kinds that are registered in the
  scheme are watched with typed informers, though the health result is only "exists / does not exist".
- Change idea: in `newObjectForHealthCheck`, return `PartialObjectMetadata` also for scheme-registered kinds that `CheckHealth`
  does not handle (for example by making `CheckHealth`'s set of handled types queryable from one place).
- Why safe: the reconciler already handles metadata-only objects (`TargetClient.Get` plus the existing watch path), and
  `HealthStatusChanged` already works for them.
- Expected effect: removes full-object caching for the majority of kinds, in particular Secrets and ConfigMaps in `kube-system`
  and cluster-scoped RBAC objects.
- Caveats to verify:
  - Other controllers may still create a typed informer for the same kind in the same target cache (for example
    `agentreconciliationdelay` lists Secrets in `kube-system`; node controllers read Node, DaemonSet, Namespace). A typed and a
    metadata informer for the same kind are separate informers, so the saving only materialises for kinds not used elsewhere.
    This needs a per-controller check of which target-cache kinds each enabled controller uses in the shoot GRM.
  - Progressing controller (`progressing/add.go`) adds its own watches; check overlap.

### 2. Strip `spec` of CustomResourceDefinitions in the target cache (medium value, low-medium risk)

- Health only needs `metadata` (annotation `resources.gardener.cloud/skip-health-check`) and `status.conditions`
  (`NamesAccepted`, `Established`, `Terminating`), see `pkg/utils/kubernetes/health/customresourcedefinition.go`.
- Metadata-only does not work (conditions are in `status`), so use a cache transform on the target cluster
  (`opts.Cache.ByObject[&CustomResourceDefinition{}].Transform`) that clears `Spec` (at least `Spec.Versions[*].Schema`).
- The informer is cluster-wide because CRDs are cluster-scoped. The local managed seed shoot has 26 CRDs in its ManagedResources;
  in this setup the same CRDs total only ~0.44 MB of JSON in the runtime cluster, so the saving per instance is small
  locally but grows with the number of CRDs deployed via ManagedResources (and the cluster's total CRDs only matter if the
  informer also receives events for all of them, which it does).
- Multiplied by up to ~120 GRMs per seed, even a small per-instance saving is meaningful.
- Risks:
  - Any typed `Update`/`Patch` built from a cached CRD would wipe the spec. No such code path was found in `pkg/resourcemanager`
    (the ManagedResource apply path uses unstructured objects, not read in detail), but add a comment at the transform.
  - The CRD deletion-protection webhook reads from `SourceReader` and decodes the admission object, so it is unaffected.
  - Decoding garbage is still transient; only live heap shrinks.

### 3. Optionally strip large fields of other kinds that have status-based checks (low-medium value, higher risk)

- Deployment, DaemonSet, StatefulSet, Job, Pod, Service, VPA, Prometheus, etc. need `status` and parts of `spec`
  (replicas, selectors, update strategy) and are used by the progressing controller, so only
  `managedFields` stripping is clearly safe here (`metadata.managedFields` via a generic transform).
- Not recommended to strip `spec.template` and similar without per-check review.

### 4. Runtime/seed GRM: cluster-wide typed caches (separate problem, no cheap fix)

- For responsibility mode `ForRuntimeCluster`, `TargetClientConnection` is unset and the target cluster is the manager itself
  (`app.go:176`), so the cache is **not** restricted to `TargetNamespaces`, and typed kinds are cluster-wide
  (Secrets, ConfigMaps, Deployments, ...).
- Source and target share one cache, and the ManagedResource controller reads referenced Secrets (compressed manifests)
  from it, so Secret data **cannot** be stripped there. Items 1 and 2 give only limited savings in this mode because the
  typed informers are needed anyway.
- Possible later work: `ByObject` namespace/label restrictions, again after profiling.

### 5. Things that do not help (documented to avoid revisiting)

- `Cache.DisableFor` for kinds in the GRM target client: the health watch uses the cache directly via `source.Kind`, so the
  informer would still run, and every health check would add a live API call.
- Stripping Secret `data` in the **source** cache: it holds the ManagedResource payload used by the actual reconcile; this
  would force live reads or break reconciles.

## Suggested next steps

1. Take heap profiles (`inuse_space`) of a shoot GRM leader (managed seed shoot) and of the runtime GRM, to confirm which typed
   kinds dominate. Enable profiling via the GRM config (`debugging.enableProfiling`).
2. List the informers actually started in the shoot GRM (log lines `Adding new watch for GroupVersionKind` with `metadataOnly`).
3. Implement item 1, re-profile, then item 2, re-profile.
4. Test on a seed with many shoots to validate the per-seed total (single-snapshot measurements on a small local cluster are
   noisy).
