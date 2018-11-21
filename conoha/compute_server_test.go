package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_ComputeServer(t *testing.T) {
	ts := createMockServer(t, []string{
		"/compute/server",
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

	server, meta, err := conoha.ComputeServer("server_id")

	assert.NoError(t, err)

	assert.IsType(t, new(ComputeServer), server)

	assert.IsType(t, new(responseMeta), meta)
}

func TestConoha_AddComputeServer(t *testing.T) {
	ts := createMockServer(t, []string{
		"/compute/server_add",
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

	server, meta, err := conoha.AddComputeServer("", "", "", "", "")

	assert.NoError(t, err)

	assert.IsType(t, new(ComputeServer), server)

	assert.IsType(t, new(responseMeta), meta)
	assert.Equal(t, 202, meta.StatusCode)
}

func TestConoha_DeleteComputeServer(t *testing.T) {
	ts := createMockServer(t, []string{
		"/compute/server_delete",
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

	meta, err := conoha.DeleteComputeServer("server_id")

	assert.NoError(t, err)

	assert.IsType(t, new(responseMeta), meta)
}

func TestConoha_StartComputeServer(t *testing.T) {
	ts := createMockServer(t, []string{
		"/compute/server_start",
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

	meta, err := conoha.StartComputeServer("server_id")

	assert.NoError(t, err)

	assert.IsType(t, new(responseMeta), meta)
}

func TestConoha_StopComputeServer(t *testing.T) {
	ts := createMockServer(t, []string{
		"/compute/server_stop",
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

	meta, err := conoha.StopComputeServer("server_id")

	assert.NoError(t, err)

	assert.IsType(t, new(responseMeta), meta)

}

func TestConoha_RebootComputeServer(t *testing.T) {
	ts := createMockServer(t, []string{
		"/compute/server_reboot",
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

	meta, err := conoha.RebootComputeServer("server_id", true)

	assert.NoError(t, err)

	assert.IsType(t, new(responseMeta), meta)
}
