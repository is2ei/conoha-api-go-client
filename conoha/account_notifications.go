package conoha

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type getNotificationsResponseParam struct {
	Notifications []*Notification `json:"notifications"`
}

// Notifications fetches notifications list for the user.
//
// ConoHa API docs: https://www.conoha.jp/docs/account-informations-list.html
func (c *Conoha) Notifications(ctx context.Context) ([]*Notification, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v1/%s/notifications", c.AccountServiceURL, c.TenantID)

	p := getNotificationsResponseParam{}

	contents, meta, err := c.buildAndExecRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Notifications, meta, err
}
