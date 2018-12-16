package conoha

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type getMailVersionResponseParam struct {
	Version *Version `json:"version"`
}

// MailVersion fetches Mail API information.
//
// ConoHa API docs: https://www.conoha.jp/docs/paas-mail-get-version-detail.html
func (c *Conoha) MailVersion(ctx context.Context) (*Version, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v1", c.MailServiceURL)

	p := getMailVersionResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
