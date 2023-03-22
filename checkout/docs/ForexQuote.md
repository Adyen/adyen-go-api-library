# ForexQuote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | The account name. | [optional] 
**AccountType** | Pointer to **string** | The account type. | [optional] 
**BaseAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**BasePoints** | **int32** | The base points. | 
**Buy** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**Interbank** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**Reference** | Pointer to **string** | The reference assigned to the forex quote request. | [optional] 
**Sell** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**Signature** | Pointer to **string** | The signature to validate the integrity. | [optional] 
**Source** | Pointer to **string** | The source of the forex quote. | [optional] 
**Type** | Pointer to **string** | The type of forex. | [optional] 
**ValidTill** | **time.Time** | The date until which the forex quote is valid. | 

## Methods

### NewForexQuote

`func NewForexQuote(basePoints int32, validTill time.Time, ) *ForexQuote`

NewForexQuote instantiates a new ForexQuote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForexQuoteWithDefaults

`func NewForexQuoteWithDefaults() *ForexQuote`

NewForexQuoteWithDefaults instantiates a new ForexQuote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ForexQuote) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ForexQuote) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ForexQuote) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ForexQuote) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccountType

`func (o *ForexQuote) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *ForexQuote) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *ForexQuote) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *ForexQuote) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBaseAmount

`func (o *ForexQuote) GetBaseAmount() Amount`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *ForexQuote) GetBaseAmountOk() (*Amount, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *ForexQuote) SetBaseAmount(v Amount)`

SetBaseAmount sets BaseAmount field to given value.

### HasBaseAmount

`func (o *ForexQuote) HasBaseAmount() bool`

HasBaseAmount returns a boolean if a field has been set.

### GetBasePoints

`func (o *ForexQuote) GetBasePoints() int32`

GetBasePoints returns the BasePoints field if non-nil, zero value otherwise.

### GetBasePointsOk

`func (o *ForexQuote) GetBasePointsOk() (*int32, bool)`

GetBasePointsOk returns a tuple with the BasePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePoints

`func (o *ForexQuote) SetBasePoints(v int32)`

SetBasePoints sets BasePoints field to given value.


### GetBuy

`func (o *ForexQuote) GetBuy() Amount`

GetBuy returns the Buy field if non-nil, zero value otherwise.

### GetBuyOk

`func (o *ForexQuote) GetBuyOk() (*Amount, bool)`

GetBuyOk returns a tuple with the Buy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuy

`func (o *ForexQuote) SetBuy(v Amount)`

SetBuy sets Buy field to given value.

### HasBuy

`func (o *ForexQuote) HasBuy() bool`

HasBuy returns a boolean if a field has been set.

### GetInterbank

`func (o *ForexQuote) GetInterbank() Amount`

GetInterbank returns the Interbank field if non-nil, zero value otherwise.

### GetInterbankOk

`func (o *ForexQuote) GetInterbankOk() (*Amount, bool)`

GetInterbankOk returns a tuple with the Interbank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterbank

`func (o *ForexQuote) SetInterbank(v Amount)`

SetInterbank sets Interbank field to given value.

### HasInterbank

`func (o *ForexQuote) HasInterbank() bool`

HasInterbank returns a boolean if a field has been set.

### GetReference

`func (o *ForexQuote) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ForexQuote) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ForexQuote) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ForexQuote) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSell

`func (o *ForexQuote) GetSell() Amount`

GetSell returns the Sell field if non-nil, zero value otherwise.

### GetSellOk

`func (o *ForexQuote) GetSellOk() (*Amount, bool)`

GetSellOk returns a tuple with the Sell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSell

`func (o *ForexQuote) SetSell(v Amount)`

SetSell sets Sell field to given value.

### HasSell

`func (o *ForexQuote) HasSell() bool`

HasSell returns a boolean if a field has been set.

### GetSignature

`func (o *ForexQuote) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ForexQuote) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ForexQuote) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *ForexQuote) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSource

`func (o *ForexQuote) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ForexQuote) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ForexQuote) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ForexQuote) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *ForexQuote) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ForexQuote) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ForexQuote) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ForexQuote) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValidTill

`func (o *ForexQuote) GetValidTill() time.Time`

GetValidTill returns the ValidTill field if non-nil, zero value otherwise.

### GetValidTillOk

`func (o *ForexQuote) GetValidTillOk() (*time.Time, bool)`

GetValidTillOk returns a tuple with the ValidTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTill

`func (o *ForexQuote) SetValidTill(v time.Time)`

SetValidTill sets ValidTill field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


