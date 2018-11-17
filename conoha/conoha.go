package conoha

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
)

// Conoha represents ConoHa APi Client.
//
// ConoHa API docs: https://www.conoha.jp/docs/
type Conoha struct {
	IdentityServiceUrl      string
	AccountServiceUrl       string
	ComputeServiceUrl       string
	BlockStorageServiceUrl  string
	ImageServiceUrl         string
	NetworkServiceUrl       string
	ObjectStorageServiceUrl string
	DatabaseServiceUrl      string
	DnsServiceUrl           string
	MailServiceUrl          string
	Username                string
	Password                string
	TenantId                string
	Token                   string
	Client                  *http.Client
}

type responseMeta struct {
	Method     string
	StatusCode int
}

// NewConoha creates ConoHa API client.
func NewConoha(
	identityServiceUrl,
	accountServiceUrl,
	computeServiceUrl,
	blockStorageServiceUrl,
	imageServiceUrl,
	networkServiceUrl,
	objectStorageServiceUrl,
	databaseServiceUrl,
	dnsServiceUrl,
	mailServiceUrl,
	username,
	password,
	tenantId,
	token string,
) *Conoha {

	client := &http.Client{}

	return &Conoha{
		IdentityServiceUrl:      identityServiceUrl,
		AccountServiceUrl:       accountServiceUrl,
		ComputeServiceUrl:       computeServiceUrl,
		BlockStorageServiceUrl:  blockStorageServiceUrl,
		ImageServiceUrl:         imageServiceUrl,
		NetworkServiceUrl:       networkServiceUrl,
		ObjectStorageServiceUrl: objectStorageServiceUrl,
		DatabaseServiceUrl:      databaseServiceUrl,
		DnsServiceUrl:           dnsServiceUrl,
		MailServiceUrl:          mailServiceUrl,
		Username:                username,
		Password:                password,
		TenantId:                tenantId,
		Token:                   token,
		Client:                  client,
	}
}

func (c *Conoha) execRequest(method, url string, body []byte) (*http.Response, error) {
	var req *http.Request
	var err error

	if body != nil {
		req, err = http.NewRequest(method, url, bytes.NewReader(body))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Auth-Token", c.Token)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Client.Do error: %q", err)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		if resp.StatusCode == 401 {
			body, _ := ioutil.ReadAll(resp.Body)
			e := errorResponseBody{}
			json.Unmarshal(body, &e)
			_, file, line, _ := runtime.Caller(1)
			var msg string
			if e.Error != nil {
				msg = fmt.Sprintf("Unauthorized Request: URL[%s], Message[%s] at %s line %d", req.URL, e.Error.Message, file, line)
			} else if e.Unauthorized != nil {
				msg = fmt.Sprintf("Unauthorized Request: URL[%s], Message[%s] at %s line %d", req.URL, e.Unauthorized.Message, file, line)
			} else {
				msg = string(body)
			}
			err = newErrorUnauthorized(msg)
		} else {
			err = fmt.Errorf("Request failed: <%d> %s %s", resp.StatusCode, req.Method, req.URL)
		}
	}

	return resp, err
}

func buildResponseMeta(resp *http.Response, method, u string) responseMeta {
	meta := responseMeta{}

	meta.Method = method
	meta.StatusCode = resp.StatusCode

	return meta
}

func (c *Conoha) buildAndExecRequest(method, u string, body []byte) ([]byte, *responseMeta, error) {
	resp, err := c.execRequest(method, u, body)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	meta := buildResponseMeta(resp, method, u)

	return contents, &meta, err
}
