# AuthenticationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptAuthentication** | Pointer to **string** | Indicates when 3D Secure authentication should be attempted. This overrides all other rules, including [Dynamic 3D Secure settings](https://docs.adyen.com/risk-management/dynamic-3d-secure).  Possible values:  * **always**: Perform 3D Secure authentication. * **never**: Don&#39;t perform 3D Secure authentication. If PSD2 SCA or other national regulations require authentication, the transaction gets declined. * **preferNo**: Do not perform 3D Secure authentication if not required by PSD2 SCA or other national regulations. | [optional] 
**AuthenticationOnly** | Pointer to **bool** | If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation. Default: *false**. | [optional] [default to false]
**ThreeDSRequestData** | Pointer to [**ThreeDSRequestData**](ThreeDSRequestData.md) |  | [optional] 

## Methods

### NewAuthenticationData

`func NewAuthenticationData() *AuthenticationData`

NewAuthenticationData instantiates a new AuthenticationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationDataWithDefaults

`func NewAuthenticationDataWithDefaults() *AuthenticationData`

NewAuthenticationDataWithDefaults instantiates a new AuthenticationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptAuthentication

`func (o *AuthenticationData) GetAttemptAuthentication() string`

GetAttemptAuthentication returns the AttemptAuthentication field if non-nil, zero value otherwise.

### GetAttemptAuthenticationOk

`func (o *AuthenticationData) GetAttemptAuthenticationOk() (*string, bool)`

GetAttemptAuthenticationOk returns a tuple with the AttemptAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptAuthentication

`func (o *AuthenticationData) SetAttemptAuthentication(v string)`

SetAttemptAuthentication sets AttemptAuthentication field to given value.

### HasAttemptAuthentication

`func (o *AuthenticationData) HasAttemptAuthentication() bool`

HasAttemptAuthentication returns a boolean if a field has been set.

### GetAuthenticationOnly

`func (o *AuthenticationData) GetAuthenticationOnly() bool`

GetAuthenticationOnly returns the AuthenticationOnly field if non-nil, zero value otherwise.

### GetAuthenticationOnlyOk

`func (o *AuthenticationData) GetAuthenticationOnlyOk() (*bool, bool)`

GetAuthenticationOnlyOk returns a tuple with the AuthenticationOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOnly

`func (o *AuthenticationData) SetAuthenticationOnly(v bool)`

SetAuthenticationOnly sets AuthenticationOnly field to given value.

### HasAuthenticationOnly

`func (o *AuthenticationData) HasAuthenticationOnly() bool`

HasAuthenticationOnly returns a boolean if a field has been set.

### GetThreeDSRequestData

`func (o *AuthenticationData) GetThreeDSRequestData() ThreeDSRequestData`

GetThreeDSRequestData returns the ThreeDSRequestData field if non-nil, zero value otherwise.

### GetThreeDSRequestDataOk

`func (o *AuthenticationData) GetThreeDSRequestDataOk() (*ThreeDSRequestData, bool)`

GetThreeDSRequestDataOk returns a tuple with the ThreeDSRequestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSRequestData

`func (o *AuthenticationData) SetThreeDSRequestData(v ThreeDSRequestData)`

SetThreeDSRequestData sets ThreeDSRequestData field to given value.

### HasThreeDSRequestData

`func (o *AuthenticationData) HasThreeDSRequestData() bool`

HasThreeDSRequestData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


