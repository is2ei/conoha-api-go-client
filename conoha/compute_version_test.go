package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_ComputeVersion(t *testing.T) {
	ts := createMockServer(t, []string{
		"compute/version",
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

	version, meta, err := conoha.ComputeVersion()

	assert.NoError(t, err)

	assert.IsType(t, new(Version), version)

	assert.IsType(t, new(ResponseMeta), meta)
}
