package conoha

import "encoding/json"

type GetComputeServersDetailResponseParam struct {
	Servers []*ComputeServer `json:"servers"`
}

func (c *Conoha) ComputeServersDetail() ([]*ComputeServer, *ResponseMeta, error) {
	u := c.ComputeServiceUrl + "/v2/" + c.TenantId + "/servers/detail"

	p := GetComputeServersDetailResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Servers, meta, err
}
