package v1

import (
	testing "testing"

	assert "github.com/stretchr/testify/assert"
	test "go.medium.engineering/kubernetes/pkg/test"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	RegisterComparators(test.DefaultComparator)
}

func RegisterComparators(comparator *test.Comparator) {
	comparator.RegisterForType(&v1.ResourceQuota{}, func(t *testing.T, a, b runtime.Object) {
		Compare_ResourceQuota(t, a.(*v1.ResourceQuota), b.(*v1.ResourceQuota))
	})

	comparator.RegisterForType(&v1.Secret{}, func(t *testing.T, a, b runtime.Object) {
		Compare_Secret(t, a.(*v1.Secret), b.(*v1.Secret))
	})

	comparator.RegisterForType(&v1.ReplicationController{}, func(t *testing.T, a, b runtime.Object) {
		Compare_ReplicationController(t, a.(*v1.ReplicationController), b.(*v1.ReplicationController))
	})

	comparator.RegisterForType(&v1.ComponentStatus{}, func(t *testing.T, a, b runtime.Object) {
		Compare_ComponentStatus(t, a.(*v1.ComponentStatus), b.(*v1.ComponentStatus))
	})

	comparator.RegisterForType(&v1.PodTemplate{}, func(t *testing.T, a, b runtime.Object) {
		Compare_PodTemplate(t, a.(*v1.PodTemplate), b.(*v1.PodTemplate))
	})

	comparator.RegisterForType(&v1.PersistentVolume{}, func(t *testing.T, a, b runtime.Object) {
		Compare_PersistentVolume(t, a.(*v1.PersistentVolume), b.(*v1.PersistentVolume))
	})

	comparator.RegisterForType(&v1.PodStatusResult{}, func(t *testing.T, a, b runtime.Object) {
		Compare_PodStatusResult(t, a.(*v1.PodStatusResult), b.(*v1.PodStatusResult))
	})

	comparator.RegisterForType(&v1.Pod{}, func(t *testing.T, a, b runtime.Object) {
		Compare_Pod(t, a.(*v1.Pod), b.(*v1.Pod))
	})

	comparator.RegisterForType(&v1.LimitRange{}, func(t *testing.T, a, b runtime.Object) {
		Compare_LimitRange(t, a.(*v1.LimitRange), b.(*v1.LimitRange))
	})

	comparator.RegisterForType(&v1.PersistentVolumeClaim{}, func(t *testing.T, a, b runtime.Object) {
		Compare_PersistentVolumeClaim(t, a.(*v1.PersistentVolumeClaim), b.(*v1.PersistentVolumeClaim))
	})

	comparator.RegisterForType(&v1.Endpoints{}, func(t *testing.T, a, b runtime.Object) {
		Compare_Endpoints(t, a.(*v1.Endpoints), b.(*v1.Endpoints))
	})

	comparator.RegisterForType(&v1.ConfigMap{}, func(t *testing.T, a, b runtime.Object) {
		Compare_ConfigMap(t, a.(*v1.ConfigMap), b.(*v1.ConfigMap))
	})

	comparator.RegisterForType(&v1.Service{}, func(t *testing.T, a, b runtime.Object) {
		Compare_Service(t, a.(*v1.Service), b.(*v1.Service))
	})

	comparator.RegisterForType(&v1.Binding{}, func(t *testing.T, a, b runtime.Object) {
		Compare_Binding(t, a.(*v1.Binding), b.(*v1.Binding))
	})

	comparator.RegisterForType(&v1.RangeAllocation{}, func(t *testing.T, a, b runtime.Object) {
		Compare_RangeAllocation(t, a.(*v1.RangeAllocation), b.(*v1.RangeAllocation))
	})

	comparator.RegisterForType(&v1.Namespace{}, func(t *testing.T, a, b runtime.Object) {
		Compare_Namespace(t, a.(*v1.Namespace), b.(*v1.Namespace))
	})

	comparator.RegisterForType(&v1.Node{}, func(t *testing.T, a, b runtime.Object) {
		Compare_Node(t, a.(*v1.Node), b.(*v1.Node))
	})

	comparator.RegisterForType(&v1.EphemeralContainers{}, func(t *testing.T, a, b runtime.Object) {
		Compare_EphemeralContainers(t, a.(*v1.EphemeralContainers), b.(*v1.EphemeralContainers))
	})

	comparator.RegisterForType(&v1.Event{}, func(t *testing.T, a, b runtime.Object) {
		Compare_Event(t, a.(*v1.Event), b.(*v1.Event))
	})

	comparator.RegisterForType(&v1.ServiceAccount{}, func(t *testing.T, a, b runtime.Object) {
		Compare_ServiceAccount(t, a.(*v1.ServiceAccount), b.(*v1.ServiceAccount))
	})

}

func Compare_ResourceQuota(t *testing.T, expected, actual *v1.ResourceQuota) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.ResourceQuota)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_Secret(t *testing.T, expected, actual *v1.Secret) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.Secret)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_ReplicationController(t *testing.T, expected, actual *v1.ReplicationController) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.ReplicationController)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_ComponentStatus(t *testing.T, expected, actual *v1.ComponentStatus) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.ComponentStatus)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_PodTemplate(t *testing.T, expected, actual *v1.PodTemplate) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.PodTemplate)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_PersistentVolume(t *testing.T, expected, actual *v1.PersistentVolume) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.PersistentVolume)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_PodStatusResult(t *testing.T, expected, actual *v1.PodStatusResult) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.PodStatusResult)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_Pod(t *testing.T, expected, actual *v1.Pod) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.Pod)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_LimitRange(t *testing.T, expected, actual *v1.LimitRange) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.LimitRange)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_PersistentVolumeClaim(t *testing.T, expected, actual *v1.PersistentVolumeClaim) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.PersistentVolumeClaim)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_Endpoints(t *testing.T, expected, actual *v1.Endpoints) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.Endpoints)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_ConfigMap(t *testing.T, expected, actual *v1.ConfigMap) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.ConfigMap)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_Service(t *testing.T, expected, actual *v1.Service) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.Service)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_Binding(t *testing.T, expected, actual *v1.Binding) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.Binding)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_RangeAllocation(t *testing.T, expected, actual *v1.RangeAllocation) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.RangeAllocation)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_Namespace(t *testing.T, expected, actual *v1.Namespace) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.Namespace)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_Node(t *testing.T, expected, actual *v1.Node) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.Node)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_EphemeralContainers(t *testing.T, expected, actual *v1.EphemeralContainers) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.EphemeralContainers)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_Event(t *testing.T, expected, actual *v1.Event) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.Event)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}

func Compare_ServiceAccount(t *testing.T, expected, actual *v1.ServiceAccount) {
	assert := assert.New(t)
	e := expected.DeepCopyObject().(*v1.ServiceAccount)
	e.ObjectMeta = test.Sanitize_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	assert.EqualValues(expected, actual)
}
