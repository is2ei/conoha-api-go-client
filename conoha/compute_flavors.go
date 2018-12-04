package conoha

import (
	"encoding/json"
	"fmt"
)

type getComputeFlavorsResponseParam struct {
	Flavors []*ComputeFlavor `json:"flavors"`
}

// ComputeFlavors fetches plans list.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_flavors_list.html
func (c *Conoha) ComputeFlavors() ([]*ComputeFlavor, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v2/%s/flavors", c.ComputeServiceURL, c.TenantID)

	p := getComputeFlavorsResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Flavors, meta, err
}
