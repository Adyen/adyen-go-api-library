# BalancePlatformNotificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationResponse** | Pointer to **string** | Respond with **HTTP 200 OK** and &#x60;[accepted]&#x60; in the response body to [accept the webhook](https://docs.adyen.com/development-resources/webhooks#accept-notifications). | [optional] 

## Methods

### NewBalancePlatformNotificationResponse

`func NewBalancePlatformNotificationResponse() *BalancePlatformNotificationResponse`

NewBalancePlatformNotificationResponse instantiates a new BalancePlatformNotificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalancePlatformNotificationResponseWithDefaults

`func NewBalancePlatformNotificationResponseWithDefaults() *BalancePlatformNotificationResponse`

NewBalancePlatformNotificationResponseWithDefaults instantiates a new BalancePlatformNotificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationResponse

`func (o *BalancePlatformNotificationResponse) GetNotificationResponse() string`

GetNotificationResponse returns the NotificationResponse field if non-nil, zero value otherwise.

### GetNotificationResponseOk

`func (o *BalancePlatformNotificationResponse) GetNotificationResponseOk() (*string, bool)`

GetNotificationResponseOk returns a tuple with the NotificationResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationResponse

`func (o *BalancePlatformNotificationResponse) SetNotificationResponse(v string)`

SetNotificationResponse sets NotificationResponse field to given value.

### HasNotificationResponse

`func (o *BalancePlatformNotificationResponse) HasNotificationResponse() bool`

HasNotificationResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


