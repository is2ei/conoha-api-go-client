package conoha

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type getAccountOrderItemsResponseParam struct {
	OrderItems []*OrderItem `json:"order_items"`
}

// OrderItems fetches the orders list.
//
// ConoHa API docs: https://www.conoha.jp/docs/account-order-item-list.html
func (c *Conoha) OrderItems(ctx context.Context) ([]*OrderItem, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v1/%s/order-items", c.AccountServiceURL, c.TenantID)

	p := getAccountOrderItemsResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.OrderItems, meta, err
}
