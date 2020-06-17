package test

import (
	testify "github.com/stretchr/testify/assert"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

var DefaultScheme = runtime.NewScheme()
var DefaultComparator = NewComparator(DefaultScheme)

func RegisterForType(obj runtime.Object, fn CompareFn) {
	DefaultComparator.RegisterForType(obj, fn)
}

func init() {
	core.AddToScheme(DefaultScheme)
	RegisterForType(&core.Secret{}, func(t *testing.T, a, b runtime.Object){
		assert := testify.New(t)
		expected, actual := a.(*core.Secret).DeepCopy(), b.(*core.Secret)
		ObjectCompare_ObjectMeta(t, expected.ObjectMeta, actual.ObjectMeta)
	})
}

func ObjectCompare_ObjectMeta(t *testing.T, expected, actual meta.ObjectMeta) {
	assert := testify.New(t)
	e := expected.DeepCopy()
	e.UID = actual.UID
	e.CreationTimestamp = actual.CreationTimestamp
	e.ResourceVersion = actual.ResourceVersion
	if e.DeletionTimestamp != nil && actual.DeletionTimestamp != nil {
		e.DeletionTimestamp = actual.DeletionTimestamp
	}
	e.Generation = actual.Generation
	assert.EqualValues(e, actual)
}
