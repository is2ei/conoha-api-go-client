package conoha

import (
	"testing"

	"github.com/is2ei/conoha-api-go-client/test"
	"github.com/stretchr/testify/assert"
)

func TestConoha_IdentityToken(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"identity/token",
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

	access, meta, err := conoha.IdentityToken()

	assert.NoError(t, err)

	assert.IsType(t, new(Access), access)

	assert.IsType(t, new(Token), access.Token)
	assert.IsType(t, "2015-05-19T07:08:21.927295", access.Token.IssuedAt)
	assert.IsType(t, "2015-05-20T07:08:21Z", access.Token.Expires)
	assert.IsType(t, "sample00d88246078f2bexample788f7", access.Token.Id)

	assert.IsType(t, new(Tenant), access.Token.Tenant)
	assert.Equal(t, "did", access.Token.Tenant.DomainId)
	assert.Equal(t, "v2", access.Token.Tenant.Description)
	assert.Equal(t, true, access.Token.Tenant.Enabled)
	assert.Equal(t, "sample00d88246078f2bexample788f7", access.Token.Tenant.Id)
	assert.Equal(t, "example00000000", access.Token.Tenant.Name)

	assert.IsType(t, new([]string), &access.Token.AuditIds)
	assert.Equal(t, 1, len(access.Token.AuditIds))
	assert.Equal(t, "sample00d88246078f2bexample788f7", access.Token.AuditIds[0])

	assert.IsType(t, new([]*ServiceCatalog), &access.ServiceCatalog)
	assert.Equal(t, 1, len(access.ServiceCatalog))

	assert.IsType(t, new([]*Endpoint), &access.ServiceCatalog[0].Endpoints)
	assert.Equal(t, "region01", access.ServiceCatalog[0].Endpoints[0].Region)
	assert.Equal(t, "https://identity.region01.conoha.io/v2.0", access.ServiceCatalog[0].Endpoints[0].PublicURL)

	assert.IsType(t, new(User), access.User)
	assert.Equal(t, "example00000000", access.User.Username)

	assert.IsType(t, new([]*RolesLink), &access.User.RolesLinks)

	assert.Equal(t, "sample00d88246078f2bexample788f7", access.User.Id)

	assert.IsType(t, new([]*Role), &access.User.Roles)
	assert.Equal(t, "example00000000", access.User.Roles[0].Name)

	assert.Equal(t, "example00000000", access.User.Name)

	assert.IsType(t, new(IdentityTokenMetadata), access.Metadata)
	assert.Equal(t, 0, access.Metadata.IsAdmin)

	assert.IsType(t, new([]string), &access.Metadata.Roles)
	assert.Equal(t, "sample00d88246078f2bexample788f7", access.Metadata.Roles[0])

	assert.IsType(t, new(responseMeta), meta)
	assert.Equal(t, 200, meta.StatusCode)
}
