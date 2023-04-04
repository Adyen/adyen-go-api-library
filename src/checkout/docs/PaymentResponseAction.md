# PaymentResponseAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentData** | Pointer to **string** | A value that must be submitted to the &#x60;/payments/details&#x60; endpoint to verify this payment. | [optional] 
**PaymentMethodType** | Pointer to **string** | Specifies the payment method. | [optional] 
**Type** | **string** | **voucher** | 
**Url** | Pointer to **string** | Specifies the URL to redirect to. | [optional] 
**Data** | Pointer to **map[string]string** | When the redirect URL must be accessed via POST, use this data to post to the redirect URL. | [optional] 
**Method** | Pointer to **string** | Specifies the HTTP method, for example GET or POST. | [optional] 
**NativeRedirectData** | Pointer to **string** | Native SDK&#39;s redirect data containing the direct issuer link and state data that must be submitted to the /v1/nativeRedirect/redirectResult. | [optional] 
**ExpiresAt** | Pointer to **string** | The date time of the voucher expiry. | [optional] 
**QrCodeData** | Pointer to **string** | The contents of the QR code as a UTF8 string. | [optional] 
**SdkData** | Pointer to **map[string]string** | The data to pass to the SDK. | [optional] 
**AuthorisationToken** | Pointer to **string** | A token needed to authorise a payment. | [optional] 
**Subtype** | Pointer to **string** | A subtype of the token. | [optional] 
**Token** | Pointer to **string** | A token to pass to the 3DS2 Component to get the fingerprint. | [optional] 
**AlternativeReference** | Pointer to **string** | The voucher alternative reference code. | [optional] 
**CollectionInstitutionNumber** | Pointer to **string** | A collection institution number (store number) for Econtext Pay-Easy ATM. | [optional] 
**DownloadUrl** | Pointer to **string** | The URL to download the voucher. | [optional] 
**Entity** | Pointer to **string** | An entity number of Multibanco. | [optional] 
**InitialAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**InstructionsUrl** | Pointer to **string** | The URL to the detailed instructions to make payment using the voucher. | [optional] 
**Issuer** | Pointer to **string** | The issuer of the voucher. | [optional] 
**MaskedTelephoneNumber** | Pointer to **string** | The shopper telephone number (partially masked). | [optional] 
**MerchantName** | Pointer to **string** | The merchant name. | [optional] 
**MerchantReference** | Pointer to **string** | The merchant reference. | [optional] 
**Reference** | Pointer to **string** | The voucher reference code. | [optional] 
**ShopperEmail** | Pointer to **string** | The shopper email. | [optional] 
**ShopperName** | Pointer to **string** | The shopper name. | [optional] 
**Surcharge** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**TotalAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 

## Methods

### NewPaymentResponseAction

`func NewPaymentResponseAction(type_ string, ) *PaymentResponseAction`

NewPaymentResponseAction instantiates a new PaymentResponseAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentResponseActionWithDefaults

`func NewPaymentResponseActionWithDefaults() *PaymentResponseAction`

NewPaymentResponseActionWithDefaults instantiates a new PaymentResponseAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentData

`func (o *PaymentResponseAction) GetPaymentData() string`

GetPaymentData returns the PaymentData field if non-nil, zero value otherwise.

### GetPaymentDataOk

`func (o *PaymentResponseAction) GetPaymentDataOk() (*string, bool)`

GetPaymentDataOk returns a tuple with the PaymentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentData

`func (o *PaymentResponseAction) SetPaymentData(v string)`

SetPaymentData sets PaymentData field to given value.

### HasPaymentData

`func (o *PaymentResponseAction) HasPaymentData() bool`

HasPaymentData returns a boolean if a field has been set.

### GetPaymentMethodType

`func (o *PaymentResponseAction) GetPaymentMethodType() string`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *PaymentResponseAction) GetPaymentMethodTypeOk() (*string, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *PaymentResponseAction) SetPaymentMethodType(v string)`

SetPaymentMethodType sets PaymentMethodType field to given value.

### HasPaymentMethodType

`func (o *PaymentResponseAction) HasPaymentMethodType() bool`

HasPaymentMethodType returns a boolean if a field has been set.

### GetType

`func (o *PaymentResponseAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentResponseAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentResponseAction) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *PaymentResponseAction) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PaymentResponseAction) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PaymentResponseAction) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PaymentResponseAction) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetData

`func (o *PaymentResponseAction) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentResponseAction) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentResponseAction) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *PaymentResponseAction) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMethod

`func (o *PaymentResponseAction) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PaymentResponseAction) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PaymentResponseAction) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PaymentResponseAction) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetNativeRedirectData

`func (o *PaymentResponseAction) GetNativeRedirectData() string`

GetNativeRedirectData returns the NativeRedirectData field if non-nil, zero value otherwise.

### GetNativeRedirectDataOk

`func (o *PaymentResponseAction) GetNativeRedirectDataOk() (*string, bool)`

GetNativeRedirectDataOk returns a tuple with the NativeRedirectData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeRedirectData

`func (o *PaymentResponseAction) SetNativeRedirectData(v string)`

SetNativeRedirectData sets NativeRedirectData field to given value.

### HasNativeRedirectData

`func (o *PaymentResponseAction) HasNativeRedirectData() bool`

HasNativeRedirectData returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PaymentResponseAction) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PaymentResponseAction) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PaymentResponseAction) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PaymentResponseAction) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetQrCodeData

`func (o *PaymentResponseAction) GetQrCodeData() string`

GetQrCodeData returns the QrCodeData field if non-nil, zero value otherwise.

### GetQrCodeDataOk

`func (o *PaymentResponseAction) GetQrCodeDataOk() (*string, bool)`

GetQrCodeDataOk returns a tuple with the QrCodeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeData

`func (o *PaymentResponseAction) SetQrCodeData(v string)`

SetQrCodeData sets QrCodeData field to given value.

### HasQrCodeData

`func (o *PaymentResponseAction) HasQrCodeData() bool`

HasQrCodeData returns a boolean if a field has been set.

### GetSdkData

`func (o *PaymentResponseAction) GetSdkData() map[string]string`

GetSdkData returns the SdkData field if non-nil, zero value otherwise.

### GetSdkDataOk

`func (o *PaymentResponseAction) GetSdkDataOk() (*map[string]string, bool)`

GetSdkDataOk returns a tuple with the SdkData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkData

`func (o *PaymentResponseAction) SetSdkData(v map[string]string)`

SetSdkData sets SdkData field to given value.

### HasSdkData

`func (o *PaymentResponseAction) HasSdkData() bool`

HasSdkData returns a boolean if a field has been set.

### GetAuthorisationToken

`func (o *PaymentResponseAction) GetAuthorisationToken() string`

GetAuthorisationToken returns the AuthorisationToken field if non-nil, zero value otherwise.

### GetAuthorisationTokenOk

`func (o *PaymentResponseAction) GetAuthorisationTokenOk() (*string, bool)`

GetAuthorisationTokenOk returns a tuple with the AuthorisationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorisationToken

`func (o *PaymentResponseAction) SetAuthorisationToken(v string)`

SetAuthorisationToken sets AuthorisationToken field to given value.

### HasAuthorisationToken

`func (o *PaymentResponseAction) HasAuthorisationToken() bool`

HasAuthorisationToken returns a boolean if a field has been set.

### GetSubtype

`func (o *PaymentResponseAction) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *PaymentResponseAction) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *PaymentResponseAction) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *PaymentResponseAction) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetToken

`func (o *PaymentResponseAction) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentResponseAction) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentResponseAction) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PaymentResponseAction) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetAlternativeReference

`func (o *PaymentResponseAction) GetAlternativeReference() string`

GetAlternativeReference returns the AlternativeReference field if non-nil, zero value otherwise.

### GetAlternativeReferenceOk

`func (o *PaymentResponseAction) GetAlternativeReferenceOk() (*string, bool)`

GetAlternativeReferenceOk returns a tuple with the AlternativeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeReference

`func (o *PaymentResponseAction) SetAlternativeReference(v string)`

SetAlternativeReference sets AlternativeReference field to given value.

### HasAlternativeReference

`func (o *PaymentResponseAction) HasAlternativeReference() bool`

HasAlternativeReference returns a boolean if a field has been set.

### GetCollectionInstitutionNumber

`func (o *PaymentResponseAction) GetCollectionInstitutionNumber() string`

GetCollectionInstitutionNumber returns the CollectionInstitutionNumber field if non-nil, zero value otherwise.

### GetCollectionInstitutionNumberOk

`func (o *PaymentResponseAction) GetCollectionInstitutionNumberOk() (*string, bool)`

GetCollectionInstitutionNumberOk returns a tuple with the CollectionInstitutionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionInstitutionNumber

`func (o *PaymentResponseAction) SetCollectionInstitutionNumber(v string)`

SetCollectionInstitutionNumber sets CollectionInstitutionNumber field to given value.

### HasCollectionInstitutionNumber

`func (o *PaymentResponseAction) HasCollectionInstitutionNumber() bool`

HasCollectionInstitutionNumber returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *PaymentResponseAction) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *PaymentResponseAction) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *PaymentResponseAction) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *PaymentResponseAction) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetEntity

`func (o *PaymentResponseAction) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *PaymentResponseAction) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *PaymentResponseAction) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *PaymentResponseAction) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetInitialAmount

`func (o *PaymentResponseAction) GetInitialAmount() Amount`

GetInitialAmount returns the InitialAmount field if non-nil, zero value otherwise.

### GetInitialAmountOk

`func (o *PaymentResponseAction) GetInitialAmountOk() (*Amount, bool)`

GetInitialAmountOk returns a tuple with the InitialAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialAmount

`func (o *PaymentResponseAction) SetInitialAmount(v Amount)`

SetInitialAmount sets InitialAmount field to given value.

### HasInitialAmount

`func (o *PaymentResponseAction) HasInitialAmount() bool`

HasInitialAmount returns a boolean if a field has been set.

### GetInstructionsUrl

`func (o *PaymentResponseAction) GetInstructionsUrl() string`

GetInstructionsUrl returns the InstructionsUrl field if non-nil, zero value otherwise.

### GetInstructionsUrlOk

`func (o *PaymentResponseAction) GetInstructionsUrlOk() (*string, bool)`

GetInstructionsUrlOk returns a tuple with the InstructionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructionsUrl

`func (o *PaymentResponseAction) SetInstructionsUrl(v string)`

SetInstructionsUrl sets InstructionsUrl field to given value.

### HasInstructionsUrl

`func (o *PaymentResponseAction) HasInstructionsUrl() bool`

HasInstructionsUrl returns a boolean if a field has been set.

### GetIssuer

`func (o *PaymentResponseAction) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *PaymentResponseAction) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *PaymentResponseAction) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *PaymentResponseAction) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetMaskedTelephoneNumber

`func (o *PaymentResponseAction) GetMaskedTelephoneNumber() string`

GetMaskedTelephoneNumber returns the MaskedTelephoneNumber field if non-nil, zero value otherwise.

### GetMaskedTelephoneNumberOk

`func (o *PaymentResponseAction) GetMaskedTelephoneNumberOk() (*string, bool)`

GetMaskedTelephoneNumberOk returns a tuple with the MaskedTelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedTelephoneNumber

`func (o *PaymentResponseAction) SetMaskedTelephoneNumber(v string)`

SetMaskedTelephoneNumber sets MaskedTelephoneNumber field to given value.

### HasMaskedTelephoneNumber

`func (o *PaymentResponseAction) HasMaskedTelephoneNumber() bool`

HasMaskedTelephoneNumber returns a boolean if a field has been set.

### GetMerchantName

`func (o *PaymentResponseAction) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *PaymentResponseAction) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *PaymentResponseAction) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *PaymentResponseAction) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### GetMerchantReference

`func (o *PaymentResponseAction) GetMerchantReference() string`

GetMerchantReference returns the MerchantReference field if non-nil, zero value otherwise.

### GetMerchantReferenceOk

`func (o *PaymentResponseAction) GetMerchantReferenceOk() (*string, bool)`

GetMerchantReferenceOk returns a tuple with the MerchantReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReference

`func (o *PaymentResponseAction) SetMerchantReference(v string)`

SetMerchantReference sets MerchantReference field to given value.

### HasMerchantReference

`func (o *PaymentResponseAction) HasMerchantReference() bool`

HasMerchantReference returns a boolean if a field has been set.

### GetReference

`func (o *PaymentResponseAction) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentResponseAction) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentResponseAction) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaymentResponseAction) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetShopperEmail

`func (o *PaymentResponseAction) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *PaymentResponseAction) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *PaymentResponseAction) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.

### HasShopperEmail

`func (o *PaymentResponseAction) HasShopperEmail() bool`

HasShopperEmail returns a boolean if a field has been set.

### GetShopperName

`func (o *PaymentResponseAction) GetShopperName() string`

GetShopperName returns the ShopperName field if non-nil, zero value otherwise.

### GetShopperNameOk

`func (o *PaymentResponseAction) GetShopperNameOk() (*string, bool)`

GetShopperNameOk returns a tuple with the ShopperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperName

`func (o *PaymentResponseAction) SetShopperName(v string)`

SetShopperName sets ShopperName field to given value.

### HasShopperName

`func (o *PaymentResponseAction) HasShopperName() bool`

HasShopperName returns a boolean if a field has been set.

### GetSurcharge

`func (o *PaymentResponseAction) GetSurcharge() Amount`

GetSurcharge returns the Surcharge field if non-nil, zero value otherwise.

### GetSurchargeOk

`func (o *PaymentResponseAction) GetSurchargeOk() (*Amount, bool)`

GetSurchargeOk returns a tuple with the Surcharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharge

`func (o *PaymentResponseAction) SetSurcharge(v Amount)`

SetSurcharge sets Surcharge field to given value.

### HasSurcharge

`func (o *PaymentResponseAction) HasSurcharge() bool`

HasSurcharge returns a boolean if a field has been set.

### GetTotalAmount

`func (o *PaymentResponseAction) GetTotalAmount() Amount`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *PaymentResponseAction) GetTotalAmountOk() (*Amount, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *PaymentResponseAction) SetTotalAmount(v Amount)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *PaymentResponseAction) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


