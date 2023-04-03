# CheckoutOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderData** | **string** | The encrypted order data. | 
**PspReference** | **string** | The &#x60;pspReference&#x60; that belongs to the order. | 

## Methods

### NewCheckoutOrder

`func NewCheckoutOrder(orderData string, pspReference string, ) *CheckoutOrder`

NewCheckoutOrder instantiates a new CheckoutOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutOrderWithDefaults

`func NewCheckoutOrderWithDefaults() *CheckoutOrder`

NewCheckoutOrderWithDefaults instantiates a new CheckoutOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderData

`func (o *CheckoutOrder) GetOrderData() string`

GetOrderData returns the OrderData field if non-nil, zero value otherwise.

### GetOrderDataOk

`func (o *CheckoutOrder) GetOrderDataOk() (*string, bool)`

GetOrderDataOk returns a tuple with the OrderData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderData

`func (o *CheckoutOrder) SetOrderData(v string)`

SetOrderData sets OrderData field to given value.


### GetPspReference

`func (o *CheckoutOrder) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *CheckoutOrder) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *CheckoutOrder) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


