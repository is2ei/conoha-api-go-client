package conoha

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_IsoImages(t *testing.T) {
	ts := createMockServer(t, []string{
		"compute/iso_images",
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

	isoImages, meta, err := conoha.IsoImages(context.Background())

	assert.NoError(t, err)

	assert.IsType(t, new([]*IsoImage), &isoImages)

	assert.IsType(t, new(ResponseMeta), meta)
}
