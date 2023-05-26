# AULocalAccountIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | The bank account number, without separators or whitespace. | 
**BsbCode** | **string** | The 6-digit [Bank State Branch (BSB) code](https://en.wikipedia.org/wiki/Bank_state_branch), without separators or whitespace. | 
**Type** | **string** | **auLocal** | [default to "auLocal"]

## Methods

### NewAULocalAccountIdentification

`func NewAULocalAccountIdentification(accountNumber string, bsbCode string, type_ string, ) *AULocalAccountIdentification`

NewAULocalAccountIdentification instantiates a new AULocalAccountIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAULocalAccountIdentificationWithDefaults

`func NewAULocalAccountIdentificationWithDefaults() *AULocalAccountIdentification`

NewAULocalAccountIdentificationWithDefaults instantiates a new AULocalAccountIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *AULocalAccountIdentification) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AULocalAccountIdentification) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AULocalAccountIdentification) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetBsbCode

`func (o *AULocalAccountIdentification) GetBsbCode() string`

GetBsbCode returns the BsbCode field if non-nil, zero value otherwise.

### GetBsbCodeOk

`func (o *AULocalAccountIdentification) GetBsbCodeOk() (*string, bool)`

GetBsbCodeOk returns a tuple with the BsbCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsbCode

`func (o *AULocalAccountIdentification) SetBsbCode(v string)`

SetBsbCode sets BsbCode field to given value.


### GetType

`func (o *AULocalAccountIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AULocalAccountIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AULocalAccountIdentification) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


