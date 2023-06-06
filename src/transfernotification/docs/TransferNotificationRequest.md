# TransferNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**TransferNotificationData**](TransferNotificationData.md) |  | 
**Environment** | **string** | The environment from which the webhook originated.  Possible values: **test**, **live**. | 
**Type** | Pointer to **string** | The type of webhook. | [optional] 

## Methods

### NewTransferNotificationRequest

`func NewTransferNotificationRequest(data TransferNotificationData, environment string, ) *TransferNotificationRequest`

NewTransferNotificationRequest instantiates a new TransferNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferNotificationRequestWithDefaults

`func NewTransferNotificationRequestWithDefaults() *TransferNotificationRequest`

NewTransferNotificationRequestWithDefaults instantiates a new TransferNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TransferNotificationRequest) GetData() TransferNotificationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransferNotificationRequest) GetDataOk() (*TransferNotificationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransferNotificationRequest) SetData(v TransferNotificationData)`

SetData sets Data field to given value.


### GetEnvironment

`func (o *TransferNotificationRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *TransferNotificationRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *TransferNotificationRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetType

`func (o *TransferNotificationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransferNotificationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransferNotificationRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransferNotificationRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


