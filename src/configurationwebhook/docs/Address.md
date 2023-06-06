# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | **string** | The name of the city. Maximum length: 3000 characters. | 
**Country** | **string** | The two-character ISO-3166-1 alpha-2 country code. For example, **US**. &gt; If you don&#39;t know the country or are not collecting the country from the shopper, provide &#x60;country&#x60; as &#x60;ZZ&#x60;. | 
**HouseNumberOrName** | **string** | The number or name of the house. Maximum length: 3000 characters. | 
**PostalCode** | **string** | A maximum of five digits for an address in the US, or a maximum of ten characters for an address in all other countries. | 
**StateOrProvince** | Pointer to **string** | The two-character ISO 3166-2 state or province code. For example, **CA** in the US or **ON** in Canada. &gt; Required for the US and Canada. | [optional] 
**Street** | **string** | The name of the street. Maximum length: 3000 characters. &gt; The house number should not be included in this field; it should be separately provided via &#x60;houseNumberOrName&#x60;. | 

## Methods

### NewAddress

`func NewAddress(city string, country string, houseNumberOrName string, postalCode string, street string, ) *Address`

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


### GetHouseNumberOrName

`func (o *Address) GetHouseNumberOrName() string`

GetHouseNumberOrName returns the HouseNumberOrName field if non-nil, zero value otherwise.

### GetHouseNumberOrNameOk

`func (o *Address) GetHouseNumberOrNameOk() (*string, bool)`

GetHouseNumberOrNameOk returns a tuple with the HouseNumberOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumberOrName

`func (o *Address) SetHouseNumberOrName(v string)`

SetHouseNumberOrName sets HouseNumberOrName field to given value.


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

### GetStreet

`func (o *Address) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *Address) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *Address) SetStreet(v string)`

SetStreet sets Street field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


