package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_ForceShutdownServer(t *testing.T) {
	ts := createMockServer(t, []string{
		"/compute/server_force_shutdown",
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

	meta, err := conoha.ForceShutdownServer("server_id")

	assert.NoError(t, err)

	assert.IsType(t, new(ResponseMeta), meta)
}
