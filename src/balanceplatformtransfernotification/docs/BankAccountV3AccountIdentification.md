# BankAccountV3AccountIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | The bank account number, without separators or whitespace. | 
**BsbCode** | **string** | The 6-digit [Bank State Branch (BSB) code](https://en.wikipedia.org/wiki/Bank_state_branch), without separators or whitespace. | 
**Type** | **string** | **usLocal** | [default to "usLocal"]
**BankCode** | **string** | The 4-digit bank code (Registreringsnummer) (without separators or whitespace). | 
**BranchNumber** | **string** | The bank account branch number (without separators or whitespace). | 
**AccountType** | Pointer to **string** | The bank account type.  Possible values: **checking** or **savings**. Defaults to **checking**. | [optional] [default to "checking"]
**InstitutionNumber** | **string** | The 3-digit institution number, without separators or whitespace. | 
**TransitNumber** | **string** | The 5-digit transit number, without separators or whitespace. | 
**Iban** | **string** | The international bank account number as defined in the [ISO-13616](https://www.iso.org/standard/81090.html) standard. | 
**AdditionalBankIdentification** | Pointer to [**AdditionalBankIdentification**](AdditionalBankIdentification.md) |  | [optional] 
**Bic** | **string** | The bank&#39;s 8- or 11-character BIC or SWIFT code. | 
**ClearingNumber** | **string** | The 4- to 5-digit clearing number ([Clearingnummer](https://sv.wikipedia.org/wiki/Clearingnummer)), without separators or whitespace. | 
**SortCode** | **string** | The 6-digit [sort code](https://en.wikipedia.org/wiki/Sort_code), without separators or whitespace. | 
**RoutingNumber** | **string** | The 9-digit [routing number](https://en.wikipedia.org/wiki/ABA_routing_transit_number), without separators or whitespace. | 

## Methods

### NewBankAccountV3AccountIdentification

`func NewBankAccountV3AccountIdentification(accountNumber string, bsbCode string, type_ string, bankCode string, branchNumber string, institutionNumber string, transitNumber string, iban string, bic string, clearingNumber string, sortCode string, routingNumber string, ) *BankAccountV3AccountIdentification`

NewBankAccountV3AccountIdentification instantiates a new BankAccountV3AccountIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountV3AccountIdentificationWithDefaults

`func NewBankAccountV3AccountIdentificationWithDefaults() *BankAccountV3AccountIdentification`

NewBankAccountV3AccountIdentificationWithDefaults instantiates a new BankAccountV3AccountIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *BankAccountV3AccountIdentification) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *BankAccountV3AccountIdentification) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *BankAccountV3AccountIdentification) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetBsbCode

`func (o *BankAccountV3AccountIdentification) GetBsbCode() string`

GetBsbCode returns the BsbCode field if non-nil, zero value otherwise.

### GetBsbCodeOk

`func (o *BankAccountV3AccountIdentification) GetBsbCodeOk() (*string, bool)`

GetBsbCodeOk returns a tuple with the BsbCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsbCode

`func (o *BankAccountV3AccountIdentification) SetBsbCode(v string)`

SetBsbCode sets BsbCode field to given value.


### GetType

`func (o *BankAccountV3AccountIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BankAccountV3AccountIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BankAccountV3AccountIdentification) SetType(v string)`

SetType sets Type field to given value.


### GetBankCode

`func (o *BankAccountV3AccountIdentification) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *BankAccountV3AccountIdentification) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *BankAccountV3AccountIdentification) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.


### GetBranchNumber

`func (o *BankAccountV3AccountIdentification) GetBranchNumber() string`

GetBranchNumber returns the BranchNumber field if non-nil, zero value otherwise.

### GetBranchNumberOk

`func (o *BankAccountV3AccountIdentification) GetBranchNumberOk() (*string, bool)`

GetBranchNumberOk returns a tuple with the BranchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchNumber

`func (o *BankAccountV3AccountIdentification) SetBranchNumber(v string)`

SetBranchNumber sets BranchNumber field to given value.


### GetAccountType

`func (o *BankAccountV3AccountIdentification) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BankAccountV3AccountIdentification) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BankAccountV3AccountIdentification) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *BankAccountV3AccountIdentification) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetInstitutionNumber

`func (o *BankAccountV3AccountIdentification) GetInstitutionNumber() string`

GetInstitutionNumber returns the InstitutionNumber field if non-nil, zero value otherwise.

### GetInstitutionNumberOk

`func (o *BankAccountV3AccountIdentification) GetInstitutionNumberOk() (*string, bool)`

GetInstitutionNumberOk returns a tuple with the InstitutionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionNumber

`func (o *BankAccountV3AccountIdentification) SetInstitutionNumber(v string)`

SetInstitutionNumber sets InstitutionNumber field to given value.


### GetTransitNumber

`func (o *BankAccountV3AccountIdentification) GetTransitNumber() string`

GetTransitNumber returns the TransitNumber field if non-nil, zero value otherwise.

### GetTransitNumberOk

`func (o *BankAccountV3AccountIdentification) GetTransitNumberOk() (*string, bool)`

GetTransitNumberOk returns a tuple with the TransitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitNumber

`func (o *BankAccountV3AccountIdentification) SetTransitNumber(v string)`

SetTransitNumber sets TransitNumber field to given value.


### GetIban

`func (o *BankAccountV3AccountIdentification) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *BankAccountV3AccountIdentification) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *BankAccountV3AccountIdentification) SetIban(v string)`

SetIban sets Iban field to given value.


### GetAdditionalBankIdentification

`func (o *BankAccountV3AccountIdentification) GetAdditionalBankIdentification() AdditionalBankIdentification`

GetAdditionalBankIdentification returns the AdditionalBankIdentification field if non-nil, zero value otherwise.

### GetAdditionalBankIdentificationOk

`func (o *BankAccountV3AccountIdentification) GetAdditionalBankIdentificationOk() (*AdditionalBankIdentification, bool)`

GetAdditionalBankIdentificationOk returns a tuple with the AdditionalBankIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalBankIdentification

`func (o *BankAccountV3AccountIdentification) SetAdditionalBankIdentification(v AdditionalBankIdentification)`

SetAdditionalBankIdentification sets AdditionalBankIdentification field to given value.

### HasAdditionalBankIdentification

`func (o *BankAccountV3AccountIdentification) HasAdditionalBankIdentification() bool`

HasAdditionalBankIdentification returns a boolean if a field has been set.

### GetBic

`func (o *BankAccountV3AccountIdentification) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *BankAccountV3AccountIdentification) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *BankAccountV3AccountIdentification) SetBic(v string)`

SetBic sets Bic field to given value.


### GetClearingNumber

`func (o *BankAccountV3AccountIdentification) GetClearingNumber() string`

GetClearingNumber returns the ClearingNumber field if non-nil, zero value otherwise.

### GetClearingNumberOk

`func (o *BankAccountV3AccountIdentification) GetClearingNumberOk() (*string, bool)`

GetClearingNumberOk returns a tuple with the ClearingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearingNumber

`func (o *BankAccountV3AccountIdentification) SetClearingNumber(v string)`

SetClearingNumber sets ClearingNumber field to given value.


### GetSortCode

`func (o *BankAccountV3AccountIdentification) GetSortCode() string`

GetSortCode returns the SortCode field if non-nil, zero value otherwise.

### GetSortCodeOk

`func (o *BankAccountV3AccountIdentification) GetSortCodeOk() (*string, bool)`

GetSortCodeOk returns a tuple with the SortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCode

`func (o *BankAccountV3AccountIdentification) SetSortCode(v string)`

SetSortCode sets SortCode field to given value.


### GetRoutingNumber

`func (o *BankAccountV3AccountIdentification) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *BankAccountV3AccountIdentification) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *BankAccountV3AccountIdentification) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


