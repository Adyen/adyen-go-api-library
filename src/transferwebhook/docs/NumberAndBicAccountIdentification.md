# NumberAndBicAccountIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | The bank account number, without separators or whitespace. The length and format depends on the bank or country. | 
**AdditionalBankIdentification** | Pointer to [**AdditionalBankIdentification**](AdditionalBankIdentification.md) |  | [optional] 
**Bic** | **string** | The bank&#39;s 8- or 11-character BIC or SWIFT code. | 
**Type** | **string** | **numberAndBic** | [default to "numberAndBic"]

## Methods

### NewNumberAndBicAccountIdentification

`func NewNumberAndBicAccountIdentification(accountNumber string, bic string, type_ string, ) *NumberAndBicAccountIdentification`

NewNumberAndBicAccountIdentification instantiates a new NumberAndBicAccountIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberAndBicAccountIdentificationWithDefaults

`func NewNumberAndBicAccountIdentificationWithDefaults() *NumberAndBicAccountIdentification`

NewNumberAndBicAccountIdentificationWithDefaults instantiates a new NumberAndBicAccountIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *NumberAndBicAccountIdentification) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *NumberAndBicAccountIdentification) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *NumberAndBicAccountIdentification) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAdditionalBankIdentification

`func (o *NumberAndBicAccountIdentification) GetAdditionalBankIdentification() AdditionalBankIdentification`

GetAdditionalBankIdentification returns the AdditionalBankIdentification field if non-nil, zero value otherwise.

### GetAdditionalBankIdentificationOk

`func (o *NumberAndBicAccountIdentification) GetAdditionalBankIdentificationOk() (*AdditionalBankIdentification, bool)`

GetAdditionalBankIdentificationOk returns a tuple with the AdditionalBankIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalBankIdentification

`func (o *NumberAndBicAccountIdentification) SetAdditionalBankIdentification(v AdditionalBankIdentification)`

SetAdditionalBankIdentification sets AdditionalBankIdentification field to given value.

### HasAdditionalBankIdentification

`func (o *NumberAndBicAccountIdentification) HasAdditionalBankIdentification() bool`

HasAdditionalBankIdentification returns a boolean if a field has been set.

### GetBic

`func (o *NumberAndBicAccountIdentification) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *NumberAndBicAccountIdentification) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *NumberAndBicAccountIdentification) SetBic(v string)`

SetBic sets Bic field to given value.


### GetType

`func (o *NumberAndBicAccountIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NumberAndBicAccountIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NumberAndBicAccountIdentification) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


