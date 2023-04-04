# InstallmentOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxValue** | Pointer to **int32** | The maximum number of installments offered for this payment method. | [optional] 
**Plans** | Pointer to **[]string** | Defines the type of installment plan. If not set, defaults to **regular**.  Possible values: * **regular** * **revolving** | [optional] 
**PreselectedValue** | Pointer to **int32** | Preselected number of installments offered for this payment method. | [optional] 
**Values** | Pointer to **[]int32** | An array of the number of installments that the shopper can choose from. For example, **[2,3,5]**. This cannot be specified simultaneously with &#x60;maxValue&#x60;. | [optional] 

## Methods

### NewInstallmentOption

`func NewInstallmentOption() *InstallmentOption`

NewInstallmentOption instantiates a new InstallmentOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallmentOptionWithDefaults

`func NewInstallmentOptionWithDefaults() *InstallmentOption`

NewInstallmentOptionWithDefaults instantiates a new InstallmentOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxValue

`func (o *InstallmentOption) GetMaxValue() int32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *InstallmentOption) GetMaxValueOk() (*int32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *InstallmentOption) SetMaxValue(v int32)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *InstallmentOption) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetPlans

`func (o *InstallmentOption) GetPlans() []string`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *InstallmentOption) GetPlansOk() (*[]string, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *InstallmentOption) SetPlans(v []string)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *InstallmentOption) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetPreselectedValue

`func (o *InstallmentOption) GetPreselectedValue() int32`

GetPreselectedValue returns the PreselectedValue field if non-nil, zero value otherwise.

### GetPreselectedValueOk

`func (o *InstallmentOption) GetPreselectedValueOk() (*int32, bool)`

GetPreselectedValueOk returns a tuple with the PreselectedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreselectedValue

`func (o *InstallmentOption) SetPreselectedValue(v int32)`

SetPreselectedValue sets PreselectedValue field to given value.

### HasPreselectedValue

`func (o *InstallmentOption) HasPreselectedValue() bool`

HasPreselectedValue returns a boolean if a field has been set.

### GetValues

`func (o *InstallmentOption) GetValues() []int32`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *InstallmentOption) GetValuesOk() (*[]int32, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *InstallmentOption) SetValues(v []int32)`

SetValues sets Values field to given value.

### HasValues

`func (o *InstallmentOption) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


