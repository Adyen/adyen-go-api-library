# BacsDirectDebitDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccountNumber** | Pointer to **string** | The bank account number (without separators). | [optional] 
**BankLocationId** | Pointer to **string** | The bank routing number of the account. | [optional] 
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**HolderName** | Pointer to **string** | The name of the bank account holder. | [optional] 
**RecurringDetailReference** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**StoredPaymentMethodId** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**Type** | Pointer to **string** | **directdebit_GB** | [optional] [default to "directdebit_GB"]

## Methods

### NewBacsDirectDebitDetails

`func NewBacsDirectDebitDetails() *BacsDirectDebitDetails`

NewBacsDirectDebitDetails instantiates a new BacsDirectDebitDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBacsDirectDebitDetailsWithDefaults

`func NewBacsDirectDebitDetailsWithDefaults() *BacsDirectDebitDetails`

NewBacsDirectDebitDetailsWithDefaults instantiates a new BacsDirectDebitDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAccountNumber

`func (o *BacsDirectDebitDetails) GetBankAccountNumber() string`

GetBankAccountNumber returns the BankAccountNumber field if non-nil, zero value otherwise.

### GetBankAccountNumberOk

`func (o *BacsDirectDebitDetails) GetBankAccountNumberOk() (*string, bool)`

GetBankAccountNumberOk returns a tuple with the BankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountNumber

`func (o *BacsDirectDebitDetails) SetBankAccountNumber(v string)`

SetBankAccountNumber sets BankAccountNumber field to given value.

### HasBankAccountNumber

`func (o *BacsDirectDebitDetails) HasBankAccountNumber() bool`

HasBankAccountNumber returns a boolean if a field has been set.

### GetBankLocationId

`func (o *BacsDirectDebitDetails) GetBankLocationId() string`

GetBankLocationId returns the BankLocationId field if non-nil, zero value otherwise.

### GetBankLocationIdOk

`func (o *BacsDirectDebitDetails) GetBankLocationIdOk() (*string, bool)`

GetBankLocationIdOk returns a tuple with the BankLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankLocationId

`func (o *BacsDirectDebitDetails) SetBankLocationId(v string)`

SetBankLocationId sets BankLocationId field to given value.

### HasBankLocationId

`func (o *BacsDirectDebitDetails) HasBankLocationId() bool`

HasBankLocationId returns a boolean if a field has been set.

### GetCheckoutAttemptId

`func (o *BacsDirectDebitDetails) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *BacsDirectDebitDetails) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *BacsDirectDebitDetails) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *BacsDirectDebitDetails) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetHolderName

`func (o *BacsDirectDebitDetails) GetHolderName() string`

GetHolderName returns the HolderName field if non-nil, zero value otherwise.

### GetHolderNameOk

`func (o *BacsDirectDebitDetails) GetHolderNameOk() (*string, bool)`

GetHolderNameOk returns a tuple with the HolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderName

`func (o *BacsDirectDebitDetails) SetHolderName(v string)`

SetHolderName sets HolderName field to given value.

### HasHolderName

`func (o *BacsDirectDebitDetails) HasHolderName() bool`

HasHolderName returns a boolean if a field has been set.

### GetRecurringDetailReference

`func (o *BacsDirectDebitDetails) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *BacsDirectDebitDetails) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *BacsDirectDebitDetails) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.

### HasRecurringDetailReference

`func (o *BacsDirectDebitDetails) HasRecurringDetailReference() bool`

HasRecurringDetailReference returns a boolean if a field has been set.

### GetStoredPaymentMethodId

`func (o *BacsDirectDebitDetails) GetStoredPaymentMethodId() string`

GetStoredPaymentMethodId returns the StoredPaymentMethodId field if non-nil, zero value otherwise.

### GetStoredPaymentMethodIdOk

`func (o *BacsDirectDebitDetails) GetStoredPaymentMethodIdOk() (*string, bool)`

GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredPaymentMethodId

`func (o *BacsDirectDebitDetails) SetStoredPaymentMethodId(v string)`

SetStoredPaymentMethodId sets StoredPaymentMethodId field to given value.

### HasStoredPaymentMethodId

`func (o *BacsDirectDebitDetails) HasStoredPaymentMethodId() bool`

HasStoredPaymentMethodId returns a boolean if a field has been set.

### GetType

`func (o *BacsDirectDebitDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BacsDirectDebitDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BacsDirectDebitDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BacsDirectDebitDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


