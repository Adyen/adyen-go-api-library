# TransactionRulesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advice** | Pointer to **string** | The advice given by the Risk analysis. | [optional] 
**AllRulesPassed** | Pointer to **bool** | Indicates whether the transaction passed the evaluation for all transaction rules. | [optional] 
**FailedTransactionRules** | Pointer to [**[]TransactionEventViolation**](TransactionEventViolation.md) | Array containing all the transaction rules that the transaction violated. This list is only sent when &#x60;allRulesPassed&#x60; is **false**. | [optional] 
**Score** | Pointer to **int32** | The score of the Risk analysis. | [optional] 

## Methods

### NewTransactionRulesResult

`func NewTransactionRulesResult() *TransactionRulesResult`

NewTransactionRulesResult instantiates a new TransactionRulesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRulesResultWithDefaults

`func NewTransactionRulesResultWithDefaults() *TransactionRulesResult`

NewTransactionRulesResultWithDefaults instantiates a new TransactionRulesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvice

`func (o *TransactionRulesResult) GetAdvice() string`

GetAdvice returns the Advice field if non-nil, zero value otherwise.

### GetAdviceOk

`func (o *TransactionRulesResult) GetAdviceOk() (*string, bool)`

GetAdviceOk returns a tuple with the Advice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvice

`func (o *TransactionRulesResult) SetAdvice(v string)`

SetAdvice sets Advice field to given value.

### HasAdvice

`func (o *TransactionRulesResult) HasAdvice() bool`

HasAdvice returns a boolean if a field has been set.

### GetAllRulesPassed

`func (o *TransactionRulesResult) GetAllRulesPassed() bool`

GetAllRulesPassed returns the AllRulesPassed field if non-nil, zero value otherwise.

### GetAllRulesPassedOk

`func (o *TransactionRulesResult) GetAllRulesPassedOk() (*bool, bool)`

GetAllRulesPassedOk returns a tuple with the AllRulesPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRulesPassed

`func (o *TransactionRulesResult) SetAllRulesPassed(v bool)`

SetAllRulesPassed sets AllRulesPassed field to given value.

### HasAllRulesPassed

`func (o *TransactionRulesResult) HasAllRulesPassed() bool`

HasAllRulesPassed returns a boolean if a field has been set.

### GetFailedTransactionRules

`func (o *TransactionRulesResult) GetFailedTransactionRules() []TransactionEventViolation`

GetFailedTransactionRules returns the FailedTransactionRules field if non-nil, zero value otherwise.

### GetFailedTransactionRulesOk

`func (o *TransactionRulesResult) GetFailedTransactionRulesOk() (*[]TransactionEventViolation, bool)`

GetFailedTransactionRulesOk returns a tuple with the FailedTransactionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedTransactionRules

`func (o *TransactionRulesResult) SetFailedTransactionRules(v []TransactionEventViolation)`

SetFailedTransactionRules sets FailedTransactionRules field to given value.

### HasFailedTransactionRules

`func (o *TransactionRulesResult) HasFailedTransactionRules() bool`

HasFailedTransactionRules returns a boolean if a field has been set.

### GetScore

`func (o *TransactionRulesResult) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *TransactionRulesResult) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *TransactionRulesResult) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *TransactionRulesResult) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


