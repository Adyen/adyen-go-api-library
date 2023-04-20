# ThreeDSAvailabilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinDetails** | Pointer to [**BinDetail**](BinDetail.md) |  | [optional] 
**DsPublicKeys** | Pointer to [**[]DSPublicKeyDetail**](DSPublicKeyDetail.md) | List of Directory Server (DS) public keys. | [optional] 
**ThreeDS1Supported** | Pointer to **bool** | Indicator if 3D Secure 1 is supported. | [optional] 
**ThreeDS2CardRangeDetails** | Pointer to [**[]ThreeDS2CardRangeDetail**](ThreeDS2CardRangeDetail.md) | List of brand and card range pairs. | [optional] 
**ThreeDS2supported** | Pointer to **bool** | Indicator if 3D Secure 2 is supported. | [optional] 

## Methods

### NewThreeDSAvailabilityResponse

`func NewThreeDSAvailabilityResponse() *ThreeDSAvailabilityResponse`

NewThreeDSAvailabilityResponse instantiates a new ThreeDSAvailabilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSAvailabilityResponseWithDefaults

`func NewThreeDSAvailabilityResponseWithDefaults() *ThreeDSAvailabilityResponse`

NewThreeDSAvailabilityResponseWithDefaults instantiates a new ThreeDSAvailabilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinDetails

`func (o *ThreeDSAvailabilityResponse) GetBinDetails() BinDetail`

GetBinDetails returns the BinDetails field if non-nil, zero value otherwise.

### GetBinDetailsOk

`func (o *ThreeDSAvailabilityResponse) GetBinDetailsOk() (*BinDetail, bool)`

GetBinDetailsOk returns a tuple with the BinDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinDetails

`func (o *ThreeDSAvailabilityResponse) SetBinDetails(v BinDetail)`

SetBinDetails sets BinDetails field to given value.

### HasBinDetails

`func (o *ThreeDSAvailabilityResponse) HasBinDetails() bool`

HasBinDetails returns a boolean if a field has been set.

### GetDsPublicKeys

`func (o *ThreeDSAvailabilityResponse) GetDsPublicKeys() []DSPublicKeyDetail`

GetDsPublicKeys returns the DsPublicKeys field if non-nil, zero value otherwise.

### GetDsPublicKeysOk

`func (o *ThreeDSAvailabilityResponse) GetDsPublicKeysOk() (*[]DSPublicKeyDetail, bool)`

GetDsPublicKeysOk returns a tuple with the DsPublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsPublicKeys

`func (o *ThreeDSAvailabilityResponse) SetDsPublicKeys(v []DSPublicKeyDetail)`

SetDsPublicKeys sets DsPublicKeys field to given value.

### HasDsPublicKeys

`func (o *ThreeDSAvailabilityResponse) HasDsPublicKeys() bool`

HasDsPublicKeys returns a boolean if a field has been set.

### GetThreeDS1Supported

`func (o *ThreeDSAvailabilityResponse) GetThreeDS1Supported() bool`

GetThreeDS1Supported returns the ThreeDS1Supported field if non-nil, zero value otherwise.

### GetThreeDS1SupportedOk

`func (o *ThreeDSAvailabilityResponse) GetThreeDS1SupportedOk() (*bool, bool)`

GetThreeDS1SupportedOk returns a tuple with the ThreeDS1Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDS1Supported

`func (o *ThreeDSAvailabilityResponse) SetThreeDS1Supported(v bool)`

SetThreeDS1Supported sets ThreeDS1Supported field to given value.

### HasThreeDS1Supported

`func (o *ThreeDSAvailabilityResponse) HasThreeDS1Supported() bool`

HasThreeDS1Supported returns a boolean if a field has been set.

### GetThreeDS2CardRangeDetails

`func (o *ThreeDSAvailabilityResponse) GetThreeDS2CardRangeDetails() []ThreeDS2CardRangeDetail`

GetThreeDS2CardRangeDetails returns the ThreeDS2CardRangeDetails field if non-nil, zero value otherwise.

### GetThreeDS2CardRangeDetailsOk

`func (o *ThreeDSAvailabilityResponse) GetThreeDS2CardRangeDetailsOk() (*[]ThreeDS2CardRangeDetail, bool)`

GetThreeDS2CardRangeDetailsOk returns a tuple with the ThreeDS2CardRangeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDS2CardRangeDetails

`func (o *ThreeDSAvailabilityResponse) SetThreeDS2CardRangeDetails(v []ThreeDS2CardRangeDetail)`

SetThreeDS2CardRangeDetails sets ThreeDS2CardRangeDetails field to given value.

### HasThreeDS2CardRangeDetails

`func (o *ThreeDSAvailabilityResponse) HasThreeDS2CardRangeDetails() bool`

HasThreeDS2CardRangeDetails returns a boolean if a field has been set.

### GetThreeDS2supported

`func (o *ThreeDSAvailabilityResponse) GetThreeDS2supported() bool`

GetThreeDS2supported returns the ThreeDS2supported field if non-nil, zero value otherwise.

### GetThreeDS2supportedOk

`func (o *ThreeDSAvailabilityResponse) GetThreeDS2supportedOk() (*bool, bool)`

GetThreeDS2supportedOk returns a tuple with the ThreeDS2supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDS2supported

`func (o *ThreeDSAvailabilityResponse) SetThreeDS2supported(v bool)`

SetThreeDS2supported sets ThreeDS2supported field to given value.

### HasThreeDS2supported

`func (o *ThreeDSAvailabilityResponse) HasThreeDS2supported() bool`

HasThreeDS2supported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


