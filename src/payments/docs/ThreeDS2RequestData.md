# ThreeDS2RequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctInfo** | Pointer to [**AcctInfo**](AcctInfo.md) |  | [optional] 
**AcctType** | Pointer to **string** | Indicates the type of account. For example, for a multi-account card product. Length: 2 characters. Allowed values: * **01** — Not applicable * **02** — Credit * **03** — Debit | [optional] 
**AcquirerBIN** | Pointer to **string** | Required for [authentication-only integration](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only). The acquiring BIN enrolled for 3D Secure 2. This string should match the value that you will use in the authorisation. Use 123456 on the Test platform. | [optional] 
**AcquirerMerchantID** | Pointer to **string** | Required for [authentication-only integration](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only). The merchantId that is enrolled for 3D Secure 2 by the merchant&#39;s acquirer. This string should match the value that you will use in the authorisation. Use 123456 on the Test platform. | [optional] 
**AddrMatch** | Pointer to **string** | Indicates whether the Cardholder Shipping Address and Cardholder Billing Address are the same. Allowed values: * **Y** — Shipping Address matches Billing Address. * **N** — Shipping Address does not match Billing Address. | [optional] 
**AuthenticationOnly** | Pointer to **bool** | If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation. | [optional] [default to false]
**ChallengeIndicator** | Pointer to **string** | Possibility to specify a preference for receiving a challenge from the issuer. Allowed values: * &#x60;noPreference&#x60; * &#x60;requestNoChallenge&#x60; * &#x60;requestChallenge&#x60; * &#x60;requestChallengeAsMandate&#x60;  | [optional] 
**DeviceChannel** | **string** | The environment of the shopper. Allowed values: * &#x60;app&#x60; * &#x60;browser&#x60; | 
**DeviceRenderOptions** | Pointer to [**DeviceRenderOptions**](DeviceRenderOptions.md) |  | [optional] 
**HomePhone** | Pointer to [**Phone**](Phone.md) |  | [optional] 
**Mcc** | Pointer to **string** | Required for merchants that have been enrolled for 3D Secure 2 by another party than Adyen, mostly [authentication-only integrations](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only). The &#x60;mcc&#x60; is a four-digit code with which the previously given &#x60;acquirerMerchantID&#x60; is registered at the scheme. | [optional] 
**MerchantName** | Pointer to **string** | Required for [authentication-only integration](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only). The merchant name that the issuer presents to the shopper if they get a challenge. We recommend to use the same value that you will use in the authorization. Maximum length is 40 characters. &gt; Optional for a [full 3D Secure 2 integration](https://docs.adyen.com/online-payments/3d-secure/native-3ds2/api-integration). Use this field if you are enrolled for 3D Secure 2 with us and want to override the merchant name already configured on your account. | [optional] 
**MessageVersion** | Pointer to **string** | The &#x60;messageVersion&#x60; value indicating the 3D Secure 2 protocol version. | [optional] [default to "2.1.0"]
**MobilePhone** | Pointer to [**Phone**](Phone.md) |  | [optional] 
**NotificationURL** | Pointer to **string** | URL to where the issuer should send the &#x60;CRes&#x60;. Required if you are not using components for &#x60;channel&#x60; **Web** or if you are using classic integration &#x60;deviceChannel&#x60; **browser**. | [optional] 
**PayTokenInd** | Pointer to **bool** | Value **true** indicates that the transaction was de-tokenised prior to being received by the ACS. | [optional] 
**PaymentAuthenticationUseCase** | Pointer to **string** | Indicates the type of payment for which an authentication is requested (message extension) | [optional] 
**PurchaseInstalData** | Pointer to **string** | Indicates the maximum number of authorisations permitted for instalment payments. Length: 1–3 characters. | [optional] 
**RecurringExpiry** | Pointer to **string** | Date after which no further authorisations shall be performed. Format: YYYYMMDD | [optional] 
**RecurringFrequency** | Pointer to **string** | Indicates the minimum number of days between authorisations. Maximum length: 4 characters. | [optional] 
**SdkAppID** | Pointer to **string** | The &#x60;sdkAppID&#x60; value as received from the 3D Secure 2 SDK. Required for &#x60;deviceChannel&#x60; set to **app**. | [optional] 
**SdkEncData** | Pointer to **string** | The &#x60;sdkEncData&#x60; value as received from the 3D Secure 2 SDK. Required for &#x60;deviceChannel&#x60; set to **app**. | [optional] 
**SdkEphemPubKey** | Pointer to [**SDKEphemPubKey**](SDKEphemPubKey.md) |  | [optional] 
**SdkMaxTimeout** | Pointer to **int32** | The maximum amount of time in minutes for the 3D Secure 2 authentication process. Optional and only for &#x60;deviceChannel&#x60; set to **app**. Defaults to **60** minutes. | [optional] [default to 60]
**SdkReferenceNumber** | Pointer to **string** | The &#x60;sdkReferenceNumber&#x60; value as received from the 3D Secure 2 SDK. Only for &#x60;deviceChannel&#x60; set to **app**. | [optional] 
**SdkTransID** | Pointer to **string** | The &#x60;sdkTransID&#x60; value as received from the 3D Secure 2 SDK. Only for &#x60;deviceChannel&#x60; set to **app**. | [optional] 
**SdkVersion** | Pointer to **string** | Version of the 3D Secure 2 mobile SDK.  Only for &#x60;deviceChannel&#x60; set to **app**. | [optional] 
**ThreeDSCompInd** | Pointer to **string** | Completion indicator for the device fingerprinting. | [optional] 
**ThreeDSRequestorAuthenticationInd** | Pointer to **string** | Indicates the type of Authentication request. | [optional] 
**ThreeDSRequestorAuthenticationInfo** | Pointer to [**ThreeDSRequestorAuthenticationInfo**](ThreeDSRequestorAuthenticationInfo.md) |  | [optional] 
**ThreeDSRequestorChallengeInd** | Pointer to **string** | Indicates whether a challenge is requested for this transaction. Possible values: * **01** — No preference * **02** — No challenge requested * **03** — Challenge requested (3DS Requestor preference) * **04** — Challenge requested (Mandate) * **05** — No challenge (transactional risk analysis is already performed) * **06** — Data Only | [optional] 
**ThreeDSRequestorID** | Pointer to **string** | Required for [authentication-only integration](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only) for Visa. Unique 3D Secure requestor identifier assigned by the Directory Server when you enrol for 3D Secure 2. | [optional] 
**ThreeDSRequestorName** | Pointer to **string** | Required for [authentication-only integration](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only) for Visa. Unique 3D Secure requestor name assigned by the Directory Server when you enrol for 3D Secure 2. | [optional] 
**ThreeDSRequestorPriorAuthenticationInfo** | Pointer to [**ThreeDSRequestorPriorAuthenticationInfo**](ThreeDSRequestorPriorAuthenticationInfo.md) |  | [optional] 
**ThreeDSRequestorURL** | Pointer to **string** | URL of the (customer service) website that will be shown to the shopper in case of technical errors during the 3D Secure 2 process. | [optional] 
**TransType** | Pointer to **string** | Identifies the type of transaction being authenticated. Length: 2 characters. Allowed values: * **01** — Goods/Service Purchase * **03** — Check Acceptance * **10** — Account Funding * **11** — Quasi-Cash Transaction * **28** — Prepaid Activation and Load | [optional] 
**TransactionType** | Pointer to **string** | Identify the type of the transaction being authenticated. | [optional] 
**WhiteListStatus** | Pointer to **string** | The &#x60;whiteListStatus&#x60; value returned from a previous 3D Secure 2 transaction, only applicable for 3D Secure 2 protocol version 2.2.0. | [optional] 
**WorkPhone** | Pointer to [**Phone**](Phone.md) |  | [optional] 

## Methods

### NewThreeDS2RequestData

`func NewThreeDS2RequestData(deviceChannel string, ) *ThreeDS2RequestData`

NewThreeDS2RequestData instantiates a new ThreeDS2RequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDS2RequestDataWithDefaults

`func NewThreeDS2RequestDataWithDefaults() *ThreeDS2RequestData`

NewThreeDS2RequestDataWithDefaults instantiates a new ThreeDS2RequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctInfo

`func (o *ThreeDS2RequestData) GetAcctInfo() AcctInfo`

GetAcctInfo returns the AcctInfo field if non-nil, zero value otherwise.

### GetAcctInfoOk

`func (o *ThreeDS2RequestData) GetAcctInfoOk() (*AcctInfo, bool)`

GetAcctInfoOk returns a tuple with the AcctInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctInfo

`func (o *ThreeDS2RequestData) SetAcctInfo(v AcctInfo)`

SetAcctInfo sets AcctInfo field to given value.

### HasAcctInfo

`func (o *ThreeDS2RequestData) HasAcctInfo() bool`

HasAcctInfo returns a boolean if a field has been set.

### GetAcctType

`func (o *ThreeDS2RequestData) GetAcctType() string`

GetAcctType returns the AcctType field if non-nil, zero value otherwise.

### GetAcctTypeOk

`func (o *ThreeDS2RequestData) GetAcctTypeOk() (*string, bool)`

GetAcctTypeOk returns a tuple with the AcctType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctType

`func (o *ThreeDS2RequestData) SetAcctType(v string)`

SetAcctType sets AcctType field to given value.

### HasAcctType

`func (o *ThreeDS2RequestData) HasAcctType() bool`

HasAcctType returns a boolean if a field has been set.

### GetAcquirerBIN

`func (o *ThreeDS2RequestData) GetAcquirerBIN() string`

GetAcquirerBIN returns the AcquirerBIN field if non-nil, zero value otherwise.

### GetAcquirerBINOk

`func (o *ThreeDS2RequestData) GetAcquirerBINOk() (*string, bool)`

GetAcquirerBINOk returns a tuple with the AcquirerBIN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerBIN

`func (o *ThreeDS2RequestData) SetAcquirerBIN(v string)`

SetAcquirerBIN sets AcquirerBIN field to given value.

### HasAcquirerBIN

`func (o *ThreeDS2RequestData) HasAcquirerBIN() bool`

HasAcquirerBIN returns a boolean if a field has been set.

### GetAcquirerMerchantID

`func (o *ThreeDS2RequestData) GetAcquirerMerchantID() string`

GetAcquirerMerchantID returns the AcquirerMerchantID field if non-nil, zero value otherwise.

### GetAcquirerMerchantIDOk

`func (o *ThreeDS2RequestData) GetAcquirerMerchantIDOk() (*string, bool)`

GetAcquirerMerchantIDOk returns a tuple with the AcquirerMerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerMerchantID

`func (o *ThreeDS2RequestData) SetAcquirerMerchantID(v string)`

SetAcquirerMerchantID sets AcquirerMerchantID field to given value.

### HasAcquirerMerchantID

`func (o *ThreeDS2RequestData) HasAcquirerMerchantID() bool`

HasAcquirerMerchantID returns a boolean if a field has been set.

### GetAddrMatch

`func (o *ThreeDS2RequestData) GetAddrMatch() string`

GetAddrMatch returns the AddrMatch field if non-nil, zero value otherwise.

### GetAddrMatchOk

`func (o *ThreeDS2RequestData) GetAddrMatchOk() (*string, bool)`

GetAddrMatchOk returns a tuple with the AddrMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrMatch

`func (o *ThreeDS2RequestData) SetAddrMatch(v string)`

SetAddrMatch sets AddrMatch field to given value.

### HasAddrMatch

`func (o *ThreeDS2RequestData) HasAddrMatch() bool`

HasAddrMatch returns a boolean if a field has been set.

### GetAuthenticationOnly

`func (o *ThreeDS2RequestData) GetAuthenticationOnly() bool`

GetAuthenticationOnly returns the AuthenticationOnly field if non-nil, zero value otherwise.

### GetAuthenticationOnlyOk

`func (o *ThreeDS2RequestData) GetAuthenticationOnlyOk() (*bool, bool)`

GetAuthenticationOnlyOk returns a tuple with the AuthenticationOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOnly

`func (o *ThreeDS2RequestData) SetAuthenticationOnly(v bool)`

SetAuthenticationOnly sets AuthenticationOnly field to given value.

### HasAuthenticationOnly

`func (o *ThreeDS2RequestData) HasAuthenticationOnly() bool`

HasAuthenticationOnly returns a boolean if a field has been set.

### GetChallengeIndicator

`func (o *ThreeDS2RequestData) GetChallengeIndicator() string`

GetChallengeIndicator returns the ChallengeIndicator field if non-nil, zero value otherwise.

### GetChallengeIndicatorOk

`func (o *ThreeDS2RequestData) GetChallengeIndicatorOk() (*string, bool)`

GetChallengeIndicatorOk returns a tuple with the ChallengeIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeIndicator

`func (o *ThreeDS2RequestData) SetChallengeIndicator(v string)`

SetChallengeIndicator sets ChallengeIndicator field to given value.

### HasChallengeIndicator

`func (o *ThreeDS2RequestData) HasChallengeIndicator() bool`

HasChallengeIndicator returns a boolean if a field has been set.

### GetDeviceChannel

`func (o *ThreeDS2RequestData) GetDeviceChannel() string`

GetDeviceChannel returns the DeviceChannel field if non-nil, zero value otherwise.

### GetDeviceChannelOk

`func (o *ThreeDS2RequestData) GetDeviceChannelOk() (*string, bool)`

GetDeviceChannelOk returns a tuple with the DeviceChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceChannel

`func (o *ThreeDS2RequestData) SetDeviceChannel(v string)`

SetDeviceChannel sets DeviceChannel field to given value.


### GetDeviceRenderOptions

`func (o *ThreeDS2RequestData) GetDeviceRenderOptions() DeviceRenderOptions`

GetDeviceRenderOptions returns the DeviceRenderOptions field if non-nil, zero value otherwise.

### GetDeviceRenderOptionsOk

`func (o *ThreeDS2RequestData) GetDeviceRenderOptionsOk() (*DeviceRenderOptions, bool)`

GetDeviceRenderOptionsOk returns a tuple with the DeviceRenderOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRenderOptions

`func (o *ThreeDS2RequestData) SetDeviceRenderOptions(v DeviceRenderOptions)`

SetDeviceRenderOptions sets DeviceRenderOptions field to given value.

### HasDeviceRenderOptions

`func (o *ThreeDS2RequestData) HasDeviceRenderOptions() bool`

HasDeviceRenderOptions returns a boolean if a field has been set.

### GetHomePhone

`func (o *ThreeDS2RequestData) GetHomePhone() Phone`

GetHomePhone returns the HomePhone field if non-nil, zero value otherwise.

### GetHomePhoneOk

`func (o *ThreeDS2RequestData) GetHomePhoneOk() (*Phone, bool)`

GetHomePhoneOk returns a tuple with the HomePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePhone

`func (o *ThreeDS2RequestData) SetHomePhone(v Phone)`

SetHomePhone sets HomePhone field to given value.

### HasHomePhone

`func (o *ThreeDS2RequestData) HasHomePhone() bool`

HasHomePhone returns a boolean if a field has been set.

### GetMcc

`func (o *ThreeDS2RequestData) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *ThreeDS2RequestData) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *ThreeDS2RequestData) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *ThreeDS2RequestData) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMerchantName

`func (o *ThreeDS2RequestData) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *ThreeDS2RequestData) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *ThreeDS2RequestData) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *ThreeDS2RequestData) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### GetMessageVersion

`func (o *ThreeDS2RequestData) GetMessageVersion() string`

GetMessageVersion returns the MessageVersion field if non-nil, zero value otherwise.

### GetMessageVersionOk

`func (o *ThreeDS2RequestData) GetMessageVersionOk() (*string, bool)`

GetMessageVersionOk returns a tuple with the MessageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageVersion

`func (o *ThreeDS2RequestData) SetMessageVersion(v string)`

SetMessageVersion sets MessageVersion field to given value.

### HasMessageVersion

`func (o *ThreeDS2RequestData) HasMessageVersion() bool`

HasMessageVersion returns a boolean if a field has been set.

### GetMobilePhone

`func (o *ThreeDS2RequestData) GetMobilePhone() Phone`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *ThreeDS2RequestData) GetMobilePhoneOk() (*Phone, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *ThreeDS2RequestData) SetMobilePhone(v Phone)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *ThreeDS2RequestData) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetNotificationURL

`func (o *ThreeDS2RequestData) GetNotificationURL() string`

GetNotificationURL returns the NotificationURL field if non-nil, zero value otherwise.

### GetNotificationURLOk

`func (o *ThreeDS2RequestData) GetNotificationURLOk() (*string, bool)`

GetNotificationURLOk returns a tuple with the NotificationURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationURL

`func (o *ThreeDS2RequestData) SetNotificationURL(v string)`

SetNotificationURL sets NotificationURL field to given value.

### HasNotificationURL

`func (o *ThreeDS2RequestData) HasNotificationURL() bool`

HasNotificationURL returns a boolean if a field has been set.

### GetPayTokenInd

`func (o *ThreeDS2RequestData) GetPayTokenInd() bool`

GetPayTokenInd returns the PayTokenInd field if non-nil, zero value otherwise.

### GetPayTokenIndOk

`func (o *ThreeDS2RequestData) GetPayTokenIndOk() (*bool, bool)`

GetPayTokenIndOk returns a tuple with the PayTokenInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayTokenInd

`func (o *ThreeDS2RequestData) SetPayTokenInd(v bool)`

SetPayTokenInd sets PayTokenInd field to given value.

### HasPayTokenInd

`func (o *ThreeDS2RequestData) HasPayTokenInd() bool`

HasPayTokenInd returns a boolean if a field has been set.

### GetPaymentAuthenticationUseCase

`func (o *ThreeDS2RequestData) GetPaymentAuthenticationUseCase() string`

GetPaymentAuthenticationUseCase returns the PaymentAuthenticationUseCase field if non-nil, zero value otherwise.

### GetPaymentAuthenticationUseCaseOk

`func (o *ThreeDS2RequestData) GetPaymentAuthenticationUseCaseOk() (*string, bool)`

GetPaymentAuthenticationUseCaseOk returns a tuple with the PaymentAuthenticationUseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAuthenticationUseCase

`func (o *ThreeDS2RequestData) SetPaymentAuthenticationUseCase(v string)`

SetPaymentAuthenticationUseCase sets PaymentAuthenticationUseCase field to given value.

### HasPaymentAuthenticationUseCase

`func (o *ThreeDS2RequestData) HasPaymentAuthenticationUseCase() bool`

HasPaymentAuthenticationUseCase returns a boolean if a field has been set.

### GetPurchaseInstalData

`func (o *ThreeDS2RequestData) GetPurchaseInstalData() string`

GetPurchaseInstalData returns the PurchaseInstalData field if non-nil, zero value otherwise.

### GetPurchaseInstalDataOk

`func (o *ThreeDS2RequestData) GetPurchaseInstalDataOk() (*string, bool)`

GetPurchaseInstalDataOk returns a tuple with the PurchaseInstalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseInstalData

`func (o *ThreeDS2RequestData) SetPurchaseInstalData(v string)`

SetPurchaseInstalData sets PurchaseInstalData field to given value.

### HasPurchaseInstalData

`func (o *ThreeDS2RequestData) HasPurchaseInstalData() bool`

HasPurchaseInstalData returns a boolean if a field has been set.

### GetRecurringExpiry

`func (o *ThreeDS2RequestData) GetRecurringExpiry() string`

GetRecurringExpiry returns the RecurringExpiry field if non-nil, zero value otherwise.

### GetRecurringExpiryOk

`func (o *ThreeDS2RequestData) GetRecurringExpiryOk() (*string, bool)`

GetRecurringExpiryOk returns a tuple with the RecurringExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringExpiry

`func (o *ThreeDS2RequestData) SetRecurringExpiry(v string)`

SetRecurringExpiry sets RecurringExpiry field to given value.

### HasRecurringExpiry

`func (o *ThreeDS2RequestData) HasRecurringExpiry() bool`

HasRecurringExpiry returns a boolean if a field has been set.

### GetRecurringFrequency

`func (o *ThreeDS2RequestData) GetRecurringFrequency() string`

GetRecurringFrequency returns the RecurringFrequency field if non-nil, zero value otherwise.

### GetRecurringFrequencyOk

`func (o *ThreeDS2RequestData) GetRecurringFrequencyOk() (*string, bool)`

GetRecurringFrequencyOk returns a tuple with the RecurringFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringFrequency

`func (o *ThreeDS2RequestData) SetRecurringFrequency(v string)`

SetRecurringFrequency sets RecurringFrequency field to given value.

### HasRecurringFrequency

`func (o *ThreeDS2RequestData) HasRecurringFrequency() bool`

HasRecurringFrequency returns a boolean if a field has been set.

### GetSdkAppID

`func (o *ThreeDS2RequestData) GetSdkAppID() string`

GetSdkAppID returns the SdkAppID field if non-nil, zero value otherwise.

### GetSdkAppIDOk

`func (o *ThreeDS2RequestData) GetSdkAppIDOk() (*string, bool)`

GetSdkAppIDOk returns a tuple with the SdkAppID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkAppID

`func (o *ThreeDS2RequestData) SetSdkAppID(v string)`

SetSdkAppID sets SdkAppID field to given value.

### HasSdkAppID

`func (o *ThreeDS2RequestData) HasSdkAppID() bool`

HasSdkAppID returns a boolean if a field has been set.

### GetSdkEncData

`func (o *ThreeDS2RequestData) GetSdkEncData() string`

GetSdkEncData returns the SdkEncData field if non-nil, zero value otherwise.

### GetSdkEncDataOk

`func (o *ThreeDS2RequestData) GetSdkEncDataOk() (*string, bool)`

GetSdkEncDataOk returns a tuple with the SdkEncData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkEncData

`func (o *ThreeDS2RequestData) SetSdkEncData(v string)`

SetSdkEncData sets SdkEncData field to given value.

### HasSdkEncData

`func (o *ThreeDS2RequestData) HasSdkEncData() bool`

HasSdkEncData returns a boolean if a field has been set.

### GetSdkEphemPubKey

`func (o *ThreeDS2RequestData) GetSdkEphemPubKey() SDKEphemPubKey`

GetSdkEphemPubKey returns the SdkEphemPubKey field if non-nil, zero value otherwise.

### GetSdkEphemPubKeyOk

`func (o *ThreeDS2RequestData) GetSdkEphemPubKeyOk() (*SDKEphemPubKey, bool)`

GetSdkEphemPubKeyOk returns a tuple with the SdkEphemPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkEphemPubKey

`func (o *ThreeDS2RequestData) SetSdkEphemPubKey(v SDKEphemPubKey)`

SetSdkEphemPubKey sets SdkEphemPubKey field to given value.

### HasSdkEphemPubKey

`func (o *ThreeDS2RequestData) HasSdkEphemPubKey() bool`

HasSdkEphemPubKey returns a boolean if a field has been set.

### GetSdkMaxTimeout

`func (o *ThreeDS2RequestData) GetSdkMaxTimeout() int32`

GetSdkMaxTimeout returns the SdkMaxTimeout field if non-nil, zero value otherwise.

### GetSdkMaxTimeoutOk

`func (o *ThreeDS2RequestData) GetSdkMaxTimeoutOk() (*int32, bool)`

GetSdkMaxTimeoutOk returns a tuple with the SdkMaxTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkMaxTimeout

`func (o *ThreeDS2RequestData) SetSdkMaxTimeout(v int32)`

SetSdkMaxTimeout sets SdkMaxTimeout field to given value.

### HasSdkMaxTimeout

`func (o *ThreeDS2RequestData) HasSdkMaxTimeout() bool`

HasSdkMaxTimeout returns a boolean if a field has been set.

### GetSdkReferenceNumber

`func (o *ThreeDS2RequestData) GetSdkReferenceNumber() string`

GetSdkReferenceNumber returns the SdkReferenceNumber field if non-nil, zero value otherwise.

### GetSdkReferenceNumberOk

`func (o *ThreeDS2RequestData) GetSdkReferenceNumberOk() (*string, bool)`

GetSdkReferenceNumberOk returns a tuple with the SdkReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkReferenceNumber

`func (o *ThreeDS2RequestData) SetSdkReferenceNumber(v string)`

SetSdkReferenceNumber sets SdkReferenceNumber field to given value.

### HasSdkReferenceNumber

`func (o *ThreeDS2RequestData) HasSdkReferenceNumber() bool`

HasSdkReferenceNumber returns a boolean if a field has been set.

### GetSdkTransID

`func (o *ThreeDS2RequestData) GetSdkTransID() string`

GetSdkTransID returns the SdkTransID field if non-nil, zero value otherwise.

### GetSdkTransIDOk

`func (o *ThreeDS2RequestData) GetSdkTransIDOk() (*string, bool)`

GetSdkTransIDOk returns a tuple with the SdkTransID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkTransID

`func (o *ThreeDS2RequestData) SetSdkTransID(v string)`

SetSdkTransID sets SdkTransID field to given value.

### HasSdkTransID

`func (o *ThreeDS2RequestData) HasSdkTransID() bool`

HasSdkTransID returns a boolean if a field has been set.

### GetSdkVersion

`func (o *ThreeDS2RequestData) GetSdkVersion() string`

GetSdkVersion returns the SdkVersion field if non-nil, zero value otherwise.

### GetSdkVersionOk

`func (o *ThreeDS2RequestData) GetSdkVersionOk() (*string, bool)`

GetSdkVersionOk returns a tuple with the SdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkVersion

`func (o *ThreeDS2RequestData) SetSdkVersion(v string)`

SetSdkVersion sets SdkVersion field to given value.

### HasSdkVersion

`func (o *ThreeDS2RequestData) HasSdkVersion() bool`

HasSdkVersion returns a boolean if a field has been set.

### GetThreeDSCompInd

`func (o *ThreeDS2RequestData) GetThreeDSCompInd() string`

GetThreeDSCompInd returns the ThreeDSCompInd field if non-nil, zero value otherwise.

### GetThreeDSCompIndOk

`func (o *ThreeDS2RequestData) GetThreeDSCompIndOk() (*string, bool)`

GetThreeDSCompIndOk returns a tuple with the ThreeDSCompInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSCompInd

`func (o *ThreeDS2RequestData) SetThreeDSCompInd(v string)`

SetThreeDSCompInd sets ThreeDSCompInd field to given value.

### HasThreeDSCompInd

`func (o *ThreeDS2RequestData) HasThreeDSCompInd() bool`

HasThreeDSCompInd returns a boolean if a field has been set.

### GetThreeDSRequestorAuthenticationInd

`func (o *ThreeDS2RequestData) GetThreeDSRequestorAuthenticationInd() string`

GetThreeDSRequestorAuthenticationInd returns the ThreeDSRequestorAuthenticationInd field if non-nil, zero value otherwise.

### GetThreeDSRequestorAuthenticationIndOk

`func (o *ThreeDS2RequestData) GetThreeDSRequestorAuthenticationIndOk() (*string, bool)`

GetThreeDSRequestorAuthenticationIndOk returns a tuple with the ThreeDSRequestorAuthenticationInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSRequestorAuthenticationInd

`func (o *ThreeDS2RequestData) SetThreeDSRequestorAuthenticationInd(v string)`

SetThreeDSRequestorAuthenticationInd sets ThreeDSRequestorAuthenticationInd field to given value.

### HasThreeDSRequestorAuthenticationInd

`func (o *ThreeDS2RequestData) HasThreeDSRequestorAuthenticationInd() bool`

HasThreeDSRequestorAuthenticationInd returns a boolean if a field has been set.

### GetThreeDSRequestorAuthenticationInfo

`func (o *ThreeDS2RequestData) GetThreeDSRequestorAuthenticationInfo() ThreeDSRequestorAuthenticationInfo`

GetThreeDSRequestorAuthenticationInfo returns the ThreeDSRequestorAuthenticationInfo field if non-nil, zero value otherwise.

### GetThreeDSRequestorAuthenticationInfoOk

`func (o *ThreeDS2RequestData) GetThreeDSRequestorAuthenticationInfoOk() (*ThreeDSRequestorAuthenticationInfo, bool)`

GetThreeDSRequestorAuthenticationInfoOk returns a tuple with the ThreeDSRequestorAuthenticationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSRequestorAuthenticationInfo

`func (o *ThreeDS2RequestData) SetThreeDSRequestorAuthenticationInfo(v ThreeDSRequestorAuthenticationInfo)`

SetThreeDSRequestorAuthenticationInfo sets ThreeDSRequestorAuthenticationInfo field to given value.

### HasThreeDSRequestorAuthenticationInfo

`func (o *ThreeDS2RequestData) HasThreeDSRequestorAuthenticationInfo() bool`

HasThreeDSRequestorAuthenticationInfo returns a boolean if a field has been set.

### GetThreeDSRequestorChallengeInd

`func (o *ThreeDS2RequestData) GetThreeDSRequestorChallengeInd() string`

GetThreeDSRequestorChallengeInd returns the ThreeDSRequestorChallengeInd field if non-nil, zero value otherwise.

### GetThreeDSRequestorChallengeIndOk

`func (o *ThreeDS2RequestData) GetThreeDSRequestorChallengeIndOk() (*string, bool)`

GetThreeDSRequestorChallengeIndOk returns a tuple with the ThreeDSRequestorChallengeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSRequestorChallengeInd

`func (o *ThreeDS2RequestData) SetThreeDSRequestorChallengeInd(v string)`

SetThreeDSRequestorChallengeInd sets ThreeDSRequestorChallengeInd field to given value.

### HasThreeDSRequestorChallengeInd

`func (o *ThreeDS2RequestData) HasThreeDSRequestorChallengeInd() bool`

HasThreeDSRequestorChallengeInd returns a boolean if a field has been set.

### GetThreeDSRequestorID

`func (o *ThreeDS2RequestData) GetThreeDSRequestorID() string`

GetThreeDSRequestorID returns the ThreeDSRequestorID field if non-nil, zero value otherwise.

### GetThreeDSRequestorIDOk

`func (o *ThreeDS2RequestData) GetThreeDSRequestorIDOk() (*string, bool)`

GetThreeDSRequestorIDOk returns a tuple with the ThreeDSRequestorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSRequestorID

`func (o *ThreeDS2RequestData) SetThreeDSRequestorID(v string)`

SetThreeDSRequestorID sets ThreeDSRequestorID field to given value.

### HasThreeDSRequestorID

`func (o *ThreeDS2RequestData) HasThreeDSRequestorID() bool`

HasThreeDSRequestorID returns a boolean if a field has been set.

### GetThreeDSRequestorName

`func (o *ThreeDS2RequestData) GetThreeDSRequestorName() string`

GetThreeDSRequestorName returns the ThreeDSRequestorName field if non-nil, zero value otherwise.

### GetThreeDSRequestorNameOk

`func (o *ThreeDS2RequestData) GetThreeDSRequestorNameOk() (*string, bool)`

GetThreeDSRequestorNameOk returns a tuple with the ThreeDSRequestorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSRequestorName

`func (o *ThreeDS2RequestData) SetThreeDSRequestorName(v string)`

SetThreeDSRequestorName sets ThreeDSRequestorName field to given value.

### HasThreeDSRequestorName

`func (o *ThreeDS2RequestData) HasThreeDSRequestorName() bool`

HasThreeDSRequestorName returns a boolean if a field has been set.

### GetThreeDSRequestorPriorAuthenticationInfo

`func (o *ThreeDS2RequestData) GetThreeDSRequestorPriorAuthenticationInfo() ThreeDSRequestorPriorAuthenticationInfo`

GetThreeDSRequestorPriorAuthenticationInfo returns the ThreeDSRequestorPriorAuthenticationInfo field if non-nil, zero value otherwise.

### GetThreeDSRequestorPriorAuthenticationInfoOk

`func (o *ThreeDS2RequestData) GetThreeDSRequestorPriorAuthenticationInfoOk() (*ThreeDSRequestorPriorAuthenticationInfo, bool)`

GetThreeDSRequestorPriorAuthenticationInfoOk returns a tuple with the ThreeDSRequestorPriorAuthenticationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSRequestorPriorAuthenticationInfo

`func (o *ThreeDS2RequestData) SetThreeDSRequestorPriorAuthenticationInfo(v ThreeDSRequestorPriorAuthenticationInfo)`

SetThreeDSRequestorPriorAuthenticationInfo sets ThreeDSRequestorPriorAuthenticationInfo field to given value.

### HasThreeDSRequestorPriorAuthenticationInfo

`func (o *ThreeDS2RequestData) HasThreeDSRequestorPriorAuthenticationInfo() bool`

HasThreeDSRequestorPriorAuthenticationInfo returns a boolean if a field has been set.

### GetThreeDSRequestorURL

`func (o *ThreeDS2RequestData) GetThreeDSRequestorURL() string`

GetThreeDSRequestorURL returns the ThreeDSRequestorURL field if non-nil, zero value otherwise.

### GetThreeDSRequestorURLOk

`func (o *ThreeDS2RequestData) GetThreeDSRequestorURLOk() (*string, bool)`

GetThreeDSRequestorURLOk returns a tuple with the ThreeDSRequestorURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSRequestorURL

`func (o *ThreeDS2RequestData) SetThreeDSRequestorURL(v string)`

SetThreeDSRequestorURL sets ThreeDSRequestorURL field to given value.

### HasThreeDSRequestorURL

`func (o *ThreeDS2RequestData) HasThreeDSRequestorURL() bool`

HasThreeDSRequestorURL returns a boolean if a field has been set.

### GetTransType

`func (o *ThreeDS2RequestData) GetTransType() string`

GetTransType returns the TransType field if non-nil, zero value otherwise.

### GetTransTypeOk

`func (o *ThreeDS2RequestData) GetTransTypeOk() (*string, bool)`

GetTransTypeOk returns a tuple with the TransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransType

`func (o *ThreeDS2RequestData) SetTransType(v string)`

SetTransType sets TransType field to given value.

### HasTransType

`func (o *ThreeDS2RequestData) HasTransType() bool`

HasTransType returns a boolean if a field has been set.

### GetTransactionType

`func (o *ThreeDS2RequestData) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *ThreeDS2RequestData) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *ThreeDS2RequestData) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *ThreeDS2RequestData) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetWhiteListStatus

`func (o *ThreeDS2RequestData) GetWhiteListStatus() string`

GetWhiteListStatus returns the WhiteListStatus field if non-nil, zero value otherwise.

### GetWhiteListStatusOk

`func (o *ThreeDS2RequestData) GetWhiteListStatusOk() (*string, bool)`

GetWhiteListStatusOk returns a tuple with the WhiteListStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteListStatus

`func (o *ThreeDS2RequestData) SetWhiteListStatus(v string)`

SetWhiteListStatus sets WhiteListStatus field to given value.

### HasWhiteListStatus

`func (o *ThreeDS2RequestData) HasWhiteListStatus() bool`

HasWhiteListStatus returns a boolean if a field has been set.

### GetWorkPhone

`func (o *ThreeDS2RequestData) GetWorkPhone() Phone`

GetWorkPhone returns the WorkPhone field if non-nil, zero value otherwise.

### GetWorkPhoneOk

`func (o *ThreeDS2RequestData) GetWorkPhoneOk() (*Phone, bool)`

GetWorkPhoneOk returns a tuple with the WorkPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkPhone

`func (o *ThreeDS2RequestData) SetWorkPhone(v Phone)`

SetWorkPhone sets WorkPhone field to given value.

### HasWorkPhone

`func (o *ThreeDS2RequestData) HasWorkPhone() bool`

HasWorkPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


