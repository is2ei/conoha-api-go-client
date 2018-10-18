package conoha

import "encoding/json"

type GetAccountNotificationResponseParam struct {
	Notification Notification `json:"notification"`
}

type Notification struct {
	NotificationCode int    `json:"notification_code"`
	Title            string `json:"title"`
	Type             string `json:"type"`
	Contents         string `json:"contents"`
	ReadStatus       string `json:"read_status"`
	StartDate        string `json:"start_date"`
}

func (c *Conoha) Notification(notificationCode string) (Notification, *ResponseMeta, error) {
	p := GetAccountNotificationResponseParam{}

	u := c.AccountServiceUrl + "/v1/" + c.TenantId + "/notifications/" + notificationCode

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Notification, meta, err
}
