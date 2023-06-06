# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Email** | Pointer to **string** | The e-mail address of the contact. | [optional] 
**FullPhoneNumber** | Pointer to **string** | The phone number of the contact provided as a single string.  It will be handled as a landline phone. **Examples:** \&quot;0031 6 11 22 33 44\&quot;, \&quot;+316/1122-3344\&quot;, \&quot;(0031) 611223344\&quot; | [optional] 
**Name** | Pointer to [**Name**](Name.md) |  | [optional] 
**PersonalData** | Pointer to [**PersonalData**](PersonalData.md) |  | [optional] 
**PhoneNumber** | Pointer to [**PhoneNumber**](PhoneNumber.md) |  | [optional] 
**WebAddress** | Pointer to **string** | The URL of the website of the contact. | [optional] 

## Methods

### NewContact

`func NewContact() *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Contact) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Contact) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Contact) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Contact) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetEmail

`func (o *Contact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Contact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Contact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Contact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFullPhoneNumber

`func (o *Contact) GetFullPhoneNumber() string`

GetFullPhoneNumber returns the FullPhoneNumber field if non-nil, zero value otherwise.

### GetFullPhoneNumberOk

`func (o *Contact) GetFullPhoneNumberOk() (*string, bool)`

GetFullPhoneNumberOk returns a tuple with the FullPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPhoneNumber

`func (o *Contact) SetFullPhoneNumber(v string)`

SetFullPhoneNumber sets FullPhoneNumber field to given value.

### HasFullPhoneNumber

`func (o *Contact) HasFullPhoneNumber() bool`

HasFullPhoneNumber returns a boolean if a field has been set.

### GetName

`func (o *Contact) GetName() Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Contact) GetNameOk() (*Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Contact) SetName(v Name)`

SetName sets Name field to given value.

### HasName

`func (o *Contact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPersonalData

`func (o *Contact) GetPersonalData() PersonalData`

GetPersonalData returns the PersonalData field if non-nil, zero value otherwise.

### GetPersonalDataOk

`func (o *Contact) GetPersonalDataOk() (*PersonalData, bool)`

GetPersonalDataOk returns a tuple with the PersonalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalData

`func (o *Contact) SetPersonalData(v PersonalData)`

SetPersonalData sets PersonalData field to given value.

### HasPersonalData

`func (o *Contact) HasPersonalData() bool`

HasPersonalData returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Contact) GetPhoneNumber() PhoneNumber`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Contact) GetPhoneNumberOk() (*PhoneNumber, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Contact) SetPhoneNumber(v PhoneNumber)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Contact) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetWebAddress

`func (o *Contact) GetWebAddress() string`

GetWebAddress returns the WebAddress field if non-nil, zero value otherwise.

### GetWebAddressOk

`func (o *Contact) GetWebAddressOk() (*string, bool)`

GetWebAddressOk returns a tuple with the WebAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAddress

`func (o *Contact) SetWebAddress(v string)`

SetWebAddress sets WebAddress field to given value.

### HasWebAddress

`func (o *Contact) HasWebAddress() bool`

HasWebAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


