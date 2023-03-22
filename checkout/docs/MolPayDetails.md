# MolPayDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**Issuer** | **string** | The shopper&#39;s bank. Specify this with the issuer value that corresponds to this bank. | 
**Type** | **string** | **molpay** | 

## Methods

### NewMolPayDetails

`func NewMolPayDetails(issuer string, type_ string, ) *MolPayDetails`

NewMolPayDetails instantiates a new MolPayDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMolPayDetailsWithDefaults

`func NewMolPayDetailsWithDefaults() *MolPayDetails`

NewMolPayDetailsWithDefaults instantiates a new MolPayDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutAttemptId

`func (o *MolPayDetails) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *MolPayDetails) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *MolPayDetails) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *MolPayDetails) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetIssuer

`func (o *MolPayDetails) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *MolPayDetails) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *MolPayDetails) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetType

`func (o *MolPayDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MolPayDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MolPayDetails) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


