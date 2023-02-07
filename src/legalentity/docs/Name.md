# Name

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | The individual&#39;s first name. | 
**Infix** | Pointer to **string** | The infix in the individual&#39;s name, if any. | [optional] 
**LastName** | **string** | The individual&#39;s last name. | 

## Methods

### NewName

`func NewName(firstName string, lastName string, ) *Name`

NewName instantiates a new Name object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNameWithDefaults

`func NewNameWithDefaults() *Name`

NewNameWithDefaults instantiates a new Name object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *Name) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Name) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Name) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetInfix

`func (o *Name) GetInfix() string`

GetInfix returns the Infix field if non-nil, zero value otherwise.

### GetInfixOk

`func (o *Name) GetInfixOk() (*string, bool)`

GetInfixOk returns a tuple with the Infix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfix

`func (o *Name) SetInfix(v string)`

SetInfix sets Infix field to given value.

### HasInfix

`func (o *Name) HasInfix() bool`

HasInfix returns a boolean if a field has been set.

### GetLastName

`func (o *Name) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Name) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Name) SetLastName(v string)`

SetLastName sets LastName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


