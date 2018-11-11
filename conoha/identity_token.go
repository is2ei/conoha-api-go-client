package conoha

import (
	"encoding/json"
	"fmt"
)

type postIdentityTokenRequestParam struct {
	Auth identityAuth `json:"auth"`
}

type identityAuth struct {
	IdentityPasswordCredentials identityPasswordCredentials `json:"passwordCredentials"`
	TenantId                    string                      `json:"tenantId"`
}

type identityPasswordCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type postIdentityTokenResponseParam struct {
	Access Access `json:"access"`
}

type Access struct {
	Token          *Token                 `json:"token"`
	ServiceCatalog []*ServiceCatalog      `json:"serviceCatalog"`
	User           *User                  `json:"user"`
	Metadata       *IdentityTokenMetadata `json:"metadata"`
}

// Token represents the ConoHa API access token.
type Token struct {
	IssuedAt string   `json:"issued_at"`
	Expires  string   `json:"expires"`
	Id       string   `json:"id"`
	Tenant   *Tenant  `json:"tenant"`
	AuditIds []string `json:"audit_ids"`
}

type Tenant struct {
	DomainId    string `json:"domain_id"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Id          string `json:"id"`
	Name        string `json:"name"`
}

type ServiceCatalog struct {
	Endpoints []*Endpoint `json:"endpoints"`
}

type Endpoint struct {
	Region    string `json:"region"`
	PublicURL string `json:"publicURL"`
}

// User represents the user.
type User struct {
	Username   string       `json:"username"`
	RolesLinks []*RolesLink `json:"roles_links"`
	Id         string       `json:"id"`
	Roles      []*Role      `json:"roles"`
	Name       string       `json:"name"`
}

type RolesLink struct {
}

type Role struct {
	Name string `json:"name"`
}

type IdentityTokenMetadata struct {
	IsAdmin int      `json:"is_admin"`
	Roles   []string `json:"roles"`
}

// IdentityToken issues ConoHa API access token.
func (c *Conoha) IdentityToken() (*Access, *ResponseMeta, error) {
	passwordCredentials := identityPasswordCredentials{
		Username: c.Username,
		Password: c.Password,
	}
	auth := identityAuth{
		IdentityPasswordCredentials: passwordCredentials,
		TenantId:                    c.TenantId,
	}
	param := postIdentityTokenRequestParam{
		Auth: auth,
	}

	u := fmt.Sprintf("%s/v2.0/tokens", c.IdentityServiceUrl)

	body, err := json.Marshal(param)

	contents, meta, err := c.buildAndExecRequest("POST", u, body)
	p := postIdentityTokenResponseParam{}
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return &p.Access, meta, err
}
