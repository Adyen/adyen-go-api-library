# TransactionRuleInterval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfMonth** | Pointer to **int32** | The day of month, used when the &#x60;duration.unit&#x60; is **months**. If not provided, by default, this is set to **1**, the first day of the month. | [optional] 
**DayOfWeek** | Pointer to **string** | The day of week, used when the &#x60;duration.unit&#x60; is **weeks**. If not provided, by default, this is set to **monday**.  Possible values: **sunday**, **monday**, **tuesday**, **wednesday**, **thursday**, **friday**. | [optional] 
**Duration** | Pointer to [**Duration**](Duration.md) |  | [optional] 
**TimeOfDay** | Pointer to **string** | The time of day, in **hh:mm:ss** format, used when the &#x60;duration.unit&#x60; is **hours**. If not provided, by default, this is set to **00:00:00**. | [optional] 
**TimeZone** | Pointer to **string** | The [time zone](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). For example, **Europe/Amsterdam**. By default, this is set to **UTC**. | [optional] 
**Type** | **string** | The [type of interval](https://docs.adyen.com/issuing/transaction-rules#time-intervals) during which the rule conditions and limits apply, and how often counters are reset.  Possible values:   * **perTransaction**: conditions are evaluated and the counters are reset for every transaction.  * **daily**: the counters are reset daily at 00:00:00 UTC.  * **weekly**: the counters are reset every Monday at 00:00:00 UTC.   * **monthly**: the counters reset every first day of the month at 00:00:00 UTC.   * **lifetime**: conditions are applied to the lifetime of the payment instrument.  * **rolling**: conditions are applied and the counters are reset based on a &#x60;duration&#x60;. If the reset date and time are not provided, Adyen applies the default reset time similar to fixed intervals. For example, if the duration is every two weeks, the counter resets every third Monday at 00:00:00 UTC.  * **sliding**: conditions are applied and the counters are reset based on the current time and a &#x60;duration&#x60; that you specify. | 

## Methods

### NewTransactionRuleInterval

`func NewTransactionRuleInterval(type_ string, ) *TransactionRuleInterval`

NewTransactionRuleInterval instantiates a new TransactionRuleInterval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRuleIntervalWithDefaults

`func NewTransactionRuleIntervalWithDefaults() *TransactionRuleInterval`

NewTransactionRuleIntervalWithDefaults instantiates a new TransactionRuleInterval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfMonth

`func (o *TransactionRuleInterval) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *TransactionRuleInterval) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *TransactionRuleInterval) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *TransactionRuleInterval) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *TransactionRuleInterval) GetDayOfWeek() string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *TransactionRuleInterval) GetDayOfWeekOk() (*string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *TransactionRuleInterval) SetDayOfWeek(v string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *TransactionRuleInterval) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetDuration

`func (o *TransactionRuleInterval) GetDuration() Duration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TransactionRuleInterval) GetDurationOk() (*Duration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TransactionRuleInterval) SetDuration(v Duration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TransactionRuleInterval) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTimeOfDay

`func (o *TransactionRuleInterval) GetTimeOfDay() string`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *TransactionRuleInterval) GetTimeOfDayOk() (*string, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *TransactionRuleInterval) SetTimeOfDay(v string)`

SetTimeOfDay sets TimeOfDay field to given value.

### HasTimeOfDay

`func (o *TransactionRuleInterval) HasTimeOfDay() bool`

HasTimeOfDay returns a boolean if a field has been set.

### GetTimeZone

`func (o *TransactionRuleInterval) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *TransactionRuleInterval) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *TransactionRuleInterval) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *TransactionRuleInterval) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetType

`func (o *TransactionRuleInterval) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionRuleInterval) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionRuleInterval) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


