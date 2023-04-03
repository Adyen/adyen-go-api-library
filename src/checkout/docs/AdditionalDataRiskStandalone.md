# AdditionalDataRiskStandalone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayPalCountryCode** | Pointer to **string** | Shopper&#39;s country of residence in the form of ISO standard 3166 2-character country codes. | [optional] 
**PayPalEmailId** | Pointer to **string** | Shopper&#39;s email. | [optional] 
**PayPalFirstName** | Pointer to **string** | Shopper&#39;s first name. | [optional] 
**PayPalLastName** | Pointer to **string** | Shopper&#39;s last name. | [optional] 
**PayPalPayerId** | Pointer to **string** | Unique PayPal Customer Account identification number. Character length and limitations: 13 single-byte alphanumeric characters. | [optional] 
**PayPalPhone** | Pointer to **string** | Shopper&#39;s phone number. | [optional] 
**PayPalProtectionEligibility** | Pointer to **string** | Allowed values: * **Eligible** — Merchant is protected by PayPal&#39;s Seller Protection Policy for Unauthorized Payments and Item Not Received.  * **PartiallyEligible** — Merchant is protected by PayPal&#39;s Seller Protection Policy for Item Not Received.  * **Ineligible** — Merchant is not protected under the Seller Protection Policy. | [optional] 
**PayPalTransactionId** | Pointer to **string** | Unique transaction ID of the payment. | [optional] 
**AvsResultRaw** | Pointer to **string** | Raw AVS result received from the acquirer, where available. Example: D | [optional] 
**Bin** | Pointer to **string** | The Bank Identification Number of a credit card, which is the first six digits of a card number. Required for [tokenized card request](https://docs.adyen.com/risk-management/standalone-risk#tokenised-pan-request). | [optional] 
**CvcResultRaw** | Pointer to **string** | Raw CVC result received from the acquirer, where available. Example: 1 | [optional] 
**RiskToken** | Pointer to **string** | Unique identifier or token for the shopper&#39;s card details. | [optional] 
**ThreeDAuthenticated** | Pointer to **string** | A Boolean value indicating whether 3DS authentication was completed on this payment. Example: true | [optional] 
**ThreeDOffered** | Pointer to **string** | A Boolean value indicating whether 3DS was offered for this payment. Example: true | [optional] 
**TokenDataType** | Pointer to **string** | Required for PayPal payments only. The only supported value is: **paypal**. | [optional] 

## Methods

### NewAdditionalDataRiskStandalone

`func NewAdditionalDataRiskStandalone() *AdditionalDataRiskStandalone`

NewAdditionalDataRiskStandalone instantiates a new AdditionalDataRiskStandalone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalDataRiskStandaloneWithDefaults

`func NewAdditionalDataRiskStandaloneWithDefaults() *AdditionalDataRiskStandalone`

NewAdditionalDataRiskStandaloneWithDefaults instantiates a new AdditionalDataRiskStandalone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayPalCountryCode

`func (o *AdditionalDataRiskStandalone) GetPayPalCountryCode() string`

GetPayPalCountryCode returns the PayPalCountryCode field if non-nil, zero value otherwise.

### GetPayPalCountryCodeOk

`func (o *AdditionalDataRiskStandalone) GetPayPalCountryCodeOk() (*string, bool)`

GetPayPalCountryCodeOk returns a tuple with the PayPalCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPalCountryCode

`func (o *AdditionalDataRiskStandalone) SetPayPalCountryCode(v string)`

SetPayPalCountryCode sets PayPalCountryCode field to given value.

### HasPayPalCountryCode

`func (o *AdditionalDataRiskStandalone) HasPayPalCountryCode() bool`

HasPayPalCountryCode returns a boolean if a field has been set.

### GetPayPalEmailId

`func (o *AdditionalDataRiskStandalone) GetPayPalEmailId() string`

GetPayPalEmailId returns the PayPalEmailId field if non-nil, zero value otherwise.

### GetPayPalEmailIdOk

`func (o *AdditionalDataRiskStandalone) GetPayPalEmailIdOk() (*string, bool)`

GetPayPalEmailIdOk returns a tuple with the PayPalEmailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPalEmailId

`func (o *AdditionalDataRiskStandalone) SetPayPalEmailId(v string)`

SetPayPalEmailId sets PayPalEmailId field to given value.

### HasPayPalEmailId

`func (o *AdditionalDataRiskStandalone) HasPayPalEmailId() bool`

HasPayPalEmailId returns a boolean if a field has been set.

### GetPayPalFirstName

`func (o *AdditionalDataRiskStandalone) GetPayPalFirstName() string`

GetPayPalFirstName returns the PayPalFirstName field if non-nil, zero value otherwise.

### GetPayPalFirstNameOk

`func (o *AdditionalDataRiskStandalone) GetPayPalFirstNameOk() (*string, bool)`

GetPayPalFirstNameOk returns a tuple with the PayPalFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPalFirstName

`func (o *AdditionalDataRiskStandalone) SetPayPalFirstName(v string)`

SetPayPalFirstName sets PayPalFirstName field to given value.

### HasPayPalFirstName

`func (o *AdditionalDataRiskStandalone) HasPayPalFirstName() bool`

HasPayPalFirstName returns a boolean if a field has been set.

### GetPayPalLastName

`func (o *AdditionalDataRiskStandalone) GetPayPalLastName() string`

GetPayPalLastName returns the PayPalLastName field if non-nil, zero value otherwise.

### GetPayPalLastNameOk

`func (o *AdditionalDataRiskStandalone) GetPayPalLastNameOk() (*string, bool)`

GetPayPalLastNameOk returns a tuple with the PayPalLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPalLastName

`func (o *AdditionalDataRiskStandalone) SetPayPalLastName(v string)`

SetPayPalLastName sets PayPalLastName field to given value.

### HasPayPalLastName

`func (o *AdditionalDataRiskStandalone) HasPayPalLastName() bool`

HasPayPalLastName returns a boolean if a field has been set.

### GetPayPalPayerId

`func (o *AdditionalDataRiskStandalone) GetPayPalPayerId() string`

GetPayPalPayerId returns the PayPalPayerId field if non-nil, zero value otherwise.

### GetPayPalPayerIdOk

`func (o *AdditionalDataRiskStandalone) GetPayPalPayerIdOk() (*string, bool)`

GetPayPalPayerIdOk returns a tuple with the PayPalPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPalPayerId

`func (o *AdditionalDataRiskStandalone) SetPayPalPayerId(v string)`

SetPayPalPayerId sets PayPalPayerId field to given value.

### HasPayPalPayerId

`func (o *AdditionalDataRiskStandalone) HasPayPalPayerId() bool`

HasPayPalPayerId returns a boolean if a field has been set.

### GetPayPalPhone

`func (o *AdditionalDataRiskStandalone) GetPayPalPhone() string`

GetPayPalPhone returns the PayPalPhone field if non-nil, zero value otherwise.

### GetPayPalPhoneOk

`func (o *AdditionalDataRiskStandalone) GetPayPalPhoneOk() (*string, bool)`

GetPayPalPhoneOk returns a tuple with the PayPalPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPalPhone

`func (o *AdditionalDataRiskStandalone) SetPayPalPhone(v string)`

SetPayPalPhone sets PayPalPhone field to given value.

### HasPayPalPhone

`func (o *AdditionalDataRiskStandalone) HasPayPalPhone() bool`

HasPayPalPhone returns a boolean if a field has been set.

### GetPayPalProtectionEligibility

`func (o *AdditionalDataRiskStandalone) GetPayPalProtectionEligibility() string`

GetPayPalProtectionEligibility returns the PayPalProtectionEligibility field if non-nil, zero value otherwise.

### GetPayPalProtectionEligibilityOk

`func (o *AdditionalDataRiskStandalone) GetPayPalProtectionEligibilityOk() (*string, bool)`

GetPayPalProtectionEligibilityOk returns a tuple with the PayPalProtectionEligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPalProtectionEligibility

`func (o *AdditionalDataRiskStandalone) SetPayPalProtectionEligibility(v string)`

SetPayPalProtectionEligibility sets PayPalProtectionEligibility field to given value.

### HasPayPalProtectionEligibility

`func (o *AdditionalDataRiskStandalone) HasPayPalProtectionEligibility() bool`

HasPayPalProtectionEligibility returns a boolean if a field has been set.

### GetPayPalTransactionId

`func (o *AdditionalDataRiskStandalone) GetPayPalTransactionId() string`

GetPayPalTransactionId returns the PayPalTransactionId field if non-nil, zero value otherwise.

### GetPayPalTransactionIdOk

`func (o *AdditionalDataRiskStandalone) GetPayPalTransactionIdOk() (*string, bool)`

GetPayPalTransactionIdOk returns a tuple with the PayPalTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPalTransactionId

`func (o *AdditionalDataRiskStandalone) SetPayPalTransactionId(v string)`

SetPayPalTransactionId sets PayPalTransactionId field to given value.

### HasPayPalTransactionId

`func (o *AdditionalDataRiskStandalone) HasPayPalTransactionId() bool`

HasPayPalTransactionId returns a boolean if a field has been set.

### GetAvsResultRaw

`func (o *AdditionalDataRiskStandalone) GetAvsResultRaw() string`

GetAvsResultRaw returns the AvsResultRaw field if non-nil, zero value otherwise.

### GetAvsResultRawOk

`func (o *AdditionalDataRiskStandalone) GetAvsResultRawOk() (*string, bool)`

GetAvsResultRawOk returns a tuple with the AvsResultRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsResultRaw

`func (o *AdditionalDataRiskStandalone) SetAvsResultRaw(v string)`

SetAvsResultRaw sets AvsResultRaw field to given value.

### HasAvsResultRaw

`func (o *AdditionalDataRiskStandalone) HasAvsResultRaw() bool`

HasAvsResultRaw returns a boolean if a field has been set.

### GetBin

`func (o *AdditionalDataRiskStandalone) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *AdditionalDataRiskStandalone) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *AdditionalDataRiskStandalone) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *AdditionalDataRiskStandalone) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetCvcResultRaw

`func (o *AdditionalDataRiskStandalone) GetCvcResultRaw() string`

GetCvcResultRaw returns the CvcResultRaw field if non-nil, zero value otherwise.

### GetCvcResultRawOk

`func (o *AdditionalDataRiskStandalone) GetCvcResultRawOk() (*string, bool)`

GetCvcResultRawOk returns a tuple with the CvcResultRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvcResultRaw

`func (o *AdditionalDataRiskStandalone) SetCvcResultRaw(v string)`

SetCvcResultRaw sets CvcResultRaw field to given value.

### HasCvcResultRaw

`func (o *AdditionalDataRiskStandalone) HasCvcResultRaw() bool`

HasCvcResultRaw returns a boolean if a field has been set.

### GetRiskToken

`func (o *AdditionalDataRiskStandalone) GetRiskToken() string`

GetRiskToken returns the RiskToken field if non-nil, zero value otherwise.

### GetRiskTokenOk

`func (o *AdditionalDataRiskStandalone) GetRiskTokenOk() (*string, bool)`

GetRiskTokenOk returns a tuple with the RiskToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskToken

`func (o *AdditionalDataRiskStandalone) SetRiskToken(v string)`

SetRiskToken sets RiskToken field to given value.

### HasRiskToken

`func (o *AdditionalDataRiskStandalone) HasRiskToken() bool`

HasRiskToken returns a boolean if a field has been set.

### GetThreeDAuthenticated

`func (o *AdditionalDataRiskStandalone) GetThreeDAuthenticated() string`

GetThreeDAuthenticated returns the ThreeDAuthenticated field if non-nil, zero value otherwise.

### GetThreeDAuthenticatedOk

`func (o *AdditionalDataRiskStandalone) GetThreeDAuthenticatedOk() (*string, bool)`

GetThreeDAuthenticatedOk returns a tuple with the ThreeDAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDAuthenticated

`func (o *AdditionalDataRiskStandalone) SetThreeDAuthenticated(v string)`

SetThreeDAuthenticated sets ThreeDAuthenticated field to given value.

### HasThreeDAuthenticated

`func (o *AdditionalDataRiskStandalone) HasThreeDAuthenticated() bool`

HasThreeDAuthenticated returns a boolean if a field has been set.

### GetThreeDOffered

`func (o *AdditionalDataRiskStandalone) GetThreeDOffered() string`

GetThreeDOffered returns the ThreeDOffered field if non-nil, zero value otherwise.

### GetThreeDOfferedOk

`func (o *AdditionalDataRiskStandalone) GetThreeDOfferedOk() (*string, bool)`

GetThreeDOfferedOk returns a tuple with the ThreeDOffered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDOffered

`func (o *AdditionalDataRiskStandalone) SetThreeDOffered(v string)`

SetThreeDOffered sets ThreeDOffered field to given value.

### HasThreeDOffered

`func (o *AdditionalDataRiskStandalone) HasThreeDOffered() bool`

HasThreeDOffered returns a boolean if a field has been set.

### GetTokenDataType

`func (o *AdditionalDataRiskStandalone) GetTokenDataType() string`

GetTokenDataType returns the TokenDataType field if non-nil, zero value otherwise.

### GetTokenDataTypeOk

`func (o *AdditionalDataRiskStandalone) GetTokenDataTypeOk() (*string, bool)`

GetTokenDataTypeOk returns a tuple with the TokenDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDataType

`func (o *AdditionalDataRiskStandalone) SetTokenDataType(v string)`

SetTokenDataType sets TokenDataType field to given value.

### HasTokenDataType

`func (o *AdditionalDataRiskStandalone) HasTokenDataType() bool`

HasTokenDataType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


