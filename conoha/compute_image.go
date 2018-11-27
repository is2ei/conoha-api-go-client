package conoha

import "encoding/json"

type getComputeImageResponseParam struct {
	Image *ComputeImage `json:"image"`
}

// ComputeImage represents the server information.
type ComputeImage struct {
	OsExtImgSize int                   `json:"OS-EXT-IMG-SIZE:size"`
	Created      string                `json:"created"`
	ID           string                `json:"id"`
	Links        []*Link               `json:"links"`
	Metadata     *ComputeImageMetadata `json:"metadata"`
	MinDisk      int                   `json:"minDisk"`
	MinRAM       int                   `json:"minRam"`
	Name         string                `json:"name"`
	Progress     int                   `json:"progress"`
	Status       string                `json:"status"`
	Updated      string                `json:"updated"`
}

// ComputeImageMetadata represents metadata for the server.
type ComputeImageMetadata struct {
	DisplayOrder     string `json:"display_order"`
	GncApp           string `json:"gnc_app"`
	HwQemuGuestAgent string `json:"hw_qemu_guest_agent"`
	OsType           string `json:"os_type"`
}

// ComputeImage fetches summarized information of the server.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_images_detail_specified.html
func (c *Conoha) ComputeImage(imageID string) (*ComputeImage, *ResponseMeta, error) {
	u := c.ComputeServiceURL + "/v2/" + c.TenantID + "/images/" + imageID

	p := getComputeImageResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Image, meta, err
}
