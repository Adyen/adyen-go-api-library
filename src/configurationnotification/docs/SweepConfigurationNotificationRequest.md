# SweepConfigurationNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**SweepConfigurationNotificationData**](SweepConfigurationNotificationData.md) |  | 
**Environment** | **string** | The environment from which the webhook originated.  Possible values: **test**, **live**. | 
**Type** | **string** | Type of webhook. | 

## Methods

### NewSweepConfigurationNotificationRequest

`func NewSweepConfigurationNotificationRequest(data SweepConfigurationNotificationData, environment string, type_ string, ) *SweepConfigurationNotificationRequest`

NewSweepConfigurationNotificationRequest instantiates a new SweepConfigurationNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSweepConfigurationNotificationRequestWithDefaults

`func NewSweepConfigurationNotificationRequestWithDefaults() *SweepConfigurationNotificationRequest`

NewSweepConfigurationNotificationRequestWithDefaults instantiates a new SweepConfigurationNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SweepConfigurationNotificationRequest) GetData() SweepConfigurationNotificationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SweepConfigurationNotificationRequest) GetDataOk() (*SweepConfigurationNotificationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SweepConfigurationNotificationRequest) SetData(v SweepConfigurationNotificationData)`

SetData sets Data field to given value.


### GetEnvironment

`func (o *SweepConfigurationNotificationRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SweepConfigurationNotificationRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SweepConfigurationNotificationRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetType

`func (o *SweepConfigurationNotificationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SweepConfigurationNotificationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SweepConfigurationNotificationRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


