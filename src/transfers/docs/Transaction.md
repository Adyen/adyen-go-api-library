# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolderId** | **string** | Unique identifier of the account holder. | 
**Amount** | [**Amount**](Amount.md) |  | 
**BalanceAccountId** | **string** | Unique identifier of the balance account. | 
**BalancePlatform** | **string** | Unique identifier of the balance platform. | 
**BookingDate** | **time.Time** | The date the transaction was booked to the balance account. | 
**Category** | Pointer to **string** | The category of the transaction indicating the type of activity.   Possible values:  * **platformPayment**: The transaction is a payment or payment modification made with an Adyen merchant account.  * **internal**: The transaction resulted from an internal adjustment such as a deposit correction or invoice deduction.  * **bank**: The transaction is a bank-related activity, such as sending a payout or receiving funds.  * **issuedCard**: The transaction is a card-related activity, such as using an Adyen-issued card to pay online.   | [optional] 
**Counterparty** | [**CounterpartyV3**](CounterpartyV3.md) |  | 
**CreatedAt** | **time.Time** | The date the transaction was created. | 
**Description** | Pointer to **string** | The &#x60;description&#x60; from the &#x60;/transfers&#x60; request. | [optional] 
**Id** | **string** | Unique identifier of the transaction. | 
**InstructedAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**PaymentInstrumentId** | Pointer to **string** | Unique identifier of the payment instrument that was used for the transaction. | [optional] 
**Reference** | **string** | The [&#x60;reference&#x60;](https://docs.adyen.com/api-explorer/#/transfers/latest/post/transfers__reqParam_reference) from the &#x60;/transfers&#x60; request. If you haven&#39;t provided any, Adyen generates a unique reference. | 
**ReferenceForBeneficiary** | Pointer to **string** | The reference sent to or received from the counterparty.  * For outgoing funds, this is the [&#x60;referenceForBeneficiary&#x60;](https://docs.adyen.com/api-explorer/#/transfers/latest/post/transfers__resParam_referenceForBeneficiary) from the  [&#x60;/transfers&#x60;](https://docs.adyen.com/api-explorer/#/transfers/latest/post/transfers__reqParam_referenceForBeneficiary) request.   * For incoming funds, this is the reference from the sender. | [optional] 
**Status** | **string** | The status of the transaction.   Possible values:  * **pending**: The transaction is still pending.  * **booked**: The transaction has been booked to the balance account.   | 
**TransferId** | Pointer to **string** | Unique identifier of the related transfer. | [optional] 
**Type** | Pointer to **string** | The type of the transaction.   Possible values: **payment**, **capture**, **captureReversal**, **refund** **refundReversal**, **chargeback**, **chargebackReversal**, **secondChargeback**, **atmWithdrawal**, **atmWithdrawalReversal**, **internalTransfer**, **manualCorrection**, **invoiceDeduction**, **depositCorrection**, **bankTransfer**, **miscCost**, **paymentCost**, **fee** | [optional] 
**ValueDate** | **time.Time** | The date the transfer amount becomes available in the balance account. | 

## Methods

### NewTransaction

`func NewTransaction(accountHolderId string, amount Amount, balanceAccountId string, balancePlatform string, bookingDate time.Time, counterparty CounterpartyV3, createdAt time.Time, id string, reference string, status string, valueDate time.Time, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolderId

`func (o *Transaction) GetAccountHolderId() string`

GetAccountHolderId returns the AccountHolderId field if non-nil, zero value otherwise.

### GetAccountHolderIdOk

`func (o *Transaction) GetAccountHolderIdOk() (*string, bool)`

GetAccountHolderIdOk returns a tuple with the AccountHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderId

`func (o *Transaction) SetAccountHolderId(v string)`

SetAccountHolderId sets AccountHolderId field to given value.


### GetAmount

`func (o *Transaction) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Transaction) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Transaction) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetBalanceAccountId

`func (o *Transaction) GetBalanceAccountId() string`

GetBalanceAccountId returns the BalanceAccountId field if non-nil, zero value otherwise.

### GetBalanceAccountIdOk

`func (o *Transaction) GetBalanceAccountIdOk() (*string, bool)`

GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccountId

`func (o *Transaction) SetBalanceAccountId(v string)`

SetBalanceAccountId sets BalanceAccountId field to given value.


### GetBalancePlatform

`func (o *Transaction) GetBalancePlatform() string`

GetBalancePlatform returns the BalancePlatform field if non-nil, zero value otherwise.

### GetBalancePlatformOk

`func (o *Transaction) GetBalancePlatformOk() (*string, bool)`

GetBalancePlatformOk returns a tuple with the BalancePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancePlatform

`func (o *Transaction) SetBalancePlatform(v string)`

SetBalancePlatform sets BalancePlatform field to given value.


### GetBookingDate

`func (o *Transaction) GetBookingDate() time.Time`

GetBookingDate returns the BookingDate field if non-nil, zero value otherwise.

### GetBookingDateOk

`func (o *Transaction) GetBookingDateOk() (*time.Time, bool)`

GetBookingDateOk returns a tuple with the BookingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingDate

`func (o *Transaction) SetBookingDate(v time.Time)`

SetBookingDate sets BookingDate field to given value.


### GetCategory

`func (o *Transaction) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Transaction) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Transaction) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Transaction) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCounterparty

`func (o *Transaction) GetCounterparty() CounterpartyV3`

GetCounterparty returns the Counterparty field if non-nil, zero value otherwise.

### GetCounterpartyOk

`func (o *Transaction) GetCounterpartyOk() (*CounterpartyV3, bool)`

GetCounterpartyOk returns a tuple with the Counterparty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterparty

`func (o *Transaction) SetCounterparty(v CounterpartyV3)`

SetCounterparty sets Counterparty field to given value.


### GetCreatedAt

`func (o *Transaction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Transaction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Transaction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *Transaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Transaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Transaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Transaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Transaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v string)`

SetId sets Id field to given value.


### GetInstructedAmount

`func (o *Transaction) GetInstructedAmount() Amount`

GetInstructedAmount returns the InstructedAmount field if non-nil, zero value otherwise.

### GetInstructedAmountOk

`func (o *Transaction) GetInstructedAmountOk() (*Amount, bool)`

GetInstructedAmountOk returns a tuple with the InstructedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructedAmount

`func (o *Transaction) SetInstructedAmount(v Amount)`

SetInstructedAmount sets InstructedAmount field to given value.

### HasInstructedAmount

`func (o *Transaction) HasInstructedAmount() bool`

HasInstructedAmount returns a boolean if a field has been set.

### GetPaymentInstrumentId

`func (o *Transaction) GetPaymentInstrumentId() string`

GetPaymentInstrumentId returns the PaymentInstrumentId field if non-nil, zero value otherwise.

### GetPaymentInstrumentIdOk

`func (o *Transaction) GetPaymentInstrumentIdOk() (*string, bool)`

GetPaymentInstrumentIdOk returns a tuple with the PaymentInstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrumentId

`func (o *Transaction) SetPaymentInstrumentId(v string)`

SetPaymentInstrumentId sets PaymentInstrumentId field to given value.

### HasPaymentInstrumentId

`func (o *Transaction) HasPaymentInstrumentId() bool`

HasPaymentInstrumentId returns a boolean if a field has been set.

### GetReference

`func (o *Transaction) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Transaction) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Transaction) SetReference(v string)`

SetReference sets Reference field to given value.


### GetReferenceForBeneficiary

`func (o *Transaction) GetReferenceForBeneficiary() string`

GetReferenceForBeneficiary returns the ReferenceForBeneficiary field if non-nil, zero value otherwise.

### GetReferenceForBeneficiaryOk

`func (o *Transaction) GetReferenceForBeneficiaryOk() (*string, bool)`

GetReferenceForBeneficiaryOk returns a tuple with the ReferenceForBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceForBeneficiary

`func (o *Transaction) SetReferenceForBeneficiary(v string)`

SetReferenceForBeneficiary sets ReferenceForBeneficiary field to given value.

### HasReferenceForBeneficiary

`func (o *Transaction) HasReferenceForBeneficiary() bool`

HasReferenceForBeneficiary returns a boolean if a field has been set.

### GetStatus

`func (o *Transaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTransferId

`func (o *Transaction) GetTransferId() string`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *Transaction) GetTransferIdOk() (*string, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *Transaction) SetTransferId(v string)`

SetTransferId sets TransferId field to given value.

### HasTransferId

`func (o *Transaction) HasTransferId() bool`

HasTransferId returns a boolean if a field has been set.

### GetType

`func (o *Transaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Transaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValueDate

`func (o *Transaction) GetValueDate() time.Time`

GetValueDate returns the ValueDate field if non-nil, zero value otherwise.

### GetValueDateOk

`func (o *Transaction) GetValueDateOk() (*time.Time, bool)`

GetValueDateOk returns a tuple with the ValueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDate

`func (o *Transaction) SetValueDate(v time.Time)`

SetValueDate sets ValueDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


