package conoha

import "encoding/json"

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
func (c *Conoha) ComputeKeypairs() ([]*KeypairParent, *ResponseMeta, error) {
	u := c.ComputeServiceURL + "/v2/" + c.TenantID + "/os-keypairs"

	p := getComputeKeypairsResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Keypairs, meta, err
}
