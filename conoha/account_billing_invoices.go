package conoha

import "encoding/json"

type getAccountBillingInvoicesResponseParam struct {
	BillingInvoices []BillingInvoice `json:"billing_invoices"`
}

// BillingInvoices fetches invoices list.
func (c *Conoha) BillingInvoices() ([]BillingInvoice, *responseMeta, error) {
	p := getAccountBillingInvoicesResponseParam{}

	u := c.AccountServiceUrl + "/v1/" + c.TenantId + "/billing-invoices"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.BillingInvoices, meta, err
}
