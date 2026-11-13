# Brotli compression optimization in `managedresources.Registry`

## Problem

`gardener-operator` heap profiles (see `findings.md`) showed brotli encoder allocations as the largest source of
allocation churn: ~533 MB (37% of all allocated bytes) within the first minutes after start. All of it comes from
`managedresources.(*Registry).SerializedObjects` (called via `AddAllAndSerialize` from `component.(*deploy).Deploy` in the
Garden and Extension reconcilers). Each call created a new `brotli.NewWriter`, which allocates a multi-MB ring buffer
(`ringBufferInitBuffer`) and hasher tables (`h5.Initialize`) sized for the default window (`lgwin=22`, 4 MiB). This is
short-lived garbage, but it inflates the GC heap target and therefore Go sys/RSS. The same code path is used by
gardenlet.

## Change

`pkg/utils/managedresources/registry.go`:

- Writers come from a `sync.Pool` instead of being created per call. They are reset onto a fresh buffer per call and
  reset to `io.Discard` before returning to the pool, so the pool does not hold on to the output buffer.
- Writers are created with `brotli.WriterOptions{Quality: 6, LGWin: 20}` (1 MiB window). Quality 6 is the library
  default, so only the window changes.

`pkg/utils/test/compression.go`: `BrotliCompression` uses the same options (the comment says they must match), because
tests compare compressed bytes.

`pkg/component/nodemanagement/dependencywatchdog/access_test.go`: hardcoded secret name/annotation hashes updated
(`3cb55785` -> `f41d1a86`, `e3a35806` -> `ae60d80f`) because they are derived from the compressed data.

## Measurements

Method: `gardener-operator` redeployed fresh for each variant (`make operator-up`, local KinD runtime cluster, profiling
enabled), profile taken once the Garden had been `Succeeded` for >= 60s (pod age ~160s). One sample per variant
(two for the final one), so memory numbers are noisy.

| Variant                          | brotli alloc | total alloc | Go sys | RSS    | live heap (profile) |
| -------------------------------- | ------------ | ----------- | ------ | ------ | ------------------- |
| A: baseline (no pool, window 22) | 436 MB       | 1432 MB     | 370 MB | 162 MB | 38.9 MB             |
| B: pool only                     | 218 MB       | 1124 MB     | 345 MB | 276 MB | 97.6 MB             |
| C: pool + window 20              | 60 MB        | 945 MB      | 240 MB | 198 MB | 70.8 MB             |
| D: window 20, no pool            | 118 MB       | 1072 MB     | 198 MB | 168 MB | 44.5 MB             |
| C2: repeat of C                  | 62 MB        | 941 MB      | 210 MB | 206 MB | 65.8 MB             |

Takeaways:

- Brotli allocation drops by ~85% (436 -> ~60 MB), total allocation by ~35%, Go sys by ~40%.
- The smaller window gives most of the reduction; the pool adds roughly another 2x on top of it.
- RSS and live heap showed **no improvement** and are noisy (single snapshots, GC timing). Do not claim an RSS benefit.
  The improvement is in allocation churn and Go sys.

### Compression ratio

Recompressing the 117 real ManagedResource secrets of the cluster (1.9 MB raw, largest 418 KB):

|                      | window 22 | window 18          |
| -------------------- | --------- | ------------------ |
| total compressed     | 370,476 B | 370,516 B (+0.01%) |
| largest (418 KB raw) | 30,174 B  | 30,214 B (+0.13%)  |

Negligible. Caveat: only a small dev cluster was tested. A payload near the 1 MB Secret limit with long-range
repetition could lose slightly more ratio. Brotli decoders accept any window size, so readers are unaffected.

### CPU

Benchmark over the same 117 payloads (sequential, quality 6):

| Variant                       | time per pass        | allocated per pass |
| ----------------------------- | -------------------- | ------------------ |
| window 22, no pool (original) | 36.5-39.9 ms         | 308 MB             |
| window 18, no pool            | 35.3-49.3 ms (noisy) | 277 MB             |
| window 22, pool               | 21.3-21.4 ms         | 1.09 MB            |
| window 18, pool (final)       | 21.6-22.0 ms         | 0.7 MB             |

- The window size does not change CPU time measurably.
- The pool reduces time per compression by ~40% (no repeated allocation/zeroing of buffers and tables).
- In the benchmark the pool is warm and single-threaded. In the operator, GC empties the pool and compressions can run
  concurrently, so the real saving is smaller (brotli allocation fell 55-85% there, not ~99%). Process-level CPU
  (`process_cpu_seconds_total`) was not measured.

### Window size sweep

Same 117 real payloads (1.9 MB raw), quality 6, 2 runs each. "no pool" = new writer per call, "pool" = reused writer.

| lgwin        | window  | compressed total | vs lgwin 22    | alloc/pass (no pool) | alloc/pass (pool) | time/pass (no pool) | time/pass (pool) |
| ------------ | ------- | ---------------- | -------------- | -------------------- | ----------------- | ------------------- | ---------------- |
| 16           | 64 KiB  | 372,550 B        | +0.55%         | 87 MB                | 0.62 MB           | 36-55 ms            | 27.0 ms          |
| 18           | 256 KiB | 370,516 B        | +0.01%         | 277 MB               | 0.71 MB           | 38-39 ms            | 21.3-21.7 ms     |
| 20           | 1 MiB   | 370,476 B        | 0% (identical) | 283 MB               | 0.75-0.78 MB      | 37-45 ms            | 21.3-21.6 ms     |
| 22 (default) | 4 MiB   | 370,476 B        | baseline       | 308 MB               | 1.07-1.10 MB      | 37 ms               | 21.5-21.7 ms     |

Why the default is not optimal for us: brotli sizes its ring buffer (~`2^lgwin` + tail) and hasher tables from the
window, not from the input, because the encoder is a streaming one and cannot know the total size. Our inputs are
bounded (a Secret is capped at ~1 MiB; the largest payload here is 418 KB) and written to a fresh writer once, so a
4 MiB window is allocated and zeroed but never filled, and it gives no compression benefit.

Observations:

- Without a pool, the window alone only helps a little in the benchmark (308 -> 277 MB at lgwin 18, -10%), except
  at lgwin 16, where allocation drops to 87 MB. The jump between 18 and 16 is probably the encoder choosing a
  different (smaller) hasher at small windows; this was not verified. lgwin 16 costs +0.55% size and is slower
  with the pool (27 ms vs ~21 ms), so it is not recommended.
- With a pool, the allocation per pass is ~1 MB regardless of the window, and the window only changes it by
  0.4 MB. In the benchmark the **pool is the dominant factor**. In the operator the window mattered more (window 18 without
  pool: brotli 436 -> 118 MB), likely because the pool cannot catch every call there (GC clears it, concurrent
  reconciles), so the allocations that do miss the pool are smaller with a smaller window.
- lgwin 20 compresses byte-for-byte as small as the default on this data and still covers any payload that fits into
  a Secret, but it saves only ~8% allocation without a pool and ~30% with one, compared with lgwin 22. lgwin 18 costs +0.01%.
  lgwin 18 was chosen empirically; lgwin 20 is the more conservative alternative (still changes the header byte,
  so the upgrade impact below applies to it as well).

## Upgrade impact (important for the PR description)

Changing the window changes the brotli stream header byte, so the compressed bytes of every ManagedResource secret
change once. Secret names and `reference.resources.gardener.cloud/secret-*` annotations are derived from the data hash,
so all ManagedResource secrets are rewritten/recreated one time on the first reconcile after upgrading (gardener-operator,
gardenlet, and anything else using `Registry`). The content is semantically identical.

If this one-time rewrite is unwanted, drop `LGWin` and keep only the pool (variant B): no byte changes, but a smaller
gain in the operator (brotli alloc ~436 -> ~218 MB in the measurement above, vs ~60 MB with both). Given the window
sweep above, pool-only is a reasonable, lower-risk first PR, with the window change as a follow-up.

## Testing

- `go test ./pkg/utils/... ./pkg/component/... ./pkg/operator/... ./pkg/gardenlet/... ./pkg/resourcemanager/...` passes,
  except `pkg/utils/gardener/shootstate`, which also fails on the unmodified code (unrelated).
- Not run: the rest of the repo (`extensions/`, `test/`, integration/envtest tests). Other hardcoded hashes derived from
  compressed data may exist there; grep for `reference.resources.gardener.cloud/secret-` and `managedresource-` name
  suffixes in tests.
- Deployed in the local setup: Garden stays `Succeeded`, no new errors beyond transient startup errors.

## Files for the PR

- `pkg/utils/managedresources/registry.go`
- `pkg/utils/test/compression.go`
- `pkg/component/nodemanagement/dependencywatchdog/access_test.go`

Everything else in the working tree (`cmd/dummy`, CRD changes, `findings.md`, `dev-setup/skaffold-operator.yaml`, etc.) is
unrelated to this optimization and must not be included.
