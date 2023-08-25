package managementwebhook

import (
	"encoding/json"
)

// HandleMerchantCreatedNotificationRequest creates a Notification object from the given JSON string
func HandleMerchantCreatedNotificationRequest(req string) (*MerchantCreatedNotificationRequest, error) {
	res := MerchantCreatedNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// HandleMerchantUpdatedNotificationRequest creates a Notification object from the given JSON string
func HandleMerchantUpdatedNotificationRequest(req string) (*MerchantUpdatedNotificationRequest, error) {
	res := MerchantUpdatedNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// HandlePaymentMethodCreatedNotificationRequest creates a Notification object from the given JSON string
func HandlePaymentMethodCreatedNotificationRequest(req string) (*PaymentMethodCreatedNotificationRequest, error) {
	res := PaymentMethodCreatedNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
