# PaymentInstrumentBankAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iban** | **string** | The international bank account number as defined in the [ISO-13616](https://www.iso.org/standard/81090.html) standard. | 
**Type** | **string** | **usLocal** | [default to "usLocal"]
**AccountNumber** | **string** | The bank account number, without separators or whitespace. | 
**AccountType** | Pointer to **string** | The bank account type.  Possible values: **checking** or **savings**. Defaults to **checking**. | [optional] [default to "checking"]
**RoutingNumber** | **string** | The 9-digit [routing number](https://en.wikipedia.org/wiki/ABA_routing_transit_number), without separators or whitespace. | 

## Methods

### NewPaymentInstrumentBankAccount

`func NewPaymentInstrumentBankAccount(iban string, type_ string, accountNumber string, routingNumber string, ) *PaymentInstrumentBankAccount`

NewPaymentInstrumentBankAccount instantiates a new PaymentInstrumentBankAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInstrumentBankAccountWithDefaults

`func NewPaymentInstrumentBankAccountWithDefaults() *PaymentInstrumentBankAccount`

NewPaymentInstrumentBankAccountWithDefaults instantiates a new PaymentInstrumentBankAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIban

`func (o *PaymentInstrumentBankAccount) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *PaymentInstrumentBankAccount) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *PaymentInstrumentBankAccount) SetIban(v string)`

SetIban sets Iban field to given value.


### GetType

`func (o *PaymentInstrumentBankAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentInstrumentBankAccount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentInstrumentBankAccount) SetType(v string)`

SetType sets Type field to given value.


### GetAccountNumber

`func (o *PaymentInstrumentBankAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *PaymentInstrumentBankAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *PaymentInstrumentBankAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAccountType

`func (o *PaymentInstrumentBankAccount) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *PaymentInstrumentBankAccount) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *PaymentInstrumentBankAccount) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *PaymentInstrumentBankAccount) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetRoutingNumber

`func (o *PaymentInstrumentBankAccount) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *PaymentInstrumentBankAccount) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *PaymentInstrumentBankAccount) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


