/*
 * Adyen for Platforms: Notifications
 *
 * The Notification API sends notifications to the endpoints specified in a given subscription. Subscriptions are managed through the Notification Configuration API. The API specifications listed here detail the format of each notification.  For more information, refer to our [documentation](https://docs.adyen.com/platforms/notifications).
 *
 * API version: 6
 * Contact: support@adyen.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package platformsnotificationevents

import (
	"encoding/json"
)

// AccountHolderDetails struct for AccountHolderDetails
type AccountHolderDetails struct {
	Address *ViasAddress `json:"address,omitempty"`
	// Each of the bank accounts associated with the account holder. > Each array entry should represent one bank account. > For comprehensive detail regarding the required `BankAccountDetail` fields, please refer to the [KYC Verification documentation](https://docs.adyen.com/platforms/onboarding-and-verification/verification-checks).
	BankAccountDetails *[]BankAccountDetail `json:"bankAccountDetails,omitempty"`
	// The opaque reference value returned by the Adyen API during bank account login.
	BankAggregatorDataReference *string          `json:"bankAggregatorDataReference,omitempty"`
	BusinessDetails             *BusinessDetails `json:"businessDetails,omitempty"`
	// The email address of the account holder.
	Email string `json:"email"`
	// The phone number of the account holder provided as a single string. It will be handled as a landline phone. **Examples:** \"0031 6 11 22 33 44\", \"+316/1122-3344\", \"(0031) 611223344\"
	FullPhoneNumber   string             `json:"fullPhoneNumber"`
	IndividualDetails *IndividualDetails `json:"individualDetails,omitempty"`
	// The Merchant Category Code of the account holder. > If not specified in the request, this will be derived from the platform account (which is configured by Adyen).
	MerchantCategoryCode *string `json:"merchantCategoryCode,omitempty"`
	// A set of key and value pairs for general use by the account holder or merchant. The keys do not have specific names and may be used for storing miscellaneous data as desired. > The values being stored have a maximum length of eighty (80) characters and will be truncated if necessary. > Note that during an update of metadata, the omission of existing key-value pairs will result in the deletion of those key-value pairs.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Each of the card tokens associated with the account holder. > Each array entry should represent one card token. > For comprehensive detail regarding the required `CardToken` fields, please refer to the [KYC Verification documentation](https://docs.adyen.com/platforms/onboarding-and-verification/verification-checks).
	PayoutMethods *[]PayoutMethod `json:"payoutMethods,omitempty"`
	// The URL of the website of the account holder.
	WebAddress string `json:"webAddress"`
}

// NewAccountHolderDetails instantiates a new AccountHolderDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountHolderDetails(email string, fullPhoneNumber string, webAddress string) *AccountHolderDetails {
	this := AccountHolderDetails{}
	this.Email = email
	this.FullPhoneNumber = fullPhoneNumber
	this.WebAddress = webAddress
	return &this
}

// NewAccountHolderDetailsWithDefaults instantiates a new AccountHolderDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountHolderDetailsWithDefaults() *AccountHolderDetails {
	this := AccountHolderDetails{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *AccountHolderDetails) GetAddress() ViasAddress {
	if o == nil || o.Address == nil {
		var ret ViasAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderDetails) GetAddressOk() (*ViasAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *AccountHolderDetails) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given ViasAddress and assigns it to the Address field.
func (o *AccountHolderDetails) SetAddress(v ViasAddress) {
	o.Address = &v
}

// GetBankAccountDetails returns the BankAccountDetails field value if set, zero value otherwise.
func (o *AccountHolderDetails) GetBankAccountDetails() []BankAccountDetail {
	if o == nil || o.BankAccountDetails == nil {
		var ret []BankAccountDetail
		return ret
	}
	return *o.BankAccountDetails
}

// GetBankAccountDetailsOk returns a tuple with the BankAccountDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderDetails) GetBankAccountDetailsOk() (*[]BankAccountDetail, bool) {
	if o == nil || o.BankAccountDetails == nil {
		return nil, false
	}
	return o.BankAccountDetails, true
}

// HasBankAccountDetails returns a boolean if a field has been set.
func (o *AccountHolderDetails) HasBankAccountDetails() bool {
	if o != nil && o.BankAccountDetails != nil {
		return true
	}

	return false
}

// SetBankAccountDetails gets a reference to the given []BankAccountDetail and assigns it to the BankAccountDetails field.
func (o *AccountHolderDetails) SetBankAccountDetails(v []BankAccountDetail) {
	o.BankAccountDetails = &v
}

// GetBankAggregatorDataReference returns the BankAggregatorDataReference field value if set, zero value otherwise.
func (o *AccountHolderDetails) GetBankAggregatorDataReference() string {
	if o == nil || o.BankAggregatorDataReference == nil {
		var ret string
		return ret
	}
	return *o.BankAggregatorDataReference
}

// GetBankAggregatorDataReferenceOk returns a tuple with the BankAggregatorDataReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderDetails) GetBankAggregatorDataReferenceOk() (*string, bool) {
	if o == nil || o.BankAggregatorDataReference == nil {
		return nil, false
	}
	return o.BankAggregatorDataReference, true
}

// HasBankAggregatorDataReference returns a boolean if a field has been set.
func (o *AccountHolderDetails) HasBankAggregatorDataReference() bool {
	if o != nil && o.BankAggregatorDataReference != nil {
		return true
	}

	return false
}

// SetBankAggregatorDataReference gets a reference to the given string and assigns it to the BankAggregatorDataReference field.
func (o *AccountHolderDetails) SetBankAggregatorDataReference(v string) {
	o.BankAggregatorDataReference = &v
}

// GetBusinessDetails returns the BusinessDetails field value if set, zero value otherwise.
func (o *AccountHolderDetails) GetBusinessDetails() BusinessDetails {
	if o == nil || o.BusinessDetails == nil {
		var ret BusinessDetails
		return ret
	}
	return *o.BusinessDetails
}

// GetBusinessDetailsOk returns a tuple with the BusinessDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderDetails) GetBusinessDetailsOk() (*BusinessDetails, bool) {
	if o == nil || o.BusinessDetails == nil {
		return nil, false
	}
	return o.BusinessDetails, true
}

// HasBusinessDetails returns a boolean if a field has been set.
func (o *AccountHolderDetails) HasBusinessDetails() bool {
	if o != nil && o.BusinessDetails != nil {
		return true
	}

	return false
}

// SetBusinessDetails gets a reference to the given BusinessDetails and assigns it to the BusinessDetails field.
func (o *AccountHolderDetails) SetBusinessDetails(v BusinessDetails) {
	o.BusinessDetails = &v
}

// GetEmail returns the Email field value
func (o *AccountHolderDetails) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AccountHolderDetails) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AccountHolderDetails) SetEmail(v string) {
	o.Email = v
}

// GetFullPhoneNumber returns the FullPhoneNumber field value
func (o *AccountHolderDetails) GetFullPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullPhoneNumber
}

// GetFullPhoneNumberOk returns a tuple with the FullPhoneNumber field value
// and a boolean to check if the value has been set.
func (o *AccountHolderDetails) GetFullPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullPhoneNumber, true
}

// SetFullPhoneNumber sets field value
func (o *AccountHolderDetails) SetFullPhoneNumber(v string) {
	o.FullPhoneNumber = v
}

// GetIndividualDetails returns the IndividualDetails field value if set, zero value otherwise.
func (o *AccountHolderDetails) GetIndividualDetails() IndividualDetails {
	if o == nil || o.IndividualDetails == nil {
		var ret IndividualDetails
		return ret
	}
	return *o.IndividualDetails
}

// GetIndividualDetailsOk returns a tuple with the IndividualDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderDetails) GetIndividualDetailsOk() (*IndividualDetails, bool) {
	if o == nil || o.IndividualDetails == nil {
		return nil, false
	}
	return o.IndividualDetails, true
}

// HasIndividualDetails returns a boolean if a field has been set.
func (o *AccountHolderDetails) HasIndividualDetails() bool {
	if o != nil && o.IndividualDetails != nil {
		return true
	}

	return false
}

// SetIndividualDetails gets a reference to the given IndividualDetails and assigns it to the IndividualDetails field.
func (o *AccountHolderDetails) SetIndividualDetails(v IndividualDetails) {
	o.IndividualDetails = &v
}

// GetMerchantCategoryCode returns the MerchantCategoryCode field value if set, zero value otherwise.
func (o *AccountHolderDetails) GetMerchantCategoryCode() string {
	if o == nil || o.MerchantCategoryCode == nil {
		var ret string
		return ret
	}
	return *o.MerchantCategoryCode
}

// GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderDetails) GetMerchantCategoryCodeOk() (*string, bool) {
	if o == nil || o.MerchantCategoryCode == nil {
		return nil, false
	}
	return o.MerchantCategoryCode, true
}

// HasMerchantCategoryCode returns a boolean if a field has been set.
func (o *AccountHolderDetails) HasMerchantCategoryCode() bool {
	if o != nil && o.MerchantCategoryCode != nil {
		return true
	}

	return false
}

// SetMerchantCategoryCode gets a reference to the given string and assigns it to the MerchantCategoryCode field.
func (o *AccountHolderDetails) SetMerchantCategoryCode(v string) {
	o.MerchantCategoryCode = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AccountHolderDetails) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderDetails) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AccountHolderDetails) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *AccountHolderDetails) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetPayoutMethods returns the PayoutMethods field value if set, zero value otherwise.
func (o *AccountHolderDetails) GetPayoutMethods() []PayoutMethod {
	if o == nil || o.PayoutMethods == nil {
		var ret []PayoutMethod
		return ret
	}
	return *o.PayoutMethods
}

// GetPayoutMethodsOk returns a tuple with the PayoutMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderDetails) GetPayoutMethodsOk() (*[]PayoutMethod, bool) {
	if o == nil || o.PayoutMethods == nil {
		return nil, false
	}
	return o.PayoutMethods, true
}

// HasPayoutMethods returns a boolean if a field has been set.
func (o *AccountHolderDetails) HasPayoutMethods() bool {
	if o != nil && o.PayoutMethods != nil {
		return true
	}

	return false
}

// SetPayoutMethods gets a reference to the given []PayoutMethod and assigns it to the PayoutMethods field.
func (o *AccountHolderDetails) SetPayoutMethods(v []PayoutMethod) {
	o.PayoutMethods = &v
}

// GetWebAddress returns the WebAddress field value
func (o *AccountHolderDetails) GetWebAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebAddress
}

// GetWebAddressOk returns a tuple with the WebAddress field value
// and a boolean to check if the value has been set.
func (o *AccountHolderDetails) GetWebAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebAddress, true
}

// SetWebAddress sets field value
func (o *AccountHolderDetails) SetWebAddress(v string) {
	o.WebAddress = v
}

func (o AccountHolderDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.BankAccountDetails != nil {
		toSerialize["bankAccountDetails"] = o.BankAccountDetails
	}
	if o.BankAggregatorDataReference != nil {
		toSerialize["bankAggregatorDataReference"] = o.BankAggregatorDataReference
	}
	if o.BusinessDetails != nil {
		toSerialize["businessDetails"] = o.BusinessDetails
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["fullPhoneNumber"] = o.FullPhoneNumber
	}
	if o.IndividualDetails != nil {
		toSerialize["individualDetails"] = o.IndividualDetails
	}
	if o.MerchantCategoryCode != nil {
		toSerialize["merchantCategoryCode"] = o.MerchantCategoryCode
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.PayoutMethods != nil {
		toSerialize["payoutMethods"] = o.PayoutMethods
	}
	if true {
		toSerialize["webAddress"] = o.WebAddress
	}
	return json.Marshal(toSerialize)
}

type NullableAccountHolderDetails struct {
	value *AccountHolderDetails
	isSet bool
}

func (v NullableAccountHolderDetails) Get() *AccountHolderDetails {
	return v.value
}

func (v *NullableAccountHolderDetails) Set(val *AccountHolderDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountHolderDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountHolderDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountHolderDetails(val *AccountHolderDetails) *NullableAccountHolderDetails {
	return &NullableAccountHolderDetails{value: val, isSet: true}
}

func (v NullableAccountHolderDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountHolderDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
