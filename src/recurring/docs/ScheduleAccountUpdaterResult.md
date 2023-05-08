# ScheduleAccountUpdaterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PspReference** | **string** | Adyen&#39;s 16-character unique reference associated with the transaction. This value is globally unique; quote it when communicating with us about this request. | 
**Result** | **string** | The result of scheduling an Account Updater. If scheduling was successful, this field returns **Success**; otherwise it contains the error message. | 

## Methods

### NewScheduleAccountUpdaterResult

`func NewScheduleAccountUpdaterResult(pspReference string, result string, ) *ScheduleAccountUpdaterResult`

NewScheduleAccountUpdaterResult instantiates a new ScheduleAccountUpdaterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleAccountUpdaterResultWithDefaults

`func NewScheduleAccountUpdaterResultWithDefaults() *ScheduleAccountUpdaterResult`

NewScheduleAccountUpdaterResultWithDefaults instantiates a new ScheduleAccountUpdaterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPspReference

`func (o *ScheduleAccountUpdaterResult) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *ScheduleAccountUpdaterResult) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *ScheduleAccountUpdaterResult) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.


### GetResult

`func (o *ScheduleAccountUpdaterResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ScheduleAccountUpdaterResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ScheduleAccountUpdaterResult) SetResult(v string)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


