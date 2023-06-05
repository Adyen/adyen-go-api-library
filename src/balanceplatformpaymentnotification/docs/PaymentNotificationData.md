# PaymentNotificationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolder** | Pointer to [**ResourceReference**](ResourceReference.md) |  | [optional] 
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**AuthCode** | Pointer to **string** | The authorisation code for the payment. | [optional] 
**BalanceAccount** | Pointer to [**ResourceReference**](ResourceReference.md) |  | [optional] 
**BalancePlatform** | Pointer to **string** | The unique identifier of the balance platform. | [optional] 
**CreationDate** | Pointer to **time.Time** | The date and time when the event was triggered, in ISO 8601 extended format. For example, **2020-12-18T10:15:30+01:00**. | [optional] 
**Description** | Pointer to **string** | Your description for the transfer. If you send a description longer than 140 characters, the text is truncated. | [optional] 
**Id** | Pointer to **string** | The ID of the resource. | [optional] 
**MerchantData** | Pointer to [**MerchantData**](MerchantData.md) |  | [optional] 
**Modification** | Pointer to [**NotificationModificationData**](NotificationModificationData.md) |  | [optional] 
**OriginalAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**PaymentInstrument** | Pointer to [**ResourceReference**](ResourceReference.md) |  | [optional] 
**PlatformPayment** | Pointer to [**PlatformPayment**](PlatformPayment.md) |  | [optional] 
**ProcessingType** | Pointer to **string** | Contains information about how the payment was processed. Possible values: **atmWithdraw**, **balanceInquiry**, **ecommerce**, **moto**, **pos**, **purchaseWithCashback**, **recurring**, **token**, **unknown**. | [optional] 
**Reference** | Pointer to **string** | The [&#x60;reference&#x60;](https://docs.adyen.com/api-explorer/#/transfers/latest/post/transfers__reqParam_reference) from the &#x60;/transfers&#x60; request. If you haven&#39;t provided any, Adyen generates a unique reference. | [optional] 
**ReferenceForBeneficiary** | Pointer to **string** | The reference sent to or received from the counterparty.  * For outgoing funds, this is the [&#x60;referenceForBeneficiary&#x60;](https://docs.adyen.com/api-explorer/#/transfers/latest/post/transfers__resParam_referenceForBeneficiary) from the  [&#x60;/transfers&#x60;](https://docs.adyen.com/api-explorer/#/transfers/latest/post/transfers) request.   * For incoming funds, this is the reference from the sender. | [optional] 
**RelayedAuthorisationData** | Pointer to [**RelayedAuthorisationData**](RelayedAuthorisationData.md) |  | [optional] 
**Status** | Pointer to **string** | The event status. The possible values depend on the &#x60;type&#x60;.  * **Authorised**, **Refused**, or **Error** for type &#x60;balancePlatform.payment.created&#x60;   * **Expired** or **Cancelled** or **AuthAdjustmentAuthorised** or **AuthAdjustmentRefused** for type &#x60;balancePlatform.payment.updated&#x60;  * **PendingIncomingTransfer** for type &#x60;balancePlatform.incomingTransfer.created&#x60;   * **Refunded** or **IncomingTransfer** for type &#x60;balancePlatform.incomingTransfer.updated&#x60;   * **Captured** or **OutgoingTransfer** for type &#x60;balancePlatform.outgoingTransfer.created&#x60;  * **TransferConfirmed**, **TransferSentOut**, or **TransferFailed** for type &#x60;balancePlatform.outgoingTransfer.updated&#x60;     | [optional] 
**TransactionRulesResult** | Pointer to [**TransactionRulesResult**](TransactionRulesResult.md) |  | [optional] 
**ValidationResult** | Pointer to [**[]ValidationResult**](ValidationResult.md) | Array of checks that Adyen performed to validate the payment and the result of each. | [optional] 

## Methods

### NewPaymentNotificationData

`func NewPaymentNotificationData() *PaymentNotificationData`

NewPaymentNotificationData instantiates a new PaymentNotificationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentNotificationDataWithDefaults

`func NewPaymentNotificationDataWithDefaults() *PaymentNotificationData`

NewPaymentNotificationDataWithDefaults instantiates a new PaymentNotificationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolder

`func (o *PaymentNotificationData) GetAccountHolder() ResourceReference`

GetAccountHolder returns the AccountHolder field if non-nil, zero value otherwise.

### GetAccountHolderOk

`func (o *PaymentNotificationData) GetAccountHolderOk() (*ResourceReference, bool)`

GetAccountHolderOk returns a tuple with the AccountHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolder

`func (o *PaymentNotificationData) SetAccountHolder(v ResourceReference)`

SetAccountHolder sets AccountHolder field to given value.

### HasAccountHolder

`func (o *PaymentNotificationData) HasAccountHolder() bool`

HasAccountHolder returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentNotificationData) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentNotificationData) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentNotificationData) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentNotificationData) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAuthCode

`func (o *PaymentNotificationData) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *PaymentNotificationData) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *PaymentNotificationData) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *PaymentNotificationData) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetBalanceAccount

`func (o *PaymentNotificationData) GetBalanceAccount() ResourceReference`

GetBalanceAccount returns the BalanceAccount field if non-nil, zero value otherwise.

### GetBalanceAccountOk

`func (o *PaymentNotificationData) GetBalanceAccountOk() (*ResourceReference, bool)`

GetBalanceAccountOk returns a tuple with the BalanceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccount

`func (o *PaymentNotificationData) SetBalanceAccount(v ResourceReference)`

SetBalanceAccount sets BalanceAccount field to given value.

### HasBalanceAccount

`func (o *PaymentNotificationData) HasBalanceAccount() bool`

HasBalanceAccount returns a boolean if a field has been set.

### GetBalancePlatform

`func (o *PaymentNotificationData) GetBalancePlatform() string`

GetBalancePlatform returns the BalancePlatform field if non-nil, zero value otherwise.

### GetBalancePlatformOk

`func (o *PaymentNotificationData) GetBalancePlatformOk() (*string, bool)`

GetBalancePlatformOk returns a tuple with the BalancePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancePlatform

`func (o *PaymentNotificationData) SetBalancePlatform(v string)`

SetBalancePlatform sets BalancePlatform field to given value.

### HasBalancePlatform

`func (o *PaymentNotificationData) HasBalancePlatform() bool`

HasBalancePlatform returns a boolean if a field has been set.

### GetCreationDate

`func (o *PaymentNotificationData) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *PaymentNotificationData) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *PaymentNotificationData) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *PaymentNotificationData) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentNotificationData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentNotificationData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentNotificationData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentNotificationData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *PaymentNotificationData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentNotificationData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentNotificationData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentNotificationData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantData

`func (o *PaymentNotificationData) GetMerchantData() MerchantData`

GetMerchantData returns the MerchantData field if non-nil, zero value otherwise.

### GetMerchantDataOk

`func (o *PaymentNotificationData) GetMerchantDataOk() (*MerchantData, bool)`

GetMerchantDataOk returns a tuple with the MerchantData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantData

`func (o *PaymentNotificationData) SetMerchantData(v MerchantData)`

SetMerchantData sets MerchantData field to given value.

### HasMerchantData

`func (o *PaymentNotificationData) HasMerchantData() bool`

HasMerchantData returns a boolean if a field has been set.

### GetModification

`func (o *PaymentNotificationData) GetModification() NotificationModificationData`

GetModification returns the Modification field if non-nil, zero value otherwise.

### GetModificationOk

`func (o *PaymentNotificationData) GetModificationOk() (*NotificationModificationData, bool)`

GetModificationOk returns a tuple with the Modification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModification

`func (o *PaymentNotificationData) SetModification(v NotificationModificationData)`

SetModification sets Modification field to given value.

### HasModification

`func (o *PaymentNotificationData) HasModification() bool`

HasModification returns a boolean if a field has been set.

### GetOriginalAmount

`func (o *PaymentNotificationData) GetOriginalAmount() Amount`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *PaymentNotificationData) GetOriginalAmountOk() (*Amount, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *PaymentNotificationData) SetOriginalAmount(v Amount)`

SetOriginalAmount sets OriginalAmount field to given value.

### HasOriginalAmount

`func (o *PaymentNotificationData) HasOriginalAmount() bool`

HasOriginalAmount returns a boolean if a field has been set.

### GetPaymentInstrument

`func (o *PaymentNotificationData) GetPaymentInstrument() ResourceReference`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *PaymentNotificationData) GetPaymentInstrumentOk() (*ResourceReference, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *PaymentNotificationData) SetPaymentInstrument(v ResourceReference)`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *PaymentNotificationData) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### GetPlatformPayment

`func (o *PaymentNotificationData) GetPlatformPayment() PlatformPayment`

GetPlatformPayment returns the PlatformPayment field if non-nil, zero value otherwise.

### GetPlatformPaymentOk

`func (o *PaymentNotificationData) GetPlatformPaymentOk() (*PlatformPayment, bool)`

GetPlatformPaymentOk returns a tuple with the PlatformPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPayment

`func (o *PaymentNotificationData) SetPlatformPayment(v PlatformPayment)`

SetPlatformPayment sets PlatformPayment field to given value.

### HasPlatformPayment

`func (o *PaymentNotificationData) HasPlatformPayment() bool`

HasPlatformPayment returns a boolean if a field has been set.

### GetProcessingType

`func (o *PaymentNotificationData) GetProcessingType() string`

GetProcessingType returns the ProcessingType field if non-nil, zero value otherwise.

### GetProcessingTypeOk

`func (o *PaymentNotificationData) GetProcessingTypeOk() (*string, bool)`

GetProcessingTypeOk returns a tuple with the ProcessingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingType

`func (o *PaymentNotificationData) SetProcessingType(v string)`

SetProcessingType sets ProcessingType field to given value.

### HasProcessingType

`func (o *PaymentNotificationData) HasProcessingType() bool`

HasProcessingType returns a boolean if a field has been set.

### GetReference

`func (o *PaymentNotificationData) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentNotificationData) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentNotificationData) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaymentNotificationData) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceForBeneficiary

`func (o *PaymentNotificationData) GetReferenceForBeneficiary() string`

GetReferenceForBeneficiary returns the ReferenceForBeneficiary field if non-nil, zero value otherwise.

### GetReferenceForBeneficiaryOk

`func (o *PaymentNotificationData) GetReferenceForBeneficiaryOk() (*string, bool)`

GetReferenceForBeneficiaryOk returns a tuple with the ReferenceForBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceForBeneficiary

`func (o *PaymentNotificationData) SetReferenceForBeneficiary(v string)`

SetReferenceForBeneficiary sets ReferenceForBeneficiary field to given value.

### HasReferenceForBeneficiary

`func (o *PaymentNotificationData) HasReferenceForBeneficiary() bool`

HasReferenceForBeneficiary returns a boolean if a field has been set.

### GetRelayedAuthorisationData

`func (o *PaymentNotificationData) GetRelayedAuthorisationData() RelayedAuthorisationData`

GetRelayedAuthorisationData returns the RelayedAuthorisationData field if non-nil, zero value otherwise.

### GetRelayedAuthorisationDataOk

`func (o *PaymentNotificationData) GetRelayedAuthorisationDataOk() (*RelayedAuthorisationData, bool)`

GetRelayedAuthorisationDataOk returns a tuple with the RelayedAuthorisationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayedAuthorisationData

`func (o *PaymentNotificationData) SetRelayedAuthorisationData(v RelayedAuthorisationData)`

SetRelayedAuthorisationData sets RelayedAuthorisationData field to given value.

### HasRelayedAuthorisationData

`func (o *PaymentNotificationData) HasRelayedAuthorisationData() bool`

HasRelayedAuthorisationData returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentNotificationData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentNotificationData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentNotificationData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentNotificationData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransactionRulesResult

`func (o *PaymentNotificationData) GetTransactionRulesResult() TransactionRulesResult`

GetTransactionRulesResult returns the TransactionRulesResult field if non-nil, zero value otherwise.

### GetTransactionRulesResultOk

`func (o *PaymentNotificationData) GetTransactionRulesResultOk() (*TransactionRulesResult, bool)`

GetTransactionRulesResultOk returns a tuple with the TransactionRulesResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRulesResult

`func (o *PaymentNotificationData) SetTransactionRulesResult(v TransactionRulesResult)`

SetTransactionRulesResult sets TransactionRulesResult field to given value.

### HasTransactionRulesResult

`func (o *PaymentNotificationData) HasTransactionRulesResult() bool`

HasTransactionRulesResult returns a boolean if a field has been set.

### GetValidationResult

`func (o *PaymentNotificationData) GetValidationResult() []ValidationResult`

GetValidationResult returns the ValidationResult field if non-nil, zero value otherwise.

### GetValidationResultOk

`func (o *PaymentNotificationData) GetValidationResultOk() (*[]ValidationResult, bool)`

GetValidationResultOk returns a tuple with the ValidationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationResult

`func (o *PaymentNotificationData) SetValidationResult(v []ValidationResult)`

SetValidationResult sets ValidationResult field to given value.

### HasValidationResult

`func (o *PaymentNotificationData) HasValidationResult() bool`

HasValidationResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


