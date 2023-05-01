# PaymentInstrumentRevealInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cvc** | **string** | The CVC2 value of the card. | 
**Expiration** | [**Expiry**](Expiry.md) |  | 
**Pan** | **string** | The primary account number (PAN) of the card. | 

## Methods

### NewPaymentInstrumentRevealInfo

`func NewPaymentInstrumentRevealInfo(cvc string, expiration Expiry, pan string, ) *PaymentInstrumentRevealInfo`

NewPaymentInstrumentRevealInfo instantiates a new PaymentInstrumentRevealInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInstrumentRevealInfoWithDefaults

`func NewPaymentInstrumentRevealInfoWithDefaults() *PaymentInstrumentRevealInfo`

NewPaymentInstrumentRevealInfoWithDefaults instantiates a new PaymentInstrumentRevealInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvc

`func (o *PaymentInstrumentRevealInfo) GetCvc() string`

GetCvc returns the Cvc field if non-nil, zero value otherwise.

### GetCvcOk

`func (o *PaymentInstrumentRevealInfo) GetCvcOk() (*string, bool)`

GetCvcOk returns a tuple with the Cvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvc

`func (o *PaymentInstrumentRevealInfo) SetCvc(v string)`

SetCvc sets Cvc field to given value.


### GetExpiration

`func (o *PaymentInstrumentRevealInfo) GetExpiration() Expiry`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *PaymentInstrumentRevealInfo) GetExpirationOk() (*Expiry, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *PaymentInstrumentRevealInfo) SetExpiration(v Expiry)`

SetExpiration sets Expiration field to given value.


### GetPan

`func (o *PaymentInstrumentRevealInfo) GetPan() string`

GetPan returns the Pan field if non-nil, zero value otherwise.

### GetPanOk

`func (o *PaymentInstrumentRevealInfo) GetPanOk() (*string, bool)`

GetPanOk returns a tuple with the Pan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPan

`func (o *PaymentInstrumentRevealInfo) SetPan(v string)`

SetPan sets Pan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


