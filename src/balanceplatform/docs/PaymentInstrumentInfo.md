# PaymentInstrumentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceAccountId** | **string** | The unique identifier of the [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/v1/post/balanceAccounts__resParam_id) associated with the payment instrument. | 
**Card** | Pointer to [**CardInfo**](CardInfo.md) |  | [optional] 
**Description** | Pointer to **string** | Your description for the payment instrument, maximum 300 characters. | [optional] 
**IssuingCountryCode** | **string** | The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code where the payment instrument is issued. For example, **NL** or **US**. | 
**PaymentInstrumentGroupId** | Pointer to **string** | The unique identifier of the [payment instrument group](https://docs.adyen.com/api-explorer/#/balanceplatform/v1/post/paymentInstrumentGroups__resParam_id) to which the payment instrument belongs. | [optional] 
**Reference** | Pointer to **string** | Your reference for the payment instrument, maximum 150 characters. | [optional] 
**Status** | Pointer to **string** | The status of the payment instrument. If a status is not specified when creating a payment instrument, it is set to **active** by default. However, there can be exceptions for cards based on the &#x60;card.formFactor&#x60; and the &#x60;issuingCountryCode&#x60;. For example, when issuing physical cards in the US, the default status is **inactive**.  Possible values:    * **active**:  The payment instrument is active and can be used to make payments.    * **inactive**: The payment instrument is inactive and cannot be used to make payments.    * **suspended**: The payment instrument is suspended, either because it was stolen or lost.    * **closed**: The payment instrument is permanently closed. This action cannot be undone.    | [optional] 
**StatusReason** | Pointer to **string** | The reason for updating the status of the payment instrument.  Possible values: **lost**, **stolen**, **damaged**, **suspectedFraud**, **expired**, **endOfLife**, **accountClosure**, **other**. If the reason is **other**, you must also send the &#x60;statusComment&#x60; parameter describing the status change. | [optional] 
**Type** | **string** | Type of payment instrument.  Possible value: **card**, **bankAccount**.  | 

## Methods

### NewPaymentInstrumentInfo

`func NewPaymentInstrumentInfo(balanceAccountId string, issuingCountryCode string, type_ string, ) *PaymentInstrumentInfo`

NewPaymentInstrumentInfo instantiates a new PaymentInstrumentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInstrumentInfoWithDefaults

`func NewPaymentInstrumentInfoWithDefaults() *PaymentInstrumentInfo`

NewPaymentInstrumentInfoWithDefaults instantiates a new PaymentInstrumentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceAccountId

`func (o *PaymentInstrumentInfo) GetBalanceAccountId() string`

GetBalanceAccountId returns the BalanceAccountId field if non-nil, zero value otherwise.

### GetBalanceAccountIdOk

`func (o *PaymentInstrumentInfo) GetBalanceAccountIdOk() (*string, bool)`

GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccountId

`func (o *PaymentInstrumentInfo) SetBalanceAccountId(v string)`

SetBalanceAccountId sets BalanceAccountId field to given value.


### GetCard

`func (o *PaymentInstrumentInfo) GetCard() CardInfo`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *PaymentInstrumentInfo) GetCardOk() (*CardInfo, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *PaymentInstrumentInfo) SetCard(v CardInfo)`

SetCard sets Card field to given value.

### HasCard

`func (o *PaymentInstrumentInfo) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentInstrumentInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentInstrumentInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentInstrumentInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentInstrumentInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIssuingCountryCode

`func (o *PaymentInstrumentInfo) GetIssuingCountryCode() string`

GetIssuingCountryCode returns the IssuingCountryCode field if non-nil, zero value otherwise.

### GetIssuingCountryCodeOk

`func (o *PaymentInstrumentInfo) GetIssuingCountryCodeOk() (*string, bool)`

GetIssuingCountryCodeOk returns a tuple with the IssuingCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountryCode

`func (o *PaymentInstrumentInfo) SetIssuingCountryCode(v string)`

SetIssuingCountryCode sets IssuingCountryCode field to given value.


### GetPaymentInstrumentGroupId

`func (o *PaymentInstrumentInfo) GetPaymentInstrumentGroupId() string`

GetPaymentInstrumentGroupId returns the PaymentInstrumentGroupId field if non-nil, zero value otherwise.

### GetPaymentInstrumentGroupIdOk

`func (o *PaymentInstrumentInfo) GetPaymentInstrumentGroupIdOk() (*string, bool)`

GetPaymentInstrumentGroupIdOk returns a tuple with the PaymentInstrumentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrumentGroupId

`func (o *PaymentInstrumentInfo) SetPaymentInstrumentGroupId(v string)`

SetPaymentInstrumentGroupId sets PaymentInstrumentGroupId field to given value.

### HasPaymentInstrumentGroupId

`func (o *PaymentInstrumentInfo) HasPaymentInstrumentGroupId() bool`

HasPaymentInstrumentGroupId returns a boolean if a field has been set.

### GetReference

`func (o *PaymentInstrumentInfo) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentInstrumentInfo) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentInstrumentInfo) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaymentInstrumentInfo) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentInstrumentInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentInstrumentInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentInstrumentInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentInstrumentInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *PaymentInstrumentInfo) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *PaymentInstrumentInfo) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *PaymentInstrumentInfo) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *PaymentInstrumentInfo) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetType

`func (o *PaymentInstrumentInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentInstrumentInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentInstrumentInfo) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


