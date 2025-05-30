/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the ApiCredential type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ApiCredential{}

// ApiCredential struct for ApiCredential
type ApiCredential struct {
	Links *ApiCredentialLinks `json:"_links,omitempty"`
	// Indicates if the API credential is enabled. Must be set to **true** to use the credential in your integration.
	Active bool `json:"active"`
	// List of IP addresses from which your client can make requests.  If the list is empty, we allow requests from any IP. If the list is not empty and we get a request from an IP which is not on the list, you get a security error.
	AllowedIpAddresses []string `json:"allowedIpAddresses"`
	// List containing the [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) linked to the API credential.
	AllowedOrigins []AllowedOrigin `json:"allowedOrigins,omitempty"`
	// Public key used for [client-side authentication](https://docs.adyen.com/development-resources/client-side-authentication). The client key is required for Drop-in and Components integrations.
	ClientKey string `json:"clientKey"`
	// Description of the API credential.
	Description *string `json:"description,omitempty"`
	// Unique identifier of the API credential.
	Id string `json:"id"`
	// List of [roles](https://docs.adyen.com/development-resources/api-credentials#roles-1) for the API credential.
	Roles []string `json:"roles"`
	// The name of the [API credential](https://docs.adyen.com/development-resources/api-credentials), for example **ws@Company.TestCompany**.
	Username string `json:"username"`
}

// NewApiCredential instantiates a new ApiCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCredential(active bool, allowedIpAddresses []string, clientKey string, id string, roles []string, username string) *ApiCredential {
	this := ApiCredential{}
	this.Active = active
	this.AllowedIpAddresses = allowedIpAddresses
	this.ClientKey = clientKey
	this.Id = id
	this.Roles = roles
	this.Username = username
	return &this
}

// NewApiCredentialWithDefaults instantiates a new ApiCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCredentialWithDefaults() *ApiCredential {
	this := ApiCredential{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApiCredential) GetLinks() ApiCredentialLinks {
	if o == nil || common.IsNil(o.Links) {
		var ret ApiCredentialLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCredential) GetLinksOk() (*ApiCredentialLinks, bool) {
	if o == nil || common.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApiCredential) HasLinks() bool {
	if o != nil && !common.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ApiCredentialLinks and assigns it to the Links field.
func (o *ApiCredential) SetLinks(v ApiCredentialLinks) {
	o.Links = &v
}

// GetActive returns the Active field value
func (o *ApiCredential) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *ApiCredential) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *ApiCredential) SetActive(v bool) {
	o.Active = v
}

// GetAllowedIpAddresses returns the AllowedIpAddresses field value
func (o *ApiCredential) GetAllowedIpAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowedIpAddresses
}

// GetAllowedIpAddressesOk returns a tuple with the AllowedIpAddresses field value
// and a boolean to check if the value has been set.
func (o *ApiCredential) GetAllowedIpAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedIpAddresses, true
}

// SetAllowedIpAddresses sets field value
func (o *ApiCredential) SetAllowedIpAddresses(v []string) {
	o.AllowedIpAddresses = v
}

// GetAllowedOrigins returns the AllowedOrigins field value if set, zero value otherwise.
func (o *ApiCredential) GetAllowedOrigins() []AllowedOrigin {
	if o == nil || common.IsNil(o.AllowedOrigins) {
		var ret []AllowedOrigin
		return ret
	}
	return o.AllowedOrigins
}

// GetAllowedOriginsOk returns a tuple with the AllowedOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCredential) GetAllowedOriginsOk() ([]AllowedOrigin, bool) {
	if o == nil || common.IsNil(o.AllowedOrigins) {
		return nil, false
	}
	return o.AllowedOrigins, true
}

// HasAllowedOrigins returns a boolean if a field has been set.
func (o *ApiCredential) HasAllowedOrigins() bool {
	if o != nil && !common.IsNil(o.AllowedOrigins) {
		return true
	}

	return false
}

// SetAllowedOrigins gets a reference to the given []AllowedOrigin and assigns it to the AllowedOrigins field.
func (o *ApiCredential) SetAllowedOrigins(v []AllowedOrigin) {
	o.AllowedOrigins = v
}

// GetClientKey returns the ClientKey field value
func (o *ApiCredential) GetClientKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientKey
}

// GetClientKeyOk returns a tuple with the ClientKey field value
// and a boolean to check if the value has been set.
func (o *ApiCredential) GetClientKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientKey, true
}

// SetClientKey sets field value
func (o *ApiCredential) SetClientKey(v string) {
	o.ClientKey = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiCredential) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCredential) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiCredential) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiCredential) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *ApiCredential) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiCredential) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiCredential) SetId(v string) {
	o.Id = v
}

// GetRoles returns the Roles field value
func (o *ApiCredential) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *ApiCredential) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *ApiCredential) SetRoles(v []string) {
	o.Roles = v
}

// GetUsername returns the Username field value
func (o *ApiCredential) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ApiCredential) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ApiCredential) SetUsername(v string) {
	o.Username = v
}

func (o ApiCredential) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	toSerialize["active"] = o.Active
	toSerialize["allowedIpAddresses"] = o.AllowedIpAddresses
	if !common.IsNil(o.AllowedOrigins) {
		toSerialize["allowedOrigins"] = o.AllowedOrigins
	}
	toSerialize["clientKey"] = o.ClientKey
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	toSerialize["roles"] = o.Roles
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

type NullableApiCredential struct {
	value *ApiCredential
	isSet bool
}

func (v NullableApiCredential) Get() *ApiCredential {
	return v.value
}

func (v *NullableApiCredential) Set(val *ApiCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCredential(val *ApiCredential) *NullableApiCredential {
	return &NullableApiCredential{value: val, isSet: true}
}

func (v NullableApiCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
