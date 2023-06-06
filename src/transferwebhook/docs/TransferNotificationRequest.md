# transferwebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**transferwebhookData**](transferwebhookData.md) |  | 
**Environment** | **string** | The environment from which the webhook originated.  Possible values: **test**, **live**. | 
**Type** | Pointer to **string** | The type of webhook. | [optional] 

## Methods

### NewtransferwebhookRequest

`func NewtransferwebhookRequest(data transferwebhookData, environment string, ) *transferwebhookRequest`

NewtransferwebhookRequest instantiates a new transferwebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewtransferwebhookRequestWithDefaults

`func NewtransferwebhookRequestWithDefaults() *transferwebhookRequest`

NewtransferwebhookRequestWithDefaults instantiates a new transferwebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *transferwebhookRequest) GetData() transferwebhookData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *transferwebhookRequest) GetDataOk() (*transferwebhookData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *transferwebhookRequest) SetData(v transferwebhookData)`

SetData sets Data field to given value.


### GetEnvironment

`func (o *transferwebhookRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *transferwebhookRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *transferwebhookRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetType

`func (o *transferwebhookRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *transferwebhookRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *transferwebhookRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *transferwebhookRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


