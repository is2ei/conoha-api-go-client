package conoha

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_AccountBillingInvoice(t *testing.T) {
	ts := createMockServer(t, []string{
		"/account/billing_invoice",
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

	billingInvoice, meta, err := conoha.BillingInvoice(context.Background(), "invoice_id")

	assert.NoError(t, err)

	assert.IsType(t, new(BillingInvoice), &billingInvoice)

	assert.IsType(t, new(ResponseMeta), meta)
}
