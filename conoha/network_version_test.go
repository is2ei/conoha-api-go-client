package conoha

import (
	"testing"

	"github.com/is2ei/conoha-api-go-client/test"
	"github.com/stretchr/testify/assert"
)

func TestConoha_NetworkVersion(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"network/version",
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

	version, meta, err := conoha.NetworkVersion()

	assert.NoError(t, err)

	assert.IsType(t, new(Version), &version)

	assert.IsType(t, new(ResponseMeta), meta)
}
