/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the ListMerchantResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListMerchantResponse{}

// ListMerchantResponse struct for ListMerchantResponse
type ListMerchantResponse struct {
	Links *PaginationLinks `json:"_links,omitempty"`
	// The list of merchant accounts.
	Data []Merchant `json:"data,omitempty"`
	// Total number of items.
	ItemsTotal int32 `json:"itemsTotal"`
	// Total number of pages.
	PagesTotal int32 `json:"pagesTotal"`
}

// NewListMerchantResponse instantiates a new ListMerchantResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMerchantResponse(itemsTotal int32, pagesTotal int32) *ListMerchantResponse {
	this := ListMerchantResponse{}
	this.ItemsTotal = itemsTotal
	this.PagesTotal = pagesTotal
	return &this
}

// NewListMerchantResponseWithDefaults instantiates a new ListMerchantResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMerchantResponseWithDefaults() *ListMerchantResponse {
	this := ListMerchantResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListMerchantResponse) GetLinks() PaginationLinks {
	if o == nil || common.IsNil(o.Links) {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMerchantResponse) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || common.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListMerchantResponse) HasLinks() bool {
	if o != nil && !common.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *ListMerchantResponse) SetLinks(v PaginationLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListMerchantResponse) GetData() []Merchant {
	if o == nil || common.IsNil(o.Data) {
		var ret []Merchant
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMerchantResponse) GetDataOk() ([]Merchant, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListMerchantResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Merchant and assigns it to the Data field.
func (o *ListMerchantResponse) SetData(v []Merchant) {
	o.Data = v
}

// GetItemsTotal returns the ItemsTotal field value
func (o *ListMerchantResponse) GetItemsTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ItemsTotal
}

// GetItemsTotalOk returns a tuple with the ItemsTotal field value
// and a boolean to check if the value has been set.
func (o *ListMerchantResponse) GetItemsTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemsTotal, true
}

// SetItemsTotal sets field value
func (o *ListMerchantResponse) SetItemsTotal(v int32) {
	o.ItemsTotal = v
}

// GetPagesTotal returns the PagesTotal field value
func (o *ListMerchantResponse) GetPagesTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PagesTotal
}

// GetPagesTotalOk returns a tuple with the PagesTotal field value
// and a boolean to check if the value has been set.
func (o *ListMerchantResponse) GetPagesTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PagesTotal, true
}

// SetPagesTotal sets field value
func (o *ListMerchantResponse) SetPagesTotal(v int32) {
	o.PagesTotal = v
}

func (o ListMerchantResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMerchantResponse) ToMap() (map[string]interface{}, error) {
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

type NullableListMerchantResponse struct {
	value *ListMerchantResponse
	isSet bool
}

func (v NullableListMerchantResponse) Get() *ListMerchantResponse {
	return v.value
}

func (v *NullableListMerchantResponse) Set(val *ListMerchantResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListMerchantResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListMerchantResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMerchantResponse(val *ListMerchantResponse) *NullableListMerchantResponse {
	return &NullableListMerchantResponse{value: val, isSet: true}
}

func (v NullableListMerchantResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMerchantResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}