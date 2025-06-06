/*
Adyen Recurring API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package recurring

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the DisablePermitRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DisablePermitRequest{}

// DisablePermitRequest struct for DisablePermitRequest
type DisablePermitRequest struct {
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	// The permit token to disable.
	Token string `json:"token"`
}

// NewDisablePermitRequest instantiates a new DisablePermitRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisablePermitRequest(merchantAccount string, token string) *DisablePermitRequest {
	this := DisablePermitRequest{}
	this.MerchantAccount = merchantAccount
	this.Token = token
	return &this
}

// NewDisablePermitRequestWithDefaults instantiates a new DisablePermitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisablePermitRequestWithDefaults() *DisablePermitRequest {
	this := DisablePermitRequest{}
	return &this
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *DisablePermitRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *DisablePermitRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *DisablePermitRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetToken returns the Token field value
func (o *DisablePermitRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DisablePermitRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DisablePermitRequest) SetToken(v string) {
	o.Token = v
}

func (o DisablePermitRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisablePermitRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableDisablePermitRequest struct {
	value *DisablePermitRequest
	isSet bool
}

func (v NullableDisablePermitRequest) Get() *DisablePermitRequest {
	return v.value
}

func (v *NullableDisablePermitRequest) Set(val *DisablePermitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDisablePermitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDisablePermitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisablePermitRequest(val *DisablePermitRequest) *NullableDisablePermitRequest {
	return &NullableDisablePermitRequest{value: val, isSet: true}
}

func (v NullableDisablePermitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisablePermitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
