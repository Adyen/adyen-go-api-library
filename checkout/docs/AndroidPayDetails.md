# AndroidPayDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**Type** | Pointer to **string** | **androidpay** | [optional] [default to "androidpay"]

## Methods

### NewAndroidPayDetails

`func NewAndroidPayDetails() *AndroidPayDetails`

NewAndroidPayDetails instantiates a new AndroidPayDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAndroidPayDetailsWithDefaults

`func NewAndroidPayDetailsWithDefaults() *AndroidPayDetails`

NewAndroidPayDetailsWithDefaults instantiates a new AndroidPayDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutAttemptId

`func (o *AndroidPayDetails) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *AndroidPayDetails) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *AndroidPayDetails) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *AndroidPayDetails) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetType

`func (o *AndroidPayDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AndroidPayDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AndroidPayDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AndroidPayDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


