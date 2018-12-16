package conoha

import (
	"context"
	"encoding/json"
)

type getAccountAPIVersionsResponseParam struct {
	Versions []*Version `json:"versions,omitempty"`
}

// AccountAPIVersions fetches Account API versions list.
//
// ConoHa API docs: https://www.conoha.jp/docs/account-get_version_list.html
func (c *Conoha) AccountAPIVersions(ctx context.Context) ([]*Version, *ResponseMeta, error) {

	apiEndPoint := c.AccountServiceURL

	p := getAccountAPIVersionsResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, "GET", apiEndPoint, nil)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(contents, &p)
	if err != nil {
		return nil, nil, err
	}

	return p.Versions, meta, err
}
