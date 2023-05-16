# TerminalModelsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]IdName**](IdName.md) | The terminal models that the API credential has access to. | [optional] 

## Methods

### NewTerminalModelsResponse

`func NewTerminalModelsResponse() *TerminalModelsResponse`

NewTerminalModelsResponse instantiates a new TerminalModelsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalModelsResponseWithDefaults

`func NewTerminalModelsResponseWithDefaults() *TerminalModelsResponse`

NewTerminalModelsResponseWithDefaults instantiates a new TerminalModelsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TerminalModelsResponse) GetData() []IdName`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TerminalModelsResponse) GetDataOk() (*[]IdName, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TerminalModelsResponse) SetData(v []IdName)`

SetData sets Data field to given value.

### HasData

`func (o *TerminalModelsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


