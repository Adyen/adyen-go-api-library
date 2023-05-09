# TransactionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationLevel** | Pointer to **string** | The level at which data must be accumulated, used in rules with &#x60;type&#x60; **velocity** or **maxUsage**. The level must be the [same or lower in hierarchy](https://docs.adyen.com/issuing/transaction-rules#accumulate-data) than the &#x60;entityKey&#x60;.  If not provided, by default, the rule will accumulate data at the **paymentInstrument** level.  Possible values: **paymentInstrument**, **paymentInstrumentGroup**, **balanceAccount**, **accountHolder**, **balancePlatform**. | [optional] 
**Description** | **string** | Your description for the transaction rule, maximum 300 characters. | 
**EndDate** | Pointer to **string** | The date when the rule will stop being evaluated, in ISO 8601 extended offset date-time format. For example, **2020-12-18T10:15:30+01:00**.  If not provided, the rule will be evaluated until the rule status is set to **inactive**. | [optional] 
**EntityKey** | [**TransactionRuleEntityKey**](TransactionRuleEntityKey.md) |  | 
**Id** | Pointer to **string** | The unique identifier of the transaction rule. | [optional] 
**Interval** | [**TransactionRuleInterval**](TransactionRuleInterval.md) |  | 
**OutcomeType** | Pointer to **string** | The [outcome](https://docs.adyen.com/issuing/transaction-rules#outcome) that will be applied when a transaction meets the conditions of the rule. If not provided, by default, this is set to **hardBlock**.  Possible values:   * **hardBlock**: the transaction is declined.  * **scoreBased**: the transaction is assigned the &#x60;score&#x60; you specified. Adyen calculates the total score and if it exceeds 100, the transaction is declined. | [optional] 
**Reference** | **string** | Your reference for the transaction rule, maximum 150 characters. | 
**RequestType** | Pointer to **string** | Indicates the type of request to which the rule applies.  Possible values: **authorization**, **authentication**, **tokenization**. | [optional] 
**RuleRestrictions** | [**TransactionRuleRestrictions**](TransactionRuleRestrictions.md) |  | 
**Score** | Pointer to **int32** | A positive or negative score applied to the transaction if it meets the conditions of the rule. Required when &#x60;outcomeType&#x60; is **scoreBased**.  The value must be between **-100** and **100**. | [optional] 
**StartDate** | Pointer to **string** | The date when the rule will start to be evaluated, in ISO 8601 extended offset date-time format. For example, **2020-12-18T10:15:30+01:00**.  If not provided when creating a transaction rule, the &#x60;startDate&#x60; is set to the date when the rule status is set to **active**.    | [optional] 
**Status** | Pointer to **string** | The status of the transaction rule. If you provide a &#x60;startDate&#x60; in the request, the rule is automatically created  with an **active** status.   Possible values: **active**, **inactive**. | [optional] 
**Type** | **string** | The [type of rule](https://docs.adyen.com/issuing/transaction-rules#rule-types), which defines if a rule blocks transactions based on individual characteristics or accumulates data.  Possible values:  * **blockList**: decline a transaction when the conditions are met.  * **maxUsage**: add the amount or number of transactions for the lifetime of a payment instrument, and then decline a transaction when the specified limits are met.  * **velocity**: add the amount or number of transactions based on a specified time interval, and then decline a transaction when the specified limits are met.  | 

## Methods

### NewTransactionRule

`func NewTransactionRule(description string, entityKey TransactionRuleEntityKey, interval TransactionRuleInterval, reference string, ruleRestrictions TransactionRuleRestrictions, type_ string, ) *TransactionRule`

NewTransactionRule instantiates a new TransactionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRuleWithDefaults

`func NewTransactionRuleWithDefaults() *TransactionRule`

NewTransactionRuleWithDefaults instantiates a new TransactionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationLevel

`func (o *TransactionRule) GetAggregationLevel() string`

GetAggregationLevel returns the AggregationLevel field if non-nil, zero value otherwise.

### GetAggregationLevelOk

`func (o *TransactionRule) GetAggregationLevelOk() (*string, bool)`

GetAggregationLevelOk returns a tuple with the AggregationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationLevel

`func (o *TransactionRule) SetAggregationLevel(v string)`

SetAggregationLevel sets AggregationLevel field to given value.

### HasAggregationLevel

`func (o *TransactionRule) HasAggregationLevel() bool`

HasAggregationLevel returns a boolean if a field has been set.

### GetDescription

`func (o *TransactionRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionRule) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEndDate

`func (o *TransactionRule) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TransactionRule) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TransactionRule) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TransactionRule) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEntityKey

`func (o *TransactionRule) GetEntityKey() TransactionRuleEntityKey`

GetEntityKey returns the EntityKey field if non-nil, zero value otherwise.

### GetEntityKeyOk

`func (o *TransactionRule) GetEntityKeyOk() (*TransactionRuleEntityKey, bool)`

GetEntityKeyOk returns a tuple with the EntityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityKey

`func (o *TransactionRule) SetEntityKey(v TransactionRuleEntityKey)`

SetEntityKey sets EntityKey field to given value.


### GetId

`func (o *TransactionRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *TransactionRule) GetInterval() TransactionRuleInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *TransactionRule) GetIntervalOk() (*TransactionRuleInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *TransactionRule) SetInterval(v TransactionRuleInterval)`

SetInterval sets Interval field to given value.


### GetOutcomeType

`func (o *TransactionRule) GetOutcomeType() string`

GetOutcomeType returns the OutcomeType field if non-nil, zero value otherwise.

### GetOutcomeTypeOk

`func (o *TransactionRule) GetOutcomeTypeOk() (*string, bool)`

GetOutcomeTypeOk returns a tuple with the OutcomeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeType

`func (o *TransactionRule) SetOutcomeType(v string)`

SetOutcomeType sets OutcomeType field to given value.

### HasOutcomeType

`func (o *TransactionRule) HasOutcomeType() bool`

HasOutcomeType returns a boolean if a field has been set.

### GetReference

`func (o *TransactionRule) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *TransactionRule) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *TransactionRule) SetReference(v string)`

SetReference sets Reference field to given value.


### GetRequestType

`func (o *TransactionRule) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *TransactionRule) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *TransactionRule) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *TransactionRule) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetRuleRestrictions

`func (o *TransactionRule) GetRuleRestrictions() TransactionRuleRestrictions`

GetRuleRestrictions returns the RuleRestrictions field if non-nil, zero value otherwise.

### GetRuleRestrictionsOk

`func (o *TransactionRule) GetRuleRestrictionsOk() (*TransactionRuleRestrictions, bool)`

GetRuleRestrictionsOk returns a tuple with the RuleRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleRestrictions

`func (o *TransactionRule) SetRuleRestrictions(v TransactionRuleRestrictions)`

SetRuleRestrictions sets RuleRestrictions field to given value.


### GetScore

`func (o *TransactionRule) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *TransactionRule) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *TransactionRule) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *TransactionRule) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetStartDate

`func (o *TransactionRule) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TransactionRule) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TransactionRule) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TransactionRule) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *TransactionRule) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionRule) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionRule) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionRule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *TransactionRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionRule) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


