package configurationwebhook

import (
	"encoding/json"
)

// HandleAccountNotificationRequest creates a Notification object from the given JSON string
func HandleAccountHolderNotificationRequest(req string) (*AccountHolderNotificationRequest, error) {
	res := AccountHolderNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// HandleBalanceAccountNotificationRequest creates a Notification object from the given JSON string
func HandleBalanceAccountNotificationRequest(req string) (*BalanceAccountNotificationRequest, error) {
	res := BalanceAccountNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// HandleCardOrderNotificationRequest creates a Notification object from the given JSON string
func HandleCardOrderNotificationRequest(req string) (*CardOrderNotificationRequest, error) {
	res := CardOrderNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// HandlePaymentNotificationRequest creates a Notification object from the given JSON string
func HandlePaymentNotificationRequest(req string) (*PaymentNotificationRequest, error) {
	res := PaymentNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// HandleSweepconfigurationNotificationRequest creates a Notification object from the given JSON string
func HandleSweepConfigurationNotificationRequest(req string) (*SweepConfigurationNotificationRequest, error) {
	res := SweepConfigurationNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
