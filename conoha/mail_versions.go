package conoha

import "encoding/json"

type GetMailVersionsResponseParam struct {
	Versions []Version `json:"versions"`
}

func (c *Conoha) MailVersions() ([]Version, *ResponseMeta, error) {
	p := GetMailVersionsResponseParam{}

	u := c.MailServiceUrl

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
