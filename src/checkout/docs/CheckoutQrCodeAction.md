# CheckoutQrCodeAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **string** | Expiry time of the QR code. | [optional] 
**PaymentData** | Pointer to **string** | A value that must be submitted to the &#x60;/payments/details&#x60; endpoint to verify this payment. | [optional] 
**PaymentMethodType** | Pointer to **string** | Specifies the payment method. | [optional] 
**QrCodeData** | Pointer to **string** | The contents of the QR code as a UTF8 string. | [optional] 
**Type** | **string** | **qrCode** | 
**Url** | Pointer to **string** | Specifies the URL to redirect to. | [optional] 

## Methods

### NewCheckoutQrCodeAction

`func NewCheckoutQrCodeAction(type_ string, ) *CheckoutQrCodeAction`

NewCheckoutQrCodeAction instantiates a new CheckoutQrCodeAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutQrCodeActionWithDefaults

`func NewCheckoutQrCodeActionWithDefaults() *CheckoutQrCodeAction`

NewCheckoutQrCodeActionWithDefaults instantiates a new CheckoutQrCodeAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *CheckoutQrCodeAction) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CheckoutQrCodeAction) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CheckoutQrCodeAction) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CheckoutQrCodeAction) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetPaymentData

`func (o *CheckoutQrCodeAction) GetPaymentData() string`

GetPaymentData returns the PaymentData field if non-nil, zero value otherwise.

### GetPaymentDataOk

`func (o *CheckoutQrCodeAction) GetPaymentDataOk() (*string, bool)`

GetPaymentDataOk returns a tuple with the PaymentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentData

`func (o *CheckoutQrCodeAction) SetPaymentData(v string)`

SetPaymentData sets PaymentData field to given value.

### HasPaymentData

`func (o *CheckoutQrCodeAction) HasPaymentData() bool`

HasPaymentData returns a boolean if a field has been set.

### GetPaymentMethodType

`func (o *CheckoutQrCodeAction) GetPaymentMethodType() string`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *CheckoutQrCodeAction) GetPaymentMethodTypeOk() (*string, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *CheckoutQrCodeAction) SetPaymentMethodType(v string)`

SetPaymentMethodType sets PaymentMethodType field to given value.

### HasPaymentMethodType

`func (o *CheckoutQrCodeAction) HasPaymentMethodType() bool`

HasPaymentMethodType returns a boolean if a field has been set.

### GetQrCodeData

`func (o *CheckoutQrCodeAction) GetQrCodeData() string`

GetQrCodeData returns the QrCodeData field if non-nil, zero value otherwise.

### GetQrCodeDataOk

`func (o *CheckoutQrCodeAction) GetQrCodeDataOk() (*string, bool)`

GetQrCodeDataOk returns a tuple with the QrCodeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeData

`func (o *CheckoutQrCodeAction) SetQrCodeData(v string)`

SetQrCodeData sets QrCodeData field to given value.

### HasQrCodeData

`func (o *CheckoutQrCodeAction) HasQrCodeData() bool`

HasQrCodeData returns a boolean if a field has been set.

### GetType

`func (o *CheckoutQrCodeAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutQrCodeAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutQrCodeAction) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *CheckoutQrCodeAction) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CheckoutQrCodeAction) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CheckoutQrCodeAction) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CheckoutQrCodeAction) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


