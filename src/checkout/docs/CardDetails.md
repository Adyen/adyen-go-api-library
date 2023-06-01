# CardDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** | Secondary brand of the card. For example: **plastix**, **hmclub**. | [optional] 
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**CupsecureplusSmscode** | Pointer to **string** |  | [optional] 
**Cvc** | Pointer to **string** | The card verification code. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide). | [optional] 
**EncryptedCardNumber** | Pointer to **string** | The encrypted card number. | [optional] 
**EncryptedExpiryMonth** | Pointer to **string** | The encrypted card expiry month. | [optional] 
**EncryptedExpiryYear** | Pointer to **string** | The encrypted card expiry year. | [optional] 
**EncryptedSecurityCode** | Pointer to **string** | The encrypted card verification code. | [optional] 
**ExpiryMonth** | Pointer to **string** | The card expiry month. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide). | [optional] 
**ExpiryYear** | Pointer to **string** | The card expiry year. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide). | [optional] 
**FundingSource** | Pointer to **string** | The funding source that should be used when multiple sources are available. For Brazilian combo cards, by default the funding source is credit. To use debit, set this value to **debit**. | [optional] 
**HolderName** | Pointer to **string** | The name of the card holder. | [optional] 
**NetworkPaymentReference** | Pointer to **string** | The network token reference. This is the [&#x60;networkTxReference&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_additionalData-ResponseAdditionalDataCommon-networkTxReference) from the response to the first payment. | [optional] 
**Number** | Pointer to **string** | The card number. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide). | [optional] 
**RecurringDetailReference** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**ShopperNotificationReference** | Pointer to **string** | The &#x60;shopperNotificationReference&#x60; returned in the response when you requested to notify the shopper. Used only for recurring payments in India. | [optional] 
**StoredPaymentMethodId** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**ThreeDS2SdkVersion** | Pointer to **string** | Required for mobile integrations. Version of the 3D Secure 2 mobile SDK. | [optional] 
**Type** | Pointer to **string** | Default payment method details. Common for scheme payment methods, and for simple payment method details. | [optional] [default to "scheme"]

## Methods

### NewCardDetails

`func NewCardDetails() *CardDetails`

NewCardDetails instantiates a new CardDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardDetailsWithDefaults

`func NewCardDetailsWithDefaults() *CardDetails`

NewCardDetailsWithDefaults instantiates a new CardDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *CardDetails) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CardDetails) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CardDetails) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *CardDetails) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetCheckoutAttemptId

`func (o *CardDetails) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *CardDetails) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *CardDetails) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *CardDetails) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetCupsecureplusSmscode

`func (o *CardDetails) GetCupsecureplusSmscode() string`

GetCupsecureplusSmscode returns the CupsecureplusSmscode field if non-nil, zero value otherwise.

### GetCupsecureplusSmscodeOk

`func (o *CardDetails) GetCupsecureplusSmscodeOk() (*string, bool)`

GetCupsecureplusSmscodeOk returns a tuple with the CupsecureplusSmscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCupsecureplusSmscode

`func (o *CardDetails) SetCupsecureplusSmscode(v string)`

SetCupsecureplusSmscode sets CupsecureplusSmscode field to given value.

### HasCupsecureplusSmscode

`func (o *CardDetails) HasCupsecureplusSmscode() bool`

HasCupsecureplusSmscode returns a boolean if a field has been set.

### GetCvc

`func (o *CardDetails) GetCvc() string`

GetCvc returns the Cvc field if non-nil, zero value otherwise.

### GetCvcOk

`func (o *CardDetails) GetCvcOk() (*string, bool)`

GetCvcOk returns a tuple with the Cvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvc

`func (o *CardDetails) SetCvc(v string)`

SetCvc sets Cvc field to given value.

### HasCvc

`func (o *CardDetails) HasCvc() bool`

HasCvc returns a boolean if a field has been set.

### GetEncryptedCardNumber

`func (o *CardDetails) GetEncryptedCardNumber() string`

GetEncryptedCardNumber returns the EncryptedCardNumber field if non-nil, zero value otherwise.

### GetEncryptedCardNumberOk

`func (o *CardDetails) GetEncryptedCardNumberOk() (*string, bool)`

GetEncryptedCardNumberOk returns a tuple with the EncryptedCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedCardNumber

`func (o *CardDetails) SetEncryptedCardNumber(v string)`

SetEncryptedCardNumber sets EncryptedCardNumber field to given value.

### HasEncryptedCardNumber

`func (o *CardDetails) HasEncryptedCardNumber() bool`

HasEncryptedCardNumber returns a boolean if a field has been set.

### GetEncryptedExpiryMonth

`func (o *CardDetails) GetEncryptedExpiryMonth() string`

GetEncryptedExpiryMonth returns the EncryptedExpiryMonth field if non-nil, zero value otherwise.

### GetEncryptedExpiryMonthOk

`func (o *CardDetails) GetEncryptedExpiryMonthOk() (*string, bool)`

GetEncryptedExpiryMonthOk returns a tuple with the EncryptedExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedExpiryMonth

`func (o *CardDetails) SetEncryptedExpiryMonth(v string)`

SetEncryptedExpiryMonth sets EncryptedExpiryMonth field to given value.

### HasEncryptedExpiryMonth

`func (o *CardDetails) HasEncryptedExpiryMonth() bool`

HasEncryptedExpiryMonth returns a boolean if a field has been set.

### GetEncryptedExpiryYear

`func (o *CardDetails) GetEncryptedExpiryYear() string`

GetEncryptedExpiryYear returns the EncryptedExpiryYear field if non-nil, zero value otherwise.

### GetEncryptedExpiryYearOk

`func (o *CardDetails) GetEncryptedExpiryYearOk() (*string, bool)`

GetEncryptedExpiryYearOk returns a tuple with the EncryptedExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedExpiryYear

`func (o *CardDetails) SetEncryptedExpiryYear(v string)`

SetEncryptedExpiryYear sets EncryptedExpiryYear field to given value.

### HasEncryptedExpiryYear

`func (o *CardDetails) HasEncryptedExpiryYear() bool`

HasEncryptedExpiryYear returns a boolean if a field has been set.

### GetEncryptedSecurityCode

`func (o *CardDetails) GetEncryptedSecurityCode() string`

GetEncryptedSecurityCode returns the EncryptedSecurityCode field if non-nil, zero value otherwise.

### GetEncryptedSecurityCodeOk

`func (o *CardDetails) GetEncryptedSecurityCodeOk() (*string, bool)`

GetEncryptedSecurityCodeOk returns a tuple with the EncryptedSecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedSecurityCode

`func (o *CardDetails) SetEncryptedSecurityCode(v string)`

SetEncryptedSecurityCode sets EncryptedSecurityCode field to given value.

### HasEncryptedSecurityCode

`func (o *CardDetails) HasEncryptedSecurityCode() bool`

HasEncryptedSecurityCode returns a boolean if a field has been set.

### GetExpiryMonth

`func (o *CardDetails) GetExpiryMonth() string`

GetExpiryMonth returns the ExpiryMonth field if non-nil, zero value otherwise.

### GetExpiryMonthOk

`func (o *CardDetails) GetExpiryMonthOk() (*string, bool)`

GetExpiryMonthOk returns a tuple with the ExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryMonth

`func (o *CardDetails) SetExpiryMonth(v string)`

SetExpiryMonth sets ExpiryMonth field to given value.

### HasExpiryMonth

`func (o *CardDetails) HasExpiryMonth() bool`

HasExpiryMonth returns a boolean if a field has been set.

### GetExpiryYear

`func (o *CardDetails) GetExpiryYear() string`

GetExpiryYear returns the ExpiryYear field if non-nil, zero value otherwise.

### GetExpiryYearOk

`func (o *CardDetails) GetExpiryYearOk() (*string, bool)`

GetExpiryYearOk returns a tuple with the ExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryYear

`func (o *CardDetails) SetExpiryYear(v string)`

SetExpiryYear sets ExpiryYear field to given value.

### HasExpiryYear

`func (o *CardDetails) HasExpiryYear() bool`

HasExpiryYear returns a boolean if a field has been set.

### GetFundingSource

`func (o *CardDetails) GetFundingSource() string`

GetFundingSource returns the FundingSource field if non-nil, zero value otherwise.

### GetFundingSourceOk

`func (o *CardDetails) GetFundingSourceOk() (*string, bool)`

GetFundingSourceOk returns a tuple with the FundingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSource

`func (o *CardDetails) SetFundingSource(v string)`

SetFundingSource sets FundingSource field to given value.

### HasFundingSource

`func (o *CardDetails) HasFundingSource() bool`

HasFundingSource returns a boolean if a field has been set.

### GetHolderName

`func (o *CardDetails) GetHolderName() string`

GetHolderName returns the HolderName field if non-nil, zero value otherwise.

### GetHolderNameOk

`func (o *CardDetails) GetHolderNameOk() (*string, bool)`

GetHolderNameOk returns a tuple with the HolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderName

`func (o *CardDetails) SetHolderName(v string)`

SetHolderName sets HolderName field to given value.

### HasHolderName

`func (o *CardDetails) HasHolderName() bool`

HasHolderName returns a boolean if a field has been set.

### GetNetworkPaymentReference

`func (o *CardDetails) GetNetworkPaymentReference() string`

GetNetworkPaymentReference returns the NetworkPaymentReference field if non-nil, zero value otherwise.

### GetNetworkPaymentReferenceOk

`func (o *CardDetails) GetNetworkPaymentReferenceOk() (*string, bool)`

GetNetworkPaymentReferenceOk returns a tuple with the NetworkPaymentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPaymentReference

`func (o *CardDetails) SetNetworkPaymentReference(v string)`

SetNetworkPaymentReference sets NetworkPaymentReference field to given value.

### HasNetworkPaymentReference

`func (o *CardDetails) HasNetworkPaymentReference() bool`

HasNetworkPaymentReference returns a boolean if a field has been set.

### GetNumber

`func (o *CardDetails) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CardDetails) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CardDetails) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CardDetails) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetRecurringDetailReference

`func (o *CardDetails) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *CardDetails) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *CardDetails) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.

### HasRecurringDetailReference

`func (o *CardDetails) HasRecurringDetailReference() bool`

HasRecurringDetailReference returns a boolean if a field has been set.

### GetShopperNotificationReference

`func (o *CardDetails) GetShopperNotificationReference() string`

GetShopperNotificationReference returns the ShopperNotificationReference field if non-nil, zero value otherwise.

### GetShopperNotificationReferenceOk

`func (o *CardDetails) GetShopperNotificationReferenceOk() (*string, bool)`

GetShopperNotificationReferenceOk returns a tuple with the ShopperNotificationReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperNotificationReference

`func (o *CardDetails) SetShopperNotificationReference(v string)`

SetShopperNotificationReference sets ShopperNotificationReference field to given value.

### HasShopperNotificationReference

`func (o *CardDetails) HasShopperNotificationReference() bool`

HasShopperNotificationReference returns a boolean if a field has been set.

### GetStoredPaymentMethodId

`func (o *CardDetails) GetStoredPaymentMethodId() string`

GetStoredPaymentMethodId returns the StoredPaymentMethodId field if non-nil, zero value otherwise.

### GetStoredPaymentMethodIdOk

`func (o *CardDetails) GetStoredPaymentMethodIdOk() (*string, bool)`

GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredPaymentMethodId

`func (o *CardDetails) SetStoredPaymentMethodId(v string)`

SetStoredPaymentMethodId sets StoredPaymentMethodId field to given value.

### HasStoredPaymentMethodId

`func (o *CardDetails) HasStoredPaymentMethodId() bool`

HasStoredPaymentMethodId returns a boolean if a field has been set.

### GetThreeDS2SdkVersion

`func (o *CardDetails) GetThreeDS2SdkVersion() string`

GetThreeDS2SdkVersion returns the ThreeDS2SdkVersion field if non-nil, zero value otherwise.

### GetThreeDS2SdkVersionOk

`func (o *CardDetails) GetThreeDS2SdkVersionOk() (*string, bool)`

GetThreeDS2SdkVersionOk returns a tuple with the ThreeDS2SdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDS2SdkVersion

`func (o *CardDetails) SetThreeDS2SdkVersion(v string)`

SetThreeDS2SdkVersion sets ThreeDS2SdkVersion field to given value.

### HasThreeDS2SdkVersion

`func (o *CardDetails) HasThreeDS2SdkVersion() bool`

HasThreeDS2SdkVersion returns a boolean if a field has been set.

### GetType

`func (o *CardDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CardDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CardDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CardDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


