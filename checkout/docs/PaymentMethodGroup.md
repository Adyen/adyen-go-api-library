# PaymentMethodGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the group. | [optional] 
**PaymentMethodData** | Pointer to **string** | Echo data to be used if the payment method is displayed as part of this group. | [optional] 
**Type** | Pointer to **string** | The unique code of the group. | [optional] 

## Methods

### NewPaymentMethodGroup

`func NewPaymentMethodGroup() *PaymentMethodGroup`

NewPaymentMethodGroup instantiates a new PaymentMethodGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodGroupWithDefaults

`func NewPaymentMethodGroupWithDefaults() *PaymentMethodGroup`

NewPaymentMethodGroupWithDefaults instantiates a new PaymentMethodGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PaymentMethodGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentMethodGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentMethodGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentMethodGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaymentMethodData

`func (o *PaymentMethodGroup) GetPaymentMethodData() string`

GetPaymentMethodData returns the PaymentMethodData field if non-nil, zero value otherwise.

### GetPaymentMethodDataOk

`func (o *PaymentMethodGroup) GetPaymentMethodDataOk() (*string, bool)`

GetPaymentMethodDataOk returns a tuple with the PaymentMethodData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodData

`func (o *PaymentMethodGroup) SetPaymentMethodData(v string)`

SetPaymentMethodData sets PaymentMethodData field to given value.

### HasPaymentMethodData

`func (o *PaymentMethodGroup) HasPaymentMethodData() bool`

HasPaymentMethodData returns a boolean if a field has been set.

### GetType

`func (o *PaymentMethodGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodGroup) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethodGroup) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


