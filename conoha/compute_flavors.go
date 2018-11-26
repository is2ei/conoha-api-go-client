package conoha

import "encoding/json"

type getComputeFlavorsResponseParam struct {
	Flavors []*ComputeFlavor `json:"flavors"`
}

// ComputeFlavors fetches plans list.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_flavors_list.html
func (c *Conoha) ComputeFlavors() ([]*ComputeFlavor, *responseMeta, error) {
	u := c.ComputeServiceURL + "/v2/" + c.TenantID + "/flavors"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	p := getComputeFlavorsResponseParam{}
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Flavors, meta, err
}
