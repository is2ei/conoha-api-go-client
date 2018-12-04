package conoha

import (
	"encoding/json"
	"fmt"
)

type getComputeFlavorsDetailResponseParam struct {
	Flavors []*ComputeFlavor `json:"flavors"`
}

// ComputeFlavorsDetail fetches detailed plans list.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_flavors_detail.html
func (c *Conoha) ComputeFlavorsDetail() ([]*ComputeFlavor, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v2/%s/flavors/detail", c.ComputeServiceURL, c.TenantID)

	p := getComputeFlavorsDetailResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Flavors, meta, err
}
