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
	Access access `json:"access"`
}

type access struct {
	Token          *Token                 `json:"token"`
	ServiceCatalog []*serviceCatalog      `json:"serviceCatalog"`
	User           *User                  `json:"user"`
	Metadata       *identityTokenMetadata `json:"metadata"`
}

// Token represents the ConoHa API access token.
type Token struct {
	IssuedAt string   `json:"issued_at"`
	Expires  string   `json:"expires"`
	Id       string   `json:"id"`
	Tenant   *tenant  `json:"tenant"`
	AuditIds []string `json:"audit_ids"`
}

type tenant struct {
	DomainId    string `json:"domain_id"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Id          string `json:"id"`
	Name        string `json:"name"`
}

type serviceCatalog struct {
	Endpoints []*endpoint `json:"endpoints"`
}

type endpoint struct {
	Region    string `json:"region"`
	PublicURL string `json:"publicURL"`
}

// User represents the user.
type User struct {
	Username   string       `json:"username"`
	RolesLinks []*rolesLink `json:"roles_links"`
	Id         string       `json:"id"`
	Roles      []*role      `json:"roles"`
	Name       string       `json:"name"`
}

type rolesLink struct {
}

type role struct {
	Name string `json:"name"`
}

type identityTokenMetadata struct {
	IsAdmin int      `json:"is_admin"`
	Roles   []string `json:"roles"`
}

// IdentityToken issues ConoHa API access token.
//
// ConoHa API docs: https://www.conoha.jp/docs/identity-post_tokens.html
func (c *Conoha) IdentityToken() (*access, *responseMeta, error) {
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

	body, _ := json.Marshal(param)

	contents, meta, err := c.buildAndExecRequest("POST", u, body)
	p := postIdentityTokenResponseParam{}
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return &p.Access, meta, err
}
