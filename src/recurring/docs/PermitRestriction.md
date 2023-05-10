# PermitRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**SingleTransactionLimit** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**SingleUse** | Pointer to **bool** | Only a single payment can be made using this permit if set to true, otherwise multiple payments are allowed. | [optional] 

## Methods

### NewPermitRestriction

`func NewPermitRestriction() *PermitRestriction`

NewPermitRestriction instantiates a new PermitRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermitRestrictionWithDefaults

`func NewPermitRestrictionWithDefaults() *PermitRestriction`

NewPermitRestrictionWithDefaults instantiates a new PermitRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxAmount

`func (o *PermitRestriction) GetMaxAmount() Amount`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *PermitRestriction) GetMaxAmountOk() (*Amount, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *PermitRestriction) SetMaxAmount(v Amount)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *PermitRestriction) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### GetSingleTransactionLimit

`func (o *PermitRestriction) GetSingleTransactionLimit() Amount`

GetSingleTransactionLimit returns the SingleTransactionLimit field if non-nil, zero value otherwise.

### GetSingleTransactionLimitOk

`func (o *PermitRestriction) GetSingleTransactionLimitOk() (*Amount, bool)`

GetSingleTransactionLimitOk returns a tuple with the SingleTransactionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleTransactionLimit

`func (o *PermitRestriction) SetSingleTransactionLimit(v Amount)`

SetSingleTransactionLimit sets SingleTransactionLimit field to given value.

### HasSingleTransactionLimit

`func (o *PermitRestriction) HasSingleTransactionLimit() bool`

HasSingleTransactionLimit returns a boolean if a field has been set.

### GetSingleUse

`func (o *PermitRestriction) GetSingleUse() bool`

GetSingleUse returns the SingleUse field if non-nil, zero value otherwise.

### GetSingleUseOk

`func (o *PermitRestriction) GetSingleUseOk() (*bool, bool)`

GetSingleUseOk returns a tuple with the SingleUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUse

`func (o *PermitRestriction) SetSingleUse(v bool)`

SetSingleUse sets SingleUse field to given value.

### HasSingleUse

`func (o *PermitRestriction) HasSingleUse() bool`

HasSingleUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


