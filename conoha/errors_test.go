package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_NewErrorUnauthorized(t *testing.T) {
	err := newErrorUnauthorized("error")

	assert.EqualError(t, err, "error")

	msg := err.Error()

	assert.Equal(t, "error", msg)
}
