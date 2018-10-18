package conoha

import "encoding/json"

type GetDatabaseVersionResponseParam struct {
	Version Version `json:"version"`
}

func (c *Conoha) DatabaseVersion() (Version, *ResponseMeta, error) {
	p := GetDatabaseVersionResponseParam{}

	u := c.DatabaseServiceUrl + "/v1"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
