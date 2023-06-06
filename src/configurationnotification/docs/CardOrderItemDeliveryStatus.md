# CardOrderItemDeliveryStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status of the delivery. | [optional] 
**StatusError** | Pointer to **string** | Error status, if any. | [optional] 
**StatusErrorMessage** | Pointer to **string** | Error message, if any. | [optional] 
**TrackingNumber** | Pointer to **string** | Tracking number of the delivery. | [optional] 

## Methods

### NewCardOrderItemDeliveryStatus

`func NewCardOrderItemDeliveryStatus() *CardOrderItemDeliveryStatus`

NewCardOrderItemDeliveryStatus instantiates a new CardOrderItemDeliveryStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardOrderItemDeliveryStatusWithDefaults

`func NewCardOrderItemDeliveryStatusWithDefaults() *CardOrderItemDeliveryStatus`

NewCardOrderItemDeliveryStatusWithDefaults instantiates a new CardOrderItemDeliveryStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CardOrderItemDeliveryStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CardOrderItemDeliveryStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CardOrderItemDeliveryStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CardOrderItemDeliveryStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusError

`func (o *CardOrderItemDeliveryStatus) GetStatusError() string`

GetStatusError returns the StatusError field if non-nil, zero value otherwise.

### GetStatusErrorOk

`func (o *CardOrderItemDeliveryStatus) GetStatusErrorOk() (*string, bool)`

GetStatusErrorOk returns a tuple with the StatusError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusError

`func (o *CardOrderItemDeliveryStatus) SetStatusError(v string)`

SetStatusError sets StatusError field to given value.

### HasStatusError

`func (o *CardOrderItemDeliveryStatus) HasStatusError() bool`

HasStatusError returns a boolean if a field has been set.

### GetStatusErrorMessage

`func (o *CardOrderItemDeliveryStatus) GetStatusErrorMessage() string`

GetStatusErrorMessage returns the StatusErrorMessage field if non-nil, zero value otherwise.

### GetStatusErrorMessageOk

`func (o *CardOrderItemDeliveryStatus) GetStatusErrorMessageOk() (*string, bool)`

GetStatusErrorMessageOk returns a tuple with the StatusErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusErrorMessage

`func (o *CardOrderItemDeliveryStatus) SetStatusErrorMessage(v string)`

SetStatusErrorMessage sets StatusErrorMessage field to given value.

### HasStatusErrorMessage

`func (o *CardOrderItemDeliveryStatus) HasStatusErrorMessage() bool`

HasStatusErrorMessage returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *CardOrderItemDeliveryStatus) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *CardOrderItemDeliveryStatus) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *CardOrderItemDeliveryStatus) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *CardOrderItemDeliveryStatus) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


