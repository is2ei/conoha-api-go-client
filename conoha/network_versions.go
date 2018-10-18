package conoha

import "encoding/json"

type GetNetworkVersionsResponseParam struct {
	Versions []Version `json:"versions"`
}

func (c *Conoha) NetworkVersions() ([]Version, *ResponseMeta, error) {
	p := GetNetworkVersionsResponseParam{}

	u := c.NetworkServiceUrl

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
