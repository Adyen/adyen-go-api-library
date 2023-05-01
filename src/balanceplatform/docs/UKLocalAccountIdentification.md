# UKLocalAccountIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | The 8-digit bank account number, without separators or whitespace. | 
**SortCode** | **string** | The 6-digit [sort code](https://en.wikipedia.org/wiki/Sort_code), without separators or whitespace. | 
**Type** | **string** | **ukLocal** | [default to "ukLocal"]

## Methods

### NewUKLocalAccountIdentification

`func NewUKLocalAccountIdentification(accountNumber string, sortCode string, type_ string, ) *UKLocalAccountIdentification`

NewUKLocalAccountIdentification instantiates a new UKLocalAccountIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUKLocalAccountIdentificationWithDefaults

`func NewUKLocalAccountIdentificationWithDefaults() *UKLocalAccountIdentification`

NewUKLocalAccountIdentificationWithDefaults instantiates a new UKLocalAccountIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *UKLocalAccountIdentification) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *UKLocalAccountIdentification) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *UKLocalAccountIdentification) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetSortCode

`func (o *UKLocalAccountIdentification) GetSortCode() string`

GetSortCode returns the SortCode field if non-nil, zero value otherwise.

### GetSortCodeOk

`func (o *UKLocalAccountIdentification) GetSortCodeOk() (*string, bool)`

GetSortCodeOk returns a tuple with the SortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCode

`func (o *UKLocalAccountIdentification) SetSortCode(v string)`

SetSortCode sets SortCode field to given value.


### GetType

`func (o *UKLocalAccountIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UKLocalAccountIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UKLocalAccountIdentification) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


