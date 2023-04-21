# Card

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cvc** | Pointer to **string** | The [card verification code](https://docs.adyen.com/payments-fundamentals/payment-glossary#card-security-code-cvc-cvv-cid) (1-20 characters). Depending on the card brand, it is known also as: * CVV2/CVC2 – length: 3 digits * CID – length: 4 digits &gt; If you are using [Client-Side Encryption](https://docs.adyen.com/classic-integration/cse-integration-ecommerce), the CVC code is present in the encrypted data. You must never post the card details to the server. &gt; This field must be always present in a [one-click payment request](https://docs.adyen.com/classic-integration/recurring-payments). &gt; When this value is returned in a response, it is always empty because it is not stored. | [optional] 
**ExpiryMonth** | **string** | The card expiry month. Format: 2 digits, zero-padded for single digits. For example: * 03 &#x3D; March * 11 &#x3D; November | 
**ExpiryYear** | **string** | The card expiry year. Format: 4 digits. For example: 2020 | 
**HolderName** | **string** | The name of the cardholder, as printed on the card. | 
**IssueNumber** | Pointer to **string** | The issue number of the card (for some UK debit cards only). | [optional] 
**Number** | **string** | The card number (4-19 characters). Do not use any separators. When this value is returned in a response, only the last 4 digits of the card number are returned. | 
**StartMonth** | Pointer to **string** | The month component of the start date (for some UK debit cards only). | [optional] 
**StartYear** | Pointer to **string** | The year component of the start date (for some UK debit cards only). | [optional] 

## Methods

### NewCard

`func NewCard(expiryMonth string, expiryYear string, holderName string, number string, ) *Card`

NewCard instantiates a new Card object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardWithDefaults

`func NewCardWithDefaults() *Card`

NewCardWithDefaults instantiates a new Card object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvc

`func (o *Card) GetCvc() string`

GetCvc returns the Cvc field if non-nil, zero value otherwise.

### GetCvcOk

`func (o *Card) GetCvcOk() (*string, bool)`

GetCvcOk returns a tuple with the Cvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvc

`func (o *Card) SetCvc(v string)`

SetCvc sets Cvc field to given value.

### HasCvc

`func (o *Card) HasCvc() bool`

HasCvc returns a boolean if a field has been set.

### GetExpiryMonth

`func (o *Card) GetExpiryMonth() string`

GetExpiryMonth returns the ExpiryMonth field if non-nil, zero value otherwise.

### GetExpiryMonthOk

`func (o *Card) GetExpiryMonthOk() (*string, bool)`

GetExpiryMonthOk returns a tuple with the ExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryMonth

`func (o *Card) SetExpiryMonth(v string)`

SetExpiryMonth sets ExpiryMonth field to given value.


### GetExpiryYear

`func (o *Card) GetExpiryYear() string`

GetExpiryYear returns the ExpiryYear field if non-nil, zero value otherwise.

### GetExpiryYearOk

`func (o *Card) GetExpiryYearOk() (*string, bool)`

GetExpiryYearOk returns a tuple with the ExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryYear

`func (o *Card) SetExpiryYear(v string)`

SetExpiryYear sets ExpiryYear field to given value.


### GetHolderName

`func (o *Card) GetHolderName() string`

GetHolderName returns the HolderName field if non-nil, zero value otherwise.

### GetHolderNameOk

`func (o *Card) GetHolderNameOk() (*string, bool)`

GetHolderNameOk returns a tuple with the HolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderName

`func (o *Card) SetHolderName(v string)`

SetHolderName sets HolderName field to given value.


### GetIssueNumber

`func (o *Card) GetIssueNumber() string`

GetIssueNumber returns the IssueNumber field if non-nil, zero value otherwise.

### GetIssueNumberOk

`func (o *Card) GetIssueNumberOk() (*string, bool)`

GetIssueNumberOk returns a tuple with the IssueNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueNumber

`func (o *Card) SetIssueNumber(v string)`

SetIssueNumber sets IssueNumber field to given value.

### HasIssueNumber

`func (o *Card) HasIssueNumber() bool`

HasIssueNumber returns a boolean if a field has been set.

### GetNumber

`func (o *Card) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Card) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Card) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetStartMonth

`func (o *Card) GetStartMonth() string`

GetStartMonth returns the StartMonth field if non-nil, zero value otherwise.

### GetStartMonthOk

`func (o *Card) GetStartMonthOk() (*string, bool)`

GetStartMonthOk returns a tuple with the StartMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMonth

`func (o *Card) SetStartMonth(v string)`

SetStartMonth sets StartMonth field to given value.

### HasStartMonth

`func (o *Card) HasStartMonth() bool`

HasStartMonth returns a boolean if a field has been set.

### GetStartYear

`func (o *Card) GetStartYear() string`

GetStartYear returns the StartYear field if non-nil, zero value otherwise.

### GetStartYearOk

`func (o *Card) GetStartYearOk() (*string, bool)`

GetStartYearOk returns a tuple with the StartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartYear

`func (o *Card) SetStartYear(v string)`

SetStartYear sets StartYear field to given value.

### HasStartYear

`func (o *Card) HasStartYear() bool`

HasStartYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


