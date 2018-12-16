package conoha

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type getComputeImagesResponseParam struct {
	Images []*ComputeImage `json:"images"`
}

// ComputeImages fetches server images list.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_images_list.html
func (c *Conoha) ComputeImages(ctx context.Context) ([]*ComputeImage, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v2/%s/images", c.ComputeServiceURL, c.TenantID)

	p := getComputeImagesResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Images, meta, err
}
