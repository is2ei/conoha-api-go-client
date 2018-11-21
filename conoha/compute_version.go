package conoha

import (
	"encoding/json"
	"fmt"
)

type getComputeVersionResponseParam struct {
	Version Version `json:"version"`
}

// ComputeVersion fetches Compute API information.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_version_detail.html
func (c *Conoha) ComputeVersion() (Version, *responseMeta, error) {
	p := getComputeVersionResponseParam{}

	u := fmt.Sprintf("%s/v2", c.ComputeServiceUrl)

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
