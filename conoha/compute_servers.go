package conoha

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type getComputeServersResponseParam struct {
	Servers []*ComputeServer `json:"servers"`
}

// ComputeServers fetches servers list. Each information is summarized.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_vms_list.html
func (c *Conoha) ComputeServers(ctx context.Context) ([]*ComputeServer, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v2/%s/servers", c.ComputeServiceURL, c.TenantID)

	p := getComputeServersResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Servers, meta, err
}
