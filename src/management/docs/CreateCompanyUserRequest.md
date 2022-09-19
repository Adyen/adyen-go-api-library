# CreateCompanyUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGroups** | Pointer to **[]string** | The list of [account groups](https://docs.adyen.com/account/account-structure#account-groups) associated with this user. | [optional] 
**AssociatedMerchantAccounts** | Pointer to **[]string** | The list of [merchant accounts](https://docs.adyen.com/account/account-structure#merchant-accounts) associated with this user. | [optional] 
**Email** | **string** | The email address of the user. | 
**Name** | [**Name**](Name.md) |  | 
**Roles** | Pointer to **[]string** | The list of [roles](https://docs.adyen.com/account/user-roles) for this user. | [optional] 
**TimeZoneCode** | Pointer to **string** | The [tz database name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) of the time zone of the user. For example, **Europe/Amsterdam**. | [optional] 
**Username** | **string** | The username for this user. Allowed length: 255 alphanumeric characters. | 

## Methods

### NewCreateCompanyUserRequest

`func NewCreateCompanyUserRequest(email string, name Name, username string, ) *CreateCompanyUserRequest`

NewCreateCompanyUserRequest instantiates a new CreateCompanyUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCompanyUserRequestWithDefaults

`func NewCreateCompanyUserRequestWithDefaults() *CreateCompanyUserRequest`

NewCreateCompanyUserRequestWithDefaults instantiates a new CreateCompanyUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGroups

`func (o *CreateCompanyUserRequest) GetAccountGroups() []string`

GetAccountGroups returns the AccountGroups field if non-nil, zero value otherwise.

### GetAccountGroupsOk

`func (o *CreateCompanyUserRequest) GetAccountGroupsOk() (*[]string, bool)`

GetAccountGroupsOk returns a tuple with the AccountGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroups

`func (o *CreateCompanyUserRequest) SetAccountGroups(v []string)`

SetAccountGroups sets AccountGroups field to given value.

### HasAccountGroups

`func (o *CreateCompanyUserRequest) HasAccountGroups() bool`

HasAccountGroups returns a boolean if a field has been set.

### GetAssociatedMerchantAccounts

`func (o *CreateCompanyUserRequest) GetAssociatedMerchantAccounts() []string`

GetAssociatedMerchantAccounts returns the AssociatedMerchantAccounts field if non-nil, zero value otherwise.

### GetAssociatedMerchantAccountsOk

`func (o *CreateCompanyUserRequest) GetAssociatedMerchantAccountsOk() (*[]string, bool)`

GetAssociatedMerchantAccountsOk returns a tuple with the AssociatedMerchantAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedMerchantAccounts

`func (o *CreateCompanyUserRequest) SetAssociatedMerchantAccounts(v []string)`

SetAssociatedMerchantAccounts sets AssociatedMerchantAccounts field to given value.

### HasAssociatedMerchantAccounts

`func (o *CreateCompanyUserRequest) HasAssociatedMerchantAccounts() bool`

HasAssociatedMerchantAccounts returns a boolean if a field has been set.

### GetEmail

`func (o *CreateCompanyUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateCompanyUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateCompanyUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *CreateCompanyUserRequest) GetName() Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCompanyUserRequest) GetNameOk() (*Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCompanyUserRequest) SetName(v Name)`

SetName sets Name field to given value.


### GetRoles

`func (o *CreateCompanyUserRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateCompanyUserRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateCompanyUserRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateCompanyUserRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTimeZoneCode

`func (o *CreateCompanyUserRequest) GetTimeZoneCode() string`

GetTimeZoneCode returns the TimeZoneCode field if non-nil, zero value otherwise.

### GetTimeZoneCodeOk

`func (o *CreateCompanyUserRequest) GetTimeZoneCodeOk() (*string, bool)`

GetTimeZoneCodeOk returns a tuple with the TimeZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneCode

`func (o *CreateCompanyUserRequest) SetTimeZoneCode(v string)`

SetTimeZoneCode sets TimeZoneCode field to given value.

### HasTimeZoneCode

`func (o *CreateCompanyUserRequest) HasTimeZoneCode() bool`

HasTimeZoneCode returns a boolean if a field has been set.

### GetUsername

`func (o *CreateCompanyUserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateCompanyUserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateCompanyUserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


