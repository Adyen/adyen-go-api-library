# DeliveryContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**Address2**](Address2.md) |  | 
**Email** | Pointer to **string** | The email address of the contact. | [optional] 
**FullPhoneNumber** | Pointer to **string** | The full phone number of the contact provided as a single string. It will be handled as a landline phone. **Examples:** \&quot;0031 6 11 22 33 44\&quot;, \&quot;+316/1122-3344\&quot;, \&quot;(0031) 611223344\&quot; | [optional] 
**Name** | [**Name**](Name.md) |  | 
**PhoneNumber** | Pointer to [**PhoneNumber**](PhoneNumber.md) |  | [optional] 
**WebAddress** | Pointer to **string** | The URL of the contact&#39;s website. | [optional] 

## Methods

### NewDeliveryContact

`func NewDeliveryContact(address Address2, name Name, ) *DeliveryContact`

NewDeliveryContact instantiates a new DeliveryContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryContactWithDefaults

`func NewDeliveryContactWithDefaults() *DeliveryContact`

NewDeliveryContactWithDefaults instantiates a new DeliveryContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DeliveryContact) GetAddress() Address2`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DeliveryContact) GetAddressOk() (*Address2, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DeliveryContact) SetAddress(v Address2)`

SetAddress sets Address field to given value.


### GetEmail

`func (o *DeliveryContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DeliveryContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DeliveryContact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DeliveryContact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFullPhoneNumber

`func (o *DeliveryContact) GetFullPhoneNumber() string`

GetFullPhoneNumber returns the FullPhoneNumber field if non-nil, zero value otherwise.

### GetFullPhoneNumberOk

`func (o *DeliveryContact) GetFullPhoneNumberOk() (*string, bool)`

GetFullPhoneNumberOk returns a tuple with the FullPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPhoneNumber

`func (o *DeliveryContact) SetFullPhoneNumber(v string)`

SetFullPhoneNumber sets FullPhoneNumber field to given value.

### HasFullPhoneNumber

`func (o *DeliveryContact) HasFullPhoneNumber() bool`

HasFullPhoneNumber returns a boolean if a field has been set.

### GetName

`func (o *DeliveryContact) GetName() Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeliveryContact) GetNameOk() (*Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeliveryContact) SetName(v Name)`

SetName sets Name field to given value.


### GetPhoneNumber

`func (o *DeliveryContact) GetPhoneNumber() PhoneNumber`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *DeliveryContact) GetPhoneNumberOk() (*PhoneNumber, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *DeliveryContact) SetPhoneNumber(v PhoneNumber)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *DeliveryContact) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetWebAddress

`func (o *DeliveryContact) GetWebAddress() string`

GetWebAddress returns the WebAddress field if non-nil, zero value otherwise.

### GetWebAddressOk

`func (o *DeliveryContact) GetWebAddressOk() (*string, bool)`

GetWebAddressOk returns a tuple with the WebAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAddress

`func (o *DeliveryContact) SetWebAddress(v string)`

SetWebAddress sets WebAddress field to given value.

### HasWebAddress

`func (o *DeliveryContact) HasWebAddress() bool`

HasWebAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


