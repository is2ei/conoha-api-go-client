package conoha

import "encoding/json"

type getDNSVersionsResponseParam struct {
	Versions *struct {
		Values []*Version `json:"values"`
	} `json:"versions"`
}

// DnsVersions fetches DNS API versions list.
//
// ConoHa API docs: https://www.conoha.jp/docs/paas-dns-get-version-list.html
func (c *Conoha) DNSVersions() ([]*Version, *ResponseMeta, error) {
	p := getDNSVersionsResponseParam{}

	u := c.DNSServiceURL

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions.Values, meta, err
}
