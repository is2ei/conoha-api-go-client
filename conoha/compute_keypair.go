package conoha

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type getComputeKeypairResponseParam struct {
	Keypair *Keypair `json:"keypair"`
}

// Keypair represents key pair information.
type Keypair struct {
	PublicKey   string `json:"public_key"`
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Deleted     bool   `json:"deleted"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	FingerPrint string `json:"fingerprint"`
	DeletedAt   string `json:"deleted_at"`
	ID          int    `json:"id"`
}

// ComputeKeypair fetches the key pair information.
//
// https://www.conoha.jp/docs/compute-get_keypairs_detail_specified.html
func (c *Conoha) ComputeKeypair(ctx context.Context, keypairName string) (*Keypair, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v2/%s/os-keypairs/%s", c.ComputeServiceURL, c.TenantID, keypairName)

	p := getComputeKeypairResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Keypair, meta, err
}
