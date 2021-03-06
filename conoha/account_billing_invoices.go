package conoha

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type getAccountBillingInvoicesResponseParam struct {
	BillingInvoices []*BillingInvoice `json:"billing_invoices"`
}

// BillingInvoices fetches invoices list.
//
// ConoHa API docs: https://www.conoha.jp/docs/account-billing-invoices-list.html
func (c *Conoha) BillingInvoices(ctx context.Context) ([]*BillingInvoice, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v1/%s/billing-invoices", c.AccountServiceURL, c.TenantID)

	p := getAccountBillingInvoicesResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.BillingInvoices, meta, err
}
