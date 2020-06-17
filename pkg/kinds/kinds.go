package kinds

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func Identify(scheme *runtime.Scheme, obj runtime.Object) schema.GroupVersionKind {
	kind := obj.GetObjectKind()
	gvk := kind.GroupVersionKind()
	if gvk.Kind == "" {
		kinds, _, _ := scheme.ObjectKinds(obj)
		if len(kinds) == 1 {
			gvk = kinds[0]
			kind.SetGroupVersionKind(gvk)
		}
	}
	return gvk
}
