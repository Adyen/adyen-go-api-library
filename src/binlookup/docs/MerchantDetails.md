# MerchantDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | 2-letter ISO 3166 country code of the card acceptor location. &gt; This parameter is required for the merchants who don&#39;t use Adyen as the payment authorisation gateway. | [optional] 
**EnrolledIn3DSecure** | Pointer to **bool** | If true, indicates that the merchant is enrolled in 3D Secure for the card network. | [optional] 
**Mcc** | Pointer to **string** | The merchant category code (MCC) is a four-digit number which relates to a particular market segment. This code reflects the predominant activity that is conducted by the merchant.  The list of MCCs can be found [here](https://en.wikipedia.org/wiki/Merchant_category_code). | [optional] 

## Methods

### NewMerchantDetails

`func NewMerchantDetails() *MerchantDetails`

NewMerchantDetails instantiates a new MerchantDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantDetailsWithDefaults

`func NewMerchantDetailsWithDefaults() *MerchantDetails`

NewMerchantDetailsWithDefaults instantiates a new MerchantDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *MerchantDetails) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *MerchantDetails) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *MerchantDetails) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *MerchantDetails) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetEnrolledIn3DSecure

`func (o *MerchantDetails) GetEnrolledIn3DSecure() bool`

GetEnrolledIn3DSecure returns the EnrolledIn3DSecure field if non-nil, zero value otherwise.

### GetEnrolledIn3DSecureOk

`func (o *MerchantDetails) GetEnrolledIn3DSecureOk() (*bool, bool)`

GetEnrolledIn3DSecureOk returns a tuple with the EnrolledIn3DSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolledIn3DSecure

`func (o *MerchantDetails) SetEnrolledIn3DSecure(v bool)`

SetEnrolledIn3DSecure sets EnrolledIn3DSecure field to given value.

### HasEnrolledIn3DSecure

`func (o *MerchantDetails) HasEnrolledIn3DSecure() bool`

HasEnrolledIn3DSecure returns a boolean if a field has been set.

### GetMcc

`func (o *MerchantDetails) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *MerchantDetails) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *MerchantDetails) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *MerchantDetails) HasMcc() bool`

HasMcc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


