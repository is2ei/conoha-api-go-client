package conoha

import (
	"context"
	"encoding/json"
	"fmt"
)

type getAccountPaymentSummaryResponseParam struct {
	PaymentSummary *PaymentSummary `json:"payment_summary"`
}

// PaymentSummary represencts summarized information of the payment.
type PaymentSummary struct {
	TotalDepositAmount int `json:"total_deposit_amount"`
}

// PaymentSummary fetches summarized information of the payment.
//
// ConoHa API docs: https://www.conoha.jp/docs/account-payment-summary.html
func (c *Conoha) PaymentSummary(ctx context.Context) (*PaymentSummary, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v1/%s/payment-summary", c.AccountServiceURL, c.TenantID)

	p := getAccountPaymentSummaryResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, "GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.PaymentSummary, meta, err
}
