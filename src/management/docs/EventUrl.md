# EventUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventLocalUrls** | Pointer to [**[]Url**](Url.md) | One or more local URLs to send event notifications to when using Terminal API. | [optional] 
**EventPublicUrls** | Pointer to [**[]Url**](Url.md) | One or more public URLs to send event notifications to when using Terminal API. | [optional] 

## Methods

### NewEventUrl

`func NewEventUrl() *EventUrl`

NewEventUrl instantiates a new EventUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventUrlWithDefaults

`func NewEventUrlWithDefaults() *EventUrl`

NewEventUrlWithDefaults instantiates a new EventUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventLocalUrls

`func (o *EventUrl) GetEventLocalUrls() []Url`

GetEventLocalUrls returns the EventLocalUrls field if non-nil, zero value otherwise.

### GetEventLocalUrlsOk

`func (o *EventUrl) GetEventLocalUrlsOk() (*[]Url, bool)`

GetEventLocalUrlsOk returns a tuple with the EventLocalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLocalUrls

`func (o *EventUrl) SetEventLocalUrls(v []Url)`

SetEventLocalUrls sets EventLocalUrls field to given value.

### HasEventLocalUrls

`func (o *EventUrl) HasEventLocalUrls() bool`

HasEventLocalUrls returns a boolean if a field has been set.

### GetEventPublicUrls

`func (o *EventUrl) GetEventPublicUrls() []Url`

GetEventPublicUrls returns the EventPublicUrls field if non-nil, zero value otherwise.

### GetEventPublicUrlsOk

`func (o *EventUrl) GetEventPublicUrlsOk() (*[]Url, bool)`

GetEventPublicUrlsOk returns a tuple with the EventPublicUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPublicUrls

`func (o *EventUrl) SetEventPublicUrls(v []Url)`

SetEventPublicUrls sets EventPublicUrls field to given value.

### HasEventPublicUrls

`func (o *EventUrl) HasEventPublicUrls() bool`

HasEventPublicUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


