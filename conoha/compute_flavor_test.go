package conoha

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_ComputeFlavor(t *testing.T) {
	ts := createMockServer(t, []string{
		"/compute/flavor",
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

	flavor, meta, err := conoha.ComputeFlavor(context.Background(), "flavor_id")

	assert.NoError(t, err)

	assert.IsType(t, new(ComputeFlavor), flavor)
	assert.Equal(t, false, flavor.OsFlvDisabled)
	assert.Equal(t, 0, flavor.OsFlvExtData)
	assert.Equal(t, 50, flavor.Disk)
	assert.Equal(t, "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX", flavor.ID)

	assert.IsType(t, new([]*Link), &flavor.Links)
	assert.Equal(t, "https://compute.tyo1.conoha.io/v2/1864e71d2deb46f6b47526b69c65a45d/flavors/XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX", flavor.Links[0].Href)
	assert.Equal(t, "self", flavor.Links[0].Rel)
	assert.Equal(t, "https://compute.tyo1.conoha.io/1864e71d2deb46f6b47526b69c65a45d/flavors/XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX", flavor.Links[1].Href)
	assert.Equal(t, "bookmark", flavor.Links[1].Rel)

	assert.Equal(t, "2gb-flavor", flavor.Name)
	assert.Equal(t, true, flavor.OsFlavorAccessIsPublic)
	assert.Equal(t, 2048, flavor.RAM)
	assert.Equal(t, 1, flavor.RxtxFactor)
	assert.Equal(t, "", flavor.Swap)
	assert.Equal(t, 3, flavor.Vcpus)

	assert.IsType(t, new(ResponseMeta), meta)
}
