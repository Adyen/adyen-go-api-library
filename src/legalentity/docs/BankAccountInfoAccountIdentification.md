# BankAccountInfoAccountIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | The bank account number, without separators or whitespace. | 
**BsbCode** | **string** | The 6-digit [Bank State Branch (BSB) code](https://en.wikipedia.org/wiki/Bank_state_branch), without separators or whitespace. | 
**Type** | **string** | **usLocal** | [default to "usLocal"]
**AccountType** | Pointer to **string** | The bank account type.  Possible values: **checking** or **savings**. Defaults to **checking**. | [optional] [default to "checking"]
**InstitutionNumber** | **string** | The 3-digit institution number, without separators or whitespace. | 
**TransitNumber** | **string** | The 5-digit transit number, without separators or whitespace. | 
**BankCode** | **string** | The 4-digit bank code (Registreringsnummer) (without separators or whitespace). | 
**Iban** | **string** | The international bank account number as defined in the [ISO-13616](https://www.iso.org/standard/81090.html) standard. | 
**AdditionalBankIdentification** | Pointer to [**AdditionalBankIdentification**](AdditionalBankIdentification.md) |  | [optional] 
**Bic** | **string** | The bank&#39;s 8- or 11-character BIC or SWIFT code. | 
**ClearingNumber** | **string** | The 4- to 5-digit clearing number ([Clearingnummer](https://sv.wikipedia.org/wiki/Clearingnummer)), without separators or whitespace. | 
**SortCode** | **string** | The 6-digit [sort code](https://en.wikipedia.org/wiki/Sort_code), without separators or whitespace. | 
**RoutingNumber** | **string** | The 9-digit [routing number](https://en.wikipedia.org/wiki/ABA_routing_transit_number), without separators or whitespace. | 

## Methods

### NewBankAccountInfoAccountIdentification

`func NewBankAccountInfoAccountIdentification(accountNumber string, bsbCode string, type_ string, institutionNumber string, transitNumber string, bankCode string, iban string, bic string, clearingNumber string, sortCode string, routingNumber string, ) *BankAccountInfoAccountIdentification`

NewBankAccountInfoAccountIdentification instantiates a new BankAccountInfoAccountIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountInfoAccountIdentificationWithDefaults

`func NewBankAccountInfoAccountIdentificationWithDefaults() *BankAccountInfoAccountIdentification`

NewBankAccountInfoAccountIdentificationWithDefaults instantiates a new BankAccountInfoAccountIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *BankAccountInfoAccountIdentification) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *BankAccountInfoAccountIdentification) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *BankAccountInfoAccountIdentification) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetBsbCode

`func (o *BankAccountInfoAccountIdentification) GetBsbCode() string`

GetBsbCode returns the BsbCode field if non-nil, zero value otherwise.

### GetBsbCodeOk

`func (o *BankAccountInfoAccountIdentification) GetBsbCodeOk() (*string, bool)`

GetBsbCodeOk returns a tuple with the BsbCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsbCode

`func (o *BankAccountInfoAccountIdentification) SetBsbCode(v string)`

SetBsbCode sets BsbCode field to given value.


### GetType

`func (o *BankAccountInfoAccountIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BankAccountInfoAccountIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BankAccountInfoAccountIdentification) SetType(v string)`

SetType sets Type field to given value.


### GetAccountType

`func (o *BankAccountInfoAccountIdentification) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BankAccountInfoAccountIdentification) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BankAccountInfoAccountIdentification) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *BankAccountInfoAccountIdentification) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetInstitutionNumber

`func (o *BankAccountInfoAccountIdentification) GetInstitutionNumber() string`

GetInstitutionNumber returns the InstitutionNumber field if non-nil, zero value otherwise.

### GetInstitutionNumberOk

`func (o *BankAccountInfoAccountIdentification) GetInstitutionNumberOk() (*string, bool)`

GetInstitutionNumberOk returns a tuple with the InstitutionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionNumber

`func (o *BankAccountInfoAccountIdentification) SetInstitutionNumber(v string)`

SetInstitutionNumber sets InstitutionNumber field to given value.


### GetTransitNumber

`func (o *BankAccountInfoAccountIdentification) GetTransitNumber() string`

GetTransitNumber returns the TransitNumber field if non-nil, zero value otherwise.

### GetTransitNumberOk

`func (o *BankAccountInfoAccountIdentification) GetTransitNumberOk() (*string, bool)`

GetTransitNumberOk returns a tuple with the TransitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitNumber

`func (o *BankAccountInfoAccountIdentification) SetTransitNumber(v string)`

SetTransitNumber sets TransitNumber field to given value.


### GetBankCode

`func (o *BankAccountInfoAccountIdentification) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *BankAccountInfoAccountIdentification) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *BankAccountInfoAccountIdentification) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.


### GetIban

`func (o *BankAccountInfoAccountIdentification) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *BankAccountInfoAccountIdentification) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *BankAccountInfoAccountIdentification) SetIban(v string)`

SetIban sets Iban field to given value.


### GetAdditionalBankIdentification

`func (o *BankAccountInfoAccountIdentification) GetAdditionalBankIdentification() AdditionalBankIdentification`

GetAdditionalBankIdentification returns the AdditionalBankIdentification field if non-nil, zero value otherwise.

### GetAdditionalBankIdentificationOk

`func (o *BankAccountInfoAccountIdentification) GetAdditionalBankIdentificationOk() (*AdditionalBankIdentification, bool)`

GetAdditionalBankIdentificationOk returns a tuple with the AdditionalBankIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalBankIdentification

`func (o *BankAccountInfoAccountIdentification) SetAdditionalBankIdentification(v AdditionalBankIdentification)`

SetAdditionalBankIdentification sets AdditionalBankIdentification field to given value.

### HasAdditionalBankIdentification

`func (o *BankAccountInfoAccountIdentification) HasAdditionalBankIdentification() bool`

HasAdditionalBankIdentification returns a boolean if a field has been set.

### GetBic

`func (o *BankAccountInfoAccountIdentification) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *BankAccountInfoAccountIdentification) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *BankAccountInfoAccountIdentification) SetBic(v string)`

SetBic sets Bic field to given value.


### GetClearingNumber

`func (o *BankAccountInfoAccountIdentification) GetClearingNumber() string`

GetClearingNumber returns the ClearingNumber field if non-nil, zero value otherwise.

### GetClearingNumberOk

`func (o *BankAccountInfoAccountIdentification) GetClearingNumberOk() (*string, bool)`

GetClearingNumberOk returns a tuple with the ClearingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearingNumber

`func (o *BankAccountInfoAccountIdentification) SetClearingNumber(v string)`

SetClearingNumber sets ClearingNumber field to given value.


### GetSortCode

`func (o *BankAccountInfoAccountIdentification) GetSortCode() string`

GetSortCode returns the SortCode field if non-nil, zero value otherwise.

### GetSortCodeOk

`func (o *BankAccountInfoAccountIdentification) GetSortCodeOk() (*string, bool)`

GetSortCodeOk returns a tuple with the SortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCode

`func (o *BankAccountInfoAccountIdentification) SetSortCode(v string)`

SetSortCode sets SortCode field to given value.


### GetRoutingNumber

`func (o *BankAccountInfoAccountIdentification) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *BankAccountInfoAccountIdentification) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *BankAccountInfoAccountIdentification) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


