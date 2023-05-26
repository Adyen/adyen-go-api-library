# UltimatePartyIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address2**](Address2.md) |  | [optional] 
**DateOfBirth** | Pointer to **string** | The date of birth of the individual in [ISO-8601](https://www.w3.org/TR/NOTE-datetime) format. For example, **YYYY-MM-DD**.  Allowed only when &#x60;type&#x60; is **individual**. | [optional] 
**FirstName** | Pointer to **string** | First name of the individual.  Allowed only when &#x60;type&#x60; is **individual**. | [optional] 
**FullName** | **string** | The name of the entity. | 
**LastName** | Pointer to **string** | Last name of the individual.  Allowed only when &#x60;type&#x60; is **individual**. | [optional] 
**Reference** | Pointer to **string** | Your unique reference of the party. This should be consistent for all transfers initiated to/from the same party/counterparty. e.g Your client&#39;s unique wallet or payee ID | [optional] 
**Type** | Pointer to **string** | The type of entity that owns the bank account.   Possible values: **individual**, **organization**, or **unknown**. | [optional] [default to "unknown"]

## Methods

### NewUltimatePartyIdentification

`func NewUltimatePartyIdentification(fullName string, ) *UltimatePartyIdentification`

NewUltimatePartyIdentification instantiates a new UltimatePartyIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUltimatePartyIdentificationWithDefaults

`func NewUltimatePartyIdentificationWithDefaults() *UltimatePartyIdentification`

NewUltimatePartyIdentificationWithDefaults instantiates a new UltimatePartyIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UltimatePartyIdentification) GetAddress() Address2`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UltimatePartyIdentification) GetAddressOk() (*Address2, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UltimatePartyIdentification) SetAddress(v Address2)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UltimatePartyIdentification) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *UltimatePartyIdentification) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *UltimatePartyIdentification) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *UltimatePartyIdentification) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *UltimatePartyIdentification) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetFirstName

`func (o *UltimatePartyIdentification) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UltimatePartyIdentification) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UltimatePartyIdentification) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UltimatePartyIdentification) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetFullName

`func (o *UltimatePartyIdentification) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *UltimatePartyIdentification) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *UltimatePartyIdentification) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetLastName

`func (o *UltimatePartyIdentification) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UltimatePartyIdentification) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UltimatePartyIdentification) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UltimatePartyIdentification) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetReference

`func (o *UltimatePartyIdentification) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *UltimatePartyIdentification) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *UltimatePartyIdentification) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *UltimatePartyIdentification) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetType

`func (o *UltimatePartyIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UltimatePartyIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UltimatePartyIdentification) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UltimatePartyIdentification) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


