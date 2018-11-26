package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_AccountBillingInvoices(t *testing.T) {
	ts := createMockServer(t, []string{
		"/account/billing_invoices",
	})
	defer ts.Close()

	conoha := NewConoha(
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		"username",
		"password",
		"tenant_id",
		"token",
	)

	billingInvoices, meta, err := conoha.BillingInvoices()

	assert.NoError(t, err)

	assert.IsType(t, new([]BillingInvoice), &billingInvoices)

	assert.IsType(t, new(ResponseMeta), meta)
}
