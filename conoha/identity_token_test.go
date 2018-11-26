package conoha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConoha_IdentityToken(t *testing.T) {
	ts := createMockServer(t, []string{
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

	a, meta, err := conoha.IdentityToken()

	assert.NoError(t, err)

	assert.IsType(t, new(access), a)

	assert.IsType(t, new(Token), a.Token)
	assert.IsType(t, "2015-05-19T07:08:21.927295", a.Token.IssuedAt)
	assert.IsType(t, "2015-05-20T07:08:21Z", a.Token.Expires)
	assert.IsType(t, "sample00d88246078f2bexample788f7", a.Token.Id)

	assert.IsType(t, new(tenant), a.Token.Tenant)
	assert.Equal(t, "did", a.Token.Tenant.DomainId)
	assert.Equal(t, "v2", a.Token.Tenant.Description)
	assert.Equal(t, true, a.Token.Tenant.Enabled)
	assert.Equal(t, "sample00d88246078f2bexample788f7", a.Token.Tenant.Id)
	assert.Equal(t, "example00000000", a.Token.Tenant.Name)

	assert.IsType(t, new([]string), &a.Token.AuditIds)
	assert.Equal(t, 1, len(a.Token.AuditIds))
	assert.Equal(t, "sample00d88246078f2bexample788f7", a.Token.AuditIds[0])

	assert.IsType(t, new([]*serviceCatalog), &a.ServiceCatalog)
	assert.Equal(t, 1, len(a.ServiceCatalog))

	assert.IsType(t, new([]*endpoint), &a.ServiceCatalog[0].Endpoints)
	assert.Equal(t, "region01", a.ServiceCatalog[0].Endpoints[0].Region)
	assert.Equal(t, "https://identity.region01.conoha.io/v2.0", a.ServiceCatalog[0].Endpoints[0].PublicURL)

	assert.IsType(t, new(user), a.User)
	assert.Equal(t, "example00000000", a.User.Username)

	assert.IsType(t, new([]*rolesLink), &a.User.RolesLinks)

	assert.Equal(t, "sample00d88246078f2bexample788f7", a.User.Id)

	assert.IsType(t, new([]*role), &a.User.Roles)
	assert.Equal(t, "example00000000", a.User.Roles[0].Name)

	assert.Equal(t, "example00000000", a.User.Name)

	assert.IsType(t, new(identityTokenMetadata), a.Metadata)
	assert.Equal(t, 0, a.Metadata.IsAdmin)

	assert.IsType(t, new([]string), &a.Metadata.Roles)
	assert.Equal(t, "sample00d88246078f2bexample788f7", a.Metadata.Roles[0])

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 200, meta.StatusCode)
}
