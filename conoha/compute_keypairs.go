package conoha

import (
	"context"
	"encoding/json"
	"fmt"
)

type getComputeKeypairsResponseParam struct {
	Keypairs []*KeypairParent `json:"keypairs"`
}

// KeypairParent represents keypair data wrapper.
type KeypairParent struct {
	Value *Keypair `json:"keypair"`
}

// ComputeKeypairs fetches key pairs list.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_keypairs.html
func (c *Conoha) ComputeKeypairs(ctx context.Context) ([]*KeypairParent, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v2/%s/os-keypairs", c.ComputeServiceURL, c.TenantID)

	p := getComputeKeypairsResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, "GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Keypairs, meta, err
}
