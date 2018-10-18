package conoha

import (
	"testing"

	"github.com/is2ei/conoha-api-go-client/test"
	"github.com/stretchr/testify/assert"
)

func TestConoha_IdentityVersions(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"/identity/versions",
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

	versions, meta, err := conoha.IdentityVersions()

	assert.NoError(t, err)

	assert.IsType(t, new([]Version), &versions)
	assert.Equal(t, 1, len(versions))
	assert.Equal(t, "v2.0", versions[0].Id)

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, "GET", meta.Method)
	assert.Equal(t, 300, meta.StatusCode)

	assert.IsType(t, new(Conoha), conoha)
}
