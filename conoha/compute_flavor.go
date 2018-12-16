package conoha

import (
	"context"
	"encoding/json"
	"fmt"
)

type getComputeFlavorResponseParam struct {
	Flavor *ComputeFlavor `json:"flavor"`
}

// ComputeFlavor represents the plan's information.
type ComputeFlavor struct {
	OsFlvDisabled          bool    `json:"OS-FLV-DISABLED:disabled"`
	OsFlvExtData           int     `json:"OS-FLV-EXT-DATA:ephemeral"`
	Disk                   int     `json:"disk"`
	ID                     string  `json:"id"`
	Links                  []*Link `json:"links"`
	Name                   string  `json:"name"`
	OsFlavorAccessIsPublic bool    `json:"os-flavor-access:is_public"`
	RAM                    int     `json:"ram"`
	RxtxFactor             int     `json:"rxtx_factor"`
	Swap                   string  `json:"swap"`
	Vcpus                  int     `json:"vcpus"`
}

// ComputeFlavor fetches the plan's information.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_flavors_detail_specified.html
func (c *Conoha) ComputeFlavor(ctx context.Context, flavorID string) (*ComputeFlavor, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v2/%s/flavors/%s", c.ComputeServiceURL, c.TenantID, flavorID)

	p := getComputeFlavorResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, "GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Flavor, meta, err
}
