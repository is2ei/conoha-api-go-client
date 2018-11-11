package conoha

import "encoding/json"

type getComputeFlavorsResponseParam struct {
	Flavors []*ComputeFlavor `json:"flavors"`
}

func (c *Conoha) ComputeFlavors() ([]*ComputeFlavor, *ResponseMeta, error) {
	u := c.ComputeServiceUrl + "/v2/" + c.TenantId + "/flavors"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	p := getComputeFlavorsResponseParam{}
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Flavors, meta, err
}
