# Passcodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminMenuPin** | Pointer to **string** | The pin code used to enter the admin menu | [optional] 
**RefundPin** | Pointer to **string** | Allows merchant to create a dedicated PIN for standalone refund functionality | [optional] 
**ScreenLockPin** | Pointer to **string** | Passcode to unlock screen after sleep | [optional] 
**TxMenuPin** | Pointer to **string** | The PIN code used to enter the transaction menu | [optional] 

## Methods

### NewPasscodes

`func NewPasscodes() *Passcodes`

NewPasscodes instantiates a new Passcodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasscodesWithDefaults

`func NewPasscodesWithDefaults() *Passcodes`

NewPasscodesWithDefaults instantiates a new Passcodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminMenuPin

`func (o *Passcodes) GetAdminMenuPin() string`

GetAdminMenuPin returns the AdminMenuPin field if non-nil, zero value otherwise.

### GetAdminMenuPinOk

`func (o *Passcodes) GetAdminMenuPinOk() (*string, bool)`

GetAdminMenuPinOk returns a tuple with the AdminMenuPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminMenuPin

`func (o *Passcodes) SetAdminMenuPin(v string)`

SetAdminMenuPin sets AdminMenuPin field to given value.

### HasAdminMenuPin

`func (o *Passcodes) HasAdminMenuPin() bool`

HasAdminMenuPin returns a boolean if a field has been set.

### GetRefundPin

`func (o *Passcodes) GetRefundPin() string`

GetRefundPin returns the RefundPin field if non-nil, zero value otherwise.

### GetRefundPinOk

`func (o *Passcodes) GetRefundPinOk() (*string, bool)`

GetRefundPinOk returns a tuple with the RefundPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundPin

`func (o *Passcodes) SetRefundPin(v string)`

SetRefundPin sets RefundPin field to given value.

### HasRefundPin

`func (o *Passcodes) HasRefundPin() bool`

HasRefundPin returns a boolean if a field has been set.

### GetScreenLockPin

`func (o *Passcodes) GetScreenLockPin() string`

GetScreenLockPin returns the ScreenLockPin field if non-nil, zero value otherwise.

### GetScreenLockPinOk

`func (o *Passcodes) GetScreenLockPinOk() (*string, bool)`

GetScreenLockPinOk returns a tuple with the ScreenLockPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenLockPin

`func (o *Passcodes) SetScreenLockPin(v string)`

SetScreenLockPin sets ScreenLockPin field to given value.

### HasScreenLockPin

`func (o *Passcodes) HasScreenLockPin() bool`

HasScreenLockPin returns a boolean if a field has been set.

### GetTxMenuPin

`func (o *Passcodes) GetTxMenuPin() string`

GetTxMenuPin returns the TxMenuPin field if non-nil, zero value otherwise.

### GetTxMenuPinOk

`func (o *Passcodes) GetTxMenuPinOk() (*string, bool)`

GetTxMenuPinOk returns a tuple with the TxMenuPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMenuPin

`func (o *Passcodes) SetTxMenuPin(v string)`

SetTxMenuPin sets TxMenuPin field to given value.

### HasTxMenuPin

`func (o *Passcodes) HasTxMenuPin() bool`

HasTxMenuPin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


