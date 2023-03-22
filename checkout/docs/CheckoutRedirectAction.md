# CheckoutRedirectAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]string** | When the redirect URL must be accessed via POST, use this data to post to the redirect URL. | [optional] 
**Method** | Pointer to **string** | Specifies the HTTP method, for example GET or POST. | [optional] 
**PaymentMethodType** | Pointer to **string** | Specifies the payment method. | [optional] 
**Type** | **string** | **redirect** | 
**Url** | Pointer to **string** | Specifies the URL to redirect to. | [optional] 

## Methods

### NewCheckoutRedirectAction

`func NewCheckoutRedirectAction(type_ string, ) *CheckoutRedirectAction`

NewCheckoutRedirectAction instantiates a new CheckoutRedirectAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutRedirectActionWithDefaults

`func NewCheckoutRedirectActionWithDefaults() *CheckoutRedirectAction`

NewCheckoutRedirectActionWithDefaults instantiates a new CheckoutRedirectAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CheckoutRedirectAction) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CheckoutRedirectAction) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CheckoutRedirectAction) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *CheckoutRedirectAction) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMethod

`func (o *CheckoutRedirectAction) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CheckoutRedirectAction) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CheckoutRedirectAction) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CheckoutRedirectAction) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPaymentMethodType

`func (o *CheckoutRedirectAction) GetPaymentMethodType() string`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *CheckoutRedirectAction) GetPaymentMethodTypeOk() (*string, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *CheckoutRedirectAction) SetPaymentMethodType(v string)`

SetPaymentMethodType sets PaymentMethodType field to given value.

### HasPaymentMethodType

`func (o *CheckoutRedirectAction) HasPaymentMethodType() bool`

HasPaymentMethodType returns a boolean if a field has been set.

### GetType

`func (o *CheckoutRedirectAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutRedirectAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutRedirectAction) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *CheckoutRedirectAction) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CheckoutRedirectAction) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CheckoutRedirectAction) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CheckoutRedirectAction) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


