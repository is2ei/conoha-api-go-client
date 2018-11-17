package conoha

import "encoding/json"

type getAccountBillingInvoiceResponseParam struct {
	BillingInvoice BillingInvoice `json:"billing_invoice"`
}

type BillingInvoice struct {
	Items             []Item `json:"items"`
	InvoiceId         int    `json:"invoice_id"`
	PaymentMethodType string `json:"payment_method_type"`
	InvoiceDate       string `json:"invoice_date"`
	BillPlasTax       int    `json:"bill_plas_tax"`
	DueDate           string `json:"due_date"`
}

type Item struct {
	InvoiceDetailId int     `json:"invoice_detail_id"`
	ProductName     string  `json:"product_name"`
	UnitPrice       float32 `json:"unit_price"`
	Quantity        int     `json:"quantity"`
	StartDate       string  `json:"start_date"`
	EndDate         string  `json:"EndDate"`
}

func (c *Conoha) BillingInvoice(invoiceId string) (BillingInvoice, *responseMeta, error) {
	p := getAccountBillingInvoiceResponseParam{}

	u := c.AccountServiceUrl + "/v1/" + c.TenantId + "/billing-invoices/" + invoiceId

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.BillingInvoice, meta, err
}
