package v1

import (
	"context"
	"go.medium.engineering/kubernetes/pkg/test"
	core "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"testing"
	"time"
)

var fixtures = []runtime.Object{
	&core.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name: "fixture",
		},
	},
}

func TestMatch(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1) * time.Second)
	defer cancel()
	cli := fake.NewFakeClientWithScheme(test.DefaultScheme, fixtures...)
	test.AssertMatch(ctx, t, cli, fixtures[0])
	s := *fixtures[0].(*core.Secret)
	s.Name = "baba"
	test.AssertNotFound(ctx, t, cli, &s)
}
