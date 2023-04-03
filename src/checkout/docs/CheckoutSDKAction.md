# CheckoutSDKAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentData** | Pointer to **string** | A value that must be submitted to the &#x60;/payments/details&#x60; endpoint to verify this payment. | [optional] 
**PaymentMethodType** | Pointer to **string** | Specifies the payment method. | [optional] 
**SdkData** | Pointer to **map[string]string** | The data to pass to the SDK. | [optional] 
**Type** | **string** | The type of the action. | 
**Url** | Pointer to **string** | Specifies the URL to redirect to. | [optional] 

## Methods

### NewCheckoutSDKAction

`func NewCheckoutSDKAction(type_ string, ) *CheckoutSDKAction`

NewCheckoutSDKAction instantiates a new CheckoutSDKAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSDKActionWithDefaults

`func NewCheckoutSDKActionWithDefaults() *CheckoutSDKAction`

NewCheckoutSDKActionWithDefaults instantiates a new CheckoutSDKAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentData

`func (o *CheckoutSDKAction) GetPaymentData() string`

GetPaymentData returns the PaymentData field if non-nil, zero value otherwise.

### GetPaymentDataOk

`func (o *CheckoutSDKAction) GetPaymentDataOk() (*string, bool)`

GetPaymentDataOk returns a tuple with the PaymentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentData

`func (o *CheckoutSDKAction) SetPaymentData(v string)`

SetPaymentData sets PaymentData field to given value.

### HasPaymentData

`func (o *CheckoutSDKAction) HasPaymentData() bool`

HasPaymentData returns a boolean if a field has been set.

### GetPaymentMethodType

`func (o *CheckoutSDKAction) GetPaymentMethodType() string`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *CheckoutSDKAction) GetPaymentMethodTypeOk() (*string, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *CheckoutSDKAction) SetPaymentMethodType(v string)`

SetPaymentMethodType sets PaymentMethodType field to given value.

### HasPaymentMethodType

`func (o *CheckoutSDKAction) HasPaymentMethodType() bool`

HasPaymentMethodType returns a boolean if a field has been set.

### GetSdkData

`func (o *CheckoutSDKAction) GetSdkData() map[string]string`

GetSdkData returns the SdkData field if non-nil, zero value otherwise.

### GetSdkDataOk

`func (o *CheckoutSDKAction) GetSdkDataOk() (*map[string]string, bool)`

GetSdkDataOk returns a tuple with the SdkData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkData

`func (o *CheckoutSDKAction) SetSdkData(v map[string]string)`

SetSdkData sets SdkData field to given value.

### HasSdkData

`func (o *CheckoutSDKAction) HasSdkData() bool`

HasSdkData returns a boolean if a field has been set.

### GetType

`func (o *CheckoutSDKAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutSDKAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutSDKAction) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *CheckoutSDKAction) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CheckoutSDKAction) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CheckoutSDKAction) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CheckoutSDKAction) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


