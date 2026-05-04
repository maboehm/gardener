// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

// Options holds the configuration for the flow-visualizer tool.
type Options struct {
	// Input is the path to the Go source file to analyze (--input).
	Input string
	// FuncName restricts extraction to a single function (--func-name).
	FuncName string
	// Markdown wraps the Mermaid output in a Markdown document (--markdown).
	Markdown bool
	// Destination is the path the generated output is written to (--destination).
	Destination string
}

// AddFlags registers all flags on the given FlagSet.
func (o *Options) AddFlags(flags *flag.FlagSet) {
	flags.StringVar(&o.Input, "input", "", "Path to the Go source file to analyze")
	flags.StringVar(&o.FuncName, "func-name", "", "Only emit graphs defined inside this function (optional)")
	flags.BoolVar(&o.Markdown, "markdown", false, "Wrap the Mermaid output in a Markdown document")
	flags.StringVar(&o.Destination, "destination", "", "Path to write the generated output to")
}

// Validate returns an error if the Options configuration is incomplete.
func (o *Options) Validate() error {
	var errs []error

	if len(o.Input) == 0 {
		errs = append(errs, fmt.Errorf("input is required"))
	}

	if len(o.Destination) == 0 {
		errs = append(errs, fmt.Errorf("destination is required"))
	}

	if len(errs) > 0 {
		return fmt.Errorf("%v", errs)
	}
	return nil
}
