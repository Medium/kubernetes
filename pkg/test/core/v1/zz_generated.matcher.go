package v1

import (
	testing "testing"

	assert "github.com/stretchr/testify/assert"
	test "go.medium.engineering/kubernetes/pkg/test"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	if err := v1.AddToScheme(test.DefaultScheme); err != nil {
		panic(err)
	}
	RegisterAsserts(test.DefaultComparator)
}

func RegisterAsserts(comparator *test.Comparator) {
	comparator.RegisterForType(&v1.Binding{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_Binding(t, a.(*v1.Binding), b.(*v1.Binding))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_Binding(t, a.(*v1.Binding), b.(*v1.Binding))
		},
	})

	comparator.RegisterForType(&v1.ComponentStatus{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_ComponentStatus(t, a.(*v1.ComponentStatus), b.(*v1.ComponentStatus))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_ComponentStatus(t, a.(*v1.ComponentStatus), b.(*v1.ComponentStatus))
		},
	})

	comparator.RegisterForType(&v1.ConfigMap{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_ConfigMap(t, a.(*v1.ConfigMap), b.(*v1.ConfigMap))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_ConfigMap(t, a.(*v1.ConfigMap), b.(*v1.ConfigMap))
		},
	})

	comparator.RegisterForType(&v1.Endpoints{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_Endpoints(t, a.(*v1.Endpoints), b.(*v1.Endpoints))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_Endpoints(t, a.(*v1.Endpoints), b.(*v1.Endpoints))
		},
	})

	comparator.RegisterForType(&v1.Event{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_Event(t, a.(*v1.Event), b.(*v1.Event))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_Event(t, a.(*v1.Event), b.(*v1.Event))
		},
	})

	comparator.RegisterForType(&v1.LimitRange{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_LimitRange(t, a.(*v1.LimitRange), b.(*v1.LimitRange))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_LimitRange(t, a.(*v1.LimitRange), b.(*v1.LimitRange))
		},
	})

	comparator.RegisterForType(&v1.Namespace{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_Namespace(t, a.(*v1.Namespace), b.(*v1.Namespace))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_Namespace(t, a.(*v1.Namespace), b.(*v1.Namespace))
		},
	})

	comparator.RegisterForType(&v1.Node{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_Node(t, a.(*v1.Node), b.(*v1.Node))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_Node(t, a.(*v1.Node), b.(*v1.Node))
		},
	})

	comparator.RegisterForType(&v1.PersistentVolume{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_PersistentVolume(t, a.(*v1.PersistentVolume), b.(*v1.PersistentVolume))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_PersistentVolume(t, a.(*v1.PersistentVolume), b.(*v1.PersistentVolume))
		},
	})

	comparator.RegisterForType(&v1.PersistentVolumeClaim{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_PersistentVolumeClaim(t, a.(*v1.PersistentVolumeClaim), b.(*v1.PersistentVolumeClaim))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_PersistentVolumeClaim(t, a.(*v1.PersistentVolumeClaim), b.(*v1.PersistentVolumeClaim))
		},
	})

	comparator.RegisterForType(&v1.Pod{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_Pod(t, a.(*v1.Pod), b.(*v1.Pod))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_Pod(t, a.(*v1.Pod), b.(*v1.Pod))
		},
	})

	comparator.RegisterForType(&v1.PodStatusResult{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_PodStatusResult(t, a.(*v1.PodStatusResult), b.(*v1.PodStatusResult))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_PodStatusResult(t, a.(*v1.PodStatusResult), b.(*v1.PodStatusResult))
		},
	})

	comparator.RegisterForType(&v1.PodTemplate{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_PodTemplate(t, a.(*v1.PodTemplate), b.(*v1.PodTemplate))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_PodTemplate(t, a.(*v1.PodTemplate), b.(*v1.PodTemplate))
		},
	})

	comparator.RegisterForType(&v1.RangeAllocation{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_RangeAllocation(t, a.(*v1.RangeAllocation), b.(*v1.RangeAllocation))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_RangeAllocation(t, a.(*v1.RangeAllocation), b.(*v1.RangeAllocation))
		},
	})

	comparator.RegisterForType(&v1.ReplicationController{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_ReplicationController(t, a.(*v1.ReplicationController), b.(*v1.ReplicationController))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_ReplicationController(t, a.(*v1.ReplicationController), b.(*v1.ReplicationController))
		},
	})

	comparator.RegisterForType(&v1.ResourceQuota{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_ResourceQuota(t, a.(*v1.ResourceQuota), b.(*v1.ResourceQuota))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_ResourceQuota(t, a.(*v1.ResourceQuota), b.(*v1.ResourceQuota))
		},
	})

	comparator.RegisterForType(&v1.Secret{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_Secret(t, a.(*v1.Secret), b.(*v1.Secret))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_Secret(t, a.(*v1.Secret), b.(*v1.Secret))
		},
	})

	comparator.RegisterForType(&v1.Service{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_Service(t, a.(*v1.Service), b.(*v1.Service))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_Service(t, a.(*v1.Service), b.(*v1.Service))
		},
	})

	comparator.RegisterForType(&v1.ServiceAccount{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_ServiceAccount(t, a.(*v1.ServiceAccount), b.(*v1.ServiceAccount))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_ServiceAccount(t, a.(*v1.ServiceAccount), b.(*v1.ServiceAccount))
		},
	})

}

func Assimilate_Binding(expected, actual *v1.Binding) *v1.Binding {
	e := expected.DeepCopyObject().(*v1.Binding)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_Binding(t *testing.T, expected, actual *v1.Binding) {
	assert := assert.New(t)
	e := Assimilate_Binding(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_Binding(t *testing.T, expected, actual *v1.Binding) {
	assert := assert.New(t)
	e := Assimilate_Binding(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_ComponentStatus(expected, actual *v1.ComponentStatus) *v1.ComponentStatus {
	e := expected.DeepCopyObject().(*v1.ComponentStatus)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_ComponentStatus(t *testing.T, expected, actual *v1.ComponentStatus) {
	assert := assert.New(t)
	e := Assimilate_ComponentStatus(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_ComponentStatus(t *testing.T, expected, actual *v1.ComponentStatus) {
	assert := assert.New(t)
	e := Assimilate_ComponentStatus(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_ConfigMap(expected, actual *v1.ConfigMap) *v1.ConfigMap {
	e := expected.DeepCopyObject().(*v1.ConfigMap)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_ConfigMap(t *testing.T, expected, actual *v1.ConfigMap) {
	assert := assert.New(t)
	e := Assimilate_ConfigMap(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_ConfigMap(t *testing.T, expected, actual *v1.ConfigMap) {
	assert := assert.New(t)
	e := Assimilate_ConfigMap(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_Endpoints(expected, actual *v1.Endpoints) *v1.Endpoints {
	e := expected.DeepCopyObject().(*v1.Endpoints)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_Endpoints(t *testing.T, expected, actual *v1.Endpoints) {
	assert := assert.New(t)
	e := Assimilate_Endpoints(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_Endpoints(t *testing.T, expected, actual *v1.Endpoints) {
	assert := assert.New(t)
	e := Assimilate_Endpoints(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_Event(expected, actual *v1.Event) *v1.Event {
	e := expected.DeepCopyObject().(*v1.Event)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_Event(t *testing.T, expected, actual *v1.Event) {
	assert := assert.New(t)
	e := Assimilate_Event(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_Event(t *testing.T, expected, actual *v1.Event) {
	assert := assert.New(t)
	e := Assimilate_Event(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_LimitRange(expected, actual *v1.LimitRange) *v1.LimitRange {
	e := expected.DeepCopyObject().(*v1.LimitRange)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_LimitRange(t *testing.T, expected, actual *v1.LimitRange) {
	assert := assert.New(t)
	e := Assimilate_LimitRange(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_LimitRange(t *testing.T, expected, actual *v1.LimitRange) {
	assert := assert.New(t)
	e := Assimilate_LimitRange(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_Namespace(expected, actual *v1.Namespace) *v1.Namespace {
	e := expected.DeepCopyObject().(*v1.Namespace)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_Namespace(t *testing.T, expected, actual *v1.Namespace) {
	assert := assert.New(t)
	e := Assimilate_Namespace(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_Namespace(t *testing.T, expected, actual *v1.Namespace) {
	assert := assert.New(t)
	e := Assimilate_Namespace(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_Node(expected, actual *v1.Node) *v1.Node {
	e := expected.DeepCopyObject().(*v1.Node)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_Node(t *testing.T, expected, actual *v1.Node) {
	assert := assert.New(t)
	e := Assimilate_Node(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_Node(t *testing.T, expected, actual *v1.Node) {
	assert := assert.New(t)
	e := Assimilate_Node(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_PersistentVolume(expected, actual *v1.PersistentVolume) *v1.PersistentVolume {
	e := expected.DeepCopyObject().(*v1.PersistentVolume)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_PersistentVolume(t *testing.T, expected, actual *v1.PersistentVolume) {
	assert := assert.New(t)
	e := Assimilate_PersistentVolume(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_PersistentVolume(t *testing.T, expected, actual *v1.PersistentVolume) {
	assert := assert.New(t)
	e := Assimilate_PersistentVolume(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_PersistentVolumeClaim(expected, actual *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim {
	e := expected.DeepCopyObject().(*v1.PersistentVolumeClaim)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_PersistentVolumeClaim(t *testing.T, expected, actual *v1.PersistentVolumeClaim) {
	assert := assert.New(t)
	e := Assimilate_PersistentVolumeClaim(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_PersistentVolumeClaim(t *testing.T, expected, actual *v1.PersistentVolumeClaim) {
	assert := assert.New(t)
	e := Assimilate_PersistentVolumeClaim(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_Pod(expected, actual *v1.Pod) *v1.Pod {
	e := expected.DeepCopyObject().(*v1.Pod)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_Pod(t *testing.T, expected, actual *v1.Pod) {
	assert := assert.New(t)
	e := Assimilate_Pod(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_Pod(t *testing.T, expected, actual *v1.Pod) {
	assert := assert.New(t)
	e := Assimilate_Pod(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_PodStatusResult(expected, actual *v1.PodStatusResult) *v1.PodStatusResult {
	e := expected.DeepCopyObject().(*v1.PodStatusResult)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_PodStatusResult(t *testing.T, expected, actual *v1.PodStatusResult) {
	assert := assert.New(t)
	e := Assimilate_PodStatusResult(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_PodStatusResult(t *testing.T, expected, actual *v1.PodStatusResult) {
	assert := assert.New(t)
	e := Assimilate_PodStatusResult(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_PodTemplate(expected, actual *v1.PodTemplate) *v1.PodTemplate {
	e := expected.DeepCopyObject().(*v1.PodTemplate)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_PodTemplate(t *testing.T, expected, actual *v1.PodTemplate) {
	assert := assert.New(t)
	e := Assimilate_PodTemplate(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_PodTemplate(t *testing.T, expected, actual *v1.PodTemplate) {
	assert := assert.New(t)
	e := Assimilate_PodTemplate(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_RangeAllocation(expected, actual *v1.RangeAllocation) *v1.RangeAllocation {
	e := expected.DeepCopyObject().(*v1.RangeAllocation)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_RangeAllocation(t *testing.T, expected, actual *v1.RangeAllocation) {
	assert := assert.New(t)
	e := Assimilate_RangeAllocation(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_RangeAllocation(t *testing.T, expected, actual *v1.RangeAllocation) {
	assert := assert.New(t)
	e := Assimilate_RangeAllocation(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_ReplicationController(expected, actual *v1.ReplicationController) *v1.ReplicationController {
	e := expected.DeepCopyObject().(*v1.ReplicationController)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_ReplicationController(t *testing.T, expected, actual *v1.ReplicationController) {
	assert := assert.New(t)
	e := Assimilate_ReplicationController(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_ReplicationController(t *testing.T, expected, actual *v1.ReplicationController) {
	assert := assert.New(t)
	e := Assimilate_ReplicationController(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_ResourceQuota(expected, actual *v1.ResourceQuota) *v1.ResourceQuota {
	e := expected.DeepCopyObject().(*v1.ResourceQuota)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_ResourceQuota(t *testing.T, expected, actual *v1.ResourceQuota) {
	assert := assert.New(t)
	e := Assimilate_ResourceQuota(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_ResourceQuota(t *testing.T, expected, actual *v1.ResourceQuota) {
	assert := assert.New(t)
	e := Assimilate_ResourceQuota(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_Secret(expected, actual *v1.Secret) *v1.Secret {
	e := expected.DeepCopyObject().(*v1.Secret)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_Secret(t *testing.T, expected, actual *v1.Secret) {
	assert := assert.New(t)
	e := Assimilate_Secret(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_Secret(t *testing.T, expected, actual *v1.Secret) {
	assert := assert.New(t)
	e := Assimilate_Secret(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_Service(expected, actual *v1.Service) *v1.Service {
	e := expected.DeepCopyObject().(*v1.Service)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_Service(t *testing.T, expected, actual *v1.Service) {
	assert := assert.New(t)
	e := Assimilate_Service(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_Service(t *testing.T, expected, actual *v1.Service) {
	assert := assert.New(t)
	e := Assimilate_Service(expected, actual)
	assert.NotEqualValues(e, actual)
}

func Assimilate_ServiceAccount(expected, actual *v1.ServiceAccount) *v1.ServiceAccount {
	e := expected.DeepCopyObject().(*v1.ServiceAccount)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_ServiceAccount(t *testing.T, expected, actual *v1.ServiceAccount) {
	assert := assert.New(t)
	e := Assimilate_ServiceAccount(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_ServiceAccount(t *testing.T, expected, actual *v1.ServiceAccount) {
	assert := assert.New(t)
	e := Assimilate_ServiceAccount(expected, actual)
	assert.NotEqualValues(e, actual)
}
