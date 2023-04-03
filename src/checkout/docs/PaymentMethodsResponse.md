# PaymentMethodsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethods** | Pointer to [**[]PaymentMethod**](PaymentMethod.md) | Detailed list of payment methods required to generate payment forms. | [optional] 
**StoredPaymentMethods** | Pointer to [**[]StoredPaymentMethod**](StoredPaymentMethod.md) | List of all stored payment methods. | [optional] 

## Methods

### NewPaymentMethodsResponse

`func NewPaymentMethodsResponse() *PaymentMethodsResponse`

NewPaymentMethodsResponse instantiates a new PaymentMethodsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodsResponseWithDefaults

`func NewPaymentMethodsResponseWithDefaults() *PaymentMethodsResponse`

NewPaymentMethodsResponseWithDefaults instantiates a new PaymentMethodsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethods

`func (o *PaymentMethodsResponse) GetPaymentMethods() []PaymentMethod`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *PaymentMethodsResponse) GetPaymentMethodsOk() (*[]PaymentMethod, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *PaymentMethodsResponse) SetPaymentMethods(v []PaymentMethod)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *PaymentMethodsResponse) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetStoredPaymentMethods

`func (o *PaymentMethodsResponse) GetStoredPaymentMethods() []StoredPaymentMethod`

GetStoredPaymentMethods returns the StoredPaymentMethods field if non-nil, zero value otherwise.

### GetStoredPaymentMethodsOk

`func (o *PaymentMethodsResponse) GetStoredPaymentMethodsOk() (*[]StoredPaymentMethod, bool)`

GetStoredPaymentMethodsOk returns a tuple with the StoredPaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredPaymentMethods

`func (o *PaymentMethodsResponse) SetStoredPaymentMethods(v []StoredPaymentMethod)`

SetStoredPaymentMethods sets StoredPaymentMethods field to given value.

### HasStoredPaymentMethods

`func (o *PaymentMethodsResponse) HasStoredPaymentMethods() bool`

HasStoredPaymentMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


