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

// checks if the CancelOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelOrderResponse{}

// CancelOrderResponse struct for CancelOrderResponse
type CancelOrderResponse struct {
	// A unique reference of the cancellation request.
	PspReference string `json:"pspReference"`
	// The result of the cancellation request.  Possible values:  * **Received** – Indicates the cancellation has successfully been received by Adyen, and will be processed.
	ResultCode string `json:"resultCode"`
}

// NewCancelOrderResponse instantiates a new CancelOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelOrderResponse(pspReference string, resultCode string) *CancelOrderResponse {
	this := CancelOrderResponse{}
	this.PspReference = pspReference
	this.ResultCode = resultCode
	return &this
}

// NewCancelOrderResponseWithDefaults instantiates a new CancelOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelOrderResponseWithDefaults() *CancelOrderResponse {
	this := CancelOrderResponse{}
	return &this
}

// GetPspReference returns the PspReference field value
func (o *CancelOrderResponse) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *CancelOrderResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *CancelOrderResponse) SetPspReference(v string) {
	o.PspReference = v
}

// GetResultCode returns the ResultCode field value
func (o *CancelOrderResponse) GetResultCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value
// and a boolean to check if the value has been set.
func (o *CancelOrderResponse) GetResultCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultCode, true
}

// SetResultCode sets field value
func (o *CancelOrderResponse) SetResultCode(v string) {
	o.ResultCode = v
}

func (o CancelOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pspReference"] = o.PspReference
	toSerialize["resultCode"] = o.ResultCode
	return toSerialize, nil
}

type NullableCancelOrderResponse struct {
	value *CancelOrderResponse
	isSet bool
}

func (v NullableCancelOrderResponse) Get() *CancelOrderResponse {
	return v.value
}

func (v *NullableCancelOrderResponse) Set(val *CancelOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelOrderResponse(val *CancelOrderResponse) *NullableCancelOrderResponse {
	return &NullableCancelOrderResponse{value: val, isSet: true}
}

func (v NullableCancelOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CancelOrderResponse) isValidResultCode() bool {
	var allowedEnumValues = []string{"Received"}
	for _, allowed := range allowedEnumValues {
		if o.GetResultCode() == allowed {
			return true
		}
	}
	return false
}
