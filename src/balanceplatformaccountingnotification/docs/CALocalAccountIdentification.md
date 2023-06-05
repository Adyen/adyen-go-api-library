# CALocalAccountIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | The 5- to 12-digit bank account number, without separators or whitespace. | 
**InstitutionNumber** | **string** | The 3-digit institution number, without separators or whitespace. | 
**TransitNumber** | **string** | The 5-digit transit number, without separators or whitespace. | 
**Type** | **string** | **caLocal** | [default to "caLocal"]

## Methods

### NewCALocalAccountIdentification

`func NewCALocalAccountIdentification(accountNumber string, institutionNumber string, transitNumber string, type_ string, ) *CALocalAccountIdentification`

NewCALocalAccountIdentification instantiates a new CALocalAccountIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCALocalAccountIdentificationWithDefaults

`func NewCALocalAccountIdentificationWithDefaults() *CALocalAccountIdentification`

NewCALocalAccountIdentificationWithDefaults instantiates a new CALocalAccountIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *CALocalAccountIdentification) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CALocalAccountIdentification) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CALocalAccountIdentification) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetInstitutionNumber

`func (o *CALocalAccountIdentification) GetInstitutionNumber() string`

GetInstitutionNumber returns the InstitutionNumber field if non-nil, zero value otherwise.

### GetInstitutionNumberOk

`func (o *CALocalAccountIdentification) GetInstitutionNumberOk() (*string, bool)`

GetInstitutionNumberOk returns a tuple with the InstitutionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionNumber

`func (o *CALocalAccountIdentification) SetInstitutionNumber(v string)`

SetInstitutionNumber sets InstitutionNumber field to given value.


### GetTransitNumber

`func (o *CALocalAccountIdentification) GetTransitNumber() string`

GetTransitNumber returns the TransitNumber field if non-nil, zero value otherwise.

### GetTransitNumberOk

`func (o *CALocalAccountIdentification) GetTransitNumberOk() (*string, bool)`

GetTransitNumberOk returns a tuple with the TransitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitNumber

`func (o *CALocalAccountIdentification) SetTransitNumber(v string)`

SetTransitNumber sets TransitNumber field to given value.


### GetType

`func (o *CALocalAccountIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CALocalAccountIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CALocalAccountIdentification) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


