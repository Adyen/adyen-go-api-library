# PermitResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultKey** | Pointer to **string** | The key to link permit requests to permit results. | [optional] 
**Token** | Pointer to **string** | The permit token which is used to make payments by the partner company. | [optional] 

## Methods

### NewPermitResult

`func NewPermitResult() *PermitResult`

NewPermitResult instantiates a new PermitResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermitResultWithDefaults

`func NewPermitResultWithDefaults() *PermitResult`

NewPermitResultWithDefaults instantiates a new PermitResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultKey

`func (o *PermitResult) GetResultKey() string`

GetResultKey returns the ResultKey field if non-nil, zero value otherwise.

### GetResultKeyOk

`func (o *PermitResult) GetResultKeyOk() (*string, bool)`

GetResultKeyOk returns a tuple with the ResultKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultKey

`func (o *PermitResult) SetResultKey(v string)`

SetResultKey sets ResultKey field to given value.

### HasResultKey

`func (o *PermitResult) HasResultKey() bool`

HasResultKey returns a boolean if a field has been set.

### GetToken

`func (o *PermitResult) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PermitResult) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PermitResult) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PermitResult) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


