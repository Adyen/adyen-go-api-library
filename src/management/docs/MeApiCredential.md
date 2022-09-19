# MeApiCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ApiCredentialLinks**](ApiCredentialLinks.md) |  | [optional] 
**Active** | **bool** | Indicates if the API credential is enabled. Must be set to **true** to use the credential in your integration. | 
**AllowedIpAddresses** | **[]string** | List of IP addresses from which your client can make requests.  If the list is empty, we allow requests from any IP. If the list is not empty and we get a request from an IP which is not on the list, you get a security error. | 
**AllowedOrigins** | Pointer to [**[]AllowedOrigin**](AllowedOrigin.md) | List containing the [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) linked to the API credential. | [optional] 
**AssociatedMerchantAccounts** | Pointer to **[]string** | List of merchant accounts that the API credential has access to. | [optional] 
**ClientKey** | **string** | Public key used for [client-side authentication](https://docs.adyen.com/development-resources/client-side-authentication). The client key is required for Drop-in and Components integrations. | 
**CompanyName** | Pointer to **string** | Name of the company linked to the API credential. | [optional] 
**Description** | Pointer to **string** | Description of the API credential. | [optional] 
**Id** | **string** | Unique identifier of the API credential. | 
**Roles** | **[]string** | List of [roles](https://docs.adyen.com/development-resources/api-credentials#roles-1) for the API credential. | 
**Username** | **string** | The name of the [API credential](https://docs.adyen.com/development-resources/api-credentials), for example **ws@Company.TestCompany**. | 

## Methods

### NewMeApiCredential

`func NewMeApiCredential(active bool, allowedIpAddresses []string, clientKey string, id string, roles []string, username string, ) *MeApiCredential`

NewMeApiCredential instantiates a new MeApiCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeApiCredentialWithDefaults

`func NewMeApiCredentialWithDefaults() *MeApiCredential`

NewMeApiCredentialWithDefaults instantiates a new MeApiCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *MeApiCredential) GetLinks() ApiCredentialLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MeApiCredential) GetLinksOk() (*ApiCredentialLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MeApiCredential) SetLinks(v ApiCredentialLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MeApiCredential) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetActive

`func (o *MeApiCredential) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MeApiCredential) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MeApiCredential) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAllowedIpAddresses

`func (o *MeApiCredential) GetAllowedIpAddresses() []string`

GetAllowedIpAddresses returns the AllowedIpAddresses field if non-nil, zero value otherwise.

### GetAllowedIpAddressesOk

`func (o *MeApiCredential) GetAllowedIpAddressesOk() (*[]string, bool)`

GetAllowedIpAddressesOk returns a tuple with the AllowedIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIpAddresses

`func (o *MeApiCredential) SetAllowedIpAddresses(v []string)`

SetAllowedIpAddresses sets AllowedIpAddresses field to given value.


### GetAllowedOrigins

`func (o *MeApiCredential) GetAllowedOrigins() []AllowedOrigin`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *MeApiCredential) GetAllowedOriginsOk() (*[]AllowedOrigin, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *MeApiCredential) SetAllowedOrigins(v []AllowedOrigin)`

SetAllowedOrigins sets AllowedOrigins field to given value.

### HasAllowedOrigins

`func (o *MeApiCredential) HasAllowedOrigins() bool`

HasAllowedOrigins returns a boolean if a field has been set.

### GetAssociatedMerchantAccounts

`func (o *MeApiCredential) GetAssociatedMerchantAccounts() []string`

GetAssociatedMerchantAccounts returns the AssociatedMerchantAccounts field if non-nil, zero value otherwise.

### GetAssociatedMerchantAccountsOk

`func (o *MeApiCredential) GetAssociatedMerchantAccountsOk() (*[]string, bool)`

GetAssociatedMerchantAccountsOk returns a tuple with the AssociatedMerchantAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedMerchantAccounts

`func (o *MeApiCredential) SetAssociatedMerchantAccounts(v []string)`

SetAssociatedMerchantAccounts sets AssociatedMerchantAccounts field to given value.

### HasAssociatedMerchantAccounts

`func (o *MeApiCredential) HasAssociatedMerchantAccounts() bool`

HasAssociatedMerchantAccounts returns a boolean if a field has been set.

### GetClientKey

`func (o *MeApiCredential) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *MeApiCredential) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *MeApiCredential) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.


### GetCompanyName

`func (o *MeApiCredential) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *MeApiCredential) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *MeApiCredential) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *MeApiCredential) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetDescription

`func (o *MeApiCredential) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MeApiCredential) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MeApiCredential) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MeApiCredential) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *MeApiCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeApiCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeApiCredential) SetId(v string)`

SetId sets Id field to given value.


### GetRoles

`func (o *MeApiCredential) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *MeApiCredential) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *MeApiCredential) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetUsername

`func (o *MeApiCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MeApiCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MeApiCredential) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


