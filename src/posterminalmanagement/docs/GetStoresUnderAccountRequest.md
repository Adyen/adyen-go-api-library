# GetStoresUnderAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyAccount** | **string** | The company account. If you only specify this parameter, the response includes the stores of all merchant accounts that are associated with the company account. | 
**MerchantAccount** | Pointer to **string** | The merchant account. With this parameter, the response only includes the stores of the specified merchant account. | [optional] 

## Methods

### NewGetStoresUnderAccountRequest

`func NewGetStoresUnderAccountRequest(companyAccount string, ) *GetStoresUnderAccountRequest`

NewGetStoresUnderAccountRequest instantiates a new GetStoresUnderAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStoresUnderAccountRequestWithDefaults

`func NewGetStoresUnderAccountRequestWithDefaults() *GetStoresUnderAccountRequest`

NewGetStoresUnderAccountRequestWithDefaults instantiates a new GetStoresUnderAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyAccount

`func (o *GetStoresUnderAccountRequest) GetCompanyAccount() string`

GetCompanyAccount returns the CompanyAccount field if non-nil, zero value otherwise.

### GetCompanyAccountOk

`func (o *GetStoresUnderAccountRequest) GetCompanyAccountOk() (*string, bool)`

GetCompanyAccountOk returns a tuple with the CompanyAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAccount

`func (o *GetStoresUnderAccountRequest) SetCompanyAccount(v string)`

SetCompanyAccount sets CompanyAccount field to given value.


### GetMerchantAccount

`func (o *GetStoresUnderAccountRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *GetStoresUnderAccountRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *GetStoresUnderAccountRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.

### HasMerchantAccount

`func (o *GetStoresUnderAccountRequest) HasMerchantAccount() bool`

HasMerchantAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


