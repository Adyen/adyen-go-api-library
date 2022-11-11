# UpdatePaymentMethodInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Countries** | Pointer to **[]string** | The list of countries where a payment method is available. By default, all countries supported by the payment method. | [optional] 
**Currencies** | Pointer to **[]string** | The list of currencies that a payment method supports. By default, all currencies supported by the payment method. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the payment method is enabled (**true**) or disabled (**false**). | [optional] 

## Methods

### NewUpdatePaymentMethodInfo

`func NewUpdatePaymentMethodInfo() *UpdatePaymentMethodInfo`

NewUpdatePaymentMethodInfo instantiates a new UpdatePaymentMethodInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePaymentMethodInfoWithDefaults

`func NewUpdatePaymentMethodInfoWithDefaults() *UpdatePaymentMethodInfo`

NewUpdatePaymentMethodInfoWithDefaults instantiates a new UpdatePaymentMethodInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountries

`func (o *UpdatePaymentMethodInfo) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *UpdatePaymentMethodInfo) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *UpdatePaymentMethodInfo) SetCountries(v []string)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *UpdatePaymentMethodInfo) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetCurrencies

`func (o *UpdatePaymentMethodInfo) GetCurrencies() []string`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *UpdatePaymentMethodInfo) GetCurrenciesOk() (*[]string, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *UpdatePaymentMethodInfo) SetCurrencies(v []string)`

SetCurrencies sets Currencies field to given value.

### HasCurrencies

`func (o *UpdatePaymentMethodInfo) HasCurrencies() bool`

HasCurrencies returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdatePaymentMethodInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdatePaymentMethodInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdatePaymentMethodInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdatePaymentMethodInfo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


