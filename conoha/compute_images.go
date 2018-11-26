package conoha

import "encoding/json"

type getComputeImagesResponseParam struct {
	Images []ComputeImage `json:"images"`
}

// ComputeImages fetches server images list.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_images_list.html
func (c *Conoha) ComputeImages() ([]ComputeImage, *responseMeta, error) {
	u := c.ComputeServiceURL + "/v2/" + c.TenantId + "/images"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	p := getComputeImagesResponseParam{}
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Images, meta, err
}
