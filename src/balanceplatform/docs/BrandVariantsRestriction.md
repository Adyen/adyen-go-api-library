# BrandVariantsRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | Defines how the condition must be evaluated. | 
**Value** | Pointer to **[]string** | List of card brand variants.  Possible values:   - **mc**, **mccredit**, **mccommercialcredit_b2b**, **mcdebit**, **mcbusinessdebit**, **mcbusinessworlddebit**, **mcprepaid**, **mcmaestro**   - **visa**, **visacredit**, **visadebit**, **visaprepaid**.  You can specify a rule for a generic variant. For example, to create a rule for all Mastercard payment instruments, use **mc**. The rule is applied to all payment instruments under **mc**, such as **mcbusinessdebit** and **mcdebit**.   | [optional] 

## Methods

### NewBrandVariantsRestriction

`func NewBrandVariantsRestriction(operation string, ) *BrandVariantsRestriction`

NewBrandVariantsRestriction instantiates a new BrandVariantsRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandVariantsRestrictionWithDefaults

`func NewBrandVariantsRestrictionWithDefaults() *BrandVariantsRestriction`

NewBrandVariantsRestrictionWithDefaults instantiates a new BrandVariantsRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *BrandVariantsRestriction) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BrandVariantsRestriction) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *BrandVariantsRestriction) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetValue

`func (o *BrandVariantsRestriction) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BrandVariantsRestriction) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BrandVariantsRestriction) SetValue(v []string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BrandVariantsRestriction) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


