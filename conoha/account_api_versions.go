package conoha

import "encoding/json"

type getAccountApiVersionsResponseParam struct {
	Versions []*Version `json:"versions"`
}

// AccountVersions fetches Account API versions list.
//
// ConoHa API docs: https://www.conoha.jp/docs/account-get_version_list.html
func (c *Conoha) AccountApiVersions() ([]*Version, *responseMeta, error) {

	apiEndPoint := c.AccountServiceUrl

	p := getAccountApiVersionsResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
