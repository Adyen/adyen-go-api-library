# Url

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | The password for authentication of the event notifications. | [optional] 
**Url** | Pointer to **string** | The URL in the format: http(s)://domain.com. | [optional] 
**Username** | Pointer to **string** | The username for authentication of the event notifications. | [optional] 

## Methods

### NewUrl

`func NewUrl() *Url`

NewUrl instantiates a new Url object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlWithDefaults

`func NewUrlWithDefaults() *Url`

NewUrlWithDefaults instantiates a new Url object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *Url) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Url) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Url) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Url) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUrl

`func (o *Url) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Url) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Url) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Url) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *Url) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Url) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Url) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Url) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


