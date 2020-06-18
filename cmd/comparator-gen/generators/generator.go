package generators

import (
	"io"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
	"strings"
)

type gen struct {
	generator.DefaultGen
	typesPackage string
	outputPackage string
	imports namer.ImportTracker
	types []*types.Type
}

func NewGenerator(name, typesPackage, outputPackage string, types []*types.Type) generator.Generator {
	return &gen{
		DefaultGen: generator.DefaultGen{
			OptionalName: name,
		},
		typesPackage: typesPackage,
		outputPackage: outputPackage,
		imports: generator.NewImportTracker(),
		types: types,
	}
}

func comparerFnNamer() *namer.NameStrategy {
	return &namer.NameStrategy{
		Prefix:              "Compare_",
		Join:                func(pre string, in []string, post string) string {
			return pre + strings.Join(in, "_") + post
		},
	}
}

func (g *gen) Namers(c *generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.outputPackage, g.imports),
		"comparerFn": comparerFnNamer(),
	}
}

func (g *gen) isOtherPackage(pkg string) bool {
	if pkg == g.outputPackage {
		return false
	}
	if strings.HasSuffix(pkg, `"`+g.outputPackage+`"`) {
		return false
	}
	return true
}

func (g *gen) Imports(c *generator.Context) (imports []string) {
	for _, i := range g.imports.ImportLines() {
		if g.isOtherPackage(i) {
			imports = append(imports, i)
		}
	}
	return imports
}

func (g *gen) Init(c *generator.Context, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, c, "$", "$")

	args := generator.Args{
		"test": c.Universe.Type(types.Name{Package:"testing", Name:"T"}),
		"testify": c.Universe.Type(types.Name{Package:"github.com/stretchr/testify/assert", Name:"New"}),
		"comparator": c.Universe.Type(types.Name{Package: "go.medium.engineering/kubernetes/pkg/test", Name: "Comparator"}),
		"defaultComparator": c.Universe.Type(types.Name{Package: "go.medium.engineering/kubernetes/pkg/test", Name: "DefaultComparator"}),
		"object": c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/runtime", Name: "Object"}),
		"sanitizeObjectMeta": c.Universe.Type(types.Name{Package: "go.medium.engineering/kubernetes/pkg/test", Name: "Sanitize_ObjectMeta"}),
	}

	sw.Do("func init() {\n", nil)
	sw.Do("RegisterComparators($.defaultComparator|raw$)\n", args)
	sw.Do("}\n\n", nil)

	sw.Do("func RegisterComparators(comparator *$.comparator|raw$) {\n", args)
	for _, t := range g.types {
		args["t"] = t
		sw.Do("comparator.RegisterForType(&$.t|raw${}, func(t *$.test|raw$, a, b $.object|raw$) {\n", args)
		sw.Do("$.t|comparerFn$(t, a.(*$.t|raw$), b.(*$.t|raw$))\n", args)
		sw.Do("})\n\n", nil)
	}
	sw.Do("}\n\n", nil)

	for _, t := range g.types {
		args["t"] = t
		sw.Do("func $.t|comparerFn$(t *$.test|raw$, expected, actual *$.t|raw$) {\n", args)
		sw.Do("assert := $.testify|raw$(t)\n", args)
		sw.Do("e := expected.DeepCopyObject().(*$.t|raw$)\n", args)
		sw.Do("e.ObjectMeta = $.sanitizeObjectMeta|raw$(e.ObjectMeta, actual.ObjectMeta)\n", args)
		sw.Do("assert.EqualValues(expected, actual)\n", nil)
		sw.Do("}\n\n", nil)
	}

	return sw.Error()
}
