# ContactDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**Address**](Address.md) |  | 
**Email** | **string** | The email address of the account holder. | 
**Phone** | [**Phone**](Phone.md) |  | 
**WebAddress** | Pointer to **string** | The URL of the account holder&#39;s website. | [optional] 

## Methods

### NewContactDetails

`func NewContactDetails(address Address, email string, phone Phone, ) *ContactDetails`

NewContactDetails instantiates a new ContactDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactDetailsWithDefaults

`func NewContactDetailsWithDefaults() *ContactDetails`

NewContactDetailsWithDefaults instantiates a new ContactDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ContactDetails) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContactDetails) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContactDetails) SetAddress(v Address)`

SetAddress sets Address field to given value.


### GetEmail

`func (o *ContactDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactDetails) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhone

`func (o *ContactDetails) GetPhone() Phone`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ContactDetails) GetPhoneOk() (*Phone, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ContactDetails) SetPhone(v Phone)`

SetPhone sets Phone field to given value.


### GetWebAddress

`func (o *ContactDetails) GetWebAddress() string`

GetWebAddress returns the WebAddress field if non-nil, zero value otherwise.

### GetWebAddressOk

`func (o *ContactDetails) GetWebAddressOk() (*string, bool)`

GetWebAddressOk returns a tuple with the WebAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAddress

`func (o *ContactDetails) SetWebAddress(v string)`

SetWebAddress sets WebAddress field to given value.

### HasWebAddress

`func (o *ContactDetails) HasWebAddress() bool`

HasWebAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


