# ThreeDS2Result

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationValue** | Pointer to **string** | The &#x60;authenticationValue&#x60; value as defined in the 3D Secure 2 specification. | [optional] 
**CavvAlgorithm** | Pointer to **string** | The algorithm used by the ACS to calculate the authentication value, only for Cartes Bancaires integrations. | [optional] 
**ChallengeCancel** | Pointer to **string** | Indicator informing the Access Control Server (ACS) and the Directory Server (DS) that the authentication has been cancelled. For possible values, refer to [3D Secure API reference](https://docs.adyen.com/online-payments/3d-secure/api-reference#mpidata). | [optional] 
**ChallengeIndicator** | Pointer to **string** | Specifies a preference for receiving a challenge from the issuer. Allowed values: * &#x60;noPreference&#x60; * &#x60;requestNoChallenge&#x60; * &#x60;requestChallenge&#x60; * &#x60;requestChallengeAsMandate&#x60;  | [optional] 
**DsTransID** | Pointer to **string** | The &#x60;dsTransID&#x60; value as defined in the 3D Secure 2 specification. | [optional] 
**Eci** | Pointer to **string** | The &#x60;eci&#x60; value as defined in the 3D Secure 2 specification. | [optional] 
**ExemptionIndicator** | Pointer to **string** | Indicates the exemption type that was applied by the issuer to the authentication, if exemption applied. Allowed values: * &#x60;lowValue&#x60; * &#x60;secureCorporate&#x60; * &#x60;trustedBeneficiary&#x60; * &#x60;transactionRiskAnalysis&#x60;  | [optional] 
**MessageVersion** | Pointer to **string** | The &#x60;messageVersion&#x60; value as defined in the 3D Secure 2 specification. | [optional] 
**RiskScore** | Pointer to **string** | Risk score calculated by Cartes Bancaires Directory Server (DS). | [optional] 
**ThreeDSServerTransID** | Pointer to **string** | The &#x60;threeDSServerTransID&#x60; value as defined in the 3D Secure 2 specification. | [optional] 
**Timestamp** | Pointer to **string** | The &#x60;timestamp&#x60; value of the 3D Secure 2 authentication. | [optional] 
**TransStatus** | Pointer to **string** | The &#x60;transStatus&#x60; value as defined in the 3D Secure 2 specification. | [optional] 
**TransStatusReason** | Pointer to **string** | Provides information on why the &#x60;transStatus&#x60; field has the specified value. For possible values, refer to [our docs](https://docs.adyen.com/online-payments/3d-secure/api-reference#possible-transstatusreason-values). | [optional] 
**WhiteListStatus** | Pointer to **string** | The &#x60;whiteListStatus&#x60; value as defined in the 3D Secure 2 specification. | [optional] 

## Methods

### NewThreeDS2Result

`func NewThreeDS2Result() *ThreeDS2Result`

NewThreeDS2Result instantiates a new ThreeDS2Result object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDS2ResultWithDefaults

`func NewThreeDS2ResultWithDefaults() *ThreeDS2Result`

NewThreeDS2ResultWithDefaults instantiates a new ThreeDS2Result object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationValue

`func (o *ThreeDS2Result) GetAuthenticationValue() string`

GetAuthenticationValue returns the AuthenticationValue field if non-nil, zero value otherwise.

### GetAuthenticationValueOk

`func (o *ThreeDS2Result) GetAuthenticationValueOk() (*string, bool)`

GetAuthenticationValueOk returns a tuple with the AuthenticationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationValue

`func (o *ThreeDS2Result) SetAuthenticationValue(v string)`

SetAuthenticationValue sets AuthenticationValue field to given value.

### HasAuthenticationValue

`func (o *ThreeDS2Result) HasAuthenticationValue() bool`

HasAuthenticationValue returns a boolean if a field has been set.

### GetCavvAlgorithm

`func (o *ThreeDS2Result) GetCavvAlgorithm() string`

GetCavvAlgorithm returns the CavvAlgorithm field if non-nil, zero value otherwise.

### GetCavvAlgorithmOk

`func (o *ThreeDS2Result) GetCavvAlgorithmOk() (*string, bool)`

GetCavvAlgorithmOk returns a tuple with the CavvAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavvAlgorithm

`func (o *ThreeDS2Result) SetCavvAlgorithm(v string)`

SetCavvAlgorithm sets CavvAlgorithm field to given value.

### HasCavvAlgorithm

`func (o *ThreeDS2Result) HasCavvAlgorithm() bool`

HasCavvAlgorithm returns a boolean if a field has been set.

### GetChallengeCancel

`func (o *ThreeDS2Result) GetChallengeCancel() string`

GetChallengeCancel returns the ChallengeCancel field if non-nil, zero value otherwise.

### GetChallengeCancelOk

`func (o *ThreeDS2Result) GetChallengeCancelOk() (*string, bool)`

GetChallengeCancelOk returns a tuple with the ChallengeCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeCancel

`func (o *ThreeDS2Result) SetChallengeCancel(v string)`

SetChallengeCancel sets ChallengeCancel field to given value.

### HasChallengeCancel

`func (o *ThreeDS2Result) HasChallengeCancel() bool`

HasChallengeCancel returns a boolean if a field has been set.

### GetChallengeIndicator

`func (o *ThreeDS2Result) GetChallengeIndicator() string`

GetChallengeIndicator returns the ChallengeIndicator field if non-nil, zero value otherwise.

### GetChallengeIndicatorOk

`func (o *ThreeDS2Result) GetChallengeIndicatorOk() (*string, bool)`

GetChallengeIndicatorOk returns a tuple with the ChallengeIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeIndicator

`func (o *ThreeDS2Result) SetChallengeIndicator(v string)`

SetChallengeIndicator sets ChallengeIndicator field to given value.

### HasChallengeIndicator

`func (o *ThreeDS2Result) HasChallengeIndicator() bool`

HasChallengeIndicator returns a boolean if a field has been set.

### GetDsTransID

`func (o *ThreeDS2Result) GetDsTransID() string`

GetDsTransID returns the DsTransID field if non-nil, zero value otherwise.

### GetDsTransIDOk

`func (o *ThreeDS2Result) GetDsTransIDOk() (*string, bool)`

GetDsTransIDOk returns a tuple with the DsTransID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsTransID

`func (o *ThreeDS2Result) SetDsTransID(v string)`

SetDsTransID sets DsTransID field to given value.

### HasDsTransID

`func (o *ThreeDS2Result) HasDsTransID() bool`

HasDsTransID returns a boolean if a field has been set.

### GetEci

`func (o *ThreeDS2Result) GetEci() string`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *ThreeDS2Result) GetEciOk() (*string, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *ThreeDS2Result) SetEci(v string)`

SetEci sets Eci field to given value.

### HasEci

`func (o *ThreeDS2Result) HasEci() bool`

HasEci returns a boolean if a field has been set.

### GetExemptionIndicator

`func (o *ThreeDS2Result) GetExemptionIndicator() string`

GetExemptionIndicator returns the ExemptionIndicator field if non-nil, zero value otherwise.

### GetExemptionIndicatorOk

`func (o *ThreeDS2Result) GetExemptionIndicatorOk() (*string, bool)`

GetExemptionIndicatorOk returns a tuple with the ExemptionIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptionIndicator

`func (o *ThreeDS2Result) SetExemptionIndicator(v string)`

SetExemptionIndicator sets ExemptionIndicator field to given value.

### HasExemptionIndicator

`func (o *ThreeDS2Result) HasExemptionIndicator() bool`

HasExemptionIndicator returns a boolean if a field has been set.

### GetMessageVersion

`func (o *ThreeDS2Result) GetMessageVersion() string`

GetMessageVersion returns the MessageVersion field if non-nil, zero value otherwise.

### GetMessageVersionOk

`func (o *ThreeDS2Result) GetMessageVersionOk() (*string, bool)`

GetMessageVersionOk returns a tuple with the MessageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageVersion

`func (o *ThreeDS2Result) SetMessageVersion(v string)`

SetMessageVersion sets MessageVersion field to given value.

### HasMessageVersion

`func (o *ThreeDS2Result) HasMessageVersion() bool`

HasMessageVersion returns a boolean if a field has been set.

### GetRiskScore

`func (o *ThreeDS2Result) GetRiskScore() string`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *ThreeDS2Result) GetRiskScoreOk() (*string, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *ThreeDS2Result) SetRiskScore(v string)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *ThreeDS2Result) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetThreeDSServerTransID

`func (o *ThreeDS2Result) GetThreeDSServerTransID() string`

GetThreeDSServerTransID returns the ThreeDSServerTransID field if non-nil, zero value otherwise.

### GetThreeDSServerTransIDOk

`func (o *ThreeDS2Result) GetThreeDSServerTransIDOk() (*string, bool)`

GetThreeDSServerTransIDOk returns a tuple with the ThreeDSServerTransID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSServerTransID

`func (o *ThreeDS2Result) SetThreeDSServerTransID(v string)`

SetThreeDSServerTransID sets ThreeDSServerTransID field to given value.

### HasThreeDSServerTransID

`func (o *ThreeDS2Result) HasThreeDSServerTransID() bool`

HasThreeDSServerTransID returns a boolean if a field has been set.

### GetTimestamp

`func (o *ThreeDS2Result) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ThreeDS2Result) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ThreeDS2Result) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ThreeDS2Result) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTransStatus

`func (o *ThreeDS2Result) GetTransStatus() string`

GetTransStatus returns the TransStatus field if non-nil, zero value otherwise.

### GetTransStatusOk

`func (o *ThreeDS2Result) GetTransStatusOk() (*string, bool)`

GetTransStatusOk returns a tuple with the TransStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransStatus

`func (o *ThreeDS2Result) SetTransStatus(v string)`

SetTransStatus sets TransStatus field to given value.

### HasTransStatus

`func (o *ThreeDS2Result) HasTransStatus() bool`

HasTransStatus returns a boolean if a field has been set.

### GetTransStatusReason

`func (o *ThreeDS2Result) GetTransStatusReason() string`

GetTransStatusReason returns the TransStatusReason field if non-nil, zero value otherwise.

### GetTransStatusReasonOk

`func (o *ThreeDS2Result) GetTransStatusReasonOk() (*string, bool)`

GetTransStatusReasonOk returns a tuple with the TransStatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransStatusReason

`func (o *ThreeDS2Result) SetTransStatusReason(v string)`

SetTransStatusReason sets TransStatusReason field to given value.

### HasTransStatusReason

`func (o *ThreeDS2Result) HasTransStatusReason() bool`

HasTransStatusReason returns a boolean if a field has been set.

### GetWhiteListStatus

`func (o *ThreeDS2Result) GetWhiteListStatus() string`

GetWhiteListStatus returns the WhiteListStatus field if non-nil, zero value otherwise.

### GetWhiteListStatusOk

`func (o *ThreeDS2Result) GetWhiteListStatusOk() (*string, bool)`

GetWhiteListStatusOk returns a tuple with the WhiteListStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteListStatus

`func (o *ThreeDS2Result) SetWhiteListStatus(v string)`

SetWhiteListStatus sets WhiteListStatus field to given value.

### HasWhiteListStatus

`func (o *ThreeDS2Result) HasWhiteListStatus() bool`

HasWhiteListStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


