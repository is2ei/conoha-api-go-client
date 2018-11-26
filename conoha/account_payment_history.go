package conoha

import "encoding/json"

type getAccountPaymentHistoryResponseParam struct {
	PaymentHistory []PaymentHistory `json:"payment_history"`
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
func (c *Conoha) PaymentHistory() ([]PaymentHistory, *responseMeta, error) {
	p := getAccountPaymentHistoryResponseParam{}

	u := c.AccountServiceURL + "/v1/" + c.TenantID + "/payment-history"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.PaymentHistory, meta, err
}
