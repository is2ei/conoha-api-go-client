package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_DatabaseVersion(t *testing.T) {
	ts := createMockServer(t, []string{
		"/database/version",
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

	version, meta, err := conoha.DatabaseVersion()

	assert.NoError(t, err)

	assert.IsType(t, new(Version), &version)

	assert.IsType(t, new(ResponseMeta), meta)
}
