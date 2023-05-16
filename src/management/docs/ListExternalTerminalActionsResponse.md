# ListExternalTerminalActionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ExternalTerminalAction**](ExternalTerminalAction.md) | The list of terminal actions. | [optional] 

## Methods

### NewListExternalTerminalActionsResponse

`func NewListExternalTerminalActionsResponse() *ListExternalTerminalActionsResponse`

NewListExternalTerminalActionsResponse instantiates a new ListExternalTerminalActionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListExternalTerminalActionsResponseWithDefaults

`func NewListExternalTerminalActionsResponseWithDefaults() *ListExternalTerminalActionsResponse`

NewListExternalTerminalActionsResponseWithDefaults instantiates a new ListExternalTerminalActionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListExternalTerminalActionsResponse) GetData() []ExternalTerminalAction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListExternalTerminalActionsResponse) GetDataOk() (*[]ExternalTerminalAction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListExternalTerminalActionsResponse) SetData(v []ExternalTerminalAction)`

SetData sets Data field to given value.

### HasData

`func (o *ListExternalTerminalActionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


