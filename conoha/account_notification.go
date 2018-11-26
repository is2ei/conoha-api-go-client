package conoha

import "encoding/json"

type getAccountNotificationResponseParam struct {
	Notification Notification `json:"notification"`
}

// Notification represents notification from ConoHa to the users.
type Notification struct {
	NotificationCode int    `json:"notification_code"`
	Title            string `json:"title"`
	Type             string `json:"type"`
	Contents         string `json:"contents"`
	ReadStatus       string `json:"read_status"`
	StartDate        string `json:"start_date"`
}

// Notification fetches notifications list.
//
// ConoHa API docs: https://www.conoha.jp/docs/account-informations-detail-specified.html
func (c *Conoha) Notification(notificationCode string) (Notification, *ResponseMeta, error) {
	p := getAccountNotificationResponseParam{}

	u := c.AccountServiceURL + "/v1/" + c.TenantID + "/notifications/" + notificationCode

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Notification, meta, err
}
