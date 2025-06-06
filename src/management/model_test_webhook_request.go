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

// checks if the TestWebhookRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TestWebhookRequest{}

// TestWebhookRequest struct for TestWebhookRequest
type TestWebhookRequest struct {
	Notification *CustomNotification `json:"notification,omitempty"`
	// List of event codes for which to send test notifications. Only the webhook types below are supported.   Possible values if webhook `type`: **standard**:  * **AUTHORISATION** * **CHARGEBACK_REVERSED** * **ORDER_CLOSED** * **ORDER_OPENED** * **PAIDOUT_REVERSED** * **PAYOUT_THIRDPARTY** * **REFUNDED_REVERSED** * **REFUND_WITH_DATA** * **REPORT_AVAILABLE** * **CUSTOM** - set your custom notification fields in the [`notification`](https://docs.adyen.com/api-explorer/#/ManagementService/v1/post/companies/{companyId}/webhooks/{webhookId}/test__reqParam_notification) object.  Possible values if webhook `type`: **banktransfer-notification**:  * **PENDING**  Possible values if webhook `type`: **report-notification**:  * **REPORT_AVAILABLE**  Possible values if webhook `type`: **ideal-notification**:  * **AUTHORISATION**  Possible values if webhook `type`: **pending-notification**:  * **PENDING**
	Types []string `json:"types,omitempty"`
}

// NewTestWebhookRequest instantiates a new TestWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestWebhookRequest() *TestWebhookRequest {
	this := TestWebhookRequest{}
	return &this
}

// NewTestWebhookRequestWithDefaults instantiates a new TestWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestWebhookRequestWithDefaults() *TestWebhookRequest {
	this := TestWebhookRequest{}
	return &this
}

// GetNotification returns the Notification field value if set, zero value otherwise.
func (o *TestWebhookRequest) GetNotification() CustomNotification {
	if o == nil || common.IsNil(o.Notification) {
		var ret CustomNotification
		return ret
	}
	return *o.Notification
}

// GetNotificationOk returns a tuple with the Notification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestWebhookRequest) GetNotificationOk() (*CustomNotification, bool) {
	if o == nil || common.IsNil(o.Notification) {
		return nil, false
	}
	return o.Notification, true
}

// HasNotification returns a boolean if a field has been set.
func (o *TestWebhookRequest) HasNotification() bool {
	if o != nil && !common.IsNil(o.Notification) {
		return true
	}

	return false
}

// SetNotification gets a reference to the given CustomNotification and assigns it to the Notification field.
func (o *TestWebhookRequest) SetNotification(v CustomNotification) {
	o.Notification = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *TestWebhookRequest) GetTypes() []string {
	if o == nil || common.IsNil(o.Types) {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestWebhookRequest) GetTypesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *TestWebhookRequest) HasTypes() bool {
	if o != nil && !common.IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *TestWebhookRequest) SetTypes(v []string) {
	o.Types = v
}

func (o TestWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestWebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Notification) {
		toSerialize["notification"] = o.Notification
	}
	if !common.IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	return toSerialize, nil
}

type NullableTestWebhookRequest struct {
	value *TestWebhookRequest
	isSet bool
}

func (v NullableTestWebhookRequest) Get() *TestWebhookRequest {
	return v.value
}

func (v *NullableTestWebhookRequest) Set(val *TestWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTestWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTestWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestWebhookRequest(val *TestWebhookRequest) *NullableTestWebhookRequest {
	return &NullableTestWebhookRequest{value: val, isSet: true}
}

func (v NullableTestWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
