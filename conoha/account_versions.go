package conoha

import "encoding/json"

type getAccountVersionsResponseParam struct {
	Versions []Version `json:"versions"`
}

// AccountVersions fetches Account API versions.
func (c *Conoha) AccountVersions() ([]Version, *ResponseMeta, error) {
	p := getAccountVersionsResponseParam{}

	u := c.AccountServiceUrl

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
