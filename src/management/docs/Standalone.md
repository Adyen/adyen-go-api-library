# Standalone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | Pointer to **string** | The default currency of the standalone payment terminal as an [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code. | [optional] 
**EnableStandalone** | Pointer to **bool** | Enable standalone mode. | [optional] 

## Methods

### NewStandalone

`func NewStandalone() *Standalone`

NewStandalone instantiates a new Standalone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandaloneWithDefaults

`func NewStandaloneWithDefaults() *Standalone`

NewStandaloneWithDefaults instantiates a new Standalone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *Standalone) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Standalone) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Standalone) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *Standalone) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetEnableStandalone

`func (o *Standalone) GetEnableStandalone() bool`

GetEnableStandalone returns the EnableStandalone field if non-nil, zero value otherwise.

### GetEnableStandaloneOk

`func (o *Standalone) GetEnableStandaloneOk() (*bool, bool)`

GetEnableStandaloneOk returns a tuple with the EnableStandalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStandalone

`func (o *Standalone) SetEnableStandalone(v bool)`

SetEnableStandalone sets EnableStandalone field to given value.

### HasEnableStandalone

`func (o *Standalone) HasEnableStandalone() bool`

HasEnableStandalone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


