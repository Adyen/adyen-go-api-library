# PhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneCountryCode** | Pointer to **string** | The two-character ISO-3166-1 alpha-2 country code of the phone number. For example, **US** or **NL**. | [optional] 
**PhoneNumber** | Pointer to **string** | The phone number. The inclusion of the phone number country code is not necessary. | [optional] 
**PhoneType** | Pointer to **string** | The type of the phone number. Possible values: **Landline**, **Mobile**, **SIP**, **Fax**. | [optional] 

## Methods

### NewPhoneNumber

`func NewPhoneNumber() *PhoneNumber`

NewPhoneNumber instantiates a new PhoneNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneNumberWithDefaults

`func NewPhoneNumberWithDefaults() *PhoneNumber`

NewPhoneNumberWithDefaults instantiates a new PhoneNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneCountryCode

`func (o *PhoneNumber) GetPhoneCountryCode() string`

GetPhoneCountryCode returns the PhoneCountryCode field if non-nil, zero value otherwise.

### GetPhoneCountryCodeOk

`func (o *PhoneNumber) GetPhoneCountryCodeOk() (*string, bool)`

GetPhoneCountryCodeOk returns a tuple with the PhoneCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCountryCode

`func (o *PhoneNumber) SetPhoneCountryCode(v string)`

SetPhoneCountryCode sets PhoneCountryCode field to given value.

### HasPhoneCountryCode

`func (o *PhoneNumber) HasPhoneCountryCode() bool`

HasPhoneCountryCode returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *PhoneNumber) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PhoneNumber) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PhoneNumber) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PhoneNumber) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPhoneType

`func (o *PhoneNumber) GetPhoneType() string`

GetPhoneType returns the PhoneType field if non-nil, zero value otherwise.

### GetPhoneTypeOk

`func (o *PhoneNumber) GetPhoneTypeOk() (*string, bool)`

GetPhoneTypeOk returns a tuple with the PhoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneType

`func (o *PhoneNumber) SetPhoneType(v string)`

SetPhoneType sets PhoneType field to given value.

### HasPhoneType

`func (o *PhoneNumber) HasPhoneType() bool`

HasPhoneType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


