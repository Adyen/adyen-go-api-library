# CheckoutThreeDS2Action

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorisationToken** | Pointer to **string** | A token needed to authorise a payment. | [optional] 
**PaymentData** | Pointer to **string** | A value that must be submitted to the &#x60;/payments/details&#x60; endpoint to verify this payment. | [optional] 
**PaymentMethodType** | Pointer to **string** | Specifies the payment method. | [optional] 
**Subtype** | Pointer to **string** | A subtype of the token. | [optional] 
**Token** | Pointer to **string** | A token to pass to the 3DS2 Component to get the fingerprint. | [optional] 
**Type** | **string** | **threeDS2** | 
**Url** | Pointer to **string** | Specifies the URL to redirect to. | [optional] 

## Methods

### NewCheckoutThreeDS2Action

`func NewCheckoutThreeDS2Action(type_ string, ) *CheckoutThreeDS2Action`

NewCheckoutThreeDS2Action instantiates a new CheckoutThreeDS2Action object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutThreeDS2ActionWithDefaults

`func NewCheckoutThreeDS2ActionWithDefaults() *CheckoutThreeDS2Action`

NewCheckoutThreeDS2ActionWithDefaults instantiates a new CheckoutThreeDS2Action object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorisationToken

`func (o *CheckoutThreeDS2Action) GetAuthorisationToken() string`

GetAuthorisationToken returns the AuthorisationToken field if non-nil, zero value otherwise.

### GetAuthorisationTokenOk

`func (o *CheckoutThreeDS2Action) GetAuthorisationTokenOk() (*string, bool)`

GetAuthorisationTokenOk returns a tuple with the AuthorisationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorisationToken

`func (o *CheckoutThreeDS2Action) SetAuthorisationToken(v string)`

SetAuthorisationToken sets AuthorisationToken field to given value.

### HasAuthorisationToken

`func (o *CheckoutThreeDS2Action) HasAuthorisationToken() bool`

HasAuthorisationToken returns a boolean if a field has been set.

### GetPaymentData

`func (o *CheckoutThreeDS2Action) GetPaymentData() string`

GetPaymentData returns the PaymentData field if non-nil, zero value otherwise.

### GetPaymentDataOk

`func (o *CheckoutThreeDS2Action) GetPaymentDataOk() (*string, bool)`

GetPaymentDataOk returns a tuple with the PaymentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentData

`func (o *CheckoutThreeDS2Action) SetPaymentData(v string)`

SetPaymentData sets PaymentData field to given value.

### HasPaymentData

`func (o *CheckoutThreeDS2Action) HasPaymentData() bool`

HasPaymentData returns a boolean if a field has been set.

### GetPaymentMethodType

`func (o *CheckoutThreeDS2Action) GetPaymentMethodType() string`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *CheckoutThreeDS2Action) GetPaymentMethodTypeOk() (*string, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *CheckoutThreeDS2Action) SetPaymentMethodType(v string)`

SetPaymentMethodType sets PaymentMethodType field to given value.

### HasPaymentMethodType

`func (o *CheckoutThreeDS2Action) HasPaymentMethodType() bool`

HasPaymentMethodType returns a boolean if a field has been set.

### GetSubtype

`func (o *CheckoutThreeDS2Action) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *CheckoutThreeDS2Action) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *CheckoutThreeDS2Action) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *CheckoutThreeDS2Action) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetToken

`func (o *CheckoutThreeDS2Action) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CheckoutThreeDS2Action) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CheckoutThreeDS2Action) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CheckoutThreeDS2Action) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *CheckoutThreeDS2Action) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutThreeDS2Action) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutThreeDS2Action) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *CheckoutThreeDS2Action) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CheckoutThreeDS2Action) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CheckoutThreeDS2Action) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CheckoutThreeDS2Action) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


