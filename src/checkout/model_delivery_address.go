/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the DeliveryAddress type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DeliveryAddress{}

// DeliveryAddress struct for DeliveryAddress
type DeliveryAddress struct {
	// The name of the city. Maximum length: 3000 characters.
	City string `json:"city"`
	// The two-character ISO-3166-1 alpha-2 country code. For example, **US**. > If you don't know the country or are not collecting the country from the shopper, provide `country` as `ZZ`.
	Country   string  `json:"country"`
	FirstName *string `json:"firstName,omitempty"`
	// The number or name of the house. Maximum length: 3000 characters.
	HouseNumberOrName string  `json:"houseNumberOrName"`
	LastName          *string `json:"lastName,omitempty"`
	// A maximum of five digits for an address in the US, or a maximum of ten characters for an address in all other countries.
	PostalCode string `json:"postalCode"`
	// The two-character ISO 3166-2 state or province code. For example, **CA** in the US or **ON** in Canada. > Required for the US and Canada.
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
	// The name of the street. Maximum length: 3000 characters. > The house number should not be included in this field; it should be separately provided via `houseNumberOrName`.
	Street string `json:"street"`
}

// NewDeliveryAddress instantiates a new DeliveryAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryAddress(city string, country string, houseNumberOrName string, postalCode string, street string) *DeliveryAddress {
	this := DeliveryAddress{}
	this.City = city
	this.Country = country
	this.HouseNumberOrName = houseNumberOrName
	this.PostalCode = postalCode
	this.Street = street
	return &this
}

// NewDeliveryAddressWithDefaults instantiates a new DeliveryAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryAddressWithDefaults() *DeliveryAddress {
	this := DeliveryAddress{}
	return &this
}

// GetCity returns the City field value
func (o *DeliveryAddress) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *DeliveryAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *DeliveryAddress) SetCity(v string) {
	o.City = v
}

// GetCountry returns the Country field value
func (o *DeliveryAddress) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *DeliveryAddress) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *DeliveryAddress) SetCountry(v string) {
	o.Country = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *DeliveryAddress) GetFirstName() string {
	if o == nil || common.IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryAddress) GetFirstNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *DeliveryAddress) HasFirstName() bool {
	if o != nil && !common.IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *DeliveryAddress) SetFirstName(v string) {
	o.FirstName = &v
}

// GetHouseNumberOrName returns the HouseNumberOrName field value
func (o *DeliveryAddress) GetHouseNumberOrName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HouseNumberOrName
}

// GetHouseNumberOrNameOk returns a tuple with the HouseNumberOrName field value
// and a boolean to check if the value has been set.
func (o *DeliveryAddress) GetHouseNumberOrNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HouseNumberOrName, true
}

// SetHouseNumberOrName sets field value
func (o *DeliveryAddress) SetHouseNumberOrName(v string) {
	o.HouseNumberOrName = v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *DeliveryAddress) GetLastName() string {
	if o == nil || common.IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryAddress) GetLastNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *DeliveryAddress) HasLastName() bool {
	if o != nil && !common.IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *DeliveryAddress) SetLastName(v string) {
	o.LastName = &v
}

// GetPostalCode returns the PostalCode field value
func (o *DeliveryAddress) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *DeliveryAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *DeliveryAddress) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetStateOrProvince returns the StateOrProvince field value if set, zero value otherwise.
func (o *DeliveryAddress) GetStateOrProvince() string {
	if o == nil || common.IsNil(o.StateOrProvince) {
		var ret string
		return ret
	}
	return *o.StateOrProvince
}

// GetStateOrProvinceOk returns a tuple with the StateOrProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryAddress) GetStateOrProvinceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StateOrProvince) {
		return nil, false
	}
	return o.StateOrProvince, true
}

// HasStateOrProvince returns a boolean if a field has been set.
func (o *DeliveryAddress) HasStateOrProvince() bool {
	if o != nil && !common.IsNil(o.StateOrProvince) {
		return true
	}

	return false
}

// SetStateOrProvince gets a reference to the given string and assigns it to the StateOrProvince field.
func (o *DeliveryAddress) SetStateOrProvince(v string) {
	o.StateOrProvince = &v
}

// GetStreet returns the Street field value
func (o *DeliveryAddress) GetStreet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Street
}

// GetStreetOk returns a tuple with the Street field value
// and a boolean to check if the value has been set.
func (o *DeliveryAddress) GetStreetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Street, true
}

// SetStreet sets field value
func (o *DeliveryAddress) SetStreet(v string) {
	o.Street = v
}

func (o DeliveryAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["city"] = o.City
	toSerialize["country"] = o.Country
	if !common.IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	toSerialize["houseNumberOrName"] = o.HouseNumberOrName
	if !common.IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	toSerialize["postalCode"] = o.PostalCode
	if !common.IsNil(o.StateOrProvince) {
		toSerialize["stateOrProvince"] = o.StateOrProvince
	}
	toSerialize["street"] = o.Street
	return toSerialize, nil
}

type NullableDeliveryAddress struct {
	value *DeliveryAddress
	isSet bool
}

func (v NullableDeliveryAddress) Get() *DeliveryAddress {
	return v.value
}

func (v *NullableDeliveryAddress) Set(val *DeliveryAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryAddress(val *DeliveryAddress) *NullableDeliveryAddress {
	return &NullableDeliveryAddress{value: val, isSet: true}
}

func (v NullableDeliveryAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}