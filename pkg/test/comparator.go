package test

import (
	"context"
	"testing"

	testify "github.com/stretchr/testify/assert"
	"go.medium.engineering/kubernetes/pkg/kinds"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type AssertFn func(t *testing.T, a, b runtime.Object)

type TypedAsserts struct {
	Match   AssertFn
	NoMatch AssertFn
}

type Comparator struct {
	typedAsserts map[schema.GroupVersionKind]TypedAsserts
	scheme       *runtime.Scheme
}

func (c *Comparator) RegisterForType(obj runtime.Object, asserts TypedAsserts) {
	gvk := kinds.Identify(c.scheme, obj)
	if gvk.Kind == "" {
		panic("can't identify type")
	}
	c.typedAsserts[gvk] = asserts
}

func (c *Comparator) AssertMatch(
	ctx context.Context,
	t *testing.T,
	cli client.Client,
	expected runtime.Object,
	msgAndArgs ...interface{},
) {
	assert := testify.New(t)
	actual := expected.DeepCopyObject().(client.Object)
	key := client.ObjectKeyFromObject(expected.(client.Object))
	assert.NoError(cli.Get(ctx, key, actual))
	gvk := kinds.Identify(c.scheme, expected)
	assert.NotEmpty(gvk.Kind, "Can't resolve expected value kind")
	asserts, ok := c.typedAsserts[gvk]
	assert.True(ok, "Can't find comparator for kind", gvk)
	asserts.Match(t, expected, actual)
}

func NewComparator(scheme *runtime.Scheme) *Comparator {
	return &Comparator{
		typedAsserts: map[schema.GroupVersionKind]TypedAsserts{},
		scheme:       scheme,
	}
}
