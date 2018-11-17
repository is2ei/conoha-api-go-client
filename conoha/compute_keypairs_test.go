package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_ComputeKeypairs(t *testing.T) {
	ts := createMockServer(t, []string{
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
