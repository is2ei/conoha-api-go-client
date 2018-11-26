package conoha

import "encoding/json"

type getComputeServersResponseParam struct {
	Servers []*ComputeServer `json:"servers"`
}

// ComputeServers fetches servers list. Each information is summarized.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_vms_list.html
func (c *Conoha) ComputeServers() ([]*ComputeServer, *responseMeta, error) {
	u := c.ComputeServiceURL + "/v2/" + c.TenantID + "/servers"

	p := getComputeServersResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Servers, meta, err
}
