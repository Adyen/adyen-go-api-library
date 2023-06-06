# SELocalAccountIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | The 7- to 10-digit bank account number ([Bankkontonummer](https://sv.wikipedia.org/wiki/Bankkonto)), without the clearing number, separators, or whitespace. | 
**ClearingNumber** | **string** | The 4- to 5-digit clearing number ([Clearingnummer](https://sv.wikipedia.org/wiki/Clearingnummer)), without separators or whitespace. | 
**Type** | **string** | **seLocal** | [default to "seLocal"]

## Methods

### NewSELocalAccountIdentification

`func NewSELocalAccountIdentification(accountNumber string, clearingNumber string, type_ string, ) *SELocalAccountIdentification`

NewSELocalAccountIdentification instantiates a new SELocalAccountIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSELocalAccountIdentificationWithDefaults

`func NewSELocalAccountIdentificationWithDefaults() *SELocalAccountIdentification`

NewSELocalAccountIdentificationWithDefaults instantiates a new SELocalAccountIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *SELocalAccountIdentification) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *SELocalAccountIdentification) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *SELocalAccountIdentification) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetClearingNumber

`func (o *SELocalAccountIdentification) GetClearingNumber() string`

GetClearingNumber returns the ClearingNumber field if non-nil, zero value otherwise.

### GetClearingNumberOk

`func (o *SELocalAccountIdentification) GetClearingNumberOk() (*string, bool)`

GetClearingNumberOk returns a tuple with the ClearingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearingNumber

`func (o *SELocalAccountIdentification) SetClearingNumber(v string)`

SetClearingNumber sets ClearingNumber field to given value.


### GetType

`func (o *SELocalAccountIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SELocalAccountIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SELocalAccountIdentification) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


