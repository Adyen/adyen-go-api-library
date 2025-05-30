/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the GetPciQuestionnaireInfosResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPciQuestionnaireInfosResponse{}

// GetPciQuestionnaireInfosResponse struct for GetPciQuestionnaireInfosResponse
type GetPciQuestionnaireInfosResponse struct {
	// Information about the signed PCI questionnaires.
	Data []PciDocumentInfo `json:"data,omitempty"`
}

// NewGetPciQuestionnaireInfosResponse instantiates a new GetPciQuestionnaireInfosResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPciQuestionnaireInfosResponse() *GetPciQuestionnaireInfosResponse {
	this := GetPciQuestionnaireInfosResponse{}
	return &this
}

// NewGetPciQuestionnaireInfosResponseWithDefaults instantiates a new GetPciQuestionnaireInfosResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPciQuestionnaireInfosResponseWithDefaults() *GetPciQuestionnaireInfosResponse {
	this := GetPciQuestionnaireInfosResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetPciQuestionnaireInfosResponse) GetData() []PciDocumentInfo {
	if o == nil || common.IsNil(o.Data) {
		var ret []PciDocumentInfo
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPciQuestionnaireInfosResponse) GetDataOk() ([]PciDocumentInfo, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetPciQuestionnaireInfosResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PciDocumentInfo and assigns it to the Data field.
func (o *GetPciQuestionnaireInfosResponse) SetData(v []PciDocumentInfo) {
	o.Data = v
}

func (o GetPciQuestionnaireInfosResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPciQuestionnaireInfosResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetPciQuestionnaireInfosResponse struct {
	value *GetPciQuestionnaireInfosResponse
	isSet bool
}

func (v NullableGetPciQuestionnaireInfosResponse) Get() *GetPciQuestionnaireInfosResponse {
	return v.value
}

func (v *NullableGetPciQuestionnaireInfosResponse) Set(val *GetPciQuestionnaireInfosResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPciQuestionnaireInfosResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPciQuestionnaireInfosResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPciQuestionnaireInfosResponse(val *GetPciQuestionnaireInfosResponse) *NullableGetPciQuestionnaireInfosResponse {
	return &NullableGetPciQuestionnaireInfosResponse{value: val, isSet: true}
}

func (v NullableGetPciQuestionnaireInfosResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPciQuestionnaireInfosResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
