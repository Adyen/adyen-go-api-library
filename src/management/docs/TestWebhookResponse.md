# TestWebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TestOutput**](TestOutput.md) | List with test results. Each test webhook we send has a list element with the result. | [optional] 

## Methods

### NewTestWebhookResponse

`func NewTestWebhookResponse() *TestWebhookResponse`

NewTestWebhookResponse instantiates a new TestWebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestWebhookResponseWithDefaults

`func NewTestWebhookResponseWithDefaults() *TestWebhookResponse`

NewTestWebhookResponseWithDefaults instantiates a new TestWebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TestWebhookResponse) GetData() []TestOutput`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TestWebhookResponse) GetDataOk() (*[]TestOutput, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TestWebhookResponse) SetData(v []TestOutput)`

SetData sets Data field to given value.

### HasData

`func (o *TestWebhookResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


