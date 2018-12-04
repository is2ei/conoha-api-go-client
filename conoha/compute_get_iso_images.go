package conoha

import (
	"encoding/json"
	"fmt"
)

type getComputeIsoImagesResponseParam struct {
	IsoImages []*IsoImage `json:"iso-images"`
}

// IsoImage represents ISO image.
type IsoImage struct {
	URL   string `json:"url"`
	Path  string `json:"path"`
	Ctime string `json:"ctime"`
	Name  string `json:"name"`
	Size  int    `json:"size"`
}

// IsoImages fetches ISO images list.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-iso-list-show.html
func (c *Conoha) IsoImages() ([]*IsoImage, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v2/%s/iso-images", c.ComputeServiceURL, c.TenantID)

	p := getComputeIsoImagesResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.IsoImages, meta, err
}
