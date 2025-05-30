/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the PayoutSettingsRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PayoutSettingsRequest{}

// PayoutSettingsRequest struct for PayoutSettingsRequest
type PayoutSettingsRequest struct {
	// Indicates if payouts to this bank account are enabled. Default: **true**.  To receive payouts into this bank account, both `enabled` and `allowed` must be **true**.
	Enabled *bool `json:"enabled,omitempty"`
	// The date when Adyen starts paying out to this bank account.  Format: [ISO 8601](https://www.w3.org/TR/NOTE-datetime), for example, **2019-11-23T12:25:28Z** or **2020-05-27T20:25:28+08:00**.  If not specified, the `enabled` field indicates if payouts are enabled for this bank account.  If a date is specified and:  * `enabled`: **true**, payouts are enabled starting the specified date. * `enabled`: **false**, payouts are disabled until the specified date. On the specified date, `enabled` changes to **true** and this field is reset to **null**.
	EnabledFromDate *string `json:"enabledFromDate,omitempty"`
	// The unique identifier of the [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments) that contains the details of the bank account.
	TransferInstrumentId string `json:"transferInstrumentId"`
}

// NewPayoutSettingsRequest instantiates a new PayoutSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayoutSettingsRequest(transferInstrumentId string) *PayoutSettingsRequest {
	this := PayoutSettingsRequest{}
	this.TransferInstrumentId = transferInstrumentId
	return &this
}

// NewPayoutSettingsRequestWithDefaults instantiates a new PayoutSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayoutSettingsRequestWithDefaults() *PayoutSettingsRequest {
	this := PayoutSettingsRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PayoutSettingsRequest) GetEnabled() bool {
	if o == nil || common.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSettingsRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PayoutSettingsRequest) HasEnabled() bool {
	if o != nil && !common.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PayoutSettingsRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEnabledFromDate returns the EnabledFromDate field value if set, zero value otherwise.
func (o *PayoutSettingsRequest) GetEnabledFromDate() string {
	if o == nil || common.IsNil(o.EnabledFromDate) {
		var ret string
		return ret
	}
	return *o.EnabledFromDate
}

// GetEnabledFromDateOk returns a tuple with the EnabledFromDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSettingsRequest) GetEnabledFromDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnabledFromDate) {
		return nil, false
	}
	return o.EnabledFromDate, true
}

// HasEnabledFromDate returns a boolean if a field has been set.
func (o *PayoutSettingsRequest) HasEnabledFromDate() bool {
	if o != nil && !common.IsNil(o.EnabledFromDate) {
		return true
	}

	return false
}

// SetEnabledFromDate gets a reference to the given string and assigns it to the EnabledFromDate field.
func (o *PayoutSettingsRequest) SetEnabledFromDate(v string) {
	o.EnabledFromDate = &v
}

// GetTransferInstrumentId returns the TransferInstrumentId field value
func (o *PayoutSettingsRequest) GetTransferInstrumentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferInstrumentId
}

// GetTransferInstrumentIdOk returns a tuple with the TransferInstrumentId field value
// and a boolean to check if the value has been set.
func (o *PayoutSettingsRequest) GetTransferInstrumentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransferInstrumentId, true
}

// SetTransferInstrumentId sets field value
func (o *PayoutSettingsRequest) SetTransferInstrumentId(v string) {
	o.TransferInstrumentId = v
}

func (o PayoutSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayoutSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !common.IsNil(o.EnabledFromDate) {
		toSerialize["enabledFromDate"] = o.EnabledFromDate
	}
	toSerialize["transferInstrumentId"] = o.TransferInstrumentId
	return toSerialize, nil
}

type NullablePayoutSettingsRequest struct {
	value *PayoutSettingsRequest
	isSet bool
}

func (v NullablePayoutSettingsRequest) Get() *PayoutSettingsRequest {
	return v.value
}

func (v *NullablePayoutSettingsRequest) Set(val *PayoutSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePayoutSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePayoutSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayoutSettingsRequest(val *PayoutSettingsRequest) *NullablePayoutSettingsRequest {
	return &NullablePayoutSettingsRequest{value: val, isSet: true}
}

func (v NullablePayoutSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayoutSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
