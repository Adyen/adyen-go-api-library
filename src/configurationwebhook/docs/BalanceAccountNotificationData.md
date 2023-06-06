# BalanceAccountNotificationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceAccount** | Pointer to [**BalanceAccount**](BalanceAccount.md) |  | [optional] 
**BalancePlatform** | Pointer to **string** | The unique identifier of the balance platform. | [optional] 

## Methods

### NewBalanceAccountNotificationData

`func NewBalanceAccountNotificationData() *BalanceAccountNotificationData`

NewBalanceAccountNotificationData instantiates a new BalanceAccountNotificationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceAccountNotificationDataWithDefaults

`func NewBalanceAccountNotificationDataWithDefaults() *BalanceAccountNotificationData`

NewBalanceAccountNotificationDataWithDefaults instantiates a new BalanceAccountNotificationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceAccount

`func (o *BalanceAccountNotificationData) GetBalanceAccount() BalanceAccount`

GetBalanceAccount returns the BalanceAccount field if non-nil, zero value otherwise.

### GetBalanceAccountOk

`func (o *BalanceAccountNotificationData) GetBalanceAccountOk() (*BalanceAccount, bool)`

GetBalanceAccountOk returns a tuple with the BalanceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccount

`func (o *BalanceAccountNotificationData) SetBalanceAccount(v BalanceAccount)`

SetBalanceAccount sets BalanceAccount field to given value.

### HasBalanceAccount

`func (o *BalanceAccountNotificationData) HasBalanceAccount() bool`

HasBalanceAccount returns a boolean if a field has been set.

### GetBalancePlatform

`func (o *BalanceAccountNotificationData) GetBalancePlatform() string`

GetBalancePlatform returns the BalancePlatform field if non-nil, zero value otherwise.

### GetBalancePlatformOk

`func (o *BalanceAccountNotificationData) GetBalancePlatformOk() (*string, bool)`

GetBalancePlatformOk returns a tuple with the BalancePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancePlatform

`func (o *BalanceAccountNotificationData) SetBalancePlatform(v string)`

SetBalancePlatform sets BalancePlatform field to given value.

### HasBalancePlatform

`func (o *BalanceAccountNotificationData) HasBalancePlatform() bool`

HasBalancePlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


