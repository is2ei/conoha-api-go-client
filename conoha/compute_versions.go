package conoha

import "encoding/json"

type getComputeVersionsResponseParam struct {
	Versions []Version `json:"versions"`
}

func (c *Conoha) ComputeVersions() ([]Version, *ResponseMeta, error) {
	p := getComputeVersionsResponseParam{}

	u := c.ComputeServiceUrl

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
