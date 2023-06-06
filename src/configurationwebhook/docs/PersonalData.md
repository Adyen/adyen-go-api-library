# PersonalData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateOfBirth** | Pointer to **string** | The date of birth of the person. The date should be in ISO-8601 format yyyy-mm-dd (e.g. 2000-01-31). | [optional] 
**IdNumber** | Pointer to **string** | An ID number of the person. | [optional] 
**Nationality** | Pointer to **string** | The nationality of the person represented by a two-character country code. &gt;The permitted country codes are defined in ISO-3166-1 alpha-2 (e.g. &#39;NL&#39;). | [optional] 

## Methods

### NewPersonalData

`func NewPersonalData() *PersonalData`

NewPersonalData instantiates a new PersonalData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalDataWithDefaults

`func NewPersonalDataWithDefaults() *PersonalData`

NewPersonalDataWithDefaults instantiates a new PersonalData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateOfBirth

`func (o *PersonalData) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *PersonalData) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *PersonalData) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *PersonalData) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetIdNumber

`func (o *PersonalData) GetIdNumber() string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *PersonalData) GetIdNumberOk() (*string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *PersonalData) SetIdNumber(v string)`

SetIdNumber sets IdNumber field to given value.

### HasIdNumber

`func (o *PersonalData) HasIdNumber() bool`

HasIdNumber returns a boolean if a field has been set.

### GetNationality

`func (o *PersonalData) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *PersonalData) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *PersonalData) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *PersonalData) HasNationality() bool`

HasNationality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


