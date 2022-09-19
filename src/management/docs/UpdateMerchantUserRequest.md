# UpdateMerchantUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGroups** | Pointer to **[]string** | The list of [account groups](https://docs.adyen.com/account/account-structure#account-groups) associated with this user. | [optional] 
**Active** | Pointer to **bool** | Sets the status of the user to active (**true**) or inactive (**false**). | [optional] 
**Email** | Pointer to **string** | The email address of the user. | [optional] 
**Name** | Pointer to [**Name2**](Name2.md) |  | [optional] 
**Roles** | Pointer to **[]string** | The list of [roles](https://docs.adyen.com/account/user-roles) for this user. | [optional] 
**TimeZoneCode** | Pointer to **string** | The [tz database name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) of the time zone of the user. For example, **Europe/Amsterdam**. | [optional] 

## Methods

### NewUpdateMerchantUserRequest

`func NewUpdateMerchantUserRequest() *UpdateMerchantUserRequest`

NewUpdateMerchantUserRequest instantiates a new UpdateMerchantUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMerchantUserRequestWithDefaults

`func NewUpdateMerchantUserRequestWithDefaults() *UpdateMerchantUserRequest`

NewUpdateMerchantUserRequestWithDefaults instantiates a new UpdateMerchantUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGroups

`func (o *UpdateMerchantUserRequest) GetAccountGroups() []string`

GetAccountGroups returns the AccountGroups field if non-nil, zero value otherwise.

### GetAccountGroupsOk

`func (o *UpdateMerchantUserRequest) GetAccountGroupsOk() (*[]string, bool)`

GetAccountGroupsOk returns a tuple with the AccountGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroups

`func (o *UpdateMerchantUserRequest) SetAccountGroups(v []string)`

SetAccountGroups sets AccountGroups field to given value.

### HasAccountGroups

`func (o *UpdateMerchantUserRequest) HasAccountGroups() bool`

HasAccountGroups returns a boolean if a field has been set.

### GetActive

`func (o *UpdateMerchantUserRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateMerchantUserRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateMerchantUserRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateMerchantUserRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateMerchantUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateMerchantUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateMerchantUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateMerchantUserRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *UpdateMerchantUserRequest) GetName() Name2`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMerchantUserRequest) GetNameOk() (*Name2, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMerchantUserRequest) SetName(v Name2)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMerchantUserRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoles

`func (o *UpdateMerchantUserRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateMerchantUserRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateMerchantUserRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateMerchantUserRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTimeZoneCode

`func (o *UpdateMerchantUserRequest) GetTimeZoneCode() string`

GetTimeZoneCode returns the TimeZoneCode field if non-nil, zero value otherwise.

### GetTimeZoneCodeOk

`func (o *UpdateMerchantUserRequest) GetTimeZoneCodeOk() (*string, bool)`

GetTimeZoneCodeOk returns a tuple with the TimeZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneCode

`func (o *UpdateMerchantUserRequest) SetTimeZoneCode(v string)`

SetTimeZoneCode sets TimeZoneCode field to given value.

### HasTimeZoneCode

`func (o *UpdateMerchantUserRequest) HasTimeZoneCode() bool`

HasTimeZoneCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


