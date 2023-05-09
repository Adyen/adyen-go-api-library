# BankAccountIdentificationValidationRequestAccountIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | The bank account number, without separators or whitespace. | 
**BsbCode** | **string** | The 6-digit [Bank State Branch (BSB) code](https://en.wikipedia.org/wiki/Bank_state_branch), without separators or whitespace. | 
**Type** | **string** | **usLocal** | [default to "usLocal"]
**AccountType** | Pointer to **string** | The bank account type.  Possible values: **checking** or **savings**. Defaults to **checking**. | [optional] [default to "checking"]
**InstitutionNumber** | **string** | The 3-digit institution number, without separators or whitespace. | 
**TransitNumber** | **string** | The 5-digit transit number, without separators or whitespace. | 
**BankCode** | **string** | The 4-digit bank code (Kód banky), without separators or whitespace. | 
**Iban** | **string** | The international bank account number as defined in the [ISO-13616](https://www.iso.org/standard/81090.html) standard. | 
**AdditionalBankIdentification** | Pointer to [**AdditionalBankIdentification**](AdditionalBankIdentification.md) |  | [optional] 
**Bic** | **string** | The bank&#39;s 8- or 11-character BIC or SWIFT code. | 
**ClearingNumber** | **string** | The 4- to 5-digit clearing number ([Clearingnummer](https://sv.wikipedia.org/wiki/Clearingnummer)), without separators or whitespace. | 
**SortCode** | **string** | The 6-digit [sort code](https://en.wikipedia.org/wiki/Sort_code), without separators or whitespace. | 
**RoutingNumber** | **string** | The 9-digit [routing number](https://en.wikipedia.org/wiki/ABA_routing_transit_number), without separators or whitespace. | 

## Methods

### NewBankAccountIdentificationValidationRequestAccountIdentification

`func NewBankAccountIdentificationValidationRequestAccountIdentification(accountNumber string, bsbCode string, type_ string, institutionNumber string, transitNumber string, bankCode string, iban string, bic string, clearingNumber string, sortCode string, routingNumber string, ) *BankAccountIdentificationValidationRequestAccountIdentification`

NewBankAccountIdentificationValidationRequestAccountIdentification instantiates a new BankAccountIdentificationValidationRequestAccountIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountIdentificationValidationRequestAccountIdentificationWithDefaults

`func NewBankAccountIdentificationValidationRequestAccountIdentificationWithDefaults() *BankAccountIdentificationValidationRequestAccountIdentification`

NewBankAccountIdentificationValidationRequestAccountIdentificationWithDefaults instantiates a new BankAccountIdentificationValidationRequestAccountIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetBsbCode

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetBsbCode() string`

GetBsbCode returns the BsbCode field if non-nil, zero value otherwise.

### GetBsbCodeOk

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetBsbCodeOk() (*string, bool)`

GetBsbCodeOk returns a tuple with the BsbCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsbCode

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) SetBsbCode(v string)`

SetBsbCode sets BsbCode field to given value.


### GetType

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) SetType(v string)`

SetType sets Type field to given value.


### GetAccountType

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetInstitutionNumber

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetInstitutionNumber() string`

GetInstitutionNumber returns the InstitutionNumber field if non-nil, zero value otherwise.

### GetInstitutionNumberOk

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetInstitutionNumberOk() (*string, bool)`

GetInstitutionNumberOk returns a tuple with the InstitutionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionNumber

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) SetInstitutionNumber(v string)`

SetInstitutionNumber sets InstitutionNumber field to given value.


### GetTransitNumber

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetTransitNumber() string`

GetTransitNumber returns the TransitNumber field if non-nil, zero value otherwise.

### GetTransitNumberOk

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetTransitNumberOk() (*string, bool)`

GetTransitNumberOk returns a tuple with the TransitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitNumber

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) SetTransitNumber(v string)`

SetTransitNumber sets TransitNumber field to given value.


### GetBankCode

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.


### GetIban

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) SetIban(v string)`

SetIban sets Iban field to given value.


### GetAdditionalBankIdentification

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetAdditionalBankIdentification() AdditionalBankIdentification`

GetAdditionalBankIdentification returns the AdditionalBankIdentification field if non-nil, zero value otherwise.

### GetAdditionalBankIdentificationOk

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetAdditionalBankIdentificationOk() (*AdditionalBankIdentification, bool)`

GetAdditionalBankIdentificationOk returns a tuple with the AdditionalBankIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalBankIdentification

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) SetAdditionalBankIdentification(v AdditionalBankIdentification)`

SetAdditionalBankIdentification sets AdditionalBankIdentification field to given value.

### HasAdditionalBankIdentification

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) HasAdditionalBankIdentification() bool`

HasAdditionalBankIdentification returns a boolean if a field has been set.

### GetBic

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) SetBic(v string)`

SetBic sets Bic field to given value.


### GetClearingNumber

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetClearingNumber() string`

GetClearingNumber returns the ClearingNumber field if non-nil, zero value otherwise.

### GetClearingNumberOk

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetClearingNumberOk() (*string, bool)`

GetClearingNumberOk returns a tuple with the ClearingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearingNumber

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) SetClearingNumber(v string)`

SetClearingNumber sets ClearingNumber field to given value.


### GetSortCode

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetSortCode() string`

GetSortCode returns the SortCode field if non-nil, zero value otherwise.

### GetSortCodeOk

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetSortCodeOk() (*string, bool)`

GetSortCodeOk returns a tuple with the SortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCode

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) SetSortCode(v string)`

SetSortCode sets SortCode field to given value.


### GetRoutingNumber

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *BankAccountIdentificationValidationRequestAccountIdentification) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


