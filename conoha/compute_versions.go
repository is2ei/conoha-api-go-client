package conoha

import "encoding/json"

type GetComputeVersionsResponseParam struct {
	Versions []Version `json:"versions"`
}

func (c *Conoha) ComputeVersions() ([]Version, *ResponseMeta, error) {
	p := GetComputeVersionsResponseParam{}

	u := c.ComputeServiceUrl

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
