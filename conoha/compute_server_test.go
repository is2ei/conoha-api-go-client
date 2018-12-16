package conoha

import (
	"context"
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

	server, meta, err := conoha.ComputeServer(context.Background(), "server_id")

	assert.NoError(t, err)

	assert.IsType(t, new(ComputeServer), server)

	assert.IsType(t, new(ResponseMeta), meta)
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

	server, meta, err := conoha.AddComputeServer(context.Background(), "", "", "", "", "")

	assert.NoError(t, err)

	assert.IsType(t, new(ComputeServer), server)

	assert.IsType(t, new(ResponseMeta), meta)
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

	meta, err := conoha.DeleteComputeServer(context.Background(), "server_id")

	assert.NoError(t, err)

	assert.IsType(t, new(ResponseMeta), meta)
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

	meta, err := conoha.StartComputeServer(context.Background(), "server_id")

	assert.NoError(t, err)

	assert.IsType(t, new(ResponseMeta), meta)
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

	meta, err := conoha.StopComputeServer(context.Background(), "server_id")

	assert.NoError(t, err)

	assert.IsType(t, new(ResponseMeta), meta)

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

	meta, err := conoha.RebootComputeServer(context.Background(), "server_id", true)

	assert.NoError(t, err)

	assert.IsType(t, new(ResponseMeta), meta)
}
