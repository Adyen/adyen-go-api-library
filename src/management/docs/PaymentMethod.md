# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplePay** | Pointer to [**ApplePayInfo**](ApplePayInfo.md) |  | [optional] 
**Bcmc** | Pointer to [**BcmcInfo**](BcmcInfo.md) |  | [optional] 
**BusinessLineId** | Pointer to **string** | The unique identifier of the business line. | [optional] 
**CartesBancaires** | Pointer to [**CartesBancairesInfo**](CartesBancairesInfo.md) |  | [optional] 
**Countries** | Pointer to **[]string** | The list of countries where a payment method is available. By default, all countries supported by the payment method. | [optional] 
**Currencies** | Pointer to **[]string** | The list of currencies that a payment method supports. By default, all currencies supported by the payment method. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the payment method is enabled (**true**) or disabled (**false**). | [optional] 
**GiroPay** | Pointer to [**GiroPayInfo**](GiroPayInfo.md) |  | [optional] 
**Id** | **string** | The identifier of the resource. | 
**Klarna** | Pointer to [**KlarnaInfo**](KlarnaInfo.md) |  | [optional] 
**Paypal** | Pointer to [**PayPalInfo**](PayPalInfo.md) |  | [optional] 
**Sofort** | Pointer to [**SofortInfo**](SofortInfo.md) |  | [optional] 
**StoreId** | Pointer to **string** | The ID of the [store](https://docs.adyen.com/api-explorer/#/ManagementService/latest/post/stores__resParam_id), if any. | [optional] 
**Swish** | Pointer to [**SwishInfo**](SwishInfo.md) |  | [optional] 
**Type** | Pointer to **string** | Payment method [variant](https://docs.adyen.com/development-resources/paymentmethodvariant#management-api). | [optional] 

## Methods

### NewPaymentMethod

`func NewPaymentMethod(id string, ) *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplePay

`func (o *PaymentMethod) GetApplePay() ApplePayInfo`

GetApplePay returns the ApplePay field if non-nil, zero value otherwise.

### GetApplePayOk

`func (o *PaymentMethod) GetApplePayOk() (*ApplePayInfo, bool)`

GetApplePayOk returns a tuple with the ApplePay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplePay

`func (o *PaymentMethod) SetApplePay(v ApplePayInfo)`

SetApplePay sets ApplePay field to given value.

### HasApplePay

`func (o *PaymentMethod) HasApplePay() bool`

HasApplePay returns a boolean if a field has been set.

### GetBcmc

`func (o *PaymentMethod) GetBcmc() BcmcInfo`

GetBcmc returns the Bcmc field if non-nil, zero value otherwise.

### GetBcmcOk

`func (o *PaymentMethod) GetBcmcOk() (*BcmcInfo, bool)`

GetBcmcOk returns a tuple with the Bcmc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcmc

`func (o *PaymentMethod) SetBcmc(v BcmcInfo)`

SetBcmc sets Bcmc field to given value.

### HasBcmc

`func (o *PaymentMethod) HasBcmc() bool`

HasBcmc returns a boolean if a field has been set.

### GetBusinessLineId

`func (o *PaymentMethod) GetBusinessLineId() string`

GetBusinessLineId returns the BusinessLineId field if non-nil, zero value otherwise.

### GetBusinessLineIdOk

`func (o *PaymentMethod) GetBusinessLineIdOk() (*string, bool)`

GetBusinessLineIdOk returns a tuple with the BusinessLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessLineId

`func (o *PaymentMethod) SetBusinessLineId(v string)`

SetBusinessLineId sets BusinessLineId field to given value.

### HasBusinessLineId

`func (o *PaymentMethod) HasBusinessLineId() bool`

HasBusinessLineId returns a boolean if a field has been set.

### GetCartesBancaires

`func (o *PaymentMethod) GetCartesBancaires() CartesBancairesInfo`

GetCartesBancaires returns the CartesBancaires field if non-nil, zero value otherwise.

### GetCartesBancairesOk

`func (o *PaymentMethod) GetCartesBancairesOk() (*CartesBancairesInfo, bool)`

GetCartesBancairesOk returns a tuple with the CartesBancaires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartesBancaires

`func (o *PaymentMethod) SetCartesBancaires(v CartesBancairesInfo)`

SetCartesBancaires sets CartesBancaires field to given value.

### HasCartesBancaires

`func (o *PaymentMethod) HasCartesBancaires() bool`

HasCartesBancaires returns a boolean if a field has been set.

### GetCountries

`func (o *PaymentMethod) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *PaymentMethod) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *PaymentMethod) SetCountries(v []string)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *PaymentMethod) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetCurrencies

`func (o *PaymentMethod) GetCurrencies() []string`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *PaymentMethod) GetCurrenciesOk() (*[]string, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *PaymentMethod) SetCurrencies(v []string)`

SetCurrencies sets Currencies field to given value.

### HasCurrencies

`func (o *PaymentMethod) HasCurrencies() bool`

HasCurrencies returns a boolean if a field has been set.

### GetEnabled

`func (o *PaymentMethod) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PaymentMethod) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PaymentMethod) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PaymentMethod) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGiroPay

`func (o *PaymentMethod) GetGiroPay() GiroPayInfo`

GetGiroPay returns the GiroPay field if non-nil, zero value otherwise.

### GetGiroPayOk

`func (o *PaymentMethod) GetGiroPayOk() (*GiroPayInfo, bool)`

GetGiroPayOk returns a tuple with the GiroPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiroPay

`func (o *PaymentMethod) SetGiroPay(v GiroPayInfo)`

SetGiroPay sets GiroPay field to given value.

### HasGiroPay

`func (o *PaymentMethod) HasGiroPay() bool`

HasGiroPay returns a boolean if a field has been set.

### GetId

`func (o *PaymentMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethod) SetId(v string)`

SetId sets Id field to given value.


### GetKlarna

`func (o *PaymentMethod) GetKlarna() KlarnaInfo`

GetKlarna returns the Klarna field if non-nil, zero value otherwise.

### GetKlarnaOk

`func (o *PaymentMethod) GetKlarnaOk() (*KlarnaInfo, bool)`

GetKlarnaOk returns a tuple with the Klarna field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKlarna

`func (o *PaymentMethod) SetKlarna(v KlarnaInfo)`

SetKlarna sets Klarna field to given value.

### HasKlarna

`func (o *PaymentMethod) HasKlarna() bool`

HasKlarna returns a boolean if a field has been set.

### GetPaypal

`func (o *PaymentMethod) GetPaypal() PayPalInfo`

GetPaypal returns the Paypal field if non-nil, zero value otherwise.

### GetPaypalOk

`func (o *PaymentMethod) GetPaypalOk() (*PayPalInfo, bool)`

GetPaypalOk returns a tuple with the Paypal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypal

`func (o *PaymentMethod) SetPaypal(v PayPalInfo)`

SetPaypal sets Paypal field to given value.

### HasPaypal

`func (o *PaymentMethod) HasPaypal() bool`

HasPaypal returns a boolean if a field has been set.

### GetSofort

`func (o *PaymentMethod) GetSofort() SofortInfo`

GetSofort returns the Sofort field if non-nil, zero value otherwise.

### GetSofortOk

`func (o *PaymentMethod) GetSofortOk() (*SofortInfo, bool)`

GetSofortOk returns a tuple with the Sofort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSofort

`func (o *PaymentMethod) SetSofort(v SofortInfo)`

SetSofort sets Sofort field to given value.

### HasSofort

`func (o *PaymentMethod) HasSofort() bool`

HasSofort returns a boolean if a field has been set.

### GetStoreId

`func (o *PaymentMethod) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *PaymentMethod) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *PaymentMethod) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *PaymentMethod) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetSwish

`func (o *PaymentMethod) GetSwish() SwishInfo`

GetSwish returns the Swish field if non-nil, zero value otherwise.

### GetSwishOk

`func (o *PaymentMethod) GetSwishOk() (*SwishInfo, bool)`

GetSwishOk returns a tuple with the Swish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwish

`func (o *PaymentMethod) SetSwish(v SwishInfo)`

SetSwish sets Swish field to given value.

### HasSwish

`func (o *PaymentMethod) HasSwish() bool`

HasSwish returns a boolean if a field has been set.

### GetType

`func (o *PaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


