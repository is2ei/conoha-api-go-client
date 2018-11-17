package conoha

import "encoding/json"

type getAccountPaymentSummaryResponseParam struct {
	PaymentSummary PaymentSummary `json:"payment_summary"`
}

// PaymentSummary represencts summarized information of the payment.
type PaymentSummary struct {
	TotalDepositAmount int `json:"total_deposit_amount"`
}

// PaymentSummary fetches summarized information of the payment.
//
// ConoHa API docs: https://www.conoha.jp/docs/account-payment-summary.html
func (c *Conoha) PaymentSummary() (PaymentSummary, *responseMeta, error) {
	p := getAccountPaymentSummaryResponseParam{}

	u := c.AccountServiceUrl + "/v1/" + c.TenantId + "/payment-summary"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.PaymentSummary, meta, err
}
