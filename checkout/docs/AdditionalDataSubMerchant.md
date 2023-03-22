# AdditionalDataSubMerchant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubMerchantNumberOfSubSellers** | Pointer to **string** | Required for transactions performed by registered payment facilitators. Indicates the number of sub-merchants contained in the request. For example, **3**. | [optional] 
**SubMerchantSubSellerSubSellerNrCity** | Pointer to **string** | Required for transactions performed by registered payment facilitators. The city of the sub-merchant&#39;s address. * Format: Alphanumeric * Maximum length: 13 characters | [optional] 
**SubMerchantSubSellerSubSellerNrCountry** | Pointer to **string** | Required for transactions performed by registered payment facilitators. The three-letter country code of the sub-merchant&#39;s address. For example, **BRA** for Brazil.  * Format: [ISO 3166-1 alpha-3](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3) * Fixed length: 3 characters | [optional] 
**SubMerchantSubSellerSubSellerNrId** | Pointer to **string** | Required for transactions performed by registered payment facilitators. A unique identifier that you create for the sub-merchant, used by schemes to identify the sub-merchant.  * Format: Alphanumeric * Maximum length: 15 characters | [optional] 
**SubMerchantSubSellerSubSellerNrMcc** | Pointer to **string** | Required for transactions performed by registered payment facilitators. The sub-merchant&#39;s 4-digit Merchant Category Code (MCC).  * Format: Numeric * Fixed length: 4 digits | [optional] 
**SubMerchantSubSellerSubSellerNrName** | Pointer to **string** | Required for transactions performed by registered payment facilitators. The name of the sub-merchant. Based on scheme specifications, this value will overwrite the shopper statement  that will appear in the card statement. * Format: Alphanumeric * Maximum length: 22 characters | [optional] 
**SubMerchantSubSellerSubSellerNrPostalCode** | Pointer to **string** | Required for transactions performed by registered payment facilitators. The postal code of the sub-merchant&#39;s address, without dashes. * Format: Numeric * Fixed length: 8 digits | [optional] 
**SubMerchantSubSellerSubSellerNrState** | Pointer to **string** | Required for transactions performed by registered payment facilitators. The state code of the sub-merchant&#39;s address, if applicable to the country. * Format: Alphanumeric * Maximum length: 2 characters | [optional] 
**SubMerchantSubSellerSubSellerNrStreet** | Pointer to **string** | Required for transactions performed by registered payment facilitators. The street name and house number of the sub-merchant&#39;s address. * Format: Alphanumeric * Maximum length: 60 characters | [optional] 
**SubMerchantSubSellerSubSellerNrTaxId** | Pointer to **string** | Required for transactions performed by registered payment facilitators. The tax ID of the sub-merchant. * Format: Numeric * Fixed length: 11 digits for the CPF or 14 digits for the CNPJ | [optional] 

## Methods

### NewAdditionalDataSubMerchant

`func NewAdditionalDataSubMerchant() *AdditionalDataSubMerchant`

NewAdditionalDataSubMerchant instantiates a new AdditionalDataSubMerchant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalDataSubMerchantWithDefaults

`func NewAdditionalDataSubMerchantWithDefaults() *AdditionalDataSubMerchant`

NewAdditionalDataSubMerchantWithDefaults instantiates a new AdditionalDataSubMerchant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubMerchantNumberOfSubSellers

`func (o *AdditionalDataSubMerchant) GetSubMerchantNumberOfSubSellers() string`

GetSubMerchantNumberOfSubSellers returns the SubMerchantNumberOfSubSellers field if non-nil, zero value otherwise.

### GetSubMerchantNumberOfSubSellersOk

`func (o *AdditionalDataSubMerchant) GetSubMerchantNumberOfSubSellersOk() (*string, bool)`

GetSubMerchantNumberOfSubSellersOk returns a tuple with the SubMerchantNumberOfSubSellers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantNumberOfSubSellers

`func (o *AdditionalDataSubMerchant) SetSubMerchantNumberOfSubSellers(v string)`

SetSubMerchantNumberOfSubSellers sets SubMerchantNumberOfSubSellers field to given value.

### HasSubMerchantNumberOfSubSellers

`func (o *AdditionalDataSubMerchant) HasSubMerchantNumberOfSubSellers() bool`

HasSubMerchantNumberOfSubSellers returns a boolean if a field has been set.

### GetSubMerchantSubSellerSubSellerNrCity

`func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrCity() string`

GetSubMerchantSubSellerSubSellerNrCity returns the SubMerchantSubSellerSubSellerNrCity field if non-nil, zero value otherwise.

### GetSubMerchantSubSellerSubSellerNrCityOk

`func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrCityOk() (*string, bool)`

GetSubMerchantSubSellerSubSellerNrCityOk returns a tuple with the SubMerchantSubSellerSubSellerNrCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantSubSellerSubSellerNrCity

`func (o *AdditionalDataSubMerchant) SetSubMerchantSubSellerSubSellerNrCity(v string)`

SetSubMerchantSubSellerSubSellerNrCity sets SubMerchantSubSellerSubSellerNrCity field to given value.

### HasSubMerchantSubSellerSubSellerNrCity

`func (o *AdditionalDataSubMerchant) HasSubMerchantSubSellerSubSellerNrCity() bool`

HasSubMerchantSubSellerSubSellerNrCity returns a boolean if a field has been set.

### GetSubMerchantSubSellerSubSellerNrCountry

`func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrCountry() string`

GetSubMerchantSubSellerSubSellerNrCountry returns the SubMerchantSubSellerSubSellerNrCountry field if non-nil, zero value otherwise.

### GetSubMerchantSubSellerSubSellerNrCountryOk

`func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrCountryOk() (*string, bool)`

GetSubMerchantSubSellerSubSellerNrCountryOk returns a tuple with the SubMerchantSubSellerSubSellerNrCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantSubSellerSubSellerNrCountry

`func (o *AdditionalDataSubMerchant) SetSubMerchantSubSellerSubSellerNrCountry(v string)`

SetSubMerchantSubSellerSubSellerNrCountry sets SubMerchantSubSellerSubSellerNrCountry field to given value.

### HasSubMerchantSubSellerSubSellerNrCountry

`func (o *AdditionalDataSubMerchant) HasSubMerchantSubSellerSubSellerNrCountry() bool`

HasSubMerchantSubSellerSubSellerNrCountry returns a boolean if a field has been set.

### GetSubMerchantSubSellerSubSellerNrId

`func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrId() string`

GetSubMerchantSubSellerSubSellerNrId returns the SubMerchantSubSellerSubSellerNrId field if non-nil, zero value otherwise.

### GetSubMerchantSubSellerSubSellerNrIdOk

`func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrIdOk() (*string, bool)`

GetSubMerchantSubSellerSubSellerNrIdOk returns a tuple with the SubMerchantSubSellerSubSellerNrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantSubSellerSubSellerNrId

`func (o *AdditionalDataSubMerchant) SetSubMerchantSubSellerSubSellerNrId(v string)`

SetSubMerchantSubSellerSubSellerNrId sets SubMerchantSubSellerSubSellerNrId field to given value.

### HasSubMerchantSubSellerSubSellerNrId

`func (o *AdditionalDataSubMerchant) HasSubMerchantSubSellerSubSellerNrId() bool`

HasSubMerchantSubSellerSubSellerNrId returns a boolean if a field has been set.

### GetSubMerchantSubSellerSubSellerNrMcc

`func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrMcc() string`

GetSubMerchantSubSellerSubSellerNrMcc returns the SubMerchantSubSellerSubSellerNrMcc field if non-nil, zero value otherwise.

### GetSubMerchantSubSellerSubSellerNrMccOk

`func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrMccOk() (*string, bool)`

GetSubMerchantSubSellerSubSellerNrMccOk returns a tuple with the SubMerchantSubSellerSubSellerNrMcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantSubSellerSubSellerNrMcc

`func (o *AdditionalDataSubMerchant) SetSubMerchantSubSellerSubSellerNrMcc(v string)`

SetSubMerchantSubSellerSubSellerNrMcc sets SubMerchantSubSellerSubSellerNrMcc field to given value.

### HasSubMerchantSubSellerSubSellerNrMcc

`func (o *AdditionalDataSubMerchant) HasSubMerchantSubSellerSubSellerNrMcc() bool`

HasSubMerchantSubSellerSubSellerNrMcc returns a boolean if a field has been set.

### GetSubMerchantSubSellerSubSellerNrName

`func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrName() string`

GetSubMerchantSubSellerSubSellerNrName returns the SubMerchantSubSellerSubSellerNrName field if non-nil, zero value otherwise.

### GetSubMerchantSubSellerSubSellerNrNameOk

`func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrNameOk() (*string, bool)`

GetSubMerchantSubSellerSubSellerNrNameOk returns a tuple with the SubMerchantSubSellerSubSellerNrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantSubSellerSubSellerNrName

`func (o *AdditionalDataSubMerchant) SetSubMerchantSubSellerSubSellerNrName(v string)`

SetSubMerchantSubSellerSubSellerNrName sets SubMerchantSubSellerSubSellerNrName field to given value.

### HasSubMerchantSubSellerSubSellerNrName

`func (o *AdditionalDataSubMerchant) HasSubMerchantSubSellerSubSellerNrName() bool`

HasSubMerchantSubSellerSubSellerNrName returns a boolean if a field has been set.

### GetSubMerchantSubSellerSubSellerNrPostalCode

`func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrPostalCode() string`

GetSubMerchantSubSellerSubSellerNrPostalCode returns the SubMerchantSubSellerSubSellerNrPostalCode field if non-nil, zero value otherwise.

### GetSubMerchantSubSellerSubSellerNrPostalCodeOk

`func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrPostalCodeOk() (*string, bool)`

GetSubMerchantSubSellerSubSellerNrPostalCodeOk returns a tuple with the SubMerchantSubSellerSubSellerNrPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantSubSellerSubSellerNrPostalCode

`func (o *AdditionalDataSubMerchant) SetSubMerchantSubSellerSubSellerNrPostalCode(v string)`

SetSubMerchantSubSellerSubSellerNrPostalCode sets SubMerchantSubSellerSubSellerNrPostalCode field to given value.

### HasSubMerchantSubSellerSubSellerNrPostalCode

`func (o *AdditionalDataSubMerchant) HasSubMerchantSubSellerSubSellerNrPostalCode() bool`

HasSubMerchantSubSellerSubSellerNrPostalCode returns a boolean if a field has been set.

### GetSubMerchantSubSellerSubSellerNrState

`func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrState() string`

GetSubMerchantSubSellerSubSellerNrState returns the SubMerchantSubSellerSubSellerNrState field if non-nil, zero value otherwise.

### GetSubMerchantSubSellerSubSellerNrStateOk

`func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrStateOk() (*string, bool)`

GetSubMerchantSubSellerSubSellerNrStateOk returns a tuple with the SubMerchantSubSellerSubSellerNrState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantSubSellerSubSellerNrState

`func (o *AdditionalDataSubMerchant) SetSubMerchantSubSellerSubSellerNrState(v string)`

SetSubMerchantSubSellerSubSellerNrState sets SubMerchantSubSellerSubSellerNrState field to given value.

### HasSubMerchantSubSellerSubSellerNrState

`func (o *AdditionalDataSubMerchant) HasSubMerchantSubSellerSubSellerNrState() bool`

HasSubMerchantSubSellerSubSellerNrState returns a boolean if a field has been set.

### GetSubMerchantSubSellerSubSellerNrStreet

`func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrStreet() string`

GetSubMerchantSubSellerSubSellerNrStreet returns the SubMerchantSubSellerSubSellerNrStreet field if non-nil, zero value otherwise.

### GetSubMerchantSubSellerSubSellerNrStreetOk

`func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrStreetOk() (*string, bool)`

GetSubMerchantSubSellerSubSellerNrStreetOk returns a tuple with the SubMerchantSubSellerSubSellerNrStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantSubSellerSubSellerNrStreet

`func (o *AdditionalDataSubMerchant) SetSubMerchantSubSellerSubSellerNrStreet(v string)`

SetSubMerchantSubSellerSubSellerNrStreet sets SubMerchantSubSellerSubSellerNrStreet field to given value.

### HasSubMerchantSubSellerSubSellerNrStreet

`func (o *AdditionalDataSubMerchant) HasSubMerchantSubSellerSubSellerNrStreet() bool`

HasSubMerchantSubSellerSubSellerNrStreet returns a boolean if a field has been set.

### GetSubMerchantSubSellerSubSellerNrTaxId

`func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrTaxId() string`

GetSubMerchantSubSellerSubSellerNrTaxId returns the SubMerchantSubSellerSubSellerNrTaxId field if non-nil, zero value otherwise.

### GetSubMerchantSubSellerSubSellerNrTaxIdOk

`func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrTaxIdOk() (*string, bool)`

GetSubMerchantSubSellerSubSellerNrTaxIdOk returns a tuple with the SubMerchantSubSellerSubSellerNrTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantSubSellerSubSellerNrTaxId

`func (o *AdditionalDataSubMerchant) SetSubMerchantSubSellerSubSellerNrTaxId(v string)`

SetSubMerchantSubSellerSubSellerNrTaxId sets SubMerchantSubSellerSubSellerNrTaxId field to given value.

### HasSubMerchantSubSellerSubSellerNrTaxId

`func (o *AdditionalDataSubMerchant) HasSubMerchantSubSellerSubSellerNrTaxId() bool`

HasSubMerchantSubSellerSubSellerNrTaxId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


