# TransferEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**AmountAdjustments** | Pointer to [**[]AmountAdjustment**](AmountAdjustment.md) | The amount adjustments in this transfer. | [optional] 
**BookingDate** | Pointer to **time.Time** | The date when the transfer request was sent. | [optional] 
**EstimatedArrivalTime** | Pointer to **time.Time** | The estimated time the beneficiary should have access to the funds. | [optional] 
**Id** | Pointer to **string** | The unique identifier of the transfer event. | [optional] 
**Mutations** | Pointer to [**[]BalanceMutation**](BalanceMutation.md) | The list of the balance mutation per event. | [optional] 
**OriginalAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**Reason** | Pointer to **string** | The reason for the transfer status. | [optional] 
**Status** | Pointer to **string** | The status of the transfer event. | [optional] 
**TransactionId** | Pointer to **string** | The id of the transaction that is related to this accounting event. Only sent for events of type **accounting** where the balance changes. | [optional] 
**Type** | Pointer to **string** | The type of the transfer event. Possible values: **accounting**, **tracking**. | [optional] 
**UpdateDate** | Pointer to **time.Time** | The date when the tracking status was updated. | [optional] 
**ValueDate** | Pointer to **time.Time** | A future date, when the funds are expected to be deducted from or credited to the balance account. | [optional] 

## Methods

### NewTransferEvent

`func NewTransferEvent() *TransferEvent`

NewTransferEvent instantiates a new TransferEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferEventWithDefaults

`func NewTransferEventWithDefaults() *TransferEvent`

NewTransferEventWithDefaults instantiates a new TransferEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransferEvent) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferEvent) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferEvent) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransferEvent) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountAdjustments

`func (o *TransferEvent) GetAmountAdjustments() []AmountAdjustment`

GetAmountAdjustments returns the AmountAdjustments field if non-nil, zero value otherwise.

### GetAmountAdjustmentsOk

`func (o *TransferEvent) GetAmountAdjustmentsOk() (*[]AmountAdjustment, bool)`

GetAmountAdjustmentsOk returns a tuple with the AmountAdjustments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAdjustments

`func (o *TransferEvent) SetAmountAdjustments(v []AmountAdjustment)`

SetAmountAdjustments sets AmountAdjustments field to given value.

### HasAmountAdjustments

`func (o *TransferEvent) HasAmountAdjustments() bool`

HasAmountAdjustments returns a boolean if a field has been set.

### GetBookingDate

`func (o *TransferEvent) GetBookingDate() time.Time`

GetBookingDate returns the BookingDate field if non-nil, zero value otherwise.

### GetBookingDateOk

`func (o *TransferEvent) GetBookingDateOk() (*time.Time, bool)`

GetBookingDateOk returns a tuple with the BookingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingDate

`func (o *TransferEvent) SetBookingDate(v time.Time)`

SetBookingDate sets BookingDate field to given value.

### HasBookingDate

`func (o *TransferEvent) HasBookingDate() bool`

HasBookingDate returns a boolean if a field has been set.

### GetEstimatedArrivalTime

`func (o *TransferEvent) GetEstimatedArrivalTime() time.Time`

GetEstimatedArrivalTime returns the EstimatedArrivalTime field if non-nil, zero value otherwise.

### GetEstimatedArrivalTimeOk

`func (o *TransferEvent) GetEstimatedArrivalTimeOk() (*time.Time, bool)`

GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedArrivalTime

`func (o *TransferEvent) SetEstimatedArrivalTime(v time.Time)`

SetEstimatedArrivalTime sets EstimatedArrivalTime field to given value.

### HasEstimatedArrivalTime

`func (o *TransferEvent) HasEstimatedArrivalTime() bool`

HasEstimatedArrivalTime returns a boolean if a field has been set.

### GetId

`func (o *TransferEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransferEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMutations

`func (o *TransferEvent) GetMutations() []BalanceMutation`

GetMutations returns the Mutations field if non-nil, zero value otherwise.

### GetMutationsOk

`func (o *TransferEvent) GetMutationsOk() (*[]BalanceMutation, bool)`

GetMutationsOk returns a tuple with the Mutations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutations

`func (o *TransferEvent) SetMutations(v []BalanceMutation)`

SetMutations sets Mutations field to given value.

### HasMutations

`func (o *TransferEvent) HasMutations() bool`

HasMutations returns a boolean if a field has been set.

### GetOriginalAmount

`func (o *TransferEvent) GetOriginalAmount() Amount`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *TransferEvent) GetOriginalAmountOk() (*Amount, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *TransferEvent) SetOriginalAmount(v Amount)`

SetOriginalAmount sets OriginalAmount field to given value.

### HasOriginalAmount

`func (o *TransferEvent) HasOriginalAmount() bool`

HasOriginalAmount returns a boolean if a field has been set.

### GetReason

`func (o *TransferEvent) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TransferEvent) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TransferEvent) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TransferEvent) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *TransferEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransferEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransferEvent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransferEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransactionId

`func (o *TransferEvent) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TransferEvent) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TransferEvent) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *TransferEvent) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetType

`func (o *TransferEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransferEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransferEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransferEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateDate

`func (o *TransferEvent) GetUpdateDate() time.Time`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *TransferEvent) GetUpdateDateOk() (*time.Time, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *TransferEvent) SetUpdateDate(v time.Time)`

SetUpdateDate sets UpdateDate field to given value.

### HasUpdateDate

`func (o *TransferEvent) HasUpdateDate() bool`

HasUpdateDate returns a boolean if a field has been set.

### GetValueDate

`func (o *TransferEvent) GetValueDate() time.Time`

GetValueDate returns the ValueDate field if non-nil, zero value otherwise.

### GetValueDateOk

`func (o *TransferEvent) GetValueDateOk() (*time.Time, bool)`

GetValueDateOk returns a tuple with the ValueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDate

`func (o *TransferEvent) SetValueDate(v time.Time)`

SetValueDate sets ValueDate field to given value.

### HasValueDate

`func (o *TransferEvent) HasValueDate() bool`

HasValueDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


