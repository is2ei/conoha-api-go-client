package conoha

import (
	"testing"

	"github.com/is2ei/conoha-api-go-client/test"
	"github.com/stretchr/testify/assert"
)

func TestConoha_ComputeKeypairs(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"compute/keypairs",
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

	keypairs, meta, err := conoha.ComputeKeypairs()

	assert.NoError(t, err)

	assert.Equal(t, 2, len(keypairs))

	assert.IsType(t, new(responseMeta), meta)
}
