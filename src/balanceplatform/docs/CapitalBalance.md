# CapitalBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes). | 
**Fee** | **int64** | Fee amount. | 
**Principal** | **int64** | Principal amount. | 
**Total** | **int64** | Total amount. A sum of principal amount and fee amount. | 

## Methods

### NewCapitalBalance

`func NewCapitalBalance(currency string, fee int64, principal int64, total int64, ) *CapitalBalance`

NewCapitalBalance instantiates a new CapitalBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapitalBalanceWithDefaults

`func NewCapitalBalanceWithDefaults() *CapitalBalance`

NewCapitalBalanceWithDefaults instantiates a new CapitalBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CapitalBalance) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CapitalBalance) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CapitalBalance) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetFee

`func (o *CapitalBalance) GetFee() int64`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CapitalBalance) GetFeeOk() (*int64, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CapitalBalance) SetFee(v int64)`

SetFee sets Fee field to given value.


### GetPrincipal

`func (o *CapitalBalance) GetPrincipal() int64`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *CapitalBalance) GetPrincipalOk() (*int64, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *CapitalBalance) SetPrincipal(v int64)`

SetPrincipal sets Principal field to given value.


### GetTotal

`func (o *CapitalBalance) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CapitalBalance) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CapitalBalance) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


