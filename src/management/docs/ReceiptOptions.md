# ReceiptOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logo** | Pointer to **string** | The receipt logo converted to a Base64-encoded string. The image must be a .bmp file of &lt; 256 KB, dimensions 240 (H) x 384 (W) px. | [optional] 

## Methods

### NewReceiptOptions

`func NewReceiptOptions() *ReceiptOptions`

NewReceiptOptions instantiates a new ReceiptOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptOptionsWithDefaults

`func NewReceiptOptionsWithDefaults() *ReceiptOptions`

NewReceiptOptionsWithDefaults instantiates a new ReceiptOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogo

`func (o *ReceiptOptions) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ReceiptOptions) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ReceiptOptions) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ReceiptOptions) HasLogo() bool`

HasLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


