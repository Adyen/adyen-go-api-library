# CheckoutCancelOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantAccount** | **string** | The merchant account identifier that orderData belongs to. | 
**Order** | [**EncryptedOrderData**](EncryptedOrderData.md) |  | 

## Methods

### NewCheckoutCancelOrderRequest

`func NewCheckoutCancelOrderRequest(merchantAccount string, order EncryptedOrderData, ) *CheckoutCancelOrderRequest`

NewCheckoutCancelOrderRequest instantiates a new CheckoutCancelOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutCancelOrderRequestWithDefaults

`func NewCheckoutCancelOrderRequestWithDefaults() *CheckoutCancelOrderRequest`

NewCheckoutCancelOrderRequestWithDefaults instantiates a new CheckoutCancelOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantAccount

`func (o *CheckoutCancelOrderRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *CheckoutCancelOrderRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *CheckoutCancelOrderRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetOrder

`func (o *CheckoutCancelOrderRequest) GetOrder() EncryptedOrderData`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CheckoutCancelOrderRequest) GetOrderOk() (*EncryptedOrderData, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CheckoutCancelOrderRequest) SetOrder(v EncryptedOrderData)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


