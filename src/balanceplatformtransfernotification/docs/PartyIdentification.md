# PartyIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**DateOfBirth** | Pointer to **string** | The date of birth of the individual in [ISO-8601](https://www.w3.org/TR/NOTE-datetime) format. For example, **YYYY-MM-DD**.  Allowed only when &#x60;type&#x60; is **individual**. | [optional] 
**FirstName** | Pointer to **string** | First name of the individual.  Allowed only when &#x60;type&#x60; is **individual**. | [optional] 
**FullName** | **string** | The name of the entity. | 
**LastName** | Pointer to **string** | Last name of the individual.  Allowed only when &#x60;type&#x60; is **individual**. | [optional] 
**Reference** | Pointer to **string** | Your unique reference of the party. This should be consistent for all transfers initiated to/from the same party/counterparty. e.g Your client&#39;s unique wallet or payee ID | [optional] 
**Type** | Pointer to **string** | The type of entity that owns the bank account.   Possible values: **individual**, **organization**, or **unknown**. | [optional] [default to "unknown"]

## Methods

### NewPartyIdentification

`func NewPartyIdentification(fullName string, ) *PartyIdentification`

NewPartyIdentification instantiates a new PartyIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyIdentificationWithDefaults

`func NewPartyIdentificationWithDefaults() *PartyIdentification`

NewPartyIdentificationWithDefaults instantiates a new PartyIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PartyIdentification) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PartyIdentification) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PartyIdentification) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PartyIdentification) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *PartyIdentification) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *PartyIdentification) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *PartyIdentification) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *PartyIdentification) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetFirstName

`func (o *PartyIdentification) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PartyIdentification) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PartyIdentification) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PartyIdentification) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetFullName

`func (o *PartyIdentification) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *PartyIdentification) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *PartyIdentification) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetLastName

`func (o *PartyIdentification) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PartyIdentification) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PartyIdentification) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PartyIdentification) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetReference

`func (o *PartyIdentification) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PartyIdentification) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PartyIdentification) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PartyIdentification) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetType

`func (o *PartyIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartyIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartyIdentification) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PartyIdentification) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


