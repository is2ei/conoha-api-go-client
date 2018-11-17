package conoha

import (
	"testing"

	"github.com/is2ei/conoha-api-go-client/test"
	"github.com/stretchr/testify/assert"
)

func TestConoha_ComputeServers(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
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

	assert.IsType(t, new(responseMeta), meta)
}
