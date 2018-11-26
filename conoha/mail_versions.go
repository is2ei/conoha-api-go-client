package conoha

import "encoding/json"

type getMailVersionsResponseParam struct {
	Versions []Version `json:"versions"`
}

// MailVersions fetches Mail API versions list.
//
// ConoHa API docs: https://www.conoha.jp/docs/paas-mail-get-version-list.html
func (c *Conoha) MailVersions() ([]Version, *ResponseMeta, error) {
	p := getMailVersionsResponseParam{}

	u := c.MailServiceURL

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
