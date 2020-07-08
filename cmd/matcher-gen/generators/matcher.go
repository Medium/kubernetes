package generators

import (
	"k8s.io/gengo/args"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
	"k8s.io/klog/v2"
	"path/filepath"
	"sort"
	"strings"
)

func NameSystems() namer.NameSystems {
	return namer.NameSystems{}
}

func DefaultNameSystem() string {
	return "public"
}
func Packages(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
	var pkgs generator.Packages
	for _, i := range context.Inputs {
		klog.V(5).Infof("considering %q", i)
		pkg := context.Universe[i]
		if pkg == nil {
			continue
		}
		var newComparators []*types.Type
		for _, t := range pkg.Types {
			if strings.HasSuffix(t.Name.String(), "Spec") {
				continue
			}
			for _, field := range t.Members {
				if field.Name == "ObjectMeta" {
					klog.V(5).Infof("Adding %q", t)
					newComparators = append(newComparators, t)
					break
				}
			}
		}
		sort.Slice(newComparators, func(i, j int) bool {
			return newComparators[i].Name.String() < newComparators[j].Name.String()
		})
		if len(newComparators) == 0 {
			klog.V(5).Infof("No valid types found in %q", pkg.Path)
			continue
		}
		pkgs = append(pkgs, &generator.DefaultPackage{
			PackageName: filepath.Base(arguments.OutputPackagePath),
			PackagePath: arguments.OutputPackagePath,
			HeaderText: []byte(""),
			GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
				return []generator.Generator{
					NewGenerator(arguments.OutputFileBaseName, pkg.Path, arguments.OutputPackagePath, newComparators),
				}
			},
			FilterFunc: func(c *generator.Context, t *types.Type) bool {
				return t.Name.Package == pkg.Path
			},
		})
	}
	return pkgs
}
