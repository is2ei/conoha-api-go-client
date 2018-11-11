package conoha

import "encoding/json"

type getNetworkVersionResponseParam struct {
	Version Version `json:"version"`
}

func (c *Conoha) NetworkVersion() (Version, *ResponseMeta, error) {
	p := getNetworkVersionResponseParam{}

	u := c.NetworkServiceUrl + "/v2.0"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Version, meta, err
}
