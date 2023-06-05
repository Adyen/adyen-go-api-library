# IncomingTransferNotificationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolder** | Pointer to [**ResourceReference**](ResourceReference.md) |  | [optional] 
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**BalanceAccount** | Pointer to [**ResourceReference**](ResourceReference.md) |  | [optional] 
**BalancePlatform** | Pointer to **string** | The unique identifier of the balance platform. | [optional] 
**Counterparty** | Pointer to [**Counterparty**](Counterparty.md) |  | [optional] 
**CreationDate** | Pointer to **time.Time** | The date and time when the event was triggered, in ISO 8601 extended format. For example, **2020-12-18T10:15:30+01:00**. | [optional] 
**Description** | Pointer to **string** | Your description for the transfer. If you send a description longer than 140 characters, the text is truncated. | [optional] 
**Id** | Pointer to **string** | The ID of the resource. | [optional] 
**Modification** | Pointer to [**NotificationModificationData**](NotificationModificationData.md) |  | [optional] 
**OriginalAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**PaymentId** | Pointer to **string** | The ID of the original payment authorisation, refund, or funds transfer request. Use this to trace the original request from the &#x60;balancePlatform.payment.created&#x60; webhook. | [optional] 
**PaymentInstrument** | Pointer to [**ResourceReference**](ResourceReference.md) |  | [optional] 
**PlatformPayment** | Pointer to [**PlatformPayment**](PlatformPayment.md) |  | [optional] 
**Reference** | Pointer to **string** | An Adyen-generated unique reference for the transfer. | [optional] 
**ReferenceForBeneficiary** | Pointer to **string** | The reference sent to or received from the counterparty.  * For outgoing funds, this is the [&#x60;referenceForBeneficiary&#x60;](https://docs.adyen.com/api-explorer/#/transfers/latest/post/transfers__resParam_referenceForBeneficiary) from the  [&#x60;/transfers&#x60;](https://docs.adyen.com/api-explorer/#/transfers/latest/post/transfers) request.   * For incoming funds, this is the reference from the sender. | [optional] 
**Status** | Pointer to **string** | The event status. The possible values depend on the &#x60;type&#x60;.  * **Authorised**, **Refused**, or **Error** for type &#x60;balancePlatform.payment.created&#x60;   * **Expired** or **Cancelled** or **AuthAdjustmentAuthorised** or **AuthAdjustmentRefused** for type &#x60;balancePlatform.payment.updated&#x60;  * **PendingIncomingTransfer** for type &#x60;balancePlatform.incomingTransfer.created&#x60;   * **Refunded** or **IncomingTransfer** for type &#x60;balancePlatform.incomingTransfer.updated&#x60;   * **Captured** or **OutgoingTransfer** for type &#x60;balancePlatform.outgoingTransfer.created&#x60;  * **TransferConfirmed**, **TransferSentOut**, or **TransferFailed** for type &#x60;balancePlatform.outgoingTransfer.updated&#x60;     | [optional] 
**ValueDate** | Pointer to **time.Time** | Indicates the expected settlement date of this transaction, in ISO 8601 extended format. For example, **2021-08-17T15:34:37+02:00**. | [optional] 

## Methods

### NewIncomingTransferNotificationData

`func NewIncomingTransferNotificationData() *IncomingTransferNotificationData`

NewIncomingTransferNotificationData instantiates a new IncomingTransferNotificationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingTransferNotificationDataWithDefaults

`func NewIncomingTransferNotificationDataWithDefaults() *IncomingTransferNotificationData`

NewIncomingTransferNotificationDataWithDefaults instantiates a new IncomingTransferNotificationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolder

`func (o *IncomingTransferNotificationData) GetAccountHolder() ResourceReference`

GetAccountHolder returns the AccountHolder field if non-nil, zero value otherwise.

### GetAccountHolderOk

`func (o *IncomingTransferNotificationData) GetAccountHolderOk() (*ResourceReference, bool)`

GetAccountHolderOk returns a tuple with the AccountHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolder

`func (o *IncomingTransferNotificationData) SetAccountHolder(v ResourceReference)`

SetAccountHolder sets AccountHolder field to given value.

### HasAccountHolder

`func (o *IncomingTransferNotificationData) HasAccountHolder() bool`

HasAccountHolder returns a boolean if a field has been set.

### GetAmount

`func (o *IncomingTransferNotificationData) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *IncomingTransferNotificationData) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *IncomingTransferNotificationData) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *IncomingTransferNotificationData) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBalanceAccount

`func (o *IncomingTransferNotificationData) GetBalanceAccount() ResourceReference`

GetBalanceAccount returns the BalanceAccount field if non-nil, zero value otherwise.

### GetBalanceAccountOk

`func (o *IncomingTransferNotificationData) GetBalanceAccountOk() (*ResourceReference, bool)`

GetBalanceAccountOk returns a tuple with the BalanceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccount

`func (o *IncomingTransferNotificationData) SetBalanceAccount(v ResourceReference)`

SetBalanceAccount sets BalanceAccount field to given value.

### HasBalanceAccount

`func (o *IncomingTransferNotificationData) HasBalanceAccount() bool`

HasBalanceAccount returns a boolean if a field has been set.

### GetBalancePlatform

`func (o *IncomingTransferNotificationData) GetBalancePlatform() string`

GetBalancePlatform returns the BalancePlatform field if non-nil, zero value otherwise.

### GetBalancePlatformOk

`func (o *IncomingTransferNotificationData) GetBalancePlatformOk() (*string, bool)`

GetBalancePlatformOk returns a tuple with the BalancePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancePlatform

`func (o *IncomingTransferNotificationData) SetBalancePlatform(v string)`

SetBalancePlatform sets BalancePlatform field to given value.

### HasBalancePlatform

`func (o *IncomingTransferNotificationData) HasBalancePlatform() bool`

HasBalancePlatform returns a boolean if a field has been set.

### GetCounterparty

`func (o *IncomingTransferNotificationData) GetCounterparty() Counterparty`

GetCounterparty returns the Counterparty field if non-nil, zero value otherwise.

### GetCounterpartyOk

`func (o *IncomingTransferNotificationData) GetCounterpartyOk() (*Counterparty, bool)`

GetCounterpartyOk returns a tuple with the Counterparty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterparty

`func (o *IncomingTransferNotificationData) SetCounterparty(v Counterparty)`

SetCounterparty sets Counterparty field to given value.

### HasCounterparty

`func (o *IncomingTransferNotificationData) HasCounterparty() bool`

HasCounterparty returns a boolean if a field has been set.

### GetCreationDate

`func (o *IncomingTransferNotificationData) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *IncomingTransferNotificationData) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *IncomingTransferNotificationData) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *IncomingTransferNotificationData) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDescription

`func (o *IncomingTransferNotificationData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IncomingTransferNotificationData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IncomingTransferNotificationData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IncomingTransferNotificationData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *IncomingTransferNotificationData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncomingTransferNotificationData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncomingTransferNotificationData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IncomingTransferNotificationData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModification

`func (o *IncomingTransferNotificationData) GetModification() NotificationModificationData`

GetModification returns the Modification field if non-nil, zero value otherwise.

### GetModificationOk

`func (o *IncomingTransferNotificationData) GetModificationOk() (*NotificationModificationData, bool)`

GetModificationOk returns a tuple with the Modification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModification

`func (o *IncomingTransferNotificationData) SetModification(v NotificationModificationData)`

SetModification sets Modification field to given value.

### HasModification

`func (o *IncomingTransferNotificationData) HasModification() bool`

HasModification returns a boolean if a field has been set.

### GetOriginalAmount

`func (o *IncomingTransferNotificationData) GetOriginalAmount() Amount`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *IncomingTransferNotificationData) GetOriginalAmountOk() (*Amount, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *IncomingTransferNotificationData) SetOriginalAmount(v Amount)`

SetOriginalAmount sets OriginalAmount field to given value.

### HasOriginalAmount

`func (o *IncomingTransferNotificationData) HasOriginalAmount() bool`

HasOriginalAmount returns a boolean if a field has been set.

### GetPaymentId

`func (o *IncomingTransferNotificationData) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *IncomingTransferNotificationData) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *IncomingTransferNotificationData) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *IncomingTransferNotificationData) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPaymentInstrument

`func (o *IncomingTransferNotificationData) GetPaymentInstrument() ResourceReference`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *IncomingTransferNotificationData) GetPaymentInstrumentOk() (*ResourceReference, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *IncomingTransferNotificationData) SetPaymentInstrument(v ResourceReference)`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *IncomingTransferNotificationData) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### GetPlatformPayment

`func (o *IncomingTransferNotificationData) GetPlatformPayment() PlatformPayment`

GetPlatformPayment returns the PlatformPayment field if non-nil, zero value otherwise.

### GetPlatformPaymentOk

`func (o *IncomingTransferNotificationData) GetPlatformPaymentOk() (*PlatformPayment, bool)`

GetPlatformPaymentOk returns a tuple with the PlatformPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPayment

`func (o *IncomingTransferNotificationData) SetPlatformPayment(v PlatformPayment)`

SetPlatformPayment sets PlatformPayment field to given value.

### HasPlatformPayment

`func (o *IncomingTransferNotificationData) HasPlatformPayment() bool`

HasPlatformPayment returns a boolean if a field has been set.

### GetReference

`func (o *IncomingTransferNotificationData) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *IncomingTransferNotificationData) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *IncomingTransferNotificationData) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *IncomingTransferNotificationData) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceForBeneficiary

`func (o *IncomingTransferNotificationData) GetReferenceForBeneficiary() string`

GetReferenceForBeneficiary returns the ReferenceForBeneficiary field if non-nil, zero value otherwise.

### GetReferenceForBeneficiaryOk

`func (o *IncomingTransferNotificationData) GetReferenceForBeneficiaryOk() (*string, bool)`

GetReferenceForBeneficiaryOk returns a tuple with the ReferenceForBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceForBeneficiary

`func (o *IncomingTransferNotificationData) SetReferenceForBeneficiary(v string)`

SetReferenceForBeneficiary sets ReferenceForBeneficiary field to given value.

### HasReferenceForBeneficiary

`func (o *IncomingTransferNotificationData) HasReferenceForBeneficiary() bool`

HasReferenceForBeneficiary returns a boolean if a field has been set.

### GetStatus

`func (o *IncomingTransferNotificationData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IncomingTransferNotificationData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IncomingTransferNotificationData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IncomingTransferNotificationData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetValueDate

`func (o *IncomingTransferNotificationData) GetValueDate() time.Time`

GetValueDate returns the ValueDate field if non-nil, zero value otherwise.

### GetValueDateOk

`func (o *IncomingTransferNotificationData) GetValueDateOk() (*time.Time, bool)`

GetValueDateOk returns a tuple with the ValueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDate

`func (o *IncomingTransferNotificationData) SetValueDate(v time.Time)`

SetValueDate sets ValueDate field to given value.

### HasValueDate

`func (o *IncomingTransferNotificationData) HasValueDate() bool`

HasValueDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


