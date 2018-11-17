package conoha

import "encoding/json"

type getComputeServersDetailResponseParam struct {
	Servers []*ComputeServer `json:"servers"`
}

// ComputeServersDetail fetches server's detailed information (e.g. IP address).
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_vms_detail_specified.html
func (c *Conoha) ComputeServersDetail() ([]*ComputeServer, *responseMeta, error) {
	u := c.ComputeServiceUrl + "/v2/" + c.TenantId + "/servers/detail"

	p := getComputeServersDetailResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Servers, meta, err
}
