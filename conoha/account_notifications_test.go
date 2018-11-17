package conoha

import (
	"testing"

	"github.com/is2ei/conoha-api-go-client/test"
	"github.com/stretchr/testify/assert"
)

func TestConoha_AccountNotifications(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
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

	notifications, meta, err := conoha.Notifications()

	assert.NoError(t, err)

	assert.IsType(t, new([]Notification), &notifications)

	assert.IsType(t, new(responseMeta), meta)
}
