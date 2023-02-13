# WebData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebAddress** | Pointer to **string** | The URL of the website or the app store URL. | [optional] 
**WebAddressId** | Pointer to **string** | The unique identifier of the web address. | [optional] [readonly] 

## Methods

### NewWebData

`func NewWebData() *WebData`

NewWebData instantiates a new WebData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebDataWithDefaults

`func NewWebDataWithDefaults() *WebData`

NewWebDataWithDefaults instantiates a new WebData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebAddress

`func (o *WebData) GetWebAddress() string`

GetWebAddress returns the WebAddress field if non-nil, zero value otherwise.

### GetWebAddressOk

`func (o *WebData) GetWebAddressOk() (*string, bool)`

GetWebAddressOk returns a tuple with the WebAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAddress

`func (o *WebData) SetWebAddress(v string)`

SetWebAddress sets WebAddress field to given value.

### HasWebAddress

`func (o *WebData) HasWebAddress() bool`

HasWebAddress returns a boolean if a field has been set.

### GetWebAddressId

`func (o *WebData) GetWebAddressId() string`

GetWebAddressId returns the WebAddressId field if non-nil, zero value otherwise.

### GetWebAddressIdOk

`func (o *WebData) GetWebAddressIdOk() (*string, bool)`

GetWebAddressIdOk returns a tuple with the WebAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAddressId

`func (o *WebData) SetWebAddressId(v string)`

SetWebAddressId sets WebAddressId field to given value.

### HasWebAddressId

`func (o *WebData) HasWebAddressId() bool`

HasWebAddressId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


