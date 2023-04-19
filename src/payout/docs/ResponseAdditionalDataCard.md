# ResponseAdditionalDataCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardBin** | Pointer to **string** | The first six digits of the card number.  This is the [Bank Identification Number (BIN)](https://docs.adyen.com/get-started-with-adyen/payment-glossary#bank-identification-number-bin) for card numbers with a six-digit BIN.  Example: 521234 | [optional] 
**CardHolderName** | Pointer to **string** | The cardholder name passed in the payment request. | [optional] 
**CardIssuingBank** | Pointer to **string** | The bank or the financial institution granting lines of credit through card association branded payment cards. This information can be included when available. | [optional] 
**CardIssuingCountry** | Pointer to **string** | The country where the card was issued.  Example: US | [optional] 
**CardIssuingCurrency** | Pointer to **string** | The currency in which the card is issued, if this information is available. Provided as the currency code or currency number from the ISO-4217 standard.   Example: USD | [optional] 
**CardPaymentMethod** | Pointer to **string** | The card payment method used for the transaction.  Example: amex | [optional] 
**CardSummary** | Pointer to **string** | The last four digits of a card number.  &gt; Returned only in case of a card payment. | [optional] 
**IssuerBin** | Pointer to **string** | The first eight digits of the card number. Only returned if the card number is 16 digits or more.  This is the [Bank Identification Number (BIN)](https://docs.adyen.com/get-started-with-adyen/payment-glossary#bank-identification-number-bin) for card numbers with an eight-digit BIN.  Example: 52123423 | [optional] 

## Methods

### NewResponseAdditionalDataCard

`func NewResponseAdditionalDataCard() *ResponseAdditionalDataCard`

NewResponseAdditionalDataCard instantiates a new ResponseAdditionalDataCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAdditionalDataCardWithDefaults

`func NewResponseAdditionalDataCardWithDefaults() *ResponseAdditionalDataCard`

NewResponseAdditionalDataCardWithDefaults instantiates a new ResponseAdditionalDataCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardBin

`func (o *ResponseAdditionalDataCard) GetCardBin() string`

GetCardBin returns the CardBin field if non-nil, zero value otherwise.

### GetCardBinOk

`func (o *ResponseAdditionalDataCard) GetCardBinOk() (*string, bool)`

GetCardBinOk returns a tuple with the CardBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBin

`func (o *ResponseAdditionalDataCard) SetCardBin(v string)`

SetCardBin sets CardBin field to given value.

### HasCardBin

`func (o *ResponseAdditionalDataCard) HasCardBin() bool`

HasCardBin returns a boolean if a field has been set.

### GetCardHolderName

`func (o *ResponseAdditionalDataCard) GetCardHolderName() string`

GetCardHolderName returns the CardHolderName field if non-nil, zero value otherwise.

### GetCardHolderNameOk

`func (o *ResponseAdditionalDataCard) GetCardHolderNameOk() (*string, bool)`

GetCardHolderNameOk returns a tuple with the CardHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHolderName

`func (o *ResponseAdditionalDataCard) SetCardHolderName(v string)`

SetCardHolderName sets CardHolderName field to given value.

### HasCardHolderName

`func (o *ResponseAdditionalDataCard) HasCardHolderName() bool`

HasCardHolderName returns a boolean if a field has been set.

### GetCardIssuingBank

`func (o *ResponseAdditionalDataCard) GetCardIssuingBank() string`

GetCardIssuingBank returns the CardIssuingBank field if non-nil, zero value otherwise.

### GetCardIssuingBankOk

`func (o *ResponseAdditionalDataCard) GetCardIssuingBankOk() (*string, bool)`

GetCardIssuingBankOk returns a tuple with the CardIssuingBank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIssuingBank

`func (o *ResponseAdditionalDataCard) SetCardIssuingBank(v string)`

SetCardIssuingBank sets CardIssuingBank field to given value.

### HasCardIssuingBank

`func (o *ResponseAdditionalDataCard) HasCardIssuingBank() bool`

HasCardIssuingBank returns a boolean if a field has been set.

### GetCardIssuingCountry

`func (o *ResponseAdditionalDataCard) GetCardIssuingCountry() string`

GetCardIssuingCountry returns the CardIssuingCountry field if non-nil, zero value otherwise.

### GetCardIssuingCountryOk

`func (o *ResponseAdditionalDataCard) GetCardIssuingCountryOk() (*string, bool)`

GetCardIssuingCountryOk returns a tuple with the CardIssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIssuingCountry

`func (o *ResponseAdditionalDataCard) SetCardIssuingCountry(v string)`

SetCardIssuingCountry sets CardIssuingCountry field to given value.

### HasCardIssuingCountry

`func (o *ResponseAdditionalDataCard) HasCardIssuingCountry() bool`

HasCardIssuingCountry returns a boolean if a field has been set.

### GetCardIssuingCurrency

`func (o *ResponseAdditionalDataCard) GetCardIssuingCurrency() string`

GetCardIssuingCurrency returns the CardIssuingCurrency field if non-nil, zero value otherwise.

### GetCardIssuingCurrencyOk

`func (o *ResponseAdditionalDataCard) GetCardIssuingCurrencyOk() (*string, bool)`

GetCardIssuingCurrencyOk returns a tuple with the CardIssuingCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIssuingCurrency

`func (o *ResponseAdditionalDataCard) SetCardIssuingCurrency(v string)`

SetCardIssuingCurrency sets CardIssuingCurrency field to given value.

### HasCardIssuingCurrency

`func (o *ResponseAdditionalDataCard) HasCardIssuingCurrency() bool`

HasCardIssuingCurrency returns a boolean if a field has been set.

### GetCardPaymentMethod

`func (o *ResponseAdditionalDataCard) GetCardPaymentMethod() string`

GetCardPaymentMethod returns the CardPaymentMethod field if non-nil, zero value otherwise.

### GetCardPaymentMethodOk

`func (o *ResponseAdditionalDataCard) GetCardPaymentMethodOk() (*string, bool)`

GetCardPaymentMethodOk returns a tuple with the CardPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardPaymentMethod

`func (o *ResponseAdditionalDataCard) SetCardPaymentMethod(v string)`

SetCardPaymentMethod sets CardPaymentMethod field to given value.

### HasCardPaymentMethod

`func (o *ResponseAdditionalDataCard) HasCardPaymentMethod() bool`

HasCardPaymentMethod returns a boolean if a field has been set.

### GetCardSummary

`func (o *ResponseAdditionalDataCard) GetCardSummary() string`

GetCardSummary returns the CardSummary field if non-nil, zero value otherwise.

### GetCardSummaryOk

`func (o *ResponseAdditionalDataCard) GetCardSummaryOk() (*string, bool)`

GetCardSummaryOk returns a tuple with the CardSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSummary

`func (o *ResponseAdditionalDataCard) SetCardSummary(v string)`

SetCardSummary sets CardSummary field to given value.

### HasCardSummary

`func (o *ResponseAdditionalDataCard) HasCardSummary() bool`

HasCardSummary returns a boolean if a field has been set.

### GetIssuerBin

`func (o *ResponseAdditionalDataCard) GetIssuerBin() string`

GetIssuerBin returns the IssuerBin field if non-nil, zero value otherwise.

### GetIssuerBinOk

`func (o *ResponseAdditionalDataCard) GetIssuerBinOk() (*string, bool)`

GetIssuerBinOk returns a tuple with the IssuerBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerBin

`func (o *ResponseAdditionalDataCard) SetIssuerBin(v string)`

SetIssuerBin sets IssuerBin field to given value.

### HasIssuerBin

`func (o *ResponseAdditionalDataCard) HasIssuerBin() bool`

HasIssuerBin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


