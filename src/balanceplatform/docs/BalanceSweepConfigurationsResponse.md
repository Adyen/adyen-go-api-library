# BalanceSweepConfigurationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNext** | **bool** | Indicates whether there are more items on the next page. | 
**HasPrevious** | **bool** | Indicates whether there are more items on the previous page. | 
**Sweeps** | [**[]SweepConfigurationV2**](SweepConfigurationV2.md) | List of sweeps associated with the balance account. | 

## Methods

### NewBalanceSweepConfigurationsResponse

`func NewBalanceSweepConfigurationsResponse(hasNext bool, hasPrevious bool, sweeps []SweepConfigurationV2, ) *BalanceSweepConfigurationsResponse`

NewBalanceSweepConfigurationsResponse instantiates a new BalanceSweepConfigurationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceSweepConfigurationsResponseWithDefaults

`func NewBalanceSweepConfigurationsResponseWithDefaults() *BalanceSweepConfigurationsResponse`

NewBalanceSweepConfigurationsResponseWithDefaults instantiates a new BalanceSweepConfigurationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNext

`func (o *BalanceSweepConfigurationsResponse) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *BalanceSweepConfigurationsResponse) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *BalanceSweepConfigurationsResponse) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.


### GetHasPrevious

`func (o *BalanceSweepConfigurationsResponse) GetHasPrevious() bool`

GetHasPrevious returns the HasPrevious field if non-nil, zero value otherwise.

### GetHasPreviousOk

`func (o *BalanceSweepConfigurationsResponse) GetHasPreviousOk() (*bool, bool)`

GetHasPreviousOk returns a tuple with the HasPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrevious

`func (o *BalanceSweepConfigurationsResponse) SetHasPrevious(v bool)`

SetHasPrevious sets HasPrevious field to given value.


### GetSweeps

`func (o *BalanceSweepConfigurationsResponse) GetSweeps() []SweepConfigurationV2`

GetSweeps returns the Sweeps field if non-nil, zero value otherwise.

### GetSweepsOk

`func (o *BalanceSweepConfigurationsResponse) GetSweepsOk() (*[]SweepConfigurationV2, bool)`

GetSweepsOk returns a tuple with the Sweeps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweeps

`func (o *BalanceSweepConfigurationsResponse) SetSweeps(v []SweepConfigurationV2)`

SetSweeps sets Sweeps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


