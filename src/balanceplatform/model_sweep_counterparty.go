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

// checks if the SweepCounterparty type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SweepCounterparty{}

// SweepCounterparty struct for SweepCounterparty
type SweepCounterparty struct {
	// The unique identifier of the destination or source [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id).   > If you are updating the counterparty from a transfer instrument to a balance account, set `transferInstrumentId` to **null**.
	BalanceAccountId *string `json:"balanceAccountId,omitempty"`
	// The merchant account that will be the source of funds.  You can only use this parameter with sweeps of `type` **pull** and if you are processing payments with Adyen.
	MerchantAccount *string `json:"merchantAccount,omitempty"`
	// The unique identifier of the destination or source [transfer instrument](https://docs.adyen.com/api-explorer/legalentity/latest/post/transferInstruments#responses-200-id) depending on the sweep `type`  . To set up automated top-up sweeps to balance accounts in your [marketplace](https://docs.adyen.com/marketplaces/top-up-balance-account/#before-you-begin) or [platform](https://docs.adyen.com/platforms/top-up-balance-account/#before-you-begin), use this parameter in combination with a `merchantAccount` and a sweep `type` of **pull**.  Top-up sweeps start a direct debit request from the source transfer instrument. Contact Adyen Support to enable this feature.> If you are updating the counterparty from a balance account to a transfer instrument, set `balanceAccountId` to **null**.
	TransferInstrumentId *string `json:"transferInstrumentId,omitempty"`
}

// NewSweepCounterparty instantiates a new SweepCounterparty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSweepCounterparty() *SweepCounterparty {
	this := SweepCounterparty{}
	return &this
}

// NewSweepCounterpartyWithDefaults instantiates a new SweepCounterparty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSweepCounterpartyWithDefaults() *SweepCounterparty {
	this := SweepCounterparty{}
	return &this
}

// GetBalanceAccountId returns the BalanceAccountId field value if set, zero value otherwise.
func (o *SweepCounterparty) GetBalanceAccountId() string {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		var ret string
		return ret
	}
	return *o.BalanceAccountId
}

// GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SweepCounterparty) GetBalanceAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		return nil, false
	}
	return o.BalanceAccountId, true
}

// HasBalanceAccountId returns a boolean if a field has been set.
func (o *SweepCounterparty) HasBalanceAccountId() bool {
	if o != nil && !common.IsNil(o.BalanceAccountId) {
		return true
	}

	return false
}

// SetBalanceAccountId gets a reference to the given string and assigns it to the BalanceAccountId field.
func (o *SweepCounterparty) SetBalanceAccountId(v string) {
	o.BalanceAccountId = &v
}

// GetMerchantAccount returns the MerchantAccount field value if set, zero value otherwise.
func (o *SweepCounterparty) GetMerchantAccount() string {
	if o == nil || common.IsNil(o.MerchantAccount) {
		var ret string
		return ret
	}
	return *o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SweepCounterparty) GetMerchantAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantAccount) {
		return nil, false
	}
	return o.MerchantAccount, true
}

// HasMerchantAccount returns a boolean if a field has been set.
func (o *SweepCounterparty) HasMerchantAccount() bool {
	if o != nil && !common.IsNil(o.MerchantAccount) {
		return true
	}

	return false
}

// SetMerchantAccount gets a reference to the given string and assigns it to the MerchantAccount field.
func (o *SweepCounterparty) SetMerchantAccount(v string) {
	o.MerchantAccount = &v
}

// GetTransferInstrumentId returns the TransferInstrumentId field value if set, zero value otherwise.
func (o *SweepCounterparty) GetTransferInstrumentId() string {
	if o == nil || common.IsNil(o.TransferInstrumentId) {
		var ret string
		return ret
	}
	return *o.TransferInstrumentId
}

// GetTransferInstrumentIdOk returns a tuple with the TransferInstrumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SweepCounterparty) GetTransferInstrumentIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransferInstrumentId) {
		return nil, false
	}
	return o.TransferInstrumentId, true
}

// HasTransferInstrumentId returns a boolean if a field has been set.
func (o *SweepCounterparty) HasTransferInstrumentId() bool {
	if o != nil && !common.IsNil(o.TransferInstrumentId) {
		return true
	}

	return false
}

// SetTransferInstrumentId gets a reference to the given string and assigns it to the TransferInstrumentId field.
func (o *SweepCounterparty) SetTransferInstrumentId(v string) {
	o.TransferInstrumentId = &v
}

func (o SweepCounterparty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SweepCounterparty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BalanceAccountId) {
		toSerialize["balanceAccountId"] = o.BalanceAccountId
	}
	if !common.IsNil(o.MerchantAccount) {
		toSerialize["merchantAccount"] = o.MerchantAccount
	}
	if !common.IsNil(o.TransferInstrumentId) {
		toSerialize["transferInstrumentId"] = o.TransferInstrumentId
	}
	return toSerialize, nil
}

type NullableSweepCounterparty struct {
	value *SweepCounterparty
	isSet bool
}

func (v NullableSweepCounterparty) Get() *SweepCounterparty {
	return v.value
}

func (v *NullableSweepCounterparty) Set(val *SweepCounterparty) {
	v.value = val
	v.isSet = true
}

func (v NullableSweepCounterparty) IsSet() bool {
	return v.isSet
}

func (v *NullableSweepCounterparty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSweepCounterparty(val *SweepCounterparty) *NullableSweepCounterparty {
	return &NullableSweepCounterparty{value: val, isSet: true}
}

func (v NullableSweepCounterparty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSweepCounterparty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
