# ResponseAdditionalData3DSecure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardHolderInfo** | Pointer to **string** | Information provided by the issuer to the cardholder. If this field is present, you need to display this information to the cardholder.  | [optional] 
**Cavv** | Pointer to **string** | The Cardholder Authentication Verification Value (CAVV) for the 3D Secure authentication session, as a Base64-encoded 20-byte array. | [optional] 
**CavvAlgorithm** | Pointer to **string** | The CAVV algorithm used. | [optional] 
**ScaExemptionRequested** | Pointer to **string** | Shows the [exemption type](https://docs.adyen.com/payments-fundamentals/psd2-sca-compliance-and-implementation-guide#specifypreferenceinyourapirequest) that Adyen requested for the payment.   Possible values: * **lowValue**  * **secureCorporate**  * **trustedBeneficiary**  * **transactionRiskAnalysis**  | [optional] 
**Threeds2CardEnrolled** | Pointer to **bool** | Indicates whether a card is enrolled for 3D Secure 2. | [optional] 

## Methods

### NewResponseAdditionalData3DSecure

`func NewResponseAdditionalData3DSecure() *ResponseAdditionalData3DSecure`

NewResponseAdditionalData3DSecure instantiates a new ResponseAdditionalData3DSecure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAdditionalData3DSecureWithDefaults

`func NewResponseAdditionalData3DSecureWithDefaults() *ResponseAdditionalData3DSecure`

NewResponseAdditionalData3DSecureWithDefaults instantiates a new ResponseAdditionalData3DSecure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardHolderInfo

`func (o *ResponseAdditionalData3DSecure) GetCardHolderInfo() string`

GetCardHolderInfo returns the CardHolderInfo field if non-nil, zero value otherwise.

### GetCardHolderInfoOk

`func (o *ResponseAdditionalData3DSecure) GetCardHolderInfoOk() (*string, bool)`

GetCardHolderInfoOk returns a tuple with the CardHolderInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHolderInfo

`func (o *ResponseAdditionalData3DSecure) SetCardHolderInfo(v string)`

SetCardHolderInfo sets CardHolderInfo field to given value.

### HasCardHolderInfo

`func (o *ResponseAdditionalData3DSecure) HasCardHolderInfo() bool`

HasCardHolderInfo returns a boolean if a field has been set.

### GetCavv

`func (o *ResponseAdditionalData3DSecure) GetCavv() string`

GetCavv returns the Cavv field if non-nil, zero value otherwise.

### GetCavvOk

`func (o *ResponseAdditionalData3DSecure) GetCavvOk() (*string, bool)`

GetCavvOk returns a tuple with the Cavv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavv

`func (o *ResponseAdditionalData3DSecure) SetCavv(v string)`

SetCavv sets Cavv field to given value.

### HasCavv

`func (o *ResponseAdditionalData3DSecure) HasCavv() bool`

HasCavv returns a boolean if a field has been set.

### GetCavvAlgorithm

`func (o *ResponseAdditionalData3DSecure) GetCavvAlgorithm() string`

GetCavvAlgorithm returns the CavvAlgorithm field if non-nil, zero value otherwise.

### GetCavvAlgorithmOk

`func (o *ResponseAdditionalData3DSecure) GetCavvAlgorithmOk() (*string, bool)`

GetCavvAlgorithmOk returns a tuple with the CavvAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavvAlgorithm

`func (o *ResponseAdditionalData3DSecure) SetCavvAlgorithm(v string)`

SetCavvAlgorithm sets CavvAlgorithm field to given value.

### HasCavvAlgorithm

`func (o *ResponseAdditionalData3DSecure) HasCavvAlgorithm() bool`

HasCavvAlgorithm returns a boolean if a field has been set.

### GetScaExemptionRequested

`func (o *ResponseAdditionalData3DSecure) GetScaExemptionRequested() string`

GetScaExemptionRequested returns the ScaExemptionRequested field if non-nil, zero value otherwise.

### GetScaExemptionRequestedOk

`func (o *ResponseAdditionalData3DSecure) GetScaExemptionRequestedOk() (*string, bool)`

GetScaExemptionRequestedOk returns a tuple with the ScaExemptionRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaExemptionRequested

`func (o *ResponseAdditionalData3DSecure) SetScaExemptionRequested(v string)`

SetScaExemptionRequested sets ScaExemptionRequested field to given value.

### HasScaExemptionRequested

`func (o *ResponseAdditionalData3DSecure) HasScaExemptionRequested() bool`

HasScaExemptionRequested returns a boolean if a field has been set.

### GetThreeds2CardEnrolled

`func (o *ResponseAdditionalData3DSecure) GetThreeds2CardEnrolled() bool`

GetThreeds2CardEnrolled returns the Threeds2CardEnrolled field if non-nil, zero value otherwise.

### GetThreeds2CardEnrolledOk

`func (o *ResponseAdditionalData3DSecure) GetThreeds2CardEnrolledOk() (*bool, bool)`

GetThreeds2CardEnrolledOk returns a tuple with the Threeds2CardEnrolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeds2CardEnrolled

`func (o *ResponseAdditionalData3DSecure) SetThreeds2CardEnrolled(v bool)`

SetThreeds2CardEnrolled sets Threeds2CardEnrolled field to given value.

### HasThreeds2CardEnrolled

`func (o *ResponseAdditionalData3DSecure) HasThreeds2CardEnrolled() bool`

HasThreeds2CardEnrolled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


