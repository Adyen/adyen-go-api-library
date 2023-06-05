# MerchantData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mcc** | Pointer to **string** | The merchant category code. | [optional] 
**MerchantId** | Pointer to **string** | The merchant identifier. | [optional] 
**NameLocation** | Pointer to [**NameLocation**](NameLocation.md) |  | [optional] 
**PostalCode** | Pointer to **string** | The merchant postal code. | [optional] 

## Methods

### NewMerchantData

`func NewMerchantData() *MerchantData`

NewMerchantData instantiates a new MerchantData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantDataWithDefaults

`func NewMerchantDataWithDefaults() *MerchantData`

NewMerchantDataWithDefaults instantiates a new MerchantData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMcc

`func (o *MerchantData) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *MerchantData) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *MerchantData) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *MerchantData) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMerchantId

`func (o *MerchantData) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *MerchantData) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *MerchantData) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *MerchantData) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetNameLocation

`func (o *MerchantData) GetNameLocation() NameLocation`

GetNameLocation returns the NameLocation field if non-nil, zero value otherwise.

### GetNameLocationOk

`func (o *MerchantData) GetNameLocationOk() (*NameLocation, bool)`

GetNameLocationOk returns a tuple with the NameLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocation

`func (o *MerchantData) SetNameLocation(v NameLocation)`

SetNameLocation sets NameLocation field to given value.

### HasNameLocation

`func (o *MerchantData) HasNameLocation() bool`

HasNameLocation returns a boolean if a field has been set.

### GetPostalCode

`func (o *MerchantData) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *MerchantData) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *MerchantData) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *MerchantData) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


