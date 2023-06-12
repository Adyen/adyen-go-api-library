# TransferNotificationTransferTracking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstimatedArrivalTime** | Pointer to **time.Time** | The estimated time the beneficiary should have access to the funds. | [optional] 
**Status** | Pointer to **string** | The tracking status of the transfer. | [optional] 

## Methods

### NewTransferNotificationTransferTracking

`func NewTransferNotificationTransferTracking() *TransferNotificationTransferTracking`

NewTransferNotificationTransferTracking instantiates a new TransferNotificationTransferTracking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferNotificationTransferTrackingWithDefaults

`func NewTransferNotificationTransferTrackingWithDefaults() *TransferNotificationTransferTracking`

NewTransferNotificationTransferTrackingWithDefaults instantiates a new TransferNotificationTransferTracking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimatedArrivalTime

`func (o *TransferNotificationTransferTracking) GetEstimatedArrivalTime() time.Time`

GetEstimatedArrivalTime returns the EstimatedArrivalTime field if non-nil, zero value otherwise.

### GetEstimatedArrivalTimeOk

`func (o *TransferNotificationTransferTracking) GetEstimatedArrivalTimeOk() (*time.Time, bool)`

GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedArrivalTime

`func (o *TransferNotificationTransferTracking) SetEstimatedArrivalTime(v time.Time)`

SetEstimatedArrivalTime sets EstimatedArrivalTime field to given value.

### HasEstimatedArrivalTime

`func (o *TransferNotificationTransferTracking) HasEstimatedArrivalTime() bool`

HasEstimatedArrivalTime returns a boolean if a field has been set.

### GetStatus

`func (o *TransferNotificationTransferTracking) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransferNotificationTransferTracking) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransferNotificationTransferTracking) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransferNotificationTransferTracking) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


