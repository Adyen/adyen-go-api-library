/*
Transfer webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the TransferNotificationRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferNotificationRequest{}

// TransferNotificationRequest struct for TransferNotificationRequest
type TransferNotificationRequest struct {
	Data TransferData `json:"data"`
	// The environment from which the webhook originated.  Possible values: **test**, **live**.
	Environment string `json:"environment"`
	// When the event was queued.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// The type of webhook.
	Type *string `json:"type,omitempty"`
}

// NewTransferNotificationRequest instantiates a new TransferNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferNotificationRequest(data TransferData, environment string) *TransferNotificationRequest {
	this := TransferNotificationRequest{}
	this.Data = data
	this.Environment = environment
	return &this
}

// NewTransferNotificationRequestWithDefaults instantiates a new TransferNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferNotificationRequestWithDefaults() *TransferNotificationRequest {
	this := TransferNotificationRequest{}
	return &this
}

// GetData returns the Data field value
func (o *TransferNotificationRequest) GetData() TransferData {
	if o == nil {
		var ret TransferData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TransferNotificationRequest) GetDataOk() (*TransferData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TransferNotificationRequest) SetData(v TransferData) {
	o.Data = v
}

// GetEnvironment returns the Environment field value
func (o *TransferNotificationRequest) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *TransferNotificationRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *TransferNotificationRequest) SetEnvironment(v string) {
	o.Environment = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TransferNotificationRequest) GetTimestamp() time.Time {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferNotificationRequest) GetTimestampOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TransferNotificationRequest) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *TransferNotificationRequest) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransferNotificationRequest) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferNotificationRequest) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransferNotificationRequest) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransferNotificationRequest) SetType(v string) {
	o.Type = &v
}

func (o TransferNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["environment"] = o.Environment
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTransferNotificationRequest struct {
	value *TransferNotificationRequest
	isSet bool
}

func (v NullableTransferNotificationRequest) Get() *TransferNotificationRequest {
	return v.value
}

func (v *NullableTransferNotificationRequest) Set(val *TransferNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferNotificationRequest(val *TransferNotificationRequest) *NullableTransferNotificationRequest {
	return &NullableTransferNotificationRequest{value: val, isSet: true}
}

func (v NullableTransferNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *TransferNotificationRequest) isValidType() bool {
	var allowedEnumValues = []string{"balancePlatform.transfer.created", "balancePlatform.transfer.updated"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
