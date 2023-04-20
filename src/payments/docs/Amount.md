# Amount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes). | 
**Value** | **int64** | The amount of the transaction, in [minor units](https://docs.adyen.com/development-resources/currency-codes). | 

## Methods

### NewAmount

`func NewAmount(currency string, value int64, ) *Amount`

NewAmount instantiates a new Amount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmountWithDefaults

`func NewAmountWithDefaults() *Amount`

NewAmountWithDefaults instantiates a new Amount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *Amount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Amount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Amount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetValue

`func (o *Amount) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Amount) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Amount) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


