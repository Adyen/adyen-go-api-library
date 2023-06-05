# AmountAdjustment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**AmountAdjustmentType** | Pointer to **string** | The type of markup that is applied to an authorised payment.  Possible values: **exchange**, **forexMarkup**, **authHoldReserve**, **atmMarkup**. | [optional] 
**Basepoints** | Pointer to **int32** | The basepoints associated with the applied markup. | [optional] 

## Methods

### NewAmountAdjustment

`func NewAmountAdjustment() *AmountAdjustment`

NewAmountAdjustment instantiates a new AmountAdjustment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmountAdjustmentWithDefaults

`func NewAmountAdjustmentWithDefaults() *AmountAdjustment`

NewAmountAdjustmentWithDefaults instantiates a new AmountAdjustment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AmountAdjustment) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AmountAdjustment) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AmountAdjustment) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AmountAdjustment) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountAdjustmentType

`func (o *AmountAdjustment) GetAmountAdjustmentType() string`

GetAmountAdjustmentType returns the AmountAdjustmentType field if non-nil, zero value otherwise.

### GetAmountAdjustmentTypeOk

`func (o *AmountAdjustment) GetAmountAdjustmentTypeOk() (*string, bool)`

GetAmountAdjustmentTypeOk returns a tuple with the AmountAdjustmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAdjustmentType

`func (o *AmountAdjustment) SetAmountAdjustmentType(v string)`

SetAmountAdjustmentType sets AmountAdjustmentType field to given value.

### HasAmountAdjustmentType

`func (o *AmountAdjustment) HasAmountAdjustmentType() bool`

HasAmountAdjustmentType returns a boolean if a field has been set.

### GetBasepoints

`func (o *AmountAdjustment) GetBasepoints() int32`

GetBasepoints returns the Basepoints field if non-nil, zero value otherwise.

### GetBasepointsOk

`func (o *AmountAdjustment) GetBasepointsOk() (*int32, bool)`

GetBasepointsOk returns a tuple with the Basepoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasepoints

`func (o *AmountAdjustment) SetBasepoints(v int32)`

SetBasepoints sets Basepoints field to given value.

### HasBasepoints

`func (o *AmountAdjustment) HasBasepoints() bool`

HasBasepoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


