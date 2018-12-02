package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_AccountPaymentSummary(t *testing.T) {
	ts := createMockServer(t, []string{
		"/account/payment_summary",
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

	paymentSummary, meta, err := conoha.PaymentSummary()

	assert.NoError(t, err)

	assert.IsType(t, new(PaymentSummary), paymentSummary)

	assert.IsType(t, new(ResponseMeta), meta)
}
