package balanceplatformtransfernotification

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// NotificationService used to namespace this util under the client for consistency and for future prooffing if this ever requires api access
type NotificationService common.Service

// HandleTransferNotificationRequest creates a Notification object from the given JSON string
func (service *NotificationService) HandleTransferNotificationRequest(req string) (*TransferNotificationRequest, error) {
	res := TransferNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
