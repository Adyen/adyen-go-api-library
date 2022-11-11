# RestServiceError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | **string** | A human-readable explanation specific to this occurrence of the problem. | 
**ErrorCode** | **string** | A code that identifies the problem type. | 
**Instance** | Pointer to **string** | A unique URI that identifies the specific occurrence of the problem. | [optional] 
**InvalidFields** | Pointer to [**[]InvalidField**](InvalidField.md) | Detailed explanation of each validation error, when applicable. | [optional] 
**RequestId** | Pointer to **string** | A unique reference for the request, essentially the same as &#x60;pspReference&#x60;. | [optional] 
**Response** | Pointer to [**JSONObject**](JSONObject.md) |  | [optional] 
**Status** | **int32** | The HTTP status code. | 
**Title** | **string** | A short, human-readable summary of the problem type. | 
**Type** | **string** | A URI that identifies the problem type, pointing to human-readable documentation on this problem type. | 

## Methods

### NewRestServiceError

`func NewRestServiceError(detail string, errorCode string, status int32, title string, type_ string, ) *RestServiceError`

NewRestServiceError instantiates a new RestServiceError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestServiceErrorWithDefaults

`func NewRestServiceErrorWithDefaults() *RestServiceError`

NewRestServiceErrorWithDefaults instantiates a new RestServiceError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *RestServiceError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *RestServiceError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *RestServiceError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetErrorCode

`func (o *RestServiceError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *RestServiceError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *RestServiceError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetInstance

`func (o *RestServiceError) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *RestServiceError) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *RestServiceError) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *RestServiceError) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetInvalidFields

`func (o *RestServiceError) GetInvalidFields() []InvalidField`

GetInvalidFields returns the InvalidFields field if non-nil, zero value otherwise.

### GetInvalidFieldsOk

`func (o *RestServiceError) GetInvalidFieldsOk() (*[]InvalidField, bool)`

GetInvalidFieldsOk returns a tuple with the InvalidFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidFields

`func (o *RestServiceError) SetInvalidFields(v []InvalidField)`

SetInvalidFields sets InvalidFields field to given value.

### HasInvalidFields

`func (o *RestServiceError) HasInvalidFields() bool`

HasInvalidFields returns a boolean if a field has been set.

### GetRequestId

`func (o *RestServiceError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *RestServiceError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *RestServiceError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *RestServiceError) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetResponse

`func (o *RestServiceError) GetResponse() JSONObject`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *RestServiceError) GetResponseOk() (*JSONObject, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *RestServiceError) SetResponse(v JSONObject)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *RestServiceError) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetStatus

`func (o *RestServiceError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestServiceError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestServiceError) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *RestServiceError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RestServiceError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RestServiceError) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *RestServiceError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestServiceError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestServiceError) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


