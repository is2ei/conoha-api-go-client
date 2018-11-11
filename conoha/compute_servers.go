package conoha

import "encoding/json"

type GetComputeServersResponseParam struct {
	Servers []*ComputeServer `json:"servers"`
}

func (c *Conoha) ComputeServers() ([]*ComputeServer, *ResponseMeta, error) {
	u := c.ComputeServiceUrl + "/v2/" + c.TenantId + "/servers"

	p := GetComputeServersResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Servers, meta, err
}