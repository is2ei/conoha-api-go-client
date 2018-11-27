package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_ComputeKeypair(t *testing.T) {
	ts := createMockServer(t, []string{
		"compute/keypair",
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

	keypair, meta, err := conoha.ComputeKeypair("keypair_name")

	assert.NoError(t, err)

	assert.IsType(t, new(Keypair), keypair)
	assert.Equal(t, "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDx8nkQv/zgGgB4rMYmIf+6A4l6Rr+o/6lHBQdW5aYd44bd8JttDCE/F/pNRr0lRE+PiqSPO8nDPHw0010JeMH9gYgnnFlyY3/OcJ02RhIPyyxYpv9FhY+2YiUkpwFOcLImyrxEsYXpD/0d3ac30bNH6Sw9JD9UZHYcpSxsIbECHw== Generated by Nova", keypair.PublicKey)
	assert.Equal(t, "d24345400a044b73ba99c3d18ef31a4e", keypair.UserID)
	assert.Equal(t, "keypair-original-01", keypair.Name)
	assert.Equal(t, false, keypair.Deleted)
	assert.Equal(t, "2014-12-11T05:45:59.000000", keypair.CreatedAt)
	assert.Equal(t, "", keypair.UpdatedAt)
	assert.Equal(t, "1e:2c:9b:56:79:4b:45:77:f9:ca:7a:98:2c:b0:d5:3c", keypair.FingerPrint)
	assert.Equal(t, "", keypair.DeletedAt)
	assert.Equal(t, 9, keypair.ID)

	assert.IsType(t, new(ResponseMeta), meta)
}
