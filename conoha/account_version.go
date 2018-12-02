package conoha

import (
	"encoding/json"
	"fmt"
)

type getAccountVersionResponseParam struct {
	Version *Version `json:"version"`
}

// AccountVersion fetches the Account API information.
//
// ConoHa API docs: https://www.conoha.jp/docs/account-get_version_detail.html
func (c *Conoha) AccountVersion() (*Version, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v1", c.AccountServiceURL)

	p := getAccountVersionResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
