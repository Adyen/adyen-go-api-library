# RequestActivationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** | The unique identifier of the company account. | [optional] 
**MerchantId** | Pointer to **string** | The unique identifier of the merchant account you requested to activate. | [optional] 

## Methods

### NewRequestActivationResponse

`func NewRequestActivationResponse() *RequestActivationResponse`

NewRequestActivationResponse instantiates a new RequestActivationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestActivationResponseWithDefaults

`func NewRequestActivationResponseWithDefaults() *RequestActivationResponse`

NewRequestActivationResponseWithDefaults instantiates a new RequestActivationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *RequestActivationResponse) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *RequestActivationResponse) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *RequestActivationResponse) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *RequestActivationResponse) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetMerchantId

`func (o *RequestActivationResponse) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *RequestActivationResponse) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *RequestActivationResponse) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *RequestActivationResponse) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


