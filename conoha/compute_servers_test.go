package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_ComputeServers(t *testing.T) {
	ts := createMockServer(t, []string{
		"/compute/servers",
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

	servers, meta, err := conoha.ComputeServers()

	assert.NoError(t, err)

	assert.IsType(t, new([]*ComputeServer), &servers)

	assert.IsType(t, new(ResponseMeta), meta)
}
