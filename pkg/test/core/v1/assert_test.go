package v1

import (
	"context"
	"testing"
	"time"

	"go.medium.engineering/kubernetes/pkg/test"
	core "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var fixtures = []client.Object{
	&core.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name: "fixture",
		},
	},
}

func TestMatch(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1)*time.Second)
	defer cancel()
	cli := fake.NewClientBuilder().WithScheme(test.DefaultScheme).WithObjects(fixtures...).Build()
	//cli := fake.NewFakeClientWithScheme(test.DefaultScheme, fixtures...)
	test.AssertMatch(ctx, t, cli, fixtures[0])
	s := *fixtures[0].(*core.Secret)
	s.Name = "baba"
	test.AssertNotFound(ctx, t, cli, &s)
}
