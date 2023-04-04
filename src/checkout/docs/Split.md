# Split

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Unique identifier of the account where the split amount should be sent. This is required if &#x60;type&#x60; is **MarketPlace** or **BalanceAccount**.   | [optional] 
**Amount** | [**SplitAmount**](SplitAmount.md) |  | 
**Description** | Pointer to **string** | A description of this split. | [optional] 
**Reference** | Pointer to **string** | Your reference for the split, which you can use to link the split to other operations such as captures and refunds.  This is required if &#x60;type&#x60; is **MarketPlace** or **BalanceAccount**. For the other types, we also recommend sending a reference so you can reconcile the split and the associated payment in the transaction overview and in the reports. If the reference is not provided, the split is reported as part of the aggregated [TransferBalance record type](https://docs.adyen.com/reporting/marketpay-payments-accounting-report) in Adyen for Platforms. | [optional] 
**Type** | **string** | The type of split. Possible values: **Default**, **PaymentFee**, **VAT**, **Commission**, **MarketPlace**, **BalanceAccount**, **Remainder**, **Surcharge**, **Tip**. | 

## Methods

### NewSplit

`func NewSplit(amount SplitAmount, type_ string, ) *Split`

NewSplit instantiates a new Split object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplitWithDefaults

`func NewSplitWithDefaults() *Split`

NewSplitWithDefaults instantiates a new Split object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Split) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Split) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Split) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Split) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAmount

`func (o *Split) GetAmount() SplitAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Split) GetAmountOk() (*SplitAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Split) SetAmount(v SplitAmount)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *Split) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Split) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Split) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Split) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReference

`func (o *Split) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Split) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Split) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *Split) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetType

`func (o *Split) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Split) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Split) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


