# Gratuity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCustomAmount** | Pointer to **bool** | Indicates whether one of the predefined tipping options is to let the shopper enter a custom tip. If **true**, only three of the other options defined in &#x60;predefinedTipEntries&#x60; are shown. | [optional] 
**Currency** | Pointer to **string** | The currency that the tipping settings apply to. | [optional] 
**PredefinedTipEntries** | Pointer to **[]string** | Tipping options the shopper can choose from if &#x60;usePredefinedTipEntries&#x60; is **true**. The maximum number of predefined options is four, or three plus the option to enter a custom tip. The options can be a mix of:  - A percentage of the transaction amount. Example: **5%** - A tip amount in minor units. Example: **500** for a EUR 5 tip. | [optional] 
**UsePredefinedTipEntries** | Pointer to **bool** | Indicates whether the terminal shows a prompt to enter a tip (**false**), or predefined tipping options to choose from (**true**). | [optional] 

## Methods

### NewGratuity

`func NewGratuity() *Gratuity`

NewGratuity instantiates a new Gratuity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGratuityWithDefaults

`func NewGratuityWithDefaults() *Gratuity`

NewGratuityWithDefaults instantiates a new Gratuity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCustomAmount

`func (o *Gratuity) GetAllowCustomAmount() bool`

GetAllowCustomAmount returns the AllowCustomAmount field if non-nil, zero value otherwise.

### GetAllowCustomAmountOk

`func (o *Gratuity) GetAllowCustomAmountOk() (*bool, bool)`

GetAllowCustomAmountOk returns a tuple with the AllowCustomAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomAmount

`func (o *Gratuity) SetAllowCustomAmount(v bool)`

SetAllowCustomAmount sets AllowCustomAmount field to given value.

### HasAllowCustomAmount

`func (o *Gratuity) HasAllowCustomAmount() bool`

HasAllowCustomAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *Gratuity) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Gratuity) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Gratuity) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Gratuity) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPredefinedTipEntries

`func (o *Gratuity) GetPredefinedTipEntries() []string`

GetPredefinedTipEntries returns the PredefinedTipEntries field if non-nil, zero value otherwise.

### GetPredefinedTipEntriesOk

`func (o *Gratuity) GetPredefinedTipEntriesOk() (*[]string, bool)`

GetPredefinedTipEntriesOk returns a tuple with the PredefinedTipEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedTipEntries

`func (o *Gratuity) SetPredefinedTipEntries(v []string)`

SetPredefinedTipEntries sets PredefinedTipEntries field to given value.

### HasPredefinedTipEntries

`func (o *Gratuity) HasPredefinedTipEntries() bool`

HasPredefinedTipEntries returns a boolean if a field has been set.

### GetUsePredefinedTipEntries

`func (o *Gratuity) GetUsePredefinedTipEntries() bool`

GetUsePredefinedTipEntries returns the UsePredefinedTipEntries field if non-nil, zero value otherwise.

### GetUsePredefinedTipEntriesOk

`func (o *Gratuity) GetUsePredefinedTipEntriesOk() (*bool, bool)`

GetUsePredefinedTipEntriesOk returns a tuple with the UsePredefinedTipEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePredefinedTipEntries

`func (o *Gratuity) SetUsePredefinedTipEntries(v bool)`

SetUsePredefinedTipEntries sets UsePredefinedTipEntries field to given value.

### HasUsePredefinedTipEntries

`func (o *Gratuity) HasUsePredefinedTipEntries() bool`

HasUsePredefinedTipEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


