package transferwebhook

import (
	"encoding/json"
)

// HandleTransferwebhookRequest creates a Notification object from the given JSON string
func HandleTransferwebhookRequest(req string) (*TransferNotificationRequest, error) {
	res := TransferNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
