# CostEstimateAssumptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assume3DSecureAuthenticated** | Pointer to **bool** | If true, the cardholder is expected to successfully authorise via 3D Secure. | [optional] 
**AssumeLevel3Data** | Pointer to **bool** | If true, the transaction is expected to have valid Level 3 data. | [optional] 
**Installments** | Pointer to **int32** | If not zero, the number of installments. | [optional] 

## Methods

### NewCostEstimateAssumptions

`func NewCostEstimateAssumptions() *CostEstimateAssumptions`

NewCostEstimateAssumptions instantiates a new CostEstimateAssumptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostEstimateAssumptionsWithDefaults

`func NewCostEstimateAssumptionsWithDefaults() *CostEstimateAssumptions`

NewCostEstimateAssumptionsWithDefaults instantiates a new CostEstimateAssumptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssume3DSecureAuthenticated

`func (o *CostEstimateAssumptions) GetAssume3DSecureAuthenticated() bool`

GetAssume3DSecureAuthenticated returns the Assume3DSecureAuthenticated field if non-nil, zero value otherwise.

### GetAssume3DSecureAuthenticatedOk

`func (o *CostEstimateAssumptions) GetAssume3DSecureAuthenticatedOk() (*bool, bool)`

GetAssume3DSecureAuthenticatedOk returns a tuple with the Assume3DSecureAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssume3DSecureAuthenticated

`func (o *CostEstimateAssumptions) SetAssume3DSecureAuthenticated(v bool)`

SetAssume3DSecureAuthenticated sets Assume3DSecureAuthenticated field to given value.

### HasAssume3DSecureAuthenticated

`func (o *CostEstimateAssumptions) HasAssume3DSecureAuthenticated() bool`

HasAssume3DSecureAuthenticated returns a boolean if a field has been set.

### GetAssumeLevel3Data

`func (o *CostEstimateAssumptions) GetAssumeLevel3Data() bool`

GetAssumeLevel3Data returns the AssumeLevel3Data field if non-nil, zero value otherwise.

### GetAssumeLevel3DataOk

`func (o *CostEstimateAssumptions) GetAssumeLevel3DataOk() (*bool, bool)`

GetAssumeLevel3DataOk returns a tuple with the AssumeLevel3Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumeLevel3Data

`func (o *CostEstimateAssumptions) SetAssumeLevel3Data(v bool)`

SetAssumeLevel3Data sets AssumeLevel3Data field to given value.

### HasAssumeLevel3Data

`func (o *CostEstimateAssumptions) HasAssumeLevel3Data() bool`

HasAssumeLevel3Data returns a boolean if a field has been set.

### GetInstallments

`func (o *CostEstimateAssumptions) GetInstallments() int32`

GetInstallments returns the Installments field if non-nil, zero value otherwise.

### GetInstallmentsOk

`func (o *CostEstimateAssumptions) GetInstallmentsOk() (*int32, bool)`

GetInstallmentsOk returns a tuple with the Installments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallments

`func (o *CostEstimateAssumptions) SetInstallments(v int32)`

SetInstallments sets Installments field to given value.

### HasInstallments

`func (o *CostEstimateAssumptions) HasInstallments() bool`

HasInstallments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


