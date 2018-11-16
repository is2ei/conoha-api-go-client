package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_NewErrorUnauthorized(t *testing.T) {
	err := NewErrorUnauthorized("error")

	assert.EqualError(t, err, "error")

	msg := err.Error()

	assert.Equal(t, "error", msg)
}
