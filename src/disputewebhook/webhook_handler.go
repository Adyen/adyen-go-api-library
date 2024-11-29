package disputewebhook

import (
	"encoding/json"
)

// HandleDisputeNotificationRequest creates a Notification object from the given JSON string
func HandleDisputeNotificationRequest(req string) (*DisputeNotificationRequest, error) {
	res := DisputeNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
