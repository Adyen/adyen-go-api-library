# AccountHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalancePlatform** | Pointer to **string** | The unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balancePlatforms/{id}__queryParam_id) to which the account holder belongs. Required in the request if your API credentials can be used for multiple balance platforms. | [optional] 
**Capabilities** | Pointer to [**map[string]AccountHolderCapability**](AccountHolderCapability.md) | Contains key-value pairs that specify the actions that an account holder can do in your platform. The key is a capability required for your integration. For example, **issueCard** for Issuing. The value is an object containing the settings for the capability. | [optional] 
**ContactDetails** | Pointer to [**ContactDetails**](ContactDetails.md) |  | [optional] 
**Description** | Pointer to **string** | Your description for the account holder, maximum 300 characters. | [optional] 
**Id** | **string** | The unique identifier of the account holder. | [readonly] 
**LegalEntityId** | **string** | The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/legalentity/latest/post/legalEntities#responses-200-id) associated with the account holder. Adyen performs a verification process against the legal entity of the account holder. | 
**PrimaryBalanceAccount** | Pointer to **string** | The ID of the account holder&#39;s primary balance account. By default, this is set to the first balance account that you create for the account holder. To assign a different balance account, send a PATCH request. | [optional] 
**Reference** | Pointer to **string** | Your reference for the account holder, maximum 150 characters. | [optional] 
**Status** | Pointer to **string** | The status of the account holder.  Possible values:    * **active**: The account holder is active. This is the default status when creating an account holder.    * **inactive (Deprecated)**: The account holder is temporarily inactive due to missing KYC details. You can set the account back to active by providing the missing KYC details.    * **suspended**: The account holder is permanently deactivated by Adyen. This action cannot be undone.   * **closed**: The account holder is permanently deactivated by you. This action cannot be undone. | [optional] 
**TimeZone** | Pointer to **string** | The time zone of the account holder. For example, **Europe/Amsterdam**. Defaults to the time zone of the balance platform if no time zone is set. For possible values, see the [list of time zone codes](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). | [optional] 
**VerificationDeadlines** | Pointer to [**[]VerificationDeadline**](VerificationDeadline.md) | List of verification deadlines and the capabilities that will be disallowed if verification errors are not resolved. | [optional] [readonly] 

## Methods

### NewAccountHolder

`func NewAccountHolder(id string, legalEntityId string, ) *AccountHolder`

NewAccountHolder instantiates a new AccountHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountHolderWithDefaults

`func NewAccountHolderWithDefaults() *AccountHolder`

NewAccountHolderWithDefaults instantiates a new AccountHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalancePlatform

`func (o *AccountHolder) GetBalancePlatform() string`

GetBalancePlatform returns the BalancePlatform field if non-nil, zero value otherwise.

### GetBalancePlatformOk

`func (o *AccountHolder) GetBalancePlatformOk() (*string, bool)`

GetBalancePlatformOk returns a tuple with the BalancePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancePlatform

`func (o *AccountHolder) SetBalancePlatform(v string)`

SetBalancePlatform sets BalancePlatform field to given value.

### HasBalancePlatform

`func (o *AccountHolder) HasBalancePlatform() bool`

HasBalancePlatform returns a boolean if a field has been set.

### GetCapabilities

`func (o *AccountHolder) GetCapabilities() map[string]AccountHolderCapability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AccountHolder) GetCapabilitiesOk() (*map[string]AccountHolderCapability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AccountHolder) SetCapabilities(v map[string]AccountHolderCapability)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *AccountHolder) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetContactDetails

`func (o *AccountHolder) GetContactDetails() ContactDetails`

GetContactDetails returns the ContactDetails field if non-nil, zero value otherwise.

### GetContactDetailsOk

`func (o *AccountHolder) GetContactDetailsOk() (*ContactDetails, bool)`

GetContactDetailsOk returns a tuple with the ContactDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactDetails

`func (o *AccountHolder) SetContactDetails(v ContactDetails)`

SetContactDetails sets ContactDetails field to given value.

### HasContactDetails

`func (o *AccountHolder) HasContactDetails() bool`

HasContactDetails returns a boolean if a field has been set.

### GetDescription

`func (o *AccountHolder) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountHolder) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountHolder) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountHolder) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AccountHolder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountHolder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountHolder) SetId(v string)`

SetId sets Id field to given value.


### GetLegalEntityId

`func (o *AccountHolder) GetLegalEntityId() string`

GetLegalEntityId returns the LegalEntityId field if non-nil, zero value otherwise.

### GetLegalEntityIdOk

`func (o *AccountHolder) GetLegalEntityIdOk() (*string, bool)`

GetLegalEntityIdOk returns a tuple with the LegalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalEntityId

`func (o *AccountHolder) SetLegalEntityId(v string)`

SetLegalEntityId sets LegalEntityId field to given value.


### GetPrimaryBalanceAccount

`func (o *AccountHolder) GetPrimaryBalanceAccount() string`

GetPrimaryBalanceAccount returns the PrimaryBalanceAccount field if non-nil, zero value otherwise.

### GetPrimaryBalanceAccountOk

`func (o *AccountHolder) GetPrimaryBalanceAccountOk() (*string, bool)`

GetPrimaryBalanceAccountOk returns a tuple with the PrimaryBalanceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryBalanceAccount

`func (o *AccountHolder) SetPrimaryBalanceAccount(v string)`

SetPrimaryBalanceAccount sets PrimaryBalanceAccount field to given value.

### HasPrimaryBalanceAccount

`func (o *AccountHolder) HasPrimaryBalanceAccount() bool`

HasPrimaryBalanceAccount returns a boolean if a field has been set.

### GetReference

`func (o *AccountHolder) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AccountHolder) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AccountHolder) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AccountHolder) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetStatus

`func (o *AccountHolder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountHolder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountHolder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountHolder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeZone

`func (o *AccountHolder) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *AccountHolder) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *AccountHolder) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *AccountHolder) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetVerificationDeadlines

`func (o *AccountHolder) GetVerificationDeadlines() []VerificationDeadline`

GetVerificationDeadlines returns the VerificationDeadlines field if non-nil, zero value otherwise.

### GetVerificationDeadlinesOk

`func (o *AccountHolder) GetVerificationDeadlinesOk() (*[]VerificationDeadline, bool)`

GetVerificationDeadlinesOk returns a tuple with the VerificationDeadlines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationDeadlines

`func (o *AccountHolder) SetVerificationDeadlines(v []VerificationDeadline)`

SetVerificationDeadlines sets VerificationDeadlines field to given value.

### HasVerificationDeadlines

`func (o *AccountHolder) HasVerificationDeadlines() bool`

HasVerificationDeadlines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


