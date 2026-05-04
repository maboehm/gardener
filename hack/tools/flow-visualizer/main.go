// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// flow-visualizer statically analyzes Go source files that use the
// github.com/gardener/gardener/pkg/utils/flow package and emits a Mermaid
// flowchart for every reconciliation graph found in the file.
//
// Usage:
//
//	flow-visualizer --input <file.go> --destination <out.md> [--func-name <name>] [--markdown]
//
// Flags:
//
//	--input         Path to the Go source file to analyze.
//	--destination   Path to write the generated output to.
//	--func-name     Only emit graphs defined inside this function (optional).
//	--markdown      Wrap the Mermaid output in a Markdown document with a title
//	                heading and a fenced code block. Useful for writing directly
//	                to a .md file that GitHub will render natively.
package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	opts := &Options{}

	cmd := &cobra.Command{
		Use:   "flow-visualizer",
		Short: "Generate Mermaid flowcharts from Go flow graph definitions.",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true

			if err := opts.Validate(); err != nil {
				return err
			}

			return run(opts)
		},
	}

	opts.AddFlags(cmd.Flags())

	if err := cmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func run(opts *Options) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, opts.Input, nil, 0)
	if err != nil {
		return fmt.Errorf("parse error: %w", err)
	}

	graphs := extractGraphs(f, opts.FuncName)
	if len(graphs) == 0 {
		return fmt.Errorf("no flow graphs found")
	}

	out, err := os.Create(opts.Destination)
	if err != nil {
		return fmt.Errorf("cannot create destination file: %w", err)
	}
	defer out.Close()

	for i, g := range graphs {
		if i > 0 {
			fmt.Fprintln(out)
		}
		if opts.Markdown {
			emitMarkdown(out, g)
		} else {
			emitMermaid(out, g)
		}
	}
	return nil
}
