# Single `gardener` binary experiment: findings

Experiment: replace the separate component binaries/images with one `gardener <component>` binary, to see whether
sharing the binary between processes saves memory (and disk) in the local setup and, by extension, on real clusters.

In scope: `gardener-operator`, `gardener-resource-manager`, `gardener-apiserver`, `gardener-controller-manager`,
`gardener-scheduler`, `gardener-admission-controller`, `gardenlet`. Not in scope: `gardenadm`, `gardener-node-agent`, all
`-local` binaries.

## 1. Binary size

| Build                          | Size                     |
| ------------------------------ | ------------------------ |
| Single `gardener` binary       | 213,116,418 B (~203 MiB) |
| Sum of the 7 separate binaries | 963,445,998 B (~919 MiB) |

The single binary is ~22% of the separate total (-750 MB). It is only ~40 MB larger than the biggest binary (gardenlet,
172.6 MB). Default `go build`, debug info included.

## 2. How the experiment was wired

- `cmd/gardener/main.go` dispatches on `os.Args[1]` (`operator`, `apiserver`, `controller-manager`, `scheduler`,
  `admission-controller`, `resource-manager`, `gardenlet`), registers that component's feature gates, and executes its
  cobra command with `os.Args[2:]`.
- Local setup: the ko artifacts for these 7 images in `dev-setup/skaffold-operator.yaml` and
  `dev-setup/skaffold-seed.yaml` were replaced by a skaffold custom builder, `hack/hack-build-gardener.sh <component>`.
  It builds `./cmd/gardener` with ko (identical layers for every component) and only changes the image entrypoint with
  `crane mutate --entrypoint /ko-app/gardener,<component>`.
- Pods set only `args`, so the entrypoint is the only thing that had to change in the deployments.
- Version flags (`k8s.io/component-base/version.*`) have to be injected via a generated `.ko.yaml`, otherwise version
  checks break. The build date is pinned in the script so that operator and seed builds produce byte-identical binaries;
  otherwise the layers differ and nothing can be shared.
- Measurement: `hack/measure-memory.sh [label]` reads `/proc/<pid>/smaps_rollup` for every `/ko-app/*` process in the
  kind node and prints Rss, Pss, Pss_Anon, Pss_File plus VM `meminfo`.

## 3. What the measurements mean

| Value                | Meaning                                                                                                    |
| -------------------- | ---------------------------------------------------------------------------------------------------------- |
| Rss                  | all resident memory of the process; shared pages are counted fully for each process                        |
| Pss                  | resident memory with shared pages divided by the number of sharers; sums are meaningful                    |
| Pss_Anon             | private anonymous memory (heap, stacks); cannot be shared                                                  |
| Pss_File             | file-backed memory (mostly the binary); this is what sharing reduces                                       |
| `used (excl. cache)` | `MemTotal - MemFree - Buffers - Cached - SReclaimable` of the Docker VM; excludes page cache by definition |

`kubectl top` reports the cgroup working set (`memory.current - inactive_file`). Shared pages are charged to the cgroup
that touched them first, and inactive file pages are subtracted, so file-backed sharing is mostly invisible there. This
is why the `kubectl top` comparison (below) showed nothing.

## 4. Results

### `kubectl top` (3 samples each, 10 target pods)

|                    | Before                | After                 |
| ------------------ | --------------------- | --------------------- |
| Node memory        | 9171 / 9176 / 9333 Mi | 9376 / 9368 / 9341 Mi |
| Sum of target pods | 1010 / 1010 / 1087 Mi | 1064 / 1060 / 1058 Mi |

No saving visible, as expected (see above).

### `smaps_rollup`, ~10 minutes after rollout (`memory-*-10m.txt`)

Two processes (gardenlet pid 707499 and one resource-manager pid 713265) were still single-binary in the "separate"
report (not restarted); they are excluded below. 8 comparable processes, kB:

|          | Separate | Single  | Difference      |
| -------- | -------- | ------- | --------------- |
| Pss      | 867,258  | 671,248 | -196,000 (-23%) |
| Pss_File | 365,114  | 88,888  | -276,000 (-76%) |
| Pss_Anon | 502,144  | 582,360 | +80,000 (+16%)  |

Unfiltered (10 processes): SUM Pss 1,235,434 vs 929,467 (-25%), Pss_File 468,946 vs 117,031, Pss_Anon 766,488 vs 812,436.

Per component anon (separate -> single): operator +3 MB, apiserver +31 MB, admission-controller +8 MB,
controller-manager +7 MB, scheduler +6 MB, resource-manager replicas +4 to +10 MB each.

Node-wide (Docker VM): used excl. cache 10,132,568 kB (separate) vs 10,297,852 kB (single), Cached about equal. The
reports were taken ~16 minutes apart, so other pods drifted too. The file saving does not appear because it is page
cache, which the kernel fills with other data anyway.

An earlier measurement (shortly after rollout, different uptimes) showed Pss_Anon +213 MB; that was largely a timing
artefact.

## 5. Interpretation

- Page sharing works: with one binary, the ~80 MB of touched binary pages per process are mapped once and charged
  ~1/N each. The ~276 MB Pss_File saving for 8 processes is real.
- It is soft memory: page cache that the kernel can reclaim. The hard memory (anon) goes up by about 5-10 MB per process
  (+31 MB for the apiserver), because every process now runs the init code of all components.
- The net gain is about 5% of the Gardener components in the first comparison and ~23% in Pss terms after 10 minutes, but
  roughly 0 for node-wide "used excl. cache". Gardener's components are about 1 GB of ~10 GB used on a minimal local
  runtime, so the effect on Docker's memory needs is marginal.

## 6. Init cost (see `init-improvements.md` for details)

`GODEBUG=inittrace=1`: the scheduler allocates 11 MB in package inits when separate and 53 MB in the single binary
(1014 inits vs 671). The operator already pays ~49 MB. `gctrace` shows the single-binary scheduler holds ~10 MB live
heap after startup vs ~3-5 MB (peak ~19 MB), consistent with the measured +5-10 MB anon per process. Allocations are
garbage-collected, but larger heap targets and slow scavenging keep RSS up.

Top init allocators: `pkg/component/observability/monitoring/perses` (23 MB, 337k allocs; decodes and validates all
embedded dashboards in `init()`), `gobuffalo/flect` (4.7 MB, 3rd party via kube-state-metrics),
`pkg/resourcemanager/client`, `pkg/client/kubernetes`, `jsonschema/v6`, `pkg/operator/client`, `imagevector`, and the
`prometheus/{cache,garden,shoot,aggregate}` rule decoding.

Candidate fixes: move the perses fail-fast validation into a unit test (incl. DNS-1123 name check, non-empty panels,
duplicate names); decode embedded prometheus rules and the imagevector lazily (`sync.OnceValue`); investigate why
`k8s.io/client-go/kubernetes/fake` is linked into the operator/gardenlet binaries; look at the apiserver's larger anon
increase.

## 7. Dev-loop impact

A single binary is a regression for the local dev loop. Today each component has its own artifact with its own
dependency list, so a change to the scheduler rebuilds and redeploys only the scheduler. With one binary any change
rebuilds a ~200 MB binary, pushes a new layer, and rolls all components (including the apiserver) because they share
one image digest.

Options: keep local builds separate and ship the single binary only for releases (CI should then test the shipped
artifact); or offer the single binary as an opt-in skaffold profile; or drop the idea.

## 8. Cost of shipping a single binary

Every deployment needs the subcommand as entrypoint/args (helm charts and the Go code in operator/gardenlet that builds
deployments), the Dockerfile targets and image vector entries change from one image per component to one image with
different commands, and release tooling/docs need updates.

## 9. Conclusion

| Goal                        | Does the single binary help?                                                             | Better lever                                                     |
| --------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------- |
| Less Docker RAM (local dev) | marginally (soft page cache, anon slightly up)                                           | fewer replicas/components, `GOMEMLIMIT`, smaller informer caches |
| Smaller nodes (operator)    | only if requests are lowered, which is risky (shared pages are charged to the first pod) | lower anon memory (init fixes, caches), VPA-tuned requests       |
| Disk / image pull size      | yes: one ~200 MB image instead of ~920 MB                                                | n/a                                                              |

The most valuable follow-up regardless of the packaging decision is the init-cost work.

## 10. Reproduction and cleanup

- Setup: `make gardener-up` with the skaffold changes in place; wait ~15 minutes; make sure the shoot resource-manager
  and gardenlet pods were restarted (check image digests); run `hack/measure-memory.sh single`. Repeat with the changes
  stashed (without `-u`) for `separate`.
- Experiment files: `cmd/gardener/`, `hack/hack-build-gardener.sh`, `hack/measure-memory.sh`,
  `dev-setup/seed-tmp.sh` (sandbox workaround, not needed on a normal shell), modified
  `dev-setup/skaffold-operator.yaml` and `dev-setup/skaffold-seed.yaml`, `memory-*.txt`.
- Revert: `git checkout dev-setup/skaffold-*.yaml`, delete the files above, then `make gardener-up`.
