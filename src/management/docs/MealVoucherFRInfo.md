# MealVoucherFRInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConecsId** | **string** | Meal Voucher conecsId. Format: digits only | 
**Siret** | **string** | Meal Voucher siret. Format: 14 digits. | 
**SubTypes** | **[]string** | The list of additional payment methods. Allowed values: **mealVoucher_FR_endenred**, **mealVoucher_FR_groupeup**, **mealVoucher_FR_natixis**, **mealVoucher_FR_sodexo**. | 

## Methods

### NewMealVoucherFRInfo

`func NewMealVoucherFRInfo(conecsId string, siret string, subTypes []string, ) *MealVoucherFRInfo`

NewMealVoucherFRInfo instantiates a new MealVoucherFRInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMealVoucherFRInfoWithDefaults

`func NewMealVoucherFRInfoWithDefaults() *MealVoucherFRInfo`

NewMealVoucherFRInfoWithDefaults instantiates a new MealVoucherFRInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConecsId

`func (o *MealVoucherFRInfo) GetConecsId() string`

GetConecsId returns the ConecsId field if non-nil, zero value otherwise.

### GetConecsIdOk

`func (o *MealVoucherFRInfo) GetConecsIdOk() (*string, bool)`

GetConecsIdOk returns a tuple with the ConecsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConecsId

`func (o *MealVoucherFRInfo) SetConecsId(v string)`

SetConecsId sets ConecsId field to given value.


### GetSiret

`func (o *MealVoucherFRInfo) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *MealVoucherFRInfo) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *MealVoucherFRInfo) SetSiret(v string)`

SetSiret sets Siret field to given value.


### GetSubTypes

`func (o *MealVoucherFRInfo) GetSubTypes() []string`

GetSubTypes returns the SubTypes field if non-nil, zero value otherwise.

### GetSubTypesOk

`func (o *MealVoucherFRInfo) GetSubTypesOk() (*[]string, bool)`

GetSubTypesOk returns a tuple with the SubTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTypes

`func (o *MealVoucherFRInfo) SetSubTypes(v []string)`

SetSubTypes sets SubTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


