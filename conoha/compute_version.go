package conoha

import "encoding/json"

type getComputeVersionResponseParam struct {
	Version Version `json:"version"`
}

func (c *Conoha) ComputeVersion() (Version, *ResponseMeta, error) {
	p := getComputeVersionResponseParam{}

	u := c.ComputeServiceUrl + "/v2"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
