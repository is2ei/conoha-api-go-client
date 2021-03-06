package conoha

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_MailVersions(t *testing.T) {
	ts := createMockServer(t, []string{
		"/mail/versions",
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

	versions, meta, err := conoha.MailVersions(context.Background())

	assert.NoError(t, err)

	assert.IsType(t, new([]*Version), &versions)

	assert.IsType(t, new(ResponseMeta), meta)
}
