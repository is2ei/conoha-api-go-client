package conoha

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_ComputeFlavorsDetail(t *testing.T) {
	ts := createMockServer(t, []string{
		"/compute/flavors_detail",
	})
	defer ts.Close()

	conoha := NewConoha(
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		"username",
		"password",
		"tenant_id",
		"token",
	)

	flavors, meta, err := conoha.ComputeFlavorsDetail(context.Background())

	assert.NoError(t, err)

	assert.IsType(t, new([]*ComputeFlavor), &flavors)

	assert.IsType(t, new(ResponseMeta), meta)
}
