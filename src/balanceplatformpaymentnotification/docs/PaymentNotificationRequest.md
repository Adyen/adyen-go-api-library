# PaymentNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**PaymentNotificationData**](PaymentNotificationData.md) |  | 
**Environment** | **string** | The environment from which the webhook originated.  Possible values: **test**, **live**. | 
**Type** | **string** | Type of webhook. | 

## Methods

### NewPaymentNotificationRequest

`func NewPaymentNotificationRequest(data PaymentNotificationData, environment string, type_ string, ) *PaymentNotificationRequest`

NewPaymentNotificationRequest instantiates a new PaymentNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentNotificationRequestWithDefaults

`func NewPaymentNotificationRequestWithDefaults() *PaymentNotificationRequest`

NewPaymentNotificationRequestWithDefaults instantiates a new PaymentNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaymentNotificationRequest) GetData() PaymentNotificationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentNotificationRequest) GetDataOk() (*PaymentNotificationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentNotificationRequest) SetData(v PaymentNotificationData)`

SetData sets Data field to given value.


### GetEnvironment

`func (o *PaymentNotificationRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PaymentNotificationRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PaymentNotificationRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetType

`func (o *PaymentNotificationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentNotificationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentNotificationRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


