/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
	"fmt"
)

// TransferEventEventsDataInner - struct for TransferEventEventsDataInner
type TransferEventEventsDataInner struct {
	MerchantPurchaseData *MerchantPurchaseData
}

// MerchantPurchaseDataAsTransferEventEventsDataInner is a convenience function that returns MerchantPurchaseData wrapped in TransferEventEventsDataInner
func MerchantPurchaseDataAsTransferEventEventsDataInner(v *MerchantPurchaseData) TransferEventEventsDataInner {
	return TransferEventEventsDataInner{
		MerchantPurchaseData: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TransferEventEventsDataInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MerchantPurchaseData
	err = json.Unmarshal(data, &dst.MerchantPurchaseData)
	if err == nil {
		jsonMerchantPurchaseData, _ := json.Marshal(dst.MerchantPurchaseData)
		if string(jsonMerchantPurchaseData) == "{}" || !dst.MerchantPurchaseData.isValidType() { // empty struct
			dst.MerchantPurchaseData = nil
        } else {
			match++
		}
	} else {
		dst.MerchantPurchaseData = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MerchantPurchaseData = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TransferEventEventsDataInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TransferEventEventsDataInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransferEventEventsDataInner) MarshalJSON() ([]byte, error) {
	if src.MerchantPurchaseData != nil {
		return json.Marshal(&src.MerchantPurchaseData)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TransferEventEventsDataInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MerchantPurchaseData != nil {
		return obj.MerchantPurchaseData
	}

	// all schemas are nil
	return nil
}

type NullableTransferEventEventsDataInner struct {
	value *TransferEventEventsDataInner
	isSet bool
}

func (v NullableTransferEventEventsDataInner) Get() *TransferEventEventsDataInner {
	return v.value
}

func (v *NullableTransferEventEventsDataInner) Set(val *TransferEventEventsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferEventEventsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferEventEventsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferEventEventsDataInner(val *TransferEventEventsDataInner) *NullableTransferEventEventsDataInner {
	return &NullableTransferEventEventsDataInner{value: val, isSet: true}
}

func (v NullableTransferEventEventsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferEventEventsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

