package v1

import (
	testing "testing"

	v1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1"
	assert "github.com/stretchr/testify/assert"
	test "go.medium.engineering/kubernetes/pkg/test"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	if err := v1.AddToScheme(test.DefaultScheme); err != nil {
		panic(err)
	}
	RegisterAsserts(test.DefaultComparator)
}

func RegisterAsserts(comparator *test.Comparator) {
	comparator.RegisterForType(&v1.ExternalSecret{}, test.TypedAsserts{
		Match: func(t *testing.T, a, b runtime.Object) {
			Match_ExternalSecret(t, a.(*v1.ExternalSecret), b.(*v1.ExternalSecret))
		},
		NoMatch: func(t *testing.T, a, b runtime.Object) {
			NoMatch_ExternalSecret(t, a.(*v1.ExternalSecret), b.(*v1.ExternalSecret))
		},
	})
}

func Assimilate_ExternalSecret(expected, actual *v1.ExternalSecret) *v1.ExternalSecret {
	e := expected.DeepCopyObject().(*v1.ExternalSecret)
	e.ObjectMeta = test.Assimilate_ObjectMeta(e.ObjectMeta, actual.ObjectMeta)
	e.TypeMeta = test.Assimilate_TypeMeta(e.TypeMeta, actual.TypeMeta)
	return e
}

func Match_ExternalSecret(t *testing.T, expected, actual *v1.ExternalSecret) {
	assert := assert.New(t)
	e := Assimilate_ExternalSecret(expected, actual)
	assert.EqualValues(e, actual)
}

func NoMatch_ExternalSecret(t *testing.T, expected, actual *v1.ExternalSecret) {
	assert := assert.New(t)
	e := Assimilate_ExternalSecret(expected, actual)
	assert.NotEqualValues(e, actual)
}
