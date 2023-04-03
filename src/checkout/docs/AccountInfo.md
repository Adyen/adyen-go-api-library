# AccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAgeIndicator** | Pointer to **string** | Indicator for the length of time since this shopper account was created in the merchant&#39;s environment. Allowed values: * notApplicable * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days | [optional] 
**AccountChangeDate** | Pointer to **time.Time** | Date when the shopper&#39;s account was last changed. | [optional] 
**AccountChangeIndicator** | Pointer to **string** | Indicator for the length of time since the shopper&#39;s account was last updated. Allowed values: * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days | [optional] 
**AccountCreationDate** | Pointer to **time.Time** | Date when the shopper&#39;s account was created. | [optional] 
**AccountType** | Pointer to **string** | Indicates the type of account. For example, for a multi-account card product. Allowed values: * notApplicable * credit * debit | [optional] 
**AddCardAttemptsDay** | Pointer to **int32** | Number of attempts the shopper tried to add a card to their account in the last day. | [optional] 
**DeliveryAddressUsageDate** | Pointer to **time.Time** | Date the selected delivery address was first used. | [optional] 
**DeliveryAddressUsageIndicator** | Pointer to **string** | Indicator for the length of time since this delivery address was first used. Allowed values: * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days | [optional] 
**HomePhone** | Pointer to **string** | Shopper&#39;s home phone number (including the country code). | [optional] 
**MobilePhone** | Pointer to **string** | Shopper&#39;s mobile phone number (including the country code). | [optional] 
**PasswordChangeDate** | Pointer to **time.Time** | Date when the shopper last changed their password. | [optional] 
**PasswordChangeIndicator** | Pointer to **string** | Indicator when the shopper has changed their password. Allowed values: * notApplicable * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days | [optional] 
**PastTransactionsDay** | Pointer to **int32** | Number of all transactions (successful and abandoned) from this shopper in the past 24 hours. | [optional] 
**PastTransactionsYear** | Pointer to **int32** | Number of all transactions (successful and abandoned) from this shopper in the past year. | [optional] 
**PaymentAccountAge** | Pointer to **time.Time** | Date this payment method was added to the shopper&#39;s account. | [optional] 
**PaymentAccountIndicator** | Pointer to **string** | Indicator for the length of time since this payment method was added to this shopper&#39;s account. Allowed values: * notApplicable * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days | [optional] 
**PurchasesLast6Months** | Pointer to **int32** | Number of successful purchases in the last six months. | [optional] 
**SuspiciousActivity** | Pointer to **bool** | Whether suspicious activity was recorded on this account. | [optional] 
**WorkPhone** | Pointer to **string** | Shopper&#39;s work phone number (including the country code). | [optional] 

## Methods

### NewAccountInfo

`func NewAccountInfo() *AccountInfo`

NewAccountInfo instantiates a new AccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInfoWithDefaults

`func NewAccountInfoWithDefaults() *AccountInfo`

NewAccountInfoWithDefaults instantiates a new AccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountAgeIndicator

`func (o *AccountInfo) GetAccountAgeIndicator() string`

GetAccountAgeIndicator returns the AccountAgeIndicator field if non-nil, zero value otherwise.

### GetAccountAgeIndicatorOk

`func (o *AccountInfo) GetAccountAgeIndicatorOk() (*string, bool)`

GetAccountAgeIndicatorOk returns a tuple with the AccountAgeIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAgeIndicator

`func (o *AccountInfo) SetAccountAgeIndicator(v string)`

SetAccountAgeIndicator sets AccountAgeIndicator field to given value.

### HasAccountAgeIndicator

`func (o *AccountInfo) HasAccountAgeIndicator() bool`

HasAccountAgeIndicator returns a boolean if a field has been set.

### GetAccountChangeDate

`func (o *AccountInfo) GetAccountChangeDate() time.Time`

GetAccountChangeDate returns the AccountChangeDate field if non-nil, zero value otherwise.

### GetAccountChangeDateOk

`func (o *AccountInfo) GetAccountChangeDateOk() (*time.Time, bool)`

GetAccountChangeDateOk returns a tuple with the AccountChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountChangeDate

`func (o *AccountInfo) SetAccountChangeDate(v time.Time)`

SetAccountChangeDate sets AccountChangeDate field to given value.

### HasAccountChangeDate

`func (o *AccountInfo) HasAccountChangeDate() bool`

HasAccountChangeDate returns a boolean if a field has been set.

### GetAccountChangeIndicator

`func (o *AccountInfo) GetAccountChangeIndicator() string`

GetAccountChangeIndicator returns the AccountChangeIndicator field if non-nil, zero value otherwise.

### GetAccountChangeIndicatorOk

`func (o *AccountInfo) GetAccountChangeIndicatorOk() (*string, bool)`

GetAccountChangeIndicatorOk returns a tuple with the AccountChangeIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountChangeIndicator

`func (o *AccountInfo) SetAccountChangeIndicator(v string)`

SetAccountChangeIndicator sets AccountChangeIndicator field to given value.

### HasAccountChangeIndicator

`func (o *AccountInfo) HasAccountChangeIndicator() bool`

HasAccountChangeIndicator returns a boolean if a field has been set.

### GetAccountCreationDate

`func (o *AccountInfo) GetAccountCreationDate() time.Time`

GetAccountCreationDate returns the AccountCreationDate field if non-nil, zero value otherwise.

### GetAccountCreationDateOk

`func (o *AccountInfo) GetAccountCreationDateOk() (*time.Time, bool)`

GetAccountCreationDateOk returns a tuple with the AccountCreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreationDate

`func (o *AccountInfo) SetAccountCreationDate(v time.Time)`

SetAccountCreationDate sets AccountCreationDate field to given value.

### HasAccountCreationDate

`func (o *AccountInfo) HasAccountCreationDate() bool`

HasAccountCreationDate returns a boolean if a field has been set.

### GetAccountType

`func (o *AccountInfo) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountInfo) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountInfo) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountInfo) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAddCardAttemptsDay

`func (o *AccountInfo) GetAddCardAttemptsDay() int32`

GetAddCardAttemptsDay returns the AddCardAttemptsDay field if non-nil, zero value otherwise.

### GetAddCardAttemptsDayOk

`func (o *AccountInfo) GetAddCardAttemptsDayOk() (*int32, bool)`

GetAddCardAttemptsDayOk returns a tuple with the AddCardAttemptsDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCardAttemptsDay

`func (o *AccountInfo) SetAddCardAttemptsDay(v int32)`

SetAddCardAttemptsDay sets AddCardAttemptsDay field to given value.

### HasAddCardAttemptsDay

`func (o *AccountInfo) HasAddCardAttemptsDay() bool`

HasAddCardAttemptsDay returns a boolean if a field has been set.

### GetDeliveryAddressUsageDate

`func (o *AccountInfo) GetDeliveryAddressUsageDate() time.Time`

GetDeliveryAddressUsageDate returns the DeliveryAddressUsageDate field if non-nil, zero value otherwise.

### GetDeliveryAddressUsageDateOk

`func (o *AccountInfo) GetDeliveryAddressUsageDateOk() (*time.Time, bool)`

GetDeliveryAddressUsageDateOk returns a tuple with the DeliveryAddressUsageDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddressUsageDate

`func (o *AccountInfo) SetDeliveryAddressUsageDate(v time.Time)`

SetDeliveryAddressUsageDate sets DeliveryAddressUsageDate field to given value.

### HasDeliveryAddressUsageDate

`func (o *AccountInfo) HasDeliveryAddressUsageDate() bool`

HasDeliveryAddressUsageDate returns a boolean if a field has been set.

### GetDeliveryAddressUsageIndicator

`func (o *AccountInfo) GetDeliveryAddressUsageIndicator() string`

GetDeliveryAddressUsageIndicator returns the DeliveryAddressUsageIndicator field if non-nil, zero value otherwise.

### GetDeliveryAddressUsageIndicatorOk

`func (o *AccountInfo) GetDeliveryAddressUsageIndicatorOk() (*string, bool)`

GetDeliveryAddressUsageIndicatorOk returns a tuple with the DeliveryAddressUsageIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddressUsageIndicator

`func (o *AccountInfo) SetDeliveryAddressUsageIndicator(v string)`

SetDeliveryAddressUsageIndicator sets DeliveryAddressUsageIndicator field to given value.

### HasDeliveryAddressUsageIndicator

`func (o *AccountInfo) HasDeliveryAddressUsageIndicator() bool`

HasDeliveryAddressUsageIndicator returns a boolean if a field has been set.

### GetHomePhone

`func (o *AccountInfo) GetHomePhone() string`

GetHomePhone returns the HomePhone field if non-nil, zero value otherwise.

### GetHomePhoneOk

`func (o *AccountInfo) GetHomePhoneOk() (*string, bool)`

GetHomePhoneOk returns a tuple with the HomePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePhone

`func (o *AccountInfo) SetHomePhone(v string)`

SetHomePhone sets HomePhone field to given value.

### HasHomePhone

`func (o *AccountInfo) HasHomePhone() bool`

HasHomePhone returns a boolean if a field has been set.

### GetMobilePhone

`func (o *AccountInfo) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *AccountInfo) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *AccountInfo) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *AccountInfo) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetPasswordChangeDate

`func (o *AccountInfo) GetPasswordChangeDate() time.Time`

GetPasswordChangeDate returns the PasswordChangeDate field if non-nil, zero value otherwise.

### GetPasswordChangeDateOk

`func (o *AccountInfo) GetPasswordChangeDateOk() (*time.Time, bool)`

GetPasswordChangeDateOk returns a tuple with the PasswordChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChangeDate

`func (o *AccountInfo) SetPasswordChangeDate(v time.Time)`

SetPasswordChangeDate sets PasswordChangeDate field to given value.

### HasPasswordChangeDate

`func (o *AccountInfo) HasPasswordChangeDate() bool`

HasPasswordChangeDate returns a boolean if a field has been set.

### GetPasswordChangeIndicator

`func (o *AccountInfo) GetPasswordChangeIndicator() string`

GetPasswordChangeIndicator returns the PasswordChangeIndicator field if non-nil, zero value otherwise.

### GetPasswordChangeIndicatorOk

`func (o *AccountInfo) GetPasswordChangeIndicatorOk() (*string, bool)`

GetPasswordChangeIndicatorOk returns a tuple with the PasswordChangeIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChangeIndicator

`func (o *AccountInfo) SetPasswordChangeIndicator(v string)`

SetPasswordChangeIndicator sets PasswordChangeIndicator field to given value.

### HasPasswordChangeIndicator

`func (o *AccountInfo) HasPasswordChangeIndicator() bool`

HasPasswordChangeIndicator returns a boolean if a field has been set.

### GetPastTransactionsDay

`func (o *AccountInfo) GetPastTransactionsDay() int32`

GetPastTransactionsDay returns the PastTransactionsDay field if non-nil, zero value otherwise.

### GetPastTransactionsDayOk

`func (o *AccountInfo) GetPastTransactionsDayOk() (*int32, bool)`

GetPastTransactionsDayOk returns a tuple with the PastTransactionsDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastTransactionsDay

`func (o *AccountInfo) SetPastTransactionsDay(v int32)`

SetPastTransactionsDay sets PastTransactionsDay field to given value.

### HasPastTransactionsDay

`func (o *AccountInfo) HasPastTransactionsDay() bool`

HasPastTransactionsDay returns a boolean if a field has been set.

### GetPastTransactionsYear

`func (o *AccountInfo) GetPastTransactionsYear() int32`

GetPastTransactionsYear returns the PastTransactionsYear field if non-nil, zero value otherwise.

### GetPastTransactionsYearOk

`func (o *AccountInfo) GetPastTransactionsYearOk() (*int32, bool)`

GetPastTransactionsYearOk returns a tuple with the PastTransactionsYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastTransactionsYear

`func (o *AccountInfo) SetPastTransactionsYear(v int32)`

SetPastTransactionsYear sets PastTransactionsYear field to given value.

### HasPastTransactionsYear

`func (o *AccountInfo) HasPastTransactionsYear() bool`

HasPastTransactionsYear returns a boolean if a field has been set.

### GetPaymentAccountAge

`func (o *AccountInfo) GetPaymentAccountAge() time.Time`

GetPaymentAccountAge returns the PaymentAccountAge field if non-nil, zero value otherwise.

### GetPaymentAccountAgeOk

`func (o *AccountInfo) GetPaymentAccountAgeOk() (*time.Time, bool)`

GetPaymentAccountAgeOk returns a tuple with the PaymentAccountAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccountAge

`func (o *AccountInfo) SetPaymentAccountAge(v time.Time)`

SetPaymentAccountAge sets PaymentAccountAge field to given value.

### HasPaymentAccountAge

`func (o *AccountInfo) HasPaymentAccountAge() bool`

HasPaymentAccountAge returns a boolean if a field has been set.

### GetPaymentAccountIndicator

`func (o *AccountInfo) GetPaymentAccountIndicator() string`

GetPaymentAccountIndicator returns the PaymentAccountIndicator field if non-nil, zero value otherwise.

### GetPaymentAccountIndicatorOk

`func (o *AccountInfo) GetPaymentAccountIndicatorOk() (*string, bool)`

GetPaymentAccountIndicatorOk returns a tuple with the PaymentAccountIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccountIndicator

`func (o *AccountInfo) SetPaymentAccountIndicator(v string)`

SetPaymentAccountIndicator sets PaymentAccountIndicator field to given value.

### HasPaymentAccountIndicator

`func (o *AccountInfo) HasPaymentAccountIndicator() bool`

HasPaymentAccountIndicator returns a boolean if a field has been set.

### GetPurchasesLast6Months

`func (o *AccountInfo) GetPurchasesLast6Months() int32`

GetPurchasesLast6Months returns the PurchasesLast6Months field if non-nil, zero value otherwise.

### GetPurchasesLast6MonthsOk

`func (o *AccountInfo) GetPurchasesLast6MonthsOk() (*int32, bool)`

GetPurchasesLast6MonthsOk returns a tuple with the PurchasesLast6Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasesLast6Months

`func (o *AccountInfo) SetPurchasesLast6Months(v int32)`

SetPurchasesLast6Months sets PurchasesLast6Months field to given value.

### HasPurchasesLast6Months

`func (o *AccountInfo) HasPurchasesLast6Months() bool`

HasPurchasesLast6Months returns a boolean if a field has been set.

### GetSuspiciousActivity

`func (o *AccountInfo) GetSuspiciousActivity() bool`

GetSuspiciousActivity returns the SuspiciousActivity field if non-nil, zero value otherwise.

### GetSuspiciousActivityOk

`func (o *AccountInfo) GetSuspiciousActivityOk() (*bool, bool)`

GetSuspiciousActivityOk returns a tuple with the SuspiciousActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspiciousActivity

`func (o *AccountInfo) SetSuspiciousActivity(v bool)`

SetSuspiciousActivity sets SuspiciousActivity field to given value.

### HasSuspiciousActivity

`func (o *AccountInfo) HasSuspiciousActivity() bool`

HasSuspiciousActivity returns a boolean if a field has been set.

### GetWorkPhone

`func (o *AccountInfo) GetWorkPhone() string`

GetWorkPhone returns the WorkPhone field if non-nil, zero value otherwise.

### GetWorkPhoneOk

`func (o *AccountInfo) GetWorkPhoneOk() (*string, bool)`

GetWorkPhoneOk returns a tuple with the WorkPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkPhone

`func (o *AccountInfo) SetWorkPhone(v string)`

SetWorkPhone sets WorkPhone field to given value.

### HasWorkPhone

`func (o *AccountInfo) HasWorkPhone() bool`

HasWorkPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


