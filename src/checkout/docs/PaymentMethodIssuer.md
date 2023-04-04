# PaymentMethodIssuer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | A boolean value indicating whether this issuer is unavailable. Can be &#x60;true&#x60; whenever the issuer is offline. | [optional] [default to false]
**Id** | **string** | The unique identifier of this issuer, to submit in requests to /payments. | 
**Name** | **string** | A localized name of the issuer. | 

## Methods

### NewPaymentMethodIssuer

`func NewPaymentMethodIssuer(id string, name string, ) *PaymentMethodIssuer`

NewPaymentMethodIssuer instantiates a new PaymentMethodIssuer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodIssuerWithDefaults

`func NewPaymentMethodIssuerWithDefaults() *PaymentMethodIssuer`

NewPaymentMethodIssuerWithDefaults instantiates a new PaymentMethodIssuer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *PaymentMethodIssuer) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *PaymentMethodIssuer) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *PaymentMethodIssuer) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *PaymentMethodIssuer) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetId

`func (o *PaymentMethodIssuer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodIssuer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodIssuer) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PaymentMethodIssuer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentMethodIssuer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentMethodIssuer) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


