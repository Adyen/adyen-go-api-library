# BulkAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | The name of the city. | [optional] 
**Company** | Pointer to **string** | The name of the company. | [optional] 
**Country** | **string** | The two-character ISO-3166-1 alpha-2 country code. For example, **US**. | 
**Email** | Pointer to **string** | The email address. | [optional] 
**HouseNumberOrName** | Pointer to **string** | The house number or name. | [optional] 
**Mobile** | Pointer to **string** | The full telephone number. | [optional] 
**PostalCode** | Pointer to **string** | The postal code.  Maximum length:  * 5 digits for addresses in the US.  * 10 characters for all other countries. | [optional] 
**StateOrProvince** | Pointer to **string** | The two-letter ISO 3166-2 state or province code.  Maximum length: 2 characters for addresses in the US. | [optional] 
**Street** | Pointer to **string** | The streetname of the house. | [optional] 

## Methods

### NewBulkAddress

`func NewBulkAddress(country string, ) *BulkAddress`

NewBulkAddress instantiates a new BulkAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAddressWithDefaults

`func NewBulkAddressWithDefaults() *BulkAddress`

NewBulkAddressWithDefaults instantiates a new BulkAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *BulkAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BulkAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BulkAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *BulkAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompany

`func (o *BulkAddress) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *BulkAddress) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *BulkAddress) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *BulkAddress) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCountry

`func (o *BulkAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BulkAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BulkAddress) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetEmail

`func (o *BulkAddress) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BulkAddress) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BulkAddress) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BulkAddress) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHouseNumberOrName

`func (o *BulkAddress) GetHouseNumberOrName() string`

GetHouseNumberOrName returns the HouseNumberOrName field if non-nil, zero value otherwise.

### GetHouseNumberOrNameOk

`func (o *BulkAddress) GetHouseNumberOrNameOk() (*string, bool)`

GetHouseNumberOrNameOk returns a tuple with the HouseNumberOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumberOrName

`func (o *BulkAddress) SetHouseNumberOrName(v string)`

SetHouseNumberOrName sets HouseNumberOrName field to given value.

### HasHouseNumberOrName

`func (o *BulkAddress) HasHouseNumberOrName() bool`

HasHouseNumberOrName returns a boolean if a field has been set.

### GetMobile

`func (o *BulkAddress) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *BulkAddress) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *BulkAddress) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *BulkAddress) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetPostalCode

`func (o *BulkAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *BulkAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *BulkAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *BulkAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetStateOrProvince

`func (o *BulkAddress) GetStateOrProvince() string`

GetStateOrProvince returns the StateOrProvince field if non-nil, zero value otherwise.

### GetStateOrProvinceOk

`func (o *BulkAddress) GetStateOrProvinceOk() (*string, bool)`

GetStateOrProvinceOk returns a tuple with the StateOrProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvince

`func (o *BulkAddress) SetStateOrProvince(v string)`

SetStateOrProvince sets StateOrProvince field to given value.

### HasStateOrProvince

`func (o *BulkAddress) HasStateOrProvince() bool`

HasStateOrProvince returns a boolean if a field has been set.

### GetStreet

`func (o *BulkAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *BulkAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *BulkAddress) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *BulkAddress) HasStreet() bool`

HasStreet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


