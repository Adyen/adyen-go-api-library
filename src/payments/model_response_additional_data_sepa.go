/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the ResponseAdditionalDataSepa type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ResponseAdditionalDataSepa{}

// ResponseAdditionalDataSepa struct for ResponseAdditionalDataSepa
type ResponseAdditionalDataSepa struct {
	// The transaction signature date.  Format: yyyy-MM-dd
	SepadirectdebitDateOfSignature *string `json:"sepadirectdebit.dateOfSignature,omitempty"`
	// Its value corresponds to the pspReference value of the transaction.
	SepadirectdebitMandateId *string `json:"sepadirectdebit.mandateId,omitempty"`
	// This field can take one of the following values: * OneOff: (OOFF) Direct debit instruction to initiate exactly one direct debit transaction.  * First: (FRST) Initial/first collection in a series of direct debit instructions. * Recurring: (RCUR) Direct debit instruction to carry out regular direct debit transactions initiated by the creditor. * Final: (FNAL) Last/final collection in a series of direct debit instructions.  Example: OOFF
	SepadirectdebitSequenceType *string `json:"sepadirectdebit.sequenceType,omitempty"`
}

// NewResponseAdditionalDataSepa instantiates a new ResponseAdditionalDataSepa object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAdditionalDataSepa() *ResponseAdditionalDataSepa {
	this := ResponseAdditionalDataSepa{}
	return &this
}

// NewResponseAdditionalDataSepaWithDefaults instantiates a new ResponseAdditionalDataSepa object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAdditionalDataSepaWithDefaults() *ResponseAdditionalDataSepa {
	this := ResponseAdditionalDataSepa{}
	return &this
}

// GetSepadirectdebitDateOfSignature returns the SepadirectdebitDateOfSignature field value if set, zero value otherwise.
func (o *ResponseAdditionalDataSepa) GetSepadirectdebitDateOfSignature() string {
	if o == nil || common.IsNil(o.SepadirectdebitDateOfSignature) {
		var ret string
		return ret
	}
	return *o.SepadirectdebitDateOfSignature
}

// GetSepadirectdebitDateOfSignatureOk returns a tuple with the SepadirectdebitDateOfSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataSepa) GetSepadirectdebitDateOfSignatureOk() (*string, bool) {
	if o == nil || common.IsNil(o.SepadirectdebitDateOfSignature) {
		return nil, false
	}
	return o.SepadirectdebitDateOfSignature, true
}

// HasSepadirectdebitDateOfSignature returns a boolean if a field has been set.
func (o *ResponseAdditionalDataSepa) HasSepadirectdebitDateOfSignature() bool {
	if o != nil && !common.IsNil(o.SepadirectdebitDateOfSignature) {
		return true
	}

	return false
}

// SetSepadirectdebitDateOfSignature gets a reference to the given string and assigns it to the SepadirectdebitDateOfSignature field.
func (o *ResponseAdditionalDataSepa) SetSepadirectdebitDateOfSignature(v string) {
	o.SepadirectdebitDateOfSignature = &v
}

// GetSepadirectdebitMandateId returns the SepadirectdebitMandateId field value if set, zero value otherwise.
func (o *ResponseAdditionalDataSepa) GetSepadirectdebitMandateId() string {
	if o == nil || common.IsNil(o.SepadirectdebitMandateId) {
		var ret string
		return ret
	}
	return *o.SepadirectdebitMandateId
}

// GetSepadirectdebitMandateIdOk returns a tuple with the SepadirectdebitMandateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataSepa) GetSepadirectdebitMandateIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.SepadirectdebitMandateId) {
		return nil, false
	}
	return o.SepadirectdebitMandateId, true
}

// HasSepadirectdebitMandateId returns a boolean if a field has been set.
func (o *ResponseAdditionalDataSepa) HasSepadirectdebitMandateId() bool {
	if o != nil && !common.IsNil(o.SepadirectdebitMandateId) {
		return true
	}

	return false
}

// SetSepadirectdebitMandateId gets a reference to the given string and assigns it to the SepadirectdebitMandateId field.
func (o *ResponseAdditionalDataSepa) SetSepadirectdebitMandateId(v string) {
	o.SepadirectdebitMandateId = &v
}

// GetSepadirectdebitSequenceType returns the SepadirectdebitSequenceType field value if set, zero value otherwise.
func (o *ResponseAdditionalDataSepa) GetSepadirectdebitSequenceType() string {
	if o == nil || common.IsNil(o.SepadirectdebitSequenceType) {
		var ret string
		return ret
	}
	return *o.SepadirectdebitSequenceType
}

// GetSepadirectdebitSequenceTypeOk returns a tuple with the SepadirectdebitSequenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataSepa) GetSepadirectdebitSequenceTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SepadirectdebitSequenceType) {
		return nil, false
	}
	return o.SepadirectdebitSequenceType, true
}

// HasSepadirectdebitSequenceType returns a boolean if a field has been set.
func (o *ResponseAdditionalDataSepa) HasSepadirectdebitSequenceType() bool {
	if o != nil && !common.IsNil(o.SepadirectdebitSequenceType) {
		return true
	}

	return false
}

// SetSepadirectdebitSequenceType gets a reference to the given string and assigns it to the SepadirectdebitSequenceType field.
func (o *ResponseAdditionalDataSepa) SetSepadirectdebitSequenceType(v string) {
	o.SepadirectdebitSequenceType = &v
}

func (o ResponseAdditionalDataSepa) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseAdditionalDataSepa) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.SepadirectdebitDateOfSignature) {
		toSerialize["sepadirectdebit.dateOfSignature"] = o.SepadirectdebitDateOfSignature
	}
	if !common.IsNil(o.SepadirectdebitMandateId) {
		toSerialize["sepadirectdebit.mandateId"] = o.SepadirectdebitMandateId
	}
	if !common.IsNil(o.SepadirectdebitSequenceType) {
		toSerialize["sepadirectdebit.sequenceType"] = o.SepadirectdebitSequenceType
	}
	return toSerialize, nil
}

type NullableResponseAdditionalDataSepa struct {
	value *ResponseAdditionalDataSepa
	isSet bool
}

func (v NullableResponseAdditionalDataSepa) Get() *ResponseAdditionalDataSepa {
	return v.value
}

func (v *NullableResponseAdditionalDataSepa) Set(val *ResponseAdditionalDataSepa) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAdditionalDataSepa) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAdditionalDataSepa) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAdditionalDataSepa(val *ResponseAdditionalDataSepa) *NullableResponseAdditionalDataSepa {
	return &NullableResponseAdditionalDataSepa{value: val, isSet: true}
}

func (v NullableResponseAdditionalDataSepa) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAdditionalDataSepa) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
