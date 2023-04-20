# CardBin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bin** | Pointer to **string** | The first 6 digit of the card number. Enable this field via merchant account settings. | [optional] 
**Commercial** | Pointer to **bool** | If true, it indicates a commercial card. Enable this field via merchant account settings. | [optional] 
**FundingSource** | Pointer to **string** | The card funding source. Valid values are: * CHARGE * CREDIT * DEBIT * DEFERRED_DEBIT * PREPAID * PREPAID_RELOADABLE * PREPAID_NONRELOADABLE &gt; Enable this field via merchant account settings. | [optional] 
**FundsAvailability** | Pointer to **string** | Indicates availability of funds.  Visa: * \&quot;I\&quot; (fast funds are supported) * \&quot;N\&quot; (otherwise)  Mastercard: * \&quot;I\&quot; (product type is Prepaid or Debit, or issuing country is in CEE/HGEM list) * \&quot;N\&quot; (otherwise) &gt; Returned when you verify a card BIN or estimate costs, and only if &#x60;payoutEligible&#x60; is different from \&quot;N\&quot; or \&quot;U\&quot;. | [optional] 
**IssuerBin** | Pointer to **string** | The first 8 digit of the card number. Enable this field via merchant account settings. | [optional] 
**IssuingBank** | Pointer to **string** | The issuing bank of the card. | [optional] 
**IssuingCountry** | Pointer to **string** | The country where the card was issued from. | [optional] 
**IssuingCurrency** | Pointer to **string** | The currency of the card. | [optional] 
**PaymentMethod** | Pointer to **string** | The payment method associated with the card (e.g. visa, mc, or amex). | [optional] 
**PayoutEligible** | Pointer to **string** | Indicates whether a payout is eligible or not for this card.  Visa: * \&quot;Y\&quot; * \&quot;N\&quot;  Mastercard: * \&quot;Y\&quot; (domestic and cross-border) * \&quot;D\&quot; (only domestic) * \&quot;N\&quot; (no MoneySend) * \&quot;U\&quot; (unknown) &gt; Returned when you verify a card BIN or estimate costs, and only if &#x60;payoutEligible&#x60; is different from \&quot;N\&quot; or \&quot;U\&quot;. | [optional] 
**Summary** | Pointer to **string** | The last four digits of the card number. | [optional] 

## Methods

### NewCardBin

`func NewCardBin() *CardBin`

NewCardBin instantiates a new CardBin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardBinWithDefaults

`func NewCardBinWithDefaults() *CardBin`

NewCardBinWithDefaults instantiates a new CardBin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBin

`func (o *CardBin) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *CardBin) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *CardBin) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *CardBin) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetCommercial

`func (o *CardBin) GetCommercial() bool`

GetCommercial returns the Commercial field if non-nil, zero value otherwise.

### GetCommercialOk

`func (o *CardBin) GetCommercialOk() (*bool, bool)`

GetCommercialOk returns a tuple with the Commercial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercial

`func (o *CardBin) SetCommercial(v bool)`

SetCommercial sets Commercial field to given value.

### HasCommercial

`func (o *CardBin) HasCommercial() bool`

HasCommercial returns a boolean if a field has been set.

### GetFundingSource

`func (o *CardBin) GetFundingSource() string`

GetFundingSource returns the FundingSource field if non-nil, zero value otherwise.

### GetFundingSourceOk

`func (o *CardBin) GetFundingSourceOk() (*string, bool)`

GetFundingSourceOk returns a tuple with the FundingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSource

`func (o *CardBin) SetFundingSource(v string)`

SetFundingSource sets FundingSource field to given value.

### HasFundingSource

`func (o *CardBin) HasFundingSource() bool`

HasFundingSource returns a boolean if a field has been set.

### GetFundsAvailability

`func (o *CardBin) GetFundsAvailability() string`

GetFundsAvailability returns the FundsAvailability field if non-nil, zero value otherwise.

### GetFundsAvailabilityOk

`func (o *CardBin) GetFundsAvailabilityOk() (*string, bool)`

GetFundsAvailabilityOk returns a tuple with the FundsAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundsAvailability

`func (o *CardBin) SetFundsAvailability(v string)`

SetFundsAvailability sets FundsAvailability field to given value.

### HasFundsAvailability

`func (o *CardBin) HasFundsAvailability() bool`

HasFundsAvailability returns a boolean if a field has been set.

### GetIssuerBin

`func (o *CardBin) GetIssuerBin() string`

GetIssuerBin returns the IssuerBin field if non-nil, zero value otherwise.

### GetIssuerBinOk

`func (o *CardBin) GetIssuerBinOk() (*string, bool)`

GetIssuerBinOk returns a tuple with the IssuerBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerBin

`func (o *CardBin) SetIssuerBin(v string)`

SetIssuerBin sets IssuerBin field to given value.

### HasIssuerBin

`func (o *CardBin) HasIssuerBin() bool`

HasIssuerBin returns a boolean if a field has been set.

### GetIssuingBank

`func (o *CardBin) GetIssuingBank() string`

GetIssuingBank returns the IssuingBank field if non-nil, zero value otherwise.

### GetIssuingBankOk

`func (o *CardBin) GetIssuingBankOk() (*string, bool)`

GetIssuingBankOk returns a tuple with the IssuingBank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingBank

`func (o *CardBin) SetIssuingBank(v string)`

SetIssuingBank sets IssuingBank field to given value.

### HasIssuingBank

`func (o *CardBin) HasIssuingBank() bool`

HasIssuingBank returns a boolean if a field has been set.

### GetIssuingCountry

`func (o *CardBin) GetIssuingCountry() string`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *CardBin) GetIssuingCountryOk() (*string, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *CardBin) SetIssuingCountry(v string)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *CardBin) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### GetIssuingCurrency

`func (o *CardBin) GetIssuingCurrency() string`

GetIssuingCurrency returns the IssuingCurrency field if non-nil, zero value otherwise.

### GetIssuingCurrencyOk

`func (o *CardBin) GetIssuingCurrencyOk() (*string, bool)`

GetIssuingCurrencyOk returns a tuple with the IssuingCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCurrency

`func (o *CardBin) SetIssuingCurrency(v string)`

SetIssuingCurrency sets IssuingCurrency field to given value.

### HasIssuingCurrency

`func (o *CardBin) HasIssuingCurrency() bool`

HasIssuingCurrency returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *CardBin) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CardBin) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CardBin) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *CardBin) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPayoutEligible

`func (o *CardBin) GetPayoutEligible() string`

GetPayoutEligible returns the PayoutEligible field if non-nil, zero value otherwise.

### GetPayoutEligibleOk

`func (o *CardBin) GetPayoutEligibleOk() (*string, bool)`

GetPayoutEligibleOk returns a tuple with the PayoutEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutEligible

`func (o *CardBin) SetPayoutEligible(v string)`

SetPayoutEligible sets PayoutEligible field to given value.

### HasPayoutEligible

`func (o *CardBin) HasPayoutEligible() bool`

HasPayoutEligible returns a boolean if a field has been set.

### GetSummary

`func (o *CardBin) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CardBin) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CardBin) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CardBin) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


