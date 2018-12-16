package conoha

import (
	"context"
	"encoding/json"
	"net/http"
)

type getComputeVersionsResponseParam struct {
	Versions []*Version `json:"versions"`
}

// ComputeVersions fetches Compute API versions list.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_version_list.html
func (c *Conoha) ComputeVersions(ctx context.Context) ([]*Version, *ResponseMeta, error) {

	apiEndPoint := c.ComputeServiceURL

	p := getComputeVersionsResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
