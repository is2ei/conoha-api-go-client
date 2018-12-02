package conoha

import "encoding/json"

type getDNSVersionsResponseParam struct {
	Versions *struct {
		Values []*Version `json:"values"`
	} `json:"versions"`
}

// DNSVersions fetches DNS API versions list.
//
// ConoHa API docs: https://www.conoha.jp/docs/paas-dns-get-version-list.html
func (c *Conoha) DNSVersions() ([]*Version, *ResponseMeta, error) {

	apiEndPoint := c.DNSServiceURL

	p := getDNSVersionsResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions.Values, meta, err
}
