# Phone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cc** | Pointer to **string** | Country code. Length: 1–3 characters. | [optional] 
**Subscriber** | Pointer to **string** | Subscriber number. Maximum length: 15 characters. | [optional] 

## Methods

### NewPhone

`func NewPhone() *Phone`

NewPhone instantiates a new Phone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneWithDefaults

`func NewPhoneWithDefaults() *Phone`

NewPhoneWithDefaults instantiates a new Phone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCc

`func (o *Phone) GetCc() string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *Phone) GetCcOk() (*string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *Phone) SetCc(v string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *Phone) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetSubscriber

`func (o *Phone) GetSubscriber() string`

GetSubscriber returns the Subscriber field if non-nil, zero value otherwise.

### GetSubscriberOk

`func (o *Phone) GetSubscriberOk() (*string, bool)`

GetSubscriberOk returns a tuple with the Subscriber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriber

`func (o *Phone) SetSubscriber(v string)`

SetSubscriber sets Subscriber field to given value.

### HasSubscriber

`func (o *Phone) HasSubscriber() bool`

HasSubscriber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


