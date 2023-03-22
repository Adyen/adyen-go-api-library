# CheckoutVoucherAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternativeReference** | Pointer to **string** | The voucher alternative reference code. | [optional] 
**CollectionInstitutionNumber** | Pointer to **string** | A collection institution number (store number) for Econtext Pay-Easy ATM. | [optional] 
**DownloadUrl** | Pointer to **string** | The URL to download the voucher. | [optional] 
**Entity** | Pointer to **string** | An entity number of Multibanco. | [optional] 
**ExpiresAt** | Pointer to **string** | The date time of the voucher expiry. | [optional] 
**InitialAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**InstructionsUrl** | Pointer to **string** | The URL to the detailed instructions to make payment using the voucher. | [optional] 
**Issuer** | Pointer to **string** | The issuer of the voucher. | [optional] 
**MaskedTelephoneNumber** | Pointer to **string** | The shopper telephone number (partially masked). | [optional] 
**MerchantName** | Pointer to **string** | The merchant name. | [optional] 
**MerchantReference** | Pointer to **string** | The merchant reference. | [optional] 
**PaymentData** | Pointer to **string** | A value that must be submitted to the &#x60;/payments/details&#x60; endpoint to verify this payment. | [optional] 
**PaymentMethodType** | Pointer to **string** | Specifies the payment method. | [optional] 
**Reference** | Pointer to **string** | The voucher reference code. | [optional] 
**ShopperEmail** | Pointer to **string** | The shopper email. | [optional] 
**ShopperName** | Pointer to **string** | The shopper name. | [optional] 
**Surcharge** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**TotalAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**Type** | **string** | **voucher** | 
**Url** | Pointer to **string** | Specifies the URL to redirect to. | [optional] 

## Methods

### NewCheckoutVoucherAction

`func NewCheckoutVoucherAction(type_ string, ) *CheckoutVoucherAction`

NewCheckoutVoucherAction instantiates a new CheckoutVoucherAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutVoucherActionWithDefaults

`func NewCheckoutVoucherActionWithDefaults() *CheckoutVoucherAction`

NewCheckoutVoucherActionWithDefaults instantiates a new CheckoutVoucherAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternativeReference

`func (o *CheckoutVoucherAction) GetAlternativeReference() string`

GetAlternativeReference returns the AlternativeReference field if non-nil, zero value otherwise.

### GetAlternativeReferenceOk

`func (o *CheckoutVoucherAction) GetAlternativeReferenceOk() (*string, bool)`

GetAlternativeReferenceOk returns a tuple with the AlternativeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeReference

`func (o *CheckoutVoucherAction) SetAlternativeReference(v string)`

SetAlternativeReference sets AlternativeReference field to given value.

### HasAlternativeReference

`func (o *CheckoutVoucherAction) HasAlternativeReference() bool`

HasAlternativeReference returns a boolean if a field has been set.

### GetCollectionInstitutionNumber

`func (o *CheckoutVoucherAction) GetCollectionInstitutionNumber() string`

GetCollectionInstitutionNumber returns the CollectionInstitutionNumber field if non-nil, zero value otherwise.

### GetCollectionInstitutionNumberOk

`func (o *CheckoutVoucherAction) GetCollectionInstitutionNumberOk() (*string, bool)`

GetCollectionInstitutionNumberOk returns a tuple with the CollectionInstitutionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionInstitutionNumber

`func (o *CheckoutVoucherAction) SetCollectionInstitutionNumber(v string)`

SetCollectionInstitutionNumber sets CollectionInstitutionNumber field to given value.

### HasCollectionInstitutionNumber

`func (o *CheckoutVoucherAction) HasCollectionInstitutionNumber() bool`

HasCollectionInstitutionNumber returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *CheckoutVoucherAction) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *CheckoutVoucherAction) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *CheckoutVoucherAction) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *CheckoutVoucherAction) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetEntity

`func (o *CheckoutVoucherAction) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *CheckoutVoucherAction) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *CheckoutVoucherAction) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *CheckoutVoucherAction) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CheckoutVoucherAction) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CheckoutVoucherAction) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CheckoutVoucherAction) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CheckoutVoucherAction) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetInitialAmount

`func (o *CheckoutVoucherAction) GetInitialAmount() Amount`

GetInitialAmount returns the InitialAmount field if non-nil, zero value otherwise.

### GetInitialAmountOk

`func (o *CheckoutVoucherAction) GetInitialAmountOk() (*Amount, bool)`

GetInitialAmountOk returns a tuple with the InitialAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialAmount

`func (o *CheckoutVoucherAction) SetInitialAmount(v Amount)`

SetInitialAmount sets InitialAmount field to given value.

### HasInitialAmount

`func (o *CheckoutVoucherAction) HasInitialAmount() bool`

HasInitialAmount returns a boolean if a field has been set.

### GetInstructionsUrl

`func (o *CheckoutVoucherAction) GetInstructionsUrl() string`

GetInstructionsUrl returns the InstructionsUrl field if non-nil, zero value otherwise.

### GetInstructionsUrlOk

`func (o *CheckoutVoucherAction) GetInstructionsUrlOk() (*string, bool)`

GetInstructionsUrlOk returns a tuple with the InstructionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructionsUrl

`func (o *CheckoutVoucherAction) SetInstructionsUrl(v string)`

SetInstructionsUrl sets InstructionsUrl field to given value.

### HasInstructionsUrl

`func (o *CheckoutVoucherAction) HasInstructionsUrl() bool`

HasInstructionsUrl returns a boolean if a field has been set.

### GetIssuer

`func (o *CheckoutVoucherAction) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CheckoutVoucherAction) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CheckoutVoucherAction) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CheckoutVoucherAction) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetMaskedTelephoneNumber

`func (o *CheckoutVoucherAction) GetMaskedTelephoneNumber() string`

GetMaskedTelephoneNumber returns the MaskedTelephoneNumber field if non-nil, zero value otherwise.

### GetMaskedTelephoneNumberOk

`func (o *CheckoutVoucherAction) GetMaskedTelephoneNumberOk() (*string, bool)`

GetMaskedTelephoneNumberOk returns a tuple with the MaskedTelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedTelephoneNumber

`func (o *CheckoutVoucherAction) SetMaskedTelephoneNumber(v string)`

SetMaskedTelephoneNumber sets MaskedTelephoneNumber field to given value.

### HasMaskedTelephoneNumber

`func (o *CheckoutVoucherAction) HasMaskedTelephoneNumber() bool`

HasMaskedTelephoneNumber returns a boolean if a field has been set.

### GetMerchantName

`func (o *CheckoutVoucherAction) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *CheckoutVoucherAction) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *CheckoutVoucherAction) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *CheckoutVoucherAction) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### GetMerchantReference

`func (o *CheckoutVoucherAction) GetMerchantReference() string`

GetMerchantReference returns the MerchantReference field if non-nil, zero value otherwise.

### GetMerchantReferenceOk

`func (o *CheckoutVoucherAction) GetMerchantReferenceOk() (*string, bool)`

GetMerchantReferenceOk returns a tuple with the MerchantReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReference

`func (o *CheckoutVoucherAction) SetMerchantReference(v string)`

SetMerchantReference sets MerchantReference field to given value.

### HasMerchantReference

`func (o *CheckoutVoucherAction) HasMerchantReference() bool`

HasMerchantReference returns a boolean if a field has been set.

### GetPaymentData

`func (o *CheckoutVoucherAction) GetPaymentData() string`

GetPaymentData returns the PaymentData field if non-nil, zero value otherwise.

### GetPaymentDataOk

`func (o *CheckoutVoucherAction) GetPaymentDataOk() (*string, bool)`

GetPaymentDataOk returns a tuple with the PaymentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentData

`func (o *CheckoutVoucherAction) SetPaymentData(v string)`

SetPaymentData sets PaymentData field to given value.

### HasPaymentData

`func (o *CheckoutVoucherAction) HasPaymentData() bool`

HasPaymentData returns a boolean if a field has been set.

### GetPaymentMethodType

`func (o *CheckoutVoucherAction) GetPaymentMethodType() string`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *CheckoutVoucherAction) GetPaymentMethodTypeOk() (*string, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *CheckoutVoucherAction) SetPaymentMethodType(v string)`

SetPaymentMethodType sets PaymentMethodType field to given value.

### HasPaymentMethodType

`func (o *CheckoutVoucherAction) HasPaymentMethodType() bool`

HasPaymentMethodType returns a boolean if a field has been set.

### GetReference

`func (o *CheckoutVoucherAction) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CheckoutVoucherAction) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CheckoutVoucherAction) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CheckoutVoucherAction) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetShopperEmail

`func (o *CheckoutVoucherAction) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *CheckoutVoucherAction) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *CheckoutVoucherAction) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.

### HasShopperEmail

`func (o *CheckoutVoucherAction) HasShopperEmail() bool`

HasShopperEmail returns a boolean if a field has been set.

### GetShopperName

`func (o *CheckoutVoucherAction) GetShopperName() string`

GetShopperName returns the ShopperName field if non-nil, zero value otherwise.

### GetShopperNameOk

`func (o *CheckoutVoucherAction) GetShopperNameOk() (*string, bool)`

GetShopperNameOk returns a tuple with the ShopperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperName

`func (o *CheckoutVoucherAction) SetShopperName(v string)`

SetShopperName sets ShopperName field to given value.

### HasShopperName

`func (o *CheckoutVoucherAction) HasShopperName() bool`

HasShopperName returns a boolean if a field has been set.

### GetSurcharge

`func (o *CheckoutVoucherAction) GetSurcharge() Amount`

GetSurcharge returns the Surcharge field if non-nil, zero value otherwise.

### GetSurchargeOk

`func (o *CheckoutVoucherAction) GetSurchargeOk() (*Amount, bool)`

GetSurchargeOk returns a tuple with the Surcharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharge

`func (o *CheckoutVoucherAction) SetSurcharge(v Amount)`

SetSurcharge sets Surcharge field to given value.

### HasSurcharge

`func (o *CheckoutVoucherAction) HasSurcharge() bool`

HasSurcharge returns a boolean if a field has been set.

### GetTotalAmount

`func (o *CheckoutVoucherAction) GetTotalAmount() Amount`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *CheckoutVoucherAction) GetTotalAmountOk() (*Amount, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *CheckoutVoucherAction) SetTotalAmount(v Amount)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *CheckoutVoucherAction) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetType

`func (o *CheckoutVoucherAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutVoucherAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutVoucherAction) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *CheckoutVoucherAction) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CheckoutVoucherAction) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CheckoutVoucherAction) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CheckoutVoucherAction) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


