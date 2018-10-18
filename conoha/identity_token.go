package conoha

import (
	"encoding/json"
	"fmt"
)

type PostIdentityTokenRequestParam struct {
	Auth IdentityAuth `json:"auth"`
}

type IdentityAuth struct {
	IdentityPasswordCredentials IdentityPasswordCredentials `json:"passwordCredentials"`
	TenantId                    string                      `json:"tenantId"`
}

type IdentityPasswordCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostIdentityTokenResponseParam struct {
	Access Access `json:"access"`
}

type Access struct {
	Token          *Token                 `json:"token"`
	ServiceCatalog []*ServiceCatalog      `json:"serviceCatalog"`
	User           *User                  `json:"user"`
	Metadata       *IdentityTokenMetadata `json:"metadata"`
}

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

func (c *Conoha) IdentityToken() (*Access, *ResponseMeta, error) {
	passwordCredentials := IdentityPasswordCredentials{
		Username: c.Username,
		Password: c.Password,
	}
	auth := IdentityAuth{
		IdentityPasswordCredentials: passwordCredentials,
		TenantId:                    c.TenantId,
	}
	param := PostIdentityTokenRequestParam{
		Auth: auth,
	}

	u := fmt.Sprintf("%s/v2.0/tokens", c.IdentityServiceUrl)

	body, err := json.Marshal(param)

	contents, meta, err := c.buildAndExecRequest("POST", u, body)
	p := PostIdentityTokenResponseParam{}
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return &p.Access, meta, err
}
