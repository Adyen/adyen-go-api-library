package transferwebhook

import (
	"encoding/json"
)

// HandleTransferNotificationRequest creates a Notification object from the given JSON string
func HandleTransferNotificationRequest(req string) (*TransferNotificationRequest, error) {
	res := TransferNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
