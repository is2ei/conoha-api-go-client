package conoha

import (
	"bytes"
	"context"
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
	IdentityServiceURL      string
	AccountServiceURL       string
	ComputeServiceURL       string
	BlockStorageServiceURL  string
	ImageServiceURL         string
	NetworkServiceURL       string
	ObjectStorageServiceURL string
	DatabaseServiceURL      string
	DNSServiceURL           string
	MailServiceURL          string
	Username                string
	Password                string
	TenantID                string
	Token                   string
	Client                  *http.Client
}

// ResponseMeta represents response information (i.e status code).
type ResponseMeta struct {
	Method     string
	StatusCode int
}

// NewConoha creates ConoHa API client.
func NewConoha(
	identityServiceURL,
	accountServiceURL,
	computeServiceURL,
	blockStorageServiceURL,
	imageServiceURL,
	networkServiceURL,
	objectStorageServiceURL,
	databaseServiceURL,
	dnsServiceURL,
	mailServiceURL,
	username,
	password,
	tenantID,
	token string,
) *Conoha {

	client := &http.Client{}

	return &Conoha{
		IdentityServiceURL:      identityServiceURL,
		AccountServiceURL:       accountServiceURL,
		ComputeServiceURL:       computeServiceURL,
		BlockStorageServiceURL:  blockStorageServiceURL,
		ImageServiceURL:         imageServiceURL,
		NetworkServiceURL:       networkServiceURL,
		ObjectStorageServiceURL: objectStorageServiceURL,
		DatabaseServiceURL:      databaseServiceURL,
		DNSServiceURL:           dnsServiceURL,
		MailServiceURL:          mailServiceURL,
		Username:                username,
		Password:                password,
		TenantID:                tenantID,
		Token:                   token,
		Client:                  client,
	}
}

func withContext(ctx context.Context, req *http.Request) *http.Request {
	return req.WithContext(ctx)
}

func (c *Conoha) buildRequest(method, url string, body []byte) (*http.Request, error) {
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

	return req, err
}

func (c *Conoha) execRequest(ctx context.Context, req *http.Request) (*http.Response, error) {

	req = withContext(ctx, req)

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

func buildResponseMeta(resp *http.Response, method, u string) ResponseMeta {
	meta := ResponseMeta{}

	meta.Method = method
	meta.StatusCode = resp.StatusCode

	return meta
}

func (c *Conoha) buildAndExecRequest(ctx context.Context, method, u string, body []byte) ([]byte, *ResponseMeta, error) {

	req, err := c.buildRequest(method, u, body)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.execRequest(ctx, req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		default:
		}

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
