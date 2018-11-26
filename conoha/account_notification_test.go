package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_AccountNotification(t *testing.T) {
	ts := createMockServer(t, []string{
		"/account/notification",
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

	notification, meta, err := conoha.Notification("notification_code")

	assert.NoError(t, err)

	assert.IsType(t, new(Notification), &notification)

	assert.IsType(t, new(ResponseMeta), meta)
}
