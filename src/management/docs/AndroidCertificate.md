# AndroidCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description that was provided when uploading the certificate. | [optional] 
**Extension** | Pointer to **string** | The file format of the certificate, as indicated by the file extension. For example, **.cert** or **.pem**. | [optional] 
**Id** | **string** | The unique identifier of the certificate. | 
**Name** | Pointer to **string** | The file name of the certificate. For example, **mycert**. | [optional] 
**NotAfter** | Pointer to **time.Time** | The date when the certificate stops to be valid. | [optional] 
**NotBefore** | Pointer to **time.Time** | The date when the certificate starts to be valid. | [optional] 
**Status** | Pointer to **string** | The status of the certificate. | [optional] 

## Methods

### NewAndroidCertificate

`func NewAndroidCertificate(id string, ) *AndroidCertificate`

NewAndroidCertificate instantiates a new AndroidCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAndroidCertificateWithDefaults

`func NewAndroidCertificateWithDefaults() *AndroidCertificate`

NewAndroidCertificateWithDefaults instantiates a new AndroidCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AndroidCertificate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AndroidCertificate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AndroidCertificate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AndroidCertificate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtension

`func (o *AndroidCertificate) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *AndroidCertificate) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *AndroidCertificate) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *AndroidCertificate) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetId

`func (o *AndroidCertificate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AndroidCertificate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AndroidCertificate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AndroidCertificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AndroidCertificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AndroidCertificate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AndroidCertificate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotAfter

`func (o *AndroidCertificate) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *AndroidCertificate) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *AndroidCertificate) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *AndroidCertificate) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *AndroidCertificate) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *AndroidCertificate) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *AndroidCertificate) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *AndroidCertificate) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetStatus

`func (o *AndroidCertificate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AndroidCertificate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AndroidCertificate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AndroidCertificate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


