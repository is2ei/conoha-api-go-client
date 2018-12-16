package conoha

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_IdentityApiVersions(t *testing.T) {
	ts := createMockServer(t, []string{
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

	versions, meta, err := conoha.IdentityAPIVersions(context.Background())

	assert.NoError(t, err)

	assert.IsType(t, new([]*Version), &versions)
	assert.Equal(t, 1, len(versions))
	assert.Equal(t, "v2.0", versions[0].ID)

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, "GET", meta.Method)
	assert.Equal(t, 300, meta.StatusCode)

	assert.IsType(t, new(Conoha), conoha)
}
