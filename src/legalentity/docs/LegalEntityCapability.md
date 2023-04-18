# LegalEntityCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | Pointer to **bool** | Indicates whether the capability is allowed. Adyen sets this to **true** if the verification is successful  | [optional] [readonly] 
**AllowedLevel** | Pointer to **string** | The capability level that is allowed for the legal entity.  Possible values: **notApplicable**, **low**, **medium**, **high**. | [optional] [readonly] 
**AllowedSettings** | Pointer to [**CapabilitySettings**](CapabilitySettings.md) |  | [optional] 
**Requested** | Pointer to **bool** | Indicates whether the capability is requested. To check whether the Legal Entity is permitted to use the capability,  | [optional] [readonly] 
**RequestedLevel** | Pointer to **string** | The requested level of the capability. Some capabilities, such as those used in [card issuing](https://docs.adyen.com/issuing/add-capabilities#capability-levels), have different levels. Levels increase the capability, but also require additional checks and increased monitoring.  Possible values: **notApplicable**, **low**, **medium**, **high**. | [optional] [readonly] 
**RequestedSettings** | Pointer to [**CapabilitySettings**](CapabilitySettings.md) |  | [optional] 
**TransferInstruments** | Pointer to [**[]SupportingEntityCapability**](SupportingEntityCapability.md) | Capability status for transfer instruments associated with legal entity | [optional] [readonly] 
**VerificationStatus** | Pointer to **string** | The status of the verification checks for the capability.  Possible values:  * **pending**: Adyen is running the verification.  * **invalid**: The verification failed. Check if the &#x60;errors&#x60; array contains more information.  * **valid**: The verification has been successfully completed.  * **rejected**: Adyen has verified the information, but found reasons to not allow the capability.  | [optional] [readonly] 

## Methods

### NewLegalEntityCapability

`func NewLegalEntityCapability() *LegalEntityCapability`

NewLegalEntityCapability instantiates a new LegalEntityCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalEntityCapabilityWithDefaults

`func NewLegalEntityCapabilityWithDefaults() *LegalEntityCapability`

NewLegalEntityCapabilityWithDefaults instantiates a new LegalEntityCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *LegalEntityCapability) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *LegalEntityCapability) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *LegalEntityCapability) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *LegalEntityCapability) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.

### GetAllowedLevel

`func (o *LegalEntityCapability) GetAllowedLevel() string`

GetAllowedLevel returns the AllowedLevel field if non-nil, zero value otherwise.

### GetAllowedLevelOk

`func (o *LegalEntityCapability) GetAllowedLevelOk() (*string, bool)`

GetAllowedLevelOk returns a tuple with the AllowedLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedLevel

`func (o *LegalEntityCapability) SetAllowedLevel(v string)`

SetAllowedLevel sets AllowedLevel field to given value.

### HasAllowedLevel

`func (o *LegalEntityCapability) HasAllowedLevel() bool`

HasAllowedLevel returns a boolean if a field has been set.

### GetAllowedSettings

`func (o *LegalEntityCapability) GetAllowedSettings() CapabilitySettings`

GetAllowedSettings returns the AllowedSettings field if non-nil, zero value otherwise.

### GetAllowedSettingsOk

`func (o *LegalEntityCapability) GetAllowedSettingsOk() (*CapabilitySettings, bool)`

GetAllowedSettingsOk returns a tuple with the AllowedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSettings

`func (o *LegalEntityCapability) SetAllowedSettings(v CapabilitySettings)`

SetAllowedSettings sets AllowedSettings field to given value.

### HasAllowedSettings

`func (o *LegalEntityCapability) HasAllowedSettings() bool`

HasAllowedSettings returns a boolean if a field has been set.

### GetRequested

`func (o *LegalEntityCapability) GetRequested() bool`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *LegalEntityCapability) GetRequestedOk() (*bool, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *LegalEntityCapability) SetRequested(v bool)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *LegalEntityCapability) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### GetRequestedLevel

`func (o *LegalEntityCapability) GetRequestedLevel() string`

GetRequestedLevel returns the RequestedLevel field if non-nil, zero value otherwise.

### GetRequestedLevelOk

`func (o *LegalEntityCapability) GetRequestedLevelOk() (*string, bool)`

GetRequestedLevelOk returns a tuple with the RequestedLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedLevel

`func (o *LegalEntityCapability) SetRequestedLevel(v string)`

SetRequestedLevel sets RequestedLevel field to given value.

### HasRequestedLevel

`func (o *LegalEntityCapability) HasRequestedLevel() bool`

HasRequestedLevel returns a boolean if a field has been set.

### GetRequestedSettings

`func (o *LegalEntityCapability) GetRequestedSettings() CapabilitySettings`

GetRequestedSettings returns the RequestedSettings field if non-nil, zero value otherwise.

### GetRequestedSettingsOk

`func (o *LegalEntityCapability) GetRequestedSettingsOk() (*CapabilitySettings, bool)`

GetRequestedSettingsOk returns a tuple with the RequestedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedSettings

`func (o *LegalEntityCapability) SetRequestedSettings(v CapabilitySettings)`

SetRequestedSettings sets RequestedSettings field to given value.

### HasRequestedSettings

`func (o *LegalEntityCapability) HasRequestedSettings() bool`

HasRequestedSettings returns a boolean if a field has been set.

### GetTransferInstruments

`func (o *LegalEntityCapability) GetTransferInstruments() []SupportingEntityCapability`

GetTransferInstruments returns the TransferInstruments field if non-nil, zero value otherwise.

### GetTransferInstrumentsOk

`func (o *LegalEntityCapability) GetTransferInstrumentsOk() (*[]SupportingEntityCapability, bool)`

GetTransferInstrumentsOk returns a tuple with the TransferInstruments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferInstruments

`func (o *LegalEntityCapability) SetTransferInstruments(v []SupportingEntityCapability)`

SetTransferInstruments sets TransferInstruments field to given value.

### HasTransferInstruments

`func (o *LegalEntityCapability) HasTransferInstruments() bool`

HasTransferInstruments returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *LegalEntityCapability) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *LegalEntityCapability) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *LegalEntityCapability) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *LegalEntityCapability) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


