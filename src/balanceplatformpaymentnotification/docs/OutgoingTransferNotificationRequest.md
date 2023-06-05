# OutgoingTransferNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**OutgoingTransferNotificationData**](OutgoingTransferNotificationData.md) |  | 
**Environment** | **string** | The environment from which the webhook originated.  Possible values: **test**, **live**. | 
**Type** | **string** | Type of webhook. | 

## Methods

### NewOutgoingTransferNotificationRequest

`func NewOutgoingTransferNotificationRequest(data OutgoingTransferNotificationData, environment string, type_ string, ) *OutgoingTransferNotificationRequest`

NewOutgoingTransferNotificationRequest instantiates a new OutgoingTransferNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingTransferNotificationRequestWithDefaults

`func NewOutgoingTransferNotificationRequestWithDefaults() *OutgoingTransferNotificationRequest`

NewOutgoingTransferNotificationRequestWithDefaults instantiates a new OutgoingTransferNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OutgoingTransferNotificationRequest) GetData() OutgoingTransferNotificationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OutgoingTransferNotificationRequest) GetDataOk() (*OutgoingTransferNotificationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OutgoingTransferNotificationRequest) SetData(v OutgoingTransferNotificationData)`

SetData sets Data field to given value.


### GetEnvironment

`func (o *OutgoingTransferNotificationRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *OutgoingTransferNotificationRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *OutgoingTransferNotificationRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetType

`func (o *OutgoingTransferNotificationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutgoingTransferNotificationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutgoingTransferNotificationRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


