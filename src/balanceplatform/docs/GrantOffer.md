# GrantOffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolderId** | **string** | The identifier of the account holder to which the grant is offered. | 
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**ContractType** | Pointer to **string** | The contract type of the grant offer. Possible value: **cashAdvance**, **loan**. | [optional] 
**ExpiresAt** | Pointer to **map[string]interface{}** |  | [optional] 
**Fee** | Pointer to [**Fee**](Fee.md) |  | [optional] 
**Id** | Pointer to **string** | The unique identifier of the grant offer. | [optional] 
**Repayment** | Pointer to [**Repayment**](Repayment.md) |  | [optional] 
**StartsAt** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGrantOffer

`func NewGrantOffer(accountHolderId string, ) *GrantOffer`

NewGrantOffer instantiates a new GrantOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantOfferWithDefaults

`func NewGrantOfferWithDefaults() *GrantOffer`

NewGrantOfferWithDefaults instantiates a new GrantOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolderId

`func (o *GrantOffer) GetAccountHolderId() string`

GetAccountHolderId returns the AccountHolderId field if non-nil, zero value otherwise.

### GetAccountHolderIdOk

`func (o *GrantOffer) GetAccountHolderIdOk() (*string, bool)`

GetAccountHolderIdOk returns a tuple with the AccountHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderId

`func (o *GrantOffer) SetAccountHolderId(v string)`

SetAccountHolderId sets AccountHolderId field to given value.


### GetAmount

`func (o *GrantOffer) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GrantOffer) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GrantOffer) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GrantOffer) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetContractType

`func (o *GrantOffer) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *GrantOffer) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *GrantOffer) SetContractType(v string)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *GrantOffer) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GrantOffer) GetExpiresAt() map[string]interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GrantOffer) GetExpiresAtOk() (*map[string]interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GrantOffer) SetExpiresAt(v map[string]interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GrantOffer) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFee

`func (o *GrantOffer) GetFee() Fee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GrantOffer) GetFeeOk() (*Fee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GrantOffer) SetFee(v Fee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GrantOffer) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetId

`func (o *GrantOffer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GrantOffer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GrantOffer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GrantOffer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRepayment

`func (o *GrantOffer) GetRepayment() Repayment`

GetRepayment returns the Repayment field if non-nil, zero value otherwise.

### GetRepaymentOk

`func (o *GrantOffer) GetRepaymentOk() (*Repayment, bool)`

GetRepaymentOk returns a tuple with the Repayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepayment

`func (o *GrantOffer) SetRepayment(v Repayment)`

SetRepayment sets Repayment field to given value.

### HasRepayment

`func (o *GrantOffer) HasRepayment() bool`

HasRepayment returns a boolean if a field has been set.

### GetStartsAt

`func (o *GrantOffer) GetStartsAt() map[string]interface{}`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *GrantOffer) GetStartsAtOk() (*map[string]interface{}, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *GrantOffer) SetStartsAt(v map[string]interface{})`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *GrantOffer) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


