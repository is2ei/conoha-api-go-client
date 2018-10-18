package conoha

import (
	"encoding/json"
)

type GetIdentityVersionsResponseParam struct {
	Versions IdentityVersionsValues `json:"versions"`
}

type IdentityVersionsValues struct {
	Values []Version `json:"values"`
}

func (c *Conoha) IdentityVersions() ([]Version, *ResponseMeta, error) {
	p := GetIdentityVersionsResponseParam{}

	u := c.IdentityServiceUrl

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions.Values, meta, err
}
