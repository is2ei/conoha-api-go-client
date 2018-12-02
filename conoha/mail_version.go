package conoha

import (
	"encoding/json"
	"fmt"
)

type getMailVersionResponseParam struct {
	Version *Version `json:"version"`
}

// MailVersion fetches Mail API information.
//
// ConoHa API docs: https://www.conoha.jp/docs/paas-mail-get-version-detail.html
func (c *Conoha) MailVersion() (*Version, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v1", c.MailServiceURL)

	p := getMailVersionResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
