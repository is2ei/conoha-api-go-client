package conoha

import "encoding/json"

type getComputeImagesResponseParam struct {
	Images []ComputeImage `json:"images"`
}

func (c *Conoha) ComputeImages() ([]ComputeImage, *ResponseMeta, error) {
	u := c.ComputeServiceUrl + "/v2/" + c.TenantId + "/images"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	p := getComputeImagesResponseParam{}
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Images, meta, err
}
