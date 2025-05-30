/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the PaymentMethodUPIApps type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentMethodUPIApps{}

// PaymentMethodUPIApps struct for PaymentMethodUPIApps
type PaymentMethodUPIApps struct {
	// The unique identifier of this app, to submit in requests to /payments.
	Id string `json:"id"`
	// A localized name of the app.
	Name string `json:"name"`
}

// NewPaymentMethodUPIApps instantiates a new PaymentMethodUPIApps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodUPIApps(id string, name string) *PaymentMethodUPIApps {
	this := PaymentMethodUPIApps{}
	this.Id = id
	this.Name = name
	return &this
}

// NewPaymentMethodUPIAppsWithDefaults instantiates a new PaymentMethodUPIApps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodUPIAppsWithDefaults() *PaymentMethodUPIApps {
	this := PaymentMethodUPIApps{}
	return &this
}

// GetId returns the Id field value
func (o *PaymentMethodUPIApps) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodUPIApps) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaymentMethodUPIApps) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PaymentMethodUPIApps) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodUPIApps) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PaymentMethodUPIApps) SetName(v string) {
	o.Name = v
}

func (o PaymentMethodUPIApps) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodUPIApps) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullablePaymentMethodUPIApps struct {
	value *PaymentMethodUPIApps
	isSet bool
}

func (v NullablePaymentMethodUPIApps) Get() *PaymentMethodUPIApps {
	return v.value
}

func (v *NullablePaymentMethodUPIApps) Set(val *PaymentMethodUPIApps) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodUPIApps) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodUPIApps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodUPIApps(val *PaymentMethodUPIApps) *NullablePaymentMethodUPIApps {
	return &NullablePaymentMethodUPIApps{value: val, isSet: true}
}

func (v NullablePaymentMethodUPIApps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodUPIApps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
