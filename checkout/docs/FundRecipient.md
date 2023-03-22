# FundRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**PaymentMethod** | Pointer to [**CardDetails**](CardDetails.md) |  | [optional] 
**ShopperEmail** | Pointer to **string** | the email address of the person | [optional] 
**ShopperName** | Pointer to [**Name**](Name.md) |  | [optional] 
**ShopperReference** | Pointer to **string** | Required for recurring payments.  Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. &gt; Your reference must not include personally identifiable information (PII), for example name or email address. | [optional] 
**StoredPaymentMethodId** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**SubMerchant** | Pointer to [**SubMerchant**](SubMerchant.md) |  | [optional] 
**TelephoneNumber** | Pointer to **string** | the telephone number of the person | [optional] 
**WalletIdentifier** | Pointer to **string** | indicates where the money is going | [optional] 
**WalletOwnerTaxId** | Pointer to **string** | indicates the tax identifier of the fund recepient | [optional] 

## Methods

### NewFundRecipient

`func NewFundRecipient() *FundRecipient`

NewFundRecipient instantiates a new FundRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundRecipientWithDefaults

`func NewFundRecipientWithDefaults() *FundRecipient`

NewFundRecipientWithDefaults instantiates a new FundRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAddress

`func (o *FundRecipient) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *FundRecipient) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *FundRecipient) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *FundRecipient) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *FundRecipient) GetPaymentMethod() CardDetails`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *FundRecipient) GetPaymentMethodOk() (*CardDetails, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *FundRecipient) SetPaymentMethod(v CardDetails)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *FundRecipient) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetShopperEmail

`func (o *FundRecipient) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *FundRecipient) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *FundRecipient) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.

### HasShopperEmail

`func (o *FundRecipient) HasShopperEmail() bool`

HasShopperEmail returns a boolean if a field has been set.

### GetShopperName

`func (o *FundRecipient) GetShopperName() Name`

GetShopperName returns the ShopperName field if non-nil, zero value otherwise.

### GetShopperNameOk

`func (o *FundRecipient) GetShopperNameOk() (*Name, bool)`

GetShopperNameOk returns a tuple with the ShopperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperName

`func (o *FundRecipient) SetShopperName(v Name)`

SetShopperName sets ShopperName field to given value.

### HasShopperName

`func (o *FundRecipient) HasShopperName() bool`

HasShopperName returns a boolean if a field has been set.

### GetShopperReference

`func (o *FundRecipient) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *FundRecipient) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *FundRecipient) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *FundRecipient) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.

### GetStoredPaymentMethodId

`func (o *FundRecipient) GetStoredPaymentMethodId() string`

GetStoredPaymentMethodId returns the StoredPaymentMethodId field if non-nil, zero value otherwise.

### GetStoredPaymentMethodIdOk

`func (o *FundRecipient) GetStoredPaymentMethodIdOk() (*string, bool)`

GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredPaymentMethodId

`func (o *FundRecipient) SetStoredPaymentMethodId(v string)`

SetStoredPaymentMethodId sets StoredPaymentMethodId field to given value.

### HasStoredPaymentMethodId

`func (o *FundRecipient) HasStoredPaymentMethodId() bool`

HasStoredPaymentMethodId returns a boolean if a field has been set.

### GetSubMerchant

`func (o *FundRecipient) GetSubMerchant() SubMerchant`

GetSubMerchant returns the SubMerchant field if non-nil, zero value otherwise.

### GetSubMerchantOk

`func (o *FundRecipient) GetSubMerchantOk() (*SubMerchant, bool)`

GetSubMerchantOk returns a tuple with the SubMerchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchant

`func (o *FundRecipient) SetSubMerchant(v SubMerchant)`

SetSubMerchant sets SubMerchant field to given value.

### HasSubMerchant

`func (o *FundRecipient) HasSubMerchant() bool`

HasSubMerchant returns a boolean if a field has been set.

### GetTelephoneNumber

`func (o *FundRecipient) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *FundRecipient) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *FundRecipient) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *FundRecipient) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.

### GetWalletIdentifier

`func (o *FundRecipient) GetWalletIdentifier() string`

GetWalletIdentifier returns the WalletIdentifier field if non-nil, zero value otherwise.

### GetWalletIdentifierOk

`func (o *FundRecipient) GetWalletIdentifierOk() (*string, bool)`

GetWalletIdentifierOk returns a tuple with the WalletIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletIdentifier

`func (o *FundRecipient) SetWalletIdentifier(v string)`

SetWalletIdentifier sets WalletIdentifier field to given value.

### HasWalletIdentifier

`func (o *FundRecipient) HasWalletIdentifier() bool`

HasWalletIdentifier returns a boolean if a field has been set.

### GetWalletOwnerTaxId

`func (o *FundRecipient) GetWalletOwnerTaxId() string`

GetWalletOwnerTaxId returns the WalletOwnerTaxId field if non-nil, zero value otherwise.

### GetWalletOwnerTaxIdOk

`func (o *FundRecipient) GetWalletOwnerTaxIdOk() (*string, bool)`

GetWalletOwnerTaxIdOk returns a tuple with the WalletOwnerTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletOwnerTaxId

`func (o *FundRecipient) SetWalletOwnerTaxId(v string)`

SetWalletOwnerTaxId sets WalletOwnerTaxId field to given value.

### HasWalletOwnerTaxId

`func (o *FundRecipient) HasWalletOwnerTaxId() bool`

HasWalletOwnerTaxId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


