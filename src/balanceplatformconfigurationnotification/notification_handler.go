package balanceplatformconfigurationnotification

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// NotificationService used to namespace this util under the client for consistency and for future prooffing if this ever requires api access
type NotificationService common.Service

// HandleAccountNotificationRequest creates a Notification object from the given JSON string
func (service *NotificationService) HandleAccountHolderNotificationRequest(req string) (*AccountHolderNotificationRequest, error) {
	res := AccountHolderNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// HandleBalanceAccountNotificationRequest creates a Notification object from the given JSON string
func (service *NotificationService) HandleBalanceAccountNotificationRequest(req string) (*BalanceAccountNotificationRequest, error) {
	res := BalanceAccountNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// HandleCardOrderNotificationRequest creates a Notification object from the given JSON string
func (service *NotificationService) HandleCardOrderNotificationRequest(req string) (*CardOrderNotificationRequest, error) {
	res := CardOrderNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// HandlePaymentNotificationRequest creates a Notification object from the given JSON string
func (service *NotificationService) HandlePaymentNotificationRequest(req string) (*PaymentNotificationRequest, error) {
	res := PaymentNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// HandleSweepConfigurationNotificationRequest creates a Notification object from the given JSON string
func (service *NotificationService) HandleSweepConfigurationNotificationRequest(req string) (*SweepConfigurationNotificationRequest, error) {
	res := SweepConfigurationNotificationRequest{}
	err := json.Unmarshal([]byte(req), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
