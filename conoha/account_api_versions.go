package conoha

import "encoding/json"

type getAccountAPIVersionsResponseParam struct {
	Versions []*Version `json:"versions,omitempty"`
}

// AccountAPIVersions fetches Account API versions list.
//
// ConoHa API docs: https://www.conoha.jp/docs/account-get_version_list.html
func (c *Conoha) AccountAPIVersions() ([]*Version, *ResponseMeta, error) {

	apiEndPoint := c.AccountServiceURL

	p := getAccountAPIVersionsResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", apiEndPoint, nil)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(contents, &p)
	if err != nil {
		return nil, nil, err
	}

	return p.Versions, meta, err
}
