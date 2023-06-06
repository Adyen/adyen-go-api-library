package reportwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// NotificationService used to namespace this util under the client for consistency and for future prooffing if this ever requires api access
type NotificationService common.Service

// HandlereportwebhookRequest creates a Notification object from the given JSON string
func HandlereportwebhookRequest(req string) (*reportwebhookRequest, error) {
	res := reportwebhookRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
