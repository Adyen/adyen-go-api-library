# Currency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | Surcharge amount per transaction, in [minor units](https://docs.adyen.com/development-resources/currency-codes). | [optional] 
**CurrencyCode** | **string** | Three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes). For example, **AUD**. | 
**Percentage** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewCurrency

`func NewCurrency(currencyCode string, ) *Currency`

NewCurrency instantiates a new Currency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyWithDefaults

`func NewCurrencyWithDefaults() *Currency`

NewCurrencyWithDefaults instantiates a new Currency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Currency) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Currency) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Currency) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Currency) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *Currency) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Currency) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Currency) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetPercentage

`func (o *Currency) GetPercentage() interface{}`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *Currency) GetPercentageOk() (*interface{}, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *Currency) SetPercentage(v interface{})`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *Currency) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### SetPercentageNil

`func (o *Currency) SetPercentageNil(b bool)`

 SetPercentageNil sets the value for Percentage to be an explicit nil

### UnsetPercentage
`func (o *Currency) UnsetPercentage()`

UnsetPercentage ensures that no value is present for Percentage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


