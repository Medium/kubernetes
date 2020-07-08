package v1alpha3

import (
	testing "testing"

	assert "github.com/stretchr/testify/assert"
	test "go.medium.engineering/kubernetes/pkg/test"
	v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	if err := v1alpha3.AddToScheme(test.DefaultScheme); err != nil {
		panic(err)
	}
	RegisterAsserts(test.DefaultComparator)
}

func RegisterAsserts(comparator *test.Comparator) {
	comparator.RegisterForType(&v1alpha3.DestinationRule{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_DestinationRule(t, a.(*v1alpha3.DestinationRule), b.(*v1alpha3.DestinationRule))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_DestinationRule(t, a.(*v1alpha3.DestinationRule), b.(*v1alpha3.DestinationRule))
		},
	})

	comparator.RegisterForType(&v1alpha3.EnvoyFilter{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_EnvoyFilter(t, a.(*v1alpha3.EnvoyFilter), b.(*v1alpha3.EnvoyFilter))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_EnvoyFilter(t, a.(*v1alpha3.EnvoyFilter), b.(*v1alpha3.EnvoyFilter))
		},
	})

	comparator.RegisterForType(&v1alpha3.Gateway{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_Gateway(t, a.(*v1alpha3.Gateway), b.(*v1alpha3.Gateway))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_Gateway(t, a.(*v1alpha3.Gateway), b.(*v1alpha3.Gateway))
		},
	})

	comparator.RegisterForType(&v1alpha3.ServiceEntry{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_ServiceEntry(t, a.(*v1alpha3.ServiceEntry), b.(*v1alpha3.ServiceEntry))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_ServiceEntry(t, a.(*v1alpha3.ServiceEntry), b.(*v1alpha3.ServiceEntry))
		},
	})

	comparator.RegisterForType(&v1alpha3.Sidecar{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_Sidecar(t, a.(*v1alpha3.Sidecar), b.(*v1alpha3.Sidecar))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_Sidecar(t, a.(*v1alpha3.Sidecar), b.(*v1alpha3.Sidecar))
		},
	})

	comparator.RegisterForType(&v1alpha3.VirtualService{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_VirtualService(t, a.(*v1alpha3.VirtualService), b.(*v1alpha3.VirtualService))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_VirtualService(t, a.(*v1alpha3.VirtualService), b.(*v1alpha3.VirtualService))
		},
	})

	comparator.RegisterForType(&v1alpha3.WorkloadEntry{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_WorkloadEntry(t, a.(*v1alpha3.WorkloadEntry), b.(*v1alpha3.WorkloadEntry))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_WorkloadEntry(t, a.(*v1alpha3.WorkloadEntry), b.(*v1alpha3.WorkloadEntry))
		},
	})

}

func Assimilate_DestinationRule(expected, actual *v1alpha3.DestinationRule) *v1alpha3.DestinationRule {
	e := expected.DeepCopyObject().(*v1alpha3.DestinationRule)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_DestinationRule(t *testing.T, expected, actual *v1alpha3.DestinationRule) {
	assert := assert.New(t)
	e := Assimilate_DestinationRule(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_DestinationRule(t *testing.T, expected, actual *v1alpha3.DestinationRule) {
	assert := assert.New(t)
	e := Assimilate_DestinationRule(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_EnvoyFilter(expected, actual *v1alpha3.EnvoyFilter) *v1alpha3.EnvoyFilter {
	e := expected.DeepCopyObject().(*v1alpha3.EnvoyFilter)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_EnvoyFilter(t *testing.T, expected, actual *v1alpha3.EnvoyFilter) {
	assert := assert.New(t)
	e := Assimilate_EnvoyFilter(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_EnvoyFilter(t *testing.T, expected, actual *v1alpha3.EnvoyFilter) {
	assert := assert.New(t)
	e := Assimilate_EnvoyFilter(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_Gateway(expected, actual *v1alpha3.Gateway) *v1alpha3.Gateway {
	e := expected.DeepCopyObject().(*v1alpha3.Gateway)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_Gateway(t *testing.T, expected, actual *v1alpha3.Gateway) {
	assert := assert.New(t)
	e := Assimilate_Gateway(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_Gateway(t *testing.T, expected, actual *v1alpha3.Gateway) {
	assert := assert.New(t)
	e := Assimilate_Gateway(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_ServiceEntry(expected, actual *v1alpha3.ServiceEntry) *v1alpha3.ServiceEntry {
	e := expected.DeepCopyObject().(*v1alpha3.ServiceEntry)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_ServiceEntry(t *testing.T, expected, actual *v1alpha3.ServiceEntry) {
	assert := assert.New(t)
	e := Assimilate_ServiceEntry(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_ServiceEntry(t *testing.T, expected, actual *v1alpha3.ServiceEntry) {
	assert := assert.New(t)
	e := Assimilate_ServiceEntry(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_Sidecar(expected, actual *v1alpha3.Sidecar) *v1alpha3.Sidecar {
	e := expected.DeepCopyObject().(*v1alpha3.Sidecar)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_Sidecar(t *testing.T, expected, actual *v1alpha3.Sidecar) {
	assert := assert.New(t)
	e := Assimilate_Sidecar(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_Sidecar(t *testing.T, expected, actual *v1alpha3.Sidecar) {
	assert := assert.New(t)
	e := Assimilate_Sidecar(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_VirtualService(expected, actual *v1alpha3.VirtualService) *v1alpha3.VirtualService {
	e := expected.DeepCopyObject().(*v1alpha3.VirtualService)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_VirtualService(t *testing.T, expected, actual *v1alpha3.VirtualService) {
	assert := assert.New(t)
	e := Assimilate_VirtualService(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_VirtualService(t *testing.T, expected, actual *v1alpha3.VirtualService) {
	assert := assert.New(t)
	e := Assimilate_VirtualService(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_WorkloadEntry(expected, actual *v1alpha3.WorkloadEntry) *v1alpha3.WorkloadEntry {
	e := expected.DeepCopyObject().(*v1alpha3.WorkloadEntry)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_WorkloadEntry(t *testing.T, expected, actual *v1alpha3.WorkloadEntry) {
	assert := assert.New(t)
	e := Assimilate_WorkloadEntry(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_WorkloadEntry(t *testing.T, expected, actual *v1alpha3.WorkloadEntry) {
	assert := assert.New(t)
	e := Assimilate_WorkloadEntry(expected, actual)
	assert.NotEqualValues(e, actual)
}
