# BalanceAccountNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**BalanceAccountNotificationData**](BalanceAccountNotificationData.md) |  | 
**Environment** | **string** | The environment from which the webhook originated.  Possible values: **test**, **live**. | 
**Type** | **string** | Type of webhook. | 

## Methods

### NewBalanceAccountNotificationRequest

`func NewBalanceAccountNotificationRequest(data BalanceAccountNotificationData, environment string, type_ string, ) *BalanceAccountNotificationRequest`

NewBalanceAccountNotificationRequest instantiates a new BalanceAccountNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceAccountNotificationRequestWithDefaults

`func NewBalanceAccountNotificationRequestWithDefaults() *BalanceAccountNotificationRequest`

NewBalanceAccountNotificationRequestWithDefaults instantiates a new BalanceAccountNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BalanceAccountNotificationRequest) GetData() BalanceAccountNotificationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BalanceAccountNotificationRequest) GetDataOk() (*BalanceAccountNotificationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BalanceAccountNotificationRequest) SetData(v BalanceAccountNotificationData)`

SetData sets Data field to given value.


### GetEnvironment

`func (o *BalanceAccountNotificationRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *BalanceAccountNotificationRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *BalanceAccountNotificationRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetType

`func (o *BalanceAccountNotificationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BalanceAccountNotificationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BalanceAccountNotificationRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


