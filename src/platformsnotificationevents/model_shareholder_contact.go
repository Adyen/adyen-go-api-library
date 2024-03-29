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

// ShareholderContact struct for ShareholderContact
type ShareholderContact struct {
	Address *ViasAddress `json:"address,omitempty"`
	// The e-mail address of the contact.
	Email *string `json:"email,omitempty"`
	// The phone number of the contact provided as a single string.  It will be handled as a landline phone. **Examples:** \"0031 6 11 22 33 44\", \"+316/1122-3344\", \"(0031) 611223344\"
	FullPhoneNumber *string           `json:"fullPhoneNumber,omitempty"`
	Name            *ViasName         `json:"name,omitempty"`
	PersonalData    *ViasPersonalData `json:"personalData,omitempty"`
	PhoneNumber     *ViasPhoneNumber  `json:"phoneNumber,omitempty"`
	// The unique identifier (UUID) of the Shareholder. >**If, during an Account Holder create or update request, this field is left blank (but other fields provided), a new Shareholder will be created with a procedurally-generated UUID.**  >**If, during an Account Holder create request, a UUID is provided, the creation of the Shareholder will fail while the creation of the Account Holder will continue.**  >**If, during an Account Holder update request, a UUID that is not correlated with an existing Shareholder is provided, the update of the Shareholder will fail.**  >**If, during an Account Holder update request, a UUID that is correlated with an existing Shareholder is provided, the existing Shareholder will be updated.**
	ShareholderCode *string `json:"shareholderCode,omitempty"`
	// Merchant reference to the Shareholder.
	ShareholderReference *string `json:"shareholderReference,omitempty"`
	// The URL of the website of the contact.
	WebAddress *string `json:"webAddress,omitempty"`
}

// NewShareholderContact instantiates a new ShareholderContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShareholderContact() *ShareholderContact {
	this := ShareholderContact{}
	return &this
}

// NewShareholderContactWithDefaults instantiates a new ShareholderContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShareholderContactWithDefaults() *ShareholderContact {
	this := ShareholderContact{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ShareholderContact) GetAddress() ViasAddress {
	if o == nil || o.Address == nil {
		var ret ViasAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderContact) GetAddressOk() (*ViasAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ShareholderContact) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given ViasAddress and assigns it to the Address field.
func (o *ShareholderContact) SetAddress(v ViasAddress) {
	o.Address = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ShareholderContact) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderContact) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ShareholderContact) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ShareholderContact) SetEmail(v string) {
	o.Email = &v
}

// GetFullPhoneNumber returns the FullPhoneNumber field value if set, zero value otherwise.
func (o *ShareholderContact) GetFullPhoneNumber() string {
	if o == nil || o.FullPhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.FullPhoneNumber
}

// GetFullPhoneNumberOk returns a tuple with the FullPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderContact) GetFullPhoneNumberOk() (*string, bool) {
	if o == nil || o.FullPhoneNumber == nil {
		return nil, false
	}
	return o.FullPhoneNumber, true
}

// HasFullPhoneNumber returns a boolean if a field has been set.
func (o *ShareholderContact) HasFullPhoneNumber() bool {
	if o != nil && o.FullPhoneNumber != nil {
		return true
	}

	return false
}

// SetFullPhoneNumber gets a reference to the given string and assigns it to the FullPhoneNumber field.
func (o *ShareholderContact) SetFullPhoneNumber(v string) {
	o.FullPhoneNumber = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ShareholderContact) GetName() ViasName {
	if o == nil || o.Name == nil {
		var ret ViasName
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderContact) GetNameOk() (*ViasName, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ShareholderContact) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given ViasName and assigns it to the Name field.
func (o *ShareholderContact) SetName(v ViasName) {
	o.Name = &v
}

// GetPersonalData returns the PersonalData field value if set, zero value otherwise.
func (o *ShareholderContact) GetPersonalData() ViasPersonalData {
	if o == nil || o.PersonalData == nil {
		var ret ViasPersonalData
		return ret
	}
	return *o.PersonalData
}

// GetPersonalDataOk returns a tuple with the PersonalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderContact) GetPersonalDataOk() (*ViasPersonalData, bool) {
	if o == nil || o.PersonalData == nil {
		return nil, false
	}
	return o.PersonalData, true
}

// HasPersonalData returns a boolean if a field has been set.
func (o *ShareholderContact) HasPersonalData() bool {
	if o != nil && o.PersonalData != nil {
		return true
	}

	return false
}

// SetPersonalData gets a reference to the given ViasPersonalData and assigns it to the PersonalData field.
func (o *ShareholderContact) SetPersonalData(v ViasPersonalData) {
	o.PersonalData = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *ShareholderContact) GetPhoneNumber() ViasPhoneNumber {
	if o == nil || o.PhoneNumber == nil {
		var ret ViasPhoneNumber
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderContact) GetPhoneNumberOk() (*ViasPhoneNumber, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *ShareholderContact) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given ViasPhoneNumber and assigns it to the PhoneNumber field.
func (o *ShareholderContact) SetPhoneNumber(v ViasPhoneNumber) {
	o.PhoneNumber = &v
}

// GetShareholderCode returns the ShareholderCode field value if set, zero value otherwise.
func (o *ShareholderContact) GetShareholderCode() string {
	if o == nil || o.ShareholderCode == nil {
		var ret string
		return ret
	}
	return *o.ShareholderCode
}

// GetShareholderCodeOk returns a tuple with the ShareholderCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderContact) GetShareholderCodeOk() (*string, bool) {
	if o == nil || o.ShareholderCode == nil {
		return nil, false
	}
	return o.ShareholderCode, true
}

// HasShareholderCode returns a boolean if a field has been set.
func (o *ShareholderContact) HasShareholderCode() bool {
	if o != nil && o.ShareholderCode != nil {
		return true
	}

	return false
}

// SetShareholderCode gets a reference to the given string and assigns it to the ShareholderCode field.
func (o *ShareholderContact) SetShareholderCode(v string) {
	o.ShareholderCode = &v
}

// GetShareholderReference returns the ShareholderReference field value if set, zero value otherwise.
func (o *ShareholderContact) GetShareholderReference() string {
	if o == nil || o.ShareholderReference == nil {
		var ret string
		return ret
	}
	return *o.ShareholderReference
}

// GetShareholderReferenceOk returns a tuple with the ShareholderReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderContact) GetShareholderReferenceOk() (*string, bool) {
	if o == nil || o.ShareholderReference == nil {
		return nil, false
	}
	return o.ShareholderReference, true
}

// HasShareholderReference returns a boolean if a field has been set.
func (o *ShareholderContact) HasShareholderReference() bool {
	if o != nil && o.ShareholderReference != nil {
		return true
	}

	return false
}

// SetShareholderReference gets a reference to the given string and assigns it to the ShareholderReference field.
func (o *ShareholderContact) SetShareholderReference(v string) {
	o.ShareholderReference = &v
}

// GetWebAddress returns the WebAddress field value if set, zero value otherwise.
func (o *ShareholderContact) GetWebAddress() string {
	if o == nil || o.WebAddress == nil {
		var ret string
		return ret
	}
	return *o.WebAddress
}

// GetWebAddressOk returns a tuple with the WebAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderContact) GetWebAddressOk() (*string, bool) {
	if o == nil || o.WebAddress == nil {
		return nil, false
	}
	return o.WebAddress, true
}

// HasWebAddress returns a boolean if a field has been set.
func (o *ShareholderContact) HasWebAddress() bool {
	if o != nil && o.WebAddress != nil {
		return true
	}

	return false
}

// SetWebAddress gets a reference to the given string and assigns it to the WebAddress field.
func (o *ShareholderContact) SetWebAddress(v string) {
	o.WebAddress = &v
}

func (o ShareholderContact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FullPhoneNumber != nil {
		toSerialize["fullPhoneNumber"] = o.FullPhoneNumber
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PersonalData != nil {
		toSerialize["personalData"] = o.PersonalData
	}
	if o.PhoneNumber != nil {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if o.ShareholderCode != nil {
		toSerialize["shareholderCode"] = o.ShareholderCode
	}
	if o.ShareholderReference != nil {
		toSerialize["shareholderReference"] = o.ShareholderReference
	}
	if o.WebAddress != nil {
		toSerialize["webAddress"] = o.WebAddress
	}
	return json.Marshal(toSerialize)
}

type NullableShareholderContact struct {
	value *ShareholderContact
	isSet bool
}

func (v NullableShareholderContact) Get() *ShareholderContact {
	return v.value
}

func (v *NullableShareholderContact) Set(val *ShareholderContact) {
	v.value = val
	v.isSet = true
}

func (v NullableShareholderContact) IsSet() bool {
	return v.isSet
}

func (v *NullableShareholderContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShareholderContact(val *ShareholderContact) *NullableShareholderContact {
	return &NullableShareholderContact{value: val, isSet: true}
}

func (v NullableShareholderContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShareholderContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
