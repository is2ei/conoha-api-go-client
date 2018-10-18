package conoha

import "encoding/json"

type GetDnsVersionsResponseParam struct {
	Versions DnsVersionsValues `json:"versions"`
}

type DnsVersionsValues struct {
	Values []Version `json:"values"`
}

func (c *Conoha) DnsVersions() ([]Version, *ResponseMeta, error) {
	p := GetDnsVersionsResponseParam{}

	u := c.DnsServiceUrl

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions.Values, meta, err
}
