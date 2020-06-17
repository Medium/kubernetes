package main

import (
	"go.medium.engineering/kubernetes/cmd/comparator-gen/generators"
	"k8s.io/gengo/args"

	"github.com/spf13/pflag"
	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	arguments := args.Default()

	// Override defaults.
	arguments.OutputFileBaseName = "comparator_generated"

	// Custom args.
	customArgs := &generators.CustomArgs{
		ExtraPeerDirs: []string{},
	}
	pflag.CommandLine.StringSliceVar(&customArgs.ExtraPeerDirs, "extra-peer-dirs", customArgs.ExtraPeerDirs,
		"Comma-separated list of import paths which are considered, after tag-specified peers, for conversions.")
	arguments.CustomArgs = customArgs

	// Run it.
	if err := arguments.Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		generators.Packages,
	); err != nil {
		klog.Fatalf("Error: %v", err)
	}
	klog.V(2).Info("Completed successfully.")
}