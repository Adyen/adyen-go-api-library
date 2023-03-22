# MbwayDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**ShopperEmail** | **string** |  | 
**TelephoneNumber** | **string** |  | 
**Type** | Pointer to **string** | **mbway** | [optional] [default to "mbway"]

## Methods

### NewMbwayDetails

`func NewMbwayDetails(shopperEmail string, telephoneNumber string, ) *MbwayDetails`

NewMbwayDetails instantiates a new MbwayDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbwayDetailsWithDefaults

`func NewMbwayDetailsWithDefaults() *MbwayDetails`

NewMbwayDetailsWithDefaults instantiates a new MbwayDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutAttemptId

`func (o *MbwayDetails) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *MbwayDetails) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *MbwayDetails) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *MbwayDetails) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetShopperEmail

`func (o *MbwayDetails) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *MbwayDetails) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *MbwayDetails) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.


### GetTelephoneNumber

`func (o *MbwayDetails) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *MbwayDetails) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *MbwayDetails) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.


### GetType

`func (o *MbwayDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MbwayDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MbwayDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MbwayDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


