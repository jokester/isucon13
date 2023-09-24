package scenario

import (
	"context"
	"testing"

	"github.com/isucon/isucandar/agent"
	"github.com/isucon/isucon13/bench/internal/benchtest"
	"github.com/isucon/isucon13/bench/isupipe"
	"github.com/stretchr/testify/assert"
)

func TestPretest(t *testing.T) {
	ctx := context.Background()

	testResource, err := benchtest.Setup()
	if err != nil {
		t.Fatal(err)
	}
	defer benchtest.Teardown(testResource)

	client, err := isupipe.NewClient(
		agent.WithBaseURL(testResource.WebappIPAddress()),
	)
	if err != nil {
		t.Fatal(err)
	}

	err = Pretest(ctx, client)
	assert.NoError(t, err)
}
