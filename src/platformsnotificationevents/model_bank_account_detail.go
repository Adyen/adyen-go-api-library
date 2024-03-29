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

// BankAccountDetail struct for BankAccountDetail
type BankAccountDetail struct {
	// The bank account number (without separators). >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	AccountNumber *string `json:"accountNumber,omitempty"`
	// The type of bank account. Only applicable to bank accounts held in the USA. The permitted values are: `checking`, `savings`.  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	AccountType *string `json:"accountType,omitempty"`
	// The name of the bank account.
	BankAccountName *string `json:"bankAccountName,omitempty"`
	// Merchant reference to the bank account.
	BankAccountReference *string `json:"bankAccountReference,omitempty"`
	// The unique identifier (UUID) of the Bank Account. >If, during an account holder create or update request, this field is left blank (but other fields provided), a new Bank Account will be created with a procedurally-generated UUID.  >If, during an account holder create request, a UUID is provided, the creation of the Bank Account will fail while the creation of the account holder will continue.  >If, during an account holder update request, a UUID that is not correlated with an existing Bank Account is provided, the update of the account holder will fail.  >If, during an account holder update request, a UUID that is correlated with an existing Bank Account is provided, the existing Bank Account will be updated.
	BankAccountUUID *string `json:"bankAccountUUID,omitempty"`
	// The bank identifier code. >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	BankBicSwift *string `json:"bankBicSwift,omitempty"`
	// The city in which the bank branch is located.  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	BankCity *string `json:"bankCity,omitempty"`
	// The bank code of the banking institution with which the bank account is registered.  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	BankCode *string `json:"bankCode,omitempty"`
	// The name of the banking institution with which the bank account is held.  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	BankName *string `json:"bankName,omitempty"`
	// The branch code of the branch under which the bank account is registered. The value to be specified in this parameter depends on the country of the bank account: * United States - Routing number * United Kingdom - Sort code * Germany - Bankleitzahl >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	BranchCode *string `json:"branchCode,omitempty"`
	// The check code of the bank account.  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	CheckCode *string `json:"checkCode,omitempty"`
	// The two-letter country code in which the bank account is registered. >The permitted country codes are defined in ISO-3166-1 alpha-2 (e.g. 'NL').  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	CountryCode *string `json:"countryCode,omitempty"`
	// The currency in which the bank account deals. >The permitted currency codes are defined in ISO-4217 (e.g. 'EUR').  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// The international bank account number. >The IBAN standard is defined in ISO-13616.  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	Iban *string `json:"iban,omitempty"`
	// The city of residence of the bank account owner. >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	OwnerCity *string `json:"ownerCity,omitempty"`
	// The country code of the country of residence of the bank account owner. >The permitted country codes are defined in ISO-3166-1 alpha-2 (e.g. 'NL').  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	OwnerCountryCode *string `json:"ownerCountryCode,omitempty"`
	// The date of birth of the bank account owner.
	OwnerDateOfBirth *string `json:"ownerDateOfBirth,omitempty"`
	// The house name or number of the residence of the bank account owner. >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	OwnerHouseNumberOrName *string `json:"ownerHouseNumberOrName,omitempty"`
	// The name of the bank account owner. >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	OwnerName *string `json:"ownerName,omitempty"`
	// The country code of the country of nationality of the bank account owner. >The permitted country codes are defined in ISO-3166-1 alpha-2 (e.g. 'NL').  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	OwnerNationality *string `json:"ownerNationality,omitempty"`
	// The postal code of the residence of the bank account owner. >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	OwnerPostalCode *string `json:"ownerPostalCode,omitempty"`
	// The state of residence of the bank account owner. >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	OwnerState *string `json:"ownerState,omitempty"`
	// The street name of the residence of the bank account owner. >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	OwnerStreet *string `json:"ownerStreet,omitempty"`
	// If set to true, the bank account is a primary account.
	PrimaryAccount *bool `json:"primaryAccount,omitempty"`
	// The tax ID number.  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	TaxId *string `json:"taxId,omitempty"`
	// The URL to be used for bank account verification. This may be generated on bank account creation.  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	UrlForVerification *string `json:"urlForVerification,omitempty"`
}

// NewBankAccountDetail instantiates a new BankAccountDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankAccountDetail() *BankAccountDetail {
	this := BankAccountDetail{}
	return &this
}

// NewBankAccountDetailWithDefaults instantiates a new BankAccountDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankAccountDetailWithDefaults() *BankAccountDetail {
	this := BankAccountDetail{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *BankAccountDetail) GetAccountNumber() string {
	if o == nil || o.AccountNumber == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetAccountNumberOk() (*string, bool) {
	if o == nil || o.AccountNumber == nil {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *BankAccountDetail) HasAccountNumber() bool {
	if o != nil && o.AccountNumber != nil {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *BankAccountDetail) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *BankAccountDetail) GetAccountType() string {
	if o == nil || o.AccountType == nil {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetAccountTypeOk() (*string, bool) {
	if o == nil || o.AccountType == nil {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *BankAccountDetail) HasAccountType() bool {
	if o != nil && o.AccountType != nil {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *BankAccountDetail) SetAccountType(v string) {
	o.AccountType = &v
}

// GetBankAccountName returns the BankAccountName field value if set, zero value otherwise.
func (o *BankAccountDetail) GetBankAccountName() string {
	if o == nil || o.BankAccountName == nil {
		var ret string
		return ret
	}
	return *o.BankAccountName
}

// GetBankAccountNameOk returns a tuple with the BankAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetBankAccountNameOk() (*string, bool) {
	if o == nil || o.BankAccountName == nil {
		return nil, false
	}
	return o.BankAccountName, true
}

// HasBankAccountName returns a boolean if a field has been set.
func (o *BankAccountDetail) HasBankAccountName() bool {
	if o != nil && o.BankAccountName != nil {
		return true
	}

	return false
}

// SetBankAccountName gets a reference to the given string and assigns it to the BankAccountName field.
func (o *BankAccountDetail) SetBankAccountName(v string) {
	o.BankAccountName = &v
}

// GetBankAccountReference returns the BankAccountReference field value if set, zero value otherwise.
func (o *BankAccountDetail) GetBankAccountReference() string {
	if o == nil || o.BankAccountReference == nil {
		var ret string
		return ret
	}
	return *o.BankAccountReference
}

// GetBankAccountReferenceOk returns a tuple with the BankAccountReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetBankAccountReferenceOk() (*string, bool) {
	if o == nil || o.BankAccountReference == nil {
		return nil, false
	}
	return o.BankAccountReference, true
}

// HasBankAccountReference returns a boolean if a field has been set.
func (o *BankAccountDetail) HasBankAccountReference() bool {
	if o != nil && o.BankAccountReference != nil {
		return true
	}

	return false
}

// SetBankAccountReference gets a reference to the given string and assigns it to the BankAccountReference field.
func (o *BankAccountDetail) SetBankAccountReference(v string) {
	o.BankAccountReference = &v
}

// GetBankAccountUUID returns the BankAccountUUID field value if set, zero value otherwise.
func (o *BankAccountDetail) GetBankAccountUUID() string {
	if o == nil || o.BankAccountUUID == nil {
		var ret string
		return ret
	}
	return *o.BankAccountUUID
}

// GetBankAccountUUIDOk returns a tuple with the BankAccountUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetBankAccountUUIDOk() (*string, bool) {
	if o == nil || o.BankAccountUUID == nil {
		return nil, false
	}
	return o.BankAccountUUID, true
}

// HasBankAccountUUID returns a boolean if a field has been set.
func (o *BankAccountDetail) HasBankAccountUUID() bool {
	if o != nil && o.BankAccountUUID != nil {
		return true
	}

	return false
}

// SetBankAccountUUID gets a reference to the given string and assigns it to the BankAccountUUID field.
func (o *BankAccountDetail) SetBankAccountUUID(v string) {
	o.BankAccountUUID = &v
}

// GetBankBicSwift returns the BankBicSwift field value if set, zero value otherwise.
func (o *BankAccountDetail) GetBankBicSwift() string {
	if o == nil || o.BankBicSwift == nil {
		var ret string
		return ret
	}
	return *o.BankBicSwift
}

// GetBankBicSwiftOk returns a tuple with the BankBicSwift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetBankBicSwiftOk() (*string, bool) {
	if o == nil || o.BankBicSwift == nil {
		return nil, false
	}
	return o.BankBicSwift, true
}

// HasBankBicSwift returns a boolean if a field has been set.
func (o *BankAccountDetail) HasBankBicSwift() bool {
	if o != nil && o.BankBicSwift != nil {
		return true
	}

	return false
}

// SetBankBicSwift gets a reference to the given string and assigns it to the BankBicSwift field.
func (o *BankAccountDetail) SetBankBicSwift(v string) {
	o.BankBicSwift = &v
}

// GetBankCity returns the BankCity field value if set, zero value otherwise.
func (o *BankAccountDetail) GetBankCity() string {
	if o == nil || o.BankCity == nil {
		var ret string
		return ret
	}
	return *o.BankCity
}

// GetBankCityOk returns a tuple with the BankCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetBankCityOk() (*string, bool) {
	if o == nil || o.BankCity == nil {
		return nil, false
	}
	return o.BankCity, true
}

// HasBankCity returns a boolean if a field has been set.
func (o *BankAccountDetail) HasBankCity() bool {
	if o != nil && o.BankCity != nil {
		return true
	}

	return false
}

// SetBankCity gets a reference to the given string and assigns it to the BankCity field.
func (o *BankAccountDetail) SetBankCity(v string) {
	o.BankCity = &v
}

// GetBankCode returns the BankCode field value if set, zero value otherwise.
func (o *BankAccountDetail) GetBankCode() string {
	if o == nil || o.BankCode == nil {
		var ret string
		return ret
	}
	return *o.BankCode
}

// GetBankCodeOk returns a tuple with the BankCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetBankCodeOk() (*string, bool) {
	if o == nil || o.BankCode == nil {
		return nil, false
	}
	return o.BankCode, true
}

// HasBankCode returns a boolean if a field has been set.
func (o *BankAccountDetail) HasBankCode() bool {
	if o != nil && o.BankCode != nil {
		return true
	}

	return false
}

// SetBankCode gets a reference to the given string and assigns it to the BankCode field.
func (o *BankAccountDetail) SetBankCode(v string) {
	o.BankCode = &v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *BankAccountDetail) GetBankName() string {
	if o == nil || o.BankName == nil {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetBankNameOk() (*string, bool) {
	if o == nil || o.BankName == nil {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *BankAccountDetail) HasBankName() bool {
	if o != nil && o.BankName != nil {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *BankAccountDetail) SetBankName(v string) {
	o.BankName = &v
}

// GetBranchCode returns the BranchCode field value if set, zero value otherwise.
func (o *BankAccountDetail) GetBranchCode() string {
	if o == nil || o.BranchCode == nil {
		var ret string
		return ret
	}
	return *o.BranchCode
}

// GetBranchCodeOk returns a tuple with the BranchCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetBranchCodeOk() (*string, bool) {
	if o == nil || o.BranchCode == nil {
		return nil, false
	}
	return o.BranchCode, true
}

// HasBranchCode returns a boolean if a field has been set.
func (o *BankAccountDetail) HasBranchCode() bool {
	if o != nil && o.BranchCode != nil {
		return true
	}

	return false
}

// SetBranchCode gets a reference to the given string and assigns it to the BranchCode field.
func (o *BankAccountDetail) SetBranchCode(v string) {
	o.BranchCode = &v
}

// GetCheckCode returns the CheckCode field value if set, zero value otherwise.
func (o *BankAccountDetail) GetCheckCode() string {
	if o == nil || o.CheckCode == nil {
		var ret string
		return ret
	}
	return *o.CheckCode
}

// GetCheckCodeOk returns a tuple with the CheckCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetCheckCodeOk() (*string, bool) {
	if o == nil || o.CheckCode == nil {
		return nil, false
	}
	return o.CheckCode, true
}

// HasCheckCode returns a boolean if a field has been set.
func (o *BankAccountDetail) HasCheckCode() bool {
	if o != nil && o.CheckCode != nil {
		return true
	}

	return false
}

// SetCheckCode gets a reference to the given string and assigns it to the CheckCode field.
func (o *BankAccountDetail) SetCheckCode(v string) {
	o.CheckCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *BankAccountDetail) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *BankAccountDetail) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *BankAccountDetail) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *BankAccountDetail) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *BankAccountDetail) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *BankAccountDetail) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *BankAccountDetail) GetIban() string {
	if o == nil || o.Iban == nil {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetIbanOk() (*string, bool) {
	if o == nil || o.Iban == nil {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *BankAccountDetail) HasIban() bool {
	if o != nil && o.Iban != nil {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *BankAccountDetail) SetIban(v string) {
	o.Iban = &v
}

// GetOwnerCity returns the OwnerCity field value if set, zero value otherwise.
func (o *BankAccountDetail) GetOwnerCity() string {
	if o == nil || o.OwnerCity == nil {
		var ret string
		return ret
	}
	return *o.OwnerCity
}

// GetOwnerCityOk returns a tuple with the OwnerCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetOwnerCityOk() (*string, bool) {
	if o == nil || o.OwnerCity == nil {
		return nil, false
	}
	return o.OwnerCity, true
}

// HasOwnerCity returns a boolean if a field has been set.
func (o *BankAccountDetail) HasOwnerCity() bool {
	if o != nil && o.OwnerCity != nil {
		return true
	}

	return false
}

// SetOwnerCity gets a reference to the given string and assigns it to the OwnerCity field.
func (o *BankAccountDetail) SetOwnerCity(v string) {
	o.OwnerCity = &v
}

// GetOwnerCountryCode returns the OwnerCountryCode field value if set, zero value otherwise.
func (o *BankAccountDetail) GetOwnerCountryCode() string {
	if o == nil || o.OwnerCountryCode == nil {
		var ret string
		return ret
	}
	return *o.OwnerCountryCode
}

// GetOwnerCountryCodeOk returns a tuple with the OwnerCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetOwnerCountryCodeOk() (*string, bool) {
	if o == nil || o.OwnerCountryCode == nil {
		return nil, false
	}
	return o.OwnerCountryCode, true
}

// HasOwnerCountryCode returns a boolean if a field has been set.
func (o *BankAccountDetail) HasOwnerCountryCode() bool {
	if o != nil && o.OwnerCountryCode != nil {
		return true
	}

	return false
}

// SetOwnerCountryCode gets a reference to the given string and assigns it to the OwnerCountryCode field.
func (o *BankAccountDetail) SetOwnerCountryCode(v string) {
	o.OwnerCountryCode = &v
}

// GetOwnerDateOfBirth returns the OwnerDateOfBirth field value if set, zero value otherwise.
func (o *BankAccountDetail) GetOwnerDateOfBirth() string {
	if o == nil || o.OwnerDateOfBirth == nil {
		var ret string
		return ret
	}
	return *o.OwnerDateOfBirth
}

// GetOwnerDateOfBirthOk returns a tuple with the OwnerDateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetOwnerDateOfBirthOk() (*string, bool) {
	if o == nil || o.OwnerDateOfBirth == nil {
		return nil, false
	}
	return o.OwnerDateOfBirth, true
}

// HasOwnerDateOfBirth returns a boolean if a field has been set.
func (o *BankAccountDetail) HasOwnerDateOfBirth() bool {
	if o != nil && o.OwnerDateOfBirth != nil {
		return true
	}

	return false
}

// SetOwnerDateOfBirth gets a reference to the given string and assigns it to the OwnerDateOfBirth field.
func (o *BankAccountDetail) SetOwnerDateOfBirth(v string) {
	o.OwnerDateOfBirth = &v
}

// GetOwnerHouseNumberOrName returns the OwnerHouseNumberOrName field value if set, zero value otherwise.
func (o *BankAccountDetail) GetOwnerHouseNumberOrName() string {
	if o == nil || o.OwnerHouseNumberOrName == nil {
		var ret string
		return ret
	}
	return *o.OwnerHouseNumberOrName
}

// GetOwnerHouseNumberOrNameOk returns a tuple with the OwnerHouseNumberOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetOwnerHouseNumberOrNameOk() (*string, bool) {
	if o == nil || o.OwnerHouseNumberOrName == nil {
		return nil, false
	}
	return o.OwnerHouseNumberOrName, true
}

// HasOwnerHouseNumberOrName returns a boolean if a field has been set.
func (o *BankAccountDetail) HasOwnerHouseNumberOrName() bool {
	if o != nil && o.OwnerHouseNumberOrName != nil {
		return true
	}

	return false
}

// SetOwnerHouseNumberOrName gets a reference to the given string and assigns it to the OwnerHouseNumberOrName field.
func (o *BankAccountDetail) SetOwnerHouseNumberOrName(v string) {
	o.OwnerHouseNumberOrName = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *BankAccountDetail) GetOwnerName() string {
	if o == nil || o.OwnerName == nil {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetOwnerNameOk() (*string, bool) {
	if o == nil || o.OwnerName == nil {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *BankAccountDetail) HasOwnerName() bool {
	if o != nil && o.OwnerName != nil {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *BankAccountDetail) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetOwnerNationality returns the OwnerNationality field value if set, zero value otherwise.
func (o *BankAccountDetail) GetOwnerNationality() string {
	if o == nil || o.OwnerNationality == nil {
		var ret string
		return ret
	}
	return *o.OwnerNationality
}

// GetOwnerNationalityOk returns a tuple with the OwnerNationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetOwnerNationalityOk() (*string, bool) {
	if o == nil || o.OwnerNationality == nil {
		return nil, false
	}
	return o.OwnerNationality, true
}

// HasOwnerNationality returns a boolean if a field has been set.
func (o *BankAccountDetail) HasOwnerNationality() bool {
	if o != nil && o.OwnerNationality != nil {
		return true
	}

	return false
}

// SetOwnerNationality gets a reference to the given string and assigns it to the OwnerNationality field.
func (o *BankAccountDetail) SetOwnerNationality(v string) {
	o.OwnerNationality = &v
}

// GetOwnerPostalCode returns the OwnerPostalCode field value if set, zero value otherwise.
func (o *BankAccountDetail) GetOwnerPostalCode() string {
	if o == nil || o.OwnerPostalCode == nil {
		var ret string
		return ret
	}
	return *o.OwnerPostalCode
}

// GetOwnerPostalCodeOk returns a tuple with the OwnerPostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetOwnerPostalCodeOk() (*string, bool) {
	if o == nil || o.OwnerPostalCode == nil {
		return nil, false
	}
	return o.OwnerPostalCode, true
}

// HasOwnerPostalCode returns a boolean if a field has been set.
func (o *BankAccountDetail) HasOwnerPostalCode() bool {
	if o != nil && o.OwnerPostalCode != nil {
		return true
	}

	return false
}

// SetOwnerPostalCode gets a reference to the given string and assigns it to the OwnerPostalCode field.
func (o *BankAccountDetail) SetOwnerPostalCode(v string) {
	o.OwnerPostalCode = &v
}

// GetOwnerState returns the OwnerState field value if set, zero value otherwise.
func (o *BankAccountDetail) GetOwnerState() string {
	if o == nil || o.OwnerState == nil {
		var ret string
		return ret
	}
	return *o.OwnerState
}

// GetOwnerStateOk returns a tuple with the OwnerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetOwnerStateOk() (*string, bool) {
	if o == nil || o.OwnerState == nil {
		return nil, false
	}
	return o.OwnerState, true
}

// HasOwnerState returns a boolean if a field has been set.
func (o *BankAccountDetail) HasOwnerState() bool {
	if o != nil && o.OwnerState != nil {
		return true
	}

	return false
}

// SetOwnerState gets a reference to the given string and assigns it to the OwnerState field.
func (o *BankAccountDetail) SetOwnerState(v string) {
	o.OwnerState = &v
}

// GetOwnerStreet returns the OwnerStreet field value if set, zero value otherwise.
func (o *BankAccountDetail) GetOwnerStreet() string {
	if o == nil || o.OwnerStreet == nil {
		var ret string
		return ret
	}
	return *o.OwnerStreet
}

// GetOwnerStreetOk returns a tuple with the OwnerStreet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetOwnerStreetOk() (*string, bool) {
	if o == nil || o.OwnerStreet == nil {
		return nil, false
	}
	return o.OwnerStreet, true
}

// HasOwnerStreet returns a boolean if a field has been set.
func (o *BankAccountDetail) HasOwnerStreet() bool {
	if o != nil && o.OwnerStreet != nil {
		return true
	}

	return false
}

// SetOwnerStreet gets a reference to the given string and assigns it to the OwnerStreet field.
func (o *BankAccountDetail) SetOwnerStreet(v string) {
	o.OwnerStreet = &v
}

// GetPrimaryAccount returns the PrimaryAccount field value if set, zero value otherwise.
func (o *BankAccountDetail) GetPrimaryAccount() bool {
	if o == nil || o.PrimaryAccount == nil {
		var ret bool
		return ret
	}
	return *o.PrimaryAccount
}

// GetPrimaryAccountOk returns a tuple with the PrimaryAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetPrimaryAccountOk() (*bool, bool) {
	if o == nil || o.PrimaryAccount == nil {
		return nil, false
	}
	return o.PrimaryAccount, true
}

// HasPrimaryAccount returns a boolean if a field has been set.
func (o *BankAccountDetail) HasPrimaryAccount() bool {
	if o != nil && o.PrimaryAccount != nil {
		return true
	}

	return false
}

// SetPrimaryAccount gets a reference to the given bool and assigns it to the PrimaryAccount field.
func (o *BankAccountDetail) SetPrimaryAccount(v bool) {
	o.PrimaryAccount = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *BankAccountDetail) GetTaxId() string {
	if o == nil || o.TaxId == nil {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetTaxIdOk() (*string, bool) {
	if o == nil || o.TaxId == nil {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *BankAccountDetail) HasTaxId() bool {
	if o != nil && o.TaxId != nil {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *BankAccountDetail) SetTaxId(v string) {
	o.TaxId = &v
}

// GetUrlForVerification returns the UrlForVerification field value if set, zero value otherwise.
func (o *BankAccountDetail) GetUrlForVerification() string {
	if o == nil || o.UrlForVerification == nil {
		var ret string
		return ret
	}
	return *o.UrlForVerification
}

// GetUrlForVerificationOk returns a tuple with the UrlForVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetail) GetUrlForVerificationOk() (*string, bool) {
	if o == nil || o.UrlForVerification == nil {
		return nil, false
	}
	return o.UrlForVerification, true
}

// HasUrlForVerification returns a boolean if a field has been set.
func (o *BankAccountDetail) HasUrlForVerification() bool {
	if o != nil && o.UrlForVerification != nil {
		return true
	}

	return false
}

// SetUrlForVerification gets a reference to the given string and assigns it to the UrlForVerification field.
func (o *BankAccountDetail) SetUrlForVerification(v string) {
	o.UrlForVerification = &v
}

func (o BankAccountDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountNumber != nil {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if o.AccountType != nil {
		toSerialize["accountType"] = o.AccountType
	}
	if o.BankAccountName != nil {
		toSerialize["bankAccountName"] = o.BankAccountName
	}
	if o.BankAccountReference != nil {
		toSerialize["bankAccountReference"] = o.BankAccountReference
	}
	if o.BankAccountUUID != nil {
		toSerialize["bankAccountUUID"] = o.BankAccountUUID
	}
	if o.BankBicSwift != nil {
		toSerialize["bankBicSwift"] = o.BankBicSwift
	}
	if o.BankCity != nil {
		toSerialize["bankCity"] = o.BankCity
	}
	if o.BankCode != nil {
		toSerialize["bankCode"] = o.BankCode
	}
	if o.BankName != nil {
		toSerialize["bankName"] = o.BankName
	}
	if o.BranchCode != nil {
		toSerialize["branchCode"] = o.BranchCode
	}
	if o.CheckCode != nil {
		toSerialize["checkCode"] = o.CheckCode
	}
	if o.CountryCode != nil {
		toSerialize["countryCode"] = o.CountryCode
	}
	if o.CurrencyCode != nil {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if o.Iban != nil {
		toSerialize["iban"] = o.Iban
	}
	if o.OwnerCity != nil {
		toSerialize["ownerCity"] = o.OwnerCity
	}
	if o.OwnerCountryCode != nil {
		toSerialize["ownerCountryCode"] = o.OwnerCountryCode
	}
	if o.OwnerDateOfBirth != nil {
		toSerialize["ownerDateOfBirth"] = o.OwnerDateOfBirth
	}
	if o.OwnerHouseNumberOrName != nil {
		toSerialize["ownerHouseNumberOrName"] = o.OwnerHouseNumberOrName
	}
	if o.OwnerName != nil {
		toSerialize["ownerName"] = o.OwnerName
	}
	if o.OwnerNationality != nil {
		toSerialize["ownerNationality"] = o.OwnerNationality
	}
	if o.OwnerPostalCode != nil {
		toSerialize["ownerPostalCode"] = o.OwnerPostalCode
	}
	if o.OwnerState != nil {
		toSerialize["ownerState"] = o.OwnerState
	}
	if o.OwnerStreet != nil {
		toSerialize["ownerStreet"] = o.OwnerStreet
	}
	if o.PrimaryAccount != nil {
		toSerialize["primaryAccount"] = o.PrimaryAccount
	}
	if o.TaxId != nil {
		toSerialize["taxId"] = o.TaxId
	}
	if o.UrlForVerification != nil {
		toSerialize["urlForVerification"] = o.UrlForVerification
	}
	return json.Marshal(toSerialize)
}

type NullableBankAccountDetail struct {
	value *BankAccountDetail
	isSet bool
}

func (v NullableBankAccountDetail) Get() *BankAccountDetail {
	return v.value
}

func (v *NullableBankAccountDetail) Set(val *BankAccountDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccountDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccountDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccountDetail(val *BankAccountDetail) *NullableBankAccountDetail {
	return &NullableBankAccountDetail{value: val, isSet: true}
}

func (v NullableBankAccountDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankAccountDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
