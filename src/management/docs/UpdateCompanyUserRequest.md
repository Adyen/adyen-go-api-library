# UpdateCompanyUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGroups** | Pointer to **[]string** | The list of [account groups](https://docs.adyen.com/account/account-structure#account-groups) associated with this user. | [optional] 
**Active** | Pointer to **bool** | Indicates whether this user is active. | [optional] 
**AssociatedMerchantAccounts** | Pointer to **[]string** | The list of [merchant accounts](https://docs.adyen.com/account/account-structure#merchant-accounts) to associate the user with. | [optional] 
**Email** | Pointer to **string** | The email address of the user. | [optional] 
**Name** | Pointer to [**Name2**](Name2.md) |  | [optional] 
**Roles** | Pointer to **[]string** | The list of [roles](https://docs.adyen.com/account/user-roles) for this user. | [optional] 
**TimeZoneCode** | Pointer to **string** | The [tz database name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) of the time zone of the user. For example, **Europe/Amsterdam**. | [optional] 

## Methods

### NewUpdateCompanyUserRequest

`func NewUpdateCompanyUserRequest() *UpdateCompanyUserRequest`

NewUpdateCompanyUserRequest instantiates a new UpdateCompanyUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCompanyUserRequestWithDefaults

`func NewUpdateCompanyUserRequestWithDefaults() *UpdateCompanyUserRequest`

NewUpdateCompanyUserRequestWithDefaults instantiates a new UpdateCompanyUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGroups

`func (o *UpdateCompanyUserRequest) GetAccountGroups() []string`

GetAccountGroups returns the AccountGroups field if non-nil, zero value otherwise.

### GetAccountGroupsOk

`func (o *UpdateCompanyUserRequest) GetAccountGroupsOk() (*[]string, bool)`

GetAccountGroupsOk returns a tuple with the AccountGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroups

`func (o *UpdateCompanyUserRequest) SetAccountGroups(v []string)`

SetAccountGroups sets AccountGroups field to given value.

### HasAccountGroups

`func (o *UpdateCompanyUserRequest) HasAccountGroups() bool`

HasAccountGroups returns a boolean if a field has been set.

### GetActive

`func (o *UpdateCompanyUserRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateCompanyUserRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateCompanyUserRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateCompanyUserRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAssociatedMerchantAccounts

`func (o *UpdateCompanyUserRequest) GetAssociatedMerchantAccounts() []string`

GetAssociatedMerchantAccounts returns the AssociatedMerchantAccounts field if non-nil, zero value otherwise.

### GetAssociatedMerchantAccountsOk

`func (o *UpdateCompanyUserRequest) GetAssociatedMerchantAccountsOk() (*[]string, bool)`

GetAssociatedMerchantAccountsOk returns a tuple with the AssociatedMerchantAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedMerchantAccounts

`func (o *UpdateCompanyUserRequest) SetAssociatedMerchantAccounts(v []string)`

SetAssociatedMerchantAccounts sets AssociatedMerchantAccounts field to given value.

### HasAssociatedMerchantAccounts

`func (o *UpdateCompanyUserRequest) HasAssociatedMerchantAccounts() bool`

HasAssociatedMerchantAccounts returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateCompanyUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateCompanyUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateCompanyUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateCompanyUserRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *UpdateCompanyUserRequest) GetName() Name2`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCompanyUserRequest) GetNameOk() (*Name2, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCompanyUserRequest) SetName(v Name2)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCompanyUserRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoles

`func (o *UpdateCompanyUserRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateCompanyUserRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateCompanyUserRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateCompanyUserRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTimeZoneCode

`func (o *UpdateCompanyUserRequest) GetTimeZoneCode() string`

GetTimeZoneCode returns the TimeZoneCode field if non-nil, zero value otherwise.

### GetTimeZoneCodeOk

`func (o *UpdateCompanyUserRequest) GetTimeZoneCodeOk() (*string, bool)`

GetTimeZoneCodeOk returns a tuple with the TimeZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneCode

`func (o *UpdateCompanyUserRequest) SetTimeZoneCode(v string)`

SetTimeZoneCode sets TimeZoneCode field to given value.

### HasTimeZoneCode

`func (o *UpdateCompanyUserRequest) HasTimeZoneCode() bool`

HasTimeZoneCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


