# CreateCompanyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Your description for the company account, maximum 300 characters | [optional] 
**LegalEntityId** | Pointer to **string** | Legal Entity ID | [optional] 
**Reference** | Pointer to **string** | Your reference to the account | [optional] 

## Methods

### NewCreateCompanyRequest

`func NewCreateCompanyRequest() *CreateCompanyRequest`

NewCreateCompanyRequest instantiates a new CreateCompanyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCompanyRequestWithDefaults

`func NewCreateCompanyRequestWithDefaults() *CreateCompanyRequest`

NewCreateCompanyRequestWithDefaults instantiates a new CreateCompanyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateCompanyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCompanyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCompanyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCompanyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLegalEntityId

`func (o *CreateCompanyRequest) GetLegalEntityId() string`

GetLegalEntityId returns the LegalEntityId field if non-nil, zero value otherwise.

### GetLegalEntityIdOk

`func (o *CreateCompanyRequest) GetLegalEntityIdOk() (*string, bool)`

GetLegalEntityIdOk returns a tuple with the LegalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalEntityId

`func (o *CreateCompanyRequest) SetLegalEntityId(v string)`

SetLegalEntityId sets LegalEntityId field to given value.

### HasLegalEntityId

`func (o *CreateCompanyRequest) HasLegalEntityId() bool`

HasLegalEntityId returns a boolean if a field has been set.

### GetReference

`func (o *CreateCompanyRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CreateCompanyRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CreateCompanyRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CreateCompanyRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


