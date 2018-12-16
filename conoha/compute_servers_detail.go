package conoha

import (
	"context"
	"encoding/json"
	"fmt"
)

type getComputeServersDetailResponseParam struct {
	Servers []*ComputeServer `json:"servers"`
}

// ComputeServersDetail fetches server's detailed information (e.g. IP address).
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_vms_detail_specified.html
func (c *Conoha) ComputeServersDetail(ctx context.Context) ([]*ComputeServer, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v2/%s/servers/detail", c.ComputeServiceURL, c.TenantID)

	p := getComputeServersDetailResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, "GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Servers, meta, err
}
