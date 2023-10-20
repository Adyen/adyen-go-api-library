package acswebhook

import (
	"encoding/json"
)

// HandleAuthenticationNotificationRequest creates a Notification object from the given JSON string
func HandleAuthenticationNotificationRequest(req string) (*AuthenticationNotificationRequest, error) {
	res := AuthenticationNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
