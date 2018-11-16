package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConoha(t *testing.T) {
	conoha := NewConoha(
		"http://identity_service_url",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"username",
		"password",
		"tenant_id",
		"token",
	)

	assert.IsType(t, new(Conoha), conoha)
	assert.Equal(
		t,
		conoha.IdentityServiceUrl,
		"http://identity_service_url",
	)
	assert.Equal(
		t,
		conoha.Token,
		"token",
	)
}

func TestClientDoError(t *testing.T) {
	conoha := NewConoha(
		"http://identity_service_url",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"username",
		"password",
		"tenant_id",
		"token",
	)

	_, err := conoha.execRequest("AAA", "aaa", nil)

	assert.EqualError(t, err, "Client.Do error: \"Aaa aaa: unsupported protocol scheme \\\"\\\"\"")
}
