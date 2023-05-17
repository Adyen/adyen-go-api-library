# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HideMinorUnitsInCurrencies** | Pointer to **[]string** | Hides the minor units for the listed [ISO currency codes](https://en.wikipedia.org/wiki/ISO_4217). | [optional] 

## Methods

### NewPayment

`func NewPayment() *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHideMinorUnitsInCurrencies

`func (o *Payment) GetHideMinorUnitsInCurrencies() []string`

GetHideMinorUnitsInCurrencies returns the HideMinorUnitsInCurrencies field if non-nil, zero value otherwise.

### GetHideMinorUnitsInCurrenciesOk

`func (o *Payment) GetHideMinorUnitsInCurrenciesOk() (*[]string, bool)`

GetHideMinorUnitsInCurrenciesOk returns a tuple with the HideMinorUnitsInCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideMinorUnitsInCurrencies

`func (o *Payment) SetHideMinorUnitsInCurrencies(v []string)`

SetHideMinorUnitsInCurrencies sets HideMinorUnitsInCurrencies field to given value.

### HasHideMinorUnitsInCurrencies

`func (o *Payment) HasHideMinorUnitsInCurrencies() bool`

HasHideMinorUnitsInCurrencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


