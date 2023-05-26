# GetTerminalsUnderAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyAccount** | **string** | Your company account. | 
**InventoryTerminals** | Pointer to **[]string** | Array that returns a list of all terminals that are in the inventory of the company account. | [optional] 
**MerchantAccounts** | Pointer to [**[]MerchantAccount**](MerchantAccount.md) | Array that returns a list of all merchant accounts belonging to the company account. | [optional] 

## Methods

### NewGetTerminalsUnderAccountResponse

`func NewGetTerminalsUnderAccountResponse(companyAccount string, ) *GetTerminalsUnderAccountResponse`

NewGetTerminalsUnderAccountResponse instantiates a new GetTerminalsUnderAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTerminalsUnderAccountResponseWithDefaults

`func NewGetTerminalsUnderAccountResponseWithDefaults() *GetTerminalsUnderAccountResponse`

NewGetTerminalsUnderAccountResponseWithDefaults instantiates a new GetTerminalsUnderAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyAccount

`func (o *GetTerminalsUnderAccountResponse) GetCompanyAccount() string`

GetCompanyAccount returns the CompanyAccount field if non-nil, zero value otherwise.

### GetCompanyAccountOk

`func (o *GetTerminalsUnderAccountResponse) GetCompanyAccountOk() (*string, bool)`

GetCompanyAccountOk returns a tuple with the CompanyAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAccount

`func (o *GetTerminalsUnderAccountResponse) SetCompanyAccount(v string)`

SetCompanyAccount sets CompanyAccount field to given value.


### GetInventoryTerminals

`func (o *GetTerminalsUnderAccountResponse) GetInventoryTerminals() []string`

GetInventoryTerminals returns the InventoryTerminals field if non-nil, zero value otherwise.

### GetInventoryTerminalsOk

`func (o *GetTerminalsUnderAccountResponse) GetInventoryTerminalsOk() (*[]string, bool)`

GetInventoryTerminalsOk returns a tuple with the InventoryTerminals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryTerminals

`func (o *GetTerminalsUnderAccountResponse) SetInventoryTerminals(v []string)`

SetInventoryTerminals sets InventoryTerminals field to given value.

### HasInventoryTerminals

`func (o *GetTerminalsUnderAccountResponse) HasInventoryTerminals() bool`

HasInventoryTerminals returns a boolean if a field has been set.

### GetMerchantAccounts

`func (o *GetTerminalsUnderAccountResponse) GetMerchantAccounts() []MerchantAccount`

GetMerchantAccounts returns the MerchantAccounts field if non-nil, zero value otherwise.

### GetMerchantAccountsOk

`func (o *GetTerminalsUnderAccountResponse) GetMerchantAccountsOk() (*[]MerchantAccount, bool)`

GetMerchantAccountsOk returns a tuple with the MerchantAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccounts

`func (o *GetTerminalsUnderAccountResponse) SetMerchantAccounts(v []MerchantAccount)`

SetMerchantAccounts sets MerchantAccounts field to given value.

### HasMerchantAccounts

`func (o *GetTerminalsUnderAccountResponse) HasMerchantAccounts() bool`

HasMerchantAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


