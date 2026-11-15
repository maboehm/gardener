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
(now `a55185c6` for the secret name suffix and `07af3d05` for the `reference.resources.gardener.cloud/secret-*` annotation with `LGWin: 20`) because they are derived from the compressed data.

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

## Follow-up: serialization format (YAML -> JSON)

After the brotli change, the next large source of allocation churn is the serialization of the objects in
`Registry.Add`. This section collects the analysis; nothing from it is part of the brotli PR.

### Current behavior

- Every production caller of `managedresources.NewRegistry` passes a **YAML** serializer: one of the seven shared
  package-level variables `kubernetes.{Seed,Shoot,Garden}Serializer` (`pkg/client/kubernetes/types.go`) and
  `operatorclient.{Runtime,Virtual}Serializer` (`pkg/operator/client/scheme.go`), all built with
  `SerializerOptions{Yaml: true, Pretty: false, Strict: false}`. The registry tests use `Yaml: true` as well.
- `NewRegistry` reads `serializer.Identifier()` (a JSON string describing the serializer options), unmarshals it and
  sets `isYAMLSerializer` when `yaml == "true"`. This flag is used in a single place, `Registry.Add`.
- When the flag is set, `Add` normalizes the output of the serializer: it parses the YAML with `yaml.v4` into a generic
  value and dumps it again (`yaml.WithV3Defaults()`, line width -1, indent 2, compact sequence indent). The code comment
  says this is a workaround for inconsistent/unstable map key ordering of yaml.v2 (go-yaml PR 736) in the apimachinery
  YAML serializer, to get stable bytes (and therefore stable hash-derived secret names). This is one extra full parse and
  emit per object, which is the YAML parser/emitter allocation seen in the operator profile (~150 MB alloc_space, plus
  ~105 MB `json.Marshal`, of 1450 MB total, before the brotli fix; not re-profiled since).
- If the flag is not set, the output of the serializer is stored as is (JSON for a JSON serializer). The reproduction
  below relies on that.
- The registry was not modified for this analysis other than adding constructors (next section).

### Prototype: constructors

In the working tree (not committed):

- `NewRegistryWithSerializer(scheme, codec, serializer)`: the previous `NewRegistry` implementation, unchanged.
- `NewJSONRegistry(scheme, codec)`: builds a JSON serializer (`Yaml: false, Pretty: false, Strict: false`) from the
  scheme and calls `NewRegistryWithSerializer`.
- `NewRegistry(scheme, codec, serializer)`: forwards to `NewRegistryWithSerializer` (behavior unchanged).

Cost of building a serializer per registry: negligible. `json.NewSerializerWithOptions` is a struct fill plus a
`json.Marshal` of a 4-key map (~0.6 us, 570 B, 11 allocs). `NewRegistry` itself costs ~78 us and ~235 KB per call (it
iterates `scheme.AllKnownTypes()`, dedupes and sorts group versions, and builds the codec), the same with or without the
serializer. If constructor cost ever matters, that is the part to cache.

### Helper programs: `hack/registry-bench/` (untracked in the working tree on `optimize-brotli`)

- `extract.sh <dir>`: writes the `data.yaml.br` of all Secrets of the current kubeconfig's cluster to `<dir>`.
- `main.go`: decodes all extracted payloads into objects (typed where the seed scheme knows the kind, otherwise
  unstructured, skipping kinds the codec cannot encode) and measures wall time, CPU (`getrusage`), allocation, mallocs and
  GCs per pass. `-phase write|read`, `-mode yaml|json|json-nosep|json-leadsep|kyaml|kyaml-compact`,
  `-iterations N`, `-cpuprofile f`. The `kyaml*` and `json-nosep`/`json-leadsep` modes replicate the registry's work in
  the benchmark; they are not registry implementations.

Input: 117 real ManagedResource secrets from the local KinD runtime cluster, 873 typed objects (15 objects skipped because
the seed scheme cannot encode their kind, mostly CRs). `LGWin: 20` + pool as in the current branch.
Numbers are per pass over all payloads, 3 runs per variant (20-30 passes per run).

### Write path (NewRegistry + AddAllAndSerialize)

| Variant                                        | wall per pass | CPU per pass | allocated  | mallocs | compressed size   |
| ---------------------------------------------- | ------------- | ------------ | ---------- | ------- | ----------------- |
| A: yaml (current, caller passes YAML serializer) | 221-223 ms    | 472-483 ms   | 468 MB     | 2.19 M  | 361,498 B         |
| B: json (`NewRegistry` -> `NewJSONRegistry`)   | 33.4-33.7 ms  | 51-53 ms     | 51 MB      | 60 K    | 357,954 B (-1.0%) |
| kyaml (replicated, JSON + `kyaml.Encoder`)     | 103-104 ms    | 171-174 ms   | 139-144 MB | 866 K   | 377,300 B (+4.4%) |
| kyaml-compact                                  | 90-91 ms      | 155-159 ms   | 134-137 MB | 865 K   | 356,555 B (-1.4%) |
| json, pretty (indented)                        | 42.5-42.9 ms  | not measured | 65-69 MB   | 61 K    | 385,190 B (+6.6%) |

(A and B are from the first session, 30 passes per run. A second session, which also measured the kyaml rows, gave
yaml 232-236 ms wall / 493-511 ms CPU / 468-478 MB and json 34.6-35.0 ms wall / 53-55 ms CPU / 49-51 MB, i.e. the same
picture with some machine noise. The pretty-JSON row is from an earlier benchmark via the registry tests' serializer swap.)

- JSON direct: ~6.6x less wall time, ~9x less CPU, ~9x less allocation than today, and 1% smaller compressed data.
- Output was deterministic: 20 rounds of serializing every payload twice gave byte-identical compressed output.
- KYAML (`sigs.k8s.io/yaml/kyaml`, same dependency as `sigs.k8s.io/yaml` v1.6.0) is a strict YAML subset (braces,
  quoted strings, `---` header per document) and is what `kubectl get -o kyaml` prints. The encoder takes JSON, parses it
  into `yaml.v3` nodes and renders them, so it is ~3x more expensive than JSON but ~2.3x cheaper in wall time and ~3x in
  CPU/allocation than today. It stays readable with plain `brotli -d`.

### Read path (gardener-resource-manager decoding)

Measured as the GRM `managedresource` reconciler does it: `brotli.NewReader` + `yaml.NewYAMLOrJSONDecoder(reader, 1024)`
+ `Decode(&map[string]any)` for every document (873 objects per pass). Only the decoding is measured, not the rest of a
GRM reconcile (REST mapping, unstructured, apply/diff).

| Variant (read)                       | wall per pass | CPU per pass | allocated | mallocs |
| ------------------------------------ | ------------- | ------------ | --------- | ------- |
| yaml (current)                       | 70.9-71.8 ms  | 134 ms       | 84.0 MB   | 953 K   |
| json, `---` between documents        | 67.6-68.8 ms  | 152-154 ms   | 111.8 MB  | 834 K   |
| json with an additional leading `---` | 70.5-72.7 ms  | 159-164 ms   | 120.1 MB  | 878 K   |
| json, no `---` (newline-delimited)   | 18.7-19.0 ms  | 36.5-37.8 ms | 24.4 MB   | 203 K   |
| kyaml                                | 75.0-76.2 ms  | 141-143 ms   | 89.7 MB   | 954 K   |
| kyaml-compact                        | 73.2-74.8 ms  | 166-170 ms   | 120 MB    | 878 K   |

Why: `NewYAMLOrJSONDecoder` guesses from the first bytes (`GuessJSONStream`). If the stream looks like JSON it uses a
`json.Decoder`. After the first document, `---` is a JSON syntax error; since only one object was seen the stream is
still ambiguous, so the decoder rewinds and switches to the YAML-to-JSON decoder for the rest. So JSON with `---`
separators gets **no** read benefit (and wastes a failed JSON attempt). A leading `---` makes it pick the YAML decoder
immediately, which avoids the failed attempt but is not cheaper than today's block YAML (the YAML path converting
JSON-looking documents is as expensive). KYAML starts with `---` too and therefore reads like YAML. Only a stream that
stays JSON for its whole length (first byte `{`, no `---`) uses the fast `json.Decoder` path: ~3.7x less wall time and
CPU, ~3.4x less allocation.

Reads probably outnumber writes: the `managedresource` reconciler requeues after `SyncPeriod` after every run (and runs
on changes), while the secret content only changes when a component changes. The production sync period was not looked up.

### Is newline-delimited JSON valid YAML? No.

Two JSON objects on separate lines without `---` are rejected as YAML: `yq -p yaml` fails with
`did not find expected <document start>`, PyYAML raises `ParserError`. `yq -p json` and `jq` read them. The GRM decoder
accepts it only because of its JSON-first heuristic. Storing it under `data.yaml.br` would be mislabeled. So if the
read-side speedup is wanted, use a new key, e.g. `data.json.br` (`CompressedJSONDataKey = "data.json" + BrotliCompressionSuffix`
next to `CompressedDataKey` in `pkg/apis/resources/v1alpha1/types.go`), and let readers choose the decoder from the key
(JSON stream for `.json.br`, YAML for `.yaml.br`) instead of guessing.

Readers found in this repo (non-test code):

- GRM `managedresource` reconciler (`pkg/resourcemanager/controller/managedresource/reconciler.go`): iterates all secret
  keys, decompresses keys with the `.br` suffix, lets the decoder guess. An already released GRM would read `data.json.br`
  (it does not look at the rest of the key name) and, because the stream is clean JSON, on the fast path.
- `managedresources.ExtractObjectsFromSecret` (`pkg/utils/managedresources/managedresources.go`): iterates all `.br`
  keys but splits on `---\n`; breaks for newline-delimited JSON, needs to change.
- `ExtractManifestsFromManagedResourceData` (`pkg/utils/test/compression.go`): looks up exactly `CompressedDataKey`,
  needs to change.

Suggested rollout (readers first, as there is time for one or two releases):

1. Release N: readers choose the decoder from the key and handle both `data.yaml.br` and `data.json.br`
   (`ExtractObjectsFromSecret`, the test helper, GRM). Writers still emit `data.yaml.br`.
2. Release N+1 or N+2: the registry writes `data.json.br` (newline-delimited JSON).

Open points to check before doing it:

- Does the code that writes the ManagedResource secret replace `Data` or merge it? If it merges, an updated secret
  could carry both keys and GRM would decode every object twice.
- External consumers (extensions or tooling outside this repo) that read `data.yaml.br` by name.
- Whether the registry tests and the many component tests that assert YAML text need to move to a JSON-aware assertion
  (see Tests below).

### Usability

Single-line JSON is hard to read with `base64 -d | brotli -d`. Verified on a real payload (`alertmanager-garden`,
8 objects, produced with `NewJSONRegistry`): `base64 -d | brotli -d | yq -P .` prints block-style YAML with `---`
separators and all 8 kinds. Field order follows the JSON (struct) order (`kind`, `apiVersion`, `metadata`, ...) instead
of alphabetical. For newline-delimited JSON, `jq .` or `yq -p json -P .` is needed (only `yq -p json` on a two-object
sample was tested, not a real payload). Document the pipeline wherever the old hint lives.

### Tests

With `NewRegistry` temporarily pointing to `NewJSONRegistry` (variant B), 23 test packages failed in
`./pkg/utils/managedresources/...`, `./pkg/component/...`, `./pkg/operator/...`, `./pkg/gardenlet/...`, `./extensions/...`
(excluding the already failing `shootstate`), e.g. `resourcemanager`, `coredns`, `istio`, `nodelocaldns`, `fluentbit`,
`plutono`, `seed/system`, `botanist`, `managedresources`. They assert YAML text taken from the registry. The individual
failures were not inspected, some may have other causes. `NewRegistry` is back to forwarding to
`NewRegistryWithSerializer`, so the working tree is green.

### Other things that were evaluated

- **CBOR** instead of YAML: not recommended. It changes the data contract (key, readers, tooling, docs), has the same
  version-skew and one-time-rewrite problems, is not debuggable without a tool, and brotli on top of YAML already
  equalizes sizes. Library support for Gardener's types was not checked. Emitting JSON gets most of the benefit without
  a new encoding.
- **Pooling brotli writers in `pkg/utils/test/compression.go`**: not worth it. The helper is called a handful of times per
  spec (about 0.3 ms and 2.6 MB per call unpooled); it should stay an independent oracle for the production code. The
  helper only needs to use the same writer options.

### Recommendation

1. PR 1: the brotli change in this document (pool + `LGWin: 20`), optionally pool-only to avoid the one-time rewrite.
2. PR 2 (readers): choose the decoder by key; support `data.json.br` in `ExtractObjectsFromSecret`, the test helper and
   GRM.
3. PR 3 (writers, one or two releases later): `NewJSONRegistry` as the default, newline-delimited JSON in `data.json.br`,
   test updates, docs.
4. Before PR 3, re-profile the operator after PR 1 to rank the serialization cost against the rest of a reconcile, and
   measure a real GRM (the read numbers above are a decoder micro-benchmark).
