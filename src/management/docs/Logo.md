# Logo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | The image file, converted to a Base64-encoded string, of the logo to be shown on the terminal. | [optional] 

## Methods

### NewLogo

`func NewLogo() *Logo`

NewLogo instantiates a new Logo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogoWithDefaults

`func NewLogoWithDefaults() *Logo`

NewLogoWithDefaults instantiates a new Logo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Logo) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Logo) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Logo) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *Logo) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


