package conoha

import "encoding/json"

type getComputeIsoImagesResponseParam struct {
	IsoImages []*IsoImage `json:"iso-images"`
}

type IsoImage struct {
	Url   string `json:"url"`
	Path  string `json:"path"`
	Ctime string `json:"ctime"`
	Name  string `json:"name"`
	Size  int    `json:"size"`
}

func (c *Conoha) IsoImages() ([]*IsoImage, *ResponseMeta, error) {
	u := c.ComputeServiceUrl + "/v2/" + c.TenantId + "/iso-images"

	p := getComputeIsoImagesResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.IsoImages, meta, err
}
