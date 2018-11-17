package conoha

import (
	"encoding/json"
)

type getIdentityVersionsResponseParam struct {
	Versions identityVersionsValues `json:"versions"`
}

type identityVersionsValues struct {
	Values []Version `json:"values"`
}

// IdentityVersions fetches Identity API versions list.
//
// ConoHa API docs: https://www.conoha.jp/docs/identity-get_version_list.html
func (c *Conoha) IdentityVersions() ([]Version, *responseMeta, error) {
	p := getIdentityVersionsResponseParam{}

	u := c.IdentityServiceUrl

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions.Values, meta, err
}
