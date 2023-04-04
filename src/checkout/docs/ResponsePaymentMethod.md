# ResponsePaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** | The card brand that the shopper used to pay. Only returned if &#x60;paymentMethod.type&#x60; is **scheme**. | [optional] 
**Type** | Pointer to **string** | The &#x60;paymentMethod.type&#x60; value used in the request. | [optional] 

## Methods

### NewResponsePaymentMethod

`func NewResponsePaymentMethod() *ResponsePaymentMethod`

NewResponsePaymentMethod instantiates a new ResponsePaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePaymentMethodWithDefaults

`func NewResponsePaymentMethodWithDefaults() *ResponsePaymentMethod`

NewResponsePaymentMethodWithDefaults instantiates a new ResponsePaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *ResponsePaymentMethod) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ResponsePaymentMethod) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ResponsePaymentMethod) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ResponsePaymentMethod) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetType

`func (o *ResponsePaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsePaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsePaymentMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsePaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


