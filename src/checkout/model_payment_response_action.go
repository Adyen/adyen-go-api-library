/*
Adyen Checkout API

API version: 70
Contact: developer-experience@adyen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"
	"fmt"
)

// PaymentResponseAction - Action to be taken for completing the payment.
type PaymentResponseAction struct {
	CheckoutAwaitAction          *CheckoutAwaitAction
	CheckoutNativeRedirectAction *CheckoutNativeRedirectAction
	CheckoutQrCodeAction         *CheckoutQrCodeAction
	CheckoutRedirectAction       *CheckoutRedirectAction
	CheckoutSDKAction            *CheckoutSDKAction
	CheckoutThreeDS2Action       *CheckoutThreeDS2Action
	CheckoutVoucherAction        *CheckoutVoucherAction
}

// CheckoutAwaitActionAsPaymentResponseAction is a convenience function that returns CheckoutAwaitAction wrapped in PaymentResponseAction
func CheckoutAwaitActionAsPaymentResponseAction(v *CheckoutAwaitAction) PaymentResponseAction {
	return PaymentResponseAction{
		CheckoutAwaitAction: v,
	}
}

// CheckoutNativeRedirectActionAsPaymentResponseAction is a convenience function that returns CheckoutNativeRedirectAction wrapped in PaymentResponseAction
func CheckoutNativeRedirectActionAsPaymentResponseAction(v *CheckoutNativeRedirectAction) PaymentResponseAction {
	return PaymentResponseAction{
		CheckoutNativeRedirectAction: v,
	}
}

// CheckoutQrCodeActionAsPaymentResponseAction is a convenience function that returns CheckoutQrCodeAction wrapped in PaymentResponseAction
func CheckoutQrCodeActionAsPaymentResponseAction(v *CheckoutQrCodeAction) PaymentResponseAction {
	return PaymentResponseAction{
		CheckoutQrCodeAction: v,
	}
}

// CheckoutRedirectActionAsPaymentResponseAction is a convenience function that returns CheckoutRedirectAction wrapped in PaymentResponseAction
func CheckoutRedirectActionAsPaymentResponseAction(v *CheckoutRedirectAction) PaymentResponseAction {
	return PaymentResponseAction{
		CheckoutRedirectAction: v,
	}
}

// CheckoutSDKActionAsPaymentResponseAction is a convenience function that returns CheckoutSDKAction wrapped in PaymentResponseAction
func CheckoutSDKActionAsPaymentResponseAction(v *CheckoutSDKAction) PaymentResponseAction {
	return PaymentResponseAction{
		CheckoutSDKAction: v,
	}
}

// CheckoutThreeDS2ActionAsPaymentResponseAction is a convenience function that returns CheckoutThreeDS2Action wrapped in PaymentResponseAction
func CheckoutThreeDS2ActionAsPaymentResponseAction(v *CheckoutThreeDS2Action) PaymentResponseAction {
	return PaymentResponseAction{
		CheckoutThreeDS2Action: v,
	}
}

// CheckoutVoucherActionAsPaymentResponseAction is a convenience function that returns CheckoutVoucherAction wrapped in PaymentResponseAction
func CheckoutVoucherActionAsPaymentResponseAction(v *CheckoutVoucherAction) PaymentResponseAction {
	return PaymentResponseAction{
		CheckoutVoucherAction: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PaymentResponseAction) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CheckoutAwaitAction
	err = json.Unmarshal(data, &dst.CheckoutAwaitAction)
	if err == nil {
		jsonCheckoutAwaitAction, _ := json.Marshal(dst.CheckoutAwaitAction)
		if string(jsonCheckoutAwaitAction) == "{}" { // empty struct
			dst.CheckoutAwaitAction = nil
		} else if dst.CheckoutAwaitAction.isValidType() {
			match++
		}
	} else {
		dst.CheckoutAwaitAction = nil
	}

	// try to unmarshal data into CheckoutNativeRedirectAction
	err = json.Unmarshal(data, &dst.CheckoutNativeRedirectAction)
	if err == nil {
		jsonCheckoutNativeRedirectAction, _ := json.Marshal(dst.CheckoutNativeRedirectAction)
		if string(jsonCheckoutNativeRedirectAction) == "{}" { // empty struct
			dst.CheckoutNativeRedirectAction = nil
		} else if dst.CheckoutNativeRedirectAction.isValidType() {
			match++
		}
	} else {
		dst.CheckoutNativeRedirectAction = nil
	}

	// try to unmarshal data into CheckoutQrCodeAction
	err = json.Unmarshal(data, &dst.CheckoutQrCodeAction)
	if err == nil {
		jsonCheckoutQrCodeAction, _ := json.Marshal(dst.CheckoutQrCodeAction)
		if string(jsonCheckoutQrCodeAction) == "{}" { // empty struct
			dst.CheckoutQrCodeAction = nil
		} else if dst.CheckoutQrCodeAction.isValidType() {
			match++
		}
	} else {
		dst.CheckoutQrCodeAction = nil
	}

	// try to unmarshal data into CheckoutRedirectAction
	err = json.Unmarshal(data, &dst.CheckoutRedirectAction)
	if err == nil {
		jsonCheckoutRedirectAction, _ := json.Marshal(dst.CheckoutRedirectAction)
		if string(jsonCheckoutRedirectAction) == "{}" { // empty struct
			dst.CheckoutRedirectAction = nil
		} else if dst.CheckoutRedirectAction.isValidType() {
			match++
		}
	} else {
		dst.CheckoutRedirectAction = nil
	}

	// try to unmarshal data into CheckoutSDKAction
	err = json.Unmarshal(data, &dst.CheckoutSDKAction)
	if err == nil {
		jsonCheckoutSDKAction, _ := json.Marshal(dst.CheckoutSDKAction)
		if string(jsonCheckoutSDKAction) == "{}" { // empty struct
			dst.CheckoutSDKAction = nil
		} else if dst.CheckoutSDKAction.isValidType() {
			match++
		}
	} else {
		dst.CheckoutSDKAction = nil
	}

	// try to unmarshal data into CheckoutThreeDS2Action
	err = json.Unmarshal(data, &dst.CheckoutThreeDS2Action)
	if err == nil {
		jsonCheckoutThreeDS2Action, _ := json.Marshal(dst.CheckoutThreeDS2Action)
		if string(jsonCheckoutThreeDS2Action) == "{}" { // empty struct
			dst.CheckoutThreeDS2Action = nil
		} else if dst.CheckoutThreeDS2Action.isValidType() {
			match++
		}
	} else {
		dst.CheckoutThreeDS2Action = nil
	}

	// try to unmarshal data into CheckoutVoucherAction
	err = json.Unmarshal(data, &dst.CheckoutVoucherAction)
	if err == nil {
		jsonCheckoutVoucherAction, _ := json.Marshal(dst.CheckoutVoucherAction)
		if string(jsonCheckoutVoucherAction) == "{}" { // empty struct
			dst.CheckoutVoucherAction = nil
		} else if dst.CheckoutVoucherAction.isValidType() {
			match++
		}
	} else {
		dst.CheckoutVoucherAction = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CheckoutAwaitAction = nil
		dst.CheckoutNativeRedirectAction = nil
		dst.CheckoutQrCodeAction = nil
		dst.CheckoutRedirectAction = nil
		dst.CheckoutSDKAction = nil
		dst.CheckoutThreeDS2Action = nil
		dst.CheckoutVoucherAction = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PaymentResponseAction)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PaymentResponseAction)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PaymentResponseAction) MarshalJSON() ([]byte, error) {
	if src.CheckoutAwaitAction != nil {
		return json.Marshal(&src.CheckoutAwaitAction)
	}

	if src.CheckoutNativeRedirectAction != nil {
		return json.Marshal(&src.CheckoutNativeRedirectAction)
	}

	if src.CheckoutQrCodeAction != nil {
		return json.Marshal(&src.CheckoutQrCodeAction)
	}

	if src.CheckoutRedirectAction != nil {
		return json.Marshal(&src.CheckoutRedirectAction)
	}

	if src.CheckoutSDKAction != nil {
		return json.Marshal(&src.CheckoutSDKAction)
	}

	if src.CheckoutThreeDS2Action != nil {
		return json.Marshal(&src.CheckoutThreeDS2Action)
	}

	if src.CheckoutVoucherAction != nil {
		return json.Marshal(&src.CheckoutVoucherAction)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PaymentResponseAction) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CheckoutAwaitAction != nil {
		return obj.CheckoutAwaitAction
	}

	if obj.CheckoutNativeRedirectAction != nil {
		return obj.CheckoutNativeRedirectAction
	}

	if obj.CheckoutQrCodeAction != nil {
		return obj.CheckoutQrCodeAction
	}

	if obj.CheckoutRedirectAction != nil {
		return obj.CheckoutRedirectAction
	}

	if obj.CheckoutSDKAction != nil {
		return obj.CheckoutSDKAction
	}

	if obj.CheckoutThreeDS2Action != nil {
		return obj.CheckoutThreeDS2Action
	}

	if obj.CheckoutVoucherAction != nil {
		return obj.CheckoutVoucherAction
	}

	// all schemas are nil
	return nil
}

type NullablePaymentResponseAction struct {
	value *PaymentResponseAction
	isSet bool
}

func (v NullablePaymentResponseAction) Get() *PaymentResponseAction {
	return v.value
}

func (v *NullablePaymentResponseAction) Set(val *PaymentResponseAction) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentResponseAction) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentResponseAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentResponseAction(val *PaymentResponseAction) *NullablePaymentResponseAction {
	return &NullablePaymentResponseAction{value: val, isSet: true}
}

func (v NullablePaymentResponseAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentResponseAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}