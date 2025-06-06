/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the PaymentCompletionDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentCompletionDetails{}

// PaymentCompletionDetails struct for PaymentCompletionDetails
type PaymentCompletionDetails struct {
	// A payment session identifier returned by the card issuer.
	MD *string `json:"MD,omitempty"`
	// (3D) Payment Authentication Request data for the card issuer.
	PaReq *string `json:"PaReq,omitempty"`
	// (3D) Payment Authentication Response data by the card issuer.
	PaRes              *string `json:"PaRes,omitempty"`
	AuthorizationToken *string `json:"authorization_token,omitempty"`
	// PayPal-generated token for recurring payments.
	BillingToken *string `json:"billingToken,omitempty"`
	// The SMS verification code collected from the shopper.
	CupsecureplusSmscode *string `json:"cupsecureplus.smscode,omitempty"`
	// PayPal-generated third party access token.
	FacilitatorAccessToken *string `json:"facilitatorAccessToken,omitempty"`
	// A random number sent to the mobile phone number of the shopper to verify the payment.
	OneTimePasscode *string `json:"oneTimePasscode,omitempty"`
	// PayPal-assigned ID for the order.
	OrderID *string `json:"orderID,omitempty"`
	// PayPal-assigned ID for the payer (shopper).
	PayerID *string `json:"payerID,omitempty"`
	// Payload appended to the `returnURL` as a result of the redirect.
	Payload *string `json:"payload,omitempty"`
	// PayPal-generated ID for the payment.
	PaymentID *string `json:"paymentID,omitempty"`
	// Value passed from the WeChat MiniProgram `wx.requestPayment` **complete** callback. Possible values: any value starting with `requestPayment:`.
	PaymentStatus *string `json:"paymentStatus,omitempty"`
	// The result of the redirect as appended to the `returnURL`.
	RedirectResult *string `json:"redirectResult,omitempty"`
	// Value you received from the WeChat Pay SDK.
	ResultCode *string `json:"resultCode,omitempty"`
	// The query string as appended to the `returnURL` when using direct issuer links .
	ReturnUrlQueryString *string `json:"returnUrlQueryString,omitempty"`
	// Base64-encoded string returned by the Component after the challenge flow. It contains the following parameters: `transStatus`, `authorisationToken`.
	ThreeDSResult *string `json:"threeDSResult,omitempty"`
	// Base64-encoded string returned by the Component after the challenge flow. It contains the following parameter: `transStatus`.
	Threeds2ChallengeResult *string `json:"threeds2.challengeResult,omitempty"`
	// Base64-encoded string returned by the Component after the challenge flow. It contains the following parameter: `threeDSCompInd`.
	Threeds2Fingerprint *string `json:"threeds2.fingerprint,omitempty"`
	// PayPalv2-generated token for recurring payments.
	VaultToken *string `json:"vaultToken,omitempty"`
}

// NewPaymentCompletionDetails instantiates a new PaymentCompletionDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentCompletionDetails() *PaymentCompletionDetails {
	this := PaymentCompletionDetails{}
	return &this
}

// NewPaymentCompletionDetailsWithDefaults instantiates a new PaymentCompletionDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentCompletionDetailsWithDefaults() *PaymentCompletionDetails {
	this := PaymentCompletionDetails{}
	return &this
}

// GetMD returns the MD field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetMD() string {
	if o == nil || common.IsNil(o.MD) {
		var ret string
		return ret
	}
	return *o.MD
}

// GetMDOk returns a tuple with the MD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetMDOk() (*string, bool) {
	if o == nil || common.IsNil(o.MD) {
		return nil, false
	}
	return o.MD, true
}

// HasMD returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasMD() bool {
	if o != nil && !common.IsNil(o.MD) {
		return true
	}

	return false
}

// SetMD gets a reference to the given string and assigns it to the MD field.
func (o *PaymentCompletionDetails) SetMD(v string) {
	o.MD = &v
}

// GetPaReq returns the PaReq field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetPaReq() string {
	if o == nil || common.IsNil(o.PaReq) {
		var ret string
		return ret
	}
	return *o.PaReq
}

// GetPaReqOk returns a tuple with the PaReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetPaReqOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaReq) {
		return nil, false
	}
	return o.PaReq, true
}

// HasPaReq returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasPaReq() bool {
	if o != nil && !common.IsNil(o.PaReq) {
		return true
	}

	return false
}

// SetPaReq gets a reference to the given string and assigns it to the PaReq field.
func (o *PaymentCompletionDetails) SetPaReq(v string) {
	o.PaReq = &v
}

// GetPaRes returns the PaRes field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetPaRes() string {
	if o == nil || common.IsNil(o.PaRes) {
		var ret string
		return ret
	}
	return *o.PaRes
}

// GetPaResOk returns a tuple with the PaRes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetPaResOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaRes) {
		return nil, false
	}
	return o.PaRes, true
}

// HasPaRes returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasPaRes() bool {
	if o != nil && !common.IsNil(o.PaRes) {
		return true
	}

	return false
}

// SetPaRes gets a reference to the given string and assigns it to the PaRes field.
func (o *PaymentCompletionDetails) SetPaRes(v string) {
	o.PaRes = &v
}

// GetAuthorizationToken returns the AuthorizationToken field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetAuthorizationToken() string {
	if o == nil || common.IsNil(o.AuthorizationToken) {
		var ret string
		return ret
	}
	return *o.AuthorizationToken
}

// GetAuthorizationTokenOk returns a tuple with the AuthorizationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetAuthorizationTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.AuthorizationToken) {
		return nil, false
	}
	return o.AuthorizationToken, true
}

// HasAuthorizationToken returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasAuthorizationToken() bool {
	if o != nil && !common.IsNil(o.AuthorizationToken) {
		return true
	}

	return false
}

// SetAuthorizationToken gets a reference to the given string and assigns it to the AuthorizationToken field.
func (o *PaymentCompletionDetails) SetAuthorizationToken(v string) {
	o.AuthorizationToken = &v
}

// GetBillingToken returns the BillingToken field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetBillingToken() string {
	if o == nil || common.IsNil(o.BillingToken) {
		var ret string
		return ret
	}
	return *o.BillingToken
}

// GetBillingTokenOk returns a tuple with the BillingToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetBillingTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.BillingToken) {
		return nil, false
	}
	return o.BillingToken, true
}

// HasBillingToken returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasBillingToken() bool {
	if o != nil && !common.IsNil(o.BillingToken) {
		return true
	}

	return false
}

// SetBillingToken gets a reference to the given string and assigns it to the BillingToken field.
func (o *PaymentCompletionDetails) SetBillingToken(v string) {
	o.BillingToken = &v
}

// GetCupsecureplusSmscode returns the CupsecureplusSmscode field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetCupsecureplusSmscode() string {
	if o == nil || common.IsNil(o.CupsecureplusSmscode) {
		var ret string
		return ret
	}
	return *o.CupsecureplusSmscode
}

// GetCupsecureplusSmscodeOk returns a tuple with the CupsecureplusSmscode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetCupsecureplusSmscodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CupsecureplusSmscode) {
		return nil, false
	}
	return o.CupsecureplusSmscode, true
}

// HasCupsecureplusSmscode returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasCupsecureplusSmscode() bool {
	if o != nil && !common.IsNil(o.CupsecureplusSmscode) {
		return true
	}

	return false
}

// SetCupsecureplusSmscode gets a reference to the given string and assigns it to the CupsecureplusSmscode field.
func (o *PaymentCompletionDetails) SetCupsecureplusSmscode(v string) {
	o.CupsecureplusSmscode = &v
}

// GetFacilitatorAccessToken returns the FacilitatorAccessToken field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetFacilitatorAccessToken() string {
	if o == nil || common.IsNil(o.FacilitatorAccessToken) {
		var ret string
		return ret
	}
	return *o.FacilitatorAccessToken
}

// GetFacilitatorAccessTokenOk returns a tuple with the FacilitatorAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetFacilitatorAccessTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.FacilitatorAccessToken) {
		return nil, false
	}
	return o.FacilitatorAccessToken, true
}

// HasFacilitatorAccessToken returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasFacilitatorAccessToken() bool {
	if o != nil && !common.IsNil(o.FacilitatorAccessToken) {
		return true
	}

	return false
}

// SetFacilitatorAccessToken gets a reference to the given string and assigns it to the FacilitatorAccessToken field.
func (o *PaymentCompletionDetails) SetFacilitatorAccessToken(v string) {
	o.FacilitatorAccessToken = &v
}

// GetOneTimePasscode returns the OneTimePasscode field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetOneTimePasscode() string {
	if o == nil || common.IsNil(o.OneTimePasscode) {
		var ret string
		return ret
	}
	return *o.OneTimePasscode
}

// GetOneTimePasscodeOk returns a tuple with the OneTimePasscode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetOneTimePasscodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OneTimePasscode) {
		return nil, false
	}
	return o.OneTimePasscode, true
}

// HasOneTimePasscode returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasOneTimePasscode() bool {
	if o != nil && !common.IsNil(o.OneTimePasscode) {
		return true
	}

	return false
}

// SetOneTimePasscode gets a reference to the given string and assigns it to the OneTimePasscode field.
func (o *PaymentCompletionDetails) SetOneTimePasscode(v string) {
	o.OneTimePasscode = &v
}

// GetOrderID returns the OrderID field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetOrderID() string {
	if o == nil || common.IsNil(o.OrderID) {
		var ret string
		return ret
	}
	return *o.OrderID
}

// GetOrderIDOk returns a tuple with the OrderID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetOrderIDOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderID) {
		return nil, false
	}
	return o.OrderID, true
}

// HasOrderID returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasOrderID() bool {
	if o != nil && !common.IsNil(o.OrderID) {
		return true
	}

	return false
}

// SetOrderID gets a reference to the given string and assigns it to the OrderID field.
func (o *PaymentCompletionDetails) SetOrderID(v string) {
	o.OrderID = &v
}

// GetPayerID returns the PayerID field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetPayerID() string {
	if o == nil || common.IsNil(o.PayerID) {
		var ret string
		return ret
	}
	return *o.PayerID
}

// GetPayerIDOk returns a tuple with the PayerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetPayerIDOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayerID) {
		return nil, false
	}
	return o.PayerID, true
}

// HasPayerID returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasPayerID() bool {
	if o != nil && !common.IsNil(o.PayerID) {
		return true
	}

	return false
}

// SetPayerID gets a reference to the given string and assigns it to the PayerID field.
func (o *PaymentCompletionDetails) SetPayerID(v string) {
	o.PayerID = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetPayload() string {
	if o == nil || common.IsNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetPayloadOk() (*string, bool) {
	if o == nil || common.IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasPayload() bool {
	if o != nil && !common.IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *PaymentCompletionDetails) SetPayload(v string) {
	o.Payload = &v
}

// GetPaymentID returns the PaymentID field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetPaymentID() string {
	if o == nil || common.IsNil(o.PaymentID) {
		var ret string
		return ret
	}
	return *o.PaymentID
}

// GetPaymentIDOk returns a tuple with the PaymentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetPaymentIDOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentID) {
		return nil, false
	}
	return o.PaymentID, true
}

// HasPaymentID returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasPaymentID() bool {
	if o != nil && !common.IsNil(o.PaymentID) {
		return true
	}

	return false
}

// SetPaymentID gets a reference to the given string and assigns it to the PaymentID field.
func (o *PaymentCompletionDetails) SetPaymentID(v string) {
	o.PaymentID = &v
}

// GetPaymentStatus returns the PaymentStatus field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetPaymentStatus() string {
	if o == nil || common.IsNil(o.PaymentStatus) {
		var ret string
		return ret
	}
	return *o.PaymentStatus
}

// GetPaymentStatusOk returns a tuple with the PaymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetPaymentStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentStatus) {
		return nil, false
	}
	return o.PaymentStatus, true
}

// HasPaymentStatus returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasPaymentStatus() bool {
	if o != nil && !common.IsNil(o.PaymentStatus) {
		return true
	}

	return false
}

// SetPaymentStatus gets a reference to the given string and assigns it to the PaymentStatus field.
func (o *PaymentCompletionDetails) SetPaymentStatus(v string) {
	o.PaymentStatus = &v
}

// GetRedirectResult returns the RedirectResult field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetRedirectResult() string {
	if o == nil || common.IsNil(o.RedirectResult) {
		var ret string
		return ret
	}
	return *o.RedirectResult
}

// GetRedirectResultOk returns a tuple with the RedirectResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetRedirectResultOk() (*string, bool) {
	if o == nil || common.IsNil(o.RedirectResult) {
		return nil, false
	}
	return o.RedirectResult, true
}

// HasRedirectResult returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasRedirectResult() bool {
	if o != nil && !common.IsNil(o.RedirectResult) {
		return true
	}

	return false
}

// SetRedirectResult gets a reference to the given string and assigns it to the RedirectResult field.
func (o *PaymentCompletionDetails) SetRedirectResult(v string) {
	o.RedirectResult = &v
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetResultCode() string {
	if o == nil || common.IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetResultCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasResultCode() bool {
	if o != nil && !common.IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *PaymentCompletionDetails) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetReturnUrlQueryString returns the ReturnUrlQueryString field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetReturnUrlQueryString() string {
	if o == nil || common.IsNil(o.ReturnUrlQueryString) {
		var ret string
		return ret
	}
	return *o.ReturnUrlQueryString
}

// GetReturnUrlQueryStringOk returns a tuple with the ReturnUrlQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetReturnUrlQueryStringOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReturnUrlQueryString) {
		return nil, false
	}
	return o.ReturnUrlQueryString, true
}

// HasReturnUrlQueryString returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasReturnUrlQueryString() bool {
	if o != nil && !common.IsNil(o.ReturnUrlQueryString) {
		return true
	}

	return false
}

// SetReturnUrlQueryString gets a reference to the given string and assigns it to the ReturnUrlQueryString field.
func (o *PaymentCompletionDetails) SetReturnUrlQueryString(v string) {
	o.ReturnUrlQueryString = &v
}

// GetThreeDSResult returns the ThreeDSResult field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetThreeDSResult() string {
	if o == nil || common.IsNil(o.ThreeDSResult) {
		var ret string
		return ret
	}
	return *o.ThreeDSResult
}

// GetThreeDSResultOk returns a tuple with the ThreeDSResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetThreeDSResultOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDSResult) {
		return nil, false
	}
	return o.ThreeDSResult, true
}

// HasThreeDSResult returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasThreeDSResult() bool {
	if o != nil && !common.IsNil(o.ThreeDSResult) {
		return true
	}

	return false
}

// SetThreeDSResult gets a reference to the given string and assigns it to the ThreeDSResult field.
func (o *PaymentCompletionDetails) SetThreeDSResult(v string) {
	o.ThreeDSResult = &v
}

// GetThreeds2ChallengeResult returns the Threeds2ChallengeResult field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetThreeds2ChallengeResult() string {
	if o == nil || common.IsNil(o.Threeds2ChallengeResult) {
		var ret string
		return ret
	}
	return *o.Threeds2ChallengeResult
}

// GetThreeds2ChallengeResultOk returns a tuple with the Threeds2ChallengeResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetThreeds2ChallengeResultOk() (*string, bool) {
	if o == nil || common.IsNil(o.Threeds2ChallengeResult) {
		return nil, false
	}
	return o.Threeds2ChallengeResult, true
}

// HasThreeds2ChallengeResult returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasThreeds2ChallengeResult() bool {
	if o != nil && !common.IsNil(o.Threeds2ChallengeResult) {
		return true
	}

	return false
}

// SetThreeds2ChallengeResult gets a reference to the given string and assigns it to the Threeds2ChallengeResult field.
func (o *PaymentCompletionDetails) SetThreeds2ChallengeResult(v string) {
	o.Threeds2ChallengeResult = &v
}

// GetThreeds2Fingerprint returns the Threeds2Fingerprint field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetThreeds2Fingerprint() string {
	if o == nil || common.IsNil(o.Threeds2Fingerprint) {
		var ret string
		return ret
	}
	return *o.Threeds2Fingerprint
}

// GetThreeds2FingerprintOk returns a tuple with the Threeds2Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetThreeds2FingerprintOk() (*string, bool) {
	if o == nil || common.IsNil(o.Threeds2Fingerprint) {
		return nil, false
	}
	return o.Threeds2Fingerprint, true
}

// HasThreeds2Fingerprint returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasThreeds2Fingerprint() bool {
	if o != nil && !common.IsNil(o.Threeds2Fingerprint) {
		return true
	}

	return false
}

// SetThreeds2Fingerprint gets a reference to the given string and assigns it to the Threeds2Fingerprint field.
func (o *PaymentCompletionDetails) SetThreeds2Fingerprint(v string) {
	o.Threeds2Fingerprint = &v
}

// GetVaultToken returns the VaultToken field value if set, zero value otherwise.
func (o *PaymentCompletionDetails) GetVaultToken() string {
	if o == nil || common.IsNil(o.VaultToken) {
		var ret string
		return ret
	}
	return *o.VaultToken
}

// GetVaultTokenOk returns a tuple with the VaultToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCompletionDetails) GetVaultTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.VaultToken) {
		return nil, false
	}
	return o.VaultToken, true
}

// HasVaultToken returns a boolean if a field has been set.
func (o *PaymentCompletionDetails) HasVaultToken() bool {
	if o != nil && !common.IsNil(o.VaultToken) {
		return true
	}

	return false
}

// SetVaultToken gets a reference to the given string and assigns it to the VaultToken field.
func (o *PaymentCompletionDetails) SetVaultToken(v string) {
	o.VaultToken = &v
}

func (o PaymentCompletionDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentCompletionDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MD) {
		toSerialize["MD"] = o.MD
	}
	if !common.IsNil(o.PaReq) {
		toSerialize["PaReq"] = o.PaReq
	}
	if !common.IsNil(o.PaRes) {
		toSerialize["PaRes"] = o.PaRes
	}
	if !common.IsNil(o.AuthorizationToken) {
		toSerialize["authorization_token"] = o.AuthorizationToken
	}
	if !common.IsNil(o.BillingToken) {
		toSerialize["billingToken"] = o.BillingToken
	}
	if !common.IsNil(o.CupsecureplusSmscode) {
		toSerialize["cupsecureplus.smscode"] = o.CupsecureplusSmscode
	}
	if !common.IsNil(o.FacilitatorAccessToken) {
		toSerialize["facilitatorAccessToken"] = o.FacilitatorAccessToken
	}
	if !common.IsNil(o.OneTimePasscode) {
		toSerialize["oneTimePasscode"] = o.OneTimePasscode
	}
	if !common.IsNil(o.OrderID) {
		toSerialize["orderID"] = o.OrderID
	}
	if !common.IsNil(o.PayerID) {
		toSerialize["payerID"] = o.PayerID
	}
	if !common.IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !common.IsNil(o.PaymentID) {
		toSerialize["paymentID"] = o.PaymentID
	}
	if !common.IsNil(o.PaymentStatus) {
		toSerialize["paymentStatus"] = o.PaymentStatus
	}
	if !common.IsNil(o.RedirectResult) {
		toSerialize["redirectResult"] = o.RedirectResult
	}
	if !common.IsNil(o.ResultCode) {
		toSerialize["resultCode"] = o.ResultCode
	}
	if !common.IsNil(o.ReturnUrlQueryString) {
		toSerialize["returnUrlQueryString"] = o.ReturnUrlQueryString
	}
	if !common.IsNil(o.ThreeDSResult) {
		toSerialize["threeDSResult"] = o.ThreeDSResult
	}
	if !common.IsNil(o.Threeds2ChallengeResult) {
		toSerialize["threeds2.challengeResult"] = o.Threeds2ChallengeResult
	}
	if !common.IsNil(o.Threeds2Fingerprint) {
		toSerialize["threeds2.fingerprint"] = o.Threeds2Fingerprint
	}
	if !common.IsNil(o.VaultToken) {
		toSerialize["vaultToken"] = o.VaultToken
	}
	return toSerialize, nil
}

type NullablePaymentCompletionDetails struct {
	value *PaymentCompletionDetails
	isSet bool
}

func (v NullablePaymentCompletionDetails) Get() *PaymentCompletionDetails {
	return v.value
}

func (v *NullablePaymentCompletionDetails) Set(val *PaymentCompletionDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentCompletionDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentCompletionDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentCompletionDetails(val *PaymentCompletionDetails) *NullablePaymentCompletionDetails {
	return &NullablePaymentCompletionDetails{value: val, isSet: true}
}

func (v NullablePaymentCompletionDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentCompletionDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
