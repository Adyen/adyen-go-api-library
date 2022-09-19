# ScheduleTerminalActionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionDetails** | Pointer to [**ScheduleTerminalActionsRequestActionDetails**](ScheduleTerminalActionsRequestActionDetails.md) |  | [optional] 
**ScheduledAt** | Pointer to **string** | The date and time when the action should happen.  Format: [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339), but without the **Z** before the time offset. For example, **2021-11-15T12:16:21+01:00**  The action is sent with the first [maintenance call](https://docs.adyen.com/point-of-sale/automating-terminal-management/terminal-actions-api#when-actions-take-effect) after the specified date and time in the time zone of the terminal.  An empty value causes the action to be sent as soon as possible: at the next maintenance call. | [optional] 
**StoreId** | Pointer to **string** | The unique ID of the [store](https://docs.adyen.com/api-explorer/#/ManagementService/latest/get/stores). If present, all terminals in the &#x60;terminalIds&#x60; list must be assigned to this store. | [optional] 
**TerminalIds** | Pointer to **[]string** | A list of unique IDs of the terminals to apply the action to. You can extract the IDs from the [GET &#x60;/terminals&#x60;](https://docs.adyen.com/api-explorer/#/ManagementService/latest/get/terminals) response. Maximum length: 100 IDs. | [optional] 

## Methods

### NewScheduleTerminalActionsRequest

`func NewScheduleTerminalActionsRequest() *ScheduleTerminalActionsRequest`

NewScheduleTerminalActionsRequest instantiates a new ScheduleTerminalActionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleTerminalActionsRequestWithDefaults

`func NewScheduleTerminalActionsRequestWithDefaults() *ScheduleTerminalActionsRequest`

NewScheduleTerminalActionsRequestWithDefaults instantiates a new ScheduleTerminalActionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionDetails

`func (o *ScheduleTerminalActionsRequest) GetActionDetails() ScheduleTerminalActionsRequestActionDetails`

GetActionDetails returns the ActionDetails field if non-nil, zero value otherwise.

### GetActionDetailsOk

`func (o *ScheduleTerminalActionsRequest) GetActionDetailsOk() (*ScheduleTerminalActionsRequestActionDetails, bool)`

GetActionDetailsOk returns a tuple with the ActionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDetails

`func (o *ScheduleTerminalActionsRequest) SetActionDetails(v ScheduleTerminalActionsRequestActionDetails)`

SetActionDetails sets ActionDetails field to given value.

### HasActionDetails

`func (o *ScheduleTerminalActionsRequest) HasActionDetails() bool`

HasActionDetails returns a boolean if a field has been set.

### GetScheduledAt

`func (o *ScheduleTerminalActionsRequest) GetScheduledAt() string`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *ScheduleTerminalActionsRequest) GetScheduledAtOk() (*string, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *ScheduleTerminalActionsRequest) SetScheduledAt(v string)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *ScheduleTerminalActionsRequest) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetStoreId

`func (o *ScheduleTerminalActionsRequest) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *ScheduleTerminalActionsRequest) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *ScheduleTerminalActionsRequest) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *ScheduleTerminalActionsRequest) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetTerminalIds

`func (o *ScheduleTerminalActionsRequest) GetTerminalIds() []string`

GetTerminalIds returns the TerminalIds field if non-nil, zero value otherwise.

### GetTerminalIdsOk

`func (o *ScheduleTerminalActionsRequest) GetTerminalIdsOk() (*[]string, bool)`

GetTerminalIdsOk returns a tuple with the TerminalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalIds

`func (o *ScheduleTerminalActionsRequest) SetTerminalIds(v []string)`

SetTerminalIds sets TerminalIds field to given value.

### HasTerminalIds

`func (o *ScheduleTerminalActionsRequest) HasTerminalIds() bool`

HasTerminalIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


