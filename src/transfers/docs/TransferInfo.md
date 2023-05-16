# TransferInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**Amount**](Amount.md) |  | 
**BalanceAccountId** | Pointer to **string** | The unique identifier of the source [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id). | [optional] 
**Category** | **string** | The type of transfer.  Possible values:   - **bank**: Transfer to a [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id) or a bank account.  - **internal**: Transfer to another [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id) within your platform.  - **issuedCard**: Transfer initiated by a Adyen-issued card.  - **platformPayment**: Fund movements related to payments that are acquired for your users. | 
**Counterparty** | [**CounterpartyInfoV3**](CounterpartyInfoV3.md) |  | 
**Description** | Pointer to **string** | Your description for the transfer. It is used by most banks as the transfer description. We recommend sending a maximum of 140 characters, otherwise the description may be truncated.  Supported characters: **[a-z] [A-Z] [0-9] / - ?** **: ( ) . , &#39; + Space**  Supported characters for **regular** and **fast** transfers to a US counterparty: **[a-z] [A-Z] [0-9] &amp; $ % # @** **~ &#x3D; + - _ &#39; \&quot; ! ?** | [optional] 
**Id** | Pointer to **string** | The ID of the resource. | [optional] 
**PaymentInstrumentId** | Pointer to **string** | The unique identifier of the source [payment instrument](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/paymentInstruments__resParam_id). | [optional] 
**Priority** | Pointer to **string** | The priority for the bank transfer. This sets the speed at which the transfer is sent and the fees that you have to pay. Required for transfers with &#x60;category&#x60; **bank**.  Possible values:  * **regular**: For normal, low-value transactions.  * **fast**: Faster way to transfer funds but has higher fees. Recommended for high-priority, low-value transactions.  * **wire**: Fastest way to transfer funds but has the highest fees. Recommended for high-priority, high-value transactions.  * **instant**: Instant way to transfer funds in [SEPA countries](https://www.ecb.europa.eu/paym/integration/retail/sepa/html/index.en.html).  * **crossBorder**: High-value transfer to a recipient in a different country.  * **internal**: Transfer to an Adyen-issued business bank account (by bank account number/IBAN). | [optional] 
**Reference** | Pointer to **string** | Your reference for the transfer, used internally within your platform. If you don&#39;t provide this in the request, Adyen generates a unique reference. | [optional] 
**ReferenceForBeneficiary** | Pointer to **string** |  A reference that is sent to the recipient. This reference is also sent in all notification webhooks related to the transfer, so you can use it to track statuses for both the source and recipient of funds.   Supported characters: **a-z**, **A-Z**, **0-9**. The maximum length depends on the &#x60;category&#x60;.  - **internal**: 80 characters  - **bank**: 35 characters when transferring to an IBAN, 15 characters for others. | [optional] 
**UltimateParty** | Pointer to [**UltimatePartyIdentification**](UltimatePartyIdentification.md) |  | [optional] 

## Methods

### NewTransferInfo

`func NewTransferInfo(amount Amount, category string, counterparty CounterpartyInfoV3, ) *TransferInfo`

NewTransferInfo instantiates a new TransferInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferInfoWithDefaults

`func NewTransferInfoWithDefaults() *TransferInfo`

NewTransferInfoWithDefaults instantiates a new TransferInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransferInfo) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferInfo) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferInfo) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetBalanceAccountId

`func (o *TransferInfo) GetBalanceAccountId() string`

GetBalanceAccountId returns the BalanceAccountId field if non-nil, zero value otherwise.

### GetBalanceAccountIdOk

`func (o *TransferInfo) GetBalanceAccountIdOk() (*string, bool)`

GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccountId

`func (o *TransferInfo) SetBalanceAccountId(v string)`

SetBalanceAccountId sets BalanceAccountId field to given value.

### HasBalanceAccountId

`func (o *TransferInfo) HasBalanceAccountId() bool`

HasBalanceAccountId returns a boolean if a field has been set.

### GetCategory

`func (o *TransferInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TransferInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TransferInfo) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetCounterparty

`func (o *TransferInfo) GetCounterparty() CounterpartyInfoV3`

GetCounterparty returns the Counterparty field if non-nil, zero value otherwise.

### GetCounterpartyOk

`func (o *TransferInfo) GetCounterpartyOk() (*CounterpartyInfoV3, bool)`

GetCounterpartyOk returns a tuple with the Counterparty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterparty

`func (o *TransferInfo) SetCounterparty(v CounterpartyInfoV3)`

SetCounterparty sets Counterparty field to given value.


### GetDescription

`func (o *TransferInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransferInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransferInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransferInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *TransferInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransferInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPaymentInstrumentId

`func (o *TransferInfo) GetPaymentInstrumentId() string`

GetPaymentInstrumentId returns the PaymentInstrumentId field if non-nil, zero value otherwise.

### GetPaymentInstrumentIdOk

`func (o *TransferInfo) GetPaymentInstrumentIdOk() (*string, bool)`

GetPaymentInstrumentIdOk returns a tuple with the PaymentInstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrumentId

`func (o *TransferInfo) SetPaymentInstrumentId(v string)`

SetPaymentInstrumentId sets PaymentInstrumentId field to given value.

### HasPaymentInstrumentId

`func (o *TransferInfo) HasPaymentInstrumentId() bool`

HasPaymentInstrumentId returns a boolean if a field has been set.

### GetPriority

`func (o *TransferInfo) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TransferInfo) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TransferInfo) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TransferInfo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetReference

`func (o *TransferInfo) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *TransferInfo) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *TransferInfo) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *TransferInfo) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceForBeneficiary

`func (o *TransferInfo) GetReferenceForBeneficiary() string`

GetReferenceForBeneficiary returns the ReferenceForBeneficiary field if non-nil, zero value otherwise.

### GetReferenceForBeneficiaryOk

`func (o *TransferInfo) GetReferenceForBeneficiaryOk() (*string, bool)`

GetReferenceForBeneficiaryOk returns a tuple with the ReferenceForBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceForBeneficiary

`func (o *TransferInfo) SetReferenceForBeneficiary(v string)`

SetReferenceForBeneficiary sets ReferenceForBeneficiary field to given value.

### HasReferenceForBeneficiary

`func (o *TransferInfo) HasReferenceForBeneficiary() bool`

HasReferenceForBeneficiary returns a boolean if a field has been set.

### GetUltimateParty

`func (o *TransferInfo) GetUltimateParty() UltimatePartyIdentification`

GetUltimateParty returns the UltimateParty field if non-nil, zero value otherwise.

### GetUltimatePartyOk

`func (o *TransferInfo) GetUltimatePartyOk() (*UltimatePartyIdentification, bool)`

GetUltimatePartyOk returns a tuple with the UltimateParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUltimateParty

`func (o *TransferInfo) SetUltimateParty(v UltimatePartyIdentification)`

SetUltimateParty sets UltimateParty field to given value.

### HasUltimateParty

`func (o *TransferInfo) HasUltimateParty() bool`

HasUltimateParty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


