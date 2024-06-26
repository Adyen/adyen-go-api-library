/*
Configuration webhooks

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"
	"fmt"
)

// PaymentInstrumentAdditionalBankAccountIdentificationsInner - struct for PaymentInstrumentAdditionalBankAccountIdentificationsInner
type PaymentInstrumentAdditionalBankAccountIdentificationsInner struct {
	IbanAccountIdentification *IbanAccountIdentification
}

// IbanAccountIdentificationAsPaymentInstrumentAdditionalBankAccountIdentificationsInner is a convenience function that returns IbanAccountIdentification wrapped in PaymentInstrumentAdditionalBankAccountIdentificationsInner
func IbanAccountIdentificationAsPaymentInstrumentAdditionalBankAccountIdentificationsInner(v *IbanAccountIdentification) PaymentInstrumentAdditionalBankAccountIdentificationsInner {
	return PaymentInstrumentAdditionalBankAccountIdentificationsInner{
		IbanAccountIdentification: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PaymentInstrumentAdditionalBankAccountIdentificationsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IbanAccountIdentification
	err = json.Unmarshal(data, &dst.IbanAccountIdentification)
	if err == nil {
		jsonIbanAccountIdentification, _ := json.Marshal(dst.IbanAccountIdentification)
		if string(jsonIbanAccountIdentification) == "{}" || !dst.IbanAccountIdentification.isValidType() { // empty struct
			dst.IbanAccountIdentification = nil
		} else {
			match++
		}
	} else {
		dst.IbanAccountIdentification = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.IbanAccountIdentification = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PaymentInstrumentAdditionalBankAccountIdentificationsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PaymentInstrumentAdditionalBankAccountIdentificationsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PaymentInstrumentAdditionalBankAccountIdentificationsInner) MarshalJSON() ([]byte, error) {
	if src.IbanAccountIdentification != nil {
		return json.Marshal(&src.IbanAccountIdentification)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PaymentInstrumentAdditionalBankAccountIdentificationsInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IbanAccountIdentification != nil {
		return obj.IbanAccountIdentification
	}

	// all schemas are nil
	return nil
}

type NullablePaymentInstrumentAdditionalBankAccountIdentificationsInner struct {
	value *PaymentInstrumentAdditionalBankAccountIdentificationsInner
	isSet bool
}

func (v NullablePaymentInstrumentAdditionalBankAccountIdentificationsInner) Get() *PaymentInstrumentAdditionalBankAccountIdentificationsInner {
	return v.value
}

func (v *NullablePaymentInstrumentAdditionalBankAccountIdentificationsInner) Set(val *PaymentInstrumentAdditionalBankAccountIdentificationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInstrumentAdditionalBankAccountIdentificationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInstrumentAdditionalBankAccountIdentificationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInstrumentAdditionalBankAccountIdentificationsInner(val *PaymentInstrumentAdditionalBankAccountIdentificationsInner) *NullablePaymentInstrumentAdditionalBankAccountIdentificationsInner {
	return &NullablePaymentInstrumentAdditionalBankAccountIdentificationsInner{value: val, isSet: true}
}

func (v NullablePaymentInstrumentAdditionalBankAccountIdentificationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInstrumentAdditionalBankAccountIdentificationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
