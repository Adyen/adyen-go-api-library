# UpdateCompanyWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptsExpiredCertificate** | Pointer to **bool** | Indicates if expired SSL certificates are accepted. Default value: **false**. | [optional] 
**AcceptsSelfSignedCertificate** | Pointer to **bool** | Indicates if self-signed SSL certificates are accepted. Default value: **false**. | [optional] 
**AcceptsUntrustedRootCertificate** | Pointer to **bool** | Indicates if untrusted SSL certificates are accepted. Default value: **false**. | [optional] 
**Active** | Pointer to **bool** | Indicates if the webhook configuration is active. The field must be **true** for us to send webhooks about events related an account. | [optional] 
**AdditionalSettings** | Pointer to [**AdditionalSettings**](AdditionalSettings.md) |  | [optional] 
**CommunicationFormat** | Pointer to **string** | Format or protocol for receiving webhooks. Possible values: * **soap** * **http** * **json**  | [optional] 
**Description** | Pointer to **string** | Your description for this webhook configuration. | [optional] 
**FilterMerchantAccountType** | Pointer to **string** | Shows how merchant accounts are filtered when configuring the webhook. Possible values: * **includeAccounts**: The webhook is configured for the merchant accounts listed in &#x60;filterMerchantAccounts&#x60;. * **excludeAccounts**: The webhook is not configured for the merchant accounts listed in &#x60;filterMerchantAccounts&#x60;. * **allAccounts**: Includes all merchant accounts, and does not require specifying &#x60;filterMerchantAccounts&#x60;. | [optional] 
**FilterMerchantAccounts** | Pointer to **[]string** | A list of merchant account names that are included or excluded from receiving the webhook. Inclusion or exclusion is based on the value defined for &#x60;filterMerchantAccountType&#x60;.  Required if &#x60;filterMerchantAccountType&#x60; is either: * **includeAccounts** * **excludeAccounts**  Not needed for &#x60;filterMerchantAccountType&#x60;: **allAccounts**. | [optional] 
**NetworkType** | Pointer to **string** | Network type for Terminal API notification webhooks. Possible values: * **public** * **local**  Default Value: **public**. | [optional] 
**Password** | Pointer to **string** | Password to access the webhook URL. | [optional] 
**PopulateSoapActionHeader** | Pointer to **bool** | Indicates if the SOAP action header needs to be populated. Default value: **false**.  Only applies if &#x60;communicationFormat&#x60;: **soap**. | [optional] 
**SslVersion** | Pointer to **string** | SSL version to access the public webhook URL specified in the &#x60;url&#x60; field. Possible values: * **TLSv1.2** * **HTTP** - Only allowed on Test environment.  If not specified, the webhook will use &#x60;sslVersion&#x60;: **TLSv1.2**. | [optional] 
**Url** | **string** | Public URL where webhooks will be sent, for example **https://www.domain.com/webhook-endpoint**. | 
**Username** | Pointer to **string** | Username to access the webhook URL. | [optional] 

## Methods

### NewUpdateCompanyWebhookRequest

`func NewUpdateCompanyWebhookRequest(url string, ) *UpdateCompanyWebhookRequest`

NewUpdateCompanyWebhookRequest instantiates a new UpdateCompanyWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCompanyWebhookRequestWithDefaults

`func NewUpdateCompanyWebhookRequestWithDefaults() *UpdateCompanyWebhookRequest`

NewUpdateCompanyWebhookRequestWithDefaults instantiates a new UpdateCompanyWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptsExpiredCertificate

`func (o *UpdateCompanyWebhookRequest) GetAcceptsExpiredCertificate() bool`

GetAcceptsExpiredCertificate returns the AcceptsExpiredCertificate field if non-nil, zero value otherwise.

### GetAcceptsExpiredCertificateOk

`func (o *UpdateCompanyWebhookRequest) GetAcceptsExpiredCertificateOk() (*bool, bool)`

GetAcceptsExpiredCertificateOk returns a tuple with the AcceptsExpiredCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptsExpiredCertificate

`func (o *UpdateCompanyWebhookRequest) SetAcceptsExpiredCertificate(v bool)`

SetAcceptsExpiredCertificate sets AcceptsExpiredCertificate field to given value.

### HasAcceptsExpiredCertificate

`func (o *UpdateCompanyWebhookRequest) HasAcceptsExpiredCertificate() bool`

HasAcceptsExpiredCertificate returns a boolean if a field has been set.

### GetAcceptsSelfSignedCertificate

`func (o *UpdateCompanyWebhookRequest) GetAcceptsSelfSignedCertificate() bool`

GetAcceptsSelfSignedCertificate returns the AcceptsSelfSignedCertificate field if non-nil, zero value otherwise.

### GetAcceptsSelfSignedCertificateOk

`func (o *UpdateCompanyWebhookRequest) GetAcceptsSelfSignedCertificateOk() (*bool, bool)`

GetAcceptsSelfSignedCertificateOk returns a tuple with the AcceptsSelfSignedCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptsSelfSignedCertificate

`func (o *UpdateCompanyWebhookRequest) SetAcceptsSelfSignedCertificate(v bool)`

SetAcceptsSelfSignedCertificate sets AcceptsSelfSignedCertificate field to given value.

### HasAcceptsSelfSignedCertificate

`func (o *UpdateCompanyWebhookRequest) HasAcceptsSelfSignedCertificate() bool`

HasAcceptsSelfSignedCertificate returns a boolean if a field has been set.

### GetAcceptsUntrustedRootCertificate

`func (o *UpdateCompanyWebhookRequest) GetAcceptsUntrustedRootCertificate() bool`

GetAcceptsUntrustedRootCertificate returns the AcceptsUntrustedRootCertificate field if non-nil, zero value otherwise.

### GetAcceptsUntrustedRootCertificateOk

`func (o *UpdateCompanyWebhookRequest) GetAcceptsUntrustedRootCertificateOk() (*bool, bool)`

GetAcceptsUntrustedRootCertificateOk returns a tuple with the AcceptsUntrustedRootCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptsUntrustedRootCertificate

`func (o *UpdateCompanyWebhookRequest) SetAcceptsUntrustedRootCertificate(v bool)`

SetAcceptsUntrustedRootCertificate sets AcceptsUntrustedRootCertificate field to given value.

### HasAcceptsUntrustedRootCertificate

`func (o *UpdateCompanyWebhookRequest) HasAcceptsUntrustedRootCertificate() bool`

HasAcceptsUntrustedRootCertificate returns a boolean if a field has been set.

### GetActive

`func (o *UpdateCompanyWebhookRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateCompanyWebhookRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateCompanyWebhookRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateCompanyWebhookRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAdditionalSettings

`func (o *UpdateCompanyWebhookRequest) GetAdditionalSettings() AdditionalSettings`

GetAdditionalSettings returns the AdditionalSettings field if non-nil, zero value otherwise.

### GetAdditionalSettingsOk

`func (o *UpdateCompanyWebhookRequest) GetAdditionalSettingsOk() (*AdditionalSettings, bool)`

GetAdditionalSettingsOk returns a tuple with the AdditionalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSettings

`func (o *UpdateCompanyWebhookRequest) SetAdditionalSettings(v AdditionalSettings)`

SetAdditionalSettings sets AdditionalSettings field to given value.

### HasAdditionalSettings

`func (o *UpdateCompanyWebhookRequest) HasAdditionalSettings() bool`

HasAdditionalSettings returns a boolean if a field has been set.

### GetCommunicationFormat

`func (o *UpdateCompanyWebhookRequest) GetCommunicationFormat() string`

GetCommunicationFormat returns the CommunicationFormat field if non-nil, zero value otherwise.

### GetCommunicationFormatOk

`func (o *UpdateCompanyWebhookRequest) GetCommunicationFormatOk() (*string, bool)`

GetCommunicationFormatOk returns a tuple with the CommunicationFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationFormat

`func (o *UpdateCompanyWebhookRequest) SetCommunicationFormat(v string)`

SetCommunicationFormat sets CommunicationFormat field to given value.

### HasCommunicationFormat

`func (o *UpdateCompanyWebhookRequest) HasCommunicationFormat() bool`

HasCommunicationFormat returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateCompanyWebhookRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCompanyWebhookRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCompanyWebhookRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCompanyWebhookRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilterMerchantAccountType

`func (o *UpdateCompanyWebhookRequest) GetFilterMerchantAccountType() string`

GetFilterMerchantAccountType returns the FilterMerchantAccountType field if non-nil, zero value otherwise.

### GetFilterMerchantAccountTypeOk

`func (o *UpdateCompanyWebhookRequest) GetFilterMerchantAccountTypeOk() (*string, bool)`

GetFilterMerchantAccountTypeOk returns a tuple with the FilterMerchantAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterMerchantAccountType

`func (o *UpdateCompanyWebhookRequest) SetFilterMerchantAccountType(v string)`

SetFilterMerchantAccountType sets FilterMerchantAccountType field to given value.

### HasFilterMerchantAccountType

`func (o *UpdateCompanyWebhookRequest) HasFilterMerchantAccountType() bool`

HasFilterMerchantAccountType returns a boolean if a field has been set.

### GetFilterMerchantAccounts

`func (o *UpdateCompanyWebhookRequest) GetFilterMerchantAccounts() []string`

GetFilterMerchantAccounts returns the FilterMerchantAccounts field if non-nil, zero value otherwise.

### GetFilterMerchantAccountsOk

`func (o *UpdateCompanyWebhookRequest) GetFilterMerchantAccountsOk() (*[]string, bool)`

GetFilterMerchantAccountsOk returns a tuple with the FilterMerchantAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterMerchantAccounts

`func (o *UpdateCompanyWebhookRequest) SetFilterMerchantAccounts(v []string)`

SetFilterMerchantAccounts sets FilterMerchantAccounts field to given value.

### HasFilterMerchantAccounts

`func (o *UpdateCompanyWebhookRequest) HasFilterMerchantAccounts() bool`

HasFilterMerchantAccounts returns a boolean if a field has been set.

### GetNetworkType

`func (o *UpdateCompanyWebhookRequest) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *UpdateCompanyWebhookRequest) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *UpdateCompanyWebhookRequest) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *UpdateCompanyWebhookRequest) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateCompanyWebhookRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateCompanyWebhookRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateCompanyWebhookRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateCompanyWebhookRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPopulateSoapActionHeader

`func (o *UpdateCompanyWebhookRequest) GetPopulateSoapActionHeader() bool`

GetPopulateSoapActionHeader returns the PopulateSoapActionHeader field if non-nil, zero value otherwise.

### GetPopulateSoapActionHeaderOk

`func (o *UpdateCompanyWebhookRequest) GetPopulateSoapActionHeaderOk() (*bool, bool)`

GetPopulateSoapActionHeaderOk returns a tuple with the PopulateSoapActionHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulateSoapActionHeader

`func (o *UpdateCompanyWebhookRequest) SetPopulateSoapActionHeader(v bool)`

SetPopulateSoapActionHeader sets PopulateSoapActionHeader field to given value.

### HasPopulateSoapActionHeader

`func (o *UpdateCompanyWebhookRequest) HasPopulateSoapActionHeader() bool`

HasPopulateSoapActionHeader returns a boolean if a field has been set.

### GetSslVersion

`func (o *UpdateCompanyWebhookRequest) GetSslVersion() string`

GetSslVersion returns the SslVersion field if non-nil, zero value otherwise.

### GetSslVersionOk

`func (o *UpdateCompanyWebhookRequest) GetSslVersionOk() (*string, bool)`

GetSslVersionOk returns a tuple with the SslVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVersion

`func (o *UpdateCompanyWebhookRequest) SetSslVersion(v string)`

SetSslVersion sets SslVersion field to given value.

### HasSslVersion

`func (o *UpdateCompanyWebhookRequest) HasSslVersion() bool`

HasSslVersion returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateCompanyWebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateCompanyWebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateCompanyWebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsername

`func (o *UpdateCompanyWebhookRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateCompanyWebhookRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateCompanyWebhookRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateCompanyWebhookRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


