package conoha

import "encoding/json"

type GetComputeImageResponseParam struct {
	Image *ComputeImage `json:"image"`
}

type ComputeImage struct {
	OsExtImgSize int                   `json:"OS-EXT-IMG-SIZE:size"`
	Created      string                `json:"created"`
	Id           string                `json:"id"`
	Links        []*Link               `json:"links"`
	Metadata     *ComputeImageMetadata `json:"metadata"`
	MinDisk      int                   `json:"minDisk"`
	MinRam       int                   `json:"minRam"`
	Name         string                `json:"name"`
	Progress     int                   `json:"progress"`
	Status       string                `json:"status"`
	Updated      string                `json:"updated"`
}

type ComputeImageMetadata struct {
	DisplayOrder     string `json:"display_order"`
	GncApp           string `json:"gnc_app"`
	HwQemuGuestAgent string `json:"hw_qemu_guest_agent"`
	OsType           string `json:"os_type"`
}

func (c *Conoha) ComputeImage(imageId string) (*ComputeImage, *ResponseMeta, error) {
	u := c.ComputeServiceUrl + "/v2/" + c.TenantId + "/images/" + imageId

	p := GetComputeImageResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Image, meta, err
}
