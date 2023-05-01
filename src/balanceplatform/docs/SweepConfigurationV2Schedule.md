# SweepConfigurationV2Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | **string** | A [cron expression](https://en.wikipedia.org/wiki/Cron#CRON_expression) that is used to set the sweep schedule. The schedule uses the time zone of the balance account. For example, **30 17 * * MON** schedules a sweep every Monday at 17:30.  The expression must have five values separated by a single space in the following order:  * Minute: **0-59**  * Hour: **0-23**  * Day of the month: **1-31**  * Month: **1-12** or **JAN-DEC**  * Day of the week: **0-7** (0 and 7 are Sunday) or **MON-SUN**.  The following non-standard characters are supported: **&amp;ast;**, **L**, **#**, **W** and **_/_**. See [crontab guru](https://crontab.guru/) for more examples. | 
**Type** | Pointer to **string** | The schedule type.  Possible values:  * **cron**: push out funds based on a cron expression.  * **daily**: push out funds daily at 07:00 AM CET.  * **weekly**: push out funds every Monday at 07:00 AM CET.  * **monthly**: push out funds every first of the month at 07:00 AM CET.  * **balance**: pull in funds instantly if the balance is less than or equal to the &#x60;triggerAmount&#x60;. You can only use this for sweeps of &#x60;type&#x60; **pull** and when the source is a &#x60;merchantAccount&#x60; or &#x60;transferInstrument&#x60;. | [optional] 

## Methods

### NewSweepConfigurationV2Schedule

`func NewSweepConfigurationV2Schedule(cronExpression string, ) *SweepConfigurationV2Schedule`

NewSweepConfigurationV2Schedule instantiates a new SweepConfigurationV2Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSweepConfigurationV2ScheduleWithDefaults

`func NewSweepConfigurationV2ScheduleWithDefaults() *SweepConfigurationV2Schedule`

NewSweepConfigurationV2ScheduleWithDefaults instantiates a new SweepConfigurationV2Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpression

`func (o *SweepConfigurationV2Schedule) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *SweepConfigurationV2Schedule) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *SweepConfigurationV2Schedule) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.


### GetType

`func (o *SweepConfigurationV2Schedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SweepConfigurationV2Schedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SweepConfigurationV2Schedule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SweepConfigurationV2Schedule) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


