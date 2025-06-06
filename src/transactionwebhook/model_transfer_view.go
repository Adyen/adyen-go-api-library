/*
Transaction webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transactionwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the TransferView type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferView{}

// TransferView struct for TransferView
type TransferView struct {
	CategoryData *TransferViewCategoryData `json:"categoryData,omitempty"`
	// The ID of the resource.
	Id *string `json:"id,omitempty"`
	// The [`reference`](https://docs.adyen.com/api-explorer/#/transfers/latest/post/transfers__reqParam_reference) from the `/transfers` request. If you haven't provided any, Adyen generates a unique reference.
	Reference string `json:"reference"`
}

// NewTransferView instantiates a new TransferView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferView(reference string) *TransferView {
	this := TransferView{}
	this.Reference = reference
	return &this
}

// NewTransferViewWithDefaults instantiates a new TransferView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferViewWithDefaults() *TransferView {
	this := TransferView{}
	return &this
}

// GetCategoryData returns the CategoryData field value if set, zero value otherwise.
func (o *TransferView) GetCategoryData() TransferViewCategoryData {
	if o == nil || common.IsNil(o.CategoryData) {
		var ret TransferViewCategoryData
		return ret
	}
	return *o.CategoryData
}

// GetCategoryDataOk returns a tuple with the CategoryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferView) GetCategoryDataOk() (*TransferViewCategoryData, bool) {
	if o == nil || common.IsNil(o.CategoryData) {
		return nil, false
	}
	return o.CategoryData, true
}

// HasCategoryData returns a boolean if a field has been set.
func (o *TransferView) HasCategoryData() bool {
	if o != nil && !common.IsNil(o.CategoryData) {
		return true
	}

	return false
}

// SetCategoryData gets a reference to the given TransferViewCategoryData and assigns it to the CategoryData field.
func (o *TransferView) SetCategoryData(v TransferViewCategoryData) {
	o.CategoryData = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransferView) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferView) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransferView) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TransferView) SetId(v string) {
	o.Id = &v
}

// GetReference returns the Reference field value
func (o *TransferView) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *TransferView) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *TransferView) SetReference(v string) {
	o.Reference = v
}

func (o TransferView) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CategoryData) {
		toSerialize["categoryData"] = o.CategoryData
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["reference"] = o.Reference
	return toSerialize, nil
}

type NullableTransferView struct {
	value *TransferView
	isSet bool
}

func (v NullableTransferView) Get() *TransferView {
	return v.value
}

func (v *NullableTransferView) Set(val *TransferView) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferView) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferView(val *TransferView) *NullableTransferView {
	return &NullableTransferView{value: val, isSet: true}
}

func (v NullableTransferView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
