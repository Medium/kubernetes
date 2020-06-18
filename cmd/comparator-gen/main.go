package main

import (
	"go.medium.engineering/kubernetes/cmd/comparator-gen/generators"
	"k8s.io/gengo/args"

	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	arguments := args.Default()

	// Override defaults.
	arguments.OutputFileBaseName = "zz_generated.comparator"

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