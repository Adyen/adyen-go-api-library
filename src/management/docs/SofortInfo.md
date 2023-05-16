# SofortInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | **string** | Sofort currency code. For example, **EUR**. | 
**Logo** | **string** | Sofort logo. Format: Base64-encoded string. | 

## Methods

### NewSofortInfo

`func NewSofortInfo(currencyCode string, logo string, ) *SofortInfo`

NewSofortInfo instantiates a new SofortInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSofortInfoWithDefaults

`func NewSofortInfoWithDefaults() *SofortInfo`

NewSofortInfoWithDefaults instantiates a new SofortInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *SofortInfo) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *SofortInfo) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *SofortInfo) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetLogo

`func (o *SofortInfo) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *SofortInfo) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *SofortInfo) SetLogo(v string)`

SetLogo sets Logo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


