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

// checks if the ListStoresResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListStoresResponse{}

// ListStoresResponse struct for ListStoresResponse
type ListStoresResponse struct {
	Links *PaginationLinks `json:"_links,omitempty"`
	// List of stores
	Data []Store `json:"data,omitempty"`
	// Total number of items.
	ItemsTotal int32 `json:"itemsTotal"`
	// Total number of pages.
	PagesTotal int32 `json:"pagesTotal"`
}

// NewListStoresResponse instantiates a new ListStoresResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListStoresResponse(itemsTotal int32, pagesTotal int32) *ListStoresResponse {
	this := ListStoresResponse{}
	this.ItemsTotal = itemsTotal
	this.PagesTotal = pagesTotal
	return &this
}

// NewListStoresResponseWithDefaults instantiates a new ListStoresResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListStoresResponseWithDefaults() *ListStoresResponse {
	this := ListStoresResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListStoresResponse) GetLinks() PaginationLinks {
	if o == nil || common.IsNil(o.Links) {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStoresResponse) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || common.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListStoresResponse) HasLinks() bool {
	if o != nil && !common.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *ListStoresResponse) SetLinks(v PaginationLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListStoresResponse) GetData() []Store {
	if o == nil || common.IsNil(o.Data) {
		var ret []Store
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStoresResponse) GetDataOk() ([]Store, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListStoresResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Store and assigns it to the Data field.
func (o *ListStoresResponse) SetData(v []Store) {
	o.Data = v
}

// GetItemsTotal returns the ItemsTotal field value
func (o *ListStoresResponse) GetItemsTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ItemsTotal
}

// GetItemsTotalOk returns a tuple with the ItemsTotal field value
// and a boolean to check if the value has been set.
func (o *ListStoresResponse) GetItemsTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemsTotal, true
}

// SetItemsTotal sets field value
func (o *ListStoresResponse) SetItemsTotal(v int32) {
	o.ItemsTotal = v
}

// GetPagesTotal returns the PagesTotal field value
func (o *ListStoresResponse) GetPagesTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PagesTotal
}

// GetPagesTotalOk returns a tuple with the PagesTotal field value
// and a boolean to check if the value has been set.
func (o *ListStoresResponse) GetPagesTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PagesTotal, true
}

// SetPagesTotal sets field value
func (o *ListStoresResponse) SetPagesTotal(v int32) {
	o.PagesTotal = v
}

func (o ListStoresResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListStoresResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["itemsTotal"] = o.ItemsTotal
	toSerialize["pagesTotal"] = o.PagesTotal
	return toSerialize, nil
}

type NullableListStoresResponse struct {
	value *ListStoresResponse
	isSet bool
}

func (v NullableListStoresResponse) Get() *ListStoresResponse {
	return v.value
}

func (v *NullableListStoresResponse) Set(val *ListStoresResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListStoresResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListStoresResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListStoresResponse(val *ListStoresResponse) *NullableListStoresResponse {
	return &NullableListStoresResponse{value: val, isSet: true}
}

func (v NullableListStoresResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListStoresResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
