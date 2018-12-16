package conoha

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_MailVersion(t *testing.T) {
	ts := createMockServer(t, []string{
		"mail/version",
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

	version, meta, err := conoha.MailVersion(context.Background())

	assert.NoError(t, err)

	assert.IsType(t, new(Version), version)

	assert.IsType(t, new(ResponseMeta), meta)
}
