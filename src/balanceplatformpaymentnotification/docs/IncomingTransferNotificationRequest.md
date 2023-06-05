# IncomingTransferNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**IncomingTransferNotificationData**](IncomingTransferNotificationData.md) |  | 
**Environment** | **string** | The environment from which the webhook originated.  Possible values: **test**, **live**. | 
**Type** | **string** | Type of webhook. | 

## Methods

### NewIncomingTransferNotificationRequest

`func NewIncomingTransferNotificationRequest(data IncomingTransferNotificationData, environment string, type_ string, ) *IncomingTransferNotificationRequest`

NewIncomingTransferNotificationRequest instantiates a new IncomingTransferNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingTransferNotificationRequestWithDefaults

`func NewIncomingTransferNotificationRequestWithDefaults() *IncomingTransferNotificationRequest`

NewIncomingTransferNotificationRequestWithDefaults instantiates a new IncomingTransferNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncomingTransferNotificationRequest) GetData() IncomingTransferNotificationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncomingTransferNotificationRequest) GetDataOk() (*IncomingTransferNotificationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncomingTransferNotificationRequest) SetData(v IncomingTransferNotificationData)`

SetData sets Data field to given value.


### GetEnvironment

`func (o *IncomingTransferNotificationRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *IncomingTransferNotificationRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *IncomingTransferNotificationRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetType

`func (o *IncomingTransferNotificationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncomingTransferNotificationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncomingTransferNotificationRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


