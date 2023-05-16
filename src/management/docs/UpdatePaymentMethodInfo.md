# UpdatePaymentMethodInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Countries** | Pointer to **[]string** | The list of countries where a payment method is available. By default, all countries supported by the payment method. | [optional] 
**Currencies** | Pointer to **[]string** | The list of currencies that a payment method supports. By default, all currencies supported by the payment method. | [optional] 
**CustomRoutingFlags** | Pointer to **[]string** | Custom routing flags for acquirer routing. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the payment method is enabled (**true**) or disabled (**false**). | [optional] 
**ShopperStatement** | Pointer to [**ShopperStatement**](ShopperStatement.md) |  | [optional] 
**StoreIds** | Pointer to **[]string** | The list of stores for this payment method | [optional] 

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

### GetCustomRoutingFlags

`func (o *UpdatePaymentMethodInfo) GetCustomRoutingFlags() []string`

GetCustomRoutingFlags returns the CustomRoutingFlags field if non-nil, zero value otherwise.

### GetCustomRoutingFlagsOk

`func (o *UpdatePaymentMethodInfo) GetCustomRoutingFlagsOk() (*[]string, bool)`

GetCustomRoutingFlagsOk returns a tuple with the CustomRoutingFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoutingFlags

`func (o *UpdatePaymentMethodInfo) SetCustomRoutingFlags(v []string)`

SetCustomRoutingFlags sets CustomRoutingFlags field to given value.

### HasCustomRoutingFlags

`func (o *UpdatePaymentMethodInfo) HasCustomRoutingFlags() bool`

HasCustomRoutingFlags returns a boolean if a field has been set.

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

### GetShopperStatement

`func (o *UpdatePaymentMethodInfo) GetShopperStatement() ShopperStatement`

GetShopperStatement returns the ShopperStatement field if non-nil, zero value otherwise.

### GetShopperStatementOk

`func (o *UpdatePaymentMethodInfo) GetShopperStatementOk() (*ShopperStatement, bool)`

GetShopperStatementOk returns a tuple with the ShopperStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperStatement

`func (o *UpdatePaymentMethodInfo) SetShopperStatement(v ShopperStatement)`

SetShopperStatement sets ShopperStatement field to given value.

### HasShopperStatement

`func (o *UpdatePaymentMethodInfo) HasShopperStatement() bool`

HasShopperStatement returns a boolean if a field has been set.

### GetStoreIds

`func (o *UpdatePaymentMethodInfo) GetStoreIds() []string`

GetStoreIds returns the StoreIds field if non-nil, zero value otherwise.

### GetStoreIdsOk

`func (o *UpdatePaymentMethodInfo) GetStoreIdsOk() (*[]string, bool)`

GetStoreIdsOk returns a tuple with the StoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreIds

`func (o *UpdatePaymentMethodInfo) SetStoreIds(v []string)`

SetStoreIds sets StoreIds field to given value.

### HasStoreIds

`func (o *UpdatePaymentMethodInfo) HasStoreIds() bool`

HasStoreIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


