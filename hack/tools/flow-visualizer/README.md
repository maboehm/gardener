# flow-visualizer

`flow-visualizer` is a static analysis tool that reads a Go source file, finds
reconciliation graphs built with the
[`pkg/utils/flow`](../../../pkg/utils/flow/) package, and emits a
[Mermaid](https://mermaid.js.org/) flowchart for each graph it finds.

No code is executed — the tool works purely on the source text using the Go
parser.

## Usage

```bash
go run ./hack/tools/flow-visualizer/main.go <file.go> [function-name]
```

| Argument | Required | Description |
|---|---|---|
| `<file.go>` | yes | Path to the Go source file to analyze |
| `[function-name]` | no | Only emit graphs found inside this function. Useful when a file contains more than one graph-building function. |

Output is written to stdout as Mermaid markup and can be redirected to a file
or piped directly into a renderer.

### Examples

Visualize all graphs in the shoot reconcile flow:

```bash
go run ./hack/tools/flow-visualizer/main.go \
  pkg/gardenlet/controller/shoot/shoot/reconciler_reconcile.go
```

Visualize only the graph built inside a specific function:

```bash
go run ./hack/tools/flow-visualizer/main.go \
  pkg/gardenlet/controller/shoot/shoot/reconciler_reconcile.go \
  runReconcileShootFlow
```

Render the diagram directly in the terminal using
[`mermaid-js/mermaid-cli`](https://github.com/mermaid-js/mermaid-cli):

```bash
go run ./hack/tools/flow-visualizer/main.go \
  pkg/gardenlet/controller/shoot/shoot/reconciler_reconcile.go \
  > /tmp/graph.mmd

mmdc -i /tmp/graph.mmd -o /tmp/graph.svg
```

## Output format

The tool emits a Mermaid `flowchart TD` (top-down) diagram.  Each `g.Add()`
call becomes a node and each `Dependencies:` entry becomes a directed edge.

Three visual node types are used:

### Regular task

A task whose `SkipIf` field is `false` or absent.

```
DeployingShootNamespaceInSeed["Deploying Shoot namespace in Seed"]
```

Rendered as a plain rectangle.

### Conditional task

A task whose `SkipIf` field contains a non-trivial expression (i.e. execution
depends on runtime conditions).

```
DeployingShootInfrastructure{{"Deploying Shoot infrastructure\n[CONDITIONAL]"}}:::conditional
```

Rendered as a hexagon with a dashed border and muted text.  The label includes
a `[CONDITIONAL]` suffix.

### Sync point

A named `flow.TaskIDs` set that is assigned to a variable and later referenced
as a dependency, acting as a synchronization barrier without containing any
action of its own.

```go
// Source
syncPointAllSystemComponentsDeployed = flow.NewTaskIDs(
    waitUntilNetworkIsReady,
    deployCoreDNS,
    // …
)
```

```
SyncPointAllSystemComponentsDeployed(["Sync: All System Components Deployed"]):::syncpoint
```

Rendered as a pill/stadium shape with a blue background.  All member tasks
receive an edge pointing into the sync point node; tasks that depend on the
sync point receive an edge out of it.  This keeps the diagram compact — instead
of N×M edges between every producer and every consumer, the sync point node
acts as the single intermediary.

### CSS classes

| Class | Applied to | Appearance |
|---|---|---|
| *(none)* | Regular tasks | Plain rectangle |
| `conditional` | Conditional tasks | Dashed border, muted colour |
| `syncpoint` | Sync point barriers | Blue pill, blue border |

## What the tool detects

The tool recognizes the following patterns from the `flow` package:

| Pattern | Detected as |
|---|---|
| `g := flow.NewGraph("name")` | Graph declaration |
| `taskA := g.Add(flow.Task{Name: "…", …})` | Task node |
| `_ = g.Add(flow.Task{Name: "…", …})` | Task node (no outgoing variable) |
| `Dependencies: flow.NewTaskIDs(a, b)` | Edges from `a` and `b` to the task |
| `flow.NewTaskIDs(a).InsertIf(cond, b)` | Edge from `b` (conditional dep) |
| `sync = flow.NewTaskIDs(a, b, c)` | Sync point with members `a`, `b`, `c` |
| `Dependencies: flow.NewTaskIDs(sync)` | Edge from the sync point node |

`SkipIf` expressions are detected textually: any value other than the
identifier `false` marks the task as conditional.

## Limitations

- The graph name is extracted from the string literal passed to
  `flow.NewGraph`.  When a `fmt.Sprintf` call is used, only the format string
  (without verb placeholders) is shown.
- Tasks or sync points that are declared but whose variable name cannot be
  resolved (e.g. intermediate expressions) produce no edge; a comment is not
  emitted.
- The tool does not evaluate boolean expressions in `SkipIf` or `InsertIf` —
  it only records whether a condition is present.
