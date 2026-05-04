// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

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
	kindTask nodeKind = iota
	kindSyncPoint
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
