# AccountHolderCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | Pointer to **bool** | Indicates whether the capability is allowed. Adyen sets this to **true** if the verification is successful and the account holder is permitted to use the capability. | [optional] 
**AllowedLevel** | Pointer to **string** | The capability level that is allowed for the account holder.  Possible values: **notApplicable**, **low**, **medium**, **high**. | [optional] 
**AllowedSettings** | Pointer to [**CapabilitySettings**](CapabilitySettings.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the capability is enabled. If **false**, the capability is temporarily disabled for the account holder. | [optional] 
**Problems** | Pointer to **[]map[string]interface{}** | Contains verification errors and the actions that you can take to resolve them. | [optional] 
**Requested** | Pointer to **bool** | Indicates whether the capability is requested. To check whether the account holder is permitted to use the capability, refer to the &#x60;allowed&#x60; field. | [optional] 
**RequestedLevel** | Pointer to **string** | The requested level of the capability. Some capabilities, such as those used in [card issuing](https://docs.adyen.com/issuing/add-capabilities#capability-levels), have different levels. Levels increase the capability, but also require additional checks and increased monitoring.  Possible values: **notApplicable**, **low**, **medium**, **high**. | [optional] 
**RequestedSettings** | Pointer to [**CapabilitySettings**](CapabilitySettings.md) |  | [optional] 
**TransferInstruments** | Pointer to [**[]AccountSupportingEntityCapability**](AccountSupportingEntityCapability.md) | Contains the status of the transfer instruments associated with this capability.  | [optional] 
**VerificationStatus** | Pointer to **string** | The status of the verification checks for the capability.  Possible values:  * **pending**: Adyen is running the verification.  * **invalid**: The verification failed. Check if the &#x60;errors&#x60; array contains more information.  * **valid**: The verification has been successfully completed.  * **rejected**: Adyen has verified the information, but found reasons to not allow the capability.  | [optional] 

## Methods

### NewAccountHolderCapability

`func NewAccountHolderCapability() *AccountHolderCapability`

NewAccountHolderCapability instantiates a new AccountHolderCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountHolderCapabilityWithDefaults

`func NewAccountHolderCapabilityWithDefaults() *AccountHolderCapability`

NewAccountHolderCapabilityWithDefaults instantiates a new AccountHolderCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *AccountHolderCapability) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *AccountHolderCapability) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *AccountHolderCapability) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *AccountHolderCapability) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.

### GetAllowedLevel

`func (o *AccountHolderCapability) GetAllowedLevel() string`

GetAllowedLevel returns the AllowedLevel field if non-nil, zero value otherwise.

### GetAllowedLevelOk

`func (o *AccountHolderCapability) GetAllowedLevelOk() (*string, bool)`

GetAllowedLevelOk returns a tuple with the AllowedLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedLevel

`func (o *AccountHolderCapability) SetAllowedLevel(v string)`

SetAllowedLevel sets AllowedLevel field to given value.

### HasAllowedLevel

`func (o *AccountHolderCapability) HasAllowedLevel() bool`

HasAllowedLevel returns a boolean if a field has been set.

### GetAllowedSettings

`func (o *AccountHolderCapability) GetAllowedSettings() CapabilitySettings`

GetAllowedSettings returns the AllowedSettings field if non-nil, zero value otherwise.

### GetAllowedSettingsOk

`func (o *AccountHolderCapability) GetAllowedSettingsOk() (*CapabilitySettings, bool)`

GetAllowedSettingsOk returns a tuple with the AllowedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSettings

`func (o *AccountHolderCapability) SetAllowedSettings(v CapabilitySettings)`

SetAllowedSettings sets AllowedSettings field to given value.

### HasAllowedSettings

`func (o *AccountHolderCapability) HasAllowedSettings() bool`

HasAllowedSettings returns a boolean if a field has been set.

### GetEnabled

`func (o *AccountHolderCapability) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AccountHolderCapability) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AccountHolderCapability) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AccountHolderCapability) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProblems

`func (o *AccountHolderCapability) GetProblems() []map[string]interface{}`

GetProblems returns the Problems field if non-nil, zero value otherwise.

### GetProblemsOk

`func (o *AccountHolderCapability) GetProblemsOk() (*[]map[string]interface{}, bool)`

GetProblemsOk returns a tuple with the Problems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblems

`func (o *AccountHolderCapability) SetProblems(v []map[string]interface{})`

SetProblems sets Problems field to given value.

### HasProblems

`func (o *AccountHolderCapability) HasProblems() bool`

HasProblems returns a boolean if a field has been set.

### GetRequested

`func (o *AccountHolderCapability) GetRequested() bool`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *AccountHolderCapability) GetRequestedOk() (*bool, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *AccountHolderCapability) SetRequested(v bool)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *AccountHolderCapability) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### GetRequestedLevel

`func (o *AccountHolderCapability) GetRequestedLevel() string`

GetRequestedLevel returns the RequestedLevel field if non-nil, zero value otherwise.

### GetRequestedLevelOk

`func (o *AccountHolderCapability) GetRequestedLevelOk() (*string, bool)`

GetRequestedLevelOk returns a tuple with the RequestedLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedLevel

`func (o *AccountHolderCapability) SetRequestedLevel(v string)`

SetRequestedLevel sets RequestedLevel field to given value.

### HasRequestedLevel

`func (o *AccountHolderCapability) HasRequestedLevel() bool`

HasRequestedLevel returns a boolean if a field has been set.

### GetRequestedSettings

`func (o *AccountHolderCapability) GetRequestedSettings() CapabilitySettings`

GetRequestedSettings returns the RequestedSettings field if non-nil, zero value otherwise.

### GetRequestedSettingsOk

`func (o *AccountHolderCapability) GetRequestedSettingsOk() (*CapabilitySettings, bool)`

GetRequestedSettingsOk returns a tuple with the RequestedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedSettings

`func (o *AccountHolderCapability) SetRequestedSettings(v CapabilitySettings)`

SetRequestedSettings sets RequestedSettings field to given value.

### HasRequestedSettings

`func (o *AccountHolderCapability) HasRequestedSettings() bool`

HasRequestedSettings returns a boolean if a field has been set.

### GetTransferInstruments

`func (o *AccountHolderCapability) GetTransferInstruments() []AccountSupportingEntityCapability`

GetTransferInstruments returns the TransferInstruments field if non-nil, zero value otherwise.

### GetTransferInstrumentsOk

`func (o *AccountHolderCapability) GetTransferInstrumentsOk() (*[]AccountSupportingEntityCapability, bool)`

GetTransferInstrumentsOk returns a tuple with the TransferInstruments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferInstruments

`func (o *AccountHolderCapability) SetTransferInstruments(v []AccountSupportingEntityCapability)`

SetTransferInstruments sets TransferInstruments field to given value.

### HasTransferInstruments

`func (o *AccountHolderCapability) HasTransferInstruments() bool`

HasTransferInstruments returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *AccountHolderCapability) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *AccountHolderCapability) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *AccountHolderCapability) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *AccountHolderCapability) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


