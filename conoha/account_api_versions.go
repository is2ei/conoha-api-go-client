package conoha

import "encoding/json"

type getAccountApiVersionsResponseParam struct {
	Versions []*Version `json:"versions"`
}

// AccountApiVersions fetches Account API versions list.
//
// ConoHa API docs: https://www.conoha.jp/docs/account-get_version_list.html
func (c *Conoha) AccountAPIVersions() ([]*Version, *ResponseMeta, error) {

	apiEndPoint := c.AccountServiceURL

	p := getAccountApiVersionsResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
