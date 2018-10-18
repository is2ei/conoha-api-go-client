package conoha

import "encoding/json"

type GetImageVersionsResponseParam struct {
	Versions []Version `json:"versions"`
}

func (c *Conoha) ImageVersions() ([]Version, *ResponseMeta, error) {
	p := GetImageVersionsResponseParam{}

	u := c.ImageServiceUrl

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
