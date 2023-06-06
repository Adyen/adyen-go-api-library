# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | The name of the city. | [optional] 
**Country** | **string** | The two-character ISO 3166-1 alpha-2 country code. For example, **US**, **NL**, or **GB**. | 
**Line1** | Pointer to **string** | First line of the street address. | [optional] 
**Line2** | Pointer to **string** | Second line of the street address. | [optional] 
**PostalCode** | Pointer to **string** | The postal code. Maximum length: * 5 digits for an address in the US. * 10 characters for an address in all other countries. | [optional] 
**StateOrProvince** | Pointer to **string** | The two-letter ISO 3166-2 state or province code. For example, **CA** in the US or **ON** in Canada. &gt; Required for the US and Canada. | [optional] 

## Methods

### NewAddress

`func NewAddress(country string, ) *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *Address) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Address) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Address) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Address) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *Address) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Address) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Address) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetLine1

`func (o *Address) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *Address) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *Address) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *Address) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### GetLine2

`func (o *Address) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *Address) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *Address) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *Address) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetPostalCode

`func (o *Address) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Address) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Address) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *Address) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetStateOrProvince

`func (o *Address) GetStateOrProvince() string`

GetStateOrProvince returns the StateOrProvince field if non-nil, zero value otherwise.

### GetStateOrProvinceOk

`func (o *Address) GetStateOrProvinceOk() (*string, bool)`

GetStateOrProvinceOk returns a tuple with the StateOrProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvince

`func (o *Address) SetStateOrProvince(v string)`

SetStateOrProvince sets StateOrProvince field to given value.

### HasStateOrProvince

`func (o *Address) HasStateOrProvince() bool`

HasStateOrProvince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


