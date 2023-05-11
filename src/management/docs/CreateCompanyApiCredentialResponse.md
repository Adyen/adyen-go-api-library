# CreateCompanyApiCredentialResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ApiCredentialLinks**](ApiCredentialLinks.md) |  | [optional] 
**Active** | **bool** | Indicates if the API credential is enabled. Must be set to **true** to use the credential in your integration. | 
**AllowedIpAddresses** | **[]string** | List of IP addresses from which your client can make requests.  If the list is empty, we allow requests from any IP. If the list is not empty and we get a request from an IP which is not on the list, you get a security error. | 
**AllowedOrigins** | Pointer to [**[]AllowedOrigin**](AllowedOrigin.md) | List containing the [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) linked to the API credential. | [optional] 
**ApiKey** | **string** | The API key for the API credential that was created. | 
**AssociatedMerchantAccounts** | **[]string** | List of merchant accounts that the API credential has access to. | 
**ClientKey** | **string** | Public key used for [client-side authentication](https://docs.adyen.com/development-resources/client-side-authentication). The client key is required for Drop-in and Components integrations. | 
**Description** | Pointer to **string** | Description of the API credential. | [optional] 
**Id** | **string** | Unique identifier of the API credential. | 
**Password** | **string** | The password for the API credential that was created. | 
**Roles** | **[]string** | List of [roles](https://docs.adyen.com/development-resources/api-credentials#roles-1) for the API credential. | 
**Username** | **string** | The name of the [API credential](https://docs.adyen.com/development-resources/api-credentials), for example **ws@Company.TestCompany**. | 

## Methods

### NewCreateCompanyApiCredentialResponse

`func NewCreateCompanyApiCredentialResponse(active bool, allowedIpAddresses []string, apiKey string, associatedMerchantAccounts []string, clientKey string, id string, password string, roles []string, username string, ) *CreateCompanyApiCredentialResponse`

NewCreateCompanyApiCredentialResponse instantiates a new CreateCompanyApiCredentialResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCompanyApiCredentialResponseWithDefaults

`func NewCreateCompanyApiCredentialResponseWithDefaults() *CreateCompanyApiCredentialResponse`

NewCreateCompanyApiCredentialResponseWithDefaults instantiates a new CreateCompanyApiCredentialResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *CreateCompanyApiCredentialResponse) GetLinks() ApiCredentialLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateCompanyApiCredentialResponse) GetLinksOk() (*ApiCredentialLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateCompanyApiCredentialResponse) SetLinks(v ApiCredentialLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateCompanyApiCredentialResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetActive

`func (o *CreateCompanyApiCredentialResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateCompanyApiCredentialResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateCompanyApiCredentialResponse) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAllowedIpAddresses

`func (o *CreateCompanyApiCredentialResponse) GetAllowedIpAddresses() []string`

GetAllowedIpAddresses returns the AllowedIpAddresses field if non-nil, zero value otherwise.

### GetAllowedIpAddressesOk

`func (o *CreateCompanyApiCredentialResponse) GetAllowedIpAddressesOk() (*[]string, bool)`

GetAllowedIpAddressesOk returns a tuple with the AllowedIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIpAddresses

`func (o *CreateCompanyApiCredentialResponse) SetAllowedIpAddresses(v []string)`

SetAllowedIpAddresses sets AllowedIpAddresses field to given value.


### GetAllowedOrigins

`func (o *CreateCompanyApiCredentialResponse) GetAllowedOrigins() []AllowedOrigin`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *CreateCompanyApiCredentialResponse) GetAllowedOriginsOk() (*[]AllowedOrigin, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *CreateCompanyApiCredentialResponse) SetAllowedOrigins(v []AllowedOrigin)`

SetAllowedOrigins sets AllowedOrigins field to given value.

### HasAllowedOrigins

`func (o *CreateCompanyApiCredentialResponse) HasAllowedOrigins() bool`

HasAllowedOrigins returns a boolean if a field has been set.

### GetApiKey

`func (o *CreateCompanyApiCredentialResponse) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CreateCompanyApiCredentialResponse) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CreateCompanyApiCredentialResponse) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetAssociatedMerchantAccounts

`func (o *CreateCompanyApiCredentialResponse) GetAssociatedMerchantAccounts() []string`

GetAssociatedMerchantAccounts returns the AssociatedMerchantAccounts field if non-nil, zero value otherwise.

### GetAssociatedMerchantAccountsOk

`func (o *CreateCompanyApiCredentialResponse) GetAssociatedMerchantAccountsOk() (*[]string, bool)`

GetAssociatedMerchantAccountsOk returns a tuple with the AssociatedMerchantAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedMerchantAccounts

`func (o *CreateCompanyApiCredentialResponse) SetAssociatedMerchantAccounts(v []string)`

SetAssociatedMerchantAccounts sets AssociatedMerchantAccounts field to given value.


### GetClientKey

`func (o *CreateCompanyApiCredentialResponse) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *CreateCompanyApiCredentialResponse) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *CreateCompanyApiCredentialResponse) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.


### GetDescription

`func (o *CreateCompanyApiCredentialResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCompanyApiCredentialResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCompanyApiCredentialResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCompanyApiCredentialResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *CreateCompanyApiCredentialResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateCompanyApiCredentialResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateCompanyApiCredentialResponse) SetId(v string)`

SetId sets Id field to given value.


### GetPassword

`func (o *CreateCompanyApiCredentialResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateCompanyApiCredentialResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateCompanyApiCredentialResponse) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRoles

`func (o *CreateCompanyApiCredentialResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateCompanyApiCredentialResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateCompanyApiCredentialResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetUsername

`func (o *CreateCompanyApiCredentialResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateCompanyApiCredentialResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateCompanyApiCredentialResponse) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


