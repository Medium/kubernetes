package test

import (
	"context"
	testify "github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"testing"
)

func AssertMatch(
	ctx context.Context,
	t *testing.T,
	cli client.Client,
	expected runtime.Object,
	msgAndArgs ...interface{},
) {
	DefaultComparator.AssertMatch(ctx, t, cli, expected, msgAndArgs...)
}

func AssertNotFound(
	ctx context.Context,
	t *testing.T,
	cli client.Client,
	unexpected runtime.Object,
	msgAndArgs ...interface{},
) {
	assert := testify.New(t)
	actual := unexpected.DeepCopyObject()
	key, err := client.ObjectKeyFromObject(unexpected)
	assert.NoError(err, msgAndArgs...)
	err = cli.Get(ctx, key, actual)
	assert.Error(err, "Unexpected object found")
}

func AssertAllMatch(
	ctx context.Context,
	t *testing.T,
	cli client.Client,
	expected []runtime.Object,
	msgAndArgs ...interface{},
) {
	for _, e := range expected {
		AssertMatch(ctx, t, cli, e, msgAndArgs...)
	}
}
