# ValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **string** | The result of the check.  Possible values:  - **valid**: The validation was successful.  - **invalid**: The validation failed.  - **notValidated**: The validation was not performed because some services were unreachable or Adyen does not have the information needed to perform the check.  - **notApplicable**: The validation is not applicable. | [optional] 
**Type** | Pointer to **string** | Type of check. | [optional] 

## Methods

### NewValidationResult

`func NewValidationResult() *ValidationResult`

NewValidationResult instantiates a new ValidationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationResultWithDefaults

`func NewValidationResultWithDefaults() *ValidationResult`

NewValidationResultWithDefaults instantiates a new ValidationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ValidationResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ValidationResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ValidationResult) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ValidationResult) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetType

`func (o *ValidationResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValidationResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValidationResult) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ValidationResult) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


