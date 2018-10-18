package conoha

import "encoding/json"

type GetComputeVersionResponseParam struct {
	Version Version `json:"version"`
}

func (c *Conoha) ComputeVersion() (Version, *ResponseMeta, error) {
	p := GetComputeVersionResponseParam{}

	u := c.ComputeServiceUrl + "/v2"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
