# AccountSupportingEntityCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | Pointer to **bool** | Indicates whether the supporting entity capability is allowed. Adyen sets this to **true** if the verification is successful and the account holder is permitted to use the capability. | [optional] [readonly] 
**AllowedLevel** | Pointer to **string** | The capability level that is allowed for the account holder.  Possible values: **notApplicable**, **low**, **medium**, **high**. | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Indicates whether the capability is enabled. If **false**, the capability is temporarily disabled for the account holder. | [optional] 
**Id** | Pointer to **string** | The ID of the supporting entity. | [optional] [readonly] 
**Requested** | Pointer to **bool** | Indicates whether the capability is requested. To check whether the account holder is permitted to use the capability, refer to the &#x60;allowed&#x60; field. | [optional] 
**RequestedLevel** | Pointer to **string** | The requested level of the capability. Some capabilities, such as those used in [card issuing](https://docs.adyen.com/issuing/add-capabilities#capability-levels), have different levels. Levels increase the capability, but also require additional checks and increased monitoring.  Possible values: **notApplicable**, **low**, **medium**, **high**. | [optional] 
**VerificationStatus** | Pointer to **string** | The status of the verification checks for the supporting entity capability.  Possible values:  * **pending**: Adyen is running the verification.  * **invalid**: The verification failed. Check if the &#x60;errors&#x60; array contains more information.  * **valid**: The verification has been successfully completed.  * **rejected**: Adyen has verified the information, but found reasons to not allow the capability.  | [optional] [readonly] 

## Methods

### NewAccountSupportingEntityCapability

`func NewAccountSupportingEntityCapability() *AccountSupportingEntityCapability`

NewAccountSupportingEntityCapability instantiates a new AccountSupportingEntityCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSupportingEntityCapabilityWithDefaults

`func NewAccountSupportingEntityCapabilityWithDefaults() *AccountSupportingEntityCapability`

NewAccountSupportingEntityCapabilityWithDefaults instantiates a new AccountSupportingEntityCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *AccountSupportingEntityCapability) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *AccountSupportingEntityCapability) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *AccountSupportingEntityCapability) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *AccountSupportingEntityCapability) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.

### GetAllowedLevel

`func (o *AccountSupportingEntityCapability) GetAllowedLevel() string`

GetAllowedLevel returns the AllowedLevel field if non-nil, zero value otherwise.

### GetAllowedLevelOk

`func (o *AccountSupportingEntityCapability) GetAllowedLevelOk() (*string, bool)`

GetAllowedLevelOk returns a tuple with the AllowedLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedLevel

`func (o *AccountSupportingEntityCapability) SetAllowedLevel(v string)`

SetAllowedLevel sets AllowedLevel field to given value.

### HasAllowedLevel

`func (o *AccountSupportingEntityCapability) HasAllowedLevel() bool`

HasAllowedLevel returns a boolean if a field has been set.

### GetEnabled

`func (o *AccountSupportingEntityCapability) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AccountSupportingEntityCapability) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AccountSupportingEntityCapability) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AccountSupportingEntityCapability) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *AccountSupportingEntityCapability) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountSupportingEntityCapability) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountSupportingEntityCapability) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountSupportingEntityCapability) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequested

`func (o *AccountSupportingEntityCapability) GetRequested() bool`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *AccountSupportingEntityCapability) GetRequestedOk() (*bool, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *AccountSupportingEntityCapability) SetRequested(v bool)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *AccountSupportingEntityCapability) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### GetRequestedLevel

`func (o *AccountSupportingEntityCapability) GetRequestedLevel() string`

GetRequestedLevel returns the RequestedLevel field if non-nil, zero value otherwise.

### GetRequestedLevelOk

`func (o *AccountSupportingEntityCapability) GetRequestedLevelOk() (*string, bool)`

GetRequestedLevelOk returns a tuple with the RequestedLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedLevel

`func (o *AccountSupportingEntityCapability) SetRequestedLevel(v string)`

SetRequestedLevel sets RequestedLevel field to given value.

### HasRequestedLevel

`func (o *AccountSupportingEntityCapability) HasRequestedLevel() bool`

HasRequestedLevel returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *AccountSupportingEntityCapability) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *AccountSupportingEntityCapability) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *AccountSupportingEntityCapability) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *AccountSupportingEntityCapability) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


