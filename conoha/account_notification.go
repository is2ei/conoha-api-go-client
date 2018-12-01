package conoha

import (
	"encoding/json"
	"fmt"
)

type getAccountNotificationResponseParam struct {
	Notification *Notification `json:"notification"`
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
func (c *Conoha) Notification(notificationCode string) (*Notification, *ResponseMeta, error) {

	apiEndPoint := fmt.Sprintf("%s/v1/%s/notifications/%s", c.AccountServiceURL, c.TenantID, notificationCode)

	p := getAccountNotificationResponseParam{}

	contents, meta, err := c.buildAndExecRequest("GET", apiEndPoint, nil)
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return p.Notification, meta, err
}
