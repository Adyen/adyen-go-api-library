# ThreeDSecureData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationResponse** | Pointer to **string** | In 3D Secure 1, the authentication response if the shopper was redirected.  In 3D Secure 2, this is the &#x60;transStatus&#x60; from the challenge result. If the transaction was frictionless, omit this parameter. | [optional] 
**Cavv** | Pointer to **string** | The cardholder authentication value (base64 encoded, 20 bytes in a decoded form). | [optional] 
**CavvAlgorithm** | Pointer to **string** | The CAVV algorithm used. Include this only for 3D Secure 1. | [optional] 
**ChallengeCancel** | Pointer to **string** | Indicator informing the Access Control Server (ACS) and the Directory Server (DS) that the authentication has been cancelled. For possible values, refer to [3D Secure API reference](https://docs.adyen.com/online-payments/3d-secure/api-reference#mpidata). | [optional] 
**DirectoryResponse** | Pointer to **string** | In 3D Secure 1, this is the enrollment response from the 3D directory server.  In 3D Secure 2, this is the &#x60;transStatus&#x60; from the &#x60;ARes&#x60;. | [optional] 
**DsTransID** | Pointer to **string** | Supported for 3D Secure 2. The unique transaction identifier assigned by the Directory Server (DS) to identify a single transaction. | [optional] 
**Eci** | Pointer to **string** | The electronic commerce indicator. | [optional] 
**RiskScore** | Pointer to **string** | Risk score calculated by Directory Server (DS). Required for Cartes Bancaires integrations. | [optional] 
**ThreeDSVersion** | Pointer to **string** | The version of the 3D Secure protocol. | [optional] 
**TokenAuthenticationVerificationValue** | Pointer to **string** | Network token authentication verification value (TAVV). The network token cryptogram. | [optional] 
**TransStatusReason** | Pointer to **string** | Provides information on why the &#x60;transStatus&#x60; field has the specified value. For possible values, refer to [our docs](https://docs.adyen.com/online-payments/3d-secure/api-reference#possible-transstatusreason-values). | [optional] 
**Xid** | Pointer to **string** | Supported for 3D Secure 1. The transaction identifier (Base64-encoded, 20 bytes in a decoded form). | [optional] 

## Methods

### NewThreeDSecureData

`func NewThreeDSecureData() *ThreeDSecureData`

NewThreeDSecureData instantiates a new ThreeDSecureData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSecureDataWithDefaults

`func NewThreeDSecureDataWithDefaults() *ThreeDSecureData`

NewThreeDSecureDataWithDefaults instantiates a new ThreeDSecureData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationResponse

`func (o *ThreeDSecureData) GetAuthenticationResponse() string`

GetAuthenticationResponse returns the AuthenticationResponse field if non-nil, zero value otherwise.

### GetAuthenticationResponseOk

`func (o *ThreeDSecureData) GetAuthenticationResponseOk() (*string, bool)`

GetAuthenticationResponseOk returns a tuple with the AuthenticationResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationResponse

`func (o *ThreeDSecureData) SetAuthenticationResponse(v string)`

SetAuthenticationResponse sets AuthenticationResponse field to given value.

### HasAuthenticationResponse

`func (o *ThreeDSecureData) HasAuthenticationResponse() bool`

HasAuthenticationResponse returns a boolean if a field has been set.

### GetCavv

`func (o *ThreeDSecureData) GetCavv() string`

GetCavv returns the Cavv field if non-nil, zero value otherwise.

### GetCavvOk

`func (o *ThreeDSecureData) GetCavvOk() (*string, bool)`

GetCavvOk returns a tuple with the Cavv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavv

`func (o *ThreeDSecureData) SetCavv(v string)`

SetCavv sets Cavv field to given value.

### HasCavv

`func (o *ThreeDSecureData) HasCavv() bool`

HasCavv returns a boolean if a field has been set.

### GetCavvAlgorithm

`func (o *ThreeDSecureData) GetCavvAlgorithm() string`

GetCavvAlgorithm returns the CavvAlgorithm field if non-nil, zero value otherwise.

### GetCavvAlgorithmOk

`func (o *ThreeDSecureData) GetCavvAlgorithmOk() (*string, bool)`

GetCavvAlgorithmOk returns a tuple with the CavvAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavvAlgorithm

`func (o *ThreeDSecureData) SetCavvAlgorithm(v string)`

SetCavvAlgorithm sets CavvAlgorithm field to given value.

### HasCavvAlgorithm

`func (o *ThreeDSecureData) HasCavvAlgorithm() bool`

HasCavvAlgorithm returns a boolean if a field has been set.

### GetChallengeCancel

`func (o *ThreeDSecureData) GetChallengeCancel() string`

GetChallengeCancel returns the ChallengeCancel field if non-nil, zero value otherwise.

### GetChallengeCancelOk

`func (o *ThreeDSecureData) GetChallengeCancelOk() (*string, bool)`

GetChallengeCancelOk returns a tuple with the ChallengeCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeCancel

`func (o *ThreeDSecureData) SetChallengeCancel(v string)`

SetChallengeCancel sets ChallengeCancel field to given value.

### HasChallengeCancel

`func (o *ThreeDSecureData) HasChallengeCancel() bool`

HasChallengeCancel returns a boolean if a field has been set.

### GetDirectoryResponse

`func (o *ThreeDSecureData) GetDirectoryResponse() string`

GetDirectoryResponse returns the DirectoryResponse field if non-nil, zero value otherwise.

### GetDirectoryResponseOk

`func (o *ThreeDSecureData) GetDirectoryResponseOk() (*string, bool)`

GetDirectoryResponseOk returns a tuple with the DirectoryResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryResponse

`func (o *ThreeDSecureData) SetDirectoryResponse(v string)`

SetDirectoryResponse sets DirectoryResponse field to given value.

### HasDirectoryResponse

`func (o *ThreeDSecureData) HasDirectoryResponse() bool`

HasDirectoryResponse returns a boolean if a field has been set.

### GetDsTransID

`func (o *ThreeDSecureData) GetDsTransID() string`

GetDsTransID returns the DsTransID field if non-nil, zero value otherwise.

### GetDsTransIDOk

`func (o *ThreeDSecureData) GetDsTransIDOk() (*string, bool)`

GetDsTransIDOk returns a tuple with the DsTransID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsTransID

`func (o *ThreeDSecureData) SetDsTransID(v string)`

SetDsTransID sets DsTransID field to given value.

### HasDsTransID

`func (o *ThreeDSecureData) HasDsTransID() bool`

HasDsTransID returns a boolean if a field has been set.

### GetEci

`func (o *ThreeDSecureData) GetEci() string`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *ThreeDSecureData) GetEciOk() (*string, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *ThreeDSecureData) SetEci(v string)`

SetEci sets Eci field to given value.

### HasEci

`func (o *ThreeDSecureData) HasEci() bool`

HasEci returns a boolean if a field has been set.

### GetRiskScore

`func (o *ThreeDSecureData) GetRiskScore() string`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *ThreeDSecureData) GetRiskScoreOk() (*string, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *ThreeDSecureData) SetRiskScore(v string)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *ThreeDSecureData) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetThreeDSVersion

`func (o *ThreeDSecureData) GetThreeDSVersion() string`

GetThreeDSVersion returns the ThreeDSVersion field if non-nil, zero value otherwise.

### GetThreeDSVersionOk

`func (o *ThreeDSecureData) GetThreeDSVersionOk() (*string, bool)`

GetThreeDSVersionOk returns a tuple with the ThreeDSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSVersion

`func (o *ThreeDSecureData) SetThreeDSVersion(v string)`

SetThreeDSVersion sets ThreeDSVersion field to given value.

### HasThreeDSVersion

`func (o *ThreeDSecureData) HasThreeDSVersion() bool`

HasThreeDSVersion returns a boolean if a field has been set.

### GetTokenAuthenticationVerificationValue

`func (o *ThreeDSecureData) GetTokenAuthenticationVerificationValue() string`

GetTokenAuthenticationVerificationValue returns the TokenAuthenticationVerificationValue field if non-nil, zero value otherwise.

### GetTokenAuthenticationVerificationValueOk

`func (o *ThreeDSecureData) GetTokenAuthenticationVerificationValueOk() (*string, bool)`

GetTokenAuthenticationVerificationValueOk returns a tuple with the TokenAuthenticationVerificationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAuthenticationVerificationValue

`func (o *ThreeDSecureData) SetTokenAuthenticationVerificationValue(v string)`

SetTokenAuthenticationVerificationValue sets TokenAuthenticationVerificationValue field to given value.

### HasTokenAuthenticationVerificationValue

`func (o *ThreeDSecureData) HasTokenAuthenticationVerificationValue() bool`

HasTokenAuthenticationVerificationValue returns a boolean if a field has been set.

### GetTransStatusReason

`func (o *ThreeDSecureData) GetTransStatusReason() string`

GetTransStatusReason returns the TransStatusReason field if non-nil, zero value otherwise.

### GetTransStatusReasonOk

`func (o *ThreeDSecureData) GetTransStatusReasonOk() (*string, bool)`

GetTransStatusReasonOk returns a tuple with the TransStatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransStatusReason

`func (o *ThreeDSecureData) SetTransStatusReason(v string)`

SetTransStatusReason sets TransStatusReason field to given value.

### HasTransStatusReason

`func (o *ThreeDSecureData) HasTransStatusReason() bool`

HasTransStatusReason returns a boolean if a field has been set.

### GetXid

`func (o *ThreeDSecureData) GetXid() string`

GetXid returns the Xid field if non-nil, zero value otherwise.

### GetXidOk

`func (o *ThreeDSecureData) GetXidOk() (*string, bool)`

GetXidOk returns a tuple with the Xid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXid

`func (o *ThreeDSecureData) SetXid(v string)`

SetXid sets Xid field to given value.

### HasXid

`func (o *ThreeDSecureData) HasXid() bool`

HasXid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


