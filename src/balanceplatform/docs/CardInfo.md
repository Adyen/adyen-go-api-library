# CardInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | Pointer to [**Authentication**](Authentication.md) |  | [optional] 
**Brand** | **string** | The brand of the physical or the virtual card. Possible values: **visa**, **mc**. | 
**BrandVariant** | **string** | The brand variant of the physical or the virtual card. &gt;Contact your Adyen Implementation Manager to get the values that are relevant to your integration. Examples: **visadebit**, **mcprepaid**. | 
**CardholderName** | **string** | The name of the cardholder.  Maximum length: 26 characters. | 
**Configuration** | Pointer to [**CardConfiguration**](CardConfiguration.md) |  | [optional] 
**DeliveryContact** | Pointer to [**DeliveryContact**](DeliveryContact.md) |  | [optional] 
**FormFactor** | **string** | The form factor of the card. Possible values: **virtual**, **physical**. | 

## Methods

### NewCardInfo

`func NewCardInfo(brand string, brandVariant string, cardholderName string, formFactor string, ) *CardInfo`

NewCardInfo instantiates a new CardInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardInfoWithDefaults

`func NewCardInfoWithDefaults() *CardInfo`

NewCardInfoWithDefaults instantiates a new CardInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthentication

`func (o *CardInfo) GetAuthentication() Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *CardInfo) GetAuthenticationOk() (*Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *CardInfo) SetAuthentication(v Authentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *CardInfo) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetBrand

`func (o *CardInfo) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CardInfo) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CardInfo) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetBrandVariant

`func (o *CardInfo) GetBrandVariant() string`

GetBrandVariant returns the BrandVariant field if non-nil, zero value otherwise.

### GetBrandVariantOk

`func (o *CardInfo) GetBrandVariantOk() (*string, bool)`

GetBrandVariantOk returns a tuple with the BrandVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandVariant

`func (o *CardInfo) SetBrandVariant(v string)`

SetBrandVariant sets BrandVariant field to given value.


### GetCardholderName

`func (o *CardInfo) GetCardholderName() string`

GetCardholderName returns the CardholderName field if non-nil, zero value otherwise.

### GetCardholderNameOk

`func (o *CardInfo) GetCardholderNameOk() (*string, bool)`

GetCardholderNameOk returns a tuple with the CardholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderName

`func (o *CardInfo) SetCardholderName(v string)`

SetCardholderName sets CardholderName field to given value.


### GetConfiguration

`func (o *CardInfo) GetConfiguration() CardConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CardInfo) GetConfigurationOk() (*CardConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CardInfo) SetConfiguration(v CardConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *CardInfo) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDeliveryContact

`func (o *CardInfo) GetDeliveryContact() DeliveryContact`

GetDeliveryContact returns the DeliveryContact field if non-nil, zero value otherwise.

### GetDeliveryContactOk

`func (o *CardInfo) GetDeliveryContactOk() (*DeliveryContact, bool)`

GetDeliveryContactOk returns a tuple with the DeliveryContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryContact

`func (o *CardInfo) SetDeliveryContact(v DeliveryContact)`

SetDeliveryContact sets DeliveryContact field to given value.

### HasDeliveryContact

`func (o *CardInfo) HasDeliveryContact() bool`

HasDeliveryContact returns a boolean if a field has been set.

### GetFormFactor

`func (o *CardInfo) GetFormFactor() string`

GetFormFactor returns the FormFactor field if non-nil, zero value otherwise.

### GetFormFactorOk

`func (o *CardInfo) GetFormFactorOk() (*string, bool)`

GetFormFactorOk returns a tuple with the FormFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFactor

`func (o *CardInfo) SetFormFactor(v string)`

SetFormFactor sets FormFactor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


