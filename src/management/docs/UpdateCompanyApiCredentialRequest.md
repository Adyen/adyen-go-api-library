# UpdateCompanyApiCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates if the API credential is enabled. | [optional] 
**AllowedOrigins** | Pointer to **[]string** | The new list of [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) for the API credential. | [optional] 
**AssociatedMerchantAccounts** | Pointer to **[]string** | List of merchant accounts that the API credential has access to. | [optional] 
**Description** | Pointer to **string** | Description of the API credential. | [optional] 
**Roles** | Pointer to **[]string** | List of [roles](https://docs.adyen.com/development-resources/api-credentials#roles-1) of the API credential. | [optional] 

## Methods

### NewUpdateCompanyApiCredentialRequest

`func NewUpdateCompanyApiCredentialRequest() *UpdateCompanyApiCredentialRequest`

NewUpdateCompanyApiCredentialRequest instantiates a new UpdateCompanyApiCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCompanyApiCredentialRequestWithDefaults

`func NewUpdateCompanyApiCredentialRequestWithDefaults() *UpdateCompanyApiCredentialRequest`

NewUpdateCompanyApiCredentialRequestWithDefaults instantiates a new UpdateCompanyApiCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdateCompanyApiCredentialRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateCompanyApiCredentialRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateCompanyApiCredentialRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateCompanyApiCredentialRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAllowedOrigins

`func (o *UpdateCompanyApiCredentialRequest) GetAllowedOrigins() []string`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *UpdateCompanyApiCredentialRequest) GetAllowedOriginsOk() (*[]string, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *UpdateCompanyApiCredentialRequest) SetAllowedOrigins(v []string)`

SetAllowedOrigins sets AllowedOrigins field to given value.

### HasAllowedOrigins

`func (o *UpdateCompanyApiCredentialRequest) HasAllowedOrigins() bool`

HasAllowedOrigins returns a boolean if a field has been set.

### GetAssociatedMerchantAccounts

`func (o *UpdateCompanyApiCredentialRequest) GetAssociatedMerchantAccounts() []string`

GetAssociatedMerchantAccounts returns the AssociatedMerchantAccounts field if non-nil, zero value otherwise.

### GetAssociatedMerchantAccountsOk

`func (o *UpdateCompanyApiCredentialRequest) GetAssociatedMerchantAccountsOk() (*[]string, bool)`

GetAssociatedMerchantAccountsOk returns a tuple with the AssociatedMerchantAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedMerchantAccounts

`func (o *UpdateCompanyApiCredentialRequest) SetAssociatedMerchantAccounts(v []string)`

SetAssociatedMerchantAccounts sets AssociatedMerchantAccounts field to given value.

### HasAssociatedMerchantAccounts

`func (o *UpdateCompanyApiCredentialRequest) HasAssociatedMerchantAccounts() bool`

HasAssociatedMerchantAccounts returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateCompanyApiCredentialRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCompanyApiCredentialRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCompanyApiCredentialRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCompanyApiCredentialRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRoles

`func (o *UpdateCompanyApiCredentialRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateCompanyApiCredentialRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateCompanyApiCredentialRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateCompanyApiCredentialRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


