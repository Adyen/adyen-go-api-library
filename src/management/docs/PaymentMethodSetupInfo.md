# PaymentMethodSetupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplePay** | Pointer to [**ApplePayInfo**](ApplePayInfo.md) |  | [optional] 
**Bcmc** | Pointer to [**BcmcInfo**](BcmcInfo.md) |  | [optional] 
**BusinessLineId** | Pointer to **string** | The unique identifier of the business line. | [optional] 
**CartesBancaires** | Pointer to [**CartesBancairesInfo**](CartesBancairesInfo.md) |  | [optional] 
**Countries** | Pointer to **[]string** | The list of countries where a payment method is available. By default, all countries supported by the payment method. | [optional] 
**Currencies** | Pointer to **[]string** | The list of currencies that a payment method supports. By default, all currencies supported by the payment method. | [optional] 
**CustomRoutingFlags** | Pointer to **[]string** | The list of custom routing flags to route payment to the intended acquirer. | [optional] 
**GiroPay** | Pointer to [**GiroPayInfo**](GiroPayInfo.md) |  | [optional] 
**GooglePay** | Pointer to [**GooglePayInfo**](GooglePayInfo.md) |  | [optional] 
**Klarna** | Pointer to [**KlarnaInfo**](KlarnaInfo.md) |  | [optional] 
**MealVoucherFR** | Pointer to [**MealVoucherFRInfo**](MealVoucherFRInfo.md) |  | [optional] 
**Paypal** | Pointer to [**PayPalInfo**](PayPalInfo.md) |  | [optional] 
**Reference** | Pointer to **string** | Your reference for the payment method. Supported characters a-z, A-Z, 0-9. | [optional] 
**ShopperInteraction** | Pointer to **string** | The sales channel. Required if the merchant account does not have a sales channel. When you provide this field, it overrides the default sales channel set on the merchant account.  Possible values: **eCommerce**, **pos**, **contAuth**, and **moto**.  | [optional] 
**ShopperStatement** | Pointer to [**ShopperStatement**](ShopperStatement.md) |  | [optional] 
**Sofort** | Pointer to [**SofortInfo**](SofortInfo.md) |  | [optional] 
**StoreId** | Pointer to **string** | The ID of the [store](https://docs.adyen.com/api-explorer/#/ManagementService/latest/post/stores__resParam_id), if any. | [optional] 
**Swish** | Pointer to [**SwishInfo**](SwishInfo.md) |  | [optional] 
**Type** | **string** | Payment method [variant](https://docs.adyen.com/development-resources/paymentmethodvariant#management-api). | 

## Methods

### NewPaymentMethodSetupInfo

`func NewPaymentMethodSetupInfo(type_ string, ) *PaymentMethodSetupInfo`

NewPaymentMethodSetupInfo instantiates a new PaymentMethodSetupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodSetupInfoWithDefaults

`func NewPaymentMethodSetupInfoWithDefaults() *PaymentMethodSetupInfo`

NewPaymentMethodSetupInfoWithDefaults instantiates a new PaymentMethodSetupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplePay

`func (o *PaymentMethodSetupInfo) GetApplePay() ApplePayInfo`

GetApplePay returns the ApplePay field if non-nil, zero value otherwise.

### GetApplePayOk

`func (o *PaymentMethodSetupInfo) GetApplePayOk() (*ApplePayInfo, bool)`

GetApplePayOk returns a tuple with the ApplePay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplePay

`func (o *PaymentMethodSetupInfo) SetApplePay(v ApplePayInfo)`

SetApplePay sets ApplePay field to given value.

### HasApplePay

`func (o *PaymentMethodSetupInfo) HasApplePay() bool`

HasApplePay returns a boolean if a field has been set.

### GetBcmc

`func (o *PaymentMethodSetupInfo) GetBcmc() BcmcInfo`

GetBcmc returns the Bcmc field if non-nil, zero value otherwise.

### GetBcmcOk

`func (o *PaymentMethodSetupInfo) GetBcmcOk() (*BcmcInfo, bool)`

GetBcmcOk returns a tuple with the Bcmc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcmc

`func (o *PaymentMethodSetupInfo) SetBcmc(v BcmcInfo)`

SetBcmc sets Bcmc field to given value.

### HasBcmc

`func (o *PaymentMethodSetupInfo) HasBcmc() bool`

HasBcmc returns a boolean if a field has been set.

### GetBusinessLineId

`func (o *PaymentMethodSetupInfo) GetBusinessLineId() string`

GetBusinessLineId returns the BusinessLineId field if non-nil, zero value otherwise.

### GetBusinessLineIdOk

`func (o *PaymentMethodSetupInfo) GetBusinessLineIdOk() (*string, bool)`

GetBusinessLineIdOk returns a tuple with the BusinessLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessLineId

`func (o *PaymentMethodSetupInfo) SetBusinessLineId(v string)`

SetBusinessLineId sets BusinessLineId field to given value.

### HasBusinessLineId

`func (o *PaymentMethodSetupInfo) HasBusinessLineId() bool`

HasBusinessLineId returns a boolean if a field has been set.

### GetCartesBancaires

`func (o *PaymentMethodSetupInfo) GetCartesBancaires() CartesBancairesInfo`

GetCartesBancaires returns the CartesBancaires field if non-nil, zero value otherwise.

### GetCartesBancairesOk

`func (o *PaymentMethodSetupInfo) GetCartesBancairesOk() (*CartesBancairesInfo, bool)`

GetCartesBancairesOk returns a tuple with the CartesBancaires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartesBancaires

`func (o *PaymentMethodSetupInfo) SetCartesBancaires(v CartesBancairesInfo)`

SetCartesBancaires sets CartesBancaires field to given value.

### HasCartesBancaires

`func (o *PaymentMethodSetupInfo) HasCartesBancaires() bool`

HasCartesBancaires returns a boolean if a field has been set.

### GetCountries

`func (o *PaymentMethodSetupInfo) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *PaymentMethodSetupInfo) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *PaymentMethodSetupInfo) SetCountries(v []string)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *PaymentMethodSetupInfo) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetCurrencies

`func (o *PaymentMethodSetupInfo) GetCurrencies() []string`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *PaymentMethodSetupInfo) GetCurrenciesOk() (*[]string, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *PaymentMethodSetupInfo) SetCurrencies(v []string)`

SetCurrencies sets Currencies field to given value.

### HasCurrencies

`func (o *PaymentMethodSetupInfo) HasCurrencies() bool`

HasCurrencies returns a boolean if a field has been set.

### GetCustomRoutingFlags

`func (o *PaymentMethodSetupInfo) GetCustomRoutingFlags() []string`

GetCustomRoutingFlags returns the CustomRoutingFlags field if non-nil, zero value otherwise.

### GetCustomRoutingFlagsOk

`func (o *PaymentMethodSetupInfo) GetCustomRoutingFlagsOk() (*[]string, bool)`

GetCustomRoutingFlagsOk returns a tuple with the CustomRoutingFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoutingFlags

`func (o *PaymentMethodSetupInfo) SetCustomRoutingFlags(v []string)`

SetCustomRoutingFlags sets CustomRoutingFlags field to given value.

### HasCustomRoutingFlags

`func (o *PaymentMethodSetupInfo) HasCustomRoutingFlags() bool`

HasCustomRoutingFlags returns a boolean if a field has been set.

### GetGiroPay

`func (o *PaymentMethodSetupInfo) GetGiroPay() GiroPayInfo`

GetGiroPay returns the GiroPay field if non-nil, zero value otherwise.

### GetGiroPayOk

`func (o *PaymentMethodSetupInfo) GetGiroPayOk() (*GiroPayInfo, bool)`

GetGiroPayOk returns a tuple with the GiroPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiroPay

`func (o *PaymentMethodSetupInfo) SetGiroPay(v GiroPayInfo)`

SetGiroPay sets GiroPay field to given value.

### HasGiroPay

`func (o *PaymentMethodSetupInfo) HasGiroPay() bool`

HasGiroPay returns a boolean if a field has been set.

### GetGooglePay

`func (o *PaymentMethodSetupInfo) GetGooglePay() GooglePayInfo`

GetGooglePay returns the GooglePay field if non-nil, zero value otherwise.

### GetGooglePayOk

`func (o *PaymentMethodSetupInfo) GetGooglePayOk() (*GooglePayInfo, bool)`

GetGooglePayOk returns a tuple with the GooglePay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePay

`func (o *PaymentMethodSetupInfo) SetGooglePay(v GooglePayInfo)`

SetGooglePay sets GooglePay field to given value.

### HasGooglePay

`func (o *PaymentMethodSetupInfo) HasGooglePay() bool`

HasGooglePay returns a boolean if a field has been set.

### GetKlarna

`func (o *PaymentMethodSetupInfo) GetKlarna() KlarnaInfo`

GetKlarna returns the Klarna field if non-nil, zero value otherwise.

### GetKlarnaOk

`func (o *PaymentMethodSetupInfo) GetKlarnaOk() (*KlarnaInfo, bool)`

GetKlarnaOk returns a tuple with the Klarna field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKlarna

`func (o *PaymentMethodSetupInfo) SetKlarna(v KlarnaInfo)`

SetKlarna sets Klarna field to given value.

### HasKlarna

`func (o *PaymentMethodSetupInfo) HasKlarna() bool`

HasKlarna returns a boolean if a field has been set.

### GetMealVoucherFR

`func (o *PaymentMethodSetupInfo) GetMealVoucherFR() MealVoucherFRInfo`

GetMealVoucherFR returns the MealVoucherFR field if non-nil, zero value otherwise.

### GetMealVoucherFROk

`func (o *PaymentMethodSetupInfo) GetMealVoucherFROk() (*MealVoucherFRInfo, bool)`

GetMealVoucherFROk returns a tuple with the MealVoucherFR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMealVoucherFR

`func (o *PaymentMethodSetupInfo) SetMealVoucherFR(v MealVoucherFRInfo)`

SetMealVoucherFR sets MealVoucherFR field to given value.

### HasMealVoucherFR

`func (o *PaymentMethodSetupInfo) HasMealVoucherFR() bool`

HasMealVoucherFR returns a boolean if a field has been set.

### GetPaypal

`func (o *PaymentMethodSetupInfo) GetPaypal() PayPalInfo`

GetPaypal returns the Paypal field if non-nil, zero value otherwise.

### GetPaypalOk

`func (o *PaymentMethodSetupInfo) GetPaypalOk() (*PayPalInfo, bool)`

GetPaypalOk returns a tuple with the Paypal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypal

`func (o *PaymentMethodSetupInfo) SetPaypal(v PayPalInfo)`

SetPaypal sets Paypal field to given value.

### HasPaypal

`func (o *PaymentMethodSetupInfo) HasPaypal() bool`

HasPaypal returns a boolean if a field has been set.

### GetReference

`func (o *PaymentMethodSetupInfo) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentMethodSetupInfo) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentMethodSetupInfo) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaymentMethodSetupInfo) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetShopperInteraction

`func (o *PaymentMethodSetupInfo) GetShopperInteraction() string`

GetShopperInteraction returns the ShopperInteraction field if non-nil, zero value otherwise.

### GetShopperInteractionOk

`func (o *PaymentMethodSetupInfo) GetShopperInteractionOk() (*string, bool)`

GetShopperInteractionOk returns a tuple with the ShopperInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperInteraction

`func (o *PaymentMethodSetupInfo) SetShopperInteraction(v string)`

SetShopperInteraction sets ShopperInteraction field to given value.

### HasShopperInteraction

`func (o *PaymentMethodSetupInfo) HasShopperInteraction() bool`

HasShopperInteraction returns a boolean if a field has been set.

### GetShopperStatement

`func (o *PaymentMethodSetupInfo) GetShopperStatement() ShopperStatement`

GetShopperStatement returns the ShopperStatement field if non-nil, zero value otherwise.

### GetShopperStatementOk

`func (o *PaymentMethodSetupInfo) GetShopperStatementOk() (*ShopperStatement, bool)`

GetShopperStatementOk returns a tuple with the ShopperStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperStatement

`func (o *PaymentMethodSetupInfo) SetShopperStatement(v ShopperStatement)`

SetShopperStatement sets ShopperStatement field to given value.

### HasShopperStatement

`func (o *PaymentMethodSetupInfo) HasShopperStatement() bool`

HasShopperStatement returns a boolean if a field has been set.

### GetSofort

`func (o *PaymentMethodSetupInfo) GetSofort() SofortInfo`

GetSofort returns the Sofort field if non-nil, zero value otherwise.

### GetSofortOk

`func (o *PaymentMethodSetupInfo) GetSofortOk() (*SofortInfo, bool)`

GetSofortOk returns a tuple with the Sofort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSofort

`func (o *PaymentMethodSetupInfo) SetSofort(v SofortInfo)`

SetSofort sets Sofort field to given value.

### HasSofort

`func (o *PaymentMethodSetupInfo) HasSofort() bool`

HasSofort returns a boolean if a field has been set.

### GetStoreId

`func (o *PaymentMethodSetupInfo) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *PaymentMethodSetupInfo) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *PaymentMethodSetupInfo) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *PaymentMethodSetupInfo) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetSwish

`func (o *PaymentMethodSetupInfo) GetSwish() SwishInfo`

GetSwish returns the Swish field if non-nil, zero value otherwise.

### GetSwishOk

`func (o *PaymentMethodSetupInfo) GetSwishOk() (*SwishInfo, bool)`

GetSwishOk returns a tuple with the Swish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwish

`func (o *PaymentMethodSetupInfo) SetSwish(v SwishInfo)`

SetSwish sets Swish field to given value.

### HasSwish

`func (o *PaymentMethodSetupInfo) HasSwish() bool`

HasSwish returns a boolean if a field has been set.

### GetType

`func (o *PaymentMethodSetupInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodSetupInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodSetupInfo) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


