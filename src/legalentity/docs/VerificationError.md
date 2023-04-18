# VerificationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to **[]string** | Contains key-value pairs that specify the actions that the legal entity can do in your platform. The key is a capability required for your integration. For example, **issueCard** for Issuing.The value is an object containing the settings for the capability. | [optional] 
**Code** | Pointer to **string** | The general error code. | [optional] 
**Message** | Pointer to **string** | The general error message. | [optional] 
**RemediatingActions** | Pointer to [**[]RemediatingAction**](RemediatingAction.md) | An object containing possible solutions to fix a verification error. | [optional] 
**SubErrors** | Pointer to [**[]VerificationErrorRecursive**](VerificationErrorRecursive.md) | An array containing more granular information about the cause of the verification error. | [optional] 
**Type** | Pointer to **string** | The type of error. | [optional] 

## Methods

### NewVerificationError

`func NewVerificationError() *VerificationError`

NewVerificationError instantiates a new VerificationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationErrorWithDefaults

`func NewVerificationErrorWithDefaults() *VerificationError`

NewVerificationErrorWithDefaults instantiates a new VerificationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *VerificationError) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *VerificationError) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *VerificationError) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *VerificationError) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCode

`func (o *VerificationError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VerificationError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VerificationError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VerificationError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *VerificationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VerificationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VerificationError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VerificationError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRemediatingActions

`func (o *VerificationError) GetRemediatingActions() []RemediatingAction`

GetRemediatingActions returns the RemediatingActions field if non-nil, zero value otherwise.

### GetRemediatingActionsOk

`func (o *VerificationError) GetRemediatingActionsOk() (*[]RemediatingAction, bool)`

GetRemediatingActionsOk returns a tuple with the RemediatingActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediatingActions

`func (o *VerificationError) SetRemediatingActions(v []RemediatingAction)`

SetRemediatingActions sets RemediatingActions field to given value.

### HasRemediatingActions

`func (o *VerificationError) HasRemediatingActions() bool`

HasRemediatingActions returns a boolean if a field has been set.

### GetSubErrors

`func (o *VerificationError) GetSubErrors() []VerificationErrorRecursive`

GetSubErrors returns the SubErrors field if non-nil, zero value otherwise.

### GetSubErrorsOk

`func (o *VerificationError) GetSubErrorsOk() (*[]VerificationErrorRecursive, bool)`

GetSubErrorsOk returns a tuple with the SubErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubErrors

`func (o *VerificationError) SetSubErrors(v []VerificationErrorRecursive)`

SetSubErrors sets SubErrors field to given value.

### HasSubErrors

`func (o *VerificationError) HasSubErrors() bool`

HasSubErrors returns a boolean if a field has been set.

### GetType

`func (o *VerificationError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VerificationError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VerificationError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VerificationError) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


