package conoha

import (
	"encoding/json"
)

type getIdentityAPIVersionsResponseParam struct {
	Versions identityVersionsValues `json:"versions"`
}

type identityVersionsValues struct {
	Values []*Version `json:"values"`
}

// IdentityAPIVersions fetches Identity API versions list.
//
// ConoHa API docs: https://www.conoha.jp/docs/identity-get_version_list.html
func (c *Conoha) IdentityAPIVersions() ([]*Version, *ResponseMeta, error) {

	apiEndPoint := c.IdentityServiceURL

	p := getIdentityAPIVersionsResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions.Values, meta, err
}
