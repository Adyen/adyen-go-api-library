# ReceiptOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logo** | Pointer to **string** | The receipt logo converted to a Base64-encoded string. The image must be a .bmp file of &lt; 256 KB, dimensions 240 (H) x 384 (W) px. | [optional] 
**QrCodeData** | Pointer to **string** | Data to print on the receipt as a QR code. This can include static text and the following variables:  - &#x60;${merchantreference}&#x60;: the merchant reference of the transaction. - &#x60;${pspreference}&#x60;: the PSP reference of the transaction.   For example, **http://www.example.com/order/${pspreference}/${merchantreference}**. | [optional] 

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

### GetQrCodeData

`func (o *ReceiptOptions) GetQrCodeData() string`

GetQrCodeData returns the QrCodeData field if non-nil, zero value otherwise.

### GetQrCodeDataOk

`func (o *ReceiptOptions) GetQrCodeDataOk() (*string, bool)`

GetQrCodeDataOk returns a tuple with the QrCodeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeData

`func (o *ReceiptOptions) SetQrCodeData(v string)`

SetQrCodeData sets QrCodeData field to given value.

### HasQrCodeData

`func (o *ReceiptOptions) HasQrCodeData() bool`

HasQrCodeData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


