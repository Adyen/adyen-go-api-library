# Address2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | The name of the city. | [optional] 
**Country** | **string** | The two-character ISO-3166-1 alpha-2 country code. For example, **US**. &gt;If you don&#39;t know the country or are not collecting the country from the shopper, provide &#x60;country&#x60; as &#x60;ZZ&#x60;. | 
**Line1** | Pointer to **string** | First line of the address. | [optional] 
**Line2** | Pointer to **string** | Second line of the address. | [optional] 
**Line3** | Pointer to **string** | Third line of the address. | [optional] 
**PostalCode** | Pointer to **string** | The postal code. Maximum length: * 5 digits for an address in the US. * 10 characters for an address in all other countries. | [optional] 
**StateOrProvince** | Pointer to **string** | The two-letterISO 3166-2 state or province code. For example, **CA** in the US or **ON** in Canada. &gt; Required for the US and Canada. | [optional] 

## Methods

### NewAddress2

`func NewAddress2(country string, ) *Address2`

NewAddress2 instantiates a new Address2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddress2WithDefaults

`func NewAddress2WithDefaults() *Address2`

NewAddress2WithDefaults instantiates a new Address2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *Address2) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Address2) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Address2) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Address2) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *Address2) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Address2) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Address2) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetLine1

`func (o *Address2) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *Address2) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *Address2) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *Address2) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### GetLine2

`func (o *Address2) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *Address2) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *Address2) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *Address2) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetLine3

`func (o *Address2) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *Address2) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *Address2) SetLine3(v string)`

SetLine3 sets Line3 field to given value.

### HasLine3

`func (o *Address2) HasLine3() bool`

HasLine3 returns a boolean if a field has been set.

### GetPostalCode

`func (o *Address2) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Address2) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Address2) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *Address2) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetStateOrProvince

`func (o *Address2) GetStateOrProvince() string`

GetStateOrProvince returns the StateOrProvince field if non-nil, zero value otherwise.

### GetStateOrProvinceOk

`func (o *Address2) GetStateOrProvinceOk() (*string, bool)`

GetStateOrProvinceOk returns a tuple with the StateOrProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvince

`func (o *Address2) SetStateOrProvince(v string)`

SetStateOrProvince sets StateOrProvince field to given value.

### HasStateOrProvince

`func (o *Address2) HasStateOrProvince() bool`

HasStateOrProvince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


