package conoha

import (
	"context"
	"encoding/json"
)

type getNetworkVersionsResponseParam struct {
	Versions []*Version `json:"versions"`
}

// NetworkVersions fetches Network API versions list.
//
// ConoHa API docs: https://www.conoha.jp/docs/neutron-get_version_list.html
func (c *Conoha) NetworkVersions(ctx context.Context) ([]*Version, *ResponseMeta, error) {

	apiEndPoint := c.NetworkServiceURL

	p := getNetworkVersionsResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, "GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
