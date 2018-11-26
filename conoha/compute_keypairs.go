package conoha

import "encoding/json"

type getComputeKeypairsResponseParam struct {
	Keypairs []*keypairParent `json:"keypairs"`
}

type keypairParent struct {
	Value *Keypair `json:"keypair"`
}

// ComputeKeypairs fetches key pairs list.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_keypairs.html
func (c *Conoha) ComputeKeypairs() ([]*keypairParent, *responseMeta, error) {
	u := c.ComputeServiceURL + "/v2/" + c.TenantId + "/os-keypairs"

	p := getComputeKeypairsResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Keypairs, meta, err
}
