# DetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationData** | Pointer to [**DetailsRequestAuthenticationData**](DetailsRequestAuthenticationData.md) |  | [optional] 
**Details** | [**PaymentCompletionDetails**](PaymentCompletionDetails.md) |  | 
**PaymentData** | Pointer to **string** | The &#x60;paymentData&#x60; value from the &#x60;/payments&#x60; response. Required if the &#x60;/payments&#x60; response returns this value.  | [optional] 
**ThreeDSAuthenticationOnly** | Pointer to **bool** | Change the &#x60;authenticationOnly&#x60; indicator originally set in the &#x60;/payments&#x60; request. Only needs to be set if you want to modify the value set previously. | [optional] 

## Methods

### NewDetailsRequest

`func NewDetailsRequest(details PaymentCompletionDetails, ) *DetailsRequest`

NewDetailsRequest instantiates a new DetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailsRequestWithDefaults

`func NewDetailsRequestWithDefaults() *DetailsRequest`

NewDetailsRequestWithDefaults instantiates a new DetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationData

`func (o *DetailsRequest) GetAuthenticationData() DetailsRequestAuthenticationData`

GetAuthenticationData returns the AuthenticationData field if non-nil, zero value otherwise.

### GetAuthenticationDataOk

`func (o *DetailsRequest) GetAuthenticationDataOk() (*DetailsRequestAuthenticationData, bool)`

GetAuthenticationDataOk returns a tuple with the AuthenticationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationData

`func (o *DetailsRequest) SetAuthenticationData(v DetailsRequestAuthenticationData)`

SetAuthenticationData sets AuthenticationData field to given value.

### HasAuthenticationData

`func (o *DetailsRequest) HasAuthenticationData() bool`

HasAuthenticationData returns a boolean if a field has been set.

### GetDetails

`func (o *DetailsRequest) GetDetails() PaymentCompletionDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DetailsRequest) GetDetailsOk() (*PaymentCompletionDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DetailsRequest) SetDetails(v PaymentCompletionDetails)`

SetDetails sets Details field to given value.


### GetPaymentData

`func (o *DetailsRequest) GetPaymentData() string`

GetPaymentData returns the PaymentData field if non-nil, zero value otherwise.

### GetPaymentDataOk

`func (o *DetailsRequest) GetPaymentDataOk() (*string, bool)`

GetPaymentDataOk returns a tuple with the PaymentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentData

`func (o *DetailsRequest) SetPaymentData(v string)`

SetPaymentData sets PaymentData field to given value.

### HasPaymentData

`func (o *DetailsRequest) HasPaymentData() bool`

HasPaymentData returns a boolean if a field has been set.

### GetThreeDSAuthenticationOnly

`func (o *DetailsRequest) GetThreeDSAuthenticationOnly() bool`

GetThreeDSAuthenticationOnly returns the ThreeDSAuthenticationOnly field if non-nil, zero value otherwise.

### GetThreeDSAuthenticationOnlyOk

`func (o *DetailsRequest) GetThreeDSAuthenticationOnlyOk() (*bool, bool)`

GetThreeDSAuthenticationOnlyOk returns a tuple with the ThreeDSAuthenticationOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSAuthenticationOnly

`func (o *DetailsRequest) SetThreeDSAuthenticationOnly(v bool)`

SetThreeDSAuthenticationOnly sets ThreeDSAuthenticationOnly field to given value.

### HasThreeDSAuthenticationOnly

`func (o *DetailsRequest) HasThreeDSAuthenticationOnly() bool`

HasThreeDSAuthenticationOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


