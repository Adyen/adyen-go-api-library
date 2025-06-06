/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the NZLocalAccountIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NZLocalAccountIdentification{}

// NZLocalAccountIdentification struct for NZLocalAccountIdentification
type NZLocalAccountIdentification struct {
	// The 15-16 digit bank account number. The first 2 digits are the bank number, the next 4 digits are the branch number, the next 7 digits are the account number, and the final 2-3 digits are the suffix.
	AccountNumber string `json:"accountNumber"`
	// **nzLocal**
	Type string `json:"type"`
}

// NewNZLocalAccountIdentification instantiates a new NZLocalAccountIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNZLocalAccountIdentification(accountNumber string, type_ string) *NZLocalAccountIdentification {
	this := NZLocalAccountIdentification{}
	this.AccountNumber = accountNumber
	this.Type = type_
	return &this
}

// NewNZLocalAccountIdentificationWithDefaults instantiates a new NZLocalAccountIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNZLocalAccountIdentificationWithDefaults() *NZLocalAccountIdentification {
	this := NZLocalAccountIdentification{}
	var type_ string = "nzLocal"
	this.Type = type_
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *NZLocalAccountIdentification) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *NZLocalAccountIdentification) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *NZLocalAccountIdentification) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetType returns the Type field value
func (o *NZLocalAccountIdentification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NZLocalAccountIdentification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NZLocalAccountIdentification) SetType(v string) {
	o.Type = v
}

func (o NZLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NZLocalAccountIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountNumber"] = o.AccountNumber
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableNZLocalAccountIdentification struct {
	value *NZLocalAccountIdentification
	isSet bool
}

func (v NullableNZLocalAccountIdentification) Get() *NZLocalAccountIdentification {
	return v.value
}

func (v *NullableNZLocalAccountIdentification) Set(val *NZLocalAccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableNZLocalAccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableNZLocalAccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNZLocalAccountIdentification(val *NZLocalAccountIdentification) *NullableNZLocalAccountIdentification {
	return &NullableNZLocalAccountIdentification{value: val, isSet: true}
}

func (v NullableNZLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNZLocalAccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *NZLocalAccountIdentification) isValidType() bool {
	var allowedEnumValues = []string{"nzLocal"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
