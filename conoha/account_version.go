package conoha

import "encoding/json"

type getAccountVersionResponseParam struct {
	Version Version `json:"version"`
}

// AccountVersion fetches the Account API information.
//
// ConoHa API docs: https://www.conoha.jp/docs/account-get_version_detail.html
func (c *Conoha) AccountVersion() (Version, *ResponseMeta, error) {
	p := getAccountVersionResponseParam{}

	u := c.AccountServiceURL + "/v1"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
