package conoha

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type getComputeFlavorsResponseParam struct {
	Flavors []*ComputeFlavor `json:"flavors"`
}

// ComputeFlavors fetches plans list.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_flavors_list.html
func (c *Conoha) ComputeFlavors(ctx context.Context) ([]*ComputeFlavor, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v2/%s/flavors", c.ComputeServiceURL, c.TenantID)

	p := getComputeFlavorsResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Flavors, meta, err
}
