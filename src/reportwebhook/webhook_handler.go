package reportwebhook

import (
	"encoding/json"
)

// HandleReportNotificationRequest creates a Notification object from the given JSON string
func HandleReportNotificationRequest(req string) (*ReportNotificationRequest, error) {
	res := ReportNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
