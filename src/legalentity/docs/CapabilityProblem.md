# CapabilityProblem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to [**CapabilityProblemEntity**](CapabilityProblemEntity.md) |  | [optional] 
**VerificationErrors** | Pointer to [**[]VerificationError**](VerificationError.md) |  | [optional] 

## Methods

### NewCapabilityProblem

`func NewCapabilityProblem() *CapabilityProblem`

NewCapabilityProblem instantiates a new CapabilityProblem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityProblemWithDefaults

`func NewCapabilityProblemWithDefaults() *CapabilityProblem`

NewCapabilityProblemWithDefaults instantiates a new CapabilityProblem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *CapabilityProblem) GetEntity() CapabilityProblemEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *CapabilityProblem) GetEntityOk() (*CapabilityProblemEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *CapabilityProblem) SetEntity(v CapabilityProblemEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *CapabilityProblem) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetVerificationErrors

`func (o *CapabilityProblem) GetVerificationErrors() []VerificationError`

GetVerificationErrors returns the VerificationErrors field if non-nil, zero value otherwise.

### GetVerificationErrorsOk

`func (o *CapabilityProblem) GetVerificationErrorsOk() (*[]VerificationError, bool)`

GetVerificationErrorsOk returns a tuple with the VerificationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationErrors

`func (o *CapabilityProblem) SetVerificationErrors(v []VerificationError)`

SetVerificationErrors sets VerificationErrors field to given value.

### HasVerificationErrors

`func (o *CapabilityProblem) HasVerificationErrors() bool`

HasVerificationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


