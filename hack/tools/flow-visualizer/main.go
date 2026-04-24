// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// flow-visualizer statically analyzes Go source files that use the
// github.com/gardener/gardener/pkg/utils/flow package and emits a Mermaid
// flowchart for every reconciliation graph found in the file.
//
// Usage:
//
//	go run ./hack/tools/flow-visualizer/main.go <file.go> [function-name]
//
// If [function-name] is given only the graph(s) constructed inside that
// function are emitted; otherwise every graph in the file is emitted.
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"unicode"
)

// ----------------------------------------------------------------------------
// Domain types
// ----------------------------------------------------------------------------

// taskNode holds the information extracted from a single g.Add(flow.Task{…}) call.
type taskNode struct {
	// varName is the Go identifier the return value of g.Add() was assigned to
	// (empty string / "_" when the result is discarded).
	varName string
	// name is the value of the Name field – used as the human-readable label.
	name string
	// mermaidID is a sanitized, unique identifier for the Mermaid diagram.
	mermaidID string
	// conditional is true when a non-trivial SkipIf expression was found.
	conditional bool
	// depVars is the list of Go variable names found inside the
	// Dependencies: flow.NewTaskIDs(…) / .InsertIf(…) calls.
	depVars []string
}

// syncPoint represents a bare flow.NewTaskIDs(…) assignment – a named
// synchronisation barrier that carries no action itself.
type syncPoint struct {
	// varName is the Go identifier (e.g. "syncPointAllSystemComponentsDeployed").
	varName string
	// mermaidID is the sanitized Mermaid node identifier.
	mermaidID string
	// memberVars are the Go variable names passed to flow.NewTaskIDs(…).
	memberVars []string
}

// nodeKind distinguishes what a Go variable refers to.
type nodeKind int

const (
	kindTask      nodeKind = iota
	kindSyncPoint nodeKind = iota
)

// varEntry is a value in the unified variable-to-node map.
type varEntry struct {
	kind      nodeKind
	mermaidID string
}

// graph represents one flow.NewGraph(…) invocation together with all tasks
// and sync points that belong to it.
type graph struct {
	// graphVar is the name of the *flow.Graph Go variable (typically "g").
	graphVar string
	// name is the string literal passed to flow.NewGraph("…").
	name string
	// funcName is the name of the enclosing function declaration.
	funcName string
	// tasks preserves insertion order.
	tasks []*taskNode
	// syncPoints preserves insertion order.
	syncPoints []*syncPoint
	// varMap maps every known Go variable name to its resolved node entry.
	varMap map[string]varEntry
}

func newGraph(graphVar, name, funcName string) *graph {
	return &graph{
		graphVar: graphVar,
		name:     name,
		funcName: funcName,
		varMap:   make(map[string]varEntry),
	}
}

func (g *graph) addTask(t *taskNode) {
	g.tasks = append(g.tasks, t)
	if t.varName != "" && t.varName != "_" {
		g.varMap[t.varName] = varEntry{kind: kindTask, mermaidID: t.mermaidID}
	}
}

func (g *graph) addSyncPoint(sp *syncPoint) {
	g.syncPoints = append(g.syncPoints, sp)
	g.varMap[sp.varName] = varEntry{kind: kindSyncPoint, mermaidID: sp.mermaidID}
}

func (g *graph) hasTaskName(name string) bool {
	for _, t := range g.tasks {
		if t.name == name {
			return true
		}
	}
	return false
}

func (g *graph) hasSyncPointVar(varName string) bool {
	for _, sp := range g.syncPoints {
		if sp.varName == varName {
			return true
		}
	}
	return false
}

// ----------------------------------------------------------------------------
// Entry point
// ----------------------------------------------------------------------------

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: flow-visualizer <file.go> [function-name]")
		os.Exit(1)
	}

	filePath := os.Args[1]
	filterFunc := ""
	if len(os.Args) >= 3 {
		filterFunc = os.Args[2]
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, nil, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
		os.Exit(1)
	}

	graphs := extractGraphs(f, filterFunc)
	if len(graphs) == 0 {
		fmt.Fprintln(os.Stderr, "no flow graphs found")
		os.Exit(1)
	}

	for i, g := range graphs {
		if i > 0 {
			fmt.Println()
		}
		emitMermaid(g)
	}
}

// ----------------------------------------------------------------------------
// AST extraction
// ----------------------------------------------------------------------------

func extractGraphs(f *ast.File, filterFunc string) []*graph {
	var graphs []*graph

	for _, decl := range f.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok || funcDecl.Body == nil {
			continue
		}
		if filterFunc != "" && funcDecl.Name.Name != filterFunc {
			continue
		}
		graphs = append(graphs, extractFromFunc(funcDecl)...)
	}
	return graphs
}

// extractFromFunc performs a two-stage walk of a function body:
//
//  1. Collect flow.NewGraph(…) declarations to discover graph variables.
//  2. Collect g.Add(flow.Task{…}) calls and bare flow.NewTaskIDs(…) assignments
//     (sync points), associating them with the correct graph variable.
//
// The two-stage approach is necessary because the var block containing both
// the graph variable and all the task variables is a single AST node.
func extractFromFunc(funcDecl *ast.FuncDecl) []*graph {
	funcName := funcDecl.Name.Name

	// graphsByVar maps graph-variable-name → *graph.
	graphsByVar := map[string]*graph{}
	var graphOrder []*graph

	// mermaidIDs tracks already-issued IDs to ensure uniqueness.
	mermaidIDs := map[string]int{}

	newMermaidID := func(base string) string {
		id := sanitizeID(base)
		if n, exists := mermaidIDs[id]; exists {
			mermaidIDs[id] = n + 1
			return fmt.Sprintf("%s_%d", id, n+1)
		}
		mermaidIDs[id] = 1
		return id
	}

	// --- Stage 1: discover graph variables ---
	ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
		rhsList, lhsNames := rhsLhsFromNode(n)
		for i, rhs := range rhsList {
			call, ok := rhs.(*ast.CallExpr)
			if !ok {
				continue
			}
			if !isFlowNewGraph(call) {
				continue
			}
			varName := safeIndex(lhsNames, i)
			if varName == "" || varName == "_" {
				continue
			}
			gName := graphNameFromArg(call)
			g := newGraph(varName, gName, funcName)
			graphsByVar[varName] = g
			graphOrder = append(graphOrder, g)
		}
		return true
	})

	if len(graphsByVar) == 0 {
		return nil
	}

	// --- Stage 2: collect tasks and sync points ---
	// We need to walk all assignments/value-specs in order, which means a
	// single Inspect pass collecting everything as we encounter it.
	ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
		rhsList, lhsNames := rhsLhsFromNode(n)
		for i, rhs := range rhsList {
			call, ok := rhs.(*ast.CallExpr)
			if !ok {
				continue
			}
			varName := safeIndex(lhsNames, i)

			// g.Add(flow.Task{…}) ?
			if gVar, ok := isGraphAdd(call); ok {
				g, exists := graphsByVar[gVar]
				if !exists {
					continue
				}
				task := extractTask(call, varName, newMermaidID)
				if task != nil && !g.hasTaskName(task.name) {
					g.addTask(task)
				}
				continue
			}

			// bare flow.NewTaskIDs(…) used as sync point?
			if isFlowNewTaskIDs(call) {
				if varName == "" || varName == "_" {
					continue
				}
				// Associate with the graph that was most recently declared
				// (heuristic: sync points appear inside the same var block
				// as the tasks they reference).
				g := mostRecentGraph(graphOrder)
				if g == nil || g.hasSyncPointVar(varName) {
					continue
				}
				memberVars := collectIdentArgs(call)
				sp := &syncPoint{
					varName:    varName,
					mermaidID:  newMermaidID(varName),
					memberVars: memberVars,
				}
				g.addSyncPoint(sp)
			}
		}
		return true
	})

	return graphOrder
}

// mostRecentGraph returns the last graph in the slice (the one currently being
// built), or nil.
func mostRecentGraph(graphs []*graph) *graph {
	if len(graphs) == 0 {
		return nil
	}
	return graphs[len(graphs)-1]
}

// rhsLhsFromNode extracts parallel slices of (rhs expressions, lhs names)
// from AssignStmt and ValueSpec nodes.  All other nodes return nil, nil.
func rhsLhsFromNode(n ast.Node) (rhs []ast.Expr, lhsNames []string) {
	switch node := n.(type) {
	case *ast.AssignStmt:
		names := make([]string, len(node.Lhs))
		for i, l := range node.Lhs {
			if id, ok := l.(*ast.Ident); ok {
				names[i] = id.Name
			}
		}
		return node.Rhs, names
	case *ast.ValueSpec:
		names := make([]string, len(node.Names))
		for i, id := range node.Names {
			names[i] = id.Name
		}
		return node.Values, names
	}
	return nil, nil
}

// safeIndex returns lhsNames[i] or "" when i is out of range.
func safeIndex(lhsNames []string, i int) string {
	if i < len(lhsNames) {
		return lhsNames[i]
	}
	return ""
}

// isFlowNewGraph returns true when call is flow.NewGraph(…).
func isFlowNewGraph(call *ast.CallExpr) bool {
	return isSelectorCall(call, "flow", "NewGraph")
}

// isFlowNewTaskIDs returns true when call is flow.NewTaskIDs(…).
func isFlowNewTaskIDs(call *ast.CallExpr) bool {
	return isSelectorCall(call, "flow", "NewTaskIDs")
}

// isSelectorCall returns true when call is pkg.name(…).
func isSelectorCall(call *ast.CallExpr, pkg, name string) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	id, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}
	return id.Name == pkg && sel.Sel.Name == name
}

// isGraphAdd returns (graphVarName, true) when call looks like <ident>.Add(…).
func isGraphAdd(call *ast.CallExpr) (string, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", false
	}
	if sel.Sel.Name != "Add" {
		return "", false
	}
	id, ok := sel.X.(*ast.Ident)
	if !ok {
		return "", false
	}
	return id.Name, true
}

// extractTask builds a taskNode from a g.Add(flow.Task{…}) call.
func extractTask(call *ast.CallExpr, varName string, newID func(string) string) *taskNode {
	if len(call.Args) != 1 {
		return nil
	}
	lit, ok := call.Args[0].(*ast.CompositeLit)
	if !ok {
		return nil
	}

	t := &taskNode{varName: varName}
	for _, elt := range lit.Elts {
		kv, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		key, ok := kv.Key.(*ast.Ident)
		if !ok {
			continue
		}
		switch key.Name {
		case "Name":
			t.name = stringLit(kv.Value)
		case "SkipIf":
			t.conditional = !isFalseLiteral(kv.Value)
		case "Dependencies":
			t.depVars = extractDepVarNames(kv.Value)
		}
	}

	if t.name == "" {
		return nil
	}
	t.mermaidID = newID(t.name)
	return t
}

// extractDepVarNames collects Go variable names from dependency expressions:
//
//	flow.NewTaskIDs(a, b, c)
//	flow.NewTaskIDs(a, b).InsertIf(cond, d)
//	flow.NewTaskIDs(a).InsertIf(cond, b).InsertIf(cond2, c)
func extractDepVarNames(expr ast.Expr) []string {
	var vars []string
	collectDepArgs(expr, &vars)
	return vars
}

// collectDepArgs recursively walks method chains collecting ident arguments.
func collectDepArgs(expr ast.Expr, vars *[]string) {
	call, ok := expr.(*ast.CallExpr)
	if !ok {
		return
	}
	switch fn := call.Fun.(type) {
	case *ast.SelectorExpr:
		// Recurse on the receiver first (left-to-right chain).
		collectDepArgs(fn.X, vars)
		switch fn.Sel.Name {
		case "NewTaskIDs":
			*vars = append(*vars, collectIdentArgs(call)...)
		case "InsertIf":
			// InsertIf(condition, ids...) – skip arg[0] (the condition).
			if len(call.Args) > 1 {
				for _, arg := range call.Args[1:] {
					if id, ok := arg.(*ast.Ident); ok {
						*vars = append(*vars, id.Name)
					}
				}
			}
		case "Insert":
			*vars = append(*vars, collectIdentArgs(call)...)
		}
	}
}

// collectIdentArgs returns the names of all *ast.Ident arguments of a call.
func collectIdentArgs(call *ast.CallExpr) []string {
	var names []string
	for _, arg := range call.Args {
		if id, ok := arg.(*ast.Ident); ok {
			names = append(names, id.Name)
		}
	}
	return names
}

// ----------------------------------------------------------------------------
// AST value helpers
// ----------------------------------------------------------------------------

// graphNameFromArg extracts the name passed to flow.NewGraph("…").
// The argument is commonly a plain string literal or a fmt.Sprintf(format, …)
// call.  In the latter case we use the format string (without verbs) as a
// best-effort name.
func graphNameFromArg(call *ast.CallExpr) string {
	if len(call.Args) == 0 {
		return ""
	}
	// Plain string literal: flow.NewGraph("my graph")
	if s := stringLit(call.Args[0]); s != "" {
		return s
	}
	// fmt.Sprintf("Shoot cluster %s", …) → "Shoot cluster"
	if inner, ok := call.Args[0].(*ast.CallExpr); ok {
		if sel, ok := inner.Fun.(*ast.SelectorExpr); ok && sel.Sel.Name == "Sprintf" {
			if fmt := stringArg(inner, 0); fmt != "" {
				// Strip format verbs (%s, %v, …) and trim trailing punctuation/spaces.
				clean := strings.TrimRight(strings.ReplaceAll(fmt, "%s", ""), " ,-_")
				clean = strings.TrimRight(strings.ReplaceAll(clean, "%v", ""), " ,-_")
				return strings.TrimSpace(clean)
			}
		}
	}
	return ""
}

func stringArg(call *ast.CallExpr, i int) string {
	if i >= len(call.Args) {
		return ""
	}
	return stringLit(call.Args[i])
}

func stringLit(expr ast.Expr) string {
	bl, ok := expr.(*ast.BasicLit)
	if !ok || bl.Kind != token.STRING {
		return ""
	}
	s := bl.Value
	if len(s) >= 2 && (s[0] == '"' || s[0] == '`') {
		return s[1 : len(s)-1]
	}
	return s
}

func isFalseLiteral(expr ast.Expr) bool {
	id, ok := expr.(*ast.Ident)
	return ok && id.Name == "false"
}

// ----------------------------------------------------------------------------
// Mermaid emission
// ----------------------------------------------------------------------------

func emitMermaid(g *graph) {
	fmt.Printf("---\ntitle: %s (%s)\n---\n", g.name, g.funcName)
	fmt.Println("flowchart TD")
	fmt.Println("    classDef conditional stroke-dasharray:5 5,color:#888")
	fmt.Println("    classDef syncpoint fill:#ddf4ff,stroke:#4a90d9,color:#1a5276")
	fmt.Println()

	// --- Emit task nodes ---
	for _, t := range g.tasks {
		label := escapeMermaidLabel(t.name)
		if t.conditional {
			label += "\\n[CONDITIONAL]"
			// Hexagonal shape for conditional tasks.
			fmt.Printf("    %s{{\"%s\"}}:::conditional\n", t.mermaidID, label)
		} else {
			fmt.Printf("    %s[\"%s\"]\n", t.mermaidID, label)
		}
	}

	// --- Emit sync point nodes ---
	for _, sp := range g.syncPoints {
		label := escapeMermaidLabel(prettifySyncPointName(sp.varName))
		// Stadium / pill shape: ([ … ])
		fmt.Printf("    %s([\"%s\"]):::syncpoint\n", sp.mermaidID, label)
	}

	fmt.Println()

	// --- Emit edges from sync point members → sync point node ---
	for _, sp := range g.syncPoints {
		for _, memberVar := range sp.memberVars {
			entry, ok := g.varMap[memberVar]
			if !ok {
				continue
			}
			fmt.Printf("    %s --> %s\n", entry.mermaidID, sp.mermaidID)
		}
	}

	// --- Emit edges from task dependencies → task ---
	for _, t := range g.tasks {
		for _, depVar := range t.depVars {
			entry, ok := g.varMap[depVar]
			if !ok {
				continue
			}
			fmt.Printf("    %s --> %s\n", entry.mermaidID, t.mermaidID)
		}
	}
}

// prettifySyncPointName converts camelCase/mixed variable names into a more
// readable label by inserting spaces before uppercase runs and removing common
// prefixes like "syncPoint".
func prettifySyncPointName(varName string) string {
	// Strip leading "syncPoint" prefix (case-insensitive).
	trimmed := varName
	lower := strings.ToLower(varName)
	for _, prefix := range []string{"syncpoint", "sync_point"} {
		if strings.HasPrefix(lower, prefix) {
			trimmed = varName[len(prefix):]
			break
		}
	}
	if trimmed == "" {
		trimmed = varName
	}

	// Split camelCase into words at lower→upper boundaries only.
	// This keeps acronyms like "CRDs", "APIServer" intact as single tokens.
	runes := []rune(trimmed)
	n := len(runes)
	var words []string
	start := 0
	for i := 1; i < n; i++ {
		if unicode.IsLower(runes[i-1]) && unicode.IsUpper(runes[i]) {
			words = append(words, string(runes[start:i]))
			start = i
		}
	}
	if start < n {
		words = append(words, string(runes[start:]))
	}

	label := strings.Join(words, " ")
	if label == "" {
		return varName
	}
	return "Sync: " + label
}

// escapeMermaidLabel escapes characters that would break Mermaid label syntax.
func escapeMermaidLabel(s string) string {
	// Replace double quotes with single quotes inside labels.
	return strings.ReplaceAll(s, `"`, `'`)
}

// sanitizeID turns an arbitrary name into a valid Mermaid node identifier
// (letters and digits only, CamelCase).
func sanitizeID(name string) string {
	var sb strings.Builder
	upperNext := true
	for _, r := range name {
		if r == ' ' || r == '-' || r == '/' || r == '_' {
			upperNext = true
			continue
		}
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			continue
		}
		if upperNext {
			sb.WriteRune(unicode.ToUpper(r))
			upperNext = false
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
