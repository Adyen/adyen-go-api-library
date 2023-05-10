# DisablePermitResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PspReference** | Pointer to **string** | A unique reference associated with the request. This value is globally unique; quote it when communicating with us about this request. | [optional] 
**Status** | Pointer to **string** | Status of the disable request. | [optional] 

## Methods

### NewDisablePermitResult

`func NewDisablePermitResult() *DisablePermitResult`

NewDisablePermitResult instantiates a new DisablePermitResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisablePermitResultWithDefaults

`func NewDisablePermitResultWithDefaults() *DisablePermitResult`

NewDisablePermitResultWithDefaults instantiates a new DisablePermitResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPspReference

`func (o *DisablePermitResult) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *DisablePermitResult) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *DisablePermitResult) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.

### HasPspReference

`func (o *DisablePermitResult) HasPspReference() bool`

HasPspReference returns a boolean if a field has been set.

### GetStatus

`func (o *DisablePermitResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DisablePermitResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DisablePermitResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DisablePermitResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


