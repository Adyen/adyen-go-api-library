# AuthenticationResultRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantAccount** | **string** | The merchant account identifier, with which the authentication was processed. | 
**PspReference** | **string** | The pspReference identifier for the transaction. | 

## Methods

### NewAuthenticationResultRequest

`func NewAuthenticationResultRequest(merchantAccount string, pspReference string, ) *AuthenticationResultRequest`

NewAuthenticationResultRequest instantiates a new AuthenticationResultRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationResultRequestWithDefaults

`func NewAuthenticationResultRequestWithDefaults() *AuthenticationResultRequest`

NewAuthenticationResultRequestWithDefaults instantiates a new AuthenticationResultRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantAccount

`func (o *AuthenticationResultRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *AuthenticationResultRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *AuthenticationResultRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetPspReference

`func (o *AuthenticationResultRequest) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *AuthenticationResultRequest) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *AuthenticationResultRequest) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


