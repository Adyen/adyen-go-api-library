# GetTerminalDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Terminal** | **string** | The unique terminal ID in the format &#x60;[Device model]-[Serial number]&#x60;.   For example, **V400m-324689776**. | 

## Methods

### NewGetTerminalDetailsRequest

`func NewGetTerminalDetailsRequest(terminal string, ) *GetTerminalDetailsRequest`

NewGetTerminalDetailsRequest instantiates a new GetTerminalDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTerminalDetailsRequestWithDefaults

`func NewGetTerminalDetailsRequestWithDefaults() *GetTerminalDetailsRequest`

NewGetTerminalDetailsRequestWithDefaults instantiates a new GetTerminalDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerminal

`func (o *GetTerminalDetailsRequest) GetTerminal() string`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *GetTerminalDetailsRequest) GetTerminalOk() (*string, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *GetTerminalDetailsRequest) SetTerminal(v string)`

SetTerminal sets Terminal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

