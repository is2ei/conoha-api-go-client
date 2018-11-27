package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_ComputeFlavors(t *testing.T) {
	ts := createMockServer(t, []string{
		"/compute/flavors",
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

	flavors, meta, err := conoha.ComputeFlavors()

	assert.NoError(t, err)

	assert.IsType(t, new([]*ComputeFlavor), &flavors)
	assert.Equal(t, 1, len(flavors))
	assert.Equal(t, "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX", flavors[0].ID)
	assert.Equal(t, "2gb-flavor", flavors[0].Name)

	assert.IsType(t, new(ResponseMeta), meta)
	assert.IsType(t, 200, meta.StatusCode)
}
