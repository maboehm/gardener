// SPDX-FileCopyrightText: Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	_ "k8s.io/component-base/logs/json/register" // for JSON log format registration
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	runtimemetrics "sigs.k8s.io/controller-runtime/pkg/metrics"

	admissioncontroller "github.com/gardener/gardener/cmd/gardener-admission-controller/app"
	apiserver "github.com/gardener/gardener/cmd/gardener-apiserver/app"
	controllermanager "github.com/gardener/gardener/cmd/gardener-controller-manager/app"
	operator "github.com/gardener/gardener/cmd/gardener-operator/app"
	resourcemanager "github.com/gardener/gardener/cmd/gardener-resource-manager/app"
	scheduler "github.com/gardener/gardener/cmd/gardener-scheduler/app"
	gardenlet "github.com/gardener/gardener/cmd/gardenlet/app"
	"github.com/gardener/gardener/cmd/utils"
	admissioncontrollerfeatures "github.com/gardener/gardener/pkg/admissioncontroller/features"
	apiserverfeatures "github.com/gardener/gardener/pkg/apiserver/features"
	controllermanagerfeatures "github.com/gardener/gardener/pkg/controllermanager/features"
	gardenletfeatures "github.com/gardener/gardener/pkg/gardenlet/features"
	operatorfeatures "github.com/gardener/gardener/pkg/operator/features"
	schedulerfeatures "github.com/gardener/gardener/pkg/scheduler/features"
	"github.com/gardener/gardener/pkg/utils/flow"
)

type component struct {
	newCommand func() *cobra.Command
	init       func()
}

var components = map[string]component{
	"admission-controller": {admissioncontroller.NewCommand, admissioncontrollerfeatures.RegisterFeatureGates},
	"apiserver":            {apiserver.NewCommand, apiserverfeatures.RegisterFeatureGates},
	"controller-manager":   {controllermanager.NewCommand, controllermanagerfeatures.RegisterFeatureGates},
	"operator":             {operator.NewCommand, operatorfeatures.RegisterFeatureGates},
	"resource-manager":     {resourcemanager.NewCommand, nil},
	"scheduler":            {scheduler.NewCommand, schedulerfeatures.RegisterFeatureGates},
	"gardenlet": {gardenlet.NewCommand, func() {
		gardenletfeatures.RegisterFeatureGates()
		flow.RegisterMetrics(runtimemetrics.Registry)
	}},
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	c, ok := components[os.Args[1]]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown component %q\n", os.Args[1])
		usage()
	}

	utils.DeduplicateWarnings()
	if c.init != nil {
		c.init()
	}
	debug.FreeOSMemory()

	cmd := c.newCommand()
	cmd.SetArgs(os.Args[2:])
	if err := cmd.ExecuteContext(signals.SetupSignalHandler()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func usage() {
	names := make([]string, 0, len(components))
	for n := range components {
		names = append(names, n)
	}
	sort.Strings(names)
	fmt.Fprintf(os.Stderr, "usage: gardener <%s> [flags]\n", strings.Join(names, "|"))
	os.Exit(2)
}
