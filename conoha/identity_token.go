package conoha

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type postIdentityTokenRequestParam struct {
	Auth identityAuth `json:"auth"`
}

type identityAuth struct {
	IdentityPasswordCredentials identityPasswordCredentials `json:"passwordCredentials"`
	TenantID                    string                      `json:"tenantId"`
}

type identityPasswordCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type postIdentityTokenResponseParam struct {
	Access Access `json:"access"`
}

// Access represents data contains token and metadata.
type Access struct {
	Token          *Token                 `json:"token"`
	ServiceCatalog []*serviceCatalog      `json:"serviceCatalog"`
	User           *user                  `json:"user"`
	Metadata       *identityTokenMetadata `json:"metadata"`
}

// Token represents the ConoHa API access token.
type Token struct {
	IssuedAt string   `json:"issued_at"`
	Expires  string   `json:"expires"`
	ID       string   `json:"id"`
	Tenant   *tenant  `json:"tenant"`
	AuditIds []string `json:"audit_ids"`
}

type tenant struct {
	DomainID    string `json:"domain_id"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}

type serviceCatalog struct {
	Endpoints []*endpoint `json:"endpoints"`
}

type endpoint struct {
	Region    string `json:"region"`
	PublicURL string `json:"publicURL"`
}

type user struct {
	Username   string       `json:"username"`
	RolesLinks []*rolesLink `json:"roles_links"`
	ID         string       `json:"id"`
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
func (c *Conoha) IdentityToken(ctx context.Context) (*Access, *ResponseMeta, error) {
	passwordCredentials := identityPasswordCredentials{
		Username: c.Username,
		Password: c.Password,
	}
	auth := identityAuth{
		IdentityPasswordCredentials: passwordCredentials,
		TenantID:                    c.TenantID,
	}
	param := postIdentityTokenRequestParam{
		Auth: auth,
	}

	apiEndPoint := fmt.Sprintf("%s/v2.0/tokens", c.IdentityServiceURL)

	body, _ := json.Marshal(param)

	p := postIdentityTokenResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, http.MethodPost, apiEndPoint, body)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return &p.Access, meta, err
}
