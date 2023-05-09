# PaymentInstrumentUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceAccountId** | Pointer to **string** | The unique identifier of the balance account associated with this payment instrument. &gt;You can only change the balance account ID if the payment instrument has **inactive** status. | [optional] 
**Card** | Pointer to [**CardInfo**](CardInfo.md) |  | [optional] 
**Status** | Pointer to **string** | The status of the payment instrument. If a status is not specified when creating a payment instrument, it is set to **active** by default. However, there can be exceptions for cards based on the &#x60;card.formFactor&#x60; and the &#x60;issuingCountryCode&#x60;. For example, when issuing physical cards in the US, the default status is **inactive**.  Possible values:    * **active**:  The payment instrument is active and can be used to make payments.    * **inactive**: The payment instrument is inactive and cannot be used to make payments.    * **suspended**: The payment instrument is suspended, either because it was stolen or lost.    * **closed**: The payment instrument is permanently closed. This action cannot be undone.    | [optional] 
**StatusComment** | Pointer to **string** | Comment for the status of the payment instrument.  Required if &#x60;statusReason&#x60; is **other**. | [optional] 
**StatusReason** | Pointer to **string** | The reason for updating the status of the payment instrument.  Possible values: **lost**, **stolen**, **damaged**, **suspectedFraud**, **expired**, **endOfLife**, **accountClosure**, **other**. If the reason is **other**, you must also send the &#x60;statusComment&#x60; parameter describing the status change. | [optional] 

## Methods

### NewPaymentInstrumentUpdateRequest

`func NewPaymentInstrumentUpdateRequest() *PaymentInstrumentUpdateRequest`

NewPaymentInstrumentUpdateRequest instantiates a new PaymentInstrumentUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInstrumentUpdateRequestWithDefaults

`func NewPaymentInstrumentUpdateRequestWithDefaults() *PaymentInstrumentUpdateRequest`

NewPaymentInstrumentUpdateRequestWithDefaults instantiates a new PaymentInstrumentUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceAccountId

`func (o *PaymentInstrumentUpdateRequest) GetBalanceAccountId() string`

GetBalanceAccountId returns the BalanceAccountId field if non-nil, zero value otherwise.

### GetBalanceAccountIdOk

`func (o *PaymentInstrumentUpdateRequest) GetBalanceAccountIdOk() (*string, bool)`

GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccountId

`func (o *PaymentInstrumentUpdateRequest) SetBalanceAccountId(v string)`

SetBalanceAccountId sets BalanceAccountId field to given value.

### HasBalanceAccountId

`func (o *PaymentInstrumentUpdateRequest) HasBalanceAccountId() bool`

HasBalanceAccountId returns a boolean if a field has been set.

### GetCard

`func (o *PaymentInstrumentUpdateRequest) GetCard() CardInfo`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *PaymentInstrumentUpdateRequest) GetCardOk() (*CardInfo, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *PaymentInstrumentUpdateRequest) SetCard(v CardInfo)`

SetCard sets Card field to given value.

### HasCard

`func (o *PaymentInstrumentUpdateRequest) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentInstrumentUpdateRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentInstrumentUpdateRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentInstrumentUpdateRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentInstrumentUpdateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusComment

`func (o *PaymentInstrumentUpdateRequest) GetStatusComment() string`

GetStatusComment returns the StatusComment field if non-nil, zero value otherwise.

### GetStatusCommentOk

`func (o *PaymentInstrumentUpdateRequest) GetStatusCommentOk() (*string, bool)`

GetStatusCommentOk returns a tuple with the StatusComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusComment

`func (o *PaymentInstrumentUpdateRequest) SetStatusComment(v string)`

SetStatusComment sets StatusComment field to given value.

### HasStatusComment

`func (o *PaymentInstrumentUpdateRequest) HasStatusComment() bool`

HasStatusComment returns a boolean if a field has been set.

### GetStatusReason

`func (o *PaymentInstrumentUpdateRequest) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *PaymentInstrumentUpdateRequest) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *PaymentInstrumentUpdateRequest) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *PaymentInstrumentUpdateRequest) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


