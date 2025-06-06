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

// checks if the AndroidApp type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AndroidApp{}

// AndroidApp struct for AndroidApp
type AndroidApp struct {
	// The description that was provided when uploading the app. The description is not shown on the terminal.
	Description *string `json:"description,omitempty"`
	// The error code of the Android app with the `status` of either **error** or **invalid**.
	// Deprecated since Management API v3
	// Use `errors` instead.
	ErrorCode *string `json:"errorCode,omitempty"`
	// The list of errors of the Android app.
	Errors []AndroidAppError `json:"errors,omitempty"`
	// The unique identifier of the app.
	Id string `json:"id"`
	// The app name that is shown on the terminal.
	Label *string `json:"label,omitempty"`
	// The package name that uniquely identifies the Android app.
	PackageName *string `json:"packageName,omitempty"`
	// The status of the app. Possible values:  * `processing`: the app is being signed and converted to a format that the terminal can handle. * `error`: something went wrong. Check that the app matches the [requirements](https://docs.adyen.com/point-of-sale/android-terminals/app-requirements). * `invalid`: there is something wrong with the APK file of the app. * `ready`: the app has been signed and converted. * `archived`: the app is no longer available.
	Status string `json:"status"`
	// The version number of the app.
	VersionCode *int32 `json:"versionCode,omitempty"`
	// The app version number that is shown on the terminal.
	VersionName *string `json:"versionName,omitempty"`
}

// NewAndroidApp instantiates a new AndroidApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAndroidApp(id string, status string) *AndroidApp {
	this := AndroidApp{}
	this.Id = id
	this.Status = status
	return &this
}

// NewAndroidAppWithDefaults instantiates a new AndroidApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAndroidAppWithDefaults() *AndroidApp {
	this := AndroidApp{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AndroidApp) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidApp) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AndroidApp) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AndroidApp) SetDescription(v string) {
	o.Description = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
// Deprecated since Management API v3
// Use `errors` instead.
func (o *AndroidApp) GetErrorCode() string {
	if o == nil || common.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated since Management API v3
// Use `errors` instead.
func (o *AndroidApp) GetErrorCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *AndroidApp) HasErrorCode() bool {
	if o != nil && !common.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
// Deprecated since Management API v3
// Use `errors` instead.
func (o *AndroidApp) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AndroidApp) GetErrors() []AndroidAppError {
	if o == nil || common.IsNil(o.Errors) {
		var ret []AndroidAppError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidApp) GetErrorsOk() ([]AndroidAppError, bool) {
	if o == nil || common.IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AndroidApp) HasErrors() bool {
	if o != nil && !common.IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []AndroidAppError and assigns it to the Errors field.
func (o *AndroidApp) SetErrors(v []AndroidAppError) {
	o.Errors = v
}

// GetId returns the Id field value
func (o *AndroidApp) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AndroidApp) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AndroidApp) SetId(v string) {
	o.Id = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *AndroidApp) GetLabel() string {
	if o == nil || common.IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidApp) GetLabelOk() (*string, bool) {
	if o == nil || common.IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *AndroidApp) HasLabel() bool {
	if o != nil && !common.IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *AndroidApp) SetLabel(v string) {
	o.Label = &v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *AndroidApp) GetPackageName() string {
	if o == nil || common.IsNil(o.PackageName) {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidApp) GetPackageNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.PackageName) {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *AndroidApp) HasPackageName() bool {
	if o != nil && !common.IsNil(o.PackageName) {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *AndroidApp) SetPackageName(v string) {
	o.PackageName = &v
}

// GetStatus returns the Status field value
func (o *AndroidApp) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AndroidApp) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AndroidApp) SetStatus(v string) {
	o.Status = v
}

// GetVersionCode returns the VersionCode field value if set, zero value otherwise.
func (o *AndroidApp) GetVersionCode() int32 {
	if o == nil || common.IsNil(o.VersionCode) {
		var ret int32
		return ret
	}
	return *o.VersionCode
}

// GetVersionCodeOk returns a tuple with the VersionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidApp) GetVersionCodeOk() (*int32, bool) {
	if o == nil || common.IsNil(o.VersionCode) {
		return nil, false
	}
	return o.VersionCode, true
}

// HasVersionCode returns a boolean if a field has been set.
func (o *AndroidApp) HasVersionCode() bool {
	if o != nil && !common.IsNil(o.VersionCode) {
		return true
	}

	return false
}

// SetVersionCode gets a reference to the given int32 and assigns it to the VersionCode field.
func (o *AndroidApp) SetVersionCode(v int32) {
	o.VersionCode = &v
}

// GetVersionName returns the VersionName field value if set, zero value otherwise.
func (o *AndroidApp) GetVersionName() string {
	if o == nil || common.IsNil(o.VersionName) {
		var ret string
		return ret
	}
	return *o.VersionName
}

// GetVersionNameOk returns a tuple with the VersionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidApp) GetVersionNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.VersionName) {
		return nil, false
	}
	return o.VersionName, true
}

// HasVersionName returns a boolean if a field has been set.
func (o *AndroidApp) HasVersionName() bool {
	if o != nil && !common.IsNil(o.VersionName) {
		return true
	}

	return false
}

// SetVersionName gets a reference to the given string and assigns it to the VersionName field.
func (o *AndroidApp) SetVersionName(v string) {
	o.VersionName = &v
}

func (o AndroidApp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AndroidApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !common.IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	toSerialize["id"] = o.Id
	if !common.IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !common.IsNil(o.PackageName) {
		toSerialize["packageName"] = o.PackageName
	}
	toSerialize["status"] = o.Status
	if !common.IsNil(o.VersionCode) {
		toSerialize["versionCode"] = o.VersionCode
	}
	if !common.IsNil(o.VersionName) {
		toSerialize["versionName"] = o.VersionName
	}
	return toSerialize, nil
}

type NullableAndroidApp struct {
	value *AndroidApp
	isSet bool
}

func (v NullableAndroidApp) Get() *AndroidApp {
	return v.value
}

func (v *NullableAndroidApp) Set(val *AndroidApp) {
	v.value = val
	v.isSet = true
}

func (v NullableAndroidApp) IsSet() bool {
	return v.isSet
}

func (v *NullableAndroidApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAndroidApp(val *AndroidApp) *NullableAndroidApp {
	return &NullableAndroidApp{value: val, isSet: true}
}

func (v NullableAndroidApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAndroidApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AndroidApp) isValidStatus() bool {
	var allowedEnumValues = []string{"archived", "error", "invalid", "processing", "ready"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
