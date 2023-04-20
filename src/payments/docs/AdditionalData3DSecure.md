# AdditionalData3DSecure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allow3DS2** | Pointer to **string** | Indicates if you are able to process 3D Secure 2 transactions natively on your payment page. Send this parameter when you are using &#x60;/payments&#x60; endpoint with any of our [native 3D Secure 2 solutions](https://docs.adyen.com/online-payments/3d-secure/native-3ds2).   &gt; This parameter only indicates readiness to support native 3D Secure 2 authentication. To specify if you _want_ to perform 3D Secure, use [Dynamic 3D Secure](/risk-management/dynamic-3d-secure) or send the &#x60;executeThreeD&#x60; parameter.  Possible values: * **true** - Ready to support native 3D Secure 2 authentication. Setting this to true does not mean always applying 3D Secure 2. Adyen still selects the version of 3D Secure based on configuration to optimize authorisation rates and improve the shopper&#39;s experience. * **false** – Not ready to support native 3D Secure 2 authentication. Adyen will not offer 3D Secure 2 to your shopper regardless of your configuration.  | [optional] 
**ChallengeWindowSize** | Pointer to **string** | Dimensions of the 3DS2 challenge window to be displayed to the cardholder.  Possible values:  * **01** - size of 250x400  * **02** - size of 390x400 * **03** - size of 500x600 * **04** - size of 600x400 * **05** - Fullscreen | [optional] 
**ExecuteThreeD** | Pointer to **string** | Indicates if you want to perform 3D Secure authentication on a transaction.   &gt; Alternatively, you can use [Dynamic 3D Secure](/risk-management/dynamic-3d-secure) to configure rules for applying 3D Secure.  Possible values: * **true** – Perform 3D Secure authentication. * **false** – Don&#39;t perform 3D Secure authentication. Note that this setting results in refusals if the issuer mandates 3D Secure because of the PSD2 directive  or other, national regulations.   | [optional] 
**MpiImplementationType** | Pointer to **string** | In case of Secure+, this field must be set to **CUPSecurePlus**. | [optional] 
**ScaExemption** | Pointer to **string** | Indicates the [exemption type](https://docs.adyen.com/payments-fundamentals/psd2-sca-compliance-and-implementation-guide#specifypreferenceinyourapirequest) that you want to request for the transaction.   Possible values: * **lowValue**  * **secureCorporate**  * **trustedBeneficiary**  * **transactionRiskAnalysis**  | [optional] 
**ThreeDSVersion** | Pointer to **string** | Indicates your preference for the 3D Secure version.  &gt; If you use this parameter, you override the checks from Adyen&#39;s Authentication Engine. We recommend to use this field only if you have an extensive knowledge of 3D Secure.  Possible values: * **1.0.2**: Apply 3D Secure version 1.0.2.  * **2.1.0**: Apply 3D Secure version 2.1.0.  * **2.2.0**: Apply 3D Secure version 2.2.0. If the issuer does not support version 2.2.0, we will fall back to 2.1.0.  The following rules apply: * If you prefer 2.1.0 or 2.2.0 but we receive a negative &#x60;transStatus&#x60; in the &#x60;ARes&#x60;, we will apply the fallback policy configured in your account. For example, if the configuration is to fall back to 3D Secure 1, we will apply version 1.0.2. * If you prefer 2.1.0 or 2.2.0 but the BIN is not enrolled, you will receive an error.   | [optional] 

## Methods

### NewAdditionalData3DSecure

`func NewAdditionalData3DSecure() *AdditionalData3DSecure`

NewAdditionalData3DSecure instantiates a new AdditionalData3DSecure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalData3DSecureWithDefaults

`func NewAdditionalData3DSecureWithDefaults() *AdditionalData3DSecure`

NewAdditionalData3DSecureWithDefaults instantiates a new AdditionalData3DSecure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllow3DS2

`func (o *AdditionalData3DSecure) GetAllow3DS2() string`

GetAllow3DS2 returns the Allow3DS2 field if non-nil, zero value otherwise.

### GetAllow3DS2Ok

`func (o *AdditionalData3DSecure) GetAllow3DS2Ok() (*string, bool)`

GetAllow3DS2Ok returns a tuple with the Allow3DS2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow3DS2

`func (o *AdditionalData3DSecure) SetAllow3DS2(v string)`

SetAllow3DS2 sets Allow3DS2 field to given value.

### HasAllow3DS2

`func (o *AdditionalData3DSecure) HasAllow3DS2() bool`

HasAllow3DS2 returns a boolean if a field has been set.

### GetChallengeWindowSize

`func (o *AdditionalData3DSecure) GetChallengeWindowSize() string`

GetChallengeWindowSize returns the ChallengeWindowSize field if non-nil, zero value otherwise.

### GetChallengeWindowSizeOk

`func (o *AdditionalData3DSecure) GetChallengeWindowSizeOk() (*string, bool)`

GetChallengeWindowSizeOk returns a tuple with the ChallengeWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeWindowSize

`func (o *AdditionalData3DSecure) SetChallengeWindowSize(v string)`

SetChallengeWindowSize sets ChallengeWindowSize field to given value.

### HasChallengeWindowSize

`func (o *AdditionalData3DSecure) HasChallengeWindowSize() bool`

HasChallengeWindowSize returns a boolean if a field has been set.

### GetExecuteThreeD

`func (o *AdditionalData3DSecure) GetExecuteThreeD() string`

GetExecuteThreeD returns the ExecuteThreeD field if non-nil, zero value otherwise.

### GetExecuteThreeDOk

`func (o *AdditionalData3DSecure) GetExecuteThreeDOk() (*string, bool)`

GetExecuteThreeDOk returns a tuple with the ExecuteThreeD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteThreeD

`func (o *AdditionalData3DSecure) SetExecuteThreeD(v string)`

SetExecuteThreeD sets ExecuteThreeD field to given value.

### HasExecuteThreeD

`func (o *AdditionalData3DSecure) HasExecuteThreeD() bool`

HasExecuteThreeD returns a boolean if a field has been set.

### GetMpiImplementationType

`func (o *AdditionalData3DSecure) GetMpiImplementationType() string`

GetMpiImplementationType returns the MpiImplementationType field if non-nil, zero value otherwise.

### GetMpiImplementationTypeOk

`func (o *AdditionalData3DSecure) GetMpiImplementationTypeOk() (*string, bool)`

GetMpiImplementationTypeOk returns a tuple with the MpiImplementationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpiImplementationType

`func (o *AdditionalData3DSecure) SetMpiImplementationType(v string)`

SetMpiImplementationType sets MpiImplementationType field to given value.

### HasMpiImplementationType

`func (o *AdditionalData3DSecure) HasMpiImplementationType() bool`

HasMpiImplementationType returns a boolean if a field has been set.

### GetScaExemption

`func (o *AdditionalData3DSecure) GetScaExemption() string`

GetScaExemption returns the ScaExemption field if non-nil, zero value otherwise.

### GetScaExemptionOk

`func (o *AdditionalData3DSecure) GetScaExemptionOk() (*string, bool)`

GetScaExemptionOk returns a tuple with the ScaExemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaExemption

`func (o *AdditionalData3DSecure) SetScaExemption(v string)`

SetScaExemption sets ScaExemption field to given value.

### HasScaExemption

`func (o *AdditionalData3DSecure) HasScaExemption() bool`

HasScaExemption returns a boolean if a field has been set.

### GetThreeDSVersion

`func (o *AdditionalData3DSecure) GetThreeDSVersion() string`

GetThreeDSVersion returns the ThreeDSVersion field if non-nil, zero value otherwise.

### GetThreeDSVersionOk

`func (o *AdditionalData3DSecure) GetThreeDSVersionOk() (*string, bool)`

GetThreeDSVersionOk returns a tuple with the ThreeDSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSVersion

`func (o *AdditionalData3DSecure) SetThreeDSVersion(v string)`

SetThreeDSVersion sets ThreeDSVersion field to given value.

### HasThreeDSVersion

`func (o *AdditionalData3DSecure) HasThreeDSVersion() bool`

HasThreeDSVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


