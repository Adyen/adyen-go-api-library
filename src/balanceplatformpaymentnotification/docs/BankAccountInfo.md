# BankAccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Iban** | Pointer to **string** | The international bank account number as defined in the [ISO-13616](https://www.iso.org/standard/81090.html) standard. | [optional] 
**OwnerName** | Pointer to [**Name**](Name.md) |  | [optional] 

## Methods

### NewBankAccountInfo

`func NewBankAccountInfo() *BankAccountInfo`

NewBankAccountInfo instantiates a new BankAccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountInfoWithDefaults

`func NewBankAccountInfoWithDefaults() *BankAccountInfo`

NewBankAccountInfoWithDefaults instantiates a new BankAccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *BankAccountInfo) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BankAccountInfo) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BankAccountInfo) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BankAccountInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIban

`func (o *BankAccountInfo) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *BankAccountInfo) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *BankAccountInfo) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *BankAccountInfo) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetOwnerName

`func (o *BankAccountInfo) GetOwnerName() Name`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *BankAccountInfo) GetOwnerNameOk() (*Name, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *BankAccountInfo) SetOwnerName(v Name)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *BankAccountInfo) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


