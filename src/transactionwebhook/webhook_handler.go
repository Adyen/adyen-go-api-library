package transactionwebhook

import (
	"encoding/json"
)

// HandleTransactionNotificationRequestV4 creates a Notification object from the given JSON string
func HandleTransactionNotificationRequestV4(req string) (*TransactionNotificationRequestV4, error) {
	res := TransactionNotificationRequestV4{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
