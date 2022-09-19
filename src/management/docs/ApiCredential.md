# ApiCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ApiCredentialLinks**](ApiCredentialLinks.md) |  | [optional] 
**Active** | **bool** | Indicates if the API credential is enabled. Must be set to **true** to use the credential in your integration. | 
**AllowedIpAddresses** | **[]string** | List of IP addresses from which your client can make requests.  If the list is empty, we allow requests from any IP. If the list is not empty and we get a request from an IP which is not on the list, you get a security error. | 
**AllowedOrigins** | Pointer to [**[]AllowedOrigin**](AllowedOrigin.md) | List containing the [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) linked to the API credential. | [optional] 
**ClientKey** | **string** | Public key used for [client-side authentication](https://docs.adyen.com/development-resources/client-side-authentication). The client key is required for Drop-in and Components integrations. | 
**Description** | Pointer to **string** | Description of the API credential. | [optional] 
**Id** | **string** | Unique identifier of the API credential. | 
**Roles** | **[]string** | List of [roles](https://docs.adyen.com/development-resources/api-credentials#roles-1) for the API credential. | 
**Username** | **string** | The name of the [API credential](https://docs.adyen.com/development-resources/api-credentials), for example **ws@Company.TestCompany**. | 

## Methods

### NewApiCredential

`func NewApiCredential(active bool, allowedIpAddresses []string, clientKey string, id string, roles []string, username string, ) *ApiCredential`

NewApiCredential instantiates a new ApiCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCredentialWithDefaults

`func NewApiCredentialWithDefaults() *ApiCredential`

NewApiCredentialWithDefaults instantiates a new ApiCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ApiCredential) GetLinks() ApiCredentialLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiCredential) GetLinksOk() (*ApiCredentialLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiCredential) SetLinks(v ApiCredentialLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiCredential) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetActive

`func (o *ApiCredential) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiCredential) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiCredential) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAllowedIpAddresses

`func (o *ApiCredential) GetAllowedIpAddresses() []string`

GetAllowedIpAddresses returns the AllowedIpAddresses field if non-nil, zero value otherwise.

### GetAllowedIpAddressesOk

`func (o *ApiCredential) GetAllowedIpAddressesOk() (*[]string, bool)`

GetAllowedIpAddressesOk returns a tuple with the AllowedIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIpAddresses

`func (o *ApiCredential) SetAllowedIpAddresses(v []string)`

SetAllowedIpAddresses sets AllowedIpAddresses field to given value.


### GetAllowedOrigins

`func (o *ApiCredential) GetAllowedOrigins() []AllowedOrigin`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *ApiCredential) GetAllowedOriginsOk() (*[]AllowedOrigin, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *ApiCredential) SetAllowedOrigins(v []AllowedOrigin)`

SetAllowedOrigins sets AllowedOrigins field to given value.

### HasAllowedOrigins

`func (o *ApiCredential) HasAllowedOrigins() bool`

HasAllowedOrigins returns a boolean if a field has been set.

### GetClientKey

`func (o *ApiCredential) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *ApiCredential) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *ApiCredential) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.


### GetDescription

`func (o *ApiCredential) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiCredential) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiCredential) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiCredential) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ApiCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiCredential) SetId(v string)`

SetId sets Id field to given value.


### GetRoles

`func (o *ApiCredential) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiCredential) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiCredential) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetUsername

`func (o *ApiCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiCredential) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


