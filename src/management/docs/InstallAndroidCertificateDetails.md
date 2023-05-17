# InstallAndroidCertificateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateId** | Pointer to **string** | The unique identifier of the certificate to be installed. | [optional] 
**Type** | Pointer to **string** | Type of terminal action: Install an Android certificate. | [optional] [default to "InstallAndroidCertificate"]

## Methods

### NewInstallAndroidCertificateDetails

`func NewInstallAndroidCertificateDetails() *InstallAndroidCertificateDetails`

NewInstallAndroidCertificateDetails instantiates a new InstallAndroidCertificateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallAndroidCertificateDetailsWithDefaults

`func NewInstallAndroidCertificateDetailsWithDefaults() *InstallAndroidCertificateDetails`

NewInstallAndroidCertificateDetailsWithDefaults instantiates a new InstallAndroidCertificateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateId

`func (o *InstallAndroidCertificateDetails) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *InstallAndroidCertificateDetails) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *InstallAndroidCertificateDetails) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *InstallAndroidCertificateDetails) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetType

`func (o *InstallAndroidCertificateDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstallAndroidCertificateDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstallAndroidCertificateDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InstallAndroidCertificateDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


