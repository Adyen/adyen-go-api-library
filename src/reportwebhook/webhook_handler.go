package reportwebhook

import (
	"encoding/json"
)

// HandleReportwebhookRequest creates a Notification object from the given JSON string
func HandleReportwebhookRequest(req string) (*ReportNotificationRequest, error) {
	res := ReportNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
