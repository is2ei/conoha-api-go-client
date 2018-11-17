package conoha

import "encoding/json"

type getIdentityVersionResponseParam struct {
	Version Version `json:"version"`
}

// IdentityVersion fetches Identity API information.
//
// ConoHa API docs: https://www.conoha.jp/docs/identity-get_version_detail.html
func (c *Conoha) IdentityVersion() (Version, *responseMeta, error) {
	version := getIdentityVersionResponseParam{}

	u := c.IdentityServiceUrl + "/v2.0"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &version)
	}

	return version.Version, meta, err
}
