package transferwebhook

import (
	"encoding/json"
)

// HandletransferwebhookRequest creates a Notification object from the given JSON string
func HandletransferwebhookRequest(req string) (*transferwebhookRequest, error) {
	res := transferwebhookRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
