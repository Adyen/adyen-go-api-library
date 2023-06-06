# Expiry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | Pointer to **string** | The month in which the card will expire. | [optional] 
**Year** | Pointer to **string** | The year in which the card will expire. | [optional] 

## Methods

### NewExpiry

`func NewExpiry() *Expiry`

NewExpiry instantiates a new Expiry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiryWithDefaults

`func NewExpiryWithDefaults() *Expiry`

NewExpiryWithDefaults instantiates a new Expiry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *Expiry) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *Expiry) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *Expiry) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *Expiry) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetYear

`func (o *Expiry) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *Expiry) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *Expiry) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *Expiry) HasYear() bool`

HasYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


