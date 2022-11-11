# CreateMerchantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessLineId** | Pointer to **string** | The unique identifier of the [business line](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/businessLines). Required for an Adyen for Platforms Manage integration. | [optional] 
**CompanyId** | **string** | The unique identifier of the company account. | 
**Description** | Pointer to **string** | Your description for the merchant account, maximum 300 characters. | [optional] 
**LegalEntityId** | Pointer to **string** | The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities). Required for an Adyen for Platforms Manage integration. | [optional] 
**PricingPlan** | Pointer to **string** | Sets the pricing plan for the merchant account. Required for an Adyen for Platforms Manage integration. Your Adyen contact will provide the values that you can use. | [optional] 
**Reference** | Pointer to **string** | Your reference for the merchant account. To make this reference the unique identifier of the merchant account, your Adyen contact can set up a template on your company account. The template can have 6 to 255 characters with upper- and lower-case letters, underscores, and numbers. When your company account has a template, then the &#x60;reference&#x60; is required and must be unique within the company account. | [optional] 

## Methods

### NewCreateMerchantRequest

`func NewCreateMerchantRequest(companyId string, ) *CreateMerchantRequest`

NewCreateMerchantRequest instantiates a new CreateMerchantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMerchantRequestWithDefaults

`func NewCreateMerchantRequestWithDefaults() *CreateMerchantRequest`

NewCreateMerchantRequestWithDefaults instantiates a new CreateMerchantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessLineId

`func (o *CreateMerchantRequest) GetBusinessLineId() string`

GetBusinessLineId returns the BusinessLineId field if non-nil, zero value otherwise.

### GetBusinessLineIdOk

`func (o *CreateMerchantRequest) GetBusinessLineIdOk() (*string, bool)`

GetBusinessLineIdOk returns a tuple with the BusinessLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessLineId

`func (o *CreateMerchantRequest) SetBusinessLineId(v string)`

SetBusinessLineId sets BusinessLineId field to given value.

### HasBusinessLineId

`func (o *CreateMerchantRequest) HasBusinessLineId() bool`

HasBusinessLineId returns a boolean if a field has been set.

### GetCompanyId

`func (o *CreateMerchantRequest) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CreateMerchantRequest) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CreateMerchantRequest) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetDescription

`func (o *CreateMerchantRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMerchantRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMerchantRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateMerchantRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLegalEntityId

`func (o *CreateMerchantRequest) GetLegalEntityId() string`

GetLegalEntityId returns the LegalEntityId field if non-nil, zero value otherwise.

### GetLegalEntityIdOk

`func (o *CreateMerchantRequest) GetLegalEntityIdOk() (*string, bool)`

GetLegalEntityIdOk returns a tuple with the LegalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalEntityId

`func (o *CreateMerchantRequest) SetLegalEntityId(v string)`

SetLegalEntityId sets LegalEntityId field to given value.

### HasLegalEntityId

`func (o *CreateMerchantRequest) HasLegalEntityId() bool`

HasLegalEntityId returns a boolean if a field has been set.

### GetPricingPlan

`func (o *CreateMerchantRequest) GetPricingPlan() string`

GetPricingPlan returns the PricingPlan field if non-nil, zero value otherwise.

### GetPricingPlanOk

`func (o *CreateMerchantRequest) GetPricingPlanOk() (*string, bool)`

GetPricingPlanOk returns a tuple with the PricingPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPlan

`func (o *CreateMerchantRequest) SetPricingPlan(v string)`

SetPricingPlan sets PricingPlan field to given value.

### HasPricingPlan

`func (o *CreateMerchantRequest) HasPricingPlan() bool`

HasPricingPlan returns a boolean if a field has been set.

### GetReference

`func (o *CreateMerchantRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CreateMerchantRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CreateMerchantRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CreateMerchantRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


