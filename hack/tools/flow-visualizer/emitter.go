// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

// emitMermaid writes a Mermaid flowchart for g to w.
func emitMermaid(w io.Writer, g *graph) {
	fmt.Fprintf(w, "---\ntitle: %s (%s)\n---\n", g.name, g.funcName)
	fmt.Fprintln(w, "flowchart TD")
	fmt.Fprintln(w, "    classDef conditional stroke-dasharray:5 5,color:#888")
	fmt.Fprintln(w, "    classDef syncpoint fill:#ddf4ff,stroke:#4a90d9,color:#1a5276")
	fmt.Fprintln(w)

	// Emit task nodes.
	for _, t := range g.tasks {
		label := escapeMermaidLabel(t.name)
		if t.conditional {
			label += "\\n[CONDITIONAL]"
			// Hexagonal shape for conditional tasks.
			fmt.Fprintf(w, "    %s{{\"%s\"}}:::conditional\n", t.mermaidID, label)
		} else {
			fmt.Fprintf(w, "    %s[\"%s\"]\n", t.mermaidID, label)
		}
	}

	// Emit sync point nodes.
	for _, sp := range g.syncPoints {
		label := escapeMermaidLabel(prettifySyncPointName(sp.varName))
		// Stadium / pill shape: ([ … ])
		fmt.Fprintf(w, "    %s([\"%s\"]):::syncpoint\n", sp.mermaidID, label)
	}

	fmt.Fprintln(w)

	// Emit edges from sync point members → sync point node.
	for _, sp := range g.syncPoints {
		for _, memberVar := range sp.memberVars {
			entry, ok := g.varMap[memberVar]
			if !ok {
				continue
			}
			fmt.Fprintf(w, "    %s --> %s\n", entry.mermaidID, sp.mermaidID)
		}
	}

	// Emit edges from task dependencies → task.
	for _, t := range g.tasks {
		for _, depVar := range t.depVars {
			entry, ok := g.varMap[depVar]
			if !ok {
				continue
			}
			fmt.Fprintf(w, "    %s --> %s\n", entry.mermaidID, t.mermaidID)
		}
	}
}

// emitMarkdown wraps the Mermaid output in a Markdown document suitable for
// committing to the repository and rendering natively on GitHub.
func emitMarkdown(w io.Writer, g *graph) {
	fmt.Fprintln(w, "<!-- This file is auto-generated via `make generate`. DO NOT EDIT. -->")
	fmt.Fprintln(w)
	fmt.Fprintf(w, "# %s\n\n", g.name)
	fmt.Fprintln(w, "```mermaid")
	emitMermaid(w, g)
	fmt.Fprintln(w, "```")
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
