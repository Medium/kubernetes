package test

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var DefaultScheme = runtime.NewScheme()
var DefaultComparator = NewComparator(DefaultScheme)

func RegisterForType(obj runtime.Object, asserts TypedAsserts) {
	DefaultComparator.RegisterForType(obj, asserts)
}

func Assimilate_ObjectMeta(expected, actual meta.ObjectMeta) meta.ObjectMeta{
	e := expected.DeepCopy()
	e.UID = actual.UID
	e.CreationTimestamp = actual.CreationTimestamp
	e.ResourceVersion = actual.ResourceVersion
	if e.DeletionTimestamp != nil && actual.DeletionTimestamp != nil {
		e.DeletionTimestamp = actual.DeletionTimestamp
	}
	e.Generation = actual.Generation
	return *e
}
