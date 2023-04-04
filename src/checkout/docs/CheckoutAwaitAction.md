# CheckoutAwaitAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentData** | Pointer to **string** | A value that must be submitted to the &#x60;/payments/details&#x60; endpoint to verify this payment. | [optional] 
**PaymentMethodType** | Pointer to **string** | Specifies the payment method. | [optional] 
**Type** | **string** | **await** | 
**Url** | Pointer to **string** | Specifies the URL to redirect to. | [optional] 

## Methods

### NewCheckoutAwaitAction

`func NewCheckoutAwaitAction(type_ string, ) *CheckoutAwaitAction`

NewCheckoutAwaitAction instantiates a new CheckoutAwaitAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutAwaitActionWithDefaults

`func NewCheckoutAwaitActionWithDefaults() *CheckoutAwaitAction`

NewCheckoutAwaitActionWithDefaults instantiates a new CheckoutAwaitAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentData

`func (o *CheckoutAwaitAction) GetPaymentData() string`

GetPaymentData returns the PaymentData field if non-nil, zero value otherwise.

### GetPaymentDataOk

`func (o *CheckoutAwaitAction) GetPaymentDataOk() (*string, bool)`

GetPaymentDataOk returns a tuple with the PaymentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentData

`func (o *CheckoutAwaitAction) SetPaymentData(v string)`

SetPaymentData sets PaymentData field to given value.

### HasPaymentData

`func (o *CheckoutAwaitAction) HasPaymentData() bool`

HasPaymentData returns a boolean if a field has been set.

### GetPaymentMethodType

`func (o *CheckoutAwaitAction) GetPaymentMethodType() string`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *CheckoutAwaitAction) GetPaymentMethodTypeOk() (*string, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *CheckoutAwaitAction) SetPaymentMethodType(v string)`

SetPaymentMethodType sets PaymentMethodType field to given value.

### HasPaymentMethodType

`func (o *CheckoutAwaitAction) HasPaymentMethodType() bool`

HasPaymentMethodType returns a boolean if a field has been set.

### GetType

`func (o *CheckoutAwaitAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutAwaitAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutAwaitAction) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *CheckoutAwaitAction) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CheckoutAwaitAction) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CheckoutAwaitAction) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CheckoutAwaitAction) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


