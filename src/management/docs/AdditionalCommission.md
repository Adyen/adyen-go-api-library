# AdditionalCommission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceAccountId** | Pointer to **string** | Unique identifier of the balance account to which the additional commission is booked. | [optional] 
**FixedAmount** | Pointer to **int64** | A fixed commission fee, in minor units. | [optional] 
**VariablePercentage** | Pointer to **int64** | A variable commission fee, in basis points. | [optional] 

## Methods

### NewAdditionalCommission

`func NewAdditionalCommission() *AdditionalCommission`

NewAdditionalCommission instantiates a new AdditionalCommission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalCommissionWithDefaults

`func NewAdditionalCommissionWithDefaults() *AdditionalCommission`

NewAdditionalCommissionWithDefaults instantiates a new AdditionalCommission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceAccountId

`func (o *AdditionalCommission) GetBalanceAccountId() string`

GetBalanceAccountId returns the BalanceAccountId field if non-nil, zero value otherwise.

### GetBalanceAccountIdOk

`func (o *AdditionalCommission) GetBalanceAccountIdOk() (*string, bool)`

GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccountId

`func (o *AdditionalCommission) SetBalanceAccountId(v string)`

SetBalanceAccountId sets BalanceAccountId field to given value.

### HasBalanceAccountId

`func (o *AdditionalCommission) HasBalanceAccountId() bool`

HasBalanceAccountId returns a boolean if a field has been set.

### GetFixedAmount

`func (o *AdditionalCommission) GetFixedAmount() int64`

GetFixedAmount returns the FixedAmount field if non-nil, zero value otherwise.

### GetFixedAmountOk

`func (o *AdditionalCommission) GetFixedAmountOk() (*int64, bool)`

GetFixedAmountOk returns a tuple with the FixedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmount

`func (o *AdditionalCommission) SetFixedAmount(v int64)`

SetFixedAmount sets FixedAmount field to given value.

### HasFixedAmount

`func (o *AdditionalCommission) HasFixedAmount() bool`

HasFixedAmount returns a boolean if a field has been set.

### GetVariablePercentage

`func (o *AdditionalCommission) GetVariablePercentage() int64`

GetVariablePercentage returns the VariablePercentage field if non-nil, zero value otherwise.

### GetVariablePercentageOk

`func (o *AdditionalCommission) GetVariablePercentageOk() (*int64, bool)`

GetVariablePercentageOk returns a tuple with the VariablePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablePercentage

`func (o *AdditionalCommission) SetVariablePercentage(v int64)`

SetVariablePercentage sets VariablePercentage field to given value.

### HasVariablePercentage

`func (o *AdditionalCommission) HasVariablePercentage() bool`

HasVariablePercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


