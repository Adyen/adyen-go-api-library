/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the EncryptedOrderData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EncryptedOrderData{}

// EncryptedOrderData struct for EncryptedOrderData
type EncryptedOrderData struct {
	// The encrypted order data.
	OrderData string `json:"orderData"`
	// The `pspReference` that belongs to the order.
	PspReference string `json:"pspReference"`
}

// NewEncryptedOrderData instantiates a new EncryptedOrderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptedOrderData(orderData string, pspReference string) *EncryptedOrderData {
	this := EncryptedOrderData{}
	this.OrderData = orderData
	this.PspReference = pspReference
	return &this
}

// NewEncryptedOrderDataWithDefaults instantiates a new EncryptedOrderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptedOrderDataWithDefaults() *EncryptedOrderData {
	this := EncryptedOrderData{}
	return &this
}

// GetOrderData returns the OrderData field value
func (o *EncryptedOrderData) GetOrderData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderData
}

// GetOrderDataOk returns a tuple with the OrderData field value
// and a boolean to check if the value has been set.
func (o *EncryptedOrderData) GetOrderDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderData, true
}

// SetOrderData sets field value
func (o *EncryptedOrderData) SetOrderData(v string) {
	o.OrderData = v
}

// GetPspReference returns the PspReference field value
func (o *EncryptedOrderData) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *EncryptedOrderData) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *EncryptedOrderData) SetPspReference(v string) {
	o.PspReference = v
}

func (o EncryptedOrderData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EncryptedOrderData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orderData"] = o.OrderData
	toSerialize["pspReference"] = o.PspReference
	return toSerialize, nil
}

type NullableEncryptedOrderData struct {
	value *EncryptedOrderData
	isSet bool
}

func (v NullableEncryptedOrderData) Get() *EncryptedOrderData {
	return v.value
}

func (v *NullableEncryptedOrderData) Set(val *EncryptedOrderData) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptedOrderData) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptedOrderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptedOrderData(val *EncryptedOrderData) *NullableEncryptedOrderData {
	return &NullableEncryptedOrderData{value: val, isSet: true}
}

func (v NullableEncryptedOrderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptedOrderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
