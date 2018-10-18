package conoha

import (
	"testing"

	"github.com/is2ei/conoha-api-go-client/test"
	"github.com/stretchr/testify/assert"
)

func TestConoha_AccountOrderItem(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
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

	item, meta, err := conoha.OrderItem("item_id")

	assert.NoError(t, err)

	assert.IsType(t, new(OrderItem), &item)

	assert.IsType(t, new(ResponseMeta), meta)
}
