package conoha

import "encoding/json"

type getAccountOrderItemsResponseParam struct {
	OrderItems []OrderItem `json:"order_items"`
}

// OrderItems fetches the orders list.
//
// ConoHa API docs: https://www.conoha.jp/docs/account-order-item-list.html
func (c *Conoha) OrderItems() ([]OrderItem, *responseMeta, error) {
	p := getAccountOrderItemsResponseParam{}

	u := c.AccountServiceURL + "/v1/" + c.TenantID + "/order-items"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.OrderItems, meta, err
}
