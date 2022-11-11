# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**WebhookLinks**](WebhookLinks.md) |  | [optional] 
**AcceptsExpiredCertificate** | Pointer to **bool** | Indicates if expired SSL certificates are accepted. Default value: **false**. | [optional] 
**AcceptsSelfSignedCertificate** | Pointer to **bool** | Indicates if self-signed SSL certificates are accepted. Default value: **false**. | [optional] 
**AcceptsUntrustedRootCertificate** | Pointer to **bool** | Indicates if untrusted SSL certificates are accepted. Default value: **false**. | [optional] 
**AccountReference** | Pointer to **string** | Reference to the account the webook is set on. | [optional] 
**Active** | **bool** | Indicates if the webhook configuration is active. The field must be **true** for you to receive webhooks about events related an account. | 
**AdditionalSettings** | Pointer to [**AdditionalSettingsResponse**](AdditionalSettingsResponse.md) |  | [optional] 
**CertificateAlias** | Pointer to **string** | The alias of our SSL certificate. When you receive a notification from us, the alias from the HMAC signature will match this alias. | [optional] 
**CommunicationFormat** | **string** | Format or protocol for receiving webhooks. Possible values: * **soap** * **http** * **json**  | 
**Description** | Pointer to **string** | Your description for this webhook configuration. | [optional] 
**FilterMerchantAccountType** | Pointer to **string** | Shows how merchant accounts are included in company-level webhooks. Possible values: * **includeAccounts** * **excludeAccounts** * **allAccounts**: Includes all merchant accounts, and does not require specifying &#x60;filterMerchantAccounts&#x60;. | [optional] 
**FilterMerchantAccounts** | Pointer to **[]string** | A list of merchant account names that are included or excluded from receiving the webhook. Inclusion or exclusion is based on the value defined for &#x60;filterMerchantAccountType&#x60;.  Required if &#x60;filterMerchantAccountType&#x60; is either: * **includeAccounts** * **excludeAccounts**  Not needed for &#x60;filterMerchantAccountType&#x60;: **allAccounts**. | [optional] 
**HasError** | Pointer to **bool** | Indicates if the webhook configuration has errors that need troubleshooting. If the value is **true**, troubleshoot the configuration using the [testing endpoint](https://docs.adyen.com/api-explorer/#/ManagementService/v1/post/companies/{companyId}/webhooks/{webhookid}/test). | [optional] 
**HasPassword** | Pointer to **bool** | Indicates if the webhook is password protected. | [optional] 
**HmacKeyCheckValue** | Pointer to **string** | The [checksum](https://en.wikipedia.org/wiki/Key_checksum_value) of the HMAC key generated for this webhook. You can use this value to uniquely identify the HMAC key configured for this webhook. | [optional] 
**Id** | Pointer to **string** | Unique identifier for this webhook. | [optional] 
**NetworkType** | Pointer to **string** | Network type for Terminal API details webhooks. | [optional] 
**PopulateSoapActionHeader** | Pointer to **bool** | Indicates if the SOAP action header needs to be populated. Default value: **false**.  Only applies if &#x60;communicationFormat&#x60;: **soap**. | [optional] 
**SslVersion** | Pointer to **string** | SSL version to access the public webhook URL specified in the &#x60;url&#x60; field. Possible values: * **TLSv1.2** * **HTTP** - Only allowed on Test environment.  If not specified, the webhook will use &#x60;sslVersion&#x60;: **TLSv1.2**. | [optional] 
**Type** | **string** | The type of webhook. Possible values are:  - **standard** - **account-settings-notification** - **banktransfer-notification** - **boletobancario-notification** - **directdebit-notification** - **pending-notification** - **ideal-notification** - **ideal-pending-notification** - **report-notification** - **terminal-api-notification**  Find out more about [standard notification webhooks](https://docs.adyen.com/development-resources/webhooks/understand-notifications#event-codes) and [other types of notifications](https://docs.adyen.com/development-resources/webhooks/understand-notifications#other-notifications). | 
**Url** | **string** | Public URL where webhooks will be sent, for example **https://www.domain.com/webhook-endpoint**. | 
**Username** | Pointer to **string** | Username to access the webhook URL. | [optional] 

## Methods

### NewWebhook

`func NewWebhook(active bool, communicationFormat string, type_ string, url string, ) *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Webhook) GetLinks() WebhookLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Webhook) GetLinksOk() (*WebhookLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Webhook) SetLinks(v WebhookLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Webhook) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAcceptsExpiredCertificate

`func (o *Webhook) GetAcceptsExpiredCertificate() bool`

GetAcceptsExpiredCertificate returns the AcceptsExpiredCertificate field if non-nil, zero value otherwise.

### GetAcceptsExpiredCertificateOk

`func (o *Webhook) GetAcceptsExpiredCertificateOk() (*bool, bool)`

GetAcceptsExpiredCertificateOk returns a tuple with the AcceptsExpiredCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptsExpiredCertificate

`func (o *Webhook) SetAcceptsExpiredCertificate(v bool)`

SetAcceptsExpiredCertificate sets AcceptsExpiredCertificate field to given value.

### HasAcceptsExpiredCertificate

`func (o *Webhook) HasAcceptsExpiredCertificate() bool`

HasAcceptsExpiredCertificate returns a boolean if a field has been set.

### GetAcceptsSelfSignedCertificate

`func (o *Webhook) GetAcceptsSelfSignedCertificate() bool`

GetAcceptsSelfSignedCertificate returns the AcceptsSelfSignedCertificate field if non-nil, zero value otherwise.

### GetAcceptsSelfSignedCertificateOk

`func (o *Webhook) GetAcceptsSelfSignedCertificateOk() (*bool, bool)`

GetAcceptsSelfSignedCertificateOk returns a tuple with the AcceptsSelfSignedCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptsSelfSignedCertificate

`func (o *Webhook) SetAcceptsSelfSignedCertificate(v bool)`

SetAcceptsSelfSignedCertificate sets AcceptsSelfSignedCertificate field to given value.

### HasAcceptsSelfSignedCertificate

`func (o *Webhook) HasAcceptsSelfSignedCertificate() bool`

HasAcceptsSelfSignedCertificate returns a boolean if a field has been set.

### GetAcceptsUntrustedRootCertificate

`func (o *Webhook) GetAcceptsUntrustedRootCertificate() bool`

GetAcceptsUntrustedRootCertificate returns the AcceptsUntrustedRootCertificate field if non-nil, zero value otherwise.

### GetAcceptsUntrustedRootCertificateOk

`func (o *Webhook) GetAcceptsUntrustedRootCertificateOk() (*bool, bool)`

GetAcceptsUntrustedRootCertificateOk returns a tuple with the AcceptsUntrustedRootCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptsUntrustedRootCertificate

`func (o *Webhook) SetAcceptsUntrustedRootCertificate(v bool)`

SetAcceptsUntrustedRootCertificate sets AcceptsUntrustedRootCertificate field to given value.

### HasAcceptsUntrustedRootCertificate

`func (o *Webhook) HasAcceptsUntrustedRootCertificate() bool`

HasAcceptsUntrustedRootCertificate returns a boolean if a field has been set.

### GetAccountReference

`func (o *Webhook) GetAccountReference() string`

GetAccountReference returns the AccountReference field if non-nil, zero value otherwise.

### GetAccountReferenceOk

`func (o *Webhook) GetAccountReferenceOk() (*string, bool)`

GetAccountReferenceOk returns a tuple with the AccountReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountReference

`func (o *Webhook) SetAccountReference(v string)`

SetAccountReference sets AccountReference field to given value.

### HasAccountReference

`func (o *Webhook) HasAccountReference() bool`

HasAccountReference returns a boolean if a field has been set.

### GetActive

`func (o *Webhook) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Webhook) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Webhook) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAdditionalSettings

`func (o *Webhook) GetAdditionalSettings() AdditionalSettingsResponse`

GetAdditionalSettings returns the AdditionalSettings field if non-nil, zero value otherwise.

### GetAdditionalSettingsOk

`func (o *Webhook) GetAdditionalSettingsOk() (*AdditionalSettingsResponse, bool)`

GetAdditionalSettingsOk returns a tuple with the AdditionalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSettings

`func (o *Webhook) SetAdditionalSettings(v AdditionalSettingsResponse)`

SetAdditionalSettings sets AdditionalSettings field to given value.

### HasAdditionalSettings

`func (o *Webhook) HasAdditionalSettings() bool`

HasAdditionalSettings returns a boolean if a field has been set.

### GetCertificateAlias

`func (o *Webhook) GetCertificateAlias() string`

GetCertificateAlias returns the CertificateAlias field if non-nil, zero value otherwise.

### GetCertificateAliasOk

`func (o *Webhook) GetCertificateAliasOk() (*string, bool)`

GetCertificateAliasOk returns a tuple with the CertificateAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAlias

`func (o *Webhook) SetCertificateAlias(v string)`

SetCertificateAlias sets CertificateAlias field to given value.

### HasCertificateAlias

`func (o *Webhook) HasCertificateAlias() bool`

HasCertificateAlias returns a boolean if a field has been set.

### GetCommunicationFormat

`func (o *Webhook) GetCommunicationFormat() string`

GetCommunicationFormat returns the CommunicationFormat field if non-nil, zero value otherwise.

### GetCommunicationFormatOk

`func (o *Webhook) GetCommunicationFormatOk() (*string, bool)`

GetCommunicationFormatOk returns a tuple with the CommunicationFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationFormat

`func (o *Webhook) SetCommunicationFormat(v string)`

SetCommunicationFormat sets CommunicationFormat field to given value.


### GetDescription

`func (o *Webhook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Webhook) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Webhook) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Webhook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilterMerchantAccountType

`func (o *Webhook) GetFilterMerchantAccountType() string`

GetFilterMerchantAccountType returns the FilterMerchantAccountType field if non-nil, zero value otherwise.

### GetFilterMerchantAccountTypeOk

`func (o *Webhook) GetFilterMerchantAccountTypeOk() (*string, bool)`

GetFilterMerchantAccountTypeOk returns a tuple with the FilterMerchantAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterMerchantAccountType

`func (o *Webhook) SetFilterMerchantAccountType(v string)`

SetFilterMerchantAccountType sets FilterMerchantAccountType field to given value.

### HasFilterMerchantAccountType

`func (o *Webhook) HasFilterMerchantAccountType() bool`

HasFilterMerchantAccountType returns a boolean if a field has been set.

### GetFilterMerchantAccounts

`func (o *Webhook) GetFilterMerchantAccounts() []string`

GetFilterMerchantAccounts returns the FilterMerchantAccounts field if non-nil, zero value otherwise.

### GetFilterMerchantAccountsOk

`func (o *Webhook) GetFilterMerchantAccountsOk() (*[]string, bool)`

GetFilterMerchantAccountsOk returns a tuple with the FilterMerchantAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterMerchantAccounts

`func (o *Webhook) SetFilterMerchantAccounts(v []string)`

SetFilterMerchantAccounts sets FilterMerchantAccounts field to given value.

### HasFilterMerchantAccounts

`func (o *Webhook) HasFilterMerchantAccounts() bool`

HasFilterMerchantAccounts returns a boolean if a field has been set.

### GetHasError

`func (o *Webhook) GetHasError() bool`

GetHasError returns the HasError field if non-nil, zero value otherwise.

### GetHasErrorOk

`func (o *Webhook) GetHasErrorOk() (*bool, bool)`

GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasError

`func (o *Webhook) SetHasError(v bool)`

SetHasError sets HasError field to given value.

### HasHasError

`func (o *Webhook) HasHasError() bool`

HasHasError returns a boolean if a field has been set.

### GetHasPassword

`func (o *Webhook) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *Webhook) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *Webhook) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *Webhook) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetHmacKeyCheckValue

`func (o *Webhook) GetHmacKeyCheckValue() string`

GetHmacKeyCheckValue returns the HmacKeyCheckValue field if non-nil, zero value otherwise.

### GetHmacKeyCheckValueOk

`func (o *Webhook) GetHmacKeyCheckValueOk() (*string, bool)`

GetHmacKeyCheckValueOk returns a tuple with the HmacKeyCheckValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacKeyCheckValue

`func (o *Webhook) SetHmacKeyCheckValue(v string)`

SetHmacKeyCheckValue sets HmacKeyCheckValue field to given value.

### HasHmacKeyCheckValue

`func (o *Webhook) HasHmacKeyCheckValue() bool`

HasHmacKeyCheckValue returns a boolean if a field has been set.

### GetId

`func (o *Webhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Webhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Webhook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Webhook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkType

`func (o *Webhook) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *Webhook) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *Webhook) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *Webhook) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetPopulateSoapActionHeader

`func (o *Webhook) GetPopulateSoapActionHeader() bool`

GetPopulateSoapActionHeader returns the PopulateSoapActionHeader field if non-nil, zero value otherwise.

### GetPopulateSoapActionHeaderOk

`func (o *Webhook) GetPopulateSoapActionHeaderOk() (*bool, bool)`

GetPopulateSoapActionHeaderOk returns a tuple with the PopulateSoapActionHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulateSoapActionHeader

`func (o *Webhook) SetPopulateSoapActionHeader(v bool)`

SetPopulateSoapActionHeader sets PopulateSoapActionHeader field to given value.

### HasPopulateSoapActionHeader

`func (o *Webhook) HasPopulateSoapActionHeader() bool`

HasPopulateSoapActionHeader returns a boolean if a field has been set.

### GetSslVersion

`func (o *Webhook) GetSslVersion() string`

GetSslVersion returns the SslVersion field if non-nil, zero value otherwise.

### GetSslVersionOk

`func (o *Webhook) GetSslVersionOk() (*string, bool)`

GetSslVersionOk returns a tuple with the SslVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVersion

`func (o *Webhook) SetSslVersion(v string)`

SetSslVersion sets SslVersion field to given value.

### HasSslVersion

`func (o *Webhook) HasSslVersion() bool`

HasSslVersion returns a boolean if a field has been set.

### GetType

`func (o *Webhook) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Webhook) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Webhook) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsername

`func (o *Webhook) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Webhook) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Webhook) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Webhook) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


