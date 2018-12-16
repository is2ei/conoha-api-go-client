package conoha

import (
	"context"
	"encoding/json"
	"fmt"
)

type getAccountBillingInvoiceResponseParam struct {
	BillingInvoice BillingInvoice `json:"billing_invoice"`
}

// BillingInvoice represents invoice information.
type BillingInvoice struct {
	Items             []*BillingInvoiceItem `json:"items"`
	InvoiceID         int                   `json:"invoice_id"`
	PaymentMethodType string                `json:"payment_method_type"`
	InvoiceDate       string                `json:"invoice_date"`
	BillPlasTax       int                   `json:"bill_plas_tax"`
	DueDate           string                `json:"due_date"`
}

// BillingInvoiceItem represents detailed information of the invoice.
type BillingInvoiceItem struct {
	InvoiceDetailID int     `json:"invoice_detail_id"`
	ProductName     string  `json:"product_name"`
	UnitPrice       float32 `json:"unit_price"`
	Quantity        int     `json:"quantity"`
	StartDate       string  `json:"start_date"`
	EndDate         string  `json:"EndDate"`
}

// BillingInvoice fetches invoice information.
//
// ConoHa API docs: https://www.conoha.jp/docs/account-billing-invoices-detail-specified.html
func (c *Conoha) BillingInvoice(ctx context.Context, invoiceID string) (BillingInvoice, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v1/%s/billing-invoices/%s", c.AccountServiceURL, c.TenantID, invoiceID)

	p := getAccountBillingInvoiceResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, "GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.BillingInvoice, meta, err
}
