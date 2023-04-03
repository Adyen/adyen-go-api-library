# CheckoutPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccountNumber** | **string** | The bank account number (without separators). | 
**BankLocationId** | Pointer to **string** | The bank routing number of the account. | [optional] 
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**EncryptedBankAccountNumber** | Pointer to **string** | Encrypted bank account number. The bank account number (without separators). | [optional] 
**EncryptedBankLocationId** | Pointer to **string** | Encrypted location id. The bank routing number of the account. The field value is &#x60;nil&#x60; in most cases. | [optional] 
**OwnerName** | **string** | The name of the bank account holder. | 
**RecurringDetailReference** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**StoredPaymentMethodId** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**Type** | **string** | **zip** | [default to "zip"]
**BillingAddress** | Pointer to **string** | The address where to send the invoice. | [optional] 
**DeliveryAddress** | Pointer to **string** | The address where the goods should be delivered. | [optional] 
**PersonalDetails** | Pointer to **string** | Shopper name, date of birth, phone number, and email address. | [optional] 
**AmazonPayToken** | Pointer to **string** | This is the &#x60;amazonPayToken&#x60; that you obtained from the [Get Checkout Session](https://amazon-pay-acquirer-guide.s3-eu-west-1.amazonaws.com/v1/amazon-pay-api-v2/checkout-session.html#get-checkout-session) response. | [optional] 
**ApplePayToken** | **string** | The stringified and base64 encoded &#x60;paymentData&#x60; you retrieved from the Apple framework. | 
**FundingSource** | Pointer to **string** | The funding source that should be used when multiple sources are available. For Brazilian combo cards, by default the funding source is credit. To use debit, set this value to **debit**. | [optional] 
**HolderName** | Pointer to **string** | The name of the card holder. | [optional] 
**Issuer** | **string** | The shopper&#39;s bank. Specify this with the issuer value that corresponds to this bank. | 
**BlikCode** | Pointer to **string** | BLIK code consisting of 6 digits. | [optional] 
**Brand** | Pointer to **string** | Secondary brand of the card. For example: **plastix**, **hmclub**. | [optional] 
**CupsecureplusSmscode** | Pointer to **string** |  | [optional] 
**Cvc** | Pointer to **string** | The card verification code. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide). | [optional] 
**EncryptedCardNumber** | Pointer to **string** | The encrypted card number. | [optional] 
**EncryptedExpiryMonth** | Pointer to **string** | The encrypted card expiry month. | [optional] 
**EncryptedExpiryYear** | Pointer to **string** | The encrypted card expiry year. | [optional] 
**EncryptedSecurityCode** | Pointer to **string** | The encrypted card verification code. | [optional] 
**ExpiryMonth** | Pointer to **string** | The card expiry month. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide). | [optional] 
**ExpiryYear** | Pointer to **string** | The card expiry year. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide). | [optional] 
**NetworkPaymentReference** | Pointer to **string** | The network token reference. This is the [&#x60;networkTxReference&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_additionalData-ResponseAdditionalDataCommon-networkTxReference) from the response to the first payment. | [optional] 
**Number** | Pointer to **string** | The card number. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide). | [optional] 
**ShopperNotificationReference** | Pointer to **string** | The &#x60;shopperNotificationReference&#x60; returned in the response when you requested to notify the shopper. Used for recurring payment only. | [optional] 
**ThreeDS2SdkVersion** | Pointer to **string** | Version of the 3D Secure 2 mobile SDK. | [optional] 
**FirstName** | **string** | The shopper&#39;s first name. | 
**LastName** | **string** | The shopper&#39;s last name. | 
**ShopperEmail** | **string** |  | 
**TelephoneNumber** | **string** |  | 
**GooglePayToken** | **string** | The &#x60;token&#x60; that you obtained from the [Google Pay API](https://developers.google.com/pay/api/web/reference/response-objects#PaymentData) &#x60;PaymentData&#x60; response. | 
**MasterpassTransactionId** | **string** | The Masterpass transaction ID. | 
**OrderID** | Pointer to **string** | The unique ID associated with the order. | [optional] 
**PayerID** | Pointer to **string** | The unique ID associated with the payer. | [optional] 
**Subtype** | Pointer to **string** | The type of flow to initiate. | [optional] 
**VirtualPaymentAddress** | Pointer to **string** | The virtual payment address for UPI. | [optional] 
**SamsungPayToken** | **string** | The payload you received from the Samsung Pay SDK response. | 
**Iban** | **string** | The International Bank Account Number (IBAN). | 
**BillingSequenceNumber** | **string** | The sequence number for the debit. For example, send **2** if this is the second debit for the subscription. The sequence number is included in the notification sent to the shopper. | 
**VisaCheckoutCallId** | **string** | The Visa Click to Pay Call ID value. When your shopper selects a payment and/or a shipping address from Visa Click to Pay, you will receive a Visa Click to Pay Call ID. | 
**AppId** | Pointer to **string** |  | [optional] 
**Openid** | Pointer to **string** |  | [optional] 
**ClickAndCollect** | Pointer to **string** | Set this to **true** if the shopper would like to pick up and collect their order, instead of having the goods delivered to them. | [optional] 

## Methods

### NewCheckoutPaymentMethod

`func NewCheckoutPaymentMethod(bankAccountNumber string, ownerName string, type_ string, applePayToken string, issuer string, firstName string, lastName string, shopperEmail string, telephoneNumber string, googlePayToken string, masterpassTransactionId string, samsungPayToken string, iban string, billingSequenceNumber string, visaCheckoutCallId string, ) *CheckoutPaymentMethod`

NewCheckoutPaymentMethod instantiates a new CheckoutPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutPaymentMethodWithDefaults

`func NewCheckoutPaymentMethodWithDefaults() *CheckoutPaymentMethod`

NewCheckoutPaymentMethodWithDefaults instantiates a new CheckoutPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAccountNumber

`func (o *CheckoutPaymentMethod) GetBankAccountNumber() string`

GetBankAccountNumber returns the BankAccountNumber field if non-nil, zero value otherwise.

### GetBankAccountNumberOk

`func (o *CheckoutPaymentMethod) GetBankAccountNumberOk() (*string, bool)`

GetBankAccountNumberOk returns a tuple with the BankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountNumber

`func (o *CheckoutPaymentMethod) SetBankAccountNumber(v string)`

SetBankAccountNumber sets BankAccountNumber field to given value.


### GetBankLocationId

`func (o *CheckoutPaymentMethod) GetBankLocationId() string`

GetBankLocationId returns the BankLocationId field if non-nil, zero value otherwise.

### GetBankLocationIdOk

`func (o *CheckoutPaymentMethod) GetBankLocationIdOk() (*string, bool)`

GetBankLocationIdOk returns a tuple with the BankLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankLocationId

`func (o *CheckoutPaymentMethod) SetBankLocationId(v string)`

SetBankLocationId sets BankLocationId field to given value.

### HasBankLocationId

`func (o *CheckoutPaymentMethod) HasBankLocationId() bool`

HasBankLocationId returns a boolean if a field has been set.

### GetCheckoutAttemptId

`func (o *CheckoutPaymentMethod) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *CheckoutPaymentMethod) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *CheckoutPaymentMethod) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *CheckoutPaymentMethod) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetEncryptedBankAccountNumber

`func (o *CheckoutPaymentMethod) GetEncryptedBankAccountNumber() string`

GetEncryptedBankAccountNumber returns the EncryptedBankAccountNumber field if non-nil, zero value otherwise.

### GetEncryptedBankAccountNumberOk

`func (o *CheckoutPaymentMethod) GetEncryptedBankAccountNumberOk() (*string, bool)`

GetEncryptedBankAccountNumberOk returns a tuple with the EncryptedBankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedBankAccountNumber

`func (o *CheckoutPaymentMethod) SetEncryptedBankAccountNumber(v string)`

SetEncryptedBankAccountNumber sets EncryptedBankAccountNumber field to given value.

### HasEncryptedBankAccountNumber

`func (o *CheckoutPaymentMethod) HasEncryptedBankAccountNumber() bool`

HasEncryptedBankAccountNumber returns a boolean if a field has been set.

### GetEncryptedBankLocationId

`func (o *CheckoutPaymentMethod) GetEncryptedBankLocationId() string`

GetEncryptedBankLocationId returns the EncryptedBankLocationId field if non-nil, zero value otherwise.

### GetEncryptedBankLocationIdOk

`func (o *CheckoutPaymentMethod) GetEncryptedBankLocationIdOk() (*string, bool)`

GetEncryptedBankLocationIdOk returns a tuple with the EncryptedBankLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedBankLocationId

`func (o *CheckoutPaymentMethod) SetEncryptedBankLocationId(v string)`

SetEncryptedBankLocationId sets EncryptedBankLocationId field to given value.

### HasEncryptedBankLocationId

`func (o *CheckoutPaymentMethod) HasEncryptedBankLocationId() bool`

HasEncryptedBankLocationId returns a boolean if a field has been set.

### GetOwnerName

`func (o *CheckoutPaymentMethod) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *CheckoutPaymentMethod) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *CheckoutPaymentMethod) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.


### GetRecurringDetailReference

`func (o *CheckoutPaymentMethod) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *CheckoutPaymentMethod) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *CheckoutPaymentMethod) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.

### HasRecurringDetailReference

`func (o *CheckoutPaymentMethod) HasRecurringDetailReference() bool`

HasRecurringDetailReference returns a boolean if a field has been set.

### GetStoredPaymentMethodId

`func (o *CheckoutPaymentMethod) GetStoredPaymentMethodId() string`

GetStoredPaymentMethodId returns the StoredPaymentMethodId field if non-nil, zero value otherwise.

### GetStoredPaymentMethodIdOk

`func (o *CheckoutPaymentMethod) GetStoredPaymentMethodIdOk() (*string, bool)`

GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredPaymentMethodId

`func (o *CheckoutPaymentMethod) SetStoredPaymentMethodId(v string)`

SetStoredPaymentMethodId sets StoredPaymentMethodId field to given value.

### HasStoredPaymentMethodId

`func (o *CheckoutPaymentMethod) HasStoredPaymentMethodId() bool`

HasStoredPaymentMethodId returns a boolean if a field has been set.

### GetType

`func (o *CheckoutPaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutPaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutPaymentMethod) SetType(v string)`

SetType sets Type field to given value.


### GetBillingAddress

`func (o *CheckoutPaymentMethod) GetBillingAddress() string`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *CheckoutPaymentMethod) GetBillingAddressOk() (*string, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *CheckoutPaymentMethod) SetBillingAddress(v string)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *CheckoutPaymentMethod) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetDeliveryAddress

`func (o *CheckoutPaymentMethod) GetDeliveryAddress() string`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *CheckoutPaymentMethod) GetDeliveryAddressOk() (*string, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *CheckoutPaymentMethod) SetDeliveryAddress(v string)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *CheckoutPaymentMethod) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### GetPersonalDetails

`func (o *CheckoutPaymentMethod) GetPersonalDetails() string`

GetPersonalDetails returns the PersonalDetails field if non-nil, zero value otherwise.

### GetPersonalDetailsOk

`func (o *CheckoutPaymentMethod) GetPersonalDetailsOk() (*string, bool)`

GetPersonalDetailsOk returns a tuple with the PersonalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDetails

`func (o *CheckoutPaymentMethod) SetPersonalDetails(v string)`

SetPersonalDetails sets PersonalDetails field to given value.

### HasPersonalDetails

`func (o *CheckoutPaymentMethod) HasPersonalDetails() bool`

HasPersonalDetails returns a boolean if a field has been set.

### GetAmazonPayToken

`func (o *CheckoutPaymentMethod) GetAmazonPayToken() string`

GetAmazonPayToken returns the AmazonPayToken field if non-nil, zero value otherwise.

### GetAmazonPayTokenOk

`func (o *CheckoutPaymentMethod) GetAmazonPayTokenOk() (*string, bool)`

GetAmazonPayTokenOk returns a tuple with the AmazonPayToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonPayToken

`func (o *CheckoutPaymentMethod) SetAmazonPayToken(v string)`

SetAmazonPayToken sets AmazonPayToken field to given value.

### HasAmazonPayToken

`func (o *CheckoutPaymentMethod) HasAmazonPayToken() bool`

HasAmazonPayToken returns a boolean if a field has been set.

### GetApplePayToken

`func (o *CheckoutPaymentMethod) GetApplePayToken() string`

GetApplePayToken returns the ApplePayToken field if non-nil, zero value otherwise.

### GetApplePayTokenOk

`func (o *CheckoutPaymentMethod) GetApplePayTokenOk() (*string, bool)`

GetApplePayTokenOk returns a tuple with the ApplePayToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplePayToken

`func (o *CheckoutPaymentMethod) SetApplePayToken(v string)`

SetApplePayToken sets ApplePayToken field to given value.


### GetFundingSource

`func (o *CheckoutPaymentMethod) GetFundingSource() string`

GetFundingSource returns the FundingSource field if non-nil, zero value otherwise.

### GetFundingSourceOk

`func (o *CheckoutPaymentMethod) GetFundingSourceOk() (*string, bool)`

GetFundingSourceOk returns a tuple with the FundingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSource

`func (o *CheckoutPaymentMethod) SetFundingSource(v string)`

SetFundingSource sets FundingSource field to given value.

### HasFundingSource

`func (o *CheckoutPaymentMethod) HasFundingSource() bool`

HasFundingSource returns a boolean if a field has been set.

### GetHolderName

`func (o *CheckoutPaymentMethod) GetHolderName() string`

GetHolderName returns the HolderName field if non-nil, zero value otherwise.

### GetHolderNameOk

`func (o *CheckoutPaymentMethod) GetHolderNameOk() (*string, bool)`

GetHolderNameOk returns a tuple with the HolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderName

`func (o *CheckoutPaymentMethod) SetHolderName(v string)`

SetHolderName sets HolderName field to given value.

### HasHolderName

`func (o *CheckoutPaymentMethod) HasHolderName() bool`

HasHolderName returns a boolean if a field has been set.

### GetIssuer

`func (o *CheckoutPaymentMethod) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CheckoutPaymentMethod) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CheckoutPaymentMethod) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetBlikCode

`func (o *CheckoutPaymentMethod) GetBlikCode() string`

GetBlikCode returns the BlikCode field if non-nil, zero value otherwise.

### GetBlikCodeOk

`func (o *CheckoutPaymentMethod) GetBlikCodeOk() (*string, bool)`

GetBlikCodeOk returns a tuple with the BlikCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlikCode

`func (o *CheckoutPaymentMethod) SetBlikCode(v string)`

SetBlikCode sets BlikCode field to given value.

### HasBlikCode

`func (o *CheckoutPaymentMethod) HasBlikCode() bool`

HasBlikCode returns a boolean if a field has been set.

### GetBrand

`func (o *CheckoutPaymentMethod) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CheckoutPaymentMethod) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CheckoutPaymentMethod) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *CheckoutPaymentMethod) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetCupsecureplusSmscode

`func (o *CheckoutPaymentMethod) GetCupsecureplusSmscode() string`

GetCupsecureplusSmscode returns the CupsecureplusSmscode field if non-nil, zero value otherwise.

### GetCupsecureplusSmscodeOk

`func (o *CheckoutPaymentMethod) GetCupsecureplusSmscodeOk() (*string, bool)`

GetCupsecureplusSmscodeOk returns a tuple with the CupsecureplusSmscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCupsecureplusSmscode

`func (o *CheckoutPaymentMethod) SetCupsecureplusSmscode(v string)`

SetCupsecureplusSmscode sets CupsecureplusSmscode field to given value.

### HasCupsecureplusSmscode

`func (o *CheckoutPaymentMethod) HasCupsecureplusSmscode() bool`

HasCupsecureplusSmscode returns a boolean if a field has been set.

### GetCvc

`func (o *CheckoutPaymentMethod) GetCvc() string`

GetCvc returns the Cvc field if non-nil, zero value otherwise.

### GetCvcOk

`func (o *CheckoutPaymentMethod) GetCvcOk() (*string, bool)`

GetCvcOk returns a tuple with the Cvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvc

`func (o *CheckoutPaymentMethod) SetCvc(v string)`

SetCvc sets Cvc field to given value.

### HasCvc

`func (o *CheckoutPaymentMethod) HasCvc() bool`

HasCvc returns a boolean if a field has been set.

### GetEncryptedCardNumber

`func (o *CheckoutPaymentMethod) GetEncryptedCardNumber() string`

GetEncryptedCardNumber returns the EncryptedCardNumber field if non-nil, zero value otherwise.

### GetEncryptedCardNumberOk

`func (o *CheckoutPaymentMethod) GetEncryptedCardNumberOk() (*string, bool)`

GetEncryptedCardNumberOk returns a tuple with the EncryptedCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedCardNumber

`func (o *CheckoutPaymentMethod) SetEncryptedCardNumber(v string)`

SetEncryptedCardNumber sets EncryptedCardNumber field to given value.

### HasEncryptedCardNumber

`func (o *CheckoutPaymentMethod) HasEncryptedCardNumber() bool`

HasEncryptedCardNumber returns a boolean if a field has been set.

### GetEncryptedExpiryMonth

`func (o *CheckoutPaymentMethod) GetEncryptedExpiryMonth() string`

GetEncryptedExpiryMonth returns the EncryptedExpiryMonth field if non-nil, zero value otherwise.

### GetEncryptedExpiryMonthOk

`func (o *CheckoutPaymentMethod) GetEncryptedExpiryMonthOk() (*string, bool)`

GetEncryptedExpiryMonthOk returns a tuple with the EncryptedExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedExpiryMonth

`func (o *CheckoutPaymentMethod) SetEncryptedExpiryMonth(v string)`

SetEncryptedExpiryMonth sets EncryptedExpiryMonth field to given value.

### HasEncryptedExpiryMonth

`func (o *CheckoutPaymentMethod) HasEncryptedExpiryMonth() bool`

HasEncryptedExpiryMonth returns a boolean if a field has been set.

### GetEncryptedExpiryYear

`func (o *CheckoutPaymentMethod) GetEncryptedExpiryYear() string`

GetEncryptedExpiryYear returns the EncryptedExpiryYear field if non-nil, zero value otherwise.

### GetEncryptedExpiryYearOk

`func (o *CheckoutPaymentMethod) GetEncryptedExpiryYearOk() (*string, bool)`

GetEncryptedExpiryYearOk returns a tuple with the EncryptedExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedExpiryYear

`func (o *CheckoutPaymentMethod) SetEncryptedExpiryYear(v string)`

SetEncryptedExpiryYear sets EncryptedExpiryYear field to given value.

### HasEncryptedExpiryYear

`func (o *CheckoutPaymentMethod) HasEncryptedExpiryYear() bool`

HasEncryptedExpiryYear returns a boolean if a field has been set.

### GetEncryptedSecurityCode

`func (o *CheckoutPaymentMethod) GetEncryptedSecurityCode() string`

GetEncryptedSecurityCode returns the EncryptedSecurityCode field if non-nil, zero value otherwise.

### GetEncryptedSecurityCodeOk

`func (o *CheckoutPaymentMethod) GetEncryptedSecurityCodeOk() (*string, bool)`

GetEncryptedSecurityCodeOk returns a tuple with the EncryptedSecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedSecurityCode

`func (o *CheckoutPaymentMethod) SetEncryptedSecurityCode(v string)`

SetEncryptedSecurityCode sets EncryptedSecurityCode field to given value.

### HasEncryptedSecurityCode

`func (o *CheckoutPaymentMethod) HasEncryptedSecurityCode() bool`

HasEncryptedSecurityCode returns a boolean if a field has been set.

### GetExpiryMonth

`func (o *CheckoutPaymentMethod) GetExpiryMonth() string`

GetExpiryMonth returns the ExpiryMonth field if non-nil, zero value otherwise.

### GetExpiryMonthOk

`func (o *CheckoutPaymentMethod) GetExpiryMonthOk() (*string, bool)`

GetExpiryMonthOk returns a tuple with the ExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryMonth

`func (o *CheckoutPaymentMethod) SetExpiryMonth(v string)`

SetExpiryMonth sets ExpiryMonth field to given value.

### HasExpiryMonth

`func (o *CheckoutPaymentMethod) HasExpiryMonth() bool`

HasExpiryMonth returns a boolean if a field has been set.

### GetExpiryYear

`func (o *CheckoutPaymentMethod) GetExpiryYear() string`

GetExpiryYear returns the ExpiryYear field if non-nil, zero value otherwise.

### GetExpiryYearOk

`func (o *CheckoutPaymentMethod) GetExpiryYearOk() (*string, bool)`

GetExpiryYearOk returns a tuple with the ExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryYear

`func (o *CheckoutPaymentMethod) SetExpiryYear(v string)`

SetExpiryYear sets ExpiryYear field to given value.

### HasExpiryYear

`func (o *CheckoutPaymentMethod) HasExpiryYear() bool`

HasExpiryYear returns a boolean if a field has been set.

### GetNetworkPaymentReference

`func (o *CheckoutPaymentMethod) GetNetworkPaymentReference() string`

GetNetworkPaymentReference returns the NetworkPaymentReference field if non-nil, zero value otherwise.

### GetNetworkPaymentReferenceOk

`func (o *CheckoutPaymentMethod) GetNetworkPaymentReferenceOk() (*string, bool)`

GetNetworkPaymentReferenceOk returns a tuple with the NetworkPaymentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPaymentReference

`func (o *CheckoutPaymentMethod) SetNetworkPaymentReference(v string)`

SetNetworkPaymentReference sets NetworkPaymentReference field to given value.

### HasNetworkPaymentReference

`func (o *CheckoutPaymentMethod) HasNetworkPaymentReference() bool`

HasNetworkPaymentReference returns a boolean if a field has been set.

### GetNumber

`func (o *CheckoutPaymentMethod) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CheckoutPaymentMethod) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CheckoutPaymentMethod) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CheckoutPaymentMethod) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetShopperNotificationReference

`func (o *CheckoutPaymentMethod) GetShopperNotificationReference() string`

GetShopperNotificationReference returns the ShopperNotificationReference field if non-nil, zero value otherwise.

### GetShopperNotificationReferenceOk

`func (o *CheckoutPaymentMethod) GetShopperNotificationReferenceOk() (*string, bool)`

GetShopperNotificationReferenceOk returns a tuple with the ShopperNotificationReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperNotificationReference

`func (o *CheckoutPaymentMethod) SetShopperNotificationReference(v string)`

SetShopperNotificationReference sets ShopperNotificationReference field to given value.

### HasShopperNotificationReference

`func (o *CheckoutPaymentMethod) HasShopperNotificationReference() bool`

HasShopperNotificationReference returns a boolean if a field has been set.

### GetThreeDS2SdkVersion

`func (o *CheckoutPaymentMethod) GetThreeDS2SdkVersion() string`

GetThreeDS2SdkVersion returns the ThreeDS2SdkVersion field if non-nil, zero value otherwise.

### GetThreeDS2SdkVersionOk

`func (o *CheckoutPaymentMethod) GetThreeDS2SdkVersionOk() (*string, bool)`

GetThreeDS2SdkVersionOk returns a tuple with the ThreeDS2SdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDS2SdkVersion

`func (o *CheckoutPaymentMethod) SetThreeDS2SdkVersion(v string)`

SetThreeDS2SdkVersion sets ThreeDS2SdkVersion field to given value.

### HasThreeDS2SdkVersion

`func (o *CheckoutPaymentMethod) HasThreeDS2SdkVersion() bool`

HasThreeDS2SdkVersion returns a boolean if a field has been set.

### GetFirstName

`func (o *CheckoutPaymentMethod) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CheckoutPaymentMethod) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CheckoutPaymentMethod) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *CheckoutPaymentMethod) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CheckoutPaymentMethod) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CheckoutPaymentMethod) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetShopperEmail

`func (o *CheckoutPaymentMethod) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *CheckoutPaymentMethod) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *CheckoutPaymentMethod) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.


### GetTelephoneNumber

`func (o *CheckoutPaymentMethod) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *CheckoutPaymentMethod) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *CheckoutPaymentMethod) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.


### GetGooglePayToken

`func (o *CheckoutPaymentMethod) GetGooglePayToken() string`

GetGooglePayToken returns the GooglePayToken field if non-nil, zero value otherwise.

### GetGooglePayTokenOk

`func (o *CheckoutPaymentMethod) GetGooglePayTokenOk() (*string, bool)`

GetGooglePayTokenOk returns a tuple with the GooglePayToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePayToken

`func (o *CheckoutPaymentMethod) SetGooglePayToken(v string)`

SetGooglePayToken sets GooglePayToken field to given value.


### GetMasterpassTransactionId

`func (o *CheckoutPaymentMethod) GetMasterpassTransactionId() string`

GetMasterpassTransactionId returns the MasterpassTransactionId field if non-nil, zero value otherwise.

### GetMasterpassTransactionIdOk

`func (o *CheckoutPaymentMethod) GetMasterpassTransactionIdOk() (*string, bool)`

GetMasterpassTransactionIdOk returns a tuple with the MasterpassTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterpassTransactionId

`func (o *CheckoutPaymentMethod) SetMasterpassTransactionId(v string)`

SetMasterpassTransactionId sets MasterpassTransactionId field to given value.


### GetOrderID

`func (o *CheckoutPaymentMethod) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *CheckoutPaymentMethod) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *CheckoutPaymentMethod) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.

### HasOrderID

`func (o *CheckoutPaymentMethod) HasOrderID() bool`

HasOrderID returns a boolean if a field has been set.

### GetPayerID

`func (o *CheckoutPaymentMethod) GetPayerID() string`

GetPayerID returns the PayerID field if non-nil, zero value otherwise.

### GetPayerIDOk

`func (o *CheckoutPaymentMethod) GetPayerIDOk() (*string, bool)`

GetPayerIDOk returns a tuple with the PayerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerID

`func (o *CheckoutPaymentMethod) SetPayerID(v string)`

SetPayerID sets PayerID field to given value.

### HasPayerID

`func (o *CheckoutPaymentMethod) HasPayerID() bool`

HasPayerID returns a boolean if a field has been set.

### GetSubtype

`func (o *CheckoutPaymentMethod) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *CheckoutPaymentMethod) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *CheckoutPaymentMethod) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *CheckoutPaymentMethod) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetVirtualPaymentAddress

`func (o *CheckoutPaymentMethod) GetVirtualPaymentAddress() string`

GetVirtualPaymentAddress returns the VirtualPaymentAddress field if non-nil, zero value otherwise.

### GetVirtualPaymentAddressOk

`func (o *CheckoutPaymentMethod) GetVirtualPaymentAddressOk() (*string, bool)`

GetVirtualPaymentAddressOk returns a tuple with the VirtualPaymentAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPaymentAddress

`func (o *CheckoutPaymentMethod) SetVirtualPaymentAddress(v string)`

SetVirtualPaymentAddress sets VirtualPaymentAddress field to given value.

### HasVirtualPaymentAddress

`func (o *CheckoutPaymentMethod) HasVirtualPaymentAddress() bool`

HasVirtualPaymentAddress returns a boolean if a field has been set.

### GetSamsungPayToken

`func (o *CheckoutPaymentMethod) GetSamsungPayToken() string`

GetSamsungPayToken returns the SamsungPayToken field if non-nil, zero value otherwise.

### GetSamsungPayTokenOk

`func (o *CheckoutPaymentMethod) GetSamsungPayTokenOk() (*string, bool)`

GetSamsungPayTokenOk returns a tuple with the SamsungPayToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamsungPayToken

`func (o *CheckoutPaymentMethod) SetSamsungPayToken(v string)`

SetSamsungPayToken sets SamsungPayToken field to given value.


### GetIban

`func (o *CheckoutPaymentMethod) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *CheckoutPaymentMethod) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *CheckoutPaymentMethod) SetIban(v string)`

SetIban sets Iban field to given value.


### GetBillingSequenceNumber

`func (o *CheckoutPaymentMethod) GetBillingSequenceNumber() string`

GetBillingSequenceNumber returns the BillingSequenceNumber field if non-nil, zero value otherwise.

### GetBillingSequenceNumberOk

`func (o *CheckoutPaymentMethod) GetBillingSequenceNumberOk() (*string, bool)`

GetBillingSequenceNumberOk returns a tuple with the BillingSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSequenceNumber

`func (o *CheckoutPaymentMethod) SetBillingSequenceNumber(v string)`

SetBillingSequenceNumber sets BillingSequenceNumber field to given value.


### GetVisaCheckoutCallId

`func (o *CheckoutPaymentMethod) GetVisaCheckoutCallId() string`

GetVisaCheckoutCallId returns the VisaCheckoutCallId field if non-nil, zero value otherwise.

### GetVisaCheckoutCallIdOk

`func (o *CheckoutPaymentMethod) GetVisaCheckoutCallIdOk() (*string, bool)`

GetVisaCheckoutCallIdOk returns a tuple with the VisaCheckoutCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaCheckoutCallId

`func (o *CheckoutPaymentMethod) SetVisaCheckoutCallId(v string)`

SetVisaCheckoutCallId sets VisaCheckoutCallId field to given value.


### GetAppId

`func (o *CheckoutPaymentMethod) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *CheckoutPaymentMethod) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *CheckoutPaymentMethod) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *CheckoutPaymentMethod) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetOpenid

`func (o *CheckoutPaymentMethod) GetOpenid() string`

GetOpenid returns the Openid field if non-nil, zero value otherwise.

### GetOpenidOk

`func (o *CheckoutPaymentMethod) GetOpenidOk() (*string, bool)`

GetOpenidOk returns a tuple with the Openid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenid

`func (o *CheckoutPaymentMethod) SetOpenid(v string)`

SetOpenid sets Openid field to given value.

### HasOpenid

`func (o *CheckoutPaymentMethod) HasOpenid() bool`

HasOpenid returns a boolean if a field has been set.

### GetClickAndCollect

`func (o *CheckoutPaymentMethod) GetClickAndCollect() string`

GetClickAndCollect returns the ClickAndCollect field if non-nil, zero value otherwise.

### GetClickAndCollectOk

`func (o *CheckoutPaymentMethod) GetClickAndCollectOk() (*string, bool)`

GetClickAndCollectOk returns a tuple with the ClickAndCollect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickAndCollect

`func (o *CheckoutPaymentMethod) SetClickAndCollect(v string)`

SetClickAndCollect sets ClickAndCollect field to given value.

### HasClickAndCollect

`func (o *CheckoutPaymentMethod) HasClickAndCollect() bool`

HasClickAndCollect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


