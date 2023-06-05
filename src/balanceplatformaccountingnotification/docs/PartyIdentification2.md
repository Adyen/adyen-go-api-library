# PartyIdentification2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address2**](Address2.md) |  | [optional] 
**FirstName** | Pointer to **string** | First name of the individual. Required when &#x60;type&#x60; is **individual**. | [optional] 
**FullName** | **string** | The name of the entity. | 
**LastName** | Pointer to **string** | Last name of the individual. Required when &#x60;type&#x60; is **individual**. | [optional] 
**Type** | Pointer to **string** | The type of entity that owns the bank account.   Possible values: **individual**, **organization**, or **unknown**. | [optional] [default to "unknown"]

## Methods

### NewPartyIdentification2

`func NewPartyIdentification2(fullName string, ) *PartyIdentification2`

NewPartyIdentification2 instantiates a new PartyIdentification2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyIdentification2WithDefaults

`func NewPartyIdentification2WithDefaults() *PartyIdentification2`

NewPartyIdentification2WithDefaults instantiates a new PartyIdentification2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PartyIdentification2) GetAddress() Address2`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PartyIdentification2) GetAddressOk() (*Address2, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PartyIdentification2) SetAddress(v Address2)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PartyIdentification2) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetFirstName

`func (o *PartyIdentification2) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PartyIdentification2) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PartyIdentification2) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PartyIdentification2) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetFullName

`func (o *PartyIdentification2) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *PartyIdentification2) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *PartyIdentification2) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetLastName

`func (o *PartyIdentification2) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PartyIdentification2) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PartyIdentification2) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PartyIdentification2) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetType

`func (o *PartyIdentification2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartyIdentification2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartyIdentification2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PartyIdentification2) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


