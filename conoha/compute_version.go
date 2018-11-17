package conoha

import "encoding/json"

type getComputeVersionResponseParam struct {
	Version Version `json:"version"`
}

// ComputeVersion fetches Compute API information.
func (c *Conoha) ComputeVersion() (Version, *responseMeta, error) {
	p := getComputeVersionResponseParam{}

	u := c.ComputeServiceUrl + "/v2"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
