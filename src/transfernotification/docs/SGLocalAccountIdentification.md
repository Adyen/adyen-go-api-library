# SGLocalAccountIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | The 4- to 19-digit bank account number, without separators or whitespace. | 
**Bic** | **string** | The bank&#39;s 8- or 11-character BIC or SWIFT code. | 
**Type** | Pointer to **string** | **sgLocal** | [optional] [default to "sgLocal"]

## Methods

### NewSGLocalAccountIdentification

`func NewSGLocalAccountIdentification(accountNumber string, bic string, ) *SGLocalAccountIdentification`

NewSGLocalAccountIdentification instantiates a new SGLocalAccountIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSGLocalAccountIdentificationWithDefaults

`func NewSGLocalAccountIdentificationWithDefaults() *SGLocalAccountIdentification`

NewSGLocalAccountIdentificationWithDefaults instantiates a new SGLocalAccountIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *SGLocalAccountIdentification) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *SGLocalAccountIdentification) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *SGLocalAccountIdentification) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetBic

`func (o *SGLocalAccountIdentification) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *SGLocalAccountIdentification) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *SGLocalAccountIdentification) SetBic(v string)`

SetBic sets Bic field to given value.


### GetType

`func (o *SGLocalAccountIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SGLocalAccountIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SGLocalAccountIdentification) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SGLocalAccountIdentification) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


