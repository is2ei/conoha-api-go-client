package conoha

import "encoding/json"

type getComputeFlavorsResponseParam struct {
	Flavors []*ComputeFlavor `json:"flavors"`
}

// ComputeFlavors fetches plans list.
func (c *Conoha) ComputeFlavors() ([]*ComputeFlavor, *responseMeta, error) {
	u := c.ComputeServiceUrl + "/v2/" + c.TenantId + "/flavors"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	p := getComputeFlavorsResponseParam{}
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Flavors, meta, err
}
