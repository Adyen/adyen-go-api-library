# GetTerminalsUnderAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyAccount** | **string** | Your company account. If you only specify this parameter, the response includes all terminals at all account levels. | 
**MerchantAccount** | Pointer to **string** | The merchant account. This is required if you are retrieving the terminals assigned to a store.If you don&#39;t specify a &#x60;store&#x60; the response includes the terminals assigned to the specified merchant account and the terminals assigned to the stores under this merchant account. | [optional] 
**Store** | Pointer to **string** | The store code of the store. With this parameter, the response only includes the terminals assigned to the specified store. | [optional] 

## Methods

### NewGetTerminalsUnderAccountRequest

`func NewGetTerminalsUnderAccountRequest(companyAccount string, ) *GetTerminalsUnderAccountRequest`

NewGetTerminalsUnderAccountRequest instantiates a new GetTerminalsUnderAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTerminalsUnderAccountRequestWithDefaults

`func NewGetTerminalsUnderAccountRequestWithDefaults() *GetTerminalsUnderAccountRequest`

NewGetTerminalsUnderAccountRequestWithDefaults instantiates a new GetTerminalsUnderAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyAccount

`func (o *GetTerminalsUnderAccountRequest) GetCompanyAccount() string`

GetCompanyAccount returns the CompanyAccount field if non-nil, zero value otherwise.

### GetCompanyAccountOk

`func (o *GetTerminalsUnderAccountRequest) GetCompanyAccountOk() (*string, bool)`

GetCompanyAccountOk returns a tuple with the CompanyAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAccount

`func (o *GetTerminalsUnderAccountRequest) SetCompanyAccount(v string)`

SetCompanyAccount sets CompanyAccount field to given value.


### GetMerchantAccount

`func (o *GetTerminalsUnderAccountRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *GetTerminalsUnderAccountRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *GetTerminalsUnderAccountRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.

### HasMerchantAccount

`func (o *GetTerminalsUnderAccountRequest) HasMerchantAccount() bool`

HasMerchantAccount returns a boolean if a field has been set.

### GetStore

`func (o *GetTerminalsUnderAccountRequest) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *GetTerminalsUnderAccountRequest) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *GetTerminalsUnderAccountRequest) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *GetTerminalsUnderAccountRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


