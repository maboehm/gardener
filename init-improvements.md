# init improvements

Context: experiment with a single `gardener <component>` binary (see stash "single binary"). In a single binary every
component runs the `init()` functions and package-level initialisers of *all* linked packages.

## How to measure

```
go build -o /tmp/gb/gardener ./cmd/gardener
GODEBUG=inittrace=1 /tmp/gb/gardener scheduler --help 2>&1 >/dev/null | awk '/^init /{print $(NF-3), $(NF-1), $5, $2}' | sort -rn | head -30
GODEBUG=gctrace=1   /tmp/gb/gardener scheduler --help 2>&1 >/dev/null | grep '^gc '
```

`inittrace` reports *allocated* bytes (mostly garbage later), not retained memory. Use `gctrace` ("X->Y->Z MB", Z = live)
for what stays.

## Findings so far

| Binary | init allocs | inits | live heap after startup GCs |
|---|---|---|---|
| gardener-scheduler (separate) | 11 MB / 82k allocs | 671 | ~3-5 MB |
| gardener scheduler (single) | 53 MB / 520k allocs | 1014 | ~10 MB (peak ~19 MB) |
| gardener-operator (separate) | 49 MB / 499k allocs | 827 | - |

Every component in the single binary pays roughly the operator's init cost. Garbage is collected, but the larger heap
target and the slow scavenger can keep RSS / Pss_Anon higher for a while.

Measured with `hack/measure-memory.sh` (8 like-for-like processes): Pss_File -259 MB (sharing works), Pss_Anon +213 MB
(uptime differed between runs, so not conclusive). Net -46 MB (-5%).

## Top init allocators (single binary)

| Package | Alloc | Allocs | Time | Cause |
|---|---|---|---|---|
| `pkg/component/observability/monitoring/perses` | 23.1 MB | 337k | 31-38 ms | `init()` in `dashboards.go` decodes + schema-validates every embedded dashboard ("fail fast") |
| `github.com/gobuffalo/flect` | 4.7 MB | 2k | 2 ms | 3rd party, pulled in via `kubestatemetrics` -> `kube-state-metrics/pkg/customresourcestate` |
| `pkg/resourcemanager/client` | 2.4 MB | 11k | 2-6 ms | `init()` builds source/target schemes |
| `pkg/client/kubernetes` | 2.4 MB | 11k | 3 ms | `init()`s: protobuf scheme registration (`client.go`), Garden/Seed/Shoot schemes (`types.go`) |
| `santhosh-tekuri/jsonschema/v6` | 1.7 MB | 22k | 2-3 ms | 3rd party (perses validation) |
| `pkg/operator/client` | 1.6 MB | 7k | 1.5 ms | `init()` builds Runtime/Virtual schemes |
| `imagevector` | 1.1 MB | 15k | 1-4 ms | `init()` parses `containers.yaml` + `charts.yaml` + env override |
| `prometheus/garden` | 0.7 MB | 8k | 1 ms | embedded rules decoded at init |
| `prometheus/cache` | 0.6 MB | 7k | 1.6 ms | `init()` decodes 4 PrometheusRules |
| `prometheus/shoot`, `prometheus/aggregate` | 0.4 + 0.2 MB | 7k + 2k | - | same pattern as above |
| `k8s.io/client-go/kubernetes/fake` | 0.5 MB | 2.7k | 0.9 ms | fake clientset linked into an operator/gardenlet binary; non-test importer not identified yet (`pkg/gardenadm/admission` imports it, but is not part of the single binary) |

## Candidate improvements

1. **perses: move the fail-fast validation into a unit test** (biggest win, ~43% of init allocations).
   - The embedded dashboards are static, so a test catches everything the `init()` does.
   - Test must cover all 5 embed.FS sets: decode ok, `IsDNS1123Subdomain(name)`, `len(config.Panels) > 0`.
   - Add a duplicate name check across sets (`loadDashboards` silently overwrites on collision).
   - Runtime errors are still reported by `loadDashboards` -> `decodeDashboard`.
2. **prometheus/{cache,garden,shoot,aggregate}: decode rules lazily** (`sync.OnceValue` or decode in the getter) instead
   of in `init()`. Same pattern as perses: static input, validated by unit tests.
3. **imagevector: decode lazily** (`sync.OnceValues`). Needs care: today a broken env override fails at startup; keep an
   explicit check in the commands that use it, or call it early in `main`.
4. **Scheme registration (`pkg/client/kubernetes`, `pkg/resourcemanager/client`, `pkg/operator/client`)**: schemes are
   used by almost every component, so lazy init yields less. Only worthwhile if profiling shows a measurable retained
   heap. Low priority.
5. **flect via kube-state-metrics**: check whether `kubestatemetrics` really needs `customresourcestate` (only
   operator/gardenlet need the package). Moot once components are not mixed into one binary.
6. **kubernetes/fake in production binaries**: find the importer (`go list -deps -f` / `go mod why`) and remove it.

## Open questions

- Is the Pss_Anon increase (+213 MB over 8 processes) caused by init garbage / larger heap targets, or just by uptime
  differences between the two measurements? Re-measure both states at identical uptime (>= 15 min, all pods restarted
  incl. the shoot GRM and gardenlet).
- Alternative to fixing every init: `GOMEMLIMIT` / `debug.FreeOSMemory()` after init in the single binary's main.
- Generic option for the single binary: lazy component registration so only the selected component's packages are
  linked is impossible (one binary), but inits of unused packages could be made cheap by following 1-3.
