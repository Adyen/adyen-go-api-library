# Authentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email address where the one-time password (OTP) is sent. | [optional] 
**Password** | Pointer to **string** | The password used for 3D Secure password-based authentication. The value must be between 1 to 30 characters and must only contain the following supported characters.  * Characters between **a-z**, **A-Z**, and **0-9**  * Special characters: **äöüßÄÖÜ+-*_/ç%()&#x3D;?!~#&#39;\&quot;,;:$&amp;àùòâôûáúó** | [optional] 
**Phone** | Pointer to [**Phone**](Phone.md) |  | [optional] 

## Methods

### NewAuthentication

`func NewAuthentication() *Authentication`

NewAuthentication instantiates a new Authentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationWithDefaults

`func NewAuthenticationWithDefaults() *Authentication`

NewAuthenticationWithDefaults instantiates a new Authentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *Authentication) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Authentication) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Authentication) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Authentication) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *Authentication) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Authentication) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Authentication) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Authentication) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPhone

`func (o *Authentication) GetPhone() Phone`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Authentication) GetPhoneOk() (*Phone, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Authentication) SetPhone(v Phone)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Authentication) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


