package conoha

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_AccountNotifications(t *testing.T) {
	ts := createMockServer(t, []string{
		"/account/notifications",
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

	notifications, meta, err := conoha.Notifications(context.Background())

	assert.NoError(t, err)

	assert.IsType(t, new([]*Notification), &notifications)

	assert.IsType(t, new(ResponseMeta), meta)
}
