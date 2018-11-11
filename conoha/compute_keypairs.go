package conoha

import "encoding/json"

type getComputeKeypairsResponseParam struct {
	Keypairs []*KeypairParent `json:"keypairs"`
}

type KeypairParent struct {
	Value *Keypair `json:"keypair"`
}

func (c *Conoha) ComputeKeypairs() ([]*KeypairParent, *ResponseMeta, error) {
	u := c.ComputeServiceUrl + "/v2/" + c.TenantId + "/os-keypairs"

	p := getComputeKeypairsResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Keypairs, meta, err
}
