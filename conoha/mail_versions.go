package conoha

import (
	"context"
	"encoding/json"
)

type getMailVersionsResponseParam struct {
	Versions []*Version `json:"versions"`
}

// MailVersions fetches Mail API versions list.
//
// ConoHa API docs: https://www.conoha.jp/docs/paas-mail-get-version-list.html
func (c *Conoha) MailVersions(ctx context.Context) ([]*Version, *ResponseMeta, error) {

	apiEndPoint := c.MailServiceURL

	p := getMailVersionsResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, "GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
