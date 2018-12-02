package conoha

import (
	"encoding/json"
	"fmt"
)

type getAccountPaymentHistoryResponseParam struct {
	PaymentHistory []*PaymentHistory `json:"payment_history"`
}

// PaymentHistory represents payment history.
type PaymentHistory struct {
	MoneyType     string `json:"money_type"`
	DepositAmount int    `json:"deposit_amount"`
	ReceivedDate  string `json:"received_date"`
}

// PaymentHistory fetches payment history.
//
// ConoHa API docs: https://www.conoha.jp/docs/account-payment-histories.html
func (c *Conoha) PaymentHistory() ([]*PaymentHistory, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v1/%s/payment-history", c.AccountServiceURL, c.TenantID)

	p := getAccountPaymentHistoryResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.PaymentHistory, meta, err
}
