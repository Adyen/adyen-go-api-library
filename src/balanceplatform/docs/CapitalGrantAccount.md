# CapitalGrantAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balances** | Pointer to [**[]CapitalBalance**](CapitalBalance.md) | The balances of the grant account. | [optional] 
**FundingBalanceAccountId** | Pointer to **string** | The unique identifier of the balance account used to fund the grant. | [optional] 
**Id** | Pointer to **string** | The identifier of the grant account. | [optional] 
**Limits** | Pointer to [**[]GrantLimit**](GrantLimit.md) | The limits of the grant account. | [optional] 

## Methods

### NewCapitalGrantAccount

`func NewCapitalGrantAccount() *CapitalGrantAccount`

NewCapitalGrantAccount instantiates a new CapitalGrantAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapitalGrantAccountWithDefaults

`func NewCapitalGrantAccountWithDefaults() *CapitalGrantAccount`

NewCapitalGrantAccountWithDefaults instantiates a new CapitalGrantAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalances

`func (o *CapitalGrantAccount) GetBalances() []CapitalBalance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *CapitalGrantAccount) GetBalancesOk() (*[]CapitalBalance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *CapitalGrantAccount) SetBalances(v []CapitalBalance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *CapitalGrantAccount) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetFundingBalanceAccountId

`func (o *CapitalGrantAccount) GetFundingBalanceAccountId() string`

GetFundingBalanceAccountId returns the FundingBalanceAccountId field if non-nil, zero value otherwise.

### GetFundingBalanceAccountIdOk

`func (o *CapitalGrantAccount) GetFundingBalanceAccountIdOk() (*string, bool)`

GetFundingBalanceAccountIdOk returns a tuple with the FundingBalanceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingBalanceAccountId

`func (o *CapitalGrantAccount) SetFundingBalanceAccountId(v string)`

SetFundingBalanceAccountId sets FundingBalanceAccountId field to given value.

### HasFundingBalanceAccountId

`func (o *CapitalGrantAccount) HasFundingBalanceAccountId() bool`

HasFundingBalanceAccountId returns a boolean if a field has been set.

### GetId

`func (o *CapitalGrantAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CapitalGrantAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CapitalGrantAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CapitalGrantAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLimits

`func (o *CapitalGrantAccount) GetLimits() []GrantLimit`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *CapitalGrantAccount) GetLimitsOk() (*[]GrantLimit, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *CapitalGrantAccount) SetLimits(v []GrantLimit)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *CapitalGrantAccount) HasLimits() bool`

HasLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


