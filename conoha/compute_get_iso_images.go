package conoha

import "encoding/json"

type getComputeIsoImagesResponseParam struct {
	IsoImages []*IsoImage `json:"iso-images"`
}

// IsoImage represents ISO image.
type IsoImage struct {
	Url   string `json:"url"`
	Path  string `json:"path"`
	Ctime string `json:"ctime"`
	Name  string `json:"name"`
	Size  int    `json:"size"`
}

// IsoImages fetches ISO images list.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-iso-list-show.html
func (c *Conoha) IsoImages() ([]*IsoImage, *ResponseMeta, error) {
	u := c.ComputeServiceURL + "/v2/" + c.TenantID + "/iso-images"

	p := getComputeIsoImagesResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.IsoImages, meta, err
}
