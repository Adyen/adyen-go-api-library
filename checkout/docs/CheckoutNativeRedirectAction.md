# CheckoutNativeRedirectAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]string** | When the redirect URL must be accessed via POST, use this data to post to the redirect URL. | [optional] 
**Method** | Pointer to **string** | Specifies the HTTP method, for example GET or POST. | [optional] 
**NativeRedirectData** | Pointer to **string** | Native SDK&#39;s redirect data containing the direct issuer link and state data that must be submitted to the /v1/nativeRedirect/redirectResult. | [optional] 
**PaymentMethodType** | Pointer to **string** | Specifies the payment method. | [optional] 
**Type** | **string** | **nativeRedirect** | 
**Url** | Pointer to **string** | Specifies the URL to redirect to. | [optional] 

## Methods

### NewCheckoutNativeRedirectAction

`func NewCheckoutNativeRedirectAction(type_ string, ) *CheckoutNativeRedirectAction`

NewCheckoutNativeRedirectAction instantiates a new CheckoutNativeRedirectAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutNativeRedirectActionWithDefaults

`func NewCheckoutNativeRedirectActionWithDefaults() *CheckoutNativeRedirectAction`

NewCheckoutNativeRedirectActionWithDefaults instantiates a new CheckoutNativeRedirectAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CheckoutNativeRedirectAction) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CheckoutNativeRedirectAction) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CheckoutNativeRedirectAction) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *CheckoutNativeRedirectAction) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMethod

`func (o *CheckoutNativeRedirectAction) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CheckoutNativeRedirectAction) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CheckoutNativeRedirectAction) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CheckoutNativeRedirectAction) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetNativeRedirectData

`func (o *CheckoutNativeRedirectAction) GetNativeRedirectData() string`

GetNativeRedirectData returns the NativeRedirectData field if non-nil, zero value otherwise.

### GetNativeRedirectDataOk

`func (o *CheckoutNativeRedirectAction) GetNativeRedirectDataOk() (*string, bool)`

GetNativeRedirectDataOk returns a tuple with the NativeRedirectData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeRedirectData

`func (o *CheckoutNativeRedirectAction) SetNativeRedirectData(v string)`

SetNativeRedirectData sets NativeRedirectData field to given value.

### HasNativeRedirectData

`func (o *CheckoutNativeRedirectAction) HasNativeRedirectData() bool`

HasNativeRedirectData returns a boolean if a field has been set.

### GetPaymentMethodType

`func (o *CheckoutNativeRedirectAction) GetPaymentMethodType() string`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *CheckoutNativeRedirectAction) GetPaymentMethodTypeOk() (*string, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *CheckoutNativeRedirectAction) SetPaymentMethodType(v string)`

SetPaymentMethodType sets PaymentMethodType field to given value.

### HasPaymentMethodType

`func (o *CheckoutNativeRedirectAction) HasPaymentMethodType() bool`

HasPaymentMethodType returns a boolean if a field has been set.

### GetType

`func (o *CheckoutNativeRedirectAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutNativeRedirectAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutNativeRedirectAction) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *CheckoutNativeRedirectAction) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CheckoutNativeRedirectAction) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CheckoutNativeRedirectAction) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CheckoutNativeRedirectAction) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


