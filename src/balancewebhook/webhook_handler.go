package balancewebhook

import (
	"encoding/json"
)

// HandleBalanceAccountBalanceNotificationRequest creates a Notification object from the given JSON string
func HandleBalanceAccountBalanceNotificationRequest(req string) (*BalanceAccountBalanceNotificationRequest, error) {
	res := BalanceAccountBalanceNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
