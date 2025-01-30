package webhook

import (
	"encoding/json"
)

// Event codes
const (
	EventCodeACHNotificationOfChange  = "ACH_NOTIFICATION_OF_CHANGE"
	EventCodeAuthorisation            = "AUTHORISATION"
	EventCodeAuthorisationAdjustment  = "AUTHORISATION_ADJUSTMENT"
	EventCodeAutorescue               = "AUTORESCUE"
	EventCodeAutorescueNextAttempt    = "AUTORESCUE_NEXT_ATTEMPT"
	EventCodeCancellation             = "CANCELLATION"
	EventCodeCancelAutorescue         = "CANCEL_AUTORESCUE"
	EventCodeCancelOrRefund           = "CANCEL_OR_REFUND"
	EventCodeCapture                  = "CAPTURE"
	EventCodeCaptureFailed            = "CAPTURE_FAILED"
	EventCodeChargeback               = "CHARGEBACK"
	EventCodeChargebackReversed       = "CHARGEBACK_REVERSED"
	EventCodeExpire                   = "EXPIRE"
	EventCodeHandledExternally        = "HANDLED_EXTERNALLY"
	EventCodeManualReviewAccept       = "MANUAL_REVIEW_ACCEPT"
	EventCodeManualReviewReject       = "MANUAL_REVIEW_REJECT"
	EventCodeNotificationOfChargeback = "NOTIFICATION_OF_CHARGEBACK"
	EventCodeNotificationOfFraud      = "NOTIFICATION_OF_FRAUD"
	EventCodeOfferClosed              = "OFFER_CLOSED"
	EventCodePaidoutReversed          = "PAIDOUT_REVERSED"
	EventCodePayoutDecline            = "PAYOUT_DECLINE"
	EventCodePayoutExpire             = "PAYOUT_EXPIRE"
	EventCodePayoutThirdparty         = "PAYOUT_THIRDPARTY"
	EventCodePostponedRefund          = "POSTPONED_REFUND"
	EventCodePrearbitrationLost       = "PREARBITRATION_LOST"
	EventCodePrearbitrationWon        = "PREARBITRATION_WON"
	EventCodeRecurringContract        = "RECURRING_CONTRACT"
	EventCodeRefund                   = "REFUND"
	EventCodeRefundFailed             = "REFUND_FAILED"
	EventCodeRefundWithData           = "REFUND_WITH_DATA"
	EventCodeRefundedReversed         = "REFUNDED_REVERSED"
	EventCodeReportAvailable          = "REPORT_AVAILABLE"
	EventCodeRequestForInformation    = "REQUEST_FOR_INFORMATION"
	EventCodeSecondChargeback         = "SECOND_CHARGEBACK"
	EventCodeTechnicalCancel          = "TECHNICAL_CANCEL"
	EventCodeVoidPendingRefund        = "VOID_PENDING_REFUND"
	EventOrderClosed                  = "ORDER_CLOSED"
	EventOrderOpened                  = "ORDER_OPENED"
)

// Additional Data
const (
	AdditionalDataTotalFraudScore   = "totalFraudScore"
	AdditionalDataFraudCheckPattern = "fraudCheck-(\\d+)-([A-Za-z0-9]+)"
)

// HandleRequest creates a Notification object from the given JSON string
func HandleRequest(req string) (*Webhook, error) {
	res := Webhook{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
