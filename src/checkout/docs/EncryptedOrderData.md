# EncryptedOrderData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderData** | **string** | The encrypted order data. | 
**PspReference** | **string** | The &#x60;pspReference&#x60; that belongs to the order. | 

## Methods

### NewEncryptedOrderData

`func NewEncryptedOrderData(orderData string, pspReference string, ) *EncryptedOrderData`

NewEncryptedOrderData instantiates a new EncryptedOrderData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptedOrderDataWithDefaults

`func NewEncryptedOrderDataWithDefaults() *EncryptedOrderData`

NewEncryptedOrderDataWithDefaults instantiates a new EncryptedOrderData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderData

`func (o *EncryptedOrderData) GetOrderData() string`

GetOrderData returns the OrderData field if non-nil, zero value otherwise.

### GetOrderDataOk

`func (o *EncryptedOrderData) GetOrderDataOk() (*string, bool)`

GetOrderDataOk returns a tuple with the OrderData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderData

`func (o *EncryptedOrderData) SetOrderData(v string)`

SetOrderData sets OrderData field to given value.


### GetPspReference

`func (o *EncryptedOrderData) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *EncryptedOrderData) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *EncryptedOrderData) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


