# Card

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | Pointer to [**Authentication**](Authentication.md) |  | [optional] 
**Bin** | Pointer to **string** | The bank identification number (BIN) of the card number. | [optional] 
**Brand** | **string** | The brand of the physical or the virtual card. Possible values: **visa**, **mc**. | 
**BrandVariant** | **string** | The brand variant of the physical or the virtual card. &gt;Contact your Adyen Implementation Manager to get the values that are relevant to your integration. Examples: **visadebit**, **mcprepaid**. | 
**CardholderName** | **string** | The name of the cardholder.  Maximum length: 26 characters. | 
**Configuration** | Pointer to [**CardConfiguration**](CardConfiguration.md) |  | [optional] 
**Cvc** | Pointer to **string** | The CVC2 value of the card. &gt; The CVC2 is not sent by default. This is only returned in the &#x60;POST&#x60; response for single-use virtual cards. | [optional] 
**DeliveryContact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Expiration** | Pointer to [**Expiry**](Expiry.md) |  | [optional] 
**FormFactor** | **string** | The form factor of the card. Possible values: **virtual**, **physical**. | 
**LastFour** | Pointer to **string** | Last last four digits of the card number. | [optional] 
**Number** | **string** | The primary account number (PAN) of the card. &gt; The PAN is masked by default and returned only for single-use virtual cards. | 

## Methods

### NewCard

`func NewCard(brand string, brandVariant string, cardholderName string, formFactor string, number string, ) *Card`

NewCard instantiates a new Card object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardWithDefaults

`func NewCardWithDefaults() *Card`

NewCardWithDefaults instantiates a new Card object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthentication

`func (o *Card) GetAuthentication() Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *Card) GetAuthenticationOk() (*Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *Card) SetAuthentication(v Authentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *Card) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetBin

`func (o *Card) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *Card) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *Card) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *Card) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetBrand

`func (o *Card) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *Card) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *Card) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetBrandVariant

`func (o *Card) GetBrandVariant() string`

GetBrandVariant returns the BrandVariant field if non-nil, zero value otherwise.

### GetBrandVariantOk

`func (o *Card) GetBrandVariantOk() (*string, bool)`

GetBrandVariantOk returns a tuple with the BrandVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandVariant

`func (o *Card) SetBrandVariant(v string)`

SetBrandVariant sets BrandVariant field to given value.


### GetCardholderName

`func (o *Card) GetCardholderName() string`

GetCardholderName returns the CardholderName field if non-nil, zero value otherwise.

### GetCardholderNameOk

`func (o *Card) GetCardholderNameOk() (*string, bool)`

GetCardholderNameOk returns a tuple with the CardholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderName

`func (o *Card) SetCardholderName(v string)`

SetCardholderName sets CardholderName field to given value.


### GetConfiguration

`func (o *Card) GetConfiguration() CardConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Card) GetConfigurationOk() (*CardConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Card) SetConfiguration(v CardConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *Card) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetCvc

`func (o *Card) GetCvc() string`

GetCvc returns the Cvc field if non-nil, zero value otherwise.

### GetCvcOk

`func (o *Card) GetCvcOk() (*string, bool)`

GetCvcOk returns a tuple with the Cvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvc

`func (o *Card) SetCvc(v string)`

SetCvc sets Cvc field to given value.

### HasCvc

`func (o *Card) HasCvc() bool`

HasCvc returns a boolean if a field has been set.

### GetDeliveryContact

`func (o *Card) GetDeliveryContact() Contact`

GetDeliveryContact returns the DeliveryContact field if non-nil, zero value otherwise.

### GetDeliveryContactOk

`func (o *Card) GetDeliveryContactOk() (*Contact, bool)`

GetDeliveryContactOk returns a tuple with the DeliveryContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryContact

`func (o *Card) SetDeliveryContact(v Contact)`

SetDeliveryContact sets DeliveryContact field to given value.

### HasDeliveryContact

`func (o *Card) HasDeliveryContact() bool`

HasDeliveryContact returns a boolean if a field has been set.

### GetExpiration

`func (o *Card) GetExpiration() Expiry`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *Card) GetExpirationOk() (*Expiry, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *Card) SetExpiration(v Expiry)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *Card) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetFormFactor

`func (o *Card) GetFormFactor() string`

GetFormFactor returns the FormFactor field if non-nil, zero value otherwise.

### GetFormFactorOk

`func (o *Card) GetFormFactorOk() (*string, bool)`

GetFormFactorOk returns a tuple with the FormFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFactor

`func (o *Card) SetFormFactor(v string)`

SetFormFactor sets FormFactor field to given value.


### GetLastFour

`func (o *Card) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *Card) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *Card) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *Card) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.

### GetNumber

`func (o *Card) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Card) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Card) SetNumber(v string)`

SetNumber sets Number field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


