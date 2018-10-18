package conoha

import "encoding/json"

type GetAccountVersionsResponseParam struct {
	Versions []Version `json:"versions"`
}

func (c *Conoha) AccountVersions() ([]Version, *ResponseMeta, error) {
	p := GetAccountVersionsResponseParam{}

	u := c.AccountServiceUrl

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
