# CreateCompanyApiCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedOrigins** | Pointer to **[]string** | List of [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) for the new API credential. | [optional] 
**AssociatedMerchantAccounts** | Pointer to **[]string** | List of merchant accounts that the API credential has access to. | [optional] 
**Description** | Pointer to **string** | Description of the API credential. | [optional] 
**Roles** | Pointer to **[]string** | List of [roles](https://docs.adyen.com/development-resources/api-credentials#roles-1) of the API credential. | [optional] 

## Methods

### NewCreateCompanyApiCredentialRequest

`func NewCreateCompanyApiCredentialRequest() *CreateCompanyApiCredentialRequest`

NewCreateCompanyApiCredentialRequest instantiates a new CreateCompanyApiCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCompanyApiCredentialRequestWithDefaults

`func NewCreateCompanyApiCredentialRequestWithDefaults() *CreateCompanyApiCredentialRequest`

NewCreateCompanyApiCredentialRequestWithDefaults instantiates a new CreateCompanyApiCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedOrigins

`func (o *CreateCompanyApiCredentialRequest) GetAllowedOrigins() []string`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *CreateCompanyApiCredentialRequest) GetAllowedOriginsOk() (*[]string, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *CreateCompanyApiCredentialRequest) SetAllowedOrigins(v []string)`

SetAllowedOrigins sets AllowedOrigins field to given value.

### HasAllowedOrigins

`func (o *CreateCompanyApiCredentialRequest) HasAllowedOrigins() bool`

HasAllowedOrigins returns a boolean if a field has been set.

### GetAssociatedMerchantAccounts

`func (o *CreateCompanyApiCredentialRequest) GetAssociatedMerchantAccounts() []string`

GetAssociatedMerchantAccounts returns the AssociatedMerchantAccounts field if non-nil, zero value otherwise.

### GetAssociatedMerchantAccountsOk

`func (o *CreateCompanyApiCredentialRequest) GetAssociatedMerchantAccountsOk() (*[]string, bool)`

GetAssociatedMerchantAccountsOk returns a tuple with the AssociatedMerchantAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedMerchantAccounts

`func (o *CreateCompanyApiCredentialRequest) SetAssociatedMerchantAccounts(v []string)`

SetAssociatedMerchantAccounts sets AssociatedMerchantAccounts field to given value.

### HasAssociatedMerchantAccounts

`func (o *CreateCompanyApiCredentialRequest) HasAssociatedMerchantAccounts() bool`

HasAssociatedMerchantAccounts returns a boolean if a field has been set.

### GetDescription

`func (o *CreateCompanyApiCredentialRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCompanyApiCredentialRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCompanyApiCredentialRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCompanyApiCredentialRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRoles

`func (o *CreateCompanyApiCredentialRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateCompanyApiCredentialRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateCompanyApiCredentialRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateCompanyApiCredentialRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


