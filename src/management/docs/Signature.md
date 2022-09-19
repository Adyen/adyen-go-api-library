# Signature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskSignatureOnScreen** | Pointer to **bool** | If &#x60;skipSignature&#x60; is false, indicates whether the shopper should provide a signature on the display (**true**) or on the merchant receipt (**false**). | [optional] 
**DeviceName** | Pointer to **string** | Name that identifies the terminal. | [optional] 
**SkipSignature** | Pointer to **bool** | Skip asking for a signature. This is possible because all global card schemes (American Express, Diners, Discover, JCB, MasterCard, VISA, and UnionPay) regard a signature as optional. | [optional] 

## Methods

### NewSignature

`func NewSignature() *Signature`

NewSignature instantiates a new Signature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureWithDefaults

`func NewSignatureWithDefaults() *Signature`

NewSignatureWithDefaults instantiates a new Signature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskSignatureOnScreen

`func (o *Signature) GetAskSignatureOnScreen() bool`

GetAskSignatureOnScreen returns the AskSignatureOnScreen field if non-nil, zero value otherwise.

### GetAskSignatureOnScreenOk

`func (o *Signature) GetAskSignatureOnScreenOk() (*bool, bool)`

GetAskSignatureOnScreenOk returns a tuple with the AskSignatureOnScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSignatureOnScreen

`func (o *Signature) SetAskSignatureOnScreen(v bool)`

SetAskSignatureOnScreen sets AskSignatureOnScreen field to given value.

### HasAskSignatureOnScreen

`func (o *Signature) HasAskSignatureOnScreen() bool`

HasAskSignatureOnScreen returns a boolean if a field has been set.

### GetDeviceName

`func (o *Signature) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *Signature) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *Signature) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *Signature) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetSkipSignature

`func (o *Signature) GetSkipSignature() bool`

GetSkipSignature returns the SkipSignature field if non-nil, zero value otherwise.

### GetSkipSignatureOk

`func (o *Signature) GetSkipSignatureOk() (*bool, bool)`

GetSkipSignatureOk returns a tuple with the SkipSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSignature

`func (o *Signature) SetSkipSignature(v bool)`

SetSkipSignature sets SkipSignature field to given value.

### HasSkipSignature

`func (o *Signature) HasSkipSignature() bool`

HasSkipSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


