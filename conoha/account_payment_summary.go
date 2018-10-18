package conoha

import "encoding/json"

type GetAccountPaymentSummaryResponseParam struct {
	PaymentSummary PaymentSummary `json:"payment_summary"`
}

type PaymentSummary struct {
	TotalDepositAmount int `json:"total_deposit_amount"`
}

func (c *Conoha) PaymentSummary() (PaymentSummary, *ResponseMeta, error) {
	p := GetAccountPaymentSummaryResponseParam{}

	u := c.AccountServiceUrl + "/v1/" + c.TenantId + "/payment-summary"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.PaymentSummary, meta, err
}
