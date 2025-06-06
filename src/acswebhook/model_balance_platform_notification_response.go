/*
Authentication webhooks

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package acswebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the BalancePlatformNotificationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BalancePlatformNotificationResponse{}

// BalancePlatformNotificationResponse struct for BalancePlatformNotificationResponse
type BalancePlatformNotificationResponse struct {
	// Respond with any **2xx** HTTP status code to [accept the webhook](https://docs.adyen.com/development-resources/webhooks#accept-notifications).
	NotificationResponse *string `json:"notificationResponse,omitempty"`
}

// NewBalancePlatformNotificationResponse instantiates a new BalancePlatformNotificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalancePlatformNotificationResponse() *BalancePlatformNotificationResponse {
	this := BalancePlatformNotificationResponse{}
	return &this
}

// NewBalancePlatformNotificationResponseWithDefaults instantiates a new BalancePlatformNotificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalancePlatformNotificationResponseWithDefaults() *BalancePlatformNotificationResponse {
	this := BalancePlatformNotificationResponse{}
	return &this
}

// GetNotificationResponse returns the NotificationResponse field value if set, zero value otherwise.
func (o *BalancePlatformNotificationResponse) GetNotificationResponse() string {
	if o == nil || common.IsNil(o.NotificationResponse) {
		var ret string
		return ret
	}
	return *o.NotificationResponse
}

// GetNotificationResponseOk returns a tuple with the NotificationResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancePlatformNotificationResponse) GetNotificationResponseOk() (*string, bool) {
	if o == nil || common.IsNil(o.NotificationResponse) {
		return nil, false
	}
	return o.NotificationResponse, true
}

// HasNotificationResponse returns a boolean if a field has been set.
func (o *BalancePlatformNotificationResponse) HasNotificationResponse() bool {
	if o != nil && !common.IsNil(o.NotificationResponse) {
		return true
	}

	return false
}

// SetNotificationResponse gets a reference to the given string and assigns it to the NotificationResponse field.
func (o *BalancePlatformNotificationResponse) SetNotificationResponse(v string) {
	o.NotificationResponse = &v
}

func (o BalancePlatformNotificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalancePlatformNotificationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.NotificationResponse) {
		toSerialize["notificationResponse"] = o.NotificationResponse
	}
	return toSerialize, nil
}

type NullableBalancePlatformNotificationResponse struct {
	value *BalancePlatformNotificationResponse
	isSet bool
}

func (v NullableBalancePlatformNotificationResponse) Get() *BalancePlatformNotificationResponse {
	return v.value
}

func (v *NullableBalancePlatformNotificationResponse) Set(val *BalancePlatformNotificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBalancePlatformNotificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBalancePlatformNotificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalancePlatformNotificationResponse(val *BalancePlatformNotificationResponse) *NullableBalancePlatformNotificationResponse {
	return &NullableBalancePlatformNotificationResponse{value: val, isSet: true}
}

func (v NullableBalancePlatformNotificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalancePlatformNotificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
