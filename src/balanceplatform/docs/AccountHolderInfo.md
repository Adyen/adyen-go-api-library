# AccountHolderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalancePlatform** | Pointer to **string** | The unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balancePlatforms/{id}__queryParam_id) to which the account holder belongs. Required in the request if your API credentials can be used for multiple balance platforms. | [optional] 
**Capabilities** | Pointer to [**map[string]AccountHolderCapability**](AccountHolderCapability.md) | Contains key-value pairs that specify the actions that an account holder can do in your platform. The key is a capability required for your integration. For example, **issueCard** for Issuing. The value is an object containing the settings for the capability. | [optional] 
**ContactDetails** | Pointer to [**ContactDetails**](ContactDetails.md) |  | [optional] 
**Description** | Pointer to **string** | Your description for the account holder, maximum 300 characters. | [optional] 
**LegalEntityId** | **string** | The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/legalentity/latest/post/legalEntities#responses-200-id) associated with the account holder. Adyen performs a verification process against the legal entity of the account holder. | 
**Reference** | Pointer to **string** | Your reference for the account holder, maximum 150 characters. | [optional] 
**TimeZone** | Pointer to **string** | The [time zone](https://www.iana.org/time-zones) of the account holder. For example, **Europe/Amsterdam**. Defaults to the time zone of the balance platform if no time zone is set. For possible values, see the [list of time zone codes](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). | [optional] 

## Methods

### NewAccountHolderInfo

`func NewAccountHolderInfo(legalEntityId string, ) *AccountHolderInfo`

NewAccountHolderInfo instantiates a new AccountHolderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountHolderInfoWithDefaults

`func NewAccountHolderInfoWithDefaults() *AccountHolderInfo`

NewAccountHolderInfoWithDefaults instantiates a new AccountHolderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalancePlatform

`func (o *AccountHolderInfo) GetBalancePlatform() string`

GetBalancePlatform returns the BalancePlatform field if non-nil, zero value otherwise.

### GetBalancePlatformOk

`func (o *AccountHolderInfo) GetBalancePlatformOk() (*string, bool)`

GetBalancePlatformOk returns a tuple with the BalancePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancePlatform

`func (o *AccountHolderInfo) SetBalancePlatform(v string)`

SetBalancePlatform sets BalancePlatform field to given value.

### HasBalancePlatform

`func (o *AccountHolderInfo) HasBalancePlatform() bool`

HasBalancePlatform returns a boolean if a field has been set.

### GetCapabilities

`func (o *AccountHolderInfo) GetCapabilities() map[string]AccountHolderCapability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AccountHolderInfo) GetCapabilitiesOk() (*map[string]AccountHolderCapability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AccountHolderInfo) SetCapabilities(v map[string]AccountHolderCapability)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *AccountHolderInfo) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetContactDetails

`func (o *AccountHolderInfo) GetContactDetails() ContactDetails`

GetContactDetails returns the ContactDetails field if non-nil, zero value otherwise.

### GetContactDetailsOk

`func (o *AccountHolderInfo) GetContactDetailsOk() (*ContactDetails, bool)`

GetContactDetailsOk returns a tuple with the ContactDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactDetails

`func (o *AccountHolderInfo) SetContactDetails(v ContactDetails)`

SetContactDetails sets ContactDetails field to given value.

### HasContactDetails

`func (o *AccountHolderInfo) HasContactDetails() bool`

HasContactDetails returns a boolean if a field has been set.

### GetDescription

`func (o *AccountHolderInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountHolderInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountHolderInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountHolderInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLegalEntityId

`func (o *AccountHolderInfo) GetLegalEntityId() string`

GetLegalEntityId returns the LegalEntityId field if non-nil, zero value otherwise.

### GetLegalEntityIdOk

`func (o *AccountHolderInfo) GetLegalEntityIdOk() (*string, bool)`

GetLegalEntityIdOk returns a tuple with the LegalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalEntityId

`func (o *AccountHolderInfo) SetLegalEntityId(v string)`

SetLegalEntityId sets LegalEntityId field to given value.


### GetReference

`func (o *AccountHolderInfo) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AccountHolderInfo) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AccountHolderInfo) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AccountHolderInfo) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetTimeZone

`func (o *AccountHolderInfo) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *AccountHolderInfo) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *AccountHolderInfo) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *AccountHolderInfo) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


