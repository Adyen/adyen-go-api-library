# CheckoutSessionInstallmentOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plans** | Pointer to **[]string** | Defines the type of installment plan. If not set, defaults to **regular**.  Possible values: * **regular** * **revolving** | [optional] 
**PreselectedValue** | Pointer to **int32** | Preselected number of installments offered for this payment method. | [optional] 
**Values** | Pointer to **[]int32** | An array of the number of installments that the shopper can choose from. For example, **[2,3,5]**. This cannot be specified simultaneously with &#x60;maxValue&#x60;. | [optional] 

## Methods

### NewCheckoutSessionInstallmentOption

`func NewCheckoutSessionInstallmentOption() *CheckoutSessionInstallmentOption`

NewCheckoutSessionInstallmentOption instantiates a new CheckoutSessionInstallmentOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSessionInstallmentOptionWithDefaults

`func NewCheckoutSessionInstallmentOptionWithDefaults() *CheckoutSessionInstallmentOption`

NewCheckoutSessionInstallmentOptionWithDefaults instantiates a new CheckoutSessionInstallmentOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlans

`func (o *CheckoutSessionInstallmentOption) GetPlans() []string`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *CheckoutSessionInstallmentOption) GetPlansOk() (*[]string, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *CheckoutSessionInstallmentOption) SetPlans(v []string)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *CheckoutSessionInstallmentOption) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetPreselectedValue

`func (o *CheckoutSessionInstallmentOption) GetPreselectedValue() int32`

GetPreselectedValue returns the PreselectedValue field if non-nil, zero value otherwise.

### GetPreselectedValueOk

`func (o *CheckoutSessionInstallmentOption) GetPreselectedValueOk() (*int32, bool)`

GetPreselectedValueOk returns a tuple with the PreselectedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreselectedValue

`func (o *CheckoutSessionInstallmentOption) SetPreselectedValue(v int32)`

SetPreselectedValue sets PreselectedValue field to given value.

### HasPreselectedValue

`func (o *CheckoutSessionInstallmentOption) HasPreselectedValue() bool`

HasPreselectedValue returns a boolean if a field has been set.

### GetValues

`func (o *CheckoutSessionInstallmentOption) GetValues() []int32`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CheckoutSessionInstallmentOption) GetValuesOk() (*[]int32, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CheckoutSessionInstallmentOption) SetValues(v []int32)`

SetValues sets Values field to given value.

### HasValues

`func (o *CheckoutSessionInstallmentOption) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


