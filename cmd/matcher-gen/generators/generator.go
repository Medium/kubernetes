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

func matchFnNamer() *namer.NameStrategy {
	return &namer.NameStrategy{
		Prefix:              "Match_",
		Join:                func(pre string, in []string, post string) string {
			return pre + strings.Join(in, "_") + post
		},
	}
}

func noMatchFnNamer() *namer.NameStrategy {
	return &namer.NameStrategy{
		Prefix:              "NoMatch_",
		Join:                func(pre string, in []string, post string) string {
			return pre + strings.Join(in, "_") + post
		},
	}
}

func assimilateFnNamer() *namer.NameStrategy {
	return &namer.NameStrategy{
		Prefix:              "Assimilate_",
		Join:                func(pre string, in []string, post string) string {
			return pre + strings.Join(in, "_") + post
		},
	}
}

func (g *gen) Namers(_ *generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.outputPackage, g.imports),
		"matchFn": matchFnNamer(),
		"noMatchFn": noMatchFnNamer(),
		"assimilateFn": assimilateFnNamer(),
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

func (g *gen) Imports(_ *generator.Context) (imports []string) {
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
		"defaultScheme": c.Universe.Type(types.Name{Package: "go.medium.engineering/kubernetes/pkg/test", Name: "DefaultScheme"}),
		"object": c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/runtime", Name: "Object"}),
		"assimilateObjectMeta": c.Universe.Type(types.Name{Package: "go.medium.engineering/kubernetes/pkg/test", Name: "Assimilate_ObjectMeta"}),
		"typedAsserts": c.Universe.Type(types.Name{Package: "go.medium.engineering/kubernetes/pkg/test", Name: "TypedAsserts"}),
		"addToScheme": c.Universe.Type(types.Name{Package: g.typesPackage, Name: "AddToScheme"}),
	}

	sw.Do("func init() {\n", nil)
	sw.Do("if err := $.addToScheme|raw$($.defaultScheme|raw$); err != nil { panic(err) }\n", args)
	sw.Do("RegisterAsserts($.defaultComparator|raw$)\n", args)
	sw.Do("}\n\n", nil)

	sw.Do("func RegisterAsserts(comparator *$.comparator|raw$) {\n", args)
	for _, t := range g.types {
		args["t"] = t

		sw.Do("comparator.RegisterForType(&$.t|raw${}, $.typedAsserts|raw${\n", args)
		sw.Do("Match: func(t *$.test|raw$, a, b $.object|raw$) {\n", args)
		sw.Do("$.t|matchFn$(t, a.(*$.t|raw$), b.(*$.t|raw$))\n", args)
		sw.Do("},\n", nil)
		sw.Do("NoMatch: func(t *$.test|raw$, a, b $.object|raw$) {\n", args)
		sw.Do("$.t|noMatchFn$(t, a.(*$.t|raw$), b.(*$.t|raw$))\n", args)
		sw.Do("},\n", nil)
		sw.Do("})\n\n", nil)
	}
	sw.Do("}\n\n", nil)

	for _, t := range g.types {
		args["t"] = t

		sw.Do("func $.t|assimilateFn$(expected, actual *$.t|raw$) *$.t|raw$ {\n", args)
		sw.Do("e := expected.DeepCopyObject().(*$.t|raw$)\n", args)
		sw.Do("e.ObjectMeta = $.assimilateObjectMeta|raw$(e.ObjectMeta, actual.ObjectMeta)\n", args)
		sw.Do("return e\n", nil)
		sw.Do("}\n\n", nil)

		sw.Do("func $.t|matchFn$(t *$.test|raw$, expected, actual *$.t|raw$) {\n", args)
		sw.Do("assert := $.testify|raw$(t)\n", args)
		sw.Do("e := $.t|assimilateFn$(expected, actual)\n", args)
		sw.Do("assert.EqualValues(e, actual)\n", nil)
		sw.Do("}\n\n", nil)

		sw.Do("func $.t|noMatchFn$(t *$.test|raw$, expected, actual *$.t|raw$) {\n", args)
		sw.Do("assert := $.testify|raw$(t)\n", args)
		sw.Do("e := $.t|assimilateFn$(expected, actual)\n", args)
		sw.Do("assert.NotEqualValues(e, actual)\n", nil)
		sw.Do("}\n\n", nil)
	}

	return sw.Error()
}
