package conoha

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_AccountOrderItem(t *testing.T) {
	ts := createMockServer(t, []string{
		"/account/order_item",
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

	item, meta, err := conoha.OrderItem(context.Background(), "item_id")

	assert.NoError(t, err)

	assert.IsType(t, new(OrderItem), item)

	assert.IsType(t, new(ResponseMeta), meta)
}
