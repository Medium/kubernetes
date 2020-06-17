package generators

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"reflect"
	"strings"

	"k8s.io/gengo/args"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"

	"k8s.io/klog/v2"
)

// CustomArgs is used tby the go2idl framework to pass args specific to this
// generator.
type CustomArgs struct {
	ExtraPeerDirs []string // Always consider these as last-ditch possibilities for conversions.
}

// These are the comment tags that carry parameters for comparator generation.
const tagName = "k8s:comparator-gen"
const intputTagName = "k8s:comparator-gen-input"

func extractTag(comments []string) []string {
	return types.ExtractCommentTags("+", comments)[tagName]
}

func extractInputTag(comments []string) []string {
	return types.ExtractCommentTags("+", comments)[intputTagName]
}

func checkTag(comments []string, require ...string) bool {
	values := types.ExtractCommentTags("+", comments)[tagName]
	if len(require) == 0 {
		return len(values) == 1 && values[0] == ""
	}
	return reflect.DeepEqual(values, require)
}

func comparatorFnNamer() *namer.NameStrategy {
	return &namer.NameStrategy{
		Prefix: "Compare_",
		Join: func(pre string, in []string, post string) string {
			return pre + strings.Join(in, "_") + post
		},
	}
}

func objectComparatorFnNamer() *namer.NameStrategy {
	return &namer.NameStrategy{
		Prefix: "ObjectCompare_",
		Join: func(pre string, in []string, post string) string {
			return pre + strings.Join(in, "_") + post
		},
	}
}

// NameSystems returns the name system used by the generators in this package.
func NameSystems() namer.NameSystems {
	return namer.NameSystems{
		"public":          namer.NewPublicNamer(1),
		"raw":             namer.NewRawNamer("", nil),
		"comparatorfn":       comparatorFnNamer(),
		"objectcomparatorfn": objectComparatorFnNamer(),
	}
}

// NameSystem returns the default name system for ordering the types to be
// processed by the generators in this package.
func DefaultNameSystem() string {
	return "public"
}

// comparators holds the declared comparing functions for a given type (all comparing functions
// are expected to be func(1))
type comparators struct {
	// object is the comparator function for a top level type (typically one with TypeMeta) that
	// invokes all child comparators. May be nil if the object comparator has not yet been generated.
	object *types.Type
	// base is a comparator function defined for a type Sets_Pod which does not invoke all
	// child comparators - the base comparator alone is insufficient to compare a type
	base *types.Type
	// additional is zero or more comparator functions of the form Sets_Pod_XXXX that can be
	// included in the Object comparator.
	additional []*types.Type
}

// All of the types in conversions map are of type "DeclarationOf" with
// the underlying type being "Func".
type comparatorFuncMap map[*types.Type]comparators

// Returns all manually-defined comparing functions in the package.
func getManualingFunctions(context *generator.Context, pkg *types.Package, manualMap comparatorFuncMap) {
	buffer := &bytes.Buffer{}
	sw := generator.NewSnippetWriter(buffer, context, "$", "$")

	for _, f := range pkg.Functions {
		if f.Underlying == nil || f.Underlying.Kind != types.Func {
			klog.Errorf("Malformed function: %#v", f)
			continue
		}
		if f.Underlying.Signature == nil {
			klog.Errorf("Function without signature: %#v", f)
			continue
		}
		signature := f.Underlying.Signature
		// Check whether the function is comparing function.
		// Note that all of them have signature:
		// object: func SetObjects_inType(*inType)
		// base: func Sets_inType(*inType)
		// additional: func Sets_inType_Qualifier(*inType)
		if signature.Receiver != nil {
			continue
		}
		if len(signature.Parameters) != 1 {
			continue
		}
		if len(signature.Results) != 0 {
			continue
		}
		inType := signature.Parameters[0]
		if inType.Kind != types.Pointer {
			continue
		}
		// Check if this is the primary comparator.
		cargs := comparingArgsFromType(inType.Elem)
		sw.Do("$.inType|comparatorfn$", cargs)
		switch {
		case f.Name.Name == buffer.String():
			key := inType.Elem
			// We might scan the same package twice, and that's OK.
			v, ok := manualMap[key]
			if ok && v.base != nil && v.base.Name.Package != pkg.Path {
				panic(fmt.Sprintf("duplicate static comparator defined: %#v", key))
			}
			v.base = f
			manualMap[key] = v
			klog.V(6).Infof("found base comparator function for %s from %s", key.Name, f.Name)
		// Is one of the additional comparators - a top level comparator on a type that is
		// also invoked.
		case strings.HasPrefix(f.Name.Name, buffer.String()+"_"):
			key := inType.Elem
			v, ok := manualMap[key]
			if ok {
				exists := false
				for _, existing := range v.additional {
					if existing.Name == f.Name {
						exists = true
						break
					}
				}
				if exists {
					continue
				}
			}
			v.additional = append(v.additional, f)
			manualMap[key] = v
			klog.V(6).Infof("found additional comparator function for %s from %s", key.Name, f.Name)
		}
		buffer.Reset()
		sw.Do("$.inType|objectcomparatorfn$", cargs)
		if f.Name.Name == buffer.String() {
			key := inType.Elem
			// We might scan the same package twice, and that's OK.
			v, ok := manualMap[key]
			if ok && v.base != nil && v.base.Name.Package != pkg.Path {
				panic(fmt.Sprintf("duplicate static comparator defined: %#v", key))
			}
			v.object = f
			manualMap[key] = v
			klog.V(6).Infof("found object comparator function for %s from %s", key.Name, f.Name)
		}
		buffer.Reset()
	}
}

func Packages(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
	boilerplate, err := arguments.LoadGoBoilerplate()
	if err != nil {
		klog.Fatalf("Failed loading boilerplate: %v", err)
	}

	packages := generator.Packages{}
	header := append([]byte(fmt.Sprintf("// +build !%s\n\n", arguments.GeneratedBuildTag)), boilerplate...)

	// Accumulate pre-existing comparing functions.
	// TODO: This is too ad-hoc.  We need a better way.
	existingComparators := comparatorFuncMap{}

	buffer := &bytes.Buffer{}
	sw := generator.NewSnippetWriter(buffer, context, "$", "$")

	// We are generating comparators only for packages that are explicitly
	// passed as InputDir.
	for _, i := range context.Inputs {
		klog.V(5).Infof("considering pkg %q", i)
		pkg := context.Universe[i]
		if pkg == nil {
			// If the input had no Go files, for example.
			continue
		}
		// typesPkg is where the types that needs comparator are defined.
		// Sometimes it is different from pkg. For example, kubernetes core/v1
		// types are defined in vendor/k8s.io/api/core/v1, while pkg is at
		// pkg/api/v1.
		typesPkg := pkg

		// Add comparing functions.
		getManualingFunctions(context, pkg, existingComparators)

		var peerPkgs []string
		if customArgs, ok := arguments.CustomArgs.(*CustomArgs); ok {
			for _, pkg := range customArgs.ExtraPeerDirs {
				if i := strings.Index(pkg, "/vendor/"); i != -1 {
					pkg = pkg[i+len("/vendor/"):]
				}
				peerPkgs = append(peerPkgs, pkg)
			}
		}
		// Make sure our peer-packages are added and fully parsed.
		for _, pp := range peerPkgs {
			if _, err := context.AddDirectory(pp); err != nil {
				klog.Fatalf("Failed to add directory %#v", pp)
			}
			getManualingFunctions(context, context.Universe[pp], existingComparators)
		}

		shouldCreateObjecterFn := func(t *types.Type) bool {
			if comparers, ok := existingComparators[t]; ok && comparers.object != nil {
				// A comparator generator is defined
				baseTypeName := "<unknown>"
				if comparers.base != nil {
					baseTypeName = comparers.base.Name.String()
				}
				klog.V(5).Infof("  an object comparator already exists as %s", baseTypeName)
				return false
			}
			// opt-out
			if checkTag(t.SecondClosestCommentLines, "false") {
				klog.V(5).Infof("opt-out")
				return false
			}
			// opt-in
			if checkTag(t.SecondClosestCommentLines, "true") {
				klog.V(5).Infof("opt-in")
				return true
			}
			// For every k8s:comparator-gen tag at the package level, interpret the value as a
			// field name (like TypeMeta, ListMeta, ObjectMeta) and trigger comparator generation
			// for any type with any of the matching field names. Provides a more useful package
			// level defaulting than global (because we only need comparators on a subset of objects -
			// usually those with TypeMeta).
			if t.Kind == types.Struct {
				klog.V(5).Info(t.Name)
				for _, member := range t.Members {
					klog.V(5).Infof(" - '%s'", member.Name)
					if member.Name == "ObjectMeta" {
						return true
					}
				}
			}
			return false
		}

		// if the types are not in the same package where the comparator functions to be generated
		inputTags := extractInputTag(pkg.Comments)
		if len(inputTags) > 1 {
			panic(fmt.Sprintf("there could only be one input tag, got %#v", inputTags))
		}
		if len(inputTags) == 1 {
			var err error
			typesPkg, err = context.AddDirectory(filepath.Join(pkg.Path, inputTags[0]))
			if err != nil {
				klog.Fatalf("cannot import package %s", inputTags[0])
			}
			// update context.Order to the latest context.Universe
			orderer := namer.Orderer{Namer: namer.NewPublicNamer(1)}
			context.Order = orderer.OrderUniverse(context.Universe)
		}

		newComparators := comparatorFuncMap{}
		for _, t := range typesPkg.Types {
			if t.Name.Name != "Secret" {
				continue
			}
			if !shouldCreateObjecterFn(t) {
				continue
			}
			if namer.IsPrivateGoName(t.Name.Name) {
				// We won't be able to convert to a private type.
				klog.V(5).Infof("  found a type %v, but it is a private name", t)
				continue
			}

			// create a synthetic type we can use during generation
			newComparators[t] = comparators{}
		}

		// only generate comparators for objects that actually have defined comparators
		// prevents empty comparators from being registered
		for {
			promoted := 0
			for t, d := range newComparators {
				if d.object != nil {
					continue
				}
				newCallTreeForType(existingComparators, newComparators).build(t, true)
				cargs := comparingArgsFromType(t)
				sw.Do("$.inType|objectcomparatorfn$", cargs)
				newComparators[t] = comparators{
					object: &types.Type{
						Name: types.Name{
							Package: pkg.Path,
							Name:    buffer.String(),
						},
						Kind: types.Func,
					},
				}
				buffer.Reset()
				promoted++
			}
			if promoted != 0 {
				continue
			}

			// prune any types that were not used
			for t, d := range newComparators {
				if d.object == nil {
					klog.V(6).Infof("did not generate comparator for %s because no child comparators were registered", t.Name)
					delete(newComparators, t)
				}
			}
			break
		}

		if len(newComparators) == 0 {
			klog.V(5).Infof("no comparators in package %s", pkg.Name)
		}

		path := pkg.Path
		// if the source path is within a /vendor/ directory (for example,
		// k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1), allow
		// generation to output to the proper relative path (under vendor).
		// Otherwise, the generator will create the file in the wrong location
		// in the output directory.
		// TODO: build a more fundamental concept in gengo for dealing with modifications
		// to vendored packages.
		if strings.HasPrefix(pkg.SourcePath, arguments.OutputBase) {
			expandedPath := strings.TrimPrefix(pkg.SourcePath, arguments.OutputBase)
			if strings.Contains(expandedPath, "/vendor/") {
				path = expandedPath
			}
		}

		packages = append(packages,
			&generator.DefaultPackage{
				PackageName: filepath.Base(pkg.Path),
				PackagePath: path,
				HeaderText:  header,
				GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
					return []generator.Generator{
						NewComparator(arguments.OutputFileBaseName, typesPkg.Path, pkg.Path, existingComparators, newComparators, peerPkgs),
					}
				},
				FilterFunc: func(c *generator.Context, t *types.Type) bool {
					return t.Name.Package == typesPkg.Path
				},
			})
	}
	return packages
}

// callTreeForType contains fields necessary to build a tree for types.
type callTreeForType struct {
	existingComparators     comparatorFuncMap
	newComparators          comparatorFuncMap
	currentlyBuildingTypes map[*types.Type]bool
}

func newCallTreeForType(existingComparators, newComparators comparatorFuncMap) *callTreeForType {
	return &callTreeForType{
		existingComparators:     existingComparators,
		newComparators:          newComparators,
		currentlyBuildingTypes: make(map[*types.Type]bool),
	}
}

// build creates a tree of paths to fields (based on how they would be accessed in Go - pointer, elem,
// slice, or key) and the functions that should be invoked on each field. An in-order traversal of the resulting tree
// can be used to generate a Go function that invokes each nested function on the appropriate type. The return
// value may be nil if there are no functions to call on type or the type is a primitive (ers can only be
// invoked on structs today). When root is true this function will not use a newer. existingComparators should
// contain all comparing functions by type defined in code - newComparators should contain all object comparators
// that could be or will be generated. If newComparators has an entry for a type, but the 'object' field is nil,
// this function skips adding that comparator - this allows us to avoid generating object comparator functions for
// list types that call empty comparators.
func (c *callTreeForType) build(t *types.Type, root bool) *callNode {
	parent := &callNode{}

	if root {
		// the root node is always a pointer
		parent.elem = true
	}

	comparators, _ := c.existingComparators[t]
	news, generated := c.newComparators[t]
	switch {
	case !root && generated && news.object != nil:
		parent.call = append(parent.call, news.object)
		// if we will be generating the comparator, it by definition is a covering
		// comparator, so we halt recursion
		klog.V(6).Infof("the comparator %s will be generated as an object comparator", t.Name)
		return parent

	case comparators.object != nil:
		// object comparators are always covering
		parent.call = append(parent.call, comparators.object)
		return parent

	case comparators.base != nil:
		parent.call = append(parent.call, comparators.base)
		// if the base function indicates it "covers" (it already includes comparators)
		// we can halt recursion
		if checkTag(comparators.base.CommentLines, "covers") {
			klog.V(6).Infof("the comparator %s indicates it covers all sub generators", t.Name)
			return parent
		}
	}

	// base has been added already, now add any additional comparators defined for this object
	parent.call = append(parent.call, comparators.additional...)

	// if the type already exists, don't build the tree for it and don't generate anything.
	// This is used to avoid recursion for nested recursive types.
	if c.currentlyBuildingTypes[t] {
		return nil
	}
	// if type doesn't exist, mark it as existing
	c.currentlyBuildingTypes[t] = true

	defer func() {
		// The type will now acts as a parent, not a nested recursive type.
		// We can now build the tree for it safely.
		c.currentlyBuildingTypes[t] = false
	}()

	switch t.Kind {
	case types.Pointer:
		if child := c.build(t.Elem, false); child != nil {
			child.elem = true
			parent.children = append(parent.children, *child)
		}
	case types.Slice, types.Array:
		if child := c.build(t.Elem, false); child != nil {
			child.index = true
			if t.Elem.Kind == types.Pointer {
				child.elem = true
			}
			parent.children = append(parent.children, *child)
		}
	case types.Map:
		if child := c.build(t.Elem, false); child != nil {
			child.key = true
			parent.children = append(parent.children, *child)
		}
	case types.Struct:
		for _, field := range t.Members {
			name := field.Name
			if len(name) == 0 {
				if field.Type.Kind == types.Pointer {
					name = field.Type.Elem.Name.Name
				} else {
					name = field.Type.Name.Name
				}
			}
			if child := c.build(field.Type, false); child != nil {
				child.field = name
				parent.children = append(parent.children, *child)
			}
		}
	case types.Alias:
		if child := c.build(t.Underlying, false); child != nil {
			parent.children = append(parent.children, *child)
		}
	}
	if len(parent.children) == 0 && len(parent.call) == 0 {
		klog.V(6).Infof("decided type %s needs no generation", t.Name)
		return nil
	}
	return parent
}

const (
	comparatorPackagePath    = "go.medium.engineering/kubernetes/pkg/test"
)

// gener produces a file with a autogenerated conversions.
type genComparator struct {
	generator.DefaultGen
	typesPackage       string
	outputPackage      string
	peerPackages       []string
	newComparators      comparatorFuncMap
	existingComparators comparatorFuncMap
	imports            namer.ImportTracker
	typesForInit       []*types.Type
}

func NewComparator(sanitizedName, typesPackage, outputPackage string, existingComparators, newComparators comparatorFuncMap, peerPkgs []string) generator.Generator {
	return &genComparator{
		DefaultGen: generator.DefaultGen{
			OptionalName: sanitizedName,
		},
		typesPackage:       typesPackage,
		outputPackage:      outputPackage,
		peerPackages:       peerPkgs,
		newComparators:      newComparators,
		existingComparators: existingComparators,
		imports:            generator.NewImportTracker(),
		typesForInit:       make([]*types.Type, 0),
	}
}

func (g *genComparator) Namers(_ *generator.Context) namer.NameSystems {
	// Have the raw namer for this file track what it imports.
	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.outputPackage, g.imports),
	}
}

func (g *genComparator) isOtherPackage(pkg string) bool {
	if pkg == g.outputPackage {
		return false
	}
	if strings.HasSuffix(pkg, `"`+g.outputPackage+`"`) {
		return false
	}
	return true
}

func (g *genComparator) Filter(_ *generator.Context, t *types.Type) bool {
	comparers, ok := g.newComparators[t]
	if !ok || comparers.object == nil {
		return false
	}
	g.typesForInit = append(g.typesForInit, t)
	return true
}

func (g *genComparator) Imports(_ *generator.Context) (imports []string) {
	var importLines []string
	for _, singleImport := range g.imports.ImportLines() {
		if g.isOtherPackage(singleImport) {
			importLines = append(importLines, singleImport)
		}
	}
	return importLines
}

func (g *genComparator) Init(c *generator.Context, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, c, "$", "$")

	testing := types.Ref("testing", "T")
	g.imports.AddType(testing)
	objectMeta := types.Ref("k8s.io/apimachinery/pkg/apis/meta/v1", "ObjectMeta")
	g.imports.AddType(objectMeta)
	object := types.Ref("k8s.io/apimachinery/pkg/runtime", "Object")

	// default args
	dargs := generator.Args{
		"testing": testing,
		"objectMeta": objectMeta,
		"object": object,
	}

	comparator := c.Universe.Type(types.Name{Package: comparatorPackagePath, Name: "Comparator"})
	comparatorPtr := &types.Type{
		Kind: types.Pointer,
		Elem: comparator,
	}
	sw.Do("import \"go.medium.engineering/kubernetes/pkg/test\"", nil)
	sw.Do("// RegisterComparators adds comparator functions to the given comparator.\n", nil)
	sw.Do("// Public to allow building arbitrary schemes.\n", nil)
	sw.Do("func Registers(comparator $.|raw$) error {\n", comparatorPtr)
	for _, t := range g.typesForInit {
		cargs := comparingArgsFromType(t)
		cargs = cargs.WithArgs(dargs)
		sw.Do("comparator.RegisterForType(&$.inType|raw${}, func(t *$.testing|raw$, e, a $.object|raw$) {\n", cargs)
		sw.Do("  $.inType|objectcomparatorfn$(t, e.(*$.inType|raw$), a.(*$.inType|raw$))\n", cargs)
		sw.Do("})\n\n", nil)
	}
	sw.Do("return nil\n", nil)
	sw.Do("}\n\n", nil)
	for _, t := range g.typesForInit {
		cargs := comparingArgsFromType(t)
		cargs = cargs.WithArgs(dargs)
		sw.Do("func $.inType|objectcomparatorfn$(t *$.testing|raw$, e, a *$.inType|raw$) {\n", cargs)
		sw.Do("test.ObjectCompare_ObjectMeta(t, e.ObjectMeta, a.ObjectMeta)\n", nil)
		sw.Do("}\n\n", nil)
	}
	return sw.Error()
}

func (g *genComparator) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	if _, ok := g.newComparators[t]; !ok {
		return nil
	}

	klog.V(5).Infof("generating for type %v", t)

	callTree := newCallTreeForType(g.existingComparators, g.newComparators).build(t, true)
	if callTree == nil {
		klog.V(5).Infof("  no comparators defined")
		return nil
	}
	i := 0
	callTree.VisitInOrder(func(ancestors []*callNode, current *callNode) {
		if len(current.call) == 0 {
			return
		}
		path := callPath(append(ancestors, current))
		klog.V(5).Infof("  %d: %s", i, path)
		i++
	})

	sw := generator.NewSnippetWriter(w, c, "$", "$")
	g.generateComparator(t, callTree, sw)
	return sw.Error()
}

func comparingArgsFromType(inType *types.Type) generator.Args {
	return generator.Args{
		"inType": inType,
	}
}

func (g *genComparator) generateComparator(inType *types.Type, callTree *callNode, sw *generator.SnippetWriter) {
	sw.Do("func $.inType|objectcomparatorfn$(in *$.inType|raw$) {\n", comparingArgsFromType(inType))
	callTree.WriteMethod("in", 0, nil, sw)
	sw.Do("}\n\n", nil)
}

// callNode represents an entry in a tree of Go type accessors - the path from the root to a leaf represents
// how in Go code an access would be performed. For example, if a comparing function exists on a container
// lifecycle hook, to invoke that comparator correctly would require this Go code:
//
//     for i := range pod.Spec.Containers {
//       o := &pod.Spec.Containers[i]
//       if o.LifecycleHook != nil {
//         Sets_LifecycleHook(o.LifecycleHook)
//       }
//     }
//
// That would be represented by a call tree like:
//
//   callNode
//     field: "Spec"
//     children:
//     - field: "Containers"
//       children:
//       - index: true
//         children:
//         - field: "LifecycleHook"
//           elem: true
//           call:
//           - Sets_LifecycleHook
//
// which we can traverse to build that Go struct (you must call the field Spec, then Containers, then range over
// that field, then check whether the LifecycleHook field is nil, before calling Sets_LifecycleHook on
// the pointer to that field).
type callNode struct {
	// field is the name of the Go member to access
	field string
	// key is true if this is a map and we must range over the key and values
	key bool
	// index is true if this is a slice and we must range over the slice values
	index bool
	// elem is true if the previous elements refer to a pointer (typically just field)
	elem bool

	// call is all of the functions that must be invoked on this particular node, in order
	call []*types.Type
	// children is the child call nodes that must also be traversed
	children []callNode
}

// CallNodeVisitorFunc is a function for visiting a call tree. ancestors is the list of all parents
// of this node to the root of the tree - will be empty at the root.
type CallNodeVisitorFunc func(ancestors []*callNode, node *callNode)

func (n *callNode) VisitInOrder(fn CallNodeVisitorFunc) {
	n.visitInOrder(nil, fn)
}

func (n *callNode) visitInOrder(ancestors []*callNode, fn CallNodeVisitorFunc) {
	fn(ancestors, n)
	ancestors = append(ancestors, n)
	for i := range n.children {
		n.children[i].visitInOrder(ancestors, fn)
	}
}

var (
	indexVariables = "ijklmnop"
	localVariables = "abcdefgh"
)

// varsForDepth creates temporary variables guaranteed to be unique within lexical Go scopes
// of this depth in a function. It uses canonical Go loop variables for the first 7 levels
// and then resorts to uglier prefixes.
func varsForDepth(depth int) (index, local string) {
	if depth > len(indexVariables) {
		index = fmt.Sprintf("i%d", depth)
	} else {
		index = indexVariables[depth : depth+1]
	}
	if depth > len(localVariables) {
		local = fmt.Sprintf("local%d", depth)
	} else {
		local = localVariables[depth : depth+1]
	}
	return
}

// writeCalls generates a list of function calls based on the calls field for the provided variable
// name and pointer.
func (n *callNode) writeCalls(varName string, isVarPointer bool, sw *generator.SnippetWriter) {
	accessor := varName
	if !isVarPointer {
		accessor = "&" + accessor
	}
	for _, fn := range n.call {
		sw.Do("$.fn|raw$($.var$)\n", generator.Args{
			"fn":  fn,
			"var": accessor,
		})
	}
}

// WriteMethod performs an in-order traversal of the calltree, generating loops and if blocks as necessary
// to correctly turn the call tree into a method body that invokes all calls on all child nodes of the call tree.
// Depth is used to generate local variables at the proper depth.
func (n *callNode) WriteMethod(varName string, depth int, ancestors []*callNode, sw *generator.SnippetWriter) {
	// if len(n.call) > 0 {
	// 	sw.Do(fmt.Sprintf("// %s\n", callPath(append(ancestors, n)).String()), nil)
	// }

	if len(n.field) > 0 {
		varName = varName + "." + n.field
	}

	index, local := varsForDepth(depth)
	vars := generator.Args{
		"index": index,
		"local": local,
		"var":   varName,
	}

	isPointer := n.elem && !n.index
	if isPointer && len(ancestors) > 0 {
		sw.Do("if $.var$ != nil {\n", vars)
	}

	switch {
	case n.index:
		sw.Do("for $.index$ := range $.var$ {\n", vars)
		if n.elem {
			sw.Do("$.local$ := $.var$[$.index$]\n", vars)
		} else {
			sw.Do("$.local$ := &$.var$[$.index$]\n", vars)
		}

		n.writeCalls(local, true, sw)
		for i := range n.children {
			n.children[i].WriteMethod(local, depth+1, append(ancestors, n), sw)
		}
		sw.Do("}\n", nil)
	case n.key:
	default:
		n.writeCalls(varName, isPointer, sw)
		for i := range n.children {
			n.children[i].WriteMethod(varName, depth, append(ancestors, n), sw)
		}
	}

	if isPointer && len(ancestors) > 0 {
		sw.Do("}\n", nil)
	}
}

type callPath []*callNode

// String prints a representation of a callPath that roughly approximates what a Go accessor
// would look like. Used for debugging only.
func (path callPath) String() string {
	if len(path) == 0 {
		return "<none>"
	}
	var parts []string
	for _, p := range path {
		last := len(parts) - 1
		switch {
		case p.elem:
			if len(parts) > 0 {
				parts[last] = "*" + parts[last]
			} else {
				parts = append(parts, "*")
			}
		case p.index:
			if len(parts) > 0 {
				parts[last] = parts[last] + "[i]"
			} else {
				parts = append(parts, "[i]")
			}
		case p.key:
			if len(parts) > 0 {
				parts[last] = parts[last] + "[key]"
			} else {
				parts = append(parts, "[key]")
			}
		default:
			if len(p.field) > 0 {
				parts = append(parts, p.field)
			} else {
				parts = append(parts, "<root>")
			}
		}
	}
	var calls []string
	for _, fn := range path[len(path)-1].call {
		calls = append(calls, fn.Name.String())
	}
	if len(calls) == 0 {
		calls = append(calls, "<none>")
	}

	return strings.Join(parts, ".") + " calls " + strings.Join(calls, ", ")
}