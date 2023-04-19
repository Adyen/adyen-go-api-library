# ServiceError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | Contains additional information about the payment. Some data fields are included only if you select them first. Go to **Customer Area** &gt; **Developers** &gt; **Additional data**. | [optional] 
**ErrorCode** | Pointer to **string** | The error code mapped to the error message. | [optional] 
**ErrorType** | Pointer to **string** | The category of the error. | [optional] 
**Message** | Pointer to **string** | A short explanation of the issue. | [optional] 
**PspReference** | Pointer to **string** | The PSP reference of the payment. | [optional] 
**Status** | Pointer to **int32** | The HTTP response status. | [optional] 

## Methods

### NewServiceError

`func NewServiceError() *ServiceError`

NewServiceError instantiates a new ServiceError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceErrorWithDefaults

`func NewServiceErrorWithDefaults() *ServiceError`

NewServiceErrorWithDefaults instantiates a new ServiceError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *ServiceError) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *ServiceError) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *ServiceError) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *ServiceError) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetErrorCode

`func (o *ServiceError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ServiceError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ServiceError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ServiceError) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorType

`func (o *ServiceError) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *ServiceError) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *ServiceError) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *ServiceError) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetMessage

`func (o *ServiceError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServiceError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServiceError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ServiceError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPspReference

`func (o *ServiceError) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *ServiceError) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *ServiceError) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.

### HasPspReference

`func (o *ServiceError) HasPspReference() bool`

HasPspReference returns a boolean if a field has been set.

### GetStatus

`func (o *ServiceError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceError) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


