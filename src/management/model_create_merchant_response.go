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

// checks if the CreateMerchantResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateMerchantResponse{}

// CreateMerchantResponse struct for CreateMerchantResponse
type CreateMerchantResponse struct {
	// The unique identifier of the [business line](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/businessLines).
	BusinessLineId *string `json:"businessLineId,omitempty"`
	// The unique identifier of the company account.
	CompanyId *string `json:"companyId,omitempty"`
	// Your description for the merchant account, maximum 300 characters.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the merchant account. If Adyen set up a template for the `reference`, then the `id` will have the same value as the `reference` that you sent in the request. Otherwise, the value is generated by Adyen.
	Id *string `json:"id,omitempty"`
	// The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities).
	LegalEntityId *string `json:"legalEntityId,omitempty"`
	// Partner pricing plan for the merchant, applicable for merchants under AfP managed company accounts.
	PricingPlan *string `json:"pricingPlan,omitempty"`
	// Your reference for the merchant account.
	Reference *string `json:"reference,omitempty"`
}

// NewCreateMerchantResponse instantiates a new CreateMerchantResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMerchantResponse() *CreateMerchantResponse {
	this := CreateMerchantResponse{}
	return &this
}

// NewCreateMerchantResponseWithDefaults instantiates a new CreateMerchantResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMerchantResponseWithDefaults() *CreateMerchantResponse {
	this := CreateMerchantResponse{}
	return &this
}

// GetBusinessLineId returns the BusinessLineId field value if set, zero value otherwise.
func (o *CreateMerchantResponse) GetBusinessLineId() string {
	if o == nil || common.IsNil(o.BusinessLineId) {
		var ret string
		return ret
	}
	return *o.BusinessLineId
}

// GetBusinessLineIdOk returns a tuple with the BusinessLineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantResponse) GetBusinessLineIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BusinessLineId) {
		return nil, false
	}
	return o.BusinessLineId, true
}

// HasBusinessLineId returns a boolean if a field has been set.
func (o *CreateMerchantResponse) HasBusinessLineId() bool {
	if o != nil && !common.IsNil(o.BusinessLineId) {
		return true
	}

	return false
}

// SetBusinessLineId gets a reference to the given string and assigns it to the BusinessLineId field.
func (o *CreateMerchantResponse) SetBusinessLineId(v string) {
	o.BusinessLineId = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *CreateMerchantResponse) GetCompanyId() string {
	if o == nil || common.IsNil(o.CompanyId) {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantResponse) GetCompanyIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *CreateMerchantResponse) HasCompanyId() bool {
	if o != nil && !common.IsNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *CreateMerchantResponse) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateMerchantResponse) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateMerchantResponse) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateMerchantResponse) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateMerchantResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateMerchantResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateMerchantResponse) SetId(v string) {
	o.Id = &v
}

// GetLegalEntityId returns the LegalEntityId field value if set, zero value otherwise.
func (o *CreateMerchantResponse) GetLegalEntityId() string {
	if o == nil || common.IsNil(o.LegalEntityId) {
		var ret string
		return ret
	}
	return *o.LegalEntityId
}

// GetLegalEntityIdOk returns a tuple with the LegalEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantResponse) GetLegalEntityIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.LegalEntityId) {
		return nil, false
	}
	return o.LegalEntityId, true
}

// HasLegalEntityId returns a boolean if a field has been set.
func (o *CreateMerchantResponse) HasLegalEntityId() bool {
	if o != nil && !common.IsNil(o.LegalEntityId) {
		return true
	}

	return false
}

// SetLegalEntityId gets a reference to the given string and assigns it to the LegalEntityId field.
func (o *CreateMerchantResponse) SetLegalEntityId(v string) {
	o.LegalEntityId = &v
}

// GetPricingPlan returns the PricingPlan field value if set, zero value otherwise.
func (o *CreateMerchantResponse) GetPricingPlan() string {
	if o == nil || common.IsNil(o.PricingPlan) {
		var ret string
		return ret
	}
	return *o.PricingPlan
}

// GetPricingPlanOk returns a tuple with the PricingPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantResponse) GetPricingPlanOk() (*string, bool) {
	if o == nil || common.IsNil(o.PricingPlan) {
		return nil, false
	}
	return o.PricingPlan, true
}

// HasPricingPlan returns a boolean if a field has been set.
func (o *CreateMerchantResponse) HasPricingPlan() bool {
	if o != nil && !common.IsNil(o.PricingPlan) {
		return true
	}

	return false
}

// SetPricingPlan gets a reference to the given string and assigns it to the PricingPlan field.
func (o *CreateMerchantResponse) SetPricingPlan(v string) {
	o.PricingPlan = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *CreateMerchantResponse) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantResponse) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *CreateMerchantResponse) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *CreateMerchantResponse) SetReference(v string) {
	o.Reference = &v
}

func (o CreateMerchantResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMerchantResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BusinessLineId) {
		toSerialize["businessLineId"] = o.BusinessLineId
	}
	if !common.IsNil(o.CompanyId) {
		toSerialize["companyId"] = o.CompanyId
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
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
	return toSerialize, nil
}

type NullableCreateMerchantResponse struct {
	value *CreateMerchantResponse
	isSet bool
}

func (v NullableCreateMerchantResponse) Get() *CreateMerchantResponse {
	return v.value
}

func (v *NullableCreateMerchantResponse) Set(val *CreateMerchantResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMerchantResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMerchantResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMerchantResponse(val *CreateMerchantResponse) *NullableCreateMerchantResponse {
	return &NullableCreateMerchantResponse{value: val, isSet: true}
}

func (v NullableCreateMerchantResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMerchantResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
