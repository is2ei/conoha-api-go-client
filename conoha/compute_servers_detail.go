package conoha

import (
	"encoding/json"
	"fmt"
)

type getComputeServersDetailResponseParam struct {
	Servers []*ComputeServer `json:"servers"`
}

// ComputeServersDetail fetches server's detailed information (e.g. IP address).
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_vms_detail_specified.html
func (c *Conoha) ComputeServersDetail() ([]*ComputeServer, *ResponseMeta, error) {

	p := getComputeServersDetailResponseParam{}

	apiEndPoint := fmt.Sprintf("%s/v2/%s/servers/detail", c.ComputeServiceURL, c.TenantID)

	contents, meta, err := c.buildAndExecRequest("GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Servers, meta, err
}
