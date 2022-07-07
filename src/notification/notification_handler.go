package notification

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// NotificationService used to namespace this util under the client for consistency and for future prooffing if this ever requires api access
type NotificationService common.Service

//Event codes
const (
	EventCodeAuthorisation            = "AUTHORISATION"
	EventCodeAuthorisationAdjustment  = "AUTHORISATION_ADJUSTMENT"
	EventCodeCancellation             = "CANCELLATION"
	EventCodeCancelOrRefund           = "CANCEL_OR_REFUND"
	EventCodeCapture                  = "CAPTURE"
	EventCodeCaptureFailed            = "CAPTURE_FAILED"
	EventCodeHandledExternally        = "HANDLED_EXTERNALLY"
	EventOrderOpened                  = "ORDER_OPENED"
	EventOrderClosed                  = "ORDER_CLOSED"
	EventCodeRefund                   = "REFUND"
	EventCodeRefundFailed             = "REFUND_FAILED"
	EventCodeRefundedReversed         = "REFUNDED_REVERSED"
	EventCodeRefundWithData           = "REFUND_WITH_DATA"
	EventCodeReportAvailable          = "REPORT_AVAILABLE"
	EventCodeVoidPendingRefund        = "VOID_PENDING_REFUND"
	EventCodeChargeback               = "CHARGEBACK"
	EventCodeChargebackReversed       = "CHARGEBACK_REVERSED"
	EventCodeNotificationOfChargeback = "NOTIFICATION_OF_CHARGEBACK"
	EventCodeNotificationOfFraud      = "NOTIFICATION_OF_FRAUD"
	EventCodePrearbitrationLost       = "PREARBITRATION_LOST"
	EventCodePrearbitrationWon        = "PREARBITRATION_WON"
	EventCodeRequestForInformation    = "REQUEST_FOR_INFORMATION"
	EventCodeSecondChargeback         = "SECOND_CHARGEBACK"
	EventCodePayoutExpire             = "PAYOUT_EXPIRE"
	EventCodePayoutDecline            = "PAYOUT_DECLINE"
	EventCodePayoutThirdparty         = "PAYOUT_THIRDPARTY"
	EventCodePaidoutReversed          = "PAIDOUT_REVERSED"
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
