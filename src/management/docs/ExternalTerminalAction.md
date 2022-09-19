# ExternalTerminalAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | Pointer to **string** | The type of terminal action: **InstallAndroidApp**, **UninstallAndroidApp**, **InstallAndroidCertificate**, or **UninstallAndroidCertificate**. | [optional] 
**Config** | Pointer to **string** | Technical information about the terminal action. | [optional] 
**ConfirmedAt** | Pointer to **time.Time** | The date and time when the action was carried out. | [optional] 
**Id** | Pointer to **string** | The unique ID of the terminal action. | [optional] 
**Result** | Pointer to **string** | The result message for the action. | [optional] 
**ScheduledAt** | Pointer to **time.Time** | The date and time when the action was scheduled to happen. | [optional] 
**Status** | Pointer to **string** | The status of the terminal action: **pending**, **successful**, **failed**, **cancelled**, or **tryLater**. | [optional] 
**TerminalId** | Pointer to **string** | The unique ID of the terminal that the action applies to. | [optional] 

## Methods

### NewExternalTerminalAction

`func NewExternalTerminalAction() *ExternalTerminalAction`

NewExternalTerminalAction instantiates a new ExternalTerminalAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalTerminalActionWithDefaults

`func NewExternalTerminalActionWithDefaults() *ExternalTerminalAction`

NewExternalTerminalActionWithDefaults instantiates a new ExternalTerminalAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *ExternalTerminalAction) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *ExternalTerminalAction) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *ExternalTerminalAction) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *ExternalTerminalAction) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetConfig

`func (o *ExternalTerminalAction) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ExternalTerminalAction) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ExternalTerminalAction) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ExternalTerminalAction) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConfirmedAt

`func (o *ExternalTerminalAction) GetConfirmedAt() time.Time`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *ExternalTerminalAction) GetConfirmedAtOk() (*time.Time, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *ExternalTerminalAction) SetConfirmedAt(v time.Time)`

SetConfirmedAt sets ConfirmedAt field to given value.

### HasConfirmedAt

`func (o *ExternalTerminalAction) HasConfirmedAt() bool`

HasConfirmedAt returns a boolean if a field has been set.

### GetId

`func (o *ExternalTerminalAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalTerminalAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalTerminalAction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExternalTerminalAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResult

`func (o *ExternalTerminalAction) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ExternalTerminalAction) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ExternalTerminalAction) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ExternalTerminalAction) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetScheduledAt

`func (o *ExternalTerminalAction) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *ExternalTerminalAction) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *ExternalTerminalAction) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *ExternalTerminalAction) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetStatus

`func (o *ExternalTerminalAction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalTerminalAction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalTerminalAction) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExternalTerminalAction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTerminalId

`func (o *ExternalTerminalAction) GetTerminalId() string`

GetTerminalId returns the TerminalId field if non-nil, zero value otherwise.

### GetTerminalIdOk

`func (o *ExternalTerminalAction) GetTerminalIdOk() (*string, bool)`

GetTerminalIdOk returns a tuple with the TerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalId

`func (o *ExternalTerminalAction) SetTerminalId(v string)`

SetTerminalId sets TerminalId field to given value.

### HasTerminalId

`func (o *ExternalTerminalAction) HasTerminalId() bool`

HasTerminalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


