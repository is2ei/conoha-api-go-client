package conoha

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type getDatabaseVersionResponseParam struct {
	Version *Version `json:"version"`
}

// DatabaseVersion fetches Database API information.
//
// ConoHa API docs: https://www.conoha.jp/docs/paas-database-get-version-detail.html
func (c *Conoha) DatabaseVersion(ctx context.Context) (*Version, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v1", c.DatabaseServiceURL)

	p := getDatabaseVersionResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
