# UpdateMerchantWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptsExpiredCertificate** | Pointer to **bool** | Indicates if expired SSL certificates are accepted. Default value: **false**. | [optional] 
**AcceptsSelfSignedCertificate** | Pointer to **bool** | Indicates if self-signed SSL certificates are accepted. Default value: **false**. | [optional] 
**AcceptsUntrustedRootCertificate** | Pointer to **bool** | Indicates if untrusted SSL certificates are accepted. Default value: **false**. | [optional] 
**Active** | **bool** | Indicates if the webhook configuration is active. The field must be **true** for us to send webhooks about events related an account. | 
**AdditionalSettings** | Pointer to [**AdditionalSettings**](AdditionalSettings.md) |  | [optional] 
**CommunicationFormat** | **string** | Format or protocol for receiving webhooks. Possible values: * **soap** * **http** * **json**  | 
**Description** | Pointer to **string** | Your description for this webhook configuration. | [optional] 
**NetworkType** | Pointer to **string** | Network type for Terminal API notification webhooks. Possible values: * **public** * **local**  Default Value: **public**. | [optional] 
**Password** | Pointer to **string** | Password to access the webhook URL. | [optional] 
**PopulateSoapActionHeader** | Pointer to **bool** | Indicates if the SOAP action header needs to be populated. Default value: **false**.  Only applies if &#x60;communicationFormat&#x60;: **soap**. | [optional] 
**SslVersion** | Pointer to **string** | SSL version to access the public webhook URL specified in the &#x60;url&#x60; field. Possible values: * **TLSv1.2** * **HTTP** - Only allowed on Test environment.  If not specified, the webhook will use &#x60;sslVersion&#x60;: **TLSv1.2**. | [optional] 
**Url** | **string** | Public URL where webhooks will be sent, for example **https://www.domain.com/webhook-endpoint**. | 
**Username** | Pointer to **string** | Username to access the webhook URL. | [optional] 

## Methods

### NewUpdateMerchantWebhookRequest

`func NewUpdateMerchantWebhookRequest(active bool, communicationFormat string, url string, ) *UpdateMerchantWebhookRequest`

NewUpdateMerchantWebhookRequest instantiates a new UpdateMerchantWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMerchantWebhookRequestWithDefaults

`func NewUpdateMerchantWebhookRequestWithDefaults() *UpdateMerchantWebhookRequest`

NewUpdateMerchantWebhookRequestWithDefaults instantiates a new UpdateMerchantWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptsExpiredCertificate

`func (o *UpdateMerchantWebhookRequest) GetAcceptsExpiredCertificate() bool`

GetAcceptsExpiredCertificate returns the AcceptsExpiredCertificate field if non-nil, zero value otherwise.

### GetAcceptsExpiredCertificateOk

`func (o *UpdateMerchantWebhookRequest) GetAcceptsExpiredCertificateOk() (*bool, bool)`

GetAcceptsExpiredCertificateOk returns a tuple with the AcceptsExpiredCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptsExpiredCertificate

`func (o *UpdateMerchantWebhookRequest) SetAcceptsExpiredCertificate(v bool)`

SetAcceptsExpiredCertificate sets AcceptsExpiredCertificate field to given value.

### HasAcceptsExpiredCertificate

`func (o *UpdateMerchantWebhookRequest) HasAcceptsExpiredCertificate() bool`

HasAcceptsExpiredCertificate returns a boolean if a field has been set.

### GetAcceptsSelfSignedCertificate

`func (o *UpdateMerchantWebhookRequest) GetAcceptsSelfSignedCertificate() bool`

GetAcceptsSelfSignedCertificate returns the AcceptsSelfSignedCertificate field if non-nil, zero value otherwise.

### GetAcceptsSelfSignedCertificateOk

`func (o *UpdateMerchantWebhookRequest) GetAcceptsSelfSignedCertificateOk() (*bool, bool)`

GetAcceptsSelfSignedCertificateOk returns a tuple with the AcceptsSelfSignedCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptsSelfSignedCertificate

`func (o *UpdateMerchantWebhookRequest) SetAcceptsSelfSignedCertificate(v bool)`

SetAcceptsSelfSignedCertificate sets AcceptsSelfSignedCertificate field to given value.

### HasAcceptsSelfSignedCertificate

`func (o *UpdateMerchantWebhookRequest) HasAcceptsSelfSignedCertificate() bool`

HasAcceptsSelfSignedCertificate returns a boolean if a field has been set.

### GetAcceptsUntrustedRootCertificate

`func (o *UpdateMerchantWebhookRequest) GetAcceptsUntrustedRootCertificate() bool`

GetAcceptsUntrustedRootCertificate returns the AcceptsUntrustedRootCertificate field if non-nil, zero value otherwise.

### GetAcceptsUntrustedRootCertificateOk

`func (o *UpdateMerchantWebhookRequest) GetAcceptsUntrustedRootCertificateOk() (*bool, bool)`

GetAcceptsUntrustedRootCertificateOk returns a tuple with the AcceptsUntrustedRootCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptsUntrustedRootCertificate

`func (o *UpdateMerchantWebhookRequest) SetAcceptsUntrustedRootCertificate(v bool)`

SetAcceptsUntrustedRootCertificate sets AcceptsUntrustedRootCertificate field to given value.

### HasAcceptsUntrustedRootCertificate

`func (o *UpdateMerchantWebhookRequest) HasAcceptsUntrustedRootCertificate() bool`

HasAcceptsUntrustedRootCertificate returns a boolean if a field has been set.

### GetActive

`func (o *UpdateMerchantWebhookRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateMerchantWebhookRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateMerchantWebhookRequest) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAdditionalSettings

`func (o *UpdateMerchantWebhookRequest) GetAdditionalSettings() AdditionalSettings`

GetAdditionalSettings returns the AdditionalSettings field if non-nil, zero value otherwise.

### GetAdditionalSettingsOk

`func (o *UpdateMerchantWebhookRequest) GetAdditionalSettingsOk() (*AdditionalSettings, bool)`

GetAdditionalSettingsOk returns a tuple with the AdditionalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSettings

`func (o *UpdateMerchantWebhookRequest) SetAdditionalSettings(v AdditionalSettings)`

SetAdditionalSettings sets AdditionalSettings field to given value.

### HasAdditionalSettings

`func (o *UpdateMerchantWebhookRequest) HasAdditionalSettings() bool`

HasAdditionalSettings returns a boolean if a field has been set.

### GetCommunicationFormat

`func (o *UpdateMerchantWebhookRequest) GetCommunicationFormat() string`

GetCommunicationFormat returns the CommunicationFormat field if non-nil, zero value otherwise.

### GetCommunicationFormatOk

`func (o *UpdateMerchantWebhookRequest) GetCommunicationFormatOk() (*string, bool)`

GetCommunicationFormatOk returns a tuple with the CommunicationFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationFormat

`func (o *UpdateMerchantWebhookRequest) SetCommunicationFormat(v string)`

SetCommunicationFormat sets CommunicationFormat field to given value.


### GetDescription

`func (o *UpdateMerchantWebhookRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateMerchantWebhookRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateMerchantWebhookRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateMerchantWebhookRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNetworkType

`func (o *UpdateMerchantWebhookRequest) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *UpdateMerchantWebhookRequest) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *UpdateMerchantWebhookRequest) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *UpdateMerchantWebhookRequest) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateMerchantWebhookRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateMerchantWebhookRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateMerchantWebhookRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateMerchantWebhookRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPopulateSoapActionHeader

`func (o *UpdateMerchantWebhookRequest) GetPopulateSoapActionHeader() bool`

GetPopulateSoapActionHeader returns the PopulateSoapActionHeader field if non-nil, zero value otherwise.

### GetPopulateSoapActionHeaderOk

`func (o *UpdateMerchantWebhookRequest) GetPopulateSoapActionHeaderOk() (*bool, bool)`

GetPopulateSoapActionHeaderOk returns a tuple with the PopulateSoapActionHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulateSoapActionHeader

`func (o *UpdateMerchantWebhookRequest) SetPopulateSoapActionHeader(v bool)`

SetPopulateSoapActionHeader sets PopulateSoapActionHeader field to given value.

### HasPopulateSoapActionHeader

`func (o *UpdateMerchantWebhookRequest) HasPopulateSoapActionHeader() bool`

HasPopulateSoapActionHeader returns a boolean if a field has been set.

### GetSslVersion

`func (o *UpdateMerchantWebhookRequest) GetSslVersion() string`

GetSslVersion returns the SslVersion field if non-nil, zero value otherwise.

### GetSslVersionOk

`func (o *UpdateMerchantWebhookRequest) GetSslVersionOk() (*string, bool)`

GetSslVersionOk returns a tuple with the SslVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVersion

`func (o *UpdateMerchantWebhookRequest) SetSslVersion(v string)`

SetSslVersion sets SslVersion field to given value.

### HasSslVersion

`func (o *UpdateMerchantWebhookRequest) HasSslVersion() bool`

HasSslVersion returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateMerchantWebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateMerchantWebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateMerchantWebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsername

`func (o *UpdateMerchantWebhookRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateMerchantWebhookRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateMerchantWebhookRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateMerchantWebhookRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


