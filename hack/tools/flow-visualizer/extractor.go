// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

// extractGraphs walks the top-level function declarations in f and returns all
// flow graphs found. If filterFunc is non-empty only graphs inside that
// function are returned.
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

	// Stage 1: discover graph variables.
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

	// Stage 2: collect tasks and sync points.
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
// from AssignStmt and ValueSpec nodes. All other nodes return nil, nil.
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

// graphNameFromArg extracts the name passed to flow.NewGraph("…").
// The argument is commonly a plain string literal or a fmt.Sprintf(format, …)
// call. In the latter case we use the format string (without verbs) as a
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
			if fmtStr := stringArg(inner, 0); fmtStr != "" {
				// Strip format verbs (%s, %v, …) and trim trailing punctuation/spaces.
				clean := strings.TrimRight(strings.ReplaceAll(fmtStr, "%s", ""), " ,-_")
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
