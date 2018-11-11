package conoha

import "encoding/json"

type getDnsVersionsResponseParam struct {
	Versions dnsVersionsValues `json:"versions"`
}

type dnsVersionsValues struct {
	Values []Version `json:"values"`
}

func (c *Conoha) DnsVersions() ([]Version, *ResponseMeta, error) {
	p := getDnsVersionsResponseParam{}

	u := c.DnsServiceUrl

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions.Values, meta, err
}
