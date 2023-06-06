# BalanceMutation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | Pointer to **int64** | The amount in the payment&#39;s currency that is debited or credited on the balance accounting register. | [optional] 
**Currency** | Pointer to **string** | The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes). | [optional] 
**Received** | Pointer to **int64** | The amount in the payment&#39;s currency that is debited or credited on the received accounting register. | [optional] 
**Reserved** | Pointer to **int64** | The amount in the payment&#39;s currency that is debited or credited on the reserved accounting register. | [optional] 

## Methods

### NewBalanceMutation

`func NewBalanceMutation() *BalanceMutation`

NewBalanceMutation instantiates a new BalanceMutation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceMutationWithDefaults

`func NewBalanceMutationWithDefaults() *BalanceMutation`

NewBalanceMutationWithDefaults instantiates a new BalanceMutation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *BalanceMutation) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BalanceMutation) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BalanceMutation) SetBalance(v int64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *BalanceMutation) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCurrency

`func (o *BalanceMutation) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BalanceMutation) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BalanceMutation) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BalanceMutation) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetReceived

`func (o *BalanceMutation) GetReceived() int64`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *BalanceMutation) GetReceivedOk() (*int64, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *BalanceMutation) SetReceived(v int64)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *BalanceMutation) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetReserved

`func (o *BalanceMutation) GetReserved() int64`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *BalanceMutation) GetReservedOk() (*int64, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *BalanceMutation) SetReserved(v int64)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *BalanceMutation) HasReserved() bool`

HasReserved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


