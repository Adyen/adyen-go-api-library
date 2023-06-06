# transferwebhookData

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
**Tracking** | Pointer to [**transferwebhookTransferTracking**](transferwebhookTransferTracking.md) |  | [optional] 
**TransactionId** | Pointer to **string** | The ID of the transaction that is created based on the transfer. | [optional] 
**TransactionRulesResult** | Pointer to [**TransactionRulesResult**](TransactionRulesResult.md) |  | [optional] 
**Type** | Pointer to **string** | The type of transfer or transaction. For example, **refund**, **payment**, **internalTransfer**, **bankTransfer**. | [optional] 
**ValidationFacts** | Pointer to [**[]transferwebhookValidationFact**](transferwebhookValidationFact.md) | The evaluation of the validation facts. See [validation checks](https://docs.adyen.com/issuing/validation-checks) for more information. | [optional] 

## Methods

### NewtransferwebhookData

`func NewtransferwebhookData(amount Amount, category string, status string, ) *transferwebhookData`

NewtransferwebhookData instantiates a new transferwebhookData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewtransferwebhookDataWithDefaults

`func NewtransferwebhookDataWithDefaults() *transferwebhookData`

NewtransferwebhookDataWithDefaults instantiates a new transferwebhookData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolder

`func (o *transferwebhookData) GetAccountHolder() ResourceReference`

GetAccountHolder returns the AccountHolder field if non-nil, zero value otherwise.

### GetAccountHolderOk

`func (o *transferwebhookData) GetAccountHolderOk() (*ResourceReference, bool)`

GetAccountHolderOk returns a tuple with the AccountHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolder

`func (o *transferwebhookData) SetAccountHolder(v ResourceReference)`

SetAccountHolder sets AccountHolder field to given value.

### HasAccountHolder

`func (o *transferwebhookData) HasAccountHolder() bool`

HasAccountHolder returns a boolean if a field has been set.

### GetAmount

`func (o *transferwebhookData) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *transferwebhookData) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *transferwebhookData) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetBalanceAccount

`func (o *transferwebhookData) GetBalanceAccount() ResourceReference`

GetBalanceAccount returns the BalanceAccount field if non-nil, zero value otherwise.

### GetBalanceAccountOk

`func (o *transferwebhookData) GetBalanceAccountOk() (*ResourceReference, bool)`

GetBalanceAccountOk returns a tuple with the BalanceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccount

`func (o *transferwebhookData) SetBalanceAccount(v ResourceReference)`

SetBalanceAccount sets BalanceAccount field to given value.

### HasBalanceAccount

`func (o *transferwebhookData) HasBalanceAccount() bool`

HasBalanceAccount returns a boolean if a field has been set.

### GetBalanceAccountId

`func (o *transferwebhookData) GetBalanceAccountId() string`

GetBalanceAccountId returns the BalanceAccountId field if non-nil, zero value otherwise.

### GetBalanceAccountIdOk

`func (o *transferwebhookData) GetBalanceAccountIdOk() (*string, bool)`

GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccountId

`func (o *transferwebhookData) SetBalanceAccountId(v string)`

SetBalanceAccountId sets BalanceAccountId field to given value.

### HasBalanceAccountId

`func (o *transferwebhookData) HasBalanceAccountId() bool`

HasBalanceAccountId returns a boolean if a field has been set.

### GetBalancePlatform

`func (o *transferwebhookData) GetBalancePlatform() string`

GetBalancePlatform returns the BalancePlatform field if non-nil, zero value otherwise.

### GetBalancePlatformOk

`func (o *transferwebhookData) GetBalancePlatformOk() (*string, bool)`

GetBalancePlatformOk returns a tuple with the BalancePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancePlatform

`func (o *transferwebhookData) SetBalancePlatform(v string)`

SetBalancePlatform sets BalancePlatform field to given value.

### HasBalancePlatform

`func (o *transferwebhookData) HasBalancePlatform() bool`

HasBalancePlatform returns a boolean if a field has been set.

### GetBalances

`func (o *transferwebhookData) GetBalances() []BalanceMutation`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *transferwebhookData) GetBalancesOk() (*[]BalanceMutation, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *transferwebhookData) SetBalances(v []BalanceMutation)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *transferwebhookData) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetCategory

`func (o *transferwebhookData) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *transferwebhookData) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *transferwebhookData) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetCounterparty

`func (o *transferwebhookData) GetCounterparty() CounterpartyV3`

GetCounterparty returns the Counterparty field if non-nil, zero value otherwise.

### GetCounterpartyOk

`func (o *transferwebhookData) GetCounterpartyOk() (*CounterpartyV3, bool)`

GetCounterpartyOk returns a tuple with the Counterparty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterparty

`func (o *transferwebhookData) SetCounterparty(v CounterpartyV3)`

SetCounterparty sets Counterparty field to given value.

### HasCounterparty

`func (o *transferwebhookData) HasCounterparty() bool`

HasCounterparty returns a boolean if a field has been set.

### GetCreationDate

`func (o *transferwebhookData) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *transferwebhookData) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *transferwebhookData) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *transferwebhookData) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDescription

`func (o *transferwebhookData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *transferwebhookData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *transferwebhookData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *transferwebhookData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirection

`func (o *transferwebhookData) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *transferwebhookData) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *transferwebhookData) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *transferwebhookData) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetEvents

`func (o *transferwebhookData) GetEvents() []TransferEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *transferwebhookData) GetEventsOk() (*[]TransferEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *transferwebhookData) SetEvents(v []TransferEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *transferwebhookData) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetId

`func (o *transferwebhookData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *transferwebhookData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *transferwebhookData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *transferwebhookData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModificationMerchantReference

`func (o *transferwebhookData) GetModificationMerchantReference() string`

GetModificationMerchantReference returns the ModificationMerchantReference field if non-nil, zero value otherwise.

### GetModificationMerchantReferenceOk

`func (o *transferwebhookData) GetModificationMerchantReferenceOk() (*string, bool)`

GetModificationMerchantReferenceOk returns a tuple with the ModificationMerchantReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationMerchantReference

`func (o *transferwebhookData) SetModificationMerchantReference(v string)`

SetModificationMerchantReference sets ModificationMerchantReference field to given value.

### HasModificationMerchantReference

`func (o *transferwebhookData) HasModificationMerchantReference() bool`

HasModificationMerchantReference returns a boolean if a field has been set.

### GetModificationPspReference

`func (o *transferwebhookData) GetModificationPspReference() string`

GetModificationPspReference returns the ModificationPspReference field if non-nil, zero value otherwise.

### GetModificationPspReferenceOk

`func (o *transferwebhookData) GetModificationPspReferenceOk() (*string, bool)`

GetModificationPspReferenceOk returns a tuple with the ModificationPspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationPspReference

`func (o *transferwebhookData) SetModificationPspReference(v string)`

SetModificationPspReference sets ModificationPspReference field to given value.

### HasModificationPspReference

`func (o *transferwebhookData) HasModificationPspReference() bool`

HasModificationPspReference returns a boolean if a field has been set.

### GetPanEntryMode

`func (o *transferwebhookData) GetPanEntryMode() string`

GetPanEntryMode returns the PanEntryMode field if non-nil, zero value otherwise.

### GetPanEntryModeOk

`func (o *transferwebhookData) GetPanEntryModeOk() (*string, bool)`

GetPanEntryModeOk returns a tuple with the PanEntryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanEntryMode

`func (o *transferwebhookData) SetPanEntryMode(v string)`

SetPanEntryMode sets PanEntryMode field to given value.

### HasPanEntryMode

`func (o *transferwebhookData) HasPanEntryMode() bool`

HasPanEntryMode returns a boolean if a field has been set.

### GetPaymentInstrument

`func (o *transferwebhookData) GetPaymentInstrument() PaymentInstrument`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *transferwebhookData) GetPaymentInstrumentOk() (*PaymentInstrument, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *transferwebhookData) SetPaymentInstrument(v PaymentInstrument)`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *transferwebhookData) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### GetPaymentInstrumentId

`func (o *transferwebhookData) GetPaymentInstrumentId() string`

GetPaymentInstrumentId returns the PaymentInstrumentId field if non-nil, zero value otherwise.

### GetPaymentInstrumentIdOk

`func (o *transferwebhookData) GetPaymentInstrumentIdOk() (*string, bool)`

GetPaymentInstrumentIdOk returns a tuple with the PaymentInstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrumentId

`func (o *transferwebhookData) SetPaymentInstrumentId(v string)`

SetPaymentInstrumentId sets PaymentInstrumentId field to given value.

### HasPaymentInstrumentId

`func (o *transferwebhookData) HasPaymentInstrumentId() bool`

HasPaymentInstrumentId returns a boolean if a field has been set.

### GetPaymentMerchantReference

`func (o *transferwebhookData) GetPaymentMerchantReference() string`

GetPaymentMerchantReference returns the PaymentMerchantReference field if non-nil, zero value otherwise.

### GetPaymentMerchantReferenceOk

`func (o *transferwebhookData) GetPaymentMerchantReferenceOk() (*string, bool)`

GetPaymentMerchantReferenceOk returns a tuple with the PaymentMerchantReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMerchantReference

`func (o *transferwebhookData) SetPaymentMerchantReference(v string)`

SetPaymentMerchantReference sets PaymentMerchantReference field to given value.

### HasPaymentMerchantReference

`func (o *transferwebhookData) HasPaymentMerchantReference() bool`

HasPaymentMerchantReference returns a boolean if a field has been set.

### GetPriority

`func (o *transferwebhookData) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *transferwebhookData) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *transferwebhookData) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *transferwebhookData) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProcessingType

`func (o *transferwebhookData) GetProcessingType() string`

GetProcessingType returns the ProcessingType field if non-nil, zero value otherwise.

### GetProcessingTypeOk

`func (o *transferwebhookData) GetProcessingTypeOk() (*string, bool)`

GetProcessingTypeOk returns a tuple with the ProcessingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingType

`func (o *transferwebhookData) SetProcessingType(v string)`

SetProcessingType sets ProcessingType field to given value.

### HasProcessingType

`func (o *transferwebhookData) HasProcessingType() bool`

HasProcessingType returns a boolean if a field has been set.

### GetPspPaymentReference

`func (o *transferwebhookData) GetPspPaymentReference() string`

GetPspPaymentReference returns the PspPaymentReference field if non-nil, zero value otherwise.

### GetPspPaymentReferenceOk

`func (o *transferwebhookData) GetPspPaymentReferenceOk() (*string, bool)`

GetPspPaymentReferenceOk returns a tuple with the PspPaymentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspPaymentReference

`func (o *transferwebhookData) SetPspPaymentReference(v string)`

SetPspPaymentReference sets PspPaymentReference field to given value.

### HasPspPaymentReference

`func (o *transferwebhookData) HasPspPaymentReference() bool`

HasPspPaymentReference returns a boolean if a field has been set.

### GetReason

`func (o *transferwebhookData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *transferwebhookData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *transferwebhookData) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *transferwebhookData) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReference

`func (o *transferwebhookData) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *transferwebhookData) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *transferwebhookData) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *transferwebhookData) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceForBeneficiary

`func (o *transferwebhookData) GetReferenceForBeneficiary() string`

GetReferenceForBeneficiary returns the ReferenceForBeneficiary field if non-nil, zero value otherwise.

### GetReferenceForBeneficiaryOk

`func (o *transferwebhookData) GetReferenceForBeneficiaryOk() (*string, bool)`

GetReferenceForBeneficiaryOk returns a tuple with the ReferenceForBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceForBeneficiary

`func (o *transferwebhookData) SetReferenceForBeneficiary(v string)`

SetReferenceForBeneficiary sets ReferenceForBeneficiary field to given value.

### HasReferenceForBeneficiary

`func (o *transferwebhookData) HasReferenceForBeneficiary() bool`

HasReferenceForBeneficiary returns a boolean if a field has been set.

### GetRelayedAuthorisationData

`func (o *transferwebhookData) GetRelayedAuthorisationData() RelayedAuthorisationData`

GetRelayedAuthorisationData returns the RelayedAuthorisationData field if non-nil, zero value otherwise.

### GetRelayedAuthorisationDataOk

`func (o *transferwebhookData) GetRelayedAuthorisationDataOk() (*RelayedAuthorisationData, bool)`

GetRelayedAuthorisationDataOk returns a tuple with the RelayedAuthorisationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayedAuthorisationData

`func (o *transferwebhookData) SetRelayedAuthorisationData(v RelayedAuthorisationData)`

SetRelayedAuthorisationData sets RelayedAuthorisationData field to given value.

### HasRelayedAuthorisationData

`func (o *transferwebhookData) HasRelayedAuthorisationData() bool`

HasRelayedAuthorisationData returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *transferwebhookData) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *transferwebhookData) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *transferwebhookData) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *transferwebhookData) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetStatus

`func (o *transferwebhookData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *transferwebhookData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *transferwebhookData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTracking

`func (o *transferwebhookData) GetTracking() transferwebhookTransferTracking`

GetTracking returns the Tracking field if non-nil, zero value otherwise.

### GetTrackingOk

`func (o *transferwebhookData) GetTrackingOk() (*transferwebhookTransferTracking, bool)`

GetTrackingOk returns a tuple with the Tracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracking

`func (o *transferwebhookData) SetTracking(v transferwebhookTransferTracking)`

SetTracking sets Tracking field to given value.

### HasTracking

`func (o *transferwebhookData) HasTracking() bool`

HasTracking returns a boolean if a field has been set.

### GetTransactionId

`func (o *transferwebhookData) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *transferwebhookData) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *transferwebhookData) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *transferwebhookData) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionRulesResult

`func (o *transferwebhookData) GetTransactionRulesResult() TransactionRulesResult`

GetTransactionRulesResult returns the TransactionRulesResult field if non-nil, zero value otherwise.

### GetTransactionRulesResultOk

`func (o *transferwebhookData) GetTransactionRulesResultOk() (*TransactionRulesResult, bool)`

GetTransactionRulesResultOk returns a tuple with the TransactionRulesResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRulesResult

`func (o *transferwebhookData) SetTransactionRulesResult(v TransactionRulesResult)`

SetTransactionRulesResult sets TransactionRulesResult field to given value.

### HasTransactionRulesResult

`func (o *transferwebhookData) HasTransactionRulesResult() bool`

HasTransactionRulesResult returns a boolean if a field has been set.

### GetType

`func (o *transferwebhookData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *transferwebhookData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *transferwebhookData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *transferwebhookData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValidationFacts

`func (o *transferwebhookData) GetValidationFacts() []transferwebhookValidationFact`

GetValidationFacts returns the ValidationFacts field if non-nil, zero value otherwise.

### GetValidationFactsOk

`func (o *transferwebhookData) GetValidationFactsOk() (*[]transferwebhookValidationFact, bool)`

GetValidationFactsOk returns a tuple with the ValidationFacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationFacts

`func (o *transferwebhookData) SetValidationFacts(v []transferwebhookValidationFact)`

SetValidationFacts sets ValidationFacts field to given value.

### HasValidationFacts

`func (o *transferwebhookData) HasValidationFacts() bool`

HasValidationFacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


