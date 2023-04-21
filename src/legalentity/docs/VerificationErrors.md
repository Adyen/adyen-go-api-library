# VerificationErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Problems** | Pointer to [**[]CapabilityProblem**](CapabilityProblem.md) | List of the verification errors. | [optional] 

## Methods

### NewVerificationErrors

`func NewVerificationErrors() *VerificationErrors`

NewVerificationErrors instantiates a new VerificationErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationErrorsWithDefaults

`func NewVerificationErrorsWithDefaults() *VerificationErrors`

NewVerificationErrorsWithDefaults instantiates a new VerificationErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProblems

`func (o *VerificationErrors) GetProblems() []CapabilityProblem`

GetProblems returns the Problems field if non-nil, zero value otherwise.

### GetProblemsOk

`func (o *VerificationErrors) GetProblemsOk() (*[]CapabilityProblem, bool)`

GetProblemsOk returns a tuple with the Problems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblems

`func (o *VerificationErrors) SetProblems(v []CapabilityProblem)`

SetProblems sets Problems field to given value.

### HasProblems

`func (o *VerificationErrors) HasProblems() bool`

HasProblems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


