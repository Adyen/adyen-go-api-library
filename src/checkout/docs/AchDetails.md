# AchDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccountNumber** | **string** | The bank account number (without separators). | 
**BankLocationId** | Pointer to **string** | The bank routing number of the account. The field value is &#x60;nil&#x60; in most cases. | [optional] 
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**EncryptedBankAccountNumber** | Pointer to **string** | Encrypted bank account number. The bank account number (without separators). | [optional] 
**EncryptedBankLocationId** | Pointer to **string** | Encrypted location id. The bank routing number of the account. The field value is &#x60;nil&#x60; in most cases. | [optional] 
**OwnerName** | Pointer to **string** | The name of the bank account holder. If you submit a name with non-Latin characters, we automatically replace some of them with corresponding Latin characters to meet the FATF recommendations. For example: * χ12 is converted to ch12. * üA is converted to euA. * Peter Møller is converted to Peter Mller, because banks don&#39;t accept &#39;ø&#39;. After replacement, the ownerName must have at least three alphanumeric characters (A-Z, a-z, 0-9), and at least one of them must be a valid Latin character (A-Z, a-z). For example: * John17 - allowed. * J17 - allowed. * 171 - not allowed. * John-7 - allowed. &gt; If provided details don&#39;t match the required format, the response returns the error message: 203 &#39;Invalid bank account holder name&#39;. | [optional] 
**RecurringDetailReference** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**StoredPaymentMethodId** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**Type** | Pointer to **string** | **ach** | [optional] [default to "ach"]

## Methods

### NewAchDetails

`func NewAchDetails(bankAccountNumber string, ) *AchDetails`

NewAchDetails instantiates a new AchDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchDetailsWithDefaults

`func NewAchDetailsWithDefaults() *AchDetails`

NewAchDetailsWithDefaults instantiates a new AchDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAccountNumber

`func (o *AchDetails) GetBankAccountNumber() string`

GetBankAccountNumber returns the BankAccountNumber field if non-nil, zero value otherwise.

### GetBankAccountNumberOk

`func (o *AchDetails) GetBankAccountNumberOk() (*string, bool)`

GetBankAccountNumberOk returns a tuple with the BankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountNumber

`func (o *AchDetails) SetBankAccountNumber(v string)`

SetBankAccountNumber sets BankAccountNumber field to given value.


### GetBankLocationId

`func (o *AchDetails) GetBankLocationId() string`

GetBankLocationId returns the BankLocationId field if non-nil, zero value otherwise.

### GetBankLocationIdOk

`func (o *AchDetails) GetBankLocationIdOk() (*string, bool)`

GetBankLocationIdOk returns a tuple with the BankLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankLocationId

`func (o *AchDetails) SetBankLocationId(v string)`

SetBankLocationId sets BankLocationId field to given value.

### HasBankLocationId

`func (o *AchDetails) HasBankLocationId() bool`

HasBankLocationId returns a boolean if a field has been set.

### GetCheckoutAttemptId

`func (o *AchDetails) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *AchDetails) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *AchDetails) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *AchDetails) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetEncryptedBankAccountNumber

`func (o *AchDetails) GetEncryptedBankAccountNumber() string`

GetEncryptedBankAccountNumber returns the EncryptedBankAccountNumber field if non-nil, zero value otherwise.

### GetEncryptedBankAccountNumberOk

`func (o *AchDetails) GetEncryptedBankAccountNumberOk() (*string, bool)`

GetEncryptedBankAccountNumberOk returns a tuple with the EncryptedBankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedBankAccountNumber

`func (o *AchDetails) SetEncryptedBankAccountNumber(v string)`

SetEncryptedBankAccountNumber sets EncryptedBankAccountNumber field to given value.

### HasEncryptedBankAccountNumber

`func (o *AchDetails) HasEncryptedBankAccountNumber() bool`

HasEncryptedBankAccountNumber returns a boolean if a field has been set.

### GetEncryptedBankLocationId

`func (o *AchDetails) GetEncryptedBankLocationId() string`

GetEncryptedBankLocationId returns the EncryptedBankLocationId field if non-nil, zero value otherwise.

### GetEncryptedBankLocationIdOk

`func (o *AchDetails) GetEncryptedBankLocationIdOk() (*string, bool)`

GetEncryptedBankLocationIdOk returns a tuple with the EncryptedBankLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedBankLocationId

`func (o *AchDetails) SetEncryptedBankLocationId(v string)`

SetEncryptedBankLocationId sets EncryptedBankLocationId field to given value.

### HasEncryptedBankLocationId

`func (o *AchDetails) HasEncryptedBankLocationId() bool`

HasEncryptedBankLocationId returns a boolean if a field has been set.

### GetOwnerName

`func (o *AchDetails) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *AchDetails) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *AchDetails) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *AchDetails) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetRecurringDetailReference

`func (o *AchDetails) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *AchDetails) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *AchDetails) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.

### HasRecurringDetailReference

`func (o *AchDetails) HasRecurringDetailReference() bool`

HasRecurringDetailReference returns a boolean if a field has been set.

### GetStoredPaymentMethodId

`func (o *AchDetails) GetStoredPaymentMethodId() string`

GetStoredPaymentMethodId returns the StoredPaymentMethodId field if non-nil, zero value otherwise.

### GetStoredPaymentMethodIdOk

`func (o *AchDetails) GetStoredPaymentMethodIdOk() (*string, bool)`

GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredPaymentMethodId

`func (o *AchDetails) SetStoredPaymentMethodId(v string)`

SetStoredPaymentMethodId sets StoredPaymentMethodId field to given value.

### HasStoredPaymentMethodId

`func (o *AchDetails) HasStoredPaymentMethodId() bool`

HasStoredPaymentMethodId returns a boolean if a field has been set.

### GetType

`func (o *AchDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AchDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AchDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AchDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


