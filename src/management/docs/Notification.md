# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowButton** | Pointer to **bool** | Toggle to show or hide Notification button on the terminal | [optional] 

## Methods

### NewNotification

`func NewNotification() *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowButton

`func (o *Notification) GetShowButton() bool`

GetShowButton returns the ShowButton field if non-nil, zero value otherwise.

### GetShowButtonOk

`func (o *Notification) GetShowButtonOk() (*bool, bool)`

GetShowButtonOk returns a tuple with the ShowButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowButton

`func (o *Notification) SetShowButton(v bool)`

SetShowButton sets ShowButton field to given value.

### HasShowButton

`func (o *Notification) HasShowButton() bool`

HasShowButton returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


