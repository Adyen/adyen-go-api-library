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

// checks if the CancelOrderRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelOrderRequest{}

// CancelOrderRequest struct for CancelOrderRequest
type CancelOrderRequest struct {
	// The merchant account identifier that orderData belongs to.
	MerchantAccount string             `json:"merchantAccount"`
	Order           EncryptedOrderData `json:"order"`
}

// NewCancelOrderRequest instantiates a new CancelOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelOrderRequest(merchantAccount string, order EncryptedOrderData) *CancelOrderRequest {
	this := CancelOrderRequest{}
	this.MerchantAccount = merchantAccount
	this.Order = order
	return &this
}

// NewCancelOrderRequestWithDefaults instantiates a new CancelOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelOrderRequestWithDefaults() *CancelOrderRequest {
	this := CancelOrderRequest{}
	return &this
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *CancelOrderRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *CancelOrderRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *CancelOrderRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetOrder returns the Order field value
func (o *CancelOrderRequest) GetOrder() EncryptedOrderData {
	if o == nil {
		var ret EncryptedOrderData
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *CancelOrderRequest) GetOrderOk() (*EncryptedOrderData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *CancelOrderRequest) SetOrder(v EncryptedOrderData) {
	o.Order = v
}

func (o CancelOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["order"] = o.Order
	return toSerialize, nil
}

type NullableCancelOrderRequest struct {
	value *CancelOrderRequest
	isSet bool
}

func (v NullableCancelOrderRequest) Get() *CancelOrderRequest {
	return v.value
}

func (v *NullableCancelOrderRequest) Set(val *CancelOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelOrderRequest(val *CancelOrderRequest) *NullableCancelOrderRequest {
	return &NullableCancelOrderRequest{value: val, isSet: true}
}

func (v NullableCancelOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
