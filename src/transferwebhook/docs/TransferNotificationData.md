# TransferNotificationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolder** | Pointer to [**ResourceReference**](ResourceReference.md) |  | [optional] 
**Amount** | [**Amount**](Amount.md) |  | 
**BalanceAccount** | Pointer to [**ResourceReference**](ResourceReference.md) |  | [optional] 
**BalanceAccountId** | Pointer to **string** | The unique identifier of the source [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id). | [optional] 
**BalancePlatform** | Pointer to **string** | The unique identifier of the balance platform. | [optional] 
**Balances** | Pointer to [**[]BalanceMutation**](BalanceMutation.md) | The list of the latest balance statuses in the transfer. | [optional] 
**Category** | **string** | The type of transfer.  Possible values:   - **bank**: Transfer to a [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id) or a bank account.  - **internal**: Transfer to another [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id) within your platform.  - **issuedCard**: Transfer initiated by a Adyen-issued card.  - **platformPayment**: Fund movements related to payments that are acquired for your users. | 
**Counterparty** | Pointer to [**CounterpartyV3**](CounterpartyV3.md) |  | [optional] 
**CreationDate** | Pointer to **time.Time** | The date and time when the event was triggered, in ISO 8601 extended format. For example, **2020-12-18T10:15:30+01:00**. | [optional] 
**Description** | Pointer to **string** | Your description for the transfer. It is used by most banks as the transfer description. We recommend sending a maximum of 140 characters, otherwise the description may be truncated.  Supported characters: **[a-z] [A-Z] [0-9] / - ?** **: ( ) . , &#39; + Space**  Supported characters for **regular** and **fast** transfers to a US counterparty: **[a-z] [A-Z] [0-9] &amp; $ % # @** **~ &#x3D; + - _ &#39; \&quot; ! ?** | [optional] 
**Direction** | Pointer to **string** | The direction of the transfer.  Possible values: **incoming**, **outgoing**. | [optional] 
**Events** | Pointer to [**[]TransferEvent**](TransferEvent.md) | The list of events leading up to the current status of the transfer. | [optional] 
**Id** | Pointer to **string** | The ID of the resource. | [optional] 
**ModificationMerchantReference** | Pointer to **string** | The capture&#39;s merchant reference included in the transfer. | [optional] 
**ModificationPspReference** | Pointer to **string** | The capture reference included in the transfer. | [optional] 
**PanEntryMode** | Pointer to **string** | Indicates the method used for entering the PAN to initiate a transaction.  Possible values: **manual**, **chip**, **magstripe**, **contactless**, **cof**, **ecommerce**, **token**. | [optional] 
**PaymentInstrument** | Pointer to [**PaymentInstrument**](PaymentInstrument.md) |  | [optional] 
**PaymentInstrumentId** | Pointer to **string** | The unique identifier of the [payment instrument](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id) used in the transfer. | [optional] 
**PaymentMerchantReference** | Pointer to **string** | The payment&#39;s merchant reference included in the transfer. | [optional] 
**Priority** | Pointer to **string** | The priority for the bank transfer. This sets the speed at which the transfer is sent and the fees that you have to pay. Required for transfers with &#x60;category&#x60; **bank**.  Possible values:  * **regular**: For normal, low-value transactions.  * **fast**: Faster way to transfer funds but has higher fees. Recommended for high-priority, low-value transactions.  * **wire**: Fastest way to transfer funds but has the highest fees. Recommended for high-priority, high-value transactions.  * **instant**: Instant way to transfer funds in [SEPA countries](https://www.ecb.europa.eu/paym/integration/retail/sepa/html/index.en.html).  * **crossBorder**: High-value transfer to a recipient in a different country.  * **internal**: Transfer to an Adyen-issued business bank account (by bank account number/IBAN). | [optional] 
**ProcessingType** | Pointer to **string** | Contains information about how the payment was processed. For example, **ecommerce** for online or **pos** for in-person payments. | [optional] 
**PspPaymentReference** | Pointer to **string** | The payment reference included in the transfer. | [optional] 
**Reason** | Pointer to **string** | Additional information about the status of the transfer. | [optional] 
**Reference** | Pointer to **string** | Your reference for the transfer, used internally within your platform. If you don&#39;t provide this in the request, Adyen generates a unique reference. | [optional] 
**ReferenceForBeneficiary** | Pointer to **string** |  A reference that is sent to the recipient. This reference is also sent in all webhooks related to the transfer, so you can use it to track statuses for both the source and recipient of funds.   Supported characters: **a-z**, **A-Z**, **0-9**. The maximum length depends on the &#x60;category&#x60;.  - **internal**: 80 characters  - **bank**: 35 characters when transferring to an IBAN, 15 characters for others. | [optional] 
**RelayedAuthorisationData** | Pointer to [**RelayedAuthorisationData**](RelayedAuthorisationData.md) |  | [optional] 
**SequenceNumber** | Pointer to **int32** | The sequence number of the transfer notification. The numbers start from 1 and increase with each new notification for a specific transfer.  It can help you restore the correct sequence of events even if they arrive out of order. | [optional] 
**Status** | **string** | The result of the transfer.   For example, **authorised**, **refused**, or **error**. | 
**Tracking** | Pointer to [**TransferNotificationTransferTracking**](TransferNotificationTransferTracking.md) |  | [optional] 
**TransactionId** | Pointer to **string** | The ID of the transaction that is created based on the transfer. | [optional] 
**TransactionRulesResult** | Pointer to [**TransactionRulesResult**](TransactionRulesResult.md) |  | [optional] 
**Type** | Pointer to **string** | The type of transfer or transaction. For example, **refund**, **payment**, **internalTransfer**, **bankTransfer**. | [optional] 
**ValidationFacts** | Pointer to [**[]TransferNotificationValidationFact**](TransferNotificationValidationFact.md) | The evaluation of the validation facts. See [validation checks](https://docs.adyen.com/issuing/validation-checks) for more information. | [optional] 

## Methods

### NewTransferNotificationData

`func NewTransferNotificationData(amount Amount, category string, status string, ) *TransferNotificationData`

NewTransferNotificationData instantiates a new TransferNotificationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferNotificationDataWithDefaults

`func NewTransferNotificationDataWithDefaults() *TransferNotificationData`

NewTransferNotificationDataWithDefaults instantiates a new TransferNotificationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolder

`func (o *TransferNotificationData) GetAccountHolder() ResourceReference`

GetAccountHolder returns the AccountHolder field if non-nil, zero value otherwise.

### GetAccountHolderOk

`func (o *TransferNotificationData) GetAccountHolderOk() (*ResourceReference, bool)`

GetAccountHolderOk returns a tuple with the AccountHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolder

`func (o *TransferNotificationData) SetAccountHolder(v ResourceReference)`

SetAccountHolder sets AccountHolder field to given value.

### HasAccountHolder

`func (o *TransferNotificationData) HasAccountHolder() bool`

HasAccountHolder returns a boolean if a field has been set.

### GetAmount

`func (o *TransferNotificationData) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferNotificationData) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferNotificationData) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetBalanceAccount

`func (o *TransferNotificationData) GetBalanceAccount() ResourceReference`

GetBalanceAccount returns the BalanceAccount field if non-nil, zero value otherwise.

### GetBalanceAccountOk

`func (o *TransferNotificationData) GetBalanceAccountOk() (*ResourceReference, bool)`

GetBalanceAccountOk returns a tuple with the BalanceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccount

`func (o *TransferNotificationData) SetBalanceAccount(v ResourceReference)`

SetBalanceAccount sets BalanceAccount field to given value.

### HasBalanceAccount

`func (o *TransferNotificationData) HasBalanceAccount() bool`

HasBalanceAccount returns a boolean if a field has been set.

### GetBalanceAccountId

`func (o *TransferNotificationData) GetBalanceAccountId() string`

GetBalanceAccountId returns the BalanceAccountId field if non-nil, zero value otherwise.

### GetBalanceAccountIdOk

`func (o *TransferNotificationData) GetBalanceAccountIdOk() (*string, bool)`

GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccountId

`func (o *TransferNotificationData) SetBalanceAccountId(v string)`

SetBalanceAccountId sets BalanceAccountId field to given value.

### HasBalanceAccountId

`func (o *TransferNotificationData) HasBalanceAccountId() bool`

HasBalanceAccountId returns a boolean if a field has been set.

### GetBalancePlatform

`func (o *TransferNotificationData) GetBalancePlatform() string`

GetBalancePlatform returns the BalancePlatform field if non-nil, zero value otherwise.

### GetBalancePlatformOk

`func (o *TransferNotificationData) GetBalancePlatformOk() (*string, bool)`

GetBalancePlatformOk returns a tuple with the BalancePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancePlatform

`func (o *TransferNotificationData) SetBalancePlatform(v string)`

SetBalancePlatform sets BalancePlatform field to given value.

### HasBalancePlatform

`func (o *TransferNotificationData) HasBalancePlatform() bool`

HasBalancePlatform returns a boolean if a field has been set.

### GetBalances

`func (o *TransferNotificationData) GetBalances() []BalanceMutation`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *TransferNotificationData) GetBalancesOk() (*[]BalanceMutation, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *TransferNotificationData) SetBalances(v []BalanceMutation)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *TransferNotificationData) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetCategory

`func (o *TransferNotificationData) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TransferNotificationData) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TransferNotificationData) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetCounterparty

`func (o *TransferNotificationData) GetCounterparty() CounterpartyV3`

GetCounterparty returns the Counterparty field if non-nil, zero value otherwise.

### GetCounterpartyOk

`func (o *TransferNotificationData) GetCounterpartyOk() (*CounterpartyV3, bool)`

GetCounterpartyOk returns a tuple with the Counterparty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterparty

`func (o *TransferNotificationData) SetCounterparty(v CounterpartyV3)`

SetCounterparty sets Counterparty field to given value.

### HasCounterparty

`func (o *TransferNotificationData) HasCounterparty() bool`

HasCounterparty returns a boolean if a field has been set.

### GetCreationDate

`func (o *TransferNotificationData) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *TransferNotificationData) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *TransferNotificationData) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *TransferNotificationData) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDescription

`func (o *TransferNotificationData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransferNotificationData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransferNotificationData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransferNotificationData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirection

`func (o *TransferNotificationData) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *TransferNotificationData) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *TransferNotificationData) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *TransferNotificationData) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetEvents

`func (o *TransferNotificationData) GetEvents() []TransferEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TransferNotificationData) GetEventsOk() (*[]TransferEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TransferNotificationData) SetEvents(v []TransferEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *TransferNotificationData) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetId

`func (o *TransferNotificationData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferNotificationData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferNotificationData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransferNotificationData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModificationMerchantReference

`func (o *TransferNotificationData) GetModificationMerchantReference() string`

GetModificationMerchantReference returns the ModificationMerchantReference field if non-nil, zero value otherwise.

### GetModificationMerchantReferenceOk

`func (o *TransferNotificationData) GetModificationMerchantReferenceOk() (*string, bool)`

GetModificationMerchantReferenceOk returns a tuple with the ModificationMerchantReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationMerchantReference

`func (o *TransferNotificationData) SetModificationMerchantReference(v string)`

SetModificationMerchantReference sets ModificationMerchantReference field to given value.

### HasModificationMerchantReference

`func (o *TransferNotificationData) HasModificationMerchantReference() bool`

HasModificationMerchantReference returns a boolean if a field has been set.

### GetModificationPspReference

`func (o *TransferNotificationData) GetModificationPspReference() string`

GetModificationPspReference returns the ModificationPspReference field if non-nil, zero value otherwise.

### GetModificationPspReferenceOk

`func (o *TransferNotificationData) GetModificationPspReferenceOk() (*string, bool)`

GetModificationPspReferenceOk returns a tuple with the ModificationPspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationPspReference

`func (o *TransferNotificationData) SetModificationPspReference(v string)`

SetModificationPspReference sets ModificationPspReference field to given value.

### HasModificationPspReference

`func (o *TransferNotificationData) HasModificationPspReference() bool`

HasModificationPspReference returns a boolean if a field has been set.

### GetPanEntryMode

`func (o *TransferNotificationData) GetPanEntryMode() string`

GetPanEntryMode returns the PanEntryMode field if non-nil, zero value otherwise.

### GetPanEntryModeOk

`func (o *TransferNotificationData) GetPanEntryModeOk() (*string, bool)`

GetPanEntryModeOk returns a tuple with the PanEntryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanEntryMode

`func (o *TransferNotificationData) SetPanEntryMode(v string)`

SetPanEntryMode sets PanEntryMode field to given value.

### HasPanEntryMode

`func (o *TransferNotificationData) HasPanEntryMode() bool`

HasPanEntryMode returns a boolean if a field has been set.

### GetPaymentInstrument

`func (o *TransferNotificationData) GetPaymentInstrument() PaymentInstrument`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *TransferNotificationData) GetPaymentInstrumentOk() (*PaymentInstrument, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *TransferNotificationData) SetPaymentInstrument(v PaymentInstrument)`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *TransferNotificationData) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### GetPaymentInstrumentId

`func (o *TransferNotificationData) GetPaymentInstrumentId() string`

GetPaymentInstrumentId returns the PaymentInstrumentId field if non-nil, zero value otherwise.

### GetPaymentInstrumentIdOk

`func (o *TransferNotificationData) GetPaymentInstrumentIdOk() (*string, bool)`

GetPaymentInstrumentIdOk returns a tuple with the PaymentInstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrumentId

`func (o *TransferNotificationData) SetPaymentInstrumentId(v string)`

SetPaymentInstrumentId sets PaymentInstrumentId field to given value.

### HasPaymentInstrumentId

`func (o *TransferNotificationData) HasPaymentInstrumentId() bool`

HasPaymentInstrumentId returns a boolean if a field has been set.

### GetPaymentMerchantReference

`func (o *TransferNotificationData) GetPaymentMerchantReference() string`

GetPaymentMerchantReference returns the PaymentMerchantReference field if non-nil, zero value otherwise.

### GetPaymentMerchantReferenceOk

`func (o *TransferNotificationData) GetPaymentMerchantReferenceOk() (*string, bool)`

GetPaymentMerchantReferenceOk returns a tuple with the PaymentMerchantReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMerchantReference

`func (o *TransferNotificationData) SetPaymentMerchantReference(v string)`

SetPaymentMerchantReference sets PaymentMerchantReference field to given value.

### HasPaymentMerchantReference

`func (o *TransferNotificationData) HasPaymentMerchantReference() bool`

HasPaymentMerchantReference returns a boolean if a field has been set.

### GetPriority

`func (o *TransferNotificationData) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TransferNotificationData) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TransferNotificationData) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TransferNotificationData) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProcessingType

`func (o *TransferNotificationData) GetProcessingType() string`

GetProcessingType returns the ProcessingType field if non-nil, zero value otherwise.

### GetProcessingTypeOk

`func (o *TransferNotificationData) GetProcessingTypeOk() (*string, bool)`

GetProcessingTypeOk returns a tuple with the ProcessingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingType

`func (o *TransferNotificationData) SetProcessingType(v string)`

SetProcessingType sets ProcessingType field to given value.

### HasProcessingType

`func (o *TransferNotificationData) HasProcessingType() bool`

HasProcessingType returns a boolean if a field has been set.

### GetPspPaymentReference

`func (o *TransferNotificationData) GetPspPaymentReference() string`

GetPspPaymentReference returns the PspPaymentReference field if non-nil, zero value otherwise.

### GetPspPaymentReferenceOk

`func (o *TransferNotificationData) GetPspPaymentReferenceOk() (*string, bool)`

GetPspPaymentReferenceOk returns a tuple with the PspPaymentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspPaymentReference

`func (o *TransferNotificationData) SetPspPaymentReference(v string)`

SetPspPaymentReference sets PspPaymentReference field to given value.

### HasPspPaymentReference

`func (o *TransferNotificationData) HasPspPaymentReference() bool`

HasPspPaymentReference returns a boolean if a field has been set.

### GetReason

`func (o *TransferNotificationData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TransferNotificationData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TransferNotificationData) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TransferNotificationData) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReference

`func (o *TransferNotificationData) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *TransferNotificationData) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *TransferNotificationData) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *TransferNotificationData) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceForBeneficiary

`func (o *TransferNotificationData) GetReferenceForBeneficiary() string`

GetReferenceForBeneficiary returns the ReferenceForBeneficiary field if non-nil, zero value otherwise.

### GetReferenceForBeneficiaryOk

`func (o *TransferNotificationData) GetReferenceForBeneficiaryOk() (*string, bool)`

GetReferenceForBeneficiaryOk returns a tuple with the ReferenceForBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceForBeneficiary

`func (o *TransferNotificationData) SetReferenceForBeneficiary(v string)`

SetReferenceForBeneficiary sets ReferenceForBeneficiary field to given value.

### HasReferenceForBeneficiary

`func (o *TransferNotificationData) HasReferenceForBeneficiary() bool`

HasReferenceForBeneficiary returns a boolean if a field has been set.

### GetRelayedAuthorisationData

`func (o *TransferNotificationData) GetRelayedAuthorisationData() RelayedAuthorisationData`

GetRelayedAuthorisationData returns the RelayedAuthorisationData field if non-nil, zero value otherwise.

### GetRelayedAuthorisationDataOk

`func (o *TransferNotificationData) GetRelayedAuthorisationDataOk() (*RelayedAuthorisationData, bool)`

GetRelayedAuthorisationDataOk returns a tuple with the RelayedAuthorisationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayedAuthorisationData

`func (o *TransferNotificationData) SetRelayedAuthorisationData(v RelayedAuthorisationData)`

SetRelayedAuthorisationData sets RelayedAuthorisationData field to given value.

### HasRelayedAuthorisationData

`func (o *TransferNotificationData) HasRelayedAuthorisationData() bool`

HasRelayedAuthorisationData returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *TransferNotificationData) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *TransferNotificationData) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *TransferNotificationData) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *TransferNotificationData) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetStatus

`func (o *TransferNotificationData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransferNotificationData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransferNotificationData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTracking

`func (o *TransferNotificationData) GetTracking() TransferNotificationTransferTracking`

GetTracking returns the Tracking field if non-nil, zero value otherwise.

### GetTrackingOk

`func (o *TransferNotificationData) GetTrackingOk() (*TransferNotificationTransferTracking, bool)`

GetTrackingOk returns a tuple with the Tracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracking

`func (o *TransferNotificationData) SetTracking(v TransferNotificationTransferTracking)`

SetTracking sets Tracking field to given value.

### HasTracking

`func (o *TransferNotificationData) HasTracking() bool`

HasTracking returns a boolean if a field has been set.

### GetTransactionId

`func (o *TransferNotificationData) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TransferNotificationData) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TransferNotificationData) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *TransferNotificationData) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionRulesResult

`func (o *TransferNotificationData) GetTransactionRulesResult() TransactionRulesResult`

GetTransactionRulesResult returns the TransactionRulesResult field if non-nil, zero value otherwise.

### GetTransactionRulesResultOk

`func (o *TransferNotificationData) GetTransactionRulesResultOk() (*TransactionRulesResult, bool)`

GetTransactionRulesResultOk returns a tuple with the TransactionRulesResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRulesResult

`func (o *TransferNotificationData) SetTransactionRulesResult(v TransactionRulesResult)`

SetTransactionRulesResult sets TransactionRulesResult field to given value.

### HasTransactionRulesResult

`func (o *TransferNotificationData) HasTransactionRulesResult() bool`

HasTransactionRulesResult returns a boolean if a field has been set.

### GetType

`func (o *TransferNotificationData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransferNotificationData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransferNotificationData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransferNotificationData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValidationFacts

`func (o *TransferNotificationData) GetValidationFacts() []TransferNotificationValidationFact`

GetValidationFacts returns the ValidationFacts field if non-nil, zero value otherwise.

### GetValidationFactsOk

`func (o *TransferNotificationData) GetValidationFactsOk() (*[]TransferNotificationValidationFact, bool)`

GetValidationFactsOk returns a tuple with the ValidationFacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationFacts

`func (o *TransferNotificationData) SetValidationFacts(v []TransferNotificationValidationFact)`

SetValidationFacts sets ValidationFacts field to given value.

### HasValidationFacts

`func (o *TransferNotificationData) HasValidationFacts() bool`

HasValidationFacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


