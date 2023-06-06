# CardOrderNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**CardOrderItem**](CardOrderItem.md) |  | 
**Environment** | **string** | The environment from which the webhook originated.  Possible values: **test**, **live**. | 
**Type** | **string** | Type of webhook. | 

## Methods

### NewCardOrderNotificationRequest

`func NewCardOrderNotificationRequest(data CardOrderItem, environment string, type_ string, ) *CardOrderNotificationRequest`

NewCardOrderNotificationRequest instantiates a new CardOrderNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardOrderNotificationRequestWithDefaults

`func NewCardOrderNotificationRequestWithDefaults() *CardOrderNotificationRequest`

NewCardOrderNotificationRequestWithDefaults instantiates a new CardOrderNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CardOrderNotificationRequest) GetData() CardOrderItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CardOrderNotificationRequest) GetDataOk() (*CardOrderItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CardOrderNotificationRequest) SetData(v CardOrderItem)`

SetData sets Data field to given value.


### GetEnvironment

`func (o *CardOrderNotificationRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CardOrderNotificationRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CardOrderNotificationRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetType

`func (o *CardOrderNotificationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CardOrderNotificationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CardOrderNotificationRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


