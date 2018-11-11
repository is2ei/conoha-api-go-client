package conoha

import "encoding/json"

type getMailVersionsResponseParam struct {
	Versions []Version `json:"versions"`
}

func (c *Conoha) MailVersions() ([]Version, *ResponseMeta, error) {
	p := getMailVersionsResponseParam{}

	u := c.MailServiceUrl

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
