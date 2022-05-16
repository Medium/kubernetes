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

type AssertFn func(t *testing.T, a, b client.Object)

type TypedAsserts struct {
	Match   AssertFn
	NoMatch AssertFn
}

type Comparator struct {
	typedAsserts map[schema.GroupVersionKind]TypedAsserts
	scheme       *runtime.Scheme
}

func (c *Comparator) RegisterForType(obj client.Object, asserts TypedAsserts) {
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
	expected client.Object,
	msgAndArgs ...interface{},
) {
	assert := testify.New(t)
	actual := expected.DeepCopyObject().(client.Object)
	CopyMeta(actual, expected)
	key := client.ObjectKeyFromObject(expected)
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

func CopyMeta(obj1, obj2 client.Object) {
	obj2.SetAnnotations(obj1.GetAnnotations())
	obj2.SetClusterName(obj1.GetClusterName())
	obj2.SetCreationTimestamp(obj1.GetCreationTimestamp())
	obj2.SetDeletionGracePeriodSeconds(obj1.GetDeletionGracePeriodSeconds())
	obj2.SetDeletionTimestamp(obj1.GetDeletionTimestamp())
	obj2.SetFinalizers(obj1.GetFinalizers())
	obj2.SetGenerateName(obj1.GetGenerateName())
	obj2.SetGeneration(obj1.GetGeneration())
	obj2.SetLabels(obj1.GetLabels())
	obj2.SetManagedFields(obj1.GetManagedFields())
	obj2.SetName(obj1.GetName())
	obj2.SetNamespace(obj1.GetNamespace())
	obj2.SetOwnerReferences(obj1.GetOwnerReferences())
	obj2.SetResourceVersion(obj1.GetResourceVersion())
	obj2.SetSelfLink(obj1.GetSelfLink())
	obj2.SetUID(obj1.GetUID())
}
