# MinorUnitsMonetaryValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  Amount of this monetary value, in minor units | [optional] 
**CurrencyCode** | Pointer to **string** |  Currency of this monetary value, Format: [ISO currency code](https://docs.adyen.com/development-resources/currency-codes). | [optional] 

## Methods

### NewMinorUnitsMonetaryValue

`func NewMinorUnitsMonetaryValue() *MinorUnitsMonetaryValue`

NewMinorUnitsMonetaryValue instantiates a new MinorUnitsMonetaryValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinorUnitsMonetaryValueWithDefaults

`func NewMinorUnitsMonetaryValueWithDefaults() *MinorUnitsMonetaryValue`

NewMinorUnitsMonetaryValueWithDefaults instantiates a new MinorUnitsMonetaryValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *MinorUnitsMonetaryValue) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *MinorUnitsMonetaryValue) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *MinorUnitsMonetaryValue) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *MinorUnitsMonetaryValue) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *MinorUnitsMonetaryValue) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *MinorUnitsMonetaryValue) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *MinorUnitsMonetaryValue) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *MinorUnitsMonetaryValue) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


