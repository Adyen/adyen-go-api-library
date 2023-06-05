# NameLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | The city where the merchant is located. | [optional] 
**Country** | Pointer to **string** | The country where the merchant is located in [three-letter country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3) format. | [optional] 
**CountryOfOrigin** | Pointer to **string** | The home country in [three-digit country code](https://en.wikipedia.org/wiki/ISO_3166-1_numeric) format, used for government-controlled merchants such as embassies. | [optional] 
**Name** | Pointer to **string** | The name of the merchant&#39;s shop or service. | [optional] 
**RawData** | Pointer to **string** | The raw data. | [optional] 
**State** | Pointer to **string** | The state where the merchant is located. | [optional] 

## Methods

### NewNameLocation

`func NewNameLocation() *NameLocation`

NewNameLocation instantiates a new NameLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNameLocationWithDefaults

`func NewNameLocationWithDefaults() *NameLocation`

NewNameLocationWithDefaults instantiates a new NameLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *NameLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *NameLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *NameLocation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *NameLocation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *NameLocation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *NameLocation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *NameLocation) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *NameLocation) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryOfOrigin

`func (o *NameLocation) GetCountryOfOrigin() string`

GetCountryOfOrigin returns the CountryOfOrigin field if non-nil, zero value otherwise.

### GetCountryOfOriginOk

`func (o *NameLocation) GetCountryOfOriginOk() (*string, bool)`

GetCountryOfOriginOk returns a tuple with the CountryOfOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfOrigin

`func (o *NameLocation) SetCountryOfOrigin(v string)`

SetCountryOfOrigin sets CountryOfOrigin field to given value.

### HasCountryOfOrigin

`func (o *NameLocation) HasCountryOfOrigin() bool`

HasCountryOfOrigin returns a boolean if a field has been set.

### GetName

`func (o *NameLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NameLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NameLocation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NameLocation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRawData

`func (o *NameLocation) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *NameLocation) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *NameLocation) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *NameLocation) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetState

`func (o *NameLocation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NameLocation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NameLocation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NameLocation) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


