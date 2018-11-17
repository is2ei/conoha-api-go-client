package conoha

import "encoding/json"

type getNetworkVersionsResponseParam struct {
	Versions []Version `json:"versions"`
}

func (c *Conoha) NetworkVersions() ([]Version, *responseMeta, error) {
	p := getNetworkVersionsResponseParam{}

	u := c.NetworkServiceUrl

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
