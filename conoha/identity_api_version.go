package conoha

import (
	"encoding/json"
	"fmt"
)

type getIdentityVersionResponseParam struct {
	Version *Version `json:"version"`
}

// IdentityAPIVersion fetches Identity API information.
//
// ConoHa API docs: https://www.conoha.jp/docs/identity-get_version_detail.html
func (c *Conoha) IdentityAPIVersion() (*Version, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v2.0", c.IdentityServiceURL)

	p := getIdentityVersionResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
