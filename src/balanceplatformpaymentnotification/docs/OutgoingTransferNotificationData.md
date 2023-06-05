# OutgoingTransferNotificationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolder** | Pointer to [**ResourceReference**](ResourceReference.md) |  | [optional] 
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**BalanceAccount** | Pointer to [**ResourceReference**](ResourceReference.md) |  | [optional] 
**BalancePlatform** | Pointer to **string** | The unique identifier of the balance platform. | [optional] 
**Counterparty** | [**Counterparty**](Counterparty.md) |  | 
**CreationDate** | Pointer to **time.Time** | The date and time when the event was triggered, in ISO 8601 extended format. For example, **2020-12-18T10:15:30+01:00**. | [optional] 
**Description** | Pointer to **string** | Your description for the transfer. If you send a description longer than 140 characters, the text is truncated. | [optional] 
**Id** | Pointer to **string** | The ID of the resource. | [optional] 
**MerchantData** | Pointer to [**MerchantData**](MerchantData.md) |  | [optional] 
**Modification** | Pointer to [**NotificationModificationData**](NotificationModificationData.md) |  | [optional] 
**OriginalAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**PaymentId** | Pointer to **string** | The ID of the original payment authorisation, refund, or funds transfer request. Use this to trace the original request from the &#x60;balancePlatform.payment.created&#x60; webhook. | [optional] 
**PaymentInstrument** | Pointer to [**ResourceReference**](ResourceReference.md) |  | [optional] 
**PlatformPayment** | Pointer to [**PlatformPayment**](PlatformPayment.md) |  | [optional] 
**ProcessingType** | Pointer to **string** | Contains information about how the payment was processed. Possible values: **atmWithdraw**, **balanceInquiry**, **ecommerce**, **moto**, **pos**, **purchaseWithCashback**, **recurring**, **token**, **unknown**. | [optional] 
**PurposeCode** | Pointer to **string** | Indicates the purpose of the outgoing transfer. Adyen sets this to:  * **payoutManual** when the transfer was triggered by a one-off payout using the [&#x60;/transfers&#x60;](https://docs.adyen.com/api-explorer/#/transfers/latest/post/transfers) endpoint.   * **payoutSweep** when the transfer was triggered by a scheduled payout using [&#x60;sweepConfigurations&#x60;](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_sweepConfigurations). | [optional] 
**Reference** | Pointer to **string** | The [&#x60;reference&#x60;](https://docs.adyen.com/api-explorer/#/transfers/latest/post/transfers__reqParam_reference) from the &#x60;/transfers&#x60; request. If you haven&#39;t provided any, Adyen generates a unique reference. | [optional] 
**ReferenceForBeneficiary** | Pointer to **string** | The reference sent to or received from the counterparty.  * For outgoing funds, this is the [&#x60;referenceForBeneficiary&#x60;](https://docs.adyen.com/api-explorer/#/transfers/latest/post/transfers__resParam_referenceForBeneficiary) from the  [&#x60;/transfers&#x60;](https://docs.adyen.com/api-explorer/#/transfers/latest/post/transfers) request.   * For incoming funds, this is the reference from the sender. | [optional] 
**RelayedAuthorisationData** | Pointer to [**RelayedAuthorisationData**](RelayedAuthorisationData.md) |  | [optional] 
**Status** | Pointer to **string** | The event status. The possible values depend on the &#x60;type&#x60;.  * **Authorised**, **Refused**, or **Error** for type &#x60;balancePlatform.payment.created&#x60;   * **Expired** or **Cancelled** or **AuthAdjustmentAuthorised** or **AuthAdjustmentRefused** for type &#x60;balancePlatform.payment.updated&#x60;  * **PendingIncomingTransfer** for type &#x60;balancePlatform.incomingTransfer.created&#x60;   * **Refunded** or **IncomingTransfer** for type &#x60;balancePlatform.incomingTransfer.updated&#x60;   * **Captured** or **OutgoingTransfer** for type &#x60;balancePlatform.outgoingTransfer.created&#x60;  * **TransferConfirmed**, **TransferSentOut**, or **TransferFailed** for type &#x60;balancePlatform.outgoingTransfer.updated&#x60;     | [optional] 
**TransactionRulesResult** | Pointer to [**TransactionRulesResult**](TransactionRulesResult.md) |  | [optional] 
**ValueDate** | Pointer to **time.Time** | Indicates the expected settlement date of this transaction, in ISO 8601 extended format. For example, **2021-08-17T15:34:37+02:00**. | [optional] 

## Methods

### NewOutgoingTransferNotificationData

`func NewOutgoingTransferNotificationData(counterparty Counterparty, ) *OutgoingTransferNotificationData`

NewOutgoingTransferNotificationData instantiates a new OutgoingTransferNotificationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingTransferNotificationDataWithDefaults

`func NewOutgoingTransferNotificationDataWithDefaults() *OutgoingTransferNotificationData`

NewOutgoingTransferNotificationDataWithDefaults instantiates a new OutgoingTransferNotificationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolder

`func (o *OutgoingTransferNotificationData) GetAccountHolder() ResourceReference`

GetAccountHolder returns the AccountHolder field if non-nil, zero value otherwise.

### GetAccountHolderOk

`func (o *OutgoingTransferNotificationData) GetAccountHolderOk() (*ResourceReference, bool)`

GetAccountHolderOk returns a tuple with the AccountHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolder

`func (o *OutgoingTransferNotificationData) SetAccountHolder(v ResourceReference)`

SetAccountHolder sets AccountHolder field to given value.

### HasAccountHolder

`func (o *OutgoingTransferNotificationData) HasAccountHolder() bool`

HasAccountHolder returns a boolean if a field has been set.

### GetAmount

`func (o *OutgoingTransferNotificationData) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OutgoingTransferNotificationData) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OutgoingTransferNotificationData) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *OutgoingTransferNotificationData) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBalanceAccount

`func (o *OutgoingTransferNotificationData) GetBalanceAccount() ResourceReference`

GetBalanceAccount returns the BalanceAccount field if non-nil, zero value otherwise.

### GetBalanceAccountOk

`func (o *OutgoingTransferNotificationData) GetBalanceAccountOk() (*ResourceReference, bool)`

GetBalanceAccountOk returns a tuple with the BalanceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccount

`func (o *OutgoingTransferNotificationData) SetBalanceAccount(v ResourceReference)`

SetBalanceAccount sets BalanceAccount field to given value.

### HasBalanceAccount

`func (o *OutgoingTransferNotificationData) HasBalanceAccount() bool`

HasBalanceAccount returns a boolean if a field has been set.

### GetBalancePlatform

`func (o *OutgoingTransferNotificationData) GetBalancePlatform() string`

GetBalancePlatform returns the BalancePlatform field if non-nil, zero value otherwise.

### GetBalancePlatformOk

`func (o *OutgoingTransferNotificationData) GetBalancePlatformOk() (*string, bool)`

GetBalancePlatformOk returns a tuple with the BalancePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancePlatform

`func (o *OutgoingTransferNotificationData) SetBalancePlatform(v string)`

SetBalancePlatform sets BalancePlatform field to given value.

### HasBalancePlatform

`func (o *OutgoingTransferNotificationData) HasBalancePlatform() bool`

HasBalancePlatform returns a boolean if a field has been set.

### GetCounterparty

`func (o *OutgoingTransferNotificationData) GetCounterparty() Counterparty`

GetCounterparty returns the Counterparty field if non-nil, zero value otherwise.

### GetCounterpartyOk

`func (o *OutgoingTransferNotificationData) GetCounterpartyOk() (*Counterparty, bool)`

GetCounterpartyOk returns a tuple with the Counterparty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterparty

`func (o *OutgoingTransferNotificationData) SetCounterparty(v Counterparty)`

SetCounterparty sets Counterparty field to given value.


### GetCreationDate

`func (o *OutgoingTransferNotificationData) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *OutgoingTransferNotificationData) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *OutgoingTransferNotificationData) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *OutgoingTransferNotificationData) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDescription

`func (o *OutgoingTransferNotificationData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutgoingTransferNotificationData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutgoingTransferNotificationData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OutgoingTransferNotificationData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *OutgoingTransferNotificationData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutgoingTransferNotificationData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutgoingTransferNotificationData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OutgoingTransferNotificationData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantData

`func (o *OutgoingTransferNotificationData) GetMerchantData() MerchantData`

GetMerchantData returns the MerchantData field if non-nil, zero value otherwise.

### GetMerchantDataOk

`func (o *OutgoingTransferNotificationData) GetMerchantDataOk() (*MerchantData, bool)`

GetMerchantDataOk returns a tuple with the MerchantData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantData

`func (o *OutgoingTransferNotificationData) SetMerchantData(v MerchantData)`

SetMerchantData sets MerchantData field to given value.

### HasMerchantData

`func (o *OutgoingTransferNotificationData) HasMerchantData() bool`

HasMerchantData returns a boolean if a field has been set.

### GetModification

`func (o *OutgoingTransferNotificationData) GetModification() NotificationModificationData`

GetModification returns the Modification field if non-nil, zero value otherwise.

### GetModificationOk

`func (o *OutgoingTransferNotificationData) GetModificationOk() (*NotificationModificationData, bool)`

GetModificationOk returns a tuple with the Modification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModification

`func (o *OutgoingTransferNotificationData) SetModification(v NotificationModificationData)`

SetModification sets Modification field to given value.

### HasModification

`func (o *OutgoingTransferNotificationData) HasModification() bool`

HasModification returns a boolean if a field has been set.

### GetOriginalAmount

`func (o *OutgoingTransferNotificationData) GetOriginalAmount() Amount`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *OutgoingTransferNotificationData) GetOriginalAmountOk() (*Amount, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *OutgoingTransferNotificationData) SetOriginalAmount(v Amount)`

SetOriginalAmount sets OriginalAmount field to given value.

### HasOriginalAmount

`func (o *OutgoingTransferNotificationData) HasOriginalAmount() bool`

HasOriginalAmount returns a boolean if a field has been set.

### GetPaymentId

`func (o *OutgoingTransferNotificationData) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *OutgoingTransferNotificationData) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *OutgoingTransferNotificationData) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *OutgoingTransferNotificationData) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPaymentInstrument

`func (o *OutgoingTransferNotificationData) GetPaymentInstrument() ResourceReference`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *OutgoingTransferNotificationData) GetPaymentInstrumentOk() (*ResourceReference, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *OutgoingTransferNotificationData) SetPaymentInstrument(v ResourceReference)`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *OutgoingTransferNotificationData) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### GetPlatformPayment

`func (o *OutgoingTransferNotificationData) GetPlatformPayment() PlatformPayment`

GetPlatformPayment returns the PlatformPayment field if non-nil, zero value otherwise.

### GetPlatformPaymentOk

`func (o *OutgoingTransferNotificationData) GetPlatformPaymentOk() (*PlatformPayment, bool)`

GetPlatformPaymentOk returns a tuple with the PlatformPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPayment

`func (o *OutgoingTransferNotificationData) SetPlatformPayment(v PlatformPayment)`

SetPlatformPayment sets PlatformPayment field to given value.

### HasPlatformPayment

`func (o *OutgoingTransferNotificationData) HasPlatformPayment() bool`

HasPlatformPayment returns a boolean if a field has been set.

### GetProcessingType

`func (o *OutgoingTransferNotificationData) GetProcessingType() string`

GetProcessingType returns the ProcessingType field if non-nil, zero value otherwise.

### GetProcessingTypeOk

`func (o *OutgoingTransferNotificationData) GetProcessingTypeOk() (*string, bool)`

GetProcessingTypeOk returns a tuple with the ProcessingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingType

`func (o *OutgoingTransferNotificationData) SetProcessingType(v string)`

SetProcessingType sets ProcessingType field to given value.

### HasProcessingType

`func (o *OutgoingTransferNotificationData) HasProcessingType() bool`

HasProcessingType returns a boolean if a field has been set.

### GetPurposeCode

`func (o *OutgoingTransferNotificationData) GetPurposeCode() string`

GetPurposeCode returns the PurposeCode field if non-nil, zero value otherwise.

### GetPurposeCodeOk

`func (o *OutgoingTransferNotificationData) GetPurposeCodeOk() (*string, bool)`

GetPurposeCodeOk returns a tuple with the PurposeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeCode

`func (o *OutgoingTransferNotificationData) SetPurposeCode(v string)`

SetPurposeCode sets PurposeCode field to given value.

### HasPurposeCode

`func (o *OutgoingTransferNotificationData) HasPurposeCode() bool`

HasPurposeCode returns a boolean if a field has been set.

### GetReference

`func (o *OutgoingTransferNotificationData) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *OutgoingTransferNotificationData) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *OutgoingTransferNotificationData) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *OutgoingTransferNotificationData) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceForBeneficiary

`func (o *OutgoingTransferNotificationData) GetReferenceForBeneficiary() string`

GetReferenceForBeneficiary returns the ReferenceForBeneficiary field if non-nil, zero value otherwise.

### GetReferenceForBeneficiaryOk

`func (o *OutgoingTransferNotificationData) GetReferenceForBeneficiaryOk() (*string, bool)`

GetReferenceForBeneficiaryOk returns a tuple with the ReferenceForBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceForBeneficiary

`func (o *OutgoingTransferNotificationData) SetReferenceForBeneficiary(v string)`

SetReferenceForBeneficiary sets ReferenceForBeneficiary field to given value.

### HasReferenceForBeneficiary

`func (o *OutgoingTransferNotificationData) HasReferenceForBeneficiary() bool`

HasReferenceForBeneficiary returns a boolean if a field has been set.

### GetRelayedAuthorisationData

`func (o *OutgoingTransferNotificationData) GetRelayedAuthorisationData() RelayedAuthorisationData`

GetRelayedAuthorisationData returns the RelayedAuthorisationData field if non-nil, zero value otherwise.

### GetRelayedAuthorisationDataOk

`func (o *OutgoingTransferNotificationData) GetRelayedAuthorisationDataOk() (*RelayedAuthorisationData, bool)`

GetRelayedAuthorisationDataOk returns a tuple with the RelayedAuthorisationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayedAuthorisationData

`func (o *OutgoingTransferNotificationData) SetRelayedAuthorisationData(v RelayedAuthorisationData)`

SetRelayedAuthorisationData sets RelayedAuthorisationData field to given value.

### HasRelayedAuthorisationData

`func (o *OutgoingTransferNotificationData) HasRelayedAuthorisationData() bool`

HasRelayedAuthorisationData returns a boolean if a field has been set.

### GetStatus

`func (o *OutgoingTransferNotificationData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OutgoingTransferNotificationData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OutgoingTransferNotificationData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OutgoingTransferNotificationData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransactionRulesResult

`func (o *OutgoingTransferNotificationData) GetTransactionRulesResult() TransactionRulesResult`

GetTransactionRulesResult returns the TransactionRulesResult field if non-nil, zero value otherwise.

### GetTransactionRulesResultOk

`func (o *OutgoingTransferNotificationData) GetTransactionRulesResultOk() (*TransactionRulesResult, bool)`

GetTransactionRulesResultOk returns a tuple with the TransactionRulesResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRulesResult

`func (o *OutgoingTransferNotificationData) SetTransactionRulesResult(v TransactionRulesResult)`

SetTransactionRulesResult sets TransactionRulesResult field to given value.

### HasTransactionRulesResult

`func (o *OutgoingTransferNotificationData) HasTransactionRulesResult() bool`

HasTransactionRulesResult returns a boolean if a field has been set.

### GetValueDate

`func (o *OutgoingTransferNotificationData) GetValueDate() time.Time`

GetValueDate returns the ValueDate field if non-nil, zero value otherwise.

### GetValueDateOk

`func (o *OutgoingTransferNotificationData) GetValueDateOk() (*time.Time, bool)`

GetValueDateOk returns a tuple with the ValueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDate

`func (o *OutgoingTransferNotificationData) SetValueDate(v time.Time)`

SetValueDate sets ValueDate field to given value.

### HasValueDate

`func (o *OutgoingTransferNotificationData) HasValueDate() bool`

HasValueDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


