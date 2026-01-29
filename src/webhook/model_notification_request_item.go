package webhook

import "time"

type NotificationRequestItem struct {
	AdditionalData      *map[string]interface{} `json:"additionalData,omitempty"`
	Amount              Amount                  `json:"amount"`
	EventCode           string                  `json:"eventCode"`
	EventDate           *time.Time              `json:"eventDate"`
	MerchantAccountCode string                  `json:"merchantAccountCode"`
	MerchantReference   string                  `json:"merchantReference"`
	Operations          []string                `json:"operations,omitempty"`
	OriginalReference   string                  `json:"originalReference,omitempty"`
	PaymentMethod       string                  `json:"paymentMethod"`
	PspReference        string                  `json:"pspReference"`
	Reason              string                  `json:"reason"`
	Success             string                  `json:"success"`
}
