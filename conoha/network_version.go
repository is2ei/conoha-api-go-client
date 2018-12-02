package conoha

import (
	"encoding/json"
	"fmt"
)

type getNetworkVersionResponseParam struct {
	Version *Version `json:"version"`
}

// NetworkVersion fetches Network API information.
//
// ConoHa API docs: https://www.conoha.jp/docs/neutron-get_version_detail.html
func (c *Conoha) NetworkVersion() (*Version, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v2.0", c.NetworkServiceURL)

	p := getNetworkVersionResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
