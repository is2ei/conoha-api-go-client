package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_ComputeServersDetail(t *testing.T) {
	ts := createMockServer(t, []string{
		"/compute/servers_detail",
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

	servers, meta, err := conoha.ComputeServersDetail()

	assert.NoError(t, err)

	assert.IsType(t, new([]*ComputeServer), &servers)
	assert.Equal(t, "133.130.49.xxx", servers[0].AccessIPv4)
	assert.Equal(t, "133.130.49.xxx", servers[0].Addresses["ext-133-130-48-xx-xx"][0].Addr)

	assert.IsType(t, new(ResponseMeta), meta)
}
