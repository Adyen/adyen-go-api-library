/*
Management Webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package managementwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the PaymentMethodNotificationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentMethodNotificationResponse{}

// PaymentMethodNotificationResponse struct for PaymentMethodNotificationResponse
type PaymentMethodNotificationResponse struct {
	// Respond with any **2xx** HTTP status code to [accept the webhook](https://docs.adyen.com/development-resources/webhooks#accept-notifications).
	NotificationResponse *string `json:"notificationResponse,omitempty"`
}

// NewPaymentMethodNotificationResponse instantiates a new PaymentMethodNotificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodNotificationResponse() *PaymentMethodNotificationResponse {
	this := PaymentMethodNotificationResponse{}
	return &this
}

// NewPaymentMethodNotificationResponseWithDefaults instantiates a new PaymentMethodNotificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodNotificationResponseWithDefaults() *PaymentMethodNotificationResponse {
	this := PaymentMethodNotificationResponse{}
	return &this
}

// GetNotificationResponse returns the NotificationResponse field value if set, zero value otherwise.
func (o *PaymentMethodNotificationResponse) GetNotificationResponse() string {
	if o == nil || common.IsNil(o.NotificationResponse) {
		var ret string
		return ret
	}
	return *o.NotificationResponse
}

// GetNotificationResponseOk returns a tuple with the NotificationResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodNotificationResponse) GetNotificationResponseOk() (*string, bool) {
	if o == nil || common.IsNil(o.NotificationResponse) {
		return nil, false
	}
	return o.NotificationResponse, true
}

// HasNotificationResponse returns a boolean if a field has been set.
func (o *PaymentMethodNotificationResponse) HasNotificationResponse() bool {
	if o != nil && !common.IsNil(o.NotificationResponse) {
		return true
	}

	return false
}

// SetNotificationResponse gets a reference to the given string and assigns it to the NotificationResponse field.
func (o *PaymentMethodNotificationResponse) SetNotificationResponse(v string) {
	o.NotificationResponse = &v
}

func (o PaymentMethodNotificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodNotificationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.NotificationResponse) {
		toSerialize["notificationResponse"] = o.NotificationResponse
	}
	return toSerialize, nil
}

type NullablePaymentMethodNotificationResponse struct {
	value *PaymentMethodNotificationResponse
	isSet bool
}

func (v NullablePaymentMethodNotificationResponse) Get() *PaymentMethodNotificationResponse {
	return v.value
}

func (v *NullablePaymentMethodNotificationResponse) Set(val *PaymentMethodNotificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodNotificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodNotificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodNotificationResponse(val *PaymentMethodNotificationResponse) *NullablePaymentMethodNotificationResponse {
	return &NullablePaymentMethodNotificationResponse{value: val, isSet: true}
}

func (v NullablePaymentMethodNotificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodNotificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
