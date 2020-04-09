package notification

import "time"

type NotificationRequestItem struct {
	AdditionalData      interface{} `json:"additionalData"`
	Amount              Amount      `json:"amount"`
	EventCode           string      `json:"eventCode"`
	EventDate           *time.Time  `json:"eventDate"`
	MerchantAccountCode string      `json:"merchantAccountCode"`
	Operations          []string    `json:"operations,omitEmpty"`
	OriginalReference   string      `json:"originalReference,omitEmpty"`
	PaymentMethod       string      `json:"paymentMethod"`
	PspReference        string      `json:"pspReference"`
	Reason              string      `json:"reason"`
	Success             string      `json:"success"`
}
