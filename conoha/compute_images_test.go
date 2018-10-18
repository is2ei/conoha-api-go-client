package conoha

import (
	"testing"

	"github.com/is2ei/conoha-api-go-client/test"
	"github.com/stretchr/testify/assert"
)

func TestConoha_ComputeImages(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"/compute/images",
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

	images, meta, err := conoha.ComputeImages()

	assert.NoError(t, err)

	assert.IsType(t, new([]ComputeImage), &images)
	assert.Equal(t, 1, len(images))
	assert.Equal(t, "545e54a7-77ba-4d49-8c0e-af7291bb7b63", images[0].Id)
	assert.Equal(t, "500G-0507", images[0].Name)

	assert.IsType(t, new(ResponseMeta), meta)
	assert.IsType(t, 200, meta.StatusCode)
}
