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

// checks if the CreateMerchantRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateMerchantRequest{}

// CreateMerchantRequest struct for CreateMerchantRequest
type CreateMerchantRequest struct {
	// The unique identifier of the [business line](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/businessLines). Required for an Adyen for Platforms Manage integration.
	BusinessLineId *string `json:"businessLineId,omitempty"`
	// The unique identifier of the company account.
	CompanyId string `json:"companyId"`
	// Your description for the merchant account, maximum 300 characters.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities). Required for an Adyen for Platforms Manage integration.
	LegalEntityId *string `json:"legalEntityId,omitempty"`
	// Sets the pricing plan for the merchant account. Required for an Adyen for Platforms Manage integration. Your Adyen contact will provide the values that you can use.
	PricingPlan *string `json:"pricingPlan,omitempty"`
	// Your reference for the merchant account. To make this reference the unique identifier of the merchant account, your Adyen contact can set up a template on your company account. The template can have 6 to 255 characters with upper- and lower-case letters, underscores, and numbers. When your company account has a template, then the `reference` is required and must be unique within the company account.
	Reference *string `json:"reference,omitempty"`
	// List of sales channels that the merchant will process payments with
	SalesChannels []string `json:"salesChannels,omitempty"`
}

// NewCreateMerchantRequest instantiates a new CreateMerchantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMerchantRequest(companyId string) *CreateMerchantRequest {
	this := CreateMerchantRequest{}
	this.CompanyId = companyId
	return &this
}

// NewCreateMerchantRequestWithDefaults instantiates a new CreateMerchantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMerchantRequestWithDefaults() *CreateMerchantRequest {
	this := CreateMerchantRequest{}
	return &this
}

// GetBusinessLineId returns the BusinessLineId field value if set, zero value otherwise.
func (o *CreateMerchantRequest) GetBusinessLineId() string {
	if o == nil || common.IsNil(o.BusinessLineId) {
		var ret string
		return ret
	}
	return *o.BusinessLineId
}

// GetBusinessLineIdOk returns a tuple with the BusinessLineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantRequest) GetBusinessLineIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BusinessLineId) {
		return nil, false
	}
	return o.BusinessLineId, true
}

// HasBusinessLineId returns a boolean if a field has been set.
func (o *CreateMerchantRequest) HasBusinessLineId() bool {
	if o != nil && !common.IsNil(o.BusinessLineId) {
		return true
	}

	return false
}

// SetBusinessLineId gets a reference to the given string and assigns it to the BusinessLineId field.
func (o *CreateMerchantRequest) SetBusinessLineId(v string) {
	o.BusinessLineId = &v
}

// GetCompanyId returns the CompanyId field value
func (o *CreateMerchantRequest) GetCompanyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *CreateMerchantRequest) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *CreateMerchantRequest) SetCompanyId(v string) {
	o.CompanyId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateMerchantRequest) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateMerchantRequest) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateMerchantRequest) SetDescription(v string) {
	o.Description = &v
}

// GetLegalEntityId returns the LegalEntityId field value if set, zero value otherwise.
func (o *CreateMerchantRequest) GetLegalEntityId() string {
	if o == nil || common.IsNil(o.LegalEntityId) {
		var ret string
		return ret
	}
	return *o.LegalEntityId
}

// GetLegalEntityIdOk returns a tuple with the LegalEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantRequest) GetLegalEntityIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.LegalEntityId) {
		return nil, false
	}
	return o.LegalEntityId, true
}

// HasLegalEntityId returns a boolean if a field has been set.
func (o *CreateMerchantRequest) HasLegalEntityId() bool {
	if o != nil && !common.IsNil(o.LegalEntityId) {
		return true
	}

	return false
}

// SetLegalEntityId gets a reference to the given string and assigns it to the LegalEntityId field.
func (o *CreateMerchantRequest) SetLegalEntityId(v string) {
	o.LegalEntityId = &v
}

// GetPricingPlan returns the PricingPlan field value if set, zero value otherwise.
func (o *CreateMerchantRequest) GetPricingPlan() string {
	if o == nil || common.IsNil(o.PricingPlan) {
		var ret string
		return ret
	}
	return *o.PricingPlan
}

// GetPricingPlanOk returns a tuple with the PricingPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantRequest) GetPricingPlanOk() (*string, bool) {
	if o == nil || common.IsNil(o.PricingPlan) {
		return nil, false
	}
	return o.PricingPlan, true
}

// HasPricingPlan returns a boolean if a field has been set.
func (o *CreateMerchantRequest) HasPricingPlan() bool {
	if o != nil && !common.IsNil(o.PricingPlan) {
		return true
	}

	return false
}

// SetPricingPlan gets a reference to the given string and assigns it to the PricingPlan field.
func (o *CreateMerchantRequest) SetPricingPlan(v string) {
	o.PricingPlan = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *CreateMerchantRequest) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantRequest) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *CreateMerchantRequest) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *CreateMerchantRequest) SetReference(v string) {
	o.Reference = &v
}

// GetSalesChannels returns the SalesChannels field value if set, zero value otherwise.
func (o *CreateMerchantRequest) GetSalesChannels() []string {
	if o == nil || common.IsNil(o.SalesChannels) {
		var ret []string
		return ret
	}
	return o.SalesChannels
}

// GetSalesChannelsOk returns a tuple with the SalesChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantRequest) GetSalesChannelsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.SalesChannels) {
		return nil, false
	}
	return o.SalesChannels, true
}

// HasSalesChannels returns a boolean if a field has been set.
func (o *CreateMerchantRequest) HasSalesChannels() bool {
	if o != nil && !common.IsNil(o.SalesChannels) {
		return true
	}

	return false
}

// SetSalesChannels gets a reference to the given []string and assigns it to the SalesChannels field.
func (o *CreateMerchantRequest) SetSalesChannels(v []string) {
	o.SalesChannels = v
}

func (o CreateMerchantRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMerchantRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BusinessLineId) {
		toSerialize["businessLineId"] = o.BusinessLineId
	}
	toSerialize["companyId"] = o.CompanyId
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.LegalEntityId) {
		toSerialize["legalEntityId"] = o.LegalEntityId
	}
	if !common.IsNil(o.PricingPlan) {
		toSerialize["pricingPlan"] = o.PricingPlan
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.SalesChannels) {
		toSerialize["salesChannels"] = o.SalesChannels
	}
	return toSerialize, nil
}

type NullableCreateMerchantRequest struct {
	value *CreateMerchantRequest
	isSet bool
}

func (v NullableCreateMerchantRequest) Get() *CreateMerchantRequest {
	return v.value
}

func (v *NullableCreateMerchantRequest) Set(val *CreateMerchantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMerchantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMerchantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMerchantRequest(val *CreateMerchantRequest) *NullableCreateMerchantRequest {
	return &NullableCreateMerchantRequest{value: val, isSet: true}
}

func (v NullableCreateMerchantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMerchantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
