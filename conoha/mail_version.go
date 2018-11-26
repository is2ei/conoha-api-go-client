package conoha

import "encoding/json"

type getMailVersionResponseParam struct {
	Version Version `json:"version"`
}

// MailVersion fetches Mail API information.
//
// ConoHa API docs: https://www.conoha.jp/docs/paas-mail-get-version-detail.html
func (c *Conoha) MailVersion() (Version, *ResponseMeta, error) {
	p := getMailVersionResponseParam{}

	u := c.MailServiceURL + "/v1"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
