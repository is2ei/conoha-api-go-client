package conoha

import (
	"context"
	"reflect"
	"testing"

	"github.com/k0kubun/pp"
)

func TestConoha_AccountApiVersions(t *testing.T) {
	ts := createMockServer(t, []string{
		"/account/versions",
	})
	defer ts.Close()

	conoha := NewConoha(
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		ts.URL,
		"username",
		"password",
		"tenant_id",
		"token",
	)

	got, _, err := conoha.AccountAPIVersions(context.Background())
	if err != nil {
		t.Errorf("Conoha.AccountAPIVersions returned error: %v", err)
	}
	if want := wantVersions; !reflect.DeepEqual(got, want) {
		t.Errorf("Conoha.AccountAPIVersions returned unexpected result.\n got: %s\nwant: %s\n", pp.Sprint(got), pp.Sprint(want))
	}
}

var wantVersions = []*Version{
	&Version{
		Status:  "CURRENT",
		Updated: "2015-05-12T09:00:00Z",
		ID:      "v1.0",
		Links: []*Link{
			&Link{
				Href: "https://account.tyo1.conoha.io/v1/",
				Type: "",
				Rel:  "self",
			},
		},
	},
}
