# NotificationModificationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**Type** | Pointer to **string** | The type of modification.  Possible values: **Authorised**, **Cancelled**, **Captured**, **Error**, **Expired**, **OutgoingTransfer**, **PendingIncomingTransfer**, **PendingRefund**, **IncomingTransfer**, **Refunded**, **Refused**, **AuthAdjustmentAuthorised**. | [optional] 

## Methods

### NewNotificationModificationData

`func NewNotificationModificationData() *NotificationModificationData`

NewNotificationModificationData instantiates a new NotificationModificationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationModificationDataWithDefaults

`func NewNotificationModificationDataWithDefaults() *NotificationModificationData`

NewNotificationModificationDataWithDefaults instantiates a new NotificationModificationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *NotificationModificationData) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *NotificationModificationData) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *NotificationModificationData) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *NotificationModificationData) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetType

`func (o *NotificationModificationData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationModificationData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationModificationData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NotificationModificationData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


