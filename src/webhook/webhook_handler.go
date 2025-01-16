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
	EventCodeExpire                   = "EXPIRE"
	EventCodeManualReviewAccept       = "MANUAL_REVIEW_ACCEPT"
	EventCodeManualReviewReject       = "MANUAL_REVIEW_REJECT"
	EventCodeHandledExternally        = "HANDLED_EXTERNALLY"
	EventCodeOfferClosed              = "OFFER_CLOSED"
	EventOrderOpened                  = "ORDER_OPENED"
	EventOrderClosed                  = "ORDER_CLOSED"
	EventCodePostponedRefund          = "POSTPONED_REFUND"
	EventCodeRecurringContract        = "RECURRING_CONTRACT"
	EventCodeRefund                   = "REFUND"
	EventCodeRefundFailed             = "REFUND_FAILED"
	EventCodeRefundedReversed         = "REFUNDED_REVERSED"
	EventCodeRefundWithData           = "REFUND_WITH_DATA"
	EventCodeReportAvailable          = "REPORT_AVAILABLE"
	EventCodeTechnicalCancel          = "TECHNICAL_CANCEL"
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
