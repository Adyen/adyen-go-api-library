/*
Transfer webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the BankAccountV3 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BankAccountV3{}

// BankAccountV3 struct for BankAccountV3
type BankAccountV3 struct {
	AccountHolder         PartyIdentification                `json:"accountHolder"`
	AccountIdentification BankAccountV3AccountIdentification `json:"accountIdentification"`
}

// NewBankAccountV3 instantiates a new BankAccountV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankAccountV3(accountHolder PartyIdentification, accountIdentification BankAccountV3AccountIdentification) *BankAccountV3 {
	this := BankAccountV3{}
	this.AccountHolder = accountHolder
	this.AccountIdentification = accountIdentification
	return &this
}

// NewBankAccountV3WithDefaults instantiates a new BankAccountV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankAccountV3WithDefaults() *BankAccountV3 {
	this := BankAccountV3{}
	return &this
}

// GetAccountHolder returns the AccountHolder field value
func (o *BankAccountV3) GetAccountHolder() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.AccountHolder
}

// GetAccountHolderOk returns a tuple with the AccountHolder field value
// and a boolean to check if the value has been set.
func (o *BankAccountV3) GetAccountHolderOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountHolder, true
}

// SetAccountHolder sets field value
func (o *BankAccountV3) SetAccountHolder(v PartyIdentification) {
	o.AccountHolder = v
}

// GetAccountIdentification returns the AccountIdentification field value
func (o *BankAccountV3) GetAccountIdentification() BankAccountV3AccountIdentification {
	if o == nil {
		var ret BankAccountV3AccountIdentification
		return ret
	}

	return o.AccountIdentification
}

// GetAccountIdentificationOk returns a tuple with the AccountIdentification field value
// and a boolean to check if the value has been set.
func (o *BankAccountV3) GetAccountIdentificationOk() (*BankAccountV3AccountIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountIdentification, true
}

// SetAccountIdentification sets field value
func (o *BankAccountV3) SetAccountIdentification(v BankAccountV3AccountIdentification) {
	o.AccountIdentification = v
}

func (o BankAccountV3) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankAccountV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountHolder"] = o.AccountHolder
	toSerialize["accountIdentification"] = o.AccountIdentification
	return toSerialize, nil
}

type NullableBankAccountV3 struct {
	value *BankAccountV3
	isSet bool
}

func (v NullableBankAccountV3) Get() *BankAccountV3 {
	return v.value
}

func (v *NullableBankAccountV3) Set(val *BankAccountV3) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccountV3) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccountV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccountV3(val *BankAccountV3) *NullableBankAccountV3 {
	return &NullableBankAccountV3{value: val, isSet: true}
}

func (v NullableBankAccountV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankAccountV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
