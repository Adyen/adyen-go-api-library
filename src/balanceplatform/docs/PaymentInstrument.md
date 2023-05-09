# PaymentInstrument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceAccountId** | **string** | The unique identifier of the [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/v1/post/balanceAccounts__resParam_id) associated with the payment instrument. | 
**BankAccount** | Pointer to [**PaymentInstrumentBankAccount**](PaymentInstrumentBankAccount.md) |  | [optional] 
**Card** | Pointer to [**Card**](Card.md) |  | [optional] 
**Description** | Pointer to **string** | Your description for the payment instrument, maximum 300 characters. | [optional] 
**Id** | **string** | The unique identifier of the payment instrument. | 
**IssuingCountryCode** | **string** | The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code where the payment instrument is issued. For example, **NL** or **US**. | 
**PaymentInstrumentGroupId** | Pointer to **string** | The unique identifier of the [payment instrument group](https://docs.adyen.com/api-explorer/#/balanceplatform/v1/post/paymentInstrumentGroups__resParam_id) to which the payment instrument belongs. | [optional] 
**Reference** | Pointer to **string** | Your reference for the payment instrument, maximum 150 characters. | [optional] 
**Status** | Pointer to **string** | The status of the payment instrument. If a status is not specified when creating a payment instrument, it is set to **active** by default. However, there can be exceptions for cards based on the &#x60;card.formFactor&#x60; and the &#x60;issuingCountryCode&#x60;. For example, when issuing physical cards in the US, the default status is **inactive**.  Possible values:    * **active**:  The payment instrument is active and can be used to make payments.    * **inactive**: The payment instrument is inactive and cannot be used to make payments.    * **suspended**: The payment instrument is suspended, either because it was stolen or lost.    * **closed**: The payment instrument is permanently closed. This action cannot be undone.    | [optional] 
**StatusReason** | Pointer to **string** | The reason for updating the status of the payment instrument.  Possible values: **lost**, **stolen**, **damaged**, **suspectedFraud**, **expired**, **endOfLife**, **accountClosure**, **other**. If the reason is **other**, you must also send the &#x60;statusComment&#x60; parameter describing the status change. | [optional] 
**Type** | **string** | Type of payment instrument.  Possible value: **card**, **bankAccount**.  | 

## Methods

### NewPaymentInstrument

`func NewPaymentInstrument(balanceAccountId string, id string, issuingCountryCode string, type_ string, ) *PaymentInstrument`

NewPaymentInstrument instantiates a new PaymentInstrument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInstrumentWithDefaults

`func NewPaymentInstrumentWithDefaults() *PaymentInstrument`

NewPaymentInstrumentWithDefaults instantiates a new PaymentInstrument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceAccountId

`func (o *PaymentInstrument) GetBalanceAccountId() string`

GetBalanceAccountId returns the BalanceAccountId field if non-nil, zero value otherwise.

### GetBalanceAccountIdOk

`func (o *PaymentInstrument) GetBalanceAccountIdOk() (*string, bool)`

GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccountId

`func (o *PaymentInstrument) SetBalanceAccountId(v string)`

SetBalanceAccountId sets BalanceAccountId field to given value.


### GetBankAccount

`func (o *PaymentInstrument) GetBankAccount() PaymentInstrumentBankAccount`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *PaymentInstrument) GetBankAccountOk() (*PaymentInstrumentBankAccount, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *PaymentInstrument) SetBankAccount(v PaymentInstrumentBankAccount)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *PaymentInstrument) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### GetCard

`func (o *PaymentInstrument) GetCard() Card`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *PaymentInstrument) GetCardOk() (*Card, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *PaymentInstrument) SetCard(v Card)`

SetCard sets Card field to given value.

### HasCard

`func (o *PaymentInstrument) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentInstrument) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentInstrument) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentInstrument) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentInstrument) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *PaymentInstrument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentInstrument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentInstrument) SetId(v string)`

SetId sets Id field to given value.


### GetIssuingCountryCode

`func (o *PaymentInstrument) GetIssuingCountryCode() string`

GetIssuingCountryCode returns the IssuingCountryCode field if non-nil, zero value otherwise.

### GetIssuingCountryCodeOk

`func (o *PaymentInstrument) GetIssuingCountryCodeOk() (*string, bool)`

GetIssuingCountryCodeOk returns a tuple with the IssuingCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountryCode

`func (o *PaymentInstrument) SetIssuingCountryCode(v string)`

SetIssuingCountryCode sets IssuingCountryCode field to given value.


### GetPaymentInstrumentGroupId

`func (o *PaymentInstrument) GetPaymentInstrumentGroupId() string`

GetPaymentInstrumentGroupId returns the PaymentInstrumentGroupId field if non-nil, zero value otherwise.

### GetPaymentInstrumentGroupIdOk

`func (o *PaymentInstrument) GetPaymentInstrumentGroupIdOk() (*string, bool)`

GetPaymentInstrumentGroupIdOk returns a tuple with the PaymentInstrumentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrumentGroupId

`func (o *PaymentInstrument) SetPaymentInstrumentGroupId(v string)`

SetPaymentInstrumentGroupId sets PaymentInstrumentGroupId field to given value.

### HasPaymentInstrumentGroupId

`func (o *PaymentInstrument) HasPaymentInstrumentGroupId() bool`

HasPaymentInstrumentGroupId returns a boolean if a field has been set.

### GetReference

`func (o *PaymentInstrument) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentInstrument) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentInstrument) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaymentInstrument) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentInstrument) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentInstrument) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentInstrument) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentInstrument) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *PaymentInstrument) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *PaymentInstrument) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *PaymentInstrument) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *PaymentInstrument) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetType

`func (o *PaymentInstrument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentInstrument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentInstrument) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


