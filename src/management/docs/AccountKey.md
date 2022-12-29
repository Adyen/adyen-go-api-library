# AccountKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCode** | **string** |  | 
**AccountTypeCode** | **string** |  | 
**NameSpace** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountKey

`func NewAccountKey(accountCode string, accountTypeCode string, ) *AccountKey`

NewAccountKey instantiates a new AccountKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountKeyWithDefaults

`func NewAccountKeyWithDefaults() *AccountKey`

NewAccountKeyWithDefaults instantiates a new AccountKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCode

`func (o *AccountKey) GetAccountCode() string`

GetAccountCode returns the AccountCode field if non-nil, zero value otherwise.

### GetAccountCodeOk

`func (o *AccountKey) GetAccountCodeOk() (*string, bool)`

GetAccountCodeOk returns a tuple with the AccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCode

`func (o *AccountKey) SetAccountCode(v string)`

SetAccountCode sets AccountCode field to given value.


### GetAccountTypeCode

`func (o *AccountKey) GetAccountTypeCode() string`

GetAccountTypeCode returns the AccountTypeCode field if non-nil, zero value otherwise.

### GetAccountTypeCodeOk

`func (o *AccountKey) GetAccountTypeCodeOk() (*string, bool)`

GetAccountTypeCodeOk returns a tuple with the AccountTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypeCode

`func (o *AccountKey) SetAccountTypeCode(v string)`

SetAccountTypeCode sets AccountTypeCode field to given value.


### GetNameSpace

`func (o *AccountKey) GetNameSpace() string`

GetNameSpace returns the NameSpace field if non-nil, zero value otherwise.

### GetNameSpaceOk

`func (o *AccountKey) GetNameSpaceOk() (*string, bool)`

GetNameSpaceOk returns a tuple with the NameSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSpace

`func (o *AccountKey) SetNameSpace(v string)`

SetNameSpace sets NameSpace field to given value.

### HasNameSpace

`func (o *AccountKey) HasNameSpace() bool`

HasNameSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


