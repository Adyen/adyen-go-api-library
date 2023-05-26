# Transfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolder** | Pointer to [**ResourceReference**](ResourceReference.md) |  | [optional] 
**Amount** | [**Amount**](Amount.md) |  | 
**BalanceAccount** | Pointer to [**ResourceReference**](ResourceReference.md) |  | [optional] 
**BalanceAccountId** | Pointer to **string** | The unique identifier of the source [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id). | [optional] 
**Category** | **string** | The type of transfer.  Possible values:   - **bank**: Transfer to a [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id) or a bank account.  - **internal**: Transfer to another [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id) within your platform.  - **issuedCard**: Transfer initiated by a Adyen-issued card.  - **platformPayment**: Fund movements related to payments that are acquired for your users. | 
**Counterparty** | [**CounterpartyV3**](CounterpartyV3.md) |  | 
**Description** | Pointer to **string** | Your description for the transfer. It is used by most banks as the transfer description. We recommend sending a maximum of 140 characters, otherwise the description may be truncated.  Supported characters: **[a-z] [A-Z] [0-9] / - ?** **: ( ) . , &#39; + Space**  Supported characters for **regular** and **fast** transfers to a US counterparty: **[a-z] [A-Z] [0-9] &amp; $ % # @** **~ &#x3D; + - _ &#39; \&quot; ! ?** | [optional] 
**Direction** | Pointer to **string** | The direction of the transfer.  Possible values: **incoming**, **outgoing**. | [optional] 
**Id** | Pointer to **string** | The ID of the resource. | [optional] 
**PaymentInstrument** | Pointer to [**PaymentInstrument**](PaymentInstrument.md) |  | [optional] 
**PaymentInstrumentId** | Pointer to **string** | The unique identifier of the [payment instrument](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id) used in the transfer. | [optional] 
**Priority** | Pointer to **string** | The priority for the bank transfer. This sets the speed at which the transfer is sent and the fees that you have to pay. Required for transfers with &#x60;category&#x60; **bank**.  Possible values:  * **regular**: For normal, low-value transactions.  * **fast**: Faster way to transfer funds but has higher fees. Recommended for high-priority, low-value transactions.  * **wire**: Fastest way to transfer funds but has the highest fees. Recommended for high-priority, high-value transactions.  * **instant**: Instant way to transfer funds in [SEPA countries](https://www.ecb.europa.eu/paym/integration/retail/sepa/html/index.en.html).  * **crossBorder**: High-value transfer to a recipient in a different country.  * **internal**: Transfer to an Adyen-issued business bank account (by bank account number/IBAN). | [optional] 
**Reason** | Pointer to **string** | Additional information about the status of the transfer. | [optional] 
**Reference** | Pointer to **string** | Your reference for the transfer, used internally within your platform. If you don&#39;t provide this in the request, Adyen generates a unique reference. | [optional] 
**ReferenceForBeneficiary** | Pointer to **string** |  A reference that is sent to the recipient. This reference is also sent in all webhooks related to the transfer, so you can use it to track statuses for both the source and recipient of funds.   Supported characters: **a-z**, **A-Z**, **0-9**. The maximum length depends on the &#x60;category&#x60;.  - **internal**: 80 characters  - **bank**: 35 characters when transferring to an IBAN, 15 characters for others. | [optional] 
**Status** | **string** | The result of the transfer.   For example, **authorised**, **refused**, or **error**. | 

## Methods

### NewTransfer

`func NewTransfer(amount Amount, category string, counterparty CounterpartyV3, status string, ) *Transfer`

NewTransfer instantiates a new Transfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferWithDefaults

`func NewTransferWithDefaults() *Transfer`

NewTransferWithDefaults instantiates a new Transfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolder

`func (o *Transfer) GetAccountHolder() ResourceReference`

GetAccountHolder returns the AccountHolder field if non-nil, zero value otherwise.

### GetAccountHolderOk

`func (o *Transfer) GetAccountHolderOk() (*ResourceReference, bool)`

GetAccountHolderOk returns a tuple with the AccountHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolder

`func (o *Transfer) SetAccountHolder(v ResourceReference)`

SetAccountHolder sets AccountHolder field to given value.

### HasAccountHolder

`func (o *Transfer) HasAccountHolder() bool`

HasAccountHolder returns a boolean if a field has been set.

### GetAmount

`func (o *Transfer) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Transfer) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Transfer) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetBalanceAccount

`func (o *Transfer) GetBalanceAccount() ResourceReference`

GetBalanceAccount returns the BalanceAccount field if non-nil, zero value otherwise.

### GetBalanceAccountOk

`func (o *Transfer) GetBalanceAccountOk() (*ResourceReference, bool)`

GetBalanceAccountOk returns a tuple with the BalanceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccount

`func (o *Transfer) SetBalanceAccount(v ResourceReference)`

SetBalanceAccount sets BalanceAccount field to given value.

### HasBalanceAccount

`func (o *Transfer) HasBalanceAccount() bool`

HasBalanceAccount returns a boolean if a field has been set.

### GetBalanceAccountId

`func (o *Transfer) GetBalanceAccountId() string`

GetBalanceAccountId returns the BalanceAccountId field if non-nil, zero value otherwise.

### GetBalanceAccountIdOk

`func (o *Transfer) GetBalanceAccountIdOk() (*string, bool)`

GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccountId

`func (o *Transfer) SetBalanceAccountId(v string)`

SetBalanceAccountId sets BalanceAccountId field to given value.

### HasBalanceAccountId

`func (o *Transfer) HasBalanceAccountId() bool`

HasBalanceAccountId returns a boolean if a field has been set.

### GetCategory

`func (o *Transfer) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Transfer) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Transfer) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetCounterparty

`func (o *Transfer) GetCounterparty() CounterpartyV3`

GetCounterparty returns the Counterparty field if non-nil, zero value otherwise.

### GetCounterpartyOk

`func (o *Transfer) GetCounterpartyOk() (*CounterpartyV3, bool)`

GetCounterpartyOk returns a tuple with the Counterparty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterparty

`func (o *Transfer) SetCounterparty(v CounterpartyV3)`

SetCounterparty sets Counterparty field to given value.


### GetDescription

`func (o *Transfer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Transfer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Transfer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Transfer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirection

`func (o *Transfer) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Transfer) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Transfer) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *Transfer) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *Transfer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transfer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transfer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Transfer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPaymentInstrument

`func (o *Transfer) GetPaymentInstrument() PaymentInstrument`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *Transfer) GetPaymentInstrumentOk() (*PaymentInstrument, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *Transfer) SetPaymentInstrument(v PaymentInstrument)`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *Transfer) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### GetPaymentInstrumentId

`func (o *Transfer) GetPaymentInstrumentId() string`

GetPaymentInstrumentId returns the PaymentInstrumentId field if non-nil, zero value otherwise.

### GetPaymentInstrumentIdOk

`func (o *Transfer) GetPaymentInstrumentIdOk() (*string, bool)`

GetPaymentInstrumentIdOk returns a tuple with the PaymentInstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrumentId

`func (o *Transfer) SetPaymentInstrumentId(v string)`

SetPaymentInstrumentId sets PaymentInstrumentId field to given value.

### HasPaymentInstrumentId

`func (o *Transfer) HasPaymentInstrumentId() bool`

HasPaymentInstrumentId returns a boolean if a field has been set.

### GetPriority

`func (o *Transfer) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Transfer) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Transfer) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Transfer) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetReason

`func (o *Transfer) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Transfer) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Transfer) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Transfer) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReference

`func (o *Transfer) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Transfer) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Transfer) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *Transfer) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceForBeneficiary

`func (o *Transfer) GetReferenceForBeneficiary() string`

GetReferenceForBeneficiary returns the ReferenceForBeneficiary field if non-nil, zero value otherwise.

### GetReferenceForBeneficiaryOk

`func (o *Transfer) GetReferenceForBeneficiaryOk() (*string, bool)`

GetReferenceForBeneficiaryOk returns a tuple with the ReferenceForBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceForBeneficiary

`func (o *Transfer) SetReferenceForBeneficiary(v string)`

SetReferenceForBeneficiary sets ReferenceForBeneficiary field to given value.

### HasReferenceForBeneficiary

`func (o *Transfer) HasReferenceForBeneficiary() bool`

HasReferenceForBeneficiary returns a boolean if a field has been set.

### GetStatus

`func (o *Transfer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transfer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transfer) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


