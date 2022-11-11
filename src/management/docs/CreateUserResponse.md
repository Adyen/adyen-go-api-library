# CreateUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**AccountGroups** | Pointer to **[]string** | The list of [account groups](https://docs.adyen.com/account/account-structure#account-groups) associated with this user. | [optional] 
**Active** | Pointer to **bool** | Indicates whether this user is active. | [optional] 
**Email** | **string** | The email address of the user. | 
**Id** | **string** | The unique identifier of the user. | 
**Name** | Pointer to [**Name**](Name.md) |  | [optional] 
**Roles** | **[]string** | The list of [roles](https://docs.adyen.com/account/user-roles) for this user. | 
**TimeZoneCode** | **string** | The [tz database name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) of the time zone of the user. For example, **Europe/Amsterdam**. | 
**Username** | **string** | The username for this user. | 

## Methods

### NewCreateUserResponse

`func NewCreateUserResponse(email string, id string, roles []string, timeZoneCode string, username string, ) *CreateUserResponse`

NewCreateUserResponse instantiates a new CreateUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserResponseWithDefaults

`func NewCreateUserResponseWithDefaults() *CreateUserResponse`

NewCreateUserResponseWithDefaults instantiates a new CreateUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *CreateUserResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateUserResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateUserResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateUserResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccountGroups

`func (o *CreateUserResponse) GetAccountGroups() []string`

GetAccountGroups returns the AccountGroups field if non-nil, zero value otherwise.

### GetAccountGroupsOk

`func (o *CreateUserResponse) GetAccountGroupsOk() (*[]string, bool)`

GetAccountGroupsOk returns a tuple with the AccountGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroups

`func (o *CreateUserResponse) SetAccountGroups(v []string)`

SetAccountGroups sets AccountGroups field to given value.

### HasAccountGroups

`func (o *CreateUserResponse) HasAccountGroups() bool`

HasAccountGroups returns a boolean if a field has been set.

### GetActive

`func (o *CreateUserResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateUserResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateUserResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateUserResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEmail

`func (o *CreateUserResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetId

`func (o *CreateUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateUserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateUserResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CreateUserResponse) GetName() Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUserResponse) GetNameOk() (*Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUserResponse) SetName(v Name)`

SetName sets Name field to given value.

### HasName

`func (o *CreateUserResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoles

`func (o *CreateUserResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateUserResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateUserResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetTimeZoneCode

`func (o *CreateUserResponse) GetTimeZoneCode() string`

GetTimeZoneCode returns the TimeZoneCode field if non-nil, zero value otherwise.

### GetTimeZoneCodeOk

`func (o *CreateUserResponse) GetTimeZoneCodeOk() (*string, bool)`

GetTimeZoneCodeOk returns a tuple with the TimeZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneCode

`func (o *CreateUserResponse) SetTimeZoneCode(v string)`

SetTimeZoneCode sets TimeZoneCode field to given value.


### GetUsername

`func (o *CreateUserResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateUserResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateUserResponse) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


