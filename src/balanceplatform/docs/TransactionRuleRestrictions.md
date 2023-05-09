# TransactionRuleRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveNetworkTokens** | Pointer to [**ActiveNetworkTokensRestriction**](ActiveNetworkTokensRestriction.md) |  | [optional] 
**BrandVariants** | Pointer to [**BrandVariantsRestriction**](BrandVariantsRestriction.md) |  | [optional] 
**Countries** | Pointer to [**CountriesRestriction**](CountriesRestriction.md) |  | [optional] 
**DayOfWeek** | Pointer to [**DayOfWeekRestriction**](DayOfWeekRestriction.md) |  | [optional] 
**DifferentCurrencies** | Pointer to [**DifferentCurrenciesRestriction**](DifferentCurrenciesRestriction.md) |  | [optional] 
**EntryModes** | Pointer to [**EntryModesRestriction**](EntryModesRestriction.md) |  | [optional] 
**InternationalTransaction** | Pointer to [**InternationalTransactionRestriction**](InternationalTransactionRestriction.md) |  | [optional] 
**MatchingTransactions** | Pointer to [**MatchingTransactionsRestriction**](MatchingTransactionsRestriction.md) |  | [optional] 
**Mccs** | Pointer to [**MccsRestriction**](MccsRestriction.md) |  | [optional] 
**MerchantNames** | Pointer to [**MerchantNamesRestriction**](MerchantNamesRestriction.md) |  | [optional] 
**Merchants** | Pointer to [**MerchantsRestriction**](MerchantsRestriction.md) |  | [optional] 
**ProcessingTypes** | Pointer to [**ProcessingTypesRestriction**](ProcessingTypesRestriction.md) |  | [optional] 
**TimeOfDay** | Pointer to [**TimeOfDayRestriction**](TimeOfDayRestriction.md) |  | [optional] 
**TotalAmount** | Pointer to [**TotalAmountRestriction**](TotalAmountRestriction.md) |  | [optional] 

## Methods

### NewTransactionRuleRestrictions

`func NewTransactionRuleRestrictions() *TransactionRuleRestrictions`

NewTransactionRuleRestrictions instantiates a new TransactionRuleRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRuleRestrictionsWithDefaults

`func NewTransactionRuleRestrictionsWithDefaults() *TransactionRuleRestrictions`

NewTransactionRuleRestrictionsWithDefaults instantiates a new TransactionRuleRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveNetworkTokens

`func (o *TransactionRuleRestrictions) GetActiveNetworkTokens() ActiveNetworkTokensRestriction`

GetActiveNetworkTokens returns the ActiveNetworkTokens field if non-nil, zero value otherwise.

### GetActiveNetworkTokensOk

`func (o *TransactionRuleRestrictions) GetActiveNetworkTokensOk() (*ActiveNetworkTokensRestriction, bool)`

GetActiveNetworkTokensOk returns a tuple with the ActiveNetworkTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveNetworkTokens

`func (o *TransactionRuleRestrictions) SetActiveNetworkTokens(v ActiveNetworkTokensRestriction)`

SetActiveNetworkTokens sets ActiveNetworkTokens field to given value.

### HasActiveNetworkTokens

`func (o *TransactionRuleRestrictions) HasActiveNetworkTokens() bool`

HasActiveNetworkTokens returns a boolean if a field has been set.

### GetBrandVariants

`func (o *TransactionRuleRestrictions) GetBrandVariants() BrandVariantsRestriction`

GetBrandVariants returns the BrandVariants field if non-nil, zero value otherwise.

### GetBrandVariantsOk

`func (o *TransactionRuleRestrictions) GetBrandVariantsOk() (*BrandVariantsRestriction, bool)`

GetBrandVariantsOk returns a tuple with the BrandVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandVariants

`func (o *TransactionRuleRestrictions) SetBrandVariants(v BrandVariantsRestriction)`

SetBrandVariants sets BrandVariants field to given value.

### HasBrandVariants

`func (o *TransactionRuleRestrictions) HasBrandVariants() bool`

HasBrandVariants returns a boolean if a field has been set.

### GetCountries

`func (o *TransactionRuleRestrictions) GetCountries() CountriesRestriction`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *TransactionRuleRestrictions) GetCountriesOk() (*CountriesRestriction, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *TransactionRuleRestrictions) SetCountries(v CountriesRestriction)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *TransactionRuleRestrictions) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *TransactionRuleRestrictions) GetDayOfWeek() DayOfWeekRestriction`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *TransactionRuleRestrictions) GetDayOfWeekOk() (*DayOfWeekRestriction, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *TransactionRuleRestrictions) SetDayOfWeek(v DayOfWeekRestriction)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *TransactionRuleRestrictions) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetDifferentCurrencies

`func (o *TransactionRuleRestrictions) GetDifferentCurrencies() DifferentCurrenciesRestriction`

GetDifferentCurrencies returns the DifferentCurrencies field if non-nil, zero value otherwise.

### GetDifferentCurrenciesOk

`func (o *TransactionRuleRestrictions) GetDifferentCurrenciesOk() (*DifferentCurrenciesRestriction, bool)`

GetDifferentCurrenciesOk returns a tuple with the DifferentCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifferentCurrencies

`func (o *TransactionRuleRestrictions) SetDifferentCurrencies(v DifferentCurrenciesRestriction)`

SetDifferentCurrencies sets DifferentCurrencies field to given value.

### HasDifferentCurrencies

`func (o *TransactionRuleRestrictions) HasDifferentCurrencies() bool`

HasDifferentCurrencies returns a boolean if a field has been set.

### GetEntryModes

`func (o *TransactionRuleRestrictions) GetEntryModes() EntryModesRestriction`

GetEntryModes returns the EntryModes field if non-nil, zero value otherwise.

### GetEntryModesOk

`func (o *TransactionRuleRestrictions) GetEntryModesOk() (*EntryModesRestriction, bool)`

GetEntryModesOk returns a tuple with the EntryModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryModes

`func (o *TransactionRuleRestrictions) SetEntryModes(v EntryModesRestriction)`

SetEntryModes sets EntryModes field to given value.

### HasEntryModes

`func (o *TransactionRuleRestrictions) HasEntryModes() bool`

HasEntryModes returns a boolean if a field has been set.

### GetInternationalTransaction

`func (o *TransactionRuleRestrictions) GetInternationalTransaction() InternationalTransactionRestriction`

GetInternationalTransaction returns the InternationalTransaction field if non-nil, zero value otherwise.

### GetInternationalTransactionOk

`func (o *TransactionRuleRestrictions) GetInternationalTransactionOk() (*InternationalTransactionRestriction, bool)`

GetInternationalTransactionOk returns a tuple with the InternationalTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalTransaction

`func (o *TransactionRuleRestrictions) SetInternationalTransaction(v InternationalTransactionRestriction)`

SetInternationalTransaction sets InternationalTransaction field to given value.

### HasInternationalTransaction

`func (o *TransactionRuleRestrictions) HasInternationalTransaction() bool`

HasInternationalTransaction returns a boolean if a field has been set.

### GetMatchingTransactions

`func (o *TransactionRuleRestrictions) GetMatchingTransactions() MatchingTransactionsRestriction`

GetMatchingTransactions returns the MatchingTransactions field if non-nil, zero value otherwise.

### GetMatchingTransactionsOk

`func (o *TransactionRuleRestrictions) GetMatchingTransactionsOk() (*MatchingTransactionsRestriction, bool)`

GetMatchingTransactionsOk returns a tuple with the MatchingTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingTransactions

`func (o *TransactionRuleRestrictions) SetMatchingTransactions(v MatchingTransactionsRestriction)`

SetMatchingTransactions sets MatchingTransactions field to given value.

### HasMatchingTransactions

`func (o *TransactionRuleRestrictions) HasMatchingTransactions() bool`

HasMatchingTransactions returns a boolean if a field has been set.

### GetMccs

`func (o *TransactionRuleRestrictions) GetMccs() MccsRestriction`

GetMccs returns the Mccs field if non-nil, zero value otherwise.

### GetMccsOk

`func (o *TransactionRuleRestrictions) GetMccsOk() (*MccsRestriction, bool)`

GetMccsOk returns a tuple with the Mccs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMccs

`func (o *TransactionRuleRestrictions) SetMccs(v MccsRestriction)`

SetMccs sets Mccs field to given value.

### HasMccs

`func (o *TransactionRuleRestrictions) HasMccs() bool`

HasMccs returns a boolean if a field has been set.

### GetMerchantNames

`func (o *TransactionRuleRestrictions) GetMerchantNames() MerchantNamesRestriction`

GetMerchantNames returns the MerchantNames field if non-nil, zero value otherwise.

### GetMerchantNamesOk

`func (o *TransactionRuleRestrictions) GetMerchantNamesOk() (*MerchantNamesRestriction, bool)`

GetMerchantNamesOk returns a tuple with the MerchantNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantNames

`func (o *TransactionRuleRestrictions) SetMerchantNames(v MerchantNamesRestriction)`

SetMerchantNames sets MerchantNames field to given value.

### HasMerchantNames

`func (o *TransactionRuleRestrictions) HasMerchantNames() bool`

HasMerchantNames returns a boolean if a field has been set.

### GetMerchants

`func (o *TransactionRuleRestrictions) GetMerchants() MerchantsRestriction`

GetMerchants returns the Merchants field if non-nil, zero value otherwise.

### GetMerchantsOk

`func (o *TransactionRuleRestrictions) GetMerchantsOk() (*MerchantsRestriction, bool)`

GetMerchantsOk returns a tuple with the Merchants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchants

`func (o *TransactionRuleRestrictions) SetMerchants(v MerchantsRestriction)`

SetMerchants sets Merchants field to given value.

### HasMerchants

`func (o *TransactionRuleRestrictions) HasMerchants() bool`

HasMerchants returns a boolean if a field has been set.

### GetProcessingTypes

`func (o *TransactionRuleRestrictions) GetProcessingTypes() ProcessingTypesRestriction`

GetProcessingTypes returns the ProcessingTypes field if non-nil, zero value otherwise.

### GetProcessingTypesOk

`func (o *TransactionRuleRestrictions) GetProcessingTypesOk() (*ProcessingTypesRestriction, bool)`

GetProcessingTypesOk returns a tuple with the ProcessingTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTypes

`func (o *TransactionRuleRestrictions) SetProcessingTypes(v ProcessingTypesRestriction)`

SetProcessingTypes sets ProcessingTypes field to given value.

### HasProcessingTypes

`func (o *TransactionRuleRestrictions) HasProcessingTypes() bool`

HasProcessingTypes returns a boolean if a field has been set.

### GetTimeOfDay

`func (o *TransactionRuleRestrictions) GetTimeOfDay() TimeOfDayRestriction`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *TransactionRuleRestrictions) GetTimeOfDayOk() (*TimeOfDayRestriction, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *TransactionRuleRestrictions) SetTimeOfDay(v TimeOfDayRestriction)`

SetTimeOfDay sets TimeOfDay field to given value.

### HasTimeOfDay

`func (o *TransactionRuleRestrictions) HasTimeOfDay() bool`

HasTimeOfDay returns a boolean if a field has been set.

### GetTotalAmount

`func (o *TransactionRuleRestrictions) GetTotalAmount() TotalAmountRestriction`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *TransactionRuleRestrictions) GetTotalAmountOk() (*TotalAmountRestriction, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *TransactionRuleRestrictions) SetTotalAmount(v TotalAmountRestriction)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *TransactionRuleRestrictions) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


