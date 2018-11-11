package conoha

import "encoding/json"

type getAccountPaymentHistoryResponseParam struct {
	PaymentHistory []PaymentHistory `json:"payment_history"`
}

type PaymentHistory struct {
	MoneyType     string `json:"money_type"`
	DepositAmount int    `json:"deposit_amount"`
	ReceivedDate  string `json:"received_date"`
}

func (c *Conoha) PaymentHistory() ([]PaymentHistory, *ResponseMeta, error) {
	p := getAccountPaymentHistoryResponseParam{}

	u := c.AccountServiceUrl + "/v1/" + c.TenantId + "/payment-history"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.PaymentHistory, meta, err
}
