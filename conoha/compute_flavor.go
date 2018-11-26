package conoha

import "encoding/json"

type getComputeFlavorResponseParam struct {
	Flavor *ComputeFlavor `json:"flavor"`
}

// ComputeFlavor represents the plan's information.
type ComputeFlavor struct {
	OsFlvDisabled          bool    `json:"OS-FLV-DISABLED:disabled"`
	OsFlvExtData           int     `json:"OS-FLV-EXT-DATA:ephemeral"`
	Disk                   int     `json:"disk"`
	Id                     string  `json:"id"`
	Links                  []*Link `json:"links"`
	Name                   string  `json:"name"`
	OsFlavorAccessIsPublic bool    `json:"os-flavor-access:is_public"`
	Ram                    int     `json:"ram"`
	RxtxFactor             int     `json:"rxtx_factor"`
	Swap                   string  `json:"swap"`
	Vcpus                  int     `json:"vcpus"`
}

// ComputeFlavor fetches the plan's information.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_flavors_detail_specified.html
func (c *Conoha) ComputeFlavor(flavorId string) (*ComputeFlavor, *responseMeta, error) {
	u := c.ComputeServiceURL + "/v2/" + c.TenantId + "/flavors/" + flavorId

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)

	p := getComputeFlavorResponseParam{}
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Flavor, meta, err
}
