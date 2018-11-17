package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_ComputeImage(t *testing.T) {
	ts := createMockServer(t, []string{
		"compute/image",
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

	image, meta, err := conoha.ComputeImage("image_id")

	assert.NoError(t, err)

	assert.IsType(t, new(ComputeImage), image)
	assert.Equal(t, 2733375488, image.OsExtImgSize)
	assert.Equal(t, "2015-05-05T11:25:19Z", image.Created)

	assert.IsType(t, new([]*Link), &image.Links)
	assert.Equal(t, 3, len(image.Links))
	assert.Equal(t, "https://compute.tyo1.conoha.io/v2/1864e71d2deb46f6b47526b69c65a45d/images/0ccce41e-7049-49e9-89b0-2fafad2c02a7", image.Links[0].Href)
	assert.Equal(t, "self", image.Links[0].Rel)
	assert.Equal(t, "https://compute.tyo1.conoha.io/1864e71d2deb46f6b47526b69c65a45d/images/0ccce41e-7049-49e9-89b0-2fafad2c02a7", image.Links[1].Href)
	assert.Equal(t, "bookmark", image.Links[1].Rel)
	assert.Equal(t, "https://image-service.tyo1.conoha.io/1864e71d2deb46f6b47526b69c65a45d/images/0ccce41e-7049-49e9-89b0-2fafad2c02a7", image.Links[2].Href)
	assert.Equal(t, "alternate", image.Links[2].Rel)
	assert.Equal(t, "application/vnd.openstack.image", image.Links[2].Type)

	assert.IsType(t, new(ComputeImageMetadata), image.Metadata)
	assert.Equal(t, "210", image.Metadata.DisplayOrder)
	assert.Equal(t, "Ruby on Rails-4.2-64bit", image.Metadata.GncApp)
	assert.Equal(t, "yes", image.Metadata.HwQemuGuestAgent)
	assert.Equal(t, "lin", image.Metadata.OsType)

	assert.Equal(t, 0, image.MinDisk)
	assert.Equal(t, 0, image.MinRam)
	assert.Equal(t, "vmi-ror-4.2-centos-6.6", image.Name)
	assert.Equal(t, 100, image.Progress)
	assert.Equal(t, "ACTIVE", image.Status)
	assert.Equal(t, "2015-05-07T08:14:47Z", image.Updated)

	assert.IsType(t, new(responseMeta), meta)
}
