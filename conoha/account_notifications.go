package conoha

import "encoding/json"

type getNotificationsResponseParam struct {
	Notifications []Notification `json:"notifications"`
}

// Notifications fetches notifications list for the user.
//
// ConoHa API docs: https://www.conoha.jp/docs/account-informations-list.html
func (c *Conoha) Notifications() ([]Notification, *responseMeta, error) {
	p := getNotificationsResponseParam{}

	u := c.AccountServiceURL + "/v1/" + c.TenantId + "/notifications"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Notifications, meta, err
}
