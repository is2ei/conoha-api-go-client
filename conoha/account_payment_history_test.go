package conoha

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_AccountPaymentHistory(t *testing.T) {
	ts := createMockServer(t, []string{
		"/account/payment_history",
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

	paymentHistory, meta, err := conoha.PaymentHistory(context.Background())

	assert.NoError(t, err)

	assert.IsType(t, new([]*PaymentHistory), &paymentHistory)

	assert.IsType(t, new(ResponseMeta), meta)
}
