package conoha

import (
	"testing"

	"github.com/is2ei/conoha-api-go-client/test"
	"github.com/stretchr/testify/assert"
)

func TestConoha_AccountPaymentHistory(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
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

	paymentHistory, meta, err := conoha.PaymentHistory()

	assert.NoError(t, err)

	assert.IsType(t, new([]PaymentHistory), &paymentHistory)

	assert.IsType(t, new(ResponseMeta), meta)
}
