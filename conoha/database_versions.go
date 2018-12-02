package conoha

import "encoding/json"

type getDatabaseVersionsResponseParam struct {
	Versions []*Version `json:"versions"`
}

// DatabaseVersions fetches Database API versions list.
//
// ConoHa API docs: https://www.conoha.jp/docs/paas-database-get-version-list.html
func (c *Conoha) DatabaseVersions() ([]*Version, *ResponseMeta, error) {

	apiEndPoint := c.DatabaseServiceURL

	p := getDatabaseVersionsResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
