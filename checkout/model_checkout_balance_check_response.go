/*
Adyen Checkout API

Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [online payments documentation](https://docs.adyen.com/online-payments).  ## Authentication Each request to Checkout API must be signed with an API key. For this, [get your API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key) from your Customer Area, and set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: YOUR_API_KEY\" \\ ... ``` ## Versioning Checkout API supports [versioning](https://docs.adyen.com/development-resources/versioning) using a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v70/payments ```  ## Going live  To access the live endpoints, you need an API key from your live Customer Area.  The live endpoint URLs contain a prefix which is unique to your company account, for example: ``` https://{PREFIX}-checkout-live.adyenpayments.com/checkout/v70/payments ```  Get your `{PREFIX}` from your live Customer Area under **Developers** > **API URLs** > **Prefix**.  When preparing to do live transactions with Checkout API, follow the [go-live checklist](https://docs.adyen.com/online-payments/go-live-checklist) to make sure you've got all the required configuration in place.  ## Release notes Have a look at the [release notes](https://docs.adyen.com/online-payments/release-notes?integration_type=api&version=70) to find out what changed in this version!

API version: 70
Contact: developer-experience@adyen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"
)

// checks if the CheckoutBalanceCheckResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutBalanceCheckResponse{}

// CheckoutBalanceCheckResponse struct for CheckoutBalanceCheckResponse
type CheckoutBalanceCheckResponse struct {
	// Contains additional information about the payment. Some data fields are included only if you select them first: Go to **Customer Area** > **Developers** > **Additional data**.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	Balance        Amount             `json:"balance"`
	FraudResult    *FraudResult       `json:"fraudResult,omitempty"`
	// Adyen's 16-character reference associated with the transaction/request. This value is globally unique; quote it when communicating with us about this request.
	PspReference *string `json:"pspReference,omitempty"`
	// If the payment's authorisation is refused or an error occurs during authorisation, this field holds Adyen's mapped reason for the refusal or a description of the error. When a transaction fails, the authorisation response includes `resultCode` and `refusalReason` values.  For more information, see [Refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons).
	RefusalReason *string `json:"refusalReason,omitempty"`
	// The result of the cancellation request.  Possible values:  * **Success** – Indicates that the balance check was successful. * **NotEnoughBalance** – Commonly indicates that the card did not have enough balance to pay the amount in the request, or that the currency of the balance on the card did not match the currency of the requested amount. * **Failed** – Indicates that the balance check failed.
	ResultCode       string  `json:"resultCode"`
	TransactionLimit *Amount `json:"transactionLimit,omitempty"`
}

// NewCheckoutBalanceCheckResponse instantiates a new CheckoutBalanceCheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutBalanceCheckResponse(balance Amount, resultCode string) *CheckoutBalanceCheckResponse {
	this := CheckoutBalanceCheckResponse{}
	this.Balance = balance
	this.ResultCode = resultCode
	return &this
}

// NewCheckoutBalanceCheckResponseWithDefaults instantiates a new CheckoutBalanceCheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutBalanceCheckResponseWithDefaults() *CheckoutBalanceCheckResponse {
	this := CheckoutBalanceCheckResponse{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *CheckoutBalanceCheckResponse) GetAdditionalData() map[string]string {
	if o == nil || IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBalanceCheckResponse) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *CheckoutBalanceCheckResponse) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *CheckoutBalanceCheckResponse) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetBalance returns the Balance field value
func (o *CheckoutBalanceCheckResponse) GetBalance() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *CheckoutBalanceCheckResponse) GetBalanceOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *CheckoutBalanceCheckResponse) SetBalance(v Amount) {
	o.Balance = v
}

// GetFraudResult returns the FraudResult field value if set, zero value otherwise.
func (o *CheckoutBalanceCheckResponse) GetFraudResult() FraudResult {
	if o == nil || IsNil(o.FraudResult) {
		var ret FraudResult
		return ret
	}
	return *o.FraudResult
}

// GetFraudResultOk returns a tuple with the FraudResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBalanceCheckResponse) GetFraudResultOk() (*FraudResult, bool) {
	if o == nil || IsNil(o.FraudResult) {
		return nil, false
	}
	return o.FraudResult, true
}

// HasFraudResult returns a boolean if a field has been set.
func (o *CheckoutBalanceCheckResponse) HasFraudResult() bool {
	if o != nil && !IsNil(o.FraudResult) {
		return true
	}

	return false
}

// SetFraudResult gets a reference to the given FraudResult and assigns it to the FraudResult field.
func (o *CheckoutBalanceCheckResponse) SetFraudResult(v FraudResult) {
	o.FraudResult = &v
}

// GetPspReference returns the PspReference field value if set, zero value otherwise.
func (o *CheckoutBalanceCheckResponse) GetPspReference() string {
	if o == nil || IsNil(o.PspReference) {
		var ret string
		return ret
	}
	return *o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBalanceCheckResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.PspReference) {
		return nil, false
	}
	return o.PspReference, true
}

// HasPspReference returns a boolean if a field has been set.
func (o *CheckoutBalanceCheckResponse) HasPspReference() bool {
	if o != nil && !IsNil(o.PspReference) {
		return true
	}

	return false
}

// SetPspReference gets a reference to the given string and assigns it to the PspReference field.
func (o *CheckoutBalanceCheckResponse) SetPspReference(v string) {
	o.PspReference = &v
}

// GetRefusalReason returns the RefusalReason field value if set, zero value otherwise.
func (o *CheckoutBalanceCheckResponse) GetRefusalReason() string {
	if o == nil || IsNil(o.RefusalReason) {
		var ret string
		return ret
	}
	return *o.RefusalReason
}

// GetRefusalReasonOk returns a tuple with the RefusalReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBalanceCheckResponse) GetRefusalReasonOk() (*string, bool) {
	if o == nil || IsNil(o.RefusalReason) {
		return nil, false
	}
	return o.RefusalReason, true
}

// HasRefusalReason returns a boolean if a field has been set.
func (o *CheckoutBalanceCheckResponse) HasRefusalReason() bool {
	if o != nil && !IsNil(o.RefusalReason) {
		return true
	}

	return false
}

// SetRefusalReason gets a reference to the given string and assigns it to the RefusalReason field.
func (o *CheckoutBalanceCheckResponse) SetRefusalReason(v string) {
	o.RefusalReason = &v
}

// GetResultCode returns the ResultCode field value
func (o *CheckoutBalanceCheckResponse) GetResultCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value
// and a boolean to check if the value has been set.
func (o *CheckoutBalanceCheckResponse) GetResultCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultCode, true
}

// SetResultCode sets field value
func (o *CheckoutBalanceCheckResponse) SetResultCode(v string) {
	o.ResultCode = v
}

// GetTransactionLimit returns the TransactionLimit field value if set, zero value otherwise.
func (o *CheckoutBalanceCheckResponse) GetTransactionLimit() Amount {
	if o == nil || IsNil(o.TransactionLimit) {
		var ret Amount
		return ret
	}
	return *o.TransactionLimit
}

// GetTransactionLimitOk returns a tuple with the TransactionLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBalanceCheckResponse) GetTransactionLimitOk() (*Amount, bool) {
	if o == nil || IsNil(o.TransactionLimit) {
		return nil, false
	}
	return o.TransactionLimit, true
}

// HasTransactionLimit returns a boolean if a field has been set.
func (o *CheckoutBalanceCheckResponse) HasTransactionLimit() bool {
	if o != nil && !IsNil(o.TransactionLimit) {
		return true
	}

	return false
}

// SetTransactionLimit gets a reference to the given Amount and assigns it to the TransactionLimit field.
func (o *CheckoutBalanceCheckResponse) SetTransactionLimit(v Amount) {
	o.TransactionLimit = &v
}

func (o CheckoutBalanceCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutBalanceCheckResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	toSerialize["balance"] = o.Balance
	if !IsNil(o.FraudResult) {
		toSerialize["fraudResult"] = o.FraudResult
	}
	if !IsNil(o.PspReference) {
		toSerialize["pspReference"] = o.PspReference
	}
	if !IsNil(o.RefusalReason) {
		toSerialize["refusalReason"] = o.RefusalReason
	}
	toSerialize["resultCode"] = o.ResultCode
	if !IsNil(o.TransactionLimit) {
		toSerialize["transactionLimit"] = o.TransactionLimit
	}
	return toSerialize, nil
}

type NullableCheckoutBalanceCheckResponse struct {
	value *CheckoutBalanceCheckResponse
	isSet bool
}

func (v NullableCheckoutBalanceCheckResponse) Get() *CheckoutBalanceCheckResponse {
	return v.value
}

func (v *NullableCheckoutBalanceCheckResponse) Set(val *CheckoutBalanceCheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutBalanceCheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutBalanceCheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutBalanceCheckResponse(val *CheckoutBalanceCheckResponse) *NullableCheckoutBalanceCheckResponse {
	return &NullableCheckoutBalanceCheckResponse{value: val, isSet: true}
}

func (v NullableCheckoutBalanceCheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutBalanceCheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CheckoutBalanceCheckResponse) isValidResultCode() bool {
	var allowedEnumValues = []string{"Success", "NotEnoughBalance", "Failed"}
	for _, allowed := range allowedEnumValues {
		if o.GetResultCode() == allowed {
			return true
		}
	}
	return false
}