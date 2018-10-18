package conoha

import "encoding/json"

type GetNotificationsResponseParam struct {
	Notifications []Notification `json:"notifications"`
}

func (c *Conoha) Notifications() ([]Notification, *ResponseMeta, error) {
	p := GetNotificationsResponseParam{}

	u := c.AccountServiceUrl + "/v1/" + c.TenantId + "/notifications"

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Notifications, meta, err
}
