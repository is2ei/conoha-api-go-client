package conoha

import "encoding/json"

type getImageVersionsResponseParam struct {
	Versions []Version `json:"versions"`
}

// ImageVersions fetches Image API versions list.
//
// ConoHa API docs: https://www.conoha.jp/docs/image-get_version_list.html
func (c *Conoha) ImageVersions() ([]Version, *responseMeta, error) {
	p := getImageVersionsResponseParam{}

	u := c.ImageServiceURL

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Versions, meta, err
}
