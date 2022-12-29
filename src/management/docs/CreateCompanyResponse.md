# CreateCompanyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** | Parent company ID | [optional] 
**Description** | Pointer to **string** | Your description for the company account, maximum 300 characters | [optional] 
**LegalEntityId** | Pointer to **string** | Legal Entity ID | [optional] 
**Reference** | Pointer to **string** | Your reference for the account. | [optional] 

## Methods

### NewCreateCompanyResponse

`func NewCreateCompanyResponse() *CreateCompanyResponse`

NewCreateCompanyResponse instantiates a new CreateCompanyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCompanyResponseWithDefaults

`func NewCreateCompanyResponseWithDefaults() *CreateCompanyResponse`

NewCreateCompanyResponseWithDefaults instantiates a new CreateCompanyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CreateCompanyResponse) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CreateCompanyResponse) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CreateCompanyResponse) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CreateCompanyResponse) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetDescription

`func (o *CreateCompanyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCompanyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCompanyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCompanyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLegalEntityId

`func (o *CreateCompanyResponse) GetLegalEntityId() string`

GetLegalEntityId returns the LegalEntityId field if non-nil, zero value otherwise.

### GetLegalEntityIdOk

`func (o *CreateCompanyResponse) GetLegalEntityIdOk() (*string, bool)`

GetLegalEntityIdOk returns a tuple with the LegalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalEntityId

`func (o *CreateCompanyResponse) SetLegalEntityId(v string)`

SetLegalEntityId sets LegalEntityId field to given value.

### HasLegalEntityId

`func (o *CreateCompanyResponse) HasLegalEntityId() bool`

HasLegalEntityId returns a boolean if a field has been set.

### GetReference

`func (o *CreateCompanyResponse) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CreateCompanyResponse) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CreateCompanyResponse) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CreateCompanyResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


