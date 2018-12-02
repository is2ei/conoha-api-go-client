package conoha

import "encoding/json"

type getComputeVersionsResponseParam struct {
	Versions []*Version `json:"versions"`
}

// ComputeVersions fetches Compute API versions list.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_version_list.html
func (c *Conoha) ComputeVersions() ([]*Version, *ResponseMeta, error) {

	apiEndPoint := c.ComputeServiceURL

	p := getComputeVersionsResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
