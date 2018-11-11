package conoha

import "encoding/json"

type getAccountOrderItemsResponseParam struct {
	OrderItems []OrderItem `json:"order_items"`
}

func (c *Conoha) OrderItems() ([]OrderItem, *ResponseMeta, error) {
	p := getAccountOrderItemsResponseParam{}

	u := c.AccountServiceUrl + "/v1/" + c.TenantId + "/order-items"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.OrderItems, meta, err
}
