# Balance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **int64** | The remaining amount available for spending. | 
**Balance** | **int64** | The total amount in the balance. | 
**Currency** | **string** | The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes) of the balance. | 
**Reserved** | **int64** | The amount reserved for payments that have been authorised, but have not been captured yet. | 

## Methods

### NewBalance

`func NewBalance(available int64, balance int64, currency string, reserved int64, ) *Balance`

NewBalance instantiates a new Balance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceWithDefaults

`func NewBalanceWithDefaults() *Balance`

NewBalanceWithDefaults instantiates a new Balance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *Balance) GetAvailable() int64`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *Balance) GetAvailableOk() (*int64, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *Balance) SetAvailable(v int64)`

SetAvailable sets Available field to given value.


### GetBalance

`func (o *Balance) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Balance) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Balance) SetBalance(v int64)`

SetBalance sets Balance field to given value.


### GetCurrency

`func (o *Balance) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Balance) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Balance) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReserved

`func (o *Balance) GetReserved() int64`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *Balance) GetReservedOk() (*int64, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *Balance) SetReserved(v int64)`

SetReserved sets Reserved field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


