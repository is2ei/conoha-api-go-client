package conoha

import "encoding/json"

type getImageVersionsResponseParam struct {
	Versions []Version `json:"versions"`
}

func (c *Conoha) ImageVersions() ([]Version, *responseMeta, error) {
	p := getImageVersionsResponseParam{}

	u := c.ImageServiceUrl

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
