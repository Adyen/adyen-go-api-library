package reportwebhook

import (
	"encoding/json"
)

// HandlereportwebhookRequest creates a Notification object from the given JSON string
func HandlereportwebhookRequest(req string) (*reportwebhookRequest, error) {
	res := reportwebhookRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
