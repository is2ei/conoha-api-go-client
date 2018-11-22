package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_IdentityApiVersion(t *testing.T) {
	ts := createMockServer(t, []string{
		"identity/version",
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

	version, meta, err := conoha.IdentityApiVersion()

	assert.NoError(t, err)

	assert.IsType(t, new(Version), version)
	assert.Equal(t, "stable", version.Status)
	assert.Equal(t, "2015-05-12T09:00:00Z", version.Updated)
	assert.Equal(t, 2, len(version.MediaTypes))
	assert.Equal(t, "v2.0", version.Id)
	assert.Equal(t, 2, len(version.Links))

	assert.IsType(t, new(responseMeta), meta)
	assert.Equal(t, "GET", meta.Method)
	assert.Equal(t, 200, meta.StatusCode)

	assert.IsType(t, new(Conoha), conoha)
}
