package webhook

import "time"

type NotificationRequestItem struct {
	AdditionalData      *map[string]interface{} `json:"additionalData,omitEmpty"`
	Amount              Amount                  `json:"amount"`
	EventCode           string                  `json:"eventCode"`
	EventDate           *time.Time              `json:"eventDate"`
	MerchantAccountCode string                  `json:"merchantAccountCode"`
	MerchantReference   string                  `json:"merchantReference"`
	Operations          []string                `json:"operations,omitEmpty"`
	OriginalReference   string                  `json:"originalReference,omitEmpty"`
	PaymentMethod       string                  `json:"paymentMethod"`
	PspReference        string                  `json:"pspReference"`
	Reason              string                  `json:"reason"`
	Success             string                  `json:"success"`
}
