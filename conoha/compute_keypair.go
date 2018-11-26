package conoha

import "encoding/json"

type getComputeKeypairResponseParam struct {
	Keypair *Keypair `json:"keypair"`
}

// Keypair represents key pair information.
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

// ComputeKeypair fetches the key pair information.
//
// https://www.conoha.jp/docs/compute-get_keypairs_detail_specified.html
func (c *Conoha) ComputeKeypair(keypairName string) (*Keypair, *responseMeta, error) {
	u := c.ComputeServiceURL + "/v2/" + c.TenantID + "/os-keypairs/" + keypairName

	p := getComputeKeypairResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Keypair, meta, err
}
