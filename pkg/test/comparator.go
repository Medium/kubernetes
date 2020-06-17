package test

import (
	"context"
	testify "github.com/stretchr/testify/assert"
	"go.medium.engineering/kubernetes/pkg/kinds"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"testing"
)

type CompareFn func(t *testing.T, a, b runtime.Object)

type Comparator struct {
	compareFns map[schema.GroupVersionKind]CompareFn
	scheme *runtime.Scheme
}

func (c *Comparator) RegisterForType(obj runtime.Object, fn CompareFn) {
	gvk := kinds.Identify(c.scheme, obj)
	if gvk.Kind == "" {
		panic("can't identify type")
	}
	c.compareFns[gvk] = fn
}

func (c *Comparator) AssertMatch(
	ctx context.Context,
	t *testing.T,
	cli client.Client,
	expected runtime.Object,
	msgAndArgs ...interface{},
) {
	assert := testify.New(t)
	actual := expected.DeepCopyObject()
	key, err := client.ObjectKeyFromObject(expected)
	assert.NoError(err, msgAndArgs...)
	assert.NoError(cli.Get(ctx, key, actual))
	gvk := kinds.Identify(c.scheme, expected)
	assert.NotEmpty(gvk.Kind, "Can't resolve expected value kind")
	fn, ok := c.compareFns[gvk]
	assert.True(ok, "Can't find comparator for kind", gvk)
	fn(t, expected, actual)
}

func NewComparator(scheme *runtime.Scheme) *Comparator {
	return &Comparator{
		compareFns: map[schema.GroupVersionKind]CompareFn{},
		scheme: scheme,
	}
}

