package conoha

import "encoding/json"

type getAccountBillingInvoicesResponseParam struct {
	BillingInvoices []BillingInvoice `json:"billing_invoices"`
}

// BillingInvoices fetches invoices list.
//
// ConoHa API docs: https://www.conoha.jp/docs/account-billing-invoices-list.html
func (c *Conoha) BillingInvoices() ([]BillingInvoice, *ResponseMeta, error) {
	p := getAccountBillingInvoicesResponseParam{}

	u := c.AccountServiceURL + "/v1/" + c.TenantID + "/billing-invoices"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.BillingInvoices, meta, err
}
