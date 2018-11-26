package conoha

import "encoding/json"

type getDatabaseVersionResponseParam struct {
	Version Version `json:"version"`
}

// DatabaseVersion fetches Database API information.
//
// ConoHa API docs: https://www.conoha.jp/docs/paas-database-get-version-detail.html
func (c *Conoha) DatabaseVersion() (Version, *ResponseMeta, error) {
	p := getDatabaseVersionResponseParam{}

	u := c.DatabaseServiceURL + "/v1"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
