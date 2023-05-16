# CreateMerchantApiCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedOrigins** | Pointer to **[]string** | The list of [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) for the new API credential. | [optional] 
**Description** | Pointer to **string** | Description of the API credential. | [optional] 
**Roles** | Pointer to **[]string** | List of [roles](https://docs.adyen.com/development-resources/api-credentials#roles-1) for the API credential. | [optional] 

## Methods

### NewCreateMerchantApiCredentialRequest

`func NewCreateMerchantApiCredentialRequest() *CreateMerchantApiCredentialRequest`

NewCreateMerchantApiCredentialRequest instantiates a new CreateMerchantApiCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMerchantApiCredentialRequestWithDefaults

`func NewCreateMerchantApiCredentialRequestWithDefaults() *CreateMerchantApiCredentialRequest`

NewCreateMerchantApiCredentialRequestWithDefaults instantiates a new CreateMerchantApiCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedOrigins

`func (o *CreateMerchantApiCredentialRequest) GetAllowedOrigins() []string`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *CreateMerchantApiCredentialRequest) GetAllowedOriginsOk() (*[]string, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *CreateMerchantApiCredentialRequest) SetAllowedOrigins(v []string)`

SetAllowedOrigins sets AllowedOrigins field to given value.

### HasAllowedOrigins

`func (o *CreateMerchantApiCredentialRequest) HasAllowedOrigins() bool`

HasAllowedOrigins returns a boolean if a field has been set.

### GetDescription

`func (o *CreateMerchantApiCredentialRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMerchantApiCredentialRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMerchantApiCredentialRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateMerchantApiCredentialRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRoles

`func (o *CreateMerchantApiCredentialRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateMerchantApiCredentialRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateMerchantApiCredentialRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateMerchantApiCredentialRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


