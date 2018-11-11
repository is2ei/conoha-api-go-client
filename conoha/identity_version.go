package conoha

import "encoding/json"

type getIdentityVersionResponseParam struct {
	Version Version `json:"version"`
}

func (c *Conoha) IdentityVersion() (Version, *ResponseMeta, error) {
	version := getIdentityVersionResponseParam{}

	u := c.IdentityServiceUrl + "/v2.0"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &version)
	}

	return version.Version, meta, err
}
