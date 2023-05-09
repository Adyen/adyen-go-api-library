# CardConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activation** | Pointer to **string** | Overrides the activation label design ID defined in the &#x60;configurationProfileId&#x60;. The activation label is attached to the card and contains the activation instructions. | [optional] 
**ActivationUrl** | Pointer to **string** | Your app&#39;s URL, if you want to activate cards through your app. For example, **my-app://ref1236a7d**. A QR code is created based on this URL, and is included in the carrier. Before you use this field, reach out to your Adyen contact to set up the QR code process.   Maximum length: 255 characters. | [optional] 
**BulkAddress** | Pointer to [**BulkAddress**](BulkAddress.md) |  | [optional] 
**CardImageId** | Pointer to **string** | The ID of the card image. This is the image that will be printed on the full front of the card. | [optional] 
**Carrier** | Pointer to **string** | Overrides the carrier design ID defined in the &#x60;configurationProfileId&#x60;. The carrier is the letter or packaging to which the card is attached. | [optional] 
**CarrierImageId** | Pointer to **string** | The ID of the carrier image. This is the image that will printed on the letter to which the card is attached. | [optional] 
**ConfigurationProfileId** | **string** | The ID of the card configuration profile that contains the settings of the card. For example, the envelope and PIN mailer designs or the logistics company handling the shipment. All the settings in the profile are applied to the card, unless you provide other fields to override them.  For example, send the &#x60;shipmentMethod&#x60; to override the logistics company defined in the card configuration profile. | 
**Currency** | Pointer to **string** | The three-letter [ISO-4217](https://en.wikipedia.org/wiki/ISO_4217) currency code of the card. For example, **EUR**. | [optional] 
**Envelope** | Pointer to **string** | Overrides the envelope design ID defined in the &#x60;configurationProfileId&#x60;.  | [optional] 
**Insert** | Pointer to **string** | Overrides the insert design ID defined in the &#x60;configurationProfileId&#x60;. An insert is any additional material, such as marketing materials, that are shipped together with the card. | [optional] 
**Language** | Pointer to **string** | The two-letter [ISO-639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language code of the card. For example, **en**. | [optional] 
**LogoImageId** | Pointer to **string** | The ID of the logo image. This is the image that will be printed on the partial front of the card, such as a logo on the upper right corner. | [optional] 
**PinMailer** | Pointer to **string** | Overrides the PIN mailer design ID defined in the &#x60;configurationProfileId&#x60;. The PIN mailer is the letter on which the PIN is printed. | [optional] 
**ShipmentMethod** | Pointer to **string** | Overrides the logistics company defined in the &#x60;configurationProfileId&#x60;. | [optional] 

## Methods

### NewCardConfiguration

`func NewCardConfiguration(configurationProfileId string, ) *CardConfiguration`

NewCardConfiguration instantiates a new CardConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardConfigurationWithDefaults

`func NewCardConfigurationWithDefaults() *CardConfiguration`

NewCardConfigurationWithDefaults instantiates a new CardConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivation

`func (o *CardConfiguration) GetActivation() string`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *CardConfiguration) GetActivationOk() (*string, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *CardConfiguration) SetActivation(v string)`

SetActivation sets Activation field to given value.

### HasActivation

`func (o *CardConfiguration) HasActivation() bool`

HasActivation returns a boolean if a field has been set.

### GetActivationUrl

`func (o *CardConfiguration) GetActivationUrl() string`

GetActivationUrl returns the ActivationUrl field if non-nil, zero value otherwise.

### GetActivationUrlOk

`func (o *CardConfiguration) GetActivationUrlOk() (*string, bool)`

GetActivationUrlOk returns a tuple with the ActivationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationUrl

`func (o *CardConfiguration) SetActivationUrl(v string)`

SetActivationUrl sets ActivationUrl field to given value.

### HasActivationUrl

`func (o *CardConfiguration) HasActivationUrl() bool`

HasActivationUrl returns a boolean if a field has been set.

### GetBulkAddress

`func (o *CardConfiguration) GetBulkAddress() BulkAddress`

GetBulkAddress returns the BulkAddress field if non-nil, zero value otherwise.

### GetBulkAddressOk

`func (o *CardConfiguration) GetBulkAddressOk() (*BulkAddress, bool)`

GetBulkAddressOk returns a tuple with the BulkAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkAddress

`func (o *CardConfiguration) SetBulkAddress(v BulkAddress)`

SetBulkAddress sets BulkAddress field to given value.

### HasBulkAddress

`func (o *CardConfiguration) HasBulkAddress() bool`

HasBulkAddress returns a boolean if a field has been set.

### GetCardImageId

`func (o *CardConfiguration) GetCardImageId() string`

GetCardImageId returns the CardImageId field if non-nil, zero value otherwise.

### GetCardImageIdOk

`func (o *CardConfiguration) GetCardImageIdOk() (*string, bool)`

GetCardImageIdOk returns a tuple with the CardImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardImageId

`func (o *CardConfiguration) SetCardImageId(v string)`

SetCardImageId sets CardImageId field to given value.

### HasCardImageId

`func (o *CardConfiguration) HasCardImageId() bool`

HasCardImageId returns a boolean if a field has been set.

### GetCarrier

`func (o *CardConfiguration) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *CardConfiguration) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *CardConfiguration) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *CardConfiguration) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCarrierImageId

`func (o *CardConfiguration) GetCarrierImageId() string`

GetCarrierImageId returns the CarrierImageId field if non-nil, zero value otherwise.

### GetCarrierImageIdOk

`func (o *CardConfiguration) GetCarrierImageIdOk() (*string, bool)`

GetCarrierImageIdOk returns a tuple with the CarrierImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierImageId

`func (o *CardConfiguration) SetCarrierImageId(v string)`

SetCarrierImageId sets CarrierImageId field to given value.

### HasCarrierImageId

`func (o *CardConfiguration) HasCarrierImageId() bool`

HasCarrierImageId returns a boolean if a field has been set.

### GetConfigurationProfileId

`func (o *CardConfiguration) GetConfigurationProfileId() string`

GetConfigurationProfileId returns the ConfigurationProfileId field if non-nil, zero value otherwise.

### GetConfigurationProfileIdOk

`func (o *CardConfiguration) GetConfigurationProfileIdOk() (*string, bool)`

GetConfigurationProfileIdOk returns a tuple with the ConfigurationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationProfileId

`func (o *CardConfiguration) SetConfigurationProfileId(v string)`

SetConfigurationProfileId sets ConfigurationProfileId field to given value.


### GetCurrency

`func (o *CardConfiguration) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CardConfiguration) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CardConfiguration) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CardConfiguration) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEnvelope

`func (o *CardConfiguration) GetEnvelope() string`

GetEnvelope returns the Envelope field if non-nil, zero value otherwise.

### GetEnvelopeOk

`func (o *CardConfiguration) GetEnvelopeOk() (*string, bool)`

GetEnvelopeOk returns a tuple with the Envelope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvelope

`func (o *CardConfiguration) SetEnvelope(v string)`

SetEnvelope sets Envelope field to given value.

### HasEnvelope

`func (o *CardConfiguration) HasEnvelope() bool`

HasEnvelope returns a boolean if a field has been set.

### GetInsert

`func (o *CardConfiguration) GetInsert() string`

GetInsert returns the Insert field if non-nil, zero value otherwise.

### GetInsertOk

`func (o *CardConfiguration) GetInsertOk() (*string, bool)`

GetInsertOk returns a tuple with the Insert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsert

`func (o *CardConfiguration) SetInsert(v string)`

SetInsert sets Insert field to given value.

### HasInsert

`func (o *CardConfiguration) HasInsert() bool`

HasInsert returns a boolean if a field has been set.

### GetLanguage

`func (o *CardConfiguration) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CardConfiguration) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CardConfiguration) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *CardConfiguration) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLogoImageId

`func (o *CardConfiguration) GetLogoImageId() string`

GetLogoImageId returns the LogoImageId field if non-nil, zero value otherwise.

### GetLogoImageIdOk

`func (o *CardConfiguration) GetLogoImageIdOk() (*string, bool)`

GetLogoImageIdOk returns a tuple with the LogoImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoImageId

`func (o *CardConfiguration) SetLogoImageId(v string)`

SetLogoImageId sets LogoImageId field to given value.

### HasLogoImageId

`func (o *CardConfiguration) HasLogoImageId() bool`

HasLogoImageId returns a boolean if a field has been set.

### GetPinMailer

`func (o *CardConfiguration) GetPinMailer() string`

GetPinMailer returns the PinMailer field if non-nil, zero value otherwise.

### GetPinMailerOk

`func (o *CardConfiguration) GetPinMailerOk() (*string, bool)`

GetPinMailerOk returns a tuple with the PinMailer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinMailer

`func (o *CardConfiguration) SetPinMailer(v string)`

SetPinMailer sets PinMailer field to given value.

### HasPinMailer

`func (o *CardConfiguration) HasPinMailer() bool`

HasPinMailer returns a boolean if a field has been set.

### GetShipmentMethod

`func (o *CardConfiguration) GetShipmentMethod() string`

GetShipmentMethod returns the ShipmentMethod field if non-nil, zero value otherwise.

### GetShipmentMethodOk

`func (o *CardConfiguration) GetShipmentMethodOk() (*string, bool)`

GetShipmentMethodOk returns a tuple with the ShipmentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentMethod

`func (o *CardConfiguration) SetShipmentMethod(v string)`

SetShipmentMethod sets ShipmentMethod field to given value.

### HasShipmentMethod

`func (o *CardConfiguration) HasShipmentMethod() bool`

HasShipmentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


