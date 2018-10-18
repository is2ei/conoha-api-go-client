package conoha

import (
	"testing"

	"github.com/is2ei/conoha-api-go-client/test"
	"github.com/stretchr/testify/assert"
)

func TestConoha_DatabaseVersions(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"/database/versions",
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

	versions, meta, err := conoha.DatabaseVersions()

	assert.NoError(t, err)

	assert.IsType(t, new([]Version), &versions)

	assert.IsType(t, new(ResponseMeta), meta)
}
