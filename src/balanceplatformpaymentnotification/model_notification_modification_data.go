/*
Payment webhooks (deprecated)

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatformpaymentnotification

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the NotificationModificationData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NotificationModificationData{}

// NotificationModificationData struct for NotificationModificationData
type NotificationModificationData struct {
	Amount *Amount `json:"amount,omitempty"`
	// The type of modification.  Possible values: **Authorised**, **Cancelled**, **Captured**, **Error**, **Expired**, **OutgoingTransfer**, **PendingIncomingTransfer**, **PendingRefund**, **IncomingTransfer**, **Refunded**, **Refused**, **AuthAdjustmentAuthorised**.
	Type *string `json:"type,omitempty"`
}

// NewNotificationModificationData instantiates a new NotificationModificationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationModificationData() *NotificationModificationData {
	this := NotificationModificationData{}
	return &this
}

// NewNotificationModificationDataWithDefaults instantiates a new NotificationModificationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationModificationDataWithDefaults() *NotificationModificationData {
	this := NotificationModificationData{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *NotificationModificationData) GetAmount() Amount {
	if o == nil || common.IsNil(o.Amount) {
		var ret Amount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationModificationData) GetAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *NotificationModificationData) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Amount and assigns it to the Amount field.
func (o *NotificationModificationData) SetAmount(v Amount) {
	o.Amount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NotificationModificationData) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationModificationData) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NotificationModificationData) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NotificationModificationData) SetType(v string) {
	o.Type = &v
}

func (o NotificationModificationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationModificationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableNotificationModificationData struct {
	value *NotificationModificationData
	isSet bool
}

func (v NullableNotificationModificationData) Get() *NotificationModificationData {
	return v.value
}

func (v *NullableNotificationModificationData) Set(val *NotificationModificationData) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationModificationData) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationModificationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationModificationData(val *NotificationModificationData) *NullableNotificationModificationData {
	return &NullableNotificationModificationData{value: val, isSet: true}
}

func (v NullableNotificationModificationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationModificationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}