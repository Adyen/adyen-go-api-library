package notification

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v3/src/common"
)

// NotificationService used to namespace this util under the client for consistency and for future prooffing if this ever requires api access
type NotificationService common.Service

//Event codes
const (
	EventCodeAuthorisation    = "AUTHORISATION"
	EventCodeCancellation     = "CANCELLATION"
	EventCodeRefund           = "REFUND"
	EventCodeCancelOrRefund   = "CANCEL_OR_REFUND"
	EventCodeCapture          = "CAPTURE"
	EventCodeCaptureFailed    = "CAPTURE_FAILED"
	EventCodeRefundFailed     = "REFUND_FAILED"
	EventCodeRefundedReversed = "REFUNDED_REVERSED"
	EventCodePaidoutReversed  = "PAIDOUT_REVERSED"
)

//Additional Data
const (
	AdditionalDataTotalFraudScore   = "totalFraudScore"
	AdditionalDataFraudCheckPattern = "fraudCheck-(\\d+)-([A-Za-z0-9]+)"
)

// HandleNotificationRequest creates a Notification object from the given JSON string
func (service *NotificationService) HandleNotificationRequest(req string) (*Notification, error) {
	res := Notification{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
