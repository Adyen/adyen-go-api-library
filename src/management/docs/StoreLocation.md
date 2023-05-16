# StoreLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | The name of the city. | [optional] 
**Country** | **string** | The two-letter country code in [ISO_3166-1_alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format. | 
**Line1** | Pointer to **string** | The street address. | [optional] 
**Line2** | Pointer to **string** | Second address line. | [optional] 
**Line3** | Pointer to **string** | Third address line. | [optional] 
**PostalCode** | Pointer to **string** | The postal code. | [optional] 
**StateOrProvince** | Pointer to **string** | The state or province code as defined in [ISO 3166-2](https://www.iso.org/standard/72483.html). For example, **ON** for Ontario, Canada.  Required for the following countries:  - Australia - Brazil - Canada - India - Mexico - New Zealand - United States | [optional] 

## Methods

### NewStoreLocation

`func NewStoreLocation(country string, ) *StoreLocation`

NewStoreLocation instantiates a new StoreLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreLocationWithDefaults

`func NewStoreLocationWithDefaults() *StoreLocation`

NewStoreLocationWithDefaults instantiates a new StoreLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *StoreLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *StoreLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *StoreLocation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *StoreLocation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *StoreLocation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *StoreLocation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *StoreLocation) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetLine1

`func (o *StoreLocation) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *StoreLocation) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *StoreLocation) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *StoreLocation) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### GetLine2

`func (o *StoreLocation) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *StoreLocation) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *StoreLocation) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *StoreLocation) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetLine3

`func (o *StoreLocation) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *StoreLocation) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *StoreLocation) SetLine3(v string)`

SetLine3 sets Line3 field to given value.

### HasLine3

`func (o *StoreLocation) HasLine3() bool`

HasLine3 returns a boolean if a field has been set.

### GetPostalCode

`func (o *StoreLocation) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *StoreLocation) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *StoreLocation) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *StoreLocation) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetStateOrProvince

`func (o *StoreLocation) GetStateOrProvince() string`

GetStateOrProvince returns the StateOrProvince field if non-nil, zero value otherwise.

### GetStateOrProvinceOk

`func (o *StoreLocation) GetStateOrProvinceOk() (*string, bool)`

GetStateOrProvinceOk returns a tuple with the StateOrProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvince

`func (o *StoreLocation) SetStateOrProvince(v string)`

SetStateOrProvince sets StateOrProvince field to given value.

### HasStateOrProvince

`func (o *StoreLocation) HasStateOrProvince() bool`

HasStateOrProvince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


