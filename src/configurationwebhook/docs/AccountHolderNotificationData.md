# AccountHolderNotificationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolder** | Pointer to [**AccountHolder**](AccountHolder.md) |  | [optional] 
**BalancePlatform** | Pointer to **string** | The unique identifier of the balance platform. | [optional] 

## Methods

### NewAccountHolderNotificationData

`func NewAccountHolderNotificationData() *AccountHolderNotificationData`

NewAccountHolderNotificationData instantiates a new AccountHolderNotificationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountHolderNotificationDataWithDefaults

`func NewAccountHolderNotificationDataWithDefaults() *AccountHolderNotificationData`

NewAccountHolderNotificationDataWithDefaults instantiates a new AccountHolderNotificationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolder

`func (o *AccountHolderNotificationData) GetAccountHolder() AccountHolder`

GetAccountHolder returns the AccountHolder field if non-nil, zero value otherwise.

### GetAccountHolderOk

`func (o *AccountHolderNotificationData) GetAccountHolderOk() (*AccountHolder, bool)`

GetAccountHolderOk returns a tuple with the AccountHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolder

`func (o *AccountHolderNotificationData) SetAccountHolder(v AccountHolder)`

SetAccountHolder sets AccountHolder field to given value.

### HasAccountHolder

`func (o *AccountHolderNotificationData) HasAccountHolder() bool`

HasAccountHolder returns a boolean if a field has been set.

### GetBalancePlatform

`func (o *AccountHolderNotificationData) GetBalancePlatform() string`

GetBalancePlatform returns the BalancePlatform field if non-nil, zero value otherwise.

### GetBalancePlatformOk

`func (o *AccountHolderNotificationData) GetBalancePlatformOk() (*string, bool)`

GetBalancePlatformOk returns a tuple with the BalancePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancePlatform

`func (o *AccountHolderNotificationData) SetBalancePlatform(v string)`

SetBalancePlatform sets BalancePlatform field to given value.

### HasBalancePlatform

`func (o *AccountHolderNotificationData) HasBalancePlatform() bool`

HasBalancePlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


