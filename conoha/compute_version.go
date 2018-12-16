package conoha

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type getComputeVersionResponseParam struct {
	Version *Version `json:"version"`
}

// ComputeVersion fetches Compute API information.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_version_detail.html
func (c *Conoha) ComputeVersion(ctx context.Context) (*Version, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v2", c.ComputeServiceURL)

	p := getComputeVersionResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
