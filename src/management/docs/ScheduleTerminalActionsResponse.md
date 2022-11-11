# ScheduleTerminalActionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionDetails** | Pointer to [**ScheduleTerminalActionsRequestActionDetails**](ScheduleTerminalActionsRequestActionDetails.md) |  | [optional] 
**Items** | Pointer to [**[]TerminalActionScheduleDetail**](TerminalActionScheduleDetail.md) |  | [optional] 
**ScheduledAt** | Pointer to **string** | The date and time when the action should happen.  Format: [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339), but without the **Z** before the time offset. For example, **2021-11-15T12:16:21+01:00**  The action is sent with the first [maintenance call](https://docs.adyen.com/point-of-sale/automating-terminal-management/terminal-actions-api#when-actions-take-effect) after the specified date and time in the time zone of the terminal.  An empty value causes the action to be sent as soon as possible: at the next maintenance call. | [optional] 
**StoreId** | Pointer to **string** | The unique ID of the [store](https://docs.adyen.com/api-explorer/#/ManagementService/latest/get/stores). If present, all terminals in the &#x60;terminalIds&#x60; list must be assigned to this store. | [optional] 
**TerminalIds** | Pointer to **[]string** | A list of unique IDs of the terminals to apply the action to. You can extract the IDs from the [GET &#x60;/terminals&#x60;](https://docs.adyen.com/api-explorer/#/ManagementService/latest/get/terminals) response. Maximum length: 100 IDs. | [optional] 
**TerminalsWithErrors** | Pointer to **map[string][]string** | The validation errors that occurred in the list of terminals, and for each error the IDs of the terminals that the error applies to. | [optional] 
**TotalErrors** | Pointer to **int32** | The number of terminals for which scheduling the action failed. | [optional] 
**TotalScheduled** | Pointer to **int32** | The number of terminals for which the action was successfully scheduled. This doesn&#39;t mean the action has happened yet. | [optional] 

## Methods

### NewScheduleTerminalActionsResponse

`func NewScheduleTerminalActionsResponse() *ScheduleTerminalActionsResponse`

NewScheduleTerminalActionsResponse instantiates a new ScheduleTerminalActionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleTerminalActionsResponseWithDefaults

`func NewScheduleTerminalActionsResponseWithDefaults() *ScheduleTerminalActionsResponse`

NewScheduleTerminalActionsResponseWithDefaults instantiates a new ScheduleTerminalActionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionDetails

`func (o *ScheduleTerminalActionsResponse) GetActionDetails() ScheduleTerminalActionsRequestActionDetails`

GetActionDetails returns the ActionDetails field if non-nil, zero value otherwise.

### GetActionDetailsOk

`func (o *ScheduleTerminalActionsResponse) GetActionDetailsOk() (*ScheduleTerminalActionsRequestActionDetails, bool)`

GetActionDetailsOk returns a tuple with the ActionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDetails

`func (o *ScheduleTerminalActionsResponse) SetActionDetails(v ScheduleTerminalActionsRequestActionDetails)`

SetActionDetails sets ActionDetails field to given value.

### HasActionDetails

`func (o *ScheduleTerminalActionsResponse) HasActionDetails() bool`

HasActionDetails returns a boolean if a field has been set.

### GetItems

`func (o *ScheduleTerminalActionsResponse) GetItems() []TerminalActionScheduleDetail`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ScheduleTerminalActionsResponse) GetItemsOk() (*[]TerminalActionScheduleDetail, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ScheduleTerminalActionsResponse) SetItems(v []TerminalActionScheduleDetail)`

SetItems sets Items field to given value.

### HasItems

`func (o *ScheduleTerminalActionsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetScheduledAt

`func (o *ScheduleTerminalActionsResponse) GetScheduledAt() string`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *ScheduleTerminalActionsResponse) GetScheduledAtOk() (*string, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *ScheduleTerminalActionsResponse) SetScheduledAt(v string)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *ScheduleTerminalActionsResponse) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetStoreId

`func (o *ScheduleTerminalActionsResponse) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *ScheduleTerminalActionsResponse) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *ScheduleTerminalActionsResponse) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *ScheduleTerminalActionsResponse) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetTerminalIds

`func (o *ScheduleTerminalActionsResponse) GetTerminalIds() []string`

GetTerminalIds returns the TerminalIds field if non-nil, zero value otherwise.

### GetTerminalIdsOk

`func (o *ScheduleTerminalActionsResponse) GetTerminalIdsOk() (*[]string, bool)`

GetTerminalIdsOk returns a tuple with the TerminalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalIds

`func (o *ScheduleTerminalActionsResponse) SetTerminalIds(v []string)`

SetTerminalIds sets TerminalIds field to given value.

### HasTerminalIds

`func (o *ScheduleTerminalActionsResponse) HasTerminalIds() bool`

HasTerminalIds returns a boolean if a field has been set.

### GetTerminalsWithErrors

`func (o *ScheduleTerminalActionsResponse) GetTerminalsWithErrors() map[string][]string`

GetTerminalsWithErrors returns the TerminalsWithErrors field if non-nil, zero value otherwise.

### GetTerminalsWithErrorsOk

`func (o *ScheduleTerminalActionsResponse) GetTerminalsWithErrorsOk() (*map[string][]string, bool)`

GetTerminalsWithErrorsOk returns a tuple with the TerminalsWithErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalsWithErrors

`func (o *ScheduleTerminalActionsResponse) SetTerminalsWithErrors(v map[string][]string)`

SetTerminalsWithErrors sets TerminalsWithErrors field to given value.

### HasTerminalsWithErrors

`func (o *ScheduleTerminalActionsResponse) HasTerminalsWithErrors() bool`

HasTerminalsWithErrors returns a boolean if a field has been set.

### GetTotalErrors

`func (o *ScheduleTerminalActionsResponse) GetTotalErrors() int32`

GetTotalErrors returns the TotalErrors field if non-nil, zero value otherwise.

### GetTotalErrorsOk

`func (o *ScheduleTerminalActionsResponse) GetTotalErrorsOk() (*int32, bool)`

GetTotalErrorsOk returns a tuple with the TotalErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalErrors

`func (o *ScheduleTerminalActionsResponse) SetTotalErrors(v int32)`

SetTotalErrors sets TotalErrors field to given value.

### HasTotalErrors

`func (o *ScheduleTerminalActionsResponse) HasTotalErrors() bool`

HasTotalErrors returns a boolean if a field has been set.

### GetTotalScheduled

`func (o *ScheduleTerminalActionsResponse) GetTotalScheduled() int32`

GetTotalScheduled returns the TotalScheduled field if non-nil, zero value otherwise.

### GetTotalScheduledOk

`func (o *ScheduleTerminalActionsResponse) GetTotalScheduledOk() (*int32, bool)`

GetTotalScheduledOk returns a tuple with the TotalScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalScheduled

`func (o *ScheduleTerminalActionsResponse) SetTotalScheduled(v int32)`

SetTotalScheduled sets TotalScheduled field to given value.

### HasTotalScheduled

`func (o *ScheduleTerminalActionsResponse) HasTotalScheduled() bool`

HasTotalScheduled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


