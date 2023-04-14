# AcctInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChAccAgeInd** | Pointer to **string** | Length of time that the cardholder has had the account with the 3DS Requestor.  Allowed values: * **01** — No account * **02** — Created during this transaction * **03** — Less than 30 days * **04** — 30–60 days * **05** — More than 60 days | [optional] 
**ChAccChange** | Pointer to **string** | Date that the cardholder’s account with the 3DS Requestor was last changed, including Billing or Shipping address, new payment account, or new user(s) added.  Format: **YYYYMMDD** | [optional] 
**ChAccChangeInd** | Pointer to **string** | Length of time since the cardholder’s account information with the 3DS Requestor was last changed, including Billing or Shipping address, new payment account, or new user(s) added.  Allowed values: * **01** — Changed during this transaction * **02** — Less than 30 days * **03** — 30–60 days * **04** — More than 60 days | [optional] 
**ChAccPwChange** | Pointer to **string** | Date that cardholder’s account with the 3DS Requestor had a password change or account reset.  Format: **YYYYMMDD** | [optional] 
**ChAccPwChangeInd** | Pointer to **string** | Indicates the length of time since the cardholder’s account with the 3DS Requestor had a password change or account reset.  Allowed values: * **01** — No change * **02** — Changed during this transaction * **03** — Less than 30 days * **04** — 30–60 days * **05** — More than 60 days | [optional] 
**ChAccString** | Pointer to **string** | Date that the cardholder opened the account with the 3DS Requestor.  Format: **YYYYMMDD** | [optional] 
**NbPurchaseAccount** | Pointer to **string** | Number of purchases with this cardholder account during the previous six months. Max length: 4 characters. | [optional] 
**PaymentAccAge** | Pointer to **string** | String that the payment account was enrolled in the cardholder’s account with the 3DS Requestor.  Format: **YYYYMMDD** | [optional] 
**PaymentAccInd** | Pointer to **string** | Indicates the length of time that the payment account was enrolled in the cardholder’s account with the 3DS Requestor.  Allowed values: * **01** — No account (guest checkout) * **02** — During this transaction * **03** — Less than 30 days * **04** — 30–60 days * **05** — More than 60 days | [optional] 
**ProvisionAttemptsDay** | Pointer to **string** | Number of Add Card attempts in the last 24 hours. Max length: 3 characters. | [optional] 
**ShipAddressUsage** | Pointer to **string** | String when the shipping address used for this transaction was first used with the 3DS Requestor.  Format: **YYYYMMDD** | [optional] 
**ShipAddressUsageInd** | Pointer to **string** | Indicates when the shipping address used for this transaction was first used with the 3DS Requestor.  Allowed values: * **01** — This transaction * **02** — Less than 30 days * **03** — 30–60 days * **04** — More than 60 days | [optional] 
**ShipNameIndicator** | Pointer to **string** | Indicates if the Cardholder Name on the account is identical to the shipping Name used for this transaction.  Allowed values: * **01** — Account Name identical to shipping Name * **02** — Account Name different to shipping Name | [optional] 
**SuspiciousAccActivity** | Pointer to **string** | Indicates whether the 3DS Requestor has experienced suspicious activity (including previous fraud) on the cardholder account.  Allowed values: * **01** — No suspicious activity has been observed * **02** — Suspicious activity has been observed | [optional] 
**TxnActivityDay** | Pointer to **string** | Number of transactions (successful and abandoned) for this cardholder account with the 3DS Requestor across all payment accounts in the previous 24 hours. Max length: 3 characters. | [optional] 
**TxnActivityYear** | Pointer to **string** | Number of transactions (successful and abandoned) for this cardholder account with the 3DS Requestor across all payment accounts in the previous year. Max length: 3 characters. | [optional] 

## Methods

### NewAcctInfo

`func NewAcctInfo() *AcctInfo`

NewAcctInfo instantiates a new AcctInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcctInfoWithDefaults

`func NewAcctInfoWithDefaults() *AcctInfo`

NewAcctInfoWithDefaults instantiates a new AcctInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChAccAgeInd

`func (o *AcctInfo) GetChAccAgeInd() string`

GetChAccAgeInd returns the ChAccAgeInd field if non-nil, zero value otherwise.

### GetChAccAgeIndOk

`func (o *AcctInfo) GetChAccAgeIndOk() (*string, bool)`

GetChAccAgeIndOk returns a tuple with the ChAccAgeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChAccAgeInd

`func (o *AcctInfo) SetChAccAgeInd(v string)`

SetChAccAgeInd sets ChAccAgeInd field to given value.

### HasChAccAgeInd

`func (o *AcctInfo) HasChAccAgeInd() bool`

HasChAccAgeInd returns a boolean if a field has been set.

### GetChAccChange

`func (o *AcctInfo) GetChAccChange() string`

GetChAccChange returns the ChAccChange field if non-nil, zero value otherwise.

### GetChAccChangeOk

`func (o *AcctInfo) GetChAccChangeOk() (*string, bool)`

GetChAccChangeOk returns a tuple with the ChAccChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChAccChange

`func (o *AcctInfo) SetChAccChange(v string)`

SetChAccChange sets ChAccChange field to given value.

### HasChAccChange

`func (o *AcctInfo) HasChAccChange() bool`

HasChAccChange returns a boolean if a field has been set.

### GetChAccChangeInd

`func (o *AcctInfo) GetChAccChangeInd() string`

GetChAccChangeInd returns the ChAccChangeInd field if non-nil, zero value otherwise.

### GetChAccChangeIndOk

`func (o *AcctInfo) GetChAccChangeIndOk() (*string, bool)`

GetChAccChangeIndOk returns a tuple with the ChAccChangeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChAccChangeInd

`func (o *AcctInfo) SetChAccChangeInd(v string)`

SetChAccChangeInd sets ChAccChangeInd field to given value.

### HasChAccChangeInd

`func (o *AcctInfo) HasChAccChangeInd() bool`

HasChAccChangeInd returns a boolean if a field has been set.

### GetChAccPwChange

`func (o *AcctInfo) GetChAccPwChange() string`

GetChAccPwChange returns the ChAccPwChange field if non-nil, zero value otherwise.

### GetChAccPwChangeOk

`func (o *AcctInfo) GetChAccPwChangeOk() (*string, bool)`

GetChAccPwChangeOk returns a tuple with the ChAccPwChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChAccPwChange

`func (o *AcctInfo) SetChAccPwChange(v string)`

SetChAccPwChange sets ChAccPwChange field to given value.

### HasChAccPwChange

`func (o *AcctInfo) HasChAccPwChange() bool`

HasChAccPwChange returns a boolean if a field has been set.

### GetChAccPwChangeInd

`func (o *AcctInfo) GetChAccPwChangeInd() string`

GetChAccPwChangeInd returns the ChAccPwChangeInd field if non-nil, zero value otherwise.

### GetChAccPwChangeIndOk

`func (o *AcctInfo) GetChAccPwChangeIndOk() (*string, bool)`

GetChAccPwChangeIndOk returns a tuple with the ChAccPwChangeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChAccPwChangeInd

`func (o *AcctInfo) SetChAccPwChangeInd(v string)`

SetChAccPwChangeInd sets ChAccPwChangeInd field to given value.

### HasChAccPwChangeInd

`func (o *AcctInfo) HasChAccPwChangeInd() bool`

HasChAccPwChangeInd returns a boolean if a field has been set.

### GetChAccString

`func (o *AcctInfo) GetChAccString() string`

GetChAccString returns the ChAccString field if non-nil, zero value otherwise.

### GetChAccStringOk

`func (o *AcctInfo) GetChAccStringOk() (*string, bool)`

GetChAccStringOk returns a tuple with the ChAccString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChAccString

`func (o *AcctInfo) SetChAccString(v string)`

SetChAccString sets ChAccString field to given value.

### HasChAccString

`func (o *AcctInfo) HasChAccString() bool`

HasChAccString returns a boolean if a field has been set.

### GetNbPurchaseAccount

`func (o *AcctInfo) GetNbPurchaseAccount() string`

GetNbPurchaseAccount returns the NbPurchaseAccount field if non-nil, zero value otherwise.

### GetNbPurchaseAccountOk

`func (o *AcctInfo) GetNbPurchaseAccountOk() (*string, bool)`

GetNbPurchaseAccountOk returns a tuple with the NbPurchaseAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbPurchaseAccount

`func (o *AcctInfo) SetNbPurchaseAccount(v string)`

SetNbPurchaseAccount sets NbPurchaseAccount field to given value.

### HasNbPurchaseAccount

`func (o *AcctInfo) HasNbPurchaseAccount() bool`

HasNbPurchaseAccount returns a boolean if a field has been set.

### GetPaymentAccAge

`func (o *AcctInfo) GetPaymentAccAge() string`

GetPaymentAccAge returns the PaymentAccAge field if non-nil, zero value otherwise.

### GetPaymentAccAgeOk

`func (o *AcctInfo) GetPaymentAccAgeOk() (*string, bool)`

GetPaymentAccAgeOk returns a tuple with the PaymentAccAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccAge

`func (o *AcctInfo) SetPaymentAccAge(v string)`

SetPaymentAccAge sets PaymentAccAge field to given value.

### HasPaymentAccAge

`func (o *AcctInfo) HasPaymentAccAge() bool`

HasPaymentAccAge returns a boolean if a field has been set.

### GetPaymentAccInd

`func (o *AcctInfo) GetPaymentAccInd() string`

GetPaymentAccInd returns the PaymentAccInd field if non-nil, zero value otherwise.

### GetPaymentAccIndOk

`func (o *AcctInfo) GetPaymentAccIndOk() (*string, bool)`

GetPaymentAccIndOk returns a tuple with the PaymentAccInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccInd

`func (o *AcctInfo) SetPaymentAccInd(v string)`

SetPaymentAccInd sets PaymentAccInd field to given value.

### HasPaymentAccInd

`func (o *AcctInfo) HasPaymentAccInd() bool`

HasPaymentAccInd returns a boolean if a field has been set.

### GetProvisionAttemptsDay

`func (o *AcctInfo) GetProvisionAttemptsDay() string`

GetProvisionAttemptsDay returns the ProvisionAttemptsDay field if non-nil, zero value otherwise.

### GetProvisionAttemptsDayOk

`func (o *AcctInfo) GetProvisionAttemptsDayOk() (*string, bool)`

GetProvisionAttemptsDayOk returns a tuple with the ProvisionAttemptsDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionAttemptsDay

`func (o *AcctInfo) SetProvisionAttemptsDay(v string)`

SetProvisionAttemptsDay sets ProvisionAttemptsDay field to given value.

### HasProvisionAttemptsDay

`func (o *AcctInfo) HasProvisionAttemptsDay() bool`

HasProvisionAttemptsDay returns a boolean if a field has been set.

### GetShipAddressUsage

`func (o *AcctInfo) GetShipAddressUsage() string`

GetShipAddressUsage returns the ShipAddressUsage field if non-nil, zero value otherwise.

### GetShipAddressUsageOk

`func (o *AcctInfo) GetShipAddressUsageOk() (*string, bool)`

GetShipAddressUsageOk returns a tuple with the ShipAddressUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipAddressUsage

`func (o *AcctInfo) SetShipAddressUsage(v string)`

SetShipAddressUsage sets ShipAddressUsage field to given value.

### HasShipAddressUsage

`func (o *AcctInfo) HasShipAddressUsage() bool`

HasShipAddressUsage returns a boolean if a field has been set.

### GetShipAddressUsageInd

`func (o *AcctInfo) GetShipAddressUsageInd() string`

GetShipAddressUsageInd returns the ShipAddressUsageInd field if non-nil, zero value otherwise.

### GetShipAddressUsageIndOk

`func (o *AcctInfo) GetShipAddressUsageIndOk() (*string, bool)`

GetShipAddressUsageIndOk returns a tuple with the ShipAddressUsageInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipAddressUsageInd

`func (o *AcctInfo) SetShipAddressUsageInd(v string)`

SetShipAddressUsageInd sets ShipAddressUsageInd field to given value.

### HasShipAddressUsageInd

`func (o *AcctInfo) HasShipAddressUsageInd() bool`

HasShipAddressUsageInd returns a boolean if a field has been set.

### GetShipNameIndicator

`func (o *AcctInfo) GetShipNameIndicator() string`

GetShipNameIndicator returns the ShipNameIndicator field if non-nil, zero value otherwise.

### GetShipNameIndicatorOk

`func (o *AcctInfo) GetShipNameIndicatorOk() (*string, bool)`

GetShipNameIndicatorOk returns a tuple with the ShipNameIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipNameIndicator

`func (o *AcctInfo) SetShipNameIndicator(v string)`

SetShipNameIndicator sets ShipNameIndicator field to given value.

### HasShipNameIndicator

`func (o *AcctInfo) HasShipNameIndicator() bool`

HasShipNameIndicator returns a boolean if a field has been set.

### GetSuspiciousAccActivity

`func (o *AcctInfo) GetSuspiciousAccActivity() string`

GetSuspiciousAccActivity returns the SuspiciousAccActivity field if non-nil, zero value otherwise.

### GetSuspiciousAccActivityOk

`func (o *AcctInfo) GetSuspiciousAccActivityOk() (*string, bool)`

GetSuspiciousAccActivityOk returns a tuple with the SuspiciousAccActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspiciousAccActivity

`func (o *AcctInfo) SetSuspiciousAccActivity(v string)`

SetSuspiciousAccActivity sets SuspiciousAccActivity field to given value.

### HasSuspiciousAccActivity

`func (o *AcctInfo) HasSuspiciousAccActivity() bool`

HasSuspiciousAccActivity returns a boolean if a field has been set.

### GetTxnActivityDay

`func (o *AcctInfo) GetTxnActivityDay() string`

GetTxnActivityDay returns the TxnActivityDay field if non-nil, zero value otherwise.

### GetTxnActivityDayOk

`func (o *AcctInfo) GetTxnActivityDayOk() (*string, bool)`

GetTxnActivityDayOk returns a tuple with the TxnActivityDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnActivityDay

`func (o *AcctInfo) SetTxnActivityDay(v string)`

SetTxnActivityDay sets TxnActivityDay field to given value.

### HasTxnActivityDay

`func (o *AcctInfo) HasTxnActivityDay() bool`

HasTxnActivityDay returns a boolean if a field has been set.

### GetTxnActivityYear

`func (o *AcctInfo) GetTxnActivityYear() string`

GetTxnActivityYear returns the TxnActivityYear field if non-nil, zero value otherwise.

### GetTxnActivityYearOk

`func (o *AcctInfo) GetTxnActivityYearOk() (*string, bool)`

GetTxnActivityYearOk returns a tuple with the TxnActivityYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnActivityYear

`func (o *AcctInfo) SetTxnActivityYear(v string)`

SetTxnActivityYear sets TxnActivityYear field to given value.

### HasTxnActivityYear

`func (o *AcctInfo) HasTxnActivityYear() bool`

HasTxnActivityYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


