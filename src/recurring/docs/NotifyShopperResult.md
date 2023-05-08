# NotifyShopperResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayedReference** | Pointer to **string** | Reference of Pre-debit notification that is displayed to the shopper | [optional] 
**Message** | Pointer to **string** | A simple description of the &#x60;resultCode&#x60;. | [optional] 
**PspReference** | Pointer to **string** | The unique reference that is associated with the request. | [optional] 
**Reference** | Pointer to **string** | Reference of Pre-debit notification sent in my the merchant | [optional] 
**ResultCode** | Pointer to **string** | The code indicating the status of notification. | [optional] 
**ShopperNotificationReference** | Pointer to **string** | The unique reference for the request sent downstream. | [optional] 
**StoredPaymentMethodId** | Pointer to **string** | This is the recurringDetailReference returned in the response when token was created | [optional] 

## Methods

### NewNotifyShopperResult

`func NewNotifyShopperResult() *NotifyShopperResult`

NewNotifyShopperResult instantiates a new NotifyShopperResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyShopperResultWithDefaults

`func NewNotifyShopperResultWithDefaults() *NotifyShopperResult`

NewNotifyShopperResultWithDefaults instantiates a new NotifyShopperResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayedReference

`func (o *NotifyShopperResult) GetDisplayedReference() string`

GetDisplayedReference returns the DisplayedReference field if non-nil, zero value otherwise.

### GetDisplayedReferenceOk

`func (o *NotifyShopperResult) GetDisplayedReferenceOk() (*string, bool)`

GetDisplayedReferenceOk returns a tuple with the DisplayedReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayedReference

`func (o *NotifyShopperResult) SetDisplayedReference(v string)`

SetDisplayedReference sets DisplayedReference field to given value.

### HasDisplayedReference

`func (o *NotifyShopperResult) HasDisplayedReference() bool`

HasDisplayedReference returns a boolean if a field has been set.

### GetMessage

`func (o *NotifyShopperResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotifyShopperResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotifyShopperResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NotifyShopperResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPspReference

`func (o *NotifyShopperResult) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *NotifyShopperResult) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *NotifyShopperResult) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.

### HasPspReference

`func (o *NotifyShopperResult) HasPspReference() bool`

HasPspReference returns a boolean if a field has been set.

### GetReference

`func (o *NotifyShopperResult) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *NotifyShopperResult) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *NotifyShopperResult) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *NotifyShopperResult) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetResultCode

`func (o *NotifyShopperResult) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *NotifyShopperResult) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *NotifyShopperResult) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *NotifyShopperResult) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetShopperNotificationReference

`func (o *NotifyShopperResult) GetShopperNotificationReference() string`

GetShopperNotificationReference returns the ShopperNotificationReference field if non-nil, zero value otherwise.

### GetShopperNotificationReferenceOk

`func (o *NotifyShopperResult) GetShopperNotificationReferenceOk() (*string, bool)`

GetShopperNotificationReferenceOk returns a tuple with the ShopperNotificationReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperNotificationReference

`func (o *NotifyShopperResult) SetShopperNotificationReference(v string)`

SetShopperNotificationReference sets ShopperNotificationReference field to given value.

### HasShopperNotificationReference

`func (o *NotifyShopperResult) HasShopperNotificationReference() bool`

HasShopperNotificationReference returns a boolean if a field has been set.

### GetStoredPaymentMethodId

`func (o *NotifyShopperResult) GetStoredPaymentMethodId() string`

GetStoredPaymentMethodId returns the StoredPaymentMethodId field if non-nil, zero value otherwise.

### GetStoredPaymentMethodIdOk

`func (o *NotifyShopperResult) GetStoredPaymentMethodIdOk() (*string, bool)`

GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredPaymentMethodId

`func (o *NotifyShopperResult) SetStoredPaymentMethodId(v string)`

SetStoredPaymentMethodId sets StoredPaymentMethodId field to given value.

### HasStoredPaymentMethodId

`func (o *NotifyShopperResult) HasStoredPaymentMethodId() bool`

HasStoredPaymentMethodId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


