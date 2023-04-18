# ThreeDS2CardRangeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcsInfoInd** | Pointer to **[]string** | Provides additional information to the 3DS Server. Possible values: - 01 (Authentication is available at ACS) - 02 (Attempts supported by ACS or DS) - 03 (Decoupled authentication supported) - 04 (Whitelisting supported) | [optional] 
**BrandCode** | Pointer to **string** | Card brand. | [optional] 
**EndRange** | Pointer to **string** | BIN end range. | [optional] 
**StartRange** | Pointer to **string** | BIN start range. | [optional] 
**ThreeDS2Versions** | Pointer to **[]string** | Supported 3D Secure protocol versions | [optional] 
**ThreeDSMethodURL** | Pointer to **string** | In a 3D Secure 2 browser-based flow, this is the URL where you should send the device fingerprint to. | [optional] 

## Methods

### NewThreeDS2CardRangeDetail

`func NewThreeDS2CardRangeDetail() *ThreeDS2CardRangeDetail`

NewThreeDS2CardRangeDetail instantiates a new ThreeDS2CardRangeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDS2CardRangeDetailWithDefaults

`func NewThreeDS2CardRangeDetailWithDefaults() *ThreeDS2CardRangeDetail`

NewThreeDS2CardRangeDetailWithDefaults instantiates a new ThreeDS2CardRangeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcsInfoInd

`func (o *ThreeDS2CardRangeDetail) GetAcsInfoInd() []string`

GetAcsInfoInd returns the AcsInfoInd field if non-nil, zero value otherwise.

### GetAcsInfoIndOk

`func (o *ThreeDS2CardRangeDetail) GetAcsInfoIndOk() (*[]string, bool)`

GetAcsInfoIndOk returns a tuple with the AcsInfoInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsInfoInd

`func (o *ThreeDS2CardRangeDetail) SetAcsInfoInd(v []string)`

SetAcsInfoInd sets AcsInfoInd field to given value.

### HasAcsInfoInd

`func (o *ThreeDS2CardRangeDetail) HasAcsInfoInd() bool`

HasAcsInfoInd returns a boolean if a field has been set.

### GetBrandCode

`func (o *ThreeDS2CardRangeDetail) GetBrandCode() string`

GetBrandCode returns the BrandCode field if non-nil, zero value otherwise.

### GetBrandCodeOk

`func (o *ThreeDS2CardRangeDetail) GetBrandCodeOk() (*string, bool)`

GetBrandCodeOk returns a tuple with the BrandCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandCode

`func (o *ThreeDS2CardRangeDetail) SetBrandCode(v string)`

SetBrandCode sets BrandCode field to given value.

### HasBrandCode

`func (o *ThreeDS2CardRangeDetail) HasBrandCode() bool`

HasBrandCode returns a boolean if a field has been set.

### GetEndRange

`func (o *ThreeDS2CardRangeDetail) GetEndRange() string`

GetEndRange returns the EndRange field if non-nil, zero value otherwise.

### GetEndRangeOk

`func (o *ThreeDS2CardRangeDetail) GetEndRangeOk() (*string, bool)`

GetEndRangeOk returns a tuple with the EndRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndRange

`func (o *ThreeDS2CardRangeDetail) SetEndRange(v string)`

SetEndRange sets EndRange field to given value.

### HasEndRange

`func (o *ThreeDS2CardRangeDetail) HasEndRange() bool`

HasEndRange returns a boolean if a field has been set.

### GetStartRange

`func (o *ThreeDS2CardRangeDetail) GetStartRange() string`

GetStartRange returns the StartRange field if non-nil, zero value otherwise.

### GetStartRangeOk

`func (o *ThreeDS2CardRangeDetail) GetStartRangeOk() (*string, bool)`

GetStartRangeOk returns a tuple with the StartRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartRange

`func (o *ThreeDS2CardRangeDetail) SetStartRange(v string)`

SetStartRange sets StartRange field to given value.

### HasStartRange

`func (o *ThreeDS2CardRangeDetail) HasStartRange() bool`

HasStartRange returns a boolean if a field has been set.

### GetThreeDS2Versions

`func (o *ThreeDS2CardRangeDetail) GetThreeDS2Versions() []string`

GetThreeDS2Versions returns the ThreeDS2Versions field if non-nil, zero value otherwise.

### GetThreeDS2VersionsOk

`func (o *ThreeDS2CardRangeDetail) GetThreeDS2VersionsOk() (*[]string, bool)`

GetThreeDS2VersionsOk returns a tuple with the ThreeDS2Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDS2Versions

`func (o *ThreeDS2CardRangeDetail) SetThreeDS2Versions(v []string)`

SetThreeDS2Versions sets ThreeDS2Versions field to given value.

### HasThreeDS2Versions

`func (o *ThreeDS2CardRangeDetail) HasThreeDS2Versions() bool`

HasThreeDS2Versions returns a boolean if a field has been set.

### GetThreeDSMethodURL

`func (o *ThreeDS2CardRangeDetail) GetThreeDSMethodURL() string`

GetThreeDSMethodURL returns the ThreeDSMethodURL field if non-nil, zero value otherwise.

### GetThreeDSMethodURLOk

`func (o *ThreeDS2CardRangeDetail) GetThreeDSMethodURLOk() (*string, bool)`

GetThreeDSMethodURLOk returns a tuple with the ThreeDSMethodURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSMethodURL

`func (o *ThreeDS2CardRangeDetail) SetThreeDSMethodURL(v string)`

SetThreeDSMethodURL sets ThreeDSMethodURL field to given value.

### HasThreeDSMethodURL

`func (o *ThreeDS2CardRangeDetail) HasThreeDSMethodURL() bool`

HasThreeDSMethodURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


