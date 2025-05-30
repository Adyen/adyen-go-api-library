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

// checks if the RegisterSCAResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RegisterSCAResponse{}

// RegisterSCAResponse struct for RegisterSCAResponse
type RegisterSCAResponse struct {
	// The unique identifier of the SCA device you are registering.
	Id *string `json:"id,omitempty"`
	// The unique identifier of the payment instrument for which you are registering the SCA device.
	PaymentInstrumentId *string `json:"paymentInstrumentId,omitempty"`
	// A string that you must pass to the authentication SDK to continue with the registration process.
	SdkInput *string `json:"sdkInput,omitempty"`
	// Specifies if the registration was initiated successfully.
	Success *bool `json:"success,omitempty"`
}

// NewRegisterSCAResponse instantiates a new RegisterSCAResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterSCAResponse() *RegisterSCAResponse {
	this := RegisterSCAResponse{}
	return &this
}

// NewRegisterSCAResponseWithDefaults instantiates a new RegisterSCAResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterSCAResponseWithDefaults() *RegisterSCAResponse {
	this := RegisterSCAResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RegisterSCAResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterSCAResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RegisterSCAResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RegisterSCAResponse) SetId(v string) {
	o.Id = &v
}

// GetPaymentInstrumentId returns the PaymentInstrumentId field value if set, zero value otherwise.
func (o *RegisterSCAResponse) GetPaymentInstrumentId() string {
	if o == nil || common.IsNil(o.PaymentInstrumentId) {
		var ret string
		return ret
	}
	return *o.PaymentInstrumentId
}

// GetPaymentInstrumentIdOk returns a tuple with the PaymentInstrumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterSCAResponse) GetPaymentInstrumentIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentInstrumentId) {
		return nil, false
	}
	return o.PaymentInstrumentId, true
}

// HasPaymentInstrumentId returns a boolean if a field has been set.
func (o *RegisterSCAResponse) HasPaymentInstrumentId() bool {
	if o != nil && !common.IsNil(o.PaymentInstrumentId) {
		return true
	}

	return false
}

// SetPaymentInstrumentId gets a reference to the given string and assigns it to the PaymentInstrumentId field.
func (o *RegisterSCAResponse) SetPaymentInstrumentId(v string) {
	o.PaymentInstrumentId = &v
}

// GetSdkInput returns the SdkInput field value if set, zero value otherwise.
func (o *RegisterSCAResponse) GetSdkInput() string {
	if o == nil || common.IsNil(o.SdkInput) {
		var ret string
		return ret
	}
	return *o.SdkInput
}

// GetSdkInputOk returns a tuple with the SdkInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterSCAResponse) GetSdkInputOk() (*string, bool) {
	if o == nil || common.IsNil(o.SdkInput) {
		return nil, false
	}
	return o.SdkInput, true
}

// HasSdkInput returns a boolean if a field has been set.
func (o *RegisterSCAResponse) HasSdkInput() bool {
	if o != nil && !common.IsNil(o.SdkInput) {
		return true
	}

	return false
}

// SetSdkInput gets a reference to the given string and assigns it to the SdkInput field.
func (o *RegisterSCAResponse) SetSdkInput(v string) {
	o.SdkInput = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *RegisterSCAResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterSCAResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *RegisterSCAResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *RegisterSCAResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o RegisterSCAResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterSCAResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.PaymentInstrumentId) {
		toSerialize["paymentInstrumentId"] = o.PaymentInstrumentId
	}
	if !common.IsNil(o.SdkInput) {
		toSerialize["sdkInput"] = o.SdkInput
	}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableRegisterSCAResponse struct {
	value *RegisterSCAResponse
	isSet bool
}

func (v NullableRegisterSCAResponse) Get() *RegisterSCAResponse {
	return v.value
}

func (v *NullableRegisterSCAResponse) Set(val *RegisterSCAResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterSCAResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterSCAResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterSCAResponse(val *RegisterSCAResponse) *NullableRegisterSCAResponse {
	return &NullableRegisterSCAResponse{value: val, isSet: true}
}

func (v NullableRegisterSCAResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterSCAResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
