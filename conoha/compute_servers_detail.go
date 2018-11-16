package conoha

import "encoding/json"

type getComputeServersDetailResponseParam struct {
	Servers []*ComputeServer `json:"servers"`
}

// ComputeServersDetail fetches server's detailed information (e.g. IP address).
func (c *Conoha) ComputeServersDetail() ([]*ComputeServer, *ResponseMeta, error) {
	u := c.ComputeServiceUrl + "/v2/" + c.TenantId + "/servers/detail"

	p := getComputeServersDetailResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Servers, meta, err
}
