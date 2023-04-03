# CardDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardNumber** | **string** | A minimum of the first 8 digits of the card number and a maximum of the full card number. 11 digits gives the best result.   You must be [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide) to collect raw card data. | 
**CountryCode** | Pointer to **string** | The shopper country.  Format: [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) Example: NL or DE | [optional] 
**EncryptedCardNumber** | Pointer to **string** | The encrypted card number. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**SupportedBrands** | Pointer to **[]string** | The card brands you support. This is the [&#x60;brands&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/paymentMethods__resParam_paymentMethods-brands) array from your [&#x60;/paymentMethods&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/paymentMethods) response.   If not included, our API uses the ones configured for your merchant account and, if provided, the country code. | [optional] 

## Methods

### NewCardDetailsRequest

`func NewCardDetailsRequest(cardNumber string, merchantAccount string, ) *CardDetailsRequest`

NewCardDetailsRequest instantiates a new CardDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardDetailsRequestWithDefaults

`func NewCardDetailsRequestWithDefaults() *CardDetailsRequest`

NewCardDetailsRequestWithDefaults instantiates a new CardDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardNumber

`func (o *CardDetailsRequest) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CardDetailsRequest) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CardDetailsRequest) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.


### GetCountryCode

`func (o *CardDetailsRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CardDetailsRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CardDetailsRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *CardDetailsRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetEncryptedCardNumber

`func (o *CardDetailsRequest) GetEncryptedCardNumber() string`

GetEncryptedCardNumber returns the EncryptedCardNumber field if non-nil, zero value otherwise.

### GetEncryptedCardNumberOk

`func (o *CardDetailsRequest) GetEncryptedCardNumberOk() (*string, bool)`

GetEncryptedCardNumberOk returns a tuple with the EncryptedCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedCardNumber

`func (o *CardDetailsRequest) SetEncryptedCardNumber(v string)`

SetEncryptedCardNumber sets EncryptedCardNumber field to given value.

### HasEncryptedCardNumber

`func (o *CardDetailsRequest) HasEncryptedCardNumber() bool`

HasEncryptedCardNumber returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *CardDetailsRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *CardDetailsRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *CardDetailsRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetSupportedBrands

`func (o *CardDetailsRequest) GetSupportedBrands() []string`

GetSupportedBrands returns the SupportedBrands field if non-nil, zero value otherwise.

### GetSupportedBrandsOk

`func (o *CardDetailsRequest) GetSupportedBrandsOk() (*[]string, bool)`

GetSupportedBrandsOk returns a tuple with the SupportedBrands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedBrands

`func (o *CardDetailsRequest) SetSupportedBrands(v []string)`

SetSupportedBrands sets SupportedBrands field to given value.

### HasSupportedBrands

`func (o *CardDetailsRequest) HasSupportedBrands() bool`

HasSupportedBrands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


