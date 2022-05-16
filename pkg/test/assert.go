package test

import (
	"context"
	"testing"

	testify "github.com/stretchr/testify/assert"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func AssertMatch(
	ctx context.Context,
	t *testing.T,
	cli client.Client,
	expected client.Object,
	msgAndArgs ...interface{},
) {
	DefaultComparator.AssertMatch(ctx, t, cli, expected, msgAndArgs...)
}

func AssertNotFound(
	ctx context.Context,
	t *testing.T,
	cli client.Client,
	unexpected client.Object,
	msgAndArgs ...interface{},
) {
	assert := testify.New(t)
	actual := unexpected.DeepCopyObject().(client.Object)
	key := client.ObjectKeyFromObject(unexpected)
	err := cli.Get(ctx, key, actual)
	assert.Error(err, "Unexpected object found")
}

func AssertAllMatch(
	ctx context.Context,
	t *testing.T,
	cli client.Client,
	expected []client.Object,
	msgAndArgs ...interface{},
) {
	for _, e := range expected {
		AssertMatch(ctx, t, cli, e, msgAndArgs...)
	}
}
