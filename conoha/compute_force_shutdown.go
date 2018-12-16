package conoha

import (
	"context"
	"encoding/json"
	"fmt"
)

type computeForceShutdownRequestParam struct {
	OsStop *osStop `json:"os-stop"`
}

type osStop struct {
	ForceShutdown bool `json:"force_shutdown"`
}

// ForceShutdownServer shutdowns the server forcefully.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-stop_forcibly_vm.html
func (c *Conoha) ForceShutdownServer(ctx context.Context, serverID string) (*ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v2/%s/servers/%s/action", c.ComputeServiceURL, c.TenantID, serverID)

	o := osStop{
		ForceShutdown: true,
	}

	p := computeForceShutdownRequestParam{
		OsStop: &o,
	}

	body, _ := json.Marshal(p)

	_, meta, err := c.buildAndExecRequest(ctx, "POST", apiEndPoint, body)

	return meta, err
}
