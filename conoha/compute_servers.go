package conoha

import "encoding/json"

type getComputeServersResponseParam struct {
	Servers []*ComputeServer `json:"servers"`
}

// ComputeServers fetches servers list. Each information is summarized.
func (c *Conoha) ComputeServers() ([]*ComputeServer, *ResponseMeta, error) {
	u := c.ComputeServiceUrl + "/v2/" + c.TenantId + "/servers"

	p := getComputeServersResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Servers, meta, err
}
