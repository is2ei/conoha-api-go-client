package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_AccountApiVersions(t *testing.T) {
	ts := createMockServer(t, []string{
		"/account/versions",
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

	versions, meta, err := conoha.AccountApiVersions()

	assert.NoError(t, err)

	assert.IsType(t, new([]*Version), &versions)
	assert.Equal(t, 1, len(versions))
	assert.Equal(t, "CURRENT", versions[0].Status)
	assert.Equal(t, "2015-05-12T09:00:00Z", versions[0].Updated)

	assert.IsType(t, new(responseMeta), meta)
	assert.Equal(t, "GET", meta.Method)
	assert.Equal(t, 300, meta.StatusCode)

	assert.IsType(t, new(Conoha), conoha)
}