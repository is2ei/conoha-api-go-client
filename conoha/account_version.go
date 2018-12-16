package conoha

import (
	"context"
	"encoding/json"
	"fmt"
)

type getAccountVersionResponseParam struct {
	Version *Version `json:"version"`
}

// AccountVersion fetches the Account API information.
//
// ConoHa API docs: https://www.conoha.jp/docs/account-get_version_detail.html
func (c *Conoha) AccountVersion(ctx context.Context) (*Version, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v1", c.AccountServiceURL)

	p := getAccountVersionResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, "GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
