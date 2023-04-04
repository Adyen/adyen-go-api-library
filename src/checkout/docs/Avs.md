# Avs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressEditable** | Pointer to **bool** | Indicates whether the shopper is allowed to modify the billing address for the current payment request. | [optional] 
**Enabled** | Pointer to **string** | Specifies whether the shopper should enter their billing address during checkout.  Allowed values: * yes — Perform AVS checks for every card payment. * automatic — Perform AVS checks only when required to optimize the conversion rate. * no — Do not perform AVS checks. | [optional] 

## Methods

### NewAvs

`func NewAvs() *Avs`

NewAvs instantiates a new Avs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvsWithDefaults

`func NewAvsWithDefaults() *Avs`

NewAvsWithDefaults instantiates a new Avs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressEditable

`func (o *Avs) GetAddressEditable() bool`

GetAddressEditable returns the AddressEditable field if non-nil, zero value otherwise.

### GetAddressEditableOk

`func (o *Avs) GetAddressEditableOk() (*bool, bool)`

GetAddressEditableOk returns a tuple with the AddressEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressEditable

`func (o *Avs) SetAddressEditable(v bool)`

SetAddressEditable sets AddressEditable field to given value.

### HasAddressEditable

`func (o *Avs) HasAddressEditable() bool`

HasAddressEditable returns a boolean if a field has been set.

### GetEnabled

`func (o *Avs) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Avs) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Avs) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Avs) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


