package conoha

import "encoding/json"

type getAccountVersionResponseParam struct {
	Version Version `json:"version"`
}

func (c *Conoha) AccountVersion() (Version, *ResponseMeta, error) {
	p := getAccountVersionResponseParam{}

	u := c.AccountServiceUrl + "/v1"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
