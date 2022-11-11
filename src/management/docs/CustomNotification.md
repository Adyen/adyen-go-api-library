# CustomNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to [**Amount2**](Amount2.md) |  | [optional] 
**EventCode** | Pointer to **string** | The event that caused the notification to be sent.Currently supported values: * **AUTHORISATION** * **CANCELLATION** * **REFUND** * **CAPTURE** * **DEACTIVATE_RECURRING** * **REPORT_AVAILABLE** * **CHARGEBACK** * **REQUEST_FOR_INFORMATION** * **NOTIFICATION_OF_CHARGEBACK** * **NOTIFICATIONTEST** * **ORDER_OPENED** * **ORDER_CLOSED** * **CHARGEBACK_REVERSED** * **REFUNDED_REVERSED** * **REFUND_WITH_DATA** | [optional] 
**EventDate** | Pointer to **time.Time** | The time of the event. Format: [ISO 8601](http://www.w3.org/TR/NOTE-datetime), YYYY-MM-DDThh:mm:ssTZD. | [optional] 
**MerchantReference** | Pointer to **string** | Your reference for the custom test notification. | [optional] 
**PaymentMethod** | Pointer to **string** | The payment method for the payment that the notification is about. Possible values: * **amex** * **visa** * **mc** * **maestro** * **bcmc** * **paypal**  * **sms**  * **bankTransfer_NL** * **bankTransfer_DE** * **bankTransfer_BE** * **ideal** * **elv** * **sepadirectdebit**  | [optional] 
**Reason** | Pointer to **string** | A descripton of what caused the notification. | [optional] 
**Success** | Pointer to **bool** | The outcome of the event which the notification is about. Set to either **true** or **false**.  | [optional] 

## Methods

### NewCustomNotification

`func NewCustomNotification() *CustomNotification`

NewCustomNotification instantiates a new CustomNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomNotificationWithDefaults

`func NewCustomNotificationWithDefaults() *CustomNotification`

NewCustomNotificationWithDefaults instantiates a new CustomNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CustomNotification) GetAmount() Amount2`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CustomNotification) GetAmountOk() (*Amount2, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CustomNotification) SetAmount(v Amount2)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CustomNotification) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetEventCode

`func (o *CustomNotification) GetEventCode() string`

GetEventCode returns the EventCode field if non-nil, zero value otherwise.

### GetEventCodeOk

`func (o *CustomNotification) GetEventCodeOk() (*string, bool)`

GetEventCodeOk returns a tuple with the EventCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCode

`func (o *CustomNotification) SetEventCode(v string)`

SetEventCode sets EventCode field to given value.

### HasEventCode

`func (o *CustomNotification) HasEventCode() bool`

HasEventCode returns a boolean if a field has been set.

### GetEventDate

`func (o *CustomNotification) GetEventDate() time.Time`

GetEventDate returns the EventDate field if non-nil, zero value otherwise.

### GetEventDateOk

`func (o *CustomNotification) GetEventDateOk() (*time.Time, bool)`

GetEventDateOk returns a tuple with the EventDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDate

`func (o *CustomNotification) SetEventDate(v time.Time)`

SetEventDate sets EventDate field to given value.

### HasEventDate

`func (o *CustomNotification) HasEventDate() bool`

HasEventDate returns a boolean if a field has been set.

### GetMerchantReference

`func (o *CustomNotification) GetMerchantReference() string`

GetMerchantReference returns the MerchantReference field if non-nil, zero value otherwise.

### GetMerchantReferenceOk

`func (o *CustomNotification) GetMerchantReferenceOk() (*string, bool)`

GetMerchantReferenceOk returns a tuple with the MerchantReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReference

`func (o *CustomNotification) SetMerchantReference(v string)`

SetMerchantReference sets MerchantReference field to given value.

### HasMerchantReference

`func (o *CustomNotification) HasMerchantReference() bool`

HasMerchantReference returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *CustomNotification) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CustomNotification) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CustomNotification) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *CustomNotification) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetReason

`func (o *CustomNotification) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CustomNotification) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CustomNotification) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CustomNotification) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSuccess

`func (o *CustomNotification) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CustomNotification) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CustomNotification) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CustomNotification) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


