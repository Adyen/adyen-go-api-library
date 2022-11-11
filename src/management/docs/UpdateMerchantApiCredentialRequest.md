# UpdateMerchantApiCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates if the API credential is enabled. | [optional] 
**AllowedOrigins** | Pointer to **[]string** | The new list of [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) for the API credential. | [optional] 
**Description** | Pointer to **string** | Description of the API credential. | [optional] 
**Roles** | Pointer to **[]string** | List of [roles](https://docs.adyen.com/development-resources/api-credentials#roles-1) for the API credential. | [optional] 

## Methods

### NewUpdateMerchantApiCredentialRequest

`func NewUpdateMerchantApiCredentialRequest() *UpdateMerchantApiCredentialRequest`

NewUpdateMerchantApiCredentialRequest instantiates a new UpdateMerchantApiCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMerchantApiCredentialRequestWithDefaults

`func NewUpdateMerchantApiCredentialRequestWithDefaults() *UpdateMerchantApiCredentialRequest`

NewUpdateMerchantApiCredentialRequestWithDefaults instantiates a new UpdateMerchantApiCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdateMerchantApiCredentialRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateMerchantApiCredentialRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateMerchantApiCredentialRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateMerchantApiCredentialRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAllowedOrigins

`func (o *UpdateMerchantApiCredentialRequest) GetAllowedOrigins() []string`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *UpdateMerchantApiCredentialRequest) GetAllowedOriginsOk() (*[]string, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *UpdateMerchantApiCredentialRequest) SetAllowedOrigins(v []string)`

SetAllowedOrigins sets AllowedOrigins field to given value.

### HasAllowedOrigins

`func (o *UpdateMerchantApiCredentialRequest) HasAllowedOrigins() bool`

HasAllowedOrigins returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateMerchantApiCredentialRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateMerchantApiCredentialRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateMerchantApiCredentialRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateMerchantApiCredentialRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRoles

`func (o *UpdateMerchantApiCredentialRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateMerchantApiCredentialRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateMerchantApiCredentialRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateMerchantApiCredentialRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


