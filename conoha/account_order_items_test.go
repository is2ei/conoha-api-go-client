package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_OrderItems(t *testing.T) {
	ts := createMockServer(t, []string{
		"/account/order_items",
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

	items, meta, err := conoha.OrderItems()

	assert.NoError(t, err)

	assert.IsType(t, new([]*OrderItem), &items)

	assert.IsType(t, new(ResponseMeta), meta)
}
