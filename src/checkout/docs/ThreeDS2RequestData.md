# ThreeDS2RequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcquirerBIN** | **string** | Required for [authentication-only integration](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only). The acquiring BIN enrolled for 3D Secure 2. This string should match the value that you will use in the authorisation. Use 123456 on the Test platform. | [optional] 
**AcquirerMerchantID** | **string** | Required for [authentication-only integration](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only). The merchantId that is enrolled for 3D Secure 2 by the merchant&#39;s acquirer. This string should match the value that you will use in the authorisation. Use 123456 on the Test platform. | [optional] 
**AuthenticationOnly** | **bool** | If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation. | [optional] 
**ChallengeIndicator** | **string** | Possibility to specify a preference for receiving a challenge from the issuer. Allowed values: * &#x60;noPreference&#x60; * &#x60;requestNoChallenge&#x60; * &#x60;requestChallenge&#x60; * &#x60;requestChallengeAsMandate&#x60;  | [optional] 
**DeviceChannel** | **string** | The environment of the shopper. Allowed values: * &#x60;app&#x60; * &#x60;browser&#x60; | 
**DeviceRenderOptions** | [**DeviceRenderOptions**](DeviceRenderOptions.md) |  | [optional] 
**Mcc** | **string** | Required for merchants that have been enrolled for 3D Secure 2 by another party than Adyen, mostly [authentication-only integrations](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only). The &#x60;mcc&#x60; is a four-digit code with which the previously given &#x60;acquirerMerchantID&#x60; is registered at the scheme. | [optional] 
**MerchantName** | **string** | Required for [authentication-only integration](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only). The merchant name that the issuer presents to the shopper if they get a challenge. We recommend to use the same value that you will use in the authorization. Maximum length is 40 characters. &gt; Optional for a [full 3D Secure 2 integration](https://docs.adyen.com/checkout/3d-secure/native-3ds2/api-integration). Use this field if you are enrolled for 3D Secure 2 with us and want to override the merchant name already configured on your account. | [optional] 
**MessageVersion** | **string** | The &#x60;messageVersion&#x60; value indicating the 3D Secure 2 protocol version. | [optional] 
**NotificationURL** | **string** | URL to where the issuer should send the &#x60;CRes&#x60;. Required if you are not using components for &#x60;channel&#x60; **Web** or if you are using classic integration &#x60;deviceChannel&#x60; **browser**. | [optional] 
**SdkAppID** | **string** | The &#x60;sdkAppID&#x60; value as received from the 3D Secure 2 SDK. Required for &#x60;deviceChannel&#x60; set to **app**. | [optional] 
**SdkEncData** | **string** | The &#x60;sdkEncData&#x60; value as received from the 3D Secure 2 SDK. Required for &#x60;deviceChannel&#x60; set to **app**. | [optional] 
**SdkEphemPubKey** | [**SDKEphemPubKey**](SDKEphemPubKey.md) |  | [optional] 
**SdkMaxTimeout** | **int32** | The maximum amount of time in minutes for the 3D Secure 2 authentication process. Optional and only for &#x60;deviceChannel&#x60; set to **app**. Defaults to **60** minutes. | [optional] 
**SdkReferenceNumber** | **string** | The &#x60;sdkReferenceNumber&#x60; value as received from the 3D Secure 2 SDK. Only for &#x60;deviceChannel&#x60; set to **app**. | [optional] 
**SdkTransID** | **string** | The &#x60;sdkTransID&#x60; value as received from the 3D Secure 2 SDK. Only for &#x60;deviceChannel&#x60; set to **app**. | [optional] 
**ThreeDSCompInd** | **string** | Completion indicator for the device fingerprinting. | [optional] 
**ThreeDSRequestorID** | **string** | Required for [authentication-only integration](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only) for Visa. Unique 3D Secure requestor identifier assigned by the Directory Server when you enrol for 3D Secure 2. | [optional] 
**ThreeDSRequestorName** | **string** | Required for [authentication-only integration](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only) for Visa. Unique 3D Secure requestor name assigned by the Directory Server when you enrol for 3D Secure 2. | [optional] 
**ThreeDSRequestorURL** | **string** | URL of the (customer service) website that will be shown to the shopper in case of technical errors during the 3D Secure 2 process. | [optional] 
**TransactionType** | **string** | Identify the type of the transaction being authenticated. | [optional] 
**WhiteListStatus** | **string** | The &#x60;whiteListStatus&#x60; value returned from a previous 3D Secure 2 transaction, only applicable for 3D Secure 2 protocol version 2.2.0. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


