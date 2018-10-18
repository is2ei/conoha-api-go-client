package conoha

import (
	"testing"

	"github.com/is2ei/conoha-api-go-client/test"
	"github.com/stretchr/testify/assert"
)

func TestConoha_AddComputeServer(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
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

	server, meta, err := conoha.AddComputeServer("", "", "", "", "")

	assert.NoError(t, err)

	assert.IsType(t, new(ComputeServer), server)

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 202, meta.StatusCode)
}
