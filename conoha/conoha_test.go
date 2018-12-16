package conoha

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mapping struct {
	Request struct {
		Method string            `json:"method"`
		Path   string            `json:"path"`
		Query  map[string]string `json:"query"`
	} `json:"request"`
	Response struct {
		Status       int               `json:"status"`
		BodyFileName string            `json:"bodyFileName"`
		BodyContent  string            `json:"bodyContent"`
		Headers      map[string]string `json:"headers"`
	} `json:"response"`
}

func (m *mapping) Match(r *http.Request) bool {
	if m.Request.Method != r.Method {
		return false
	}

	if m.Request.Path != r.URL.Path {
		return false
	}

	for k, v := range m.Request.Query {
		if r.URL.Query().Get(k) != v {
			return false
		}
	}

	return true
}

func baseDir(t *testing.T) string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func getMocksDirs(t *testing.T) (string, string) {
	d := baseDir(t)
	mocksDir := filepath.Join(filepath.Join(d, ".."), "mocks")
	return filepath.Join(mocksDir, "mappings"), filepath.Join(mocksDir, "files")
}

func serverFromMappings(t *testing.T, mappings []*mapping) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, m := range mappings {
			if m.Match(r) {
				for k, v := range m.Response.Headers {
					w.Header().Set(k, v)
				}
				w.WriteHeader(m.Response.Status)
				w.Write([]byte(m.Response.BodyContent))
				return
			}

			t.Error("no request found matching: \newery request has to be mocked!")
		}
		fmt.Fprintln(w, "Hello, World")
	}))

	return ts
}

func createMockServer(t *testing.T, mappingPaths []string) *httptest.Server {
	mappingsDir, filesDir := getMocksDirs(t)

	var mappings []*mapping
	for _, p := range mappingPaths {
		f := filepath.Join(mappingsDir, fmt.Sprintf("%s.json", p))
		contents, err := ioutil.ReadFile(f)
		if err != nil {
			t.Fatal(fmt.Sprintf("an error occurred while loading mapping file %s\n%v", f, err))
			return nil
		}

		mapping := mapping{}
		err = json.Unmarshal(contents, &mapping)
		if err != nil {
			t.Fatal(fmt.Sprintf("an error occurred while unmarshalling mapping %s\n%v", f, err))
			return nil
		}

		if mapping.Response.BodyFileName != "" {
			b := filepath.Join(filesDir, mapping.Response.BodyFileName)
			body, err := ioutil.ReadFile(b)
			if err != nil {
				t.Fatal(fmt.Sprintf("an error occurred while loading mapping response body: %s\n%v", b, err))
				return nil
			}

			mapping.Response.BodyContent = string(body)
		}

		mappings = append(mappings, &mapping)
	}

	return serverFromMappings(t, mappings)
}

func TestNewConoha(t *testing.T) {
	conoha := NewConoha(
		"http://identity_service_url",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"username",
		"password",
		"tenant_id",
		"token",
	)

	assert.IsType(t, new(Conoha), conoha)
	assert.Equal(
		t,
		conoha.IdentityServiceURL,
		"http://identity_service_url",
	)
	assert.Equal(
		t,
		conoha.Token,
		"token",
	)
}

func TestClientDoError(t *testing.T) {
	conoha := NewConoha(
		"http://identity_service_url",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"http://",
		"username",
		"password",
		"tenant_id",
		"token",
	)

	req, _ := conoha.buildRequest("AAA", "aaa", nil)
	_, err := conoha.execRequest(context.Background(), req)

	assert.EqualError(t, err, "Client.Do error: \"Aaa aaa: unsupported protocol scheme \\\"\\\"\"")
}
