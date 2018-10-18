package conoha

import "encoding/json"

type GetAccountOrderItemResponseParam struct {
	OrderItem OrderItem `json:"order_item"`
}

type OrderItem struct {
	UUID             string  `json:"uu_id"`
	ServiceName      string  `json:"service_name"`
	ProductName      string  `json:"product_name"`
	Status           string  `json:"status"`
	UnitPrice        float32 `json:"unit_price"`
	ServiceStartDate string  `json:"service_start_date"`
	BillStartDate    string  `json:"bill_start_date"`
	CancelDate       string  `json:"cancel_date"`
}

func (c *Conoha) OrderItem(itemId string) (OrderItem, *ResponseMeta, error) {
	p := GetAccountOrderItemResponseParam{}

	u := c.AccountServiceUrl + "/v1/" + c.TenantId + "/order-items/" + itemId

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.OrderItem, meta, err
}
