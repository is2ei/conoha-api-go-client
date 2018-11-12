package conoha

import "encoding/json"

type getDatabaseVersionsResponseParam struct {
	Versions []Version `json:"versions"`
}

func (c *Conoha) DatabaseVersions() ([]Version, *ResponseMeta, error) {
	p := getDatabaseVersionsResponseParam{}

	u := c.DatabaseServiceUrl

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
