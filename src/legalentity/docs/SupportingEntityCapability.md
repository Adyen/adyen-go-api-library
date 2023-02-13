# SupportingEntityCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | Pointer to **bool** | Indicates whether the supporting entity capability is allowed.  If a supporting entity is allowed but its parent legal entity is not, it means there are other supporting entities that failed validation.  **The allowed supporting entity can still be used** | [optional] [readonly] 
**Id** | Pointer to **string** | Supporting entity reference  | [optional] [readonly] 
**Problems** | Pointer to [**[]CapabilityProblem**](CapabilityProblem.md) | Contains verification errors and the actions that you can take to resolve them. | [optional] [readonly] 
**Requested** | Pointer to **bool** | Indicates whether the supporting entity capability is requested.  | [optional] [readonly] 
**VerificationStatus** | Pointer to **string** | The status of the verification checks for the supporting entity capability.  Possible values:  * **pending**: Adyen is running the verification.  * **invalid**: The verification failed. Check if the &#x60;errors&#x60; array contains more information.  * **valid**: The verification has been successfully completed.  * **rejected**: Adyen has verified the information, but found reasons to not allow the capability.  | [optional] [readonly] 

## Methods

### NewSupportingEntityCapability

`func NewSupportingEntityCapability() *SupportingEntityCapability`

NewSupportingEntityCapability instantiates a new SupportingEntityCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportingEntityCapabilityWithDefaults

`func NewSupportingEntityCapabilityWithDefaults() *SupportingEntityCapability`

NewSupportingEntityCapabilityWithDefaults instantiates a new SupportingEntityCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *SupportingEntityCapability) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *SupportingEntityCapability) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *SupportingEntityCapability) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *SupportingEntityCapability) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.

### GetId

`func (o *SupportingEntityCapability) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportingEntityCapability) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportingEntityCapability) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportingEntityCapability) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProblems

`func (o *SupportingEntityCapability) GetProblems() []CapabilityProblem`

GetProblems returns the Problems field if non-nil, zero value otherwise.

### GetProblemsOk

`func (o *SupportingEntityCapability) GetProblemsOk() (*[]CapabilityProblem, bool)`

GetProblemsOk returns a tuple with the Problems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblems

`func (o *SupportingEntityCapability) SetProblems(v []CapabilityProblem)`

SetProblems sets Problems field to given value.

### HasProblems

`func (o *SupportingEntityCapability) HasProblems() bool`

HasProblems returns a boolean if a field has been set.

### GetRequested

`func (o *SupportingEntityCapability) GetRequested() bool`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *SupportingEntityCapability) GetRequestedOk() (*bool, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *SupportingEntityCapability) SetRequested(v bool)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *SupportingEntityCapability) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *SupportingEntityCapability) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *SupportingEntityCapability) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *SupportingEntityCapability) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *SupportingEntityCapability) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


