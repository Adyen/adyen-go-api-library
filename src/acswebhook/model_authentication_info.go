/*
Authentication webhooks

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package acswebhook

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the AuthenticationInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AuthenticationInfo{}

// AuthenticationInfo struct for AuthenticationInfo
type AuthenticationInfo struct {
	// Universally unique transaction identifier assigned by the Access Control Server (ACS) to identify a single transaction.
	AcsTransId string         `json:"acsTransId"`
	Challenge  *ChallengeInfo `json:"challenge,omitempty"`
	// Specifies a preference for receiving a challenge. Possible values:  * **01**: No preference * **02**: No challenge requested * **03**: Challenge requested (preference) * **04**: Challenge requested (mandate) * **05**: No challenge requested (transactional risk analysis is already performed) * **07**: No challenge requested (SCA is already performed) * **08**: No challenge requested (trusted beneficiaries exemption of no challenge required) * **09**: Challenge requested (trusted beneficiaries prompt requested if challenge required) * **80**: No challenge requested (secure corporate payment with Mastercard) * **82**: No challenge requested (secure corporate payment with Visa)
	ChallengeIndicator string `json:"challengeIndicator"`
	// Date and time in UTC of the cardholder authentication.   [ISO 8601](https://www.w3.org/TR/NOTE-datetime) format: YYYY-MM-DDThh:mm:ss+TZD, for example, **2020-12-18T10:15:30+01:00**.
	CreatedAt time.Time `json:"createdAt"`
	// Indicates the type of channel interface being used to initiate the transaction. Possible values:  * **app** * **browser** * **3DSRequestorInitiated** (initiated by a merchant when the cardholder is not available)
	DeviceChannel string `json:"deviceChannel"`
	// Universally unique transaction identifier assigned by the DS (card scheme) to identify a single transaction.
	DsTransID string `json:"dsTransID"`
	// Indicates the exemption type that was applied to the authentication by the issuer, if exemption applied. Possible values:  * **lowValue** * **secureCorporate** * **trustedBeneficiary** * **transactionRiskAnalysis** * **acquirerExemption** * **noExemptionApplied** * **visaDAFExemption**
	ExemptionIndicator *string `json:"exemptionIndicator,omitempty"`
	// Indicates if the purchase was in the PSD2 scope.
	InPSD2Scope bool `json:"inPSD2Scope"`
	// Identifies the category of the message for a specific use case. Possible values:  * **payment** * **nonPayment**
	MessageCategory string `json:"messageCategory"`
	// The `messageVersion` value as defined in the 3D Secure 2 specification.
	MessageVersion string `json:"messageVersion"`
	// Risk score calculated from the transaction rules.
	RiskScore *int32 `json:"riskScore,omitempty"`
	// The `threeDSServerTransID` value as defined in the 3D Secure 2 specification.
	ThreeDSServerTransID string `json:"threeDSServerTransID"`
	// The `transStatus` value as defined in the 3D Secure 2 specification. Possible values:  * **Y**: Authentication / Account verification successful. * **N**: Not Authenticated / Account not verified. Transaction denied. * **U**: Authentication / Account verification could not be performed. * **I**: Informational Only / 3D Secure Requestor challenge preference acknowledged. * **R**: Authentication / Account verification rejected by the Issuer.
	TransStatus string `json:"transStatus"`
	// Provides information on why the `transStatus` field has the specified value. For possible values, refer to [our docs](https://docs.adyen.com/online-payments/3d-secure/api-reference#possible-transstatusreason-values).
	TransStatusReason *string `json:"transStatusReason,omitempty"`
	// The type of authentication performed. Possible values:  * **frictionless** * **challenge**
	Type string `json:"type"`
}

// NewAuthenticationInfo instantiates a new AuthenticationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationInfo(acsTransId string, challengeIndicator string, createdAt time.Time, deviceChannel string, dsTransID string, inPSD2Scope bool, messageCategory string, messageVersion string, threeDSServerTransID string, transStatus string, type_ string) *AuthenticationInfo {
	this := AuthenticationInfo{}
	this.AcsTransId = acsTransId
	this.ChallengeIndicator = challengeIndicator
	this.CreatedAt = createdAt
	this.DeviceChannel = deviceChannel
	this.DsTransID = dsTransID
	this.InPSD2Scope = inPSD2Scope
	this.MessageCategory = messageCategory
	this.MessageVersion = messageVersion
	this.ThreeDSServerTransID = threeDSServerTransID
	this.TransStatus = transStatus
	this.Type = type_
	return &this
}

// NewAuthenticationInfoWithDefaults instantiates a new AuthenticationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationInfoWithDefaults() *AuthenticationInfo {
	this := AuthenticationInfo{}
	return &this
}

// GetAcsTransId returns the AcsTransId field value
func (o *AuthenticationInfo) GetAcsTransId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AcsTransId
}

// GetAcsTransIdOk returns a tuple with the AcsTransId field value
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetAcsTransIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcsTransId, true
}

// SetAcsTransId sets field value
func (o *AuthenticationInfo) SetAcsTransId(v string) {
	o.AcsTransId = v
}

// GetChallenge returns the Challenge field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetChallenge() ChallengeInfo {
	if o == nil || common.IsNil(o.Challenge) {
		var ret ChallengeInfo
		return ret
	}
	return *o.Challenge
}

// GetChallengeOk returns a tuple with the Challenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetChallengeOk() (*ChallengeInfo, bool) {
	if o == nil || common.IsNil(o.Challenge) {
		return nil, false
	}
	return o.Challenge, true
}

// HasChallenge returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasChallenge() bool {
	if o != nil && !common.IsNil(o.Challenge) {
		return true
	}

	return false
}

// SetChallenge gets a reference to the given ChallengeInfo and assigns it to the Challenge field.
func (o *AuthenticationInfo) SetChallenge(v ChallengeInfo) {
	o.Challenge = &v
}

// GetChallengeIndicator returns the ChallengeIndicator field value
func (o *AuthenticationInfo) GetChallengeIndicator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChallengeIndicator
}

// GetChallengeIndicatorOk returns a tuple with the ChallengeIndicator field value
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetChallengeIndicatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChallengeIndicator, true
}

// SetChallengeIndicator sets field value
func (o *AuthenticationInfo) SetChallengeIndicator(v string) {
	o.ChallengeIndicator = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AuthenticationInfo) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AuthenticationInfo) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDeviceChannel returns the DeviceChannel field value
func (o *AuthenticationInfo) GetDeviceChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceChannel
}

// GetDeviceChannelOk returns a tuple with the DeviceChannel field value
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetDeviceChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceChannel, true
}

// SetDeviceChannel sets field value
func (o *AuthenticationInfo) SetDeviceChannel(v string) {
	o.DeviceChannel = v
}

// GetDsTransID returns the DsTransID field value
func (o *AuthenticationInfo) GetDsTransID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DsTransID
}

// GetDsTransIDOk returns a tuple with the DsTransID field value
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetDsTransIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DsTransID, true
}

// SetDsTransID sets field value
func (o *AuthenticationInfo) SetDsTransID(v string) {
	o.DsTransID = v
}

// GetExemptionIndicator returns the ExemptionIndicator field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetExemptionIndicator() string {
	if o == nil || common.IsNil(o.ExemptionIndicator) {
		var ret string
		return ret
	}
	return *o.ExemptionIndicator
}

// GetExemptionIndicatorOk returns a tuple with the ExemptionIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetExemptionIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExemptionIndicator) {
		return nil, false
	}
	return o.ExemptionIndicator, true
}

// HasExemptionIndicator returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasExemptionIndicator() bool {
	if o != nil && !common.IsNil(o.ExemptionIndicator) {
		return true
	}

	return false
}

// SetExemptionIndicator gets a reference to the given string and assigns it to the ExemptionIndicator field.
func (o *AuthenticationInfo) SetExemptionIndicator(v string) {
	o.ExemptionIndicator = &v
}

// GetInPSD2Scope returns the InPSD2Scope field value
func (o *AuthenticationInfo) GetInPSD2Scope() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InPSD2Scope
}

// GetInPSD2ScopeOk returns a tuple with the InPSD2Scope field value
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetInPSD2ScopeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InPSD2Scope, true
}

// SetInPSD2Scope sets field value
func (o *AuthenticationInfo) SetInPSD2Scope(v bool) {
	o.InPSD2Scope = v
}

// GetMessageCategory returns the MessageCategory field value
func (o *AuthenticationInfo) GetMessageCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageCategory
}

// GetMessageCategoryOk returns a tuple with the MessageCategory field value
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetMessageCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageCategory, true
}

// SetMessageCategory sets field value
func (o *AuthenticationInfo) SetMessageCategory(v string) {
	o.MessageCategory = v
}

// GetMessageVersion returns the MessageVersion field value
func (o *AuthenticationInfo) GetMessageVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageVersion
}

// GetMessageVersionOk returns a tuple with the MessageVersion field value
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetMessageVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageVersion, true
}

// SetMessageVersion sets field value
func (o *AuthenticationInfo) SetMessageVersion(v string) {
	o.MessageVersion = v
}

// GetRiskScore returns the RiskScore field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetRiskScore() int32 {
	if o == nil || common.IsNil(o.RiskScore) {
		var ret int32
		return ret
	}
	return *o.RiskScore
}

// GetRiskScoreOk returns a tuple with the RiskScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetRiskScoreOk() (*int32, bool) {
	if o == nil || common.IsNil(o.RiskScore) {
		return nil, false
	}
	return o.RiskScore, true
}

// HasRiskScore returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasRiskScore() bool {
	if o != nil && !common.IsNil(o.RiskScore) {
		return true
	}

	return false
}

// SetRiskScore gets a reference to the given int32 and assigns it to the RiskScore field.
func (o *AuthenticationInfo) SetRiskScore(v int32) {
	o.RiskScore = &v
}

// GetThreeDSServerTransID returns the ThreeDSServerTransID field value
func (o *AuthenticationInfo) GetThreeDSServerTransID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreeDSServerTransID
}

// GetThreeDSServerTransIDOk returns a tuple with the ThreeDSServerTransID field value
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetThreeDSServerTransIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreeDSServerTransID, true
}

// SetThreeDSServerTransID sets field value
func (o *AuthenticationInfo) SetThreeDSServerTransID(v string) {
	o.ThreeDSServerTransID = v
}

// GetTransStatus returns the TransStatus field value
func (o *AuthenticationInfo) GetTransStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransStatus
}

// GetTransStatusOk returns a tuple with the TransStatus field value
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetTransStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransStatus, true
}

// SetTransStatus sets field value
func (o *AuthenticationInfo) SetTransStatus(v string) {
	o.TransStatus = v
}

// GetTransStatusReason returns the TransStatusReason field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetTransStatusReason() string {
	if o == nil || common.IsNil(o.TransStatusReason) {
		var ret string
		return ret
	}
	return *o.TransStatusReason
}

// GetTransStatusReasonOk returns a tuple with the TransStatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetTransStatusReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransStatusReason) {
		return nil, false
	}
	return o.TransStatusReason, true
}

// HasTransStatusReason returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasTransStatusReason() bool {
	if o != nil && !common.IsNil(o.TransStatusReason) {
		return true
	}

	return false
}

// SetTransStatusReason gets a reference to the given string and assigns it to the TransStatusReason field.
func (o *AuthenticationInfo) SetTransStatusReason(v string) {
	o.TransStatusReason = &v
}

// GetType returns the Type field value
func (o *AuthenticationInfo) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthenticationInfo) SetType(v string) {
	o.Type = v
}

func (o AuthenticationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["acsTransId"] = o.AcsTransId
	if !common.IsNil(o.Challenge) {
		toSerialize["challenge"] = o.Challenge
	}
	toSerialize["challengeIndicator"] = o.ChallengeIndicator
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["deviceChannel"] = o.DeviceChannel
	toSerialize["dsTransID"] = o.DsTransID
	if !common.IsNil(o.ExemptionIndicator) {
		toSerialize["exemptionIndicator"] = o.ExemptionIndicator
	}
	toSerialize["inPSD2Scope"] = o.InPSD2Scope
	toSerialize["messageCategory"] = o.MessageCategory
	toSerialize["messageVersion"] = o.MessageVersion
	if !common.IsNil(o.RiskScore) {
		toSerialize["riskScore"] = o.RiskScore
	}
	toSerialize["threeDSServerTransID"] = o.ThreeDSServerTransID
	toSerialize["transStatus"] = o.TransStatus
	if !common.IsNil(o.TransStatusReason) {
		toSerialize["transStatusReason"] = o.TransStatusReason
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAuthenticationInfo struct {
	value *AuthenticationInfo
	isSet bool
}

func (v NullableAuthenticationInfo) Get() *AuthenticationInfo {
	return v.value
}

func (v *NullableAuthenticationInfo) Set(val *AuthenticationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationInfo(val *AuthenticationInfo) *NullableAuthenticationInfo {
	return &NullableAuthenticationInfo{value: val, isSet: true}
}

func (v NullableAuthenticationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AuthenticationInfo) isValidChallengeIndicator() bool {
	var allowedEnumValues = []string{"01", "02", "03", "04", "05", "07", "08", "09", "80", "82"}
	for _, allowed := range allowedEnumValues {
		if o.GetChallengeIndicator() == allowed {
			return true
		}
	}
	return false
}
func (o *AuthenticationInfo) isValidDeviceChannel() bool {
	var allowedEnumValues = []string{"app", "browser", "ThreeDSRequestorInitiated"}
	for _, allowed := range allowedEnumValues {
		if o.GetDeviceChannel() == allowed {
			return true
		}
	}
	return false
}
func (o *AuthenticationInfo) isValidExemptionIndicator() bool {
	var allowedEnumValues = []string{"lowValue", "secureCorporate", "trustedBeneficiary", "transactionRiskAnalysis", "acquirerExemption", "noExemptionApplied", "visaDAFExemption"}
	for _, allowed := range allowedEnumValues {
		if o.GetExemptionIndicator() == allowed {
			return true
		}
	}
	return false
}
func (o *AuthenticationInfo) isValidMessageCategory() bool {
	var allowedEnumValues = []string{"payment", "nonPayment"}
	for _, allowed := range allowedEnumValues {
		if o.GetMessageCategory() == allowed {
			return true
		}
	}
	return false
}
func (o *AuthenticationInfo) isValidTransStatus() bool {
	var allowedEnumValues = []string{"Y", "N", "R", "I", "U"}
	for _, allowed := range allowedEnumValues {
		if o.GetTransStatus() == allowed {
			return true
		}
	}
	return false
}
func (o *AuthenticationInfo) isValidTransStatusReason() bool {
	var allowedEnumValues = []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "80", "81", "82", "83", "84", "85", "86", "87", "88"}
	for _, allowed := range allowedEnumValues {
		if o.GetTransStatusReason() == allowed {
			return true
		}
	}
	return false
}
func (o *AuthenticationInfo) isValidType() bool {
	var allowedEnumValues = []string{"frictionless", "challenge"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}