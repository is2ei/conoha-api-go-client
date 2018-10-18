package conoha

import "encoding/json"

type GetComputeKeypairResponseParam struct {
	Keypair *Keypair `json:"keypair"`
}

type Keypair struct {
	PublicKey   string `json:"public_key"`
	UserId      string `json:"user_id"`
	Name        string `json:"name"`
	Deleted     bool   `json:"deleted"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	FingerPrint string `json:"fingerprint"`
	DeletedAt   string `json:"deleted_at"`
	Id          int    `json:"id"`
}

func (c *Conoha) ComputeKeypair(keypairName string) (*Keypair, *ResponseMeta, error) {
	u := c.ComputeServiceUrl + "/v2/" + c.TenantId + "/os-keypairs/" + keypairName

	p := GetComputeKeypairResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Keypair, meta, err
}
