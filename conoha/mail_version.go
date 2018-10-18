package conoha

import "encoding/json"

type GetMailVersionResponseParam struct {
	Version Version `json:"version"`
}

func (c *Conoha) MailVersion() (Version, *ResponseMeta, error) {
	p := GetMailVersionResponseParam{}

	u := c.MailServiceUrl + "/v1"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
