/*
 * Adyen for Platforms: Notifications
 *
 * The Notification API sends notifications to the endpoints specified in a given subscription. Subscriptions are managed through the Notification Configuration API. The API specifications listed here detail the format of each notification.  For more information, refer to our [documentation](https://docs.adyen.com/platforms/notifications).
 *
 * API version: 6
 * Contact: support@adyen.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package platformsnotificationevents

import (
	"encoding/json"
	"time"
)

// Transaction struct for Transaction
type Transaction struct {
	Amount            *Amount            `json:"amount,omitempty"`
	BankAccountDetail *BankAccountDetail `json:"bankAccountDetail,omitempty"`
	// The merchant reference of a related capture.
	CaptureMerchantReference *string `json:"captureMerchantReference,omitempty"`
	// The psp reference of a related capture.
	CapturePspReference *string `json:"capturePspReference,omitempty"`
	// The date on which the transaction was performed.
	CreationDate *time.Time `json:"creationDate,omitempty"`
	// A description of the transaction.
	Description *string `json:"description,omitempty"`
	// The code of the account to which funds were credited during an outgoing fund transfer.
	DestinationAccountCode *string `json:"destinationAccountCode,omitempty"`
	// The psp reference of the related dispute.
	DisputePspReference *string `json:"disputePspReference,omitempty"`
	// The reason code of a dispute.
	DisputeReasonCode *string `json:"disputeReasonCode,omitempty"`
	// The merchant reference of a transaction.
	MerchantReference *string `json:"merchantReference,omitempty"`
	// The psp reference of the related authorisation or transfer.
	PaymentPspReference *string `json:"paymentPspReference,omitempty"`
	// The psp reference of the related payout.
	PayoutPspReference *string `json:"payoutPspReference,omitempty"`
	// The psp reference of a transaction.
	PspReference *string `json:"pspReference,omitempty"`
	// The code of the account from which funds were debited during an incoming fund transfer.
	SourceAccountCode *string `json:"sourceAccountCode,omitempty"`
	// The status of the transaction. >Permitted values: `PendingCredit`, `CreditFailed`, `CreditClosed`, `CreditSuspended`, `Credited`, `Converted`, `PendingDebit`, `DebitFailed`, `Debited`, `DebitReversedReceived`, `DebitedReversed`, `ChargebackReceived`, `Chargeback`, `ChargebackReversedReceived`, `ChargebackReversed`, `Payout`, `PayoutReversed`, `FundTransfer`, `PendingFundTransfer`, `ManualCorrected`.
	TransactionStatus *string `json:"transactionStatus,omitempty"`
	// The transfer code of the transaction.
	TransferCode *string `json:"transferCode,omitempty"`
}

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction() *Transaction {
	this := Transaction{}
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Transaction) GetAmount() Amount {
	if o == nil || o.Amount == nil {
		var ret Amount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetAmountOk() (*Amount, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Transaction) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Amount and assigns it to the Amount field.
func (o *Transaction) SetAmount(v Amount) {
	o.Amount = &v
}

// GetBankAccountDetail returns the BankAccountDetail field value if set, zero value otherwise.
func (o *Transaction) GetBankAccountDetail() BankAccountDetail {
	if o == nil || o.BankAccountDetail == nil {
		var ret BankAccountDetail
		return ret
	}
	return *o.BankAccountDetail
}

// GetBankAccountDetailOk returns a tuple with the BankAccountDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetBankAccountDetailOk() (*BankAccountDetail, bool) {
	if o == nil || o.BankAccountDetail == nil {
		return nil, false
	}
	return o.BankAccountDetail, true
}

// HasBankAccountDetail returns a boolean if a field has been set.
func (o *Transaction) HasBankAccountDetail() bool {
	if o != nil && o.BankAccountDetail != nil {
		return true
	}

	return false
}

// SetBankAccountDetail gets a reference to the given BankAccountDetail and assigns it to the BankAccountDetail field.
func (o *Transaction) SetBankAccountDetail(v BankAccountDetail) {
	o.BankAccountDetail = &v
}

// GetCaptureMerchantReference returns the CaptureMerchantReference field value if set, zero value otherwise.
func (o *Transaction) GetCaptureMerchantReference() string {
	if o == nil || o.CaptureMerchantReference == nil {
		var ret string
		return ret
	}
	return *o.CaptureMerchantReference
}

// GetCaptureMerchantReferenceOk returns a tuple with the CaptureMerchantReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetCaptureMerchantReferenceOk() (*string, bool) {
	if o == nil || o.CaptureMerchantReference == nil {
		return nil, false
	}
	return o.CaptureMerchantReference, true
}

// HasCaptureMerchantReference returns a boolean if a field has been set.
func (o *Transaction) HasCaptureMerchantReference() bool {
	if o != nil && o.CaptureMerchantReference != nil {
		return true
	}

	return false
}

// SetCaptureMerchantReference gets a reference to the given string and assigns it to the CaptureMerchantReference field.
func (o *Transaction) SetCaptureMerchantReference(v string) {
	o.CaptureMerchantReference = &v
}

// GetCapturePspReference returns the CapturePspReference field value if set, zero value otherwise.
func (o *Transaction) GetCapturePspReference() string {
	if o == nil || o.CapturePspReference == nil {
		var ret string
		return ret
	}
	return *o.CapturePspReference
}

// GetCapturePspReferenceOk returns a tuple with the CapturePspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetCapturePspReferenceOk() (*string, bool) {
	if o == nil || o.CapturePspReference == nil {
		return nil, false
	}
	return o.CapturePspReference, true
}

// HasCapturePspReference returns a boolean if a field has been set.
func (o *Transaction) HasCapturePspReference() bool {
	if o != nil && o.CapturePspReference != nil {
		return true
	}

	return false
}

// SetCapturePspReference gets a reference to the given string and assigns it to the CapturePspReference field.
func (o *Transaction) SetCapturePspReference(v string) {
	o.CapturePspReference = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *Transaction) GetCreationDate() time.Time {
	if o == nil || o.CreationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *Transaction) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *Transaction) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Transaction) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Transaction) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Transaction) SetDescription(v string) {
	o.Description = &v
}

// GetDestinationAccountCode returns the DestinationAccountCode field value if set, zero value otherwise.
func (o *Transaction) GetDestinationAccountCode() string {
	if o == nil || o.DestinationAccountCode == nil {
		var ret string
		return ret
	}
	return *o.DestinationAccountCode
}

// GetDestinationAccountCodeOk returns a tuple with the DestinationAccountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetDestinationAccountCodeOk() (*string, bool) {
	if o == nil || o.DestinationAccountCode == nil {
		return nil, false
	}
	return o.DestinationAccountCode, true
}

// HasDestinationAccountCode returns a boolean if a field has been set.
func (o *Transaction) HasDestinationAccountCode() bool {
	if o != nil && o.DestinationAccountCode != nil {
		return true
	}

	return false
}

// SetDestinationAccountCode gets a reference to the given string and assigns it to the DestinationAccountCode field.
func (o *Transaction) SetDestinationAccountCode(v string) {
	o.DestinationAccountCode = &v
}

// GetDisputePspReference returns the DisputePspReference field value if set, zero value otherwise.
func (o *Transaction) GetDisputePspReference() string {
	if o == nil || o.DisputePspReference == nil {
		var ret string
		return ret
	}
	return *o.DisputePspReference
}

// GetDisputePspReferenceOk returns a tuple with the DisputePspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetDisputePspReferenceOk() (*string, bool) {
	if o == nil || o.DisputePspReference == nil {
		return nil, false
	}
	return o.DisputePspReference, true
}

// HasDisputePspReference returns a boolean if a field has been set.
func (o *Transaction) HasDisputePspReference() bool {
	if o != nil && o.DisputePspReference != nil {
		return true
	}

	return false
}

// SetDisputePspReference gets a reference to the given string and assigns it to the DisputePspReference field.
func (o *Transaction) SetDisputePspReference(v string) {
	o.DisputePspReference = &v
}

// GetDisputeReasonCode returns the DisputeReasonCode field value if set, zero value otherwise.
func (o *Transaction) GetDisputeReasonCode() string {
	if o == nil || o.DisputeReasonCode == nil {
		var ret string
		return ret
	}
	return *o.DisputeReasonCode
}

// GetDisputeReasonCodeOk returns a tuple with the DisputeReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetDisputeReasonCodeOk() (*string, bool) {
	if o == nil || o.DisputeReasonCode == nil {
		return nil, false
	}
	return o.DisputeReasonCode, true
}

// HasDisputeReasonCode returns a boolean if a field has been set.
func (o *Transaction) HasDisputeReasonCode() bool {
	if o != nil && o.DisputeReasonCode != nil {
		return true
	}

	return false
}

// SetDisputeReasonCode gets a reference to the given string and assigns it to the DisputeReasonCode field.
func (o *Transaction) SetDisputeReasonCode(v string) {
	o.DisputeReasonCode = &v
}

// GetMerchantReference returns the MerchantReference field value if set, zero value otherwise.
func (o *Transaction) GetMerchantReference() string {
	if o == nil || o.MerchantReference == nil {
		var ret string
		return ret
	}
	return *o.MerchantReference
}

// GetMerchantReferenceOk returns a tuple with the MerchantReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetMerchantReferenceOk() (*string, bool) {
	if o == nil || o.MerchantReference == nil {
		return nil, false
	}
	return o.MerchantReference, true
}

// HasMerchantReference returns a boolean if a field has been set.
func (o *Transaction) HasMerchantReference() bool {
	if o != nil && o.MerchantReference != nil {
		return true
	}

	return false
}

// SetMerchantReference gets a reference to the given string and assigns it to the MerchantReference field.
func (o *Transaction) SetMerchantReference(v string) {
	o.MerchantReference = &v
}

// GetPaymentPspReference returns the PaymentPspReference field value if set, zero value otherwise.
func (o *Transaction) GetPaymentPspReference() string {
	if o == nil || o.PaymentPspReference == nil {
		var ret string
		return ret
	}
	return *o.PaymentPspReference
}

// GetPaymentPspReferenceOk returns a tuple with the PaymentPspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetPaymentPspReferenceOk() (*string, bool) {
	if o == nil || o.PaymentPspReference == nil {
		return nil, false
	}
	return o.PaymentPspReference, true
}

// HasPaymentPspReference returns a boolean if a field has been set.
func (o *Transaction) HasPaymentPspReference() bool {
	if o != nil && o.PaymentPspReference != nil {
		return true
	}

	return false
}

// SetPaymentPspReference gets a reference to the given string and assigns it to the PaymentPspReference field.
func (o *Transaction) SetPaymentPspReference(v string) {
	o.PaymentPspReference = &v
}

// GetPayoutPspReference returns the PayoutPspReference field value if set, zero value otherwise.
func (o *Transaction) GetPayoutPspReference() string {
	if o == nil || o.PayoutPspReference == nil {
		var ret string
		return ret
	}
	return *o.PayoutPspReference
}

// GetPayoutPspReferenceOk returns a tuple with the PayoutPspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetPayoutPspReferenceOk() (*string, bool) {
	if o == nil || o.PayoutPspReference == nil {
		return nil, false
	}
	return o.PayoutPspReference, true
}

// HasPayoutPspReference returns a boolean if a field has been set.
func (o *Transaction) HasPayoutPspReference() bool {
	if o != nil && o.PayoutPspReference != nil {
		return true
	}

	return false
}

// SetPayoutPspReference gets a reference to the given string and assigns it to the PayoutPspReference field.
func (o *Transaction) SetPayoutPspReference(v string) {
	o.PayoutPspReference = &v
}

// GetPspReference returns the PspReference field value if set, zero value otherwise.
func (o *Transaction) GetPspReference() string {
	if o == nil || o.PspReference == nil {
		var ret string
		return ret
	}
	return *o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetPspReferenceOk() (*string, bool) {
	if o == nil || o.PspReference == nil {
		return nil, false
	}
	return o.PspReference, true
}

// HasPspReference returns a boolean if a field has been set.
func (o *Transaction) HasPspReference() bool {
	if o != nil && o.PspReference != nil {
		return true
	}

	return false
}

// SetPspReference gets a reference to the given string and assigns it to the PspReference field.
func (o *Transaction) SetPspReference(v string) {
	o.PspReference = &v
}

// GetSourceAccountCode returns the SourceAccountCode field value if set, zero value otherwise.
func (o *Transaction) GetSourceAccountCode() string {
	if o == nil || o.SourceAccountCode == nil {
		var ret string
		return ret
	}
	return *o.SourceAccountCode
}

// GetSourceAccountCodeOk returns a tuple with the SourceAccountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetSourceAccountCodeOk() (*string, bool) {
	if o == nil || o.SourceAccountCode == nil {
		return nil, false
	}
	return o.SourceAccountCode, true
}

// HasSourceAccountCode returns a boolean if a field has been set.
func (o *Transaction) HasSourceAccountCode() bool {
	if o != nil && o.SourceAccountCode != nil {
		return true
	}

	return false
}

// SetSourceAccountCode gets a reference to the given string and assigns it to the SourceAccountCode field.
func (o *Transaction) SetSourceAccountCode(v string) {
	o.SourceAccountCode = &v
}

// GetTransactionStatus returns the TransactionStatus field value if set, zero value otherwise.
func (o *Transaction) GetTransactionStatus() string {
	if o == nil || o.TransactionStatus == nil {
		var ret string
		return ret
	}
	return *o.TransactionStatus
}

// GetTransactionStatusOk returns a tuple with the TransactionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTransactionStatusOk() (*string, bool) {
	if o == nil || o.TransactionStatus == nil {
		return nil, false
	}
	return o.TransactionStatus, true
}

// HasTransactionStatus returns a boolean if a field has been set.
func (o *Transaction) HasTransactionStatus() bool {
	if o != nil && o.TransactionStatus != nil {
		return true
	}

	return false
}

// SetTransactionStatus gets a reference to the given string and assigns it to the TransactionStatus field.
func (o *Transaction) SetTransactionStatus(v string) {
	o.TransactionStatus = &v
}

// GetTransferCode returns the TransferCode field value if set, zero value otherwise.
func (o *Transaction) GetTransferCode() string {
	if o == nil || o.TransferCode == nil {
		var ret string
		return ret
	}
	return *o.TransferCode
}

// GetTransferCodeOk returns a tuple with the TransferCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTransferCodeOk() (*string, bool) {
	if o == nil || o.TransferCode == nil {
		return nil, false
	}
	return o.TransferCode, true
}

// HasTransferCode returns a boolean if a field has been set.
func (o *Transaction) HasTransferCode() bool {
	if o != nil && o.TransferCode != nil {
		return true
	}

	return false
}

// SetTransferCode gets a reference to the given string and assigns it to the TransferCode field.
func (o *Transaction) SetTransferCode(v string) {
	o.TransferCode = &v
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.BankAccountDetail != nil {
		toSerialize["bankAccountDetail"] = o.BankAccountDetail
	}
	if o.CaptureMerchantReference != nil {
		toSerialize["captureMerchantReference"] = o.CaptureMerchantReference
	}
	if o.CapturePspReference != nil {
		toSerialize["capturePspReference"] = o.CapturePspReference
	}
	if o.CreationDate != nil {
		toSerialize["creationDate"] = o.CreationDate
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DestinationAccountCode != nil {
		toSerialize["destinationAccountCode"] = o.DestinationAccountCode
	}
	if o.DisputePspReference != nil {
		toSerialize["disputePspReference"] = o.DisputePspReference
	}
	if o.DisputeReasonCode != nil {
		toSerialize["disputeReasonCode"] = o.DisputeReasonCode
	}
	if o.MerchantReference != nil {
		toSerialize["merchantReference"] = o.MerchantReference
	}
	if o.PaymentPspReference != nil {
		toSerialize["paymentPspReference"] = o.PaymentPspReference
	}
	if o.PayoutPspReference != nil {
		toSerialize["payoutPspReference"] = o.PayoutPspReference
	}
	if o.PspReference != nil {
		toSerialize["pspReference"] = o.PspReference
	}
	if o.SourceAccountCode != nil {
		toSerialize["sourceAccountCode"] = o.SourceAccountCode
	}
	if o.TransactionStatus != nil {
		toSerialize["transactionStatus"] = o.TransactionStatus
	}
	if o.TransferCode != nil {
		toSerialize["transferCode"] = o.TransferCode
	}
	return json.Marshal(toSerialize)
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
