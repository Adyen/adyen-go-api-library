# NotificationUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalUrls** | Pointer to [**[]Url**](Url.md) | One or more local URLs to send notifications to when using Terminal API. | [optional] 
**PublicUrls** | Pointer to [**[]Url**](Url.md) | One or more public URLs to send notifications to when using Terminal API. | [optional] 

## Methods

### NewNotificationUrl

`func NewNotificationUrl() *NotificationUrl`

NewNotificationUrl instantiates a new NotificationUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationUrlWithDefaults

`func NewNotificationUrlWithDefaults() *NotificationUrl`

NewNotificationUrlWithDefaults instantiates a new NotificationUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalUrls

`func (o *NotificationUrl) GetLocalUrls() []Url`

GetLocalUrls returns the LocalUrls field if non-nil, zero value otherwise.

### GetLocalUrlsOk

`func (o *NotificationUrl) GetLocalUrlsOk() (*[]Url, bool)`

GetLocalUrlsOk returns a tuple with the LocalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUrls

`func (o *NotificationUrl) SetLocalUrls(v []Url)`

SetLocalUrls sets LocalUrls field to given value.

### HasLocalUrls

`func (o *NotificationUrl) HasLocalUrls() bool`

HasLocalUrls returns a boolean if a field has been set.

### GetPublicUrls

`func (o *NotificationUrl) GetPublicUrls() []Url`

GetPublicUrls returns the PublicUrls field if non-nil, zero value otherwise.

### GetPublicUrlsOk

`func (o *NotificationUrl) GetPublicUrlsOk() (*[]Url, bool)`

GetPublicUrlsOk returns a tuple with the PublicUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicUrls

`func (o *NotificationUrl) SetPublicUrls(v []Url)`

SetPublicUrls sets PublicUrls field to given value.

### HasPublicUrls

`func (o *NotificationUrl) HasPublicUrls() bool`

HasPublicUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


