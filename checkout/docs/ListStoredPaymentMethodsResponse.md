# ListStoredPaymentMethodsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantAccount** | Pointer to **string** | Your merchant account. | [optional] 
**ShopperReference** | Pointer to **string** | Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. &gt; Your reference must not include personally identifiable information (PII), for example name or email address. | [optional] 
**StoredPaymentMethods** | Pointer to [**[]StoredPaymentMethodResource**](StoredPaymentMethodResource.md) | List of all stored payment methods. | [optional] 

## Methods

### NewListStoredPaymentMethodsResponse

`func NewListStoredPaymentMethodsResponse() *ListStoredPaymentMethodsResponse`

NewListStoredPaymentMethodsResponse instantiates a new ListStoredPaymentMethodsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStoredPaymentMethodsResponseWithDefaults

`func NewListStoredPaymentMethodsResponseWithDefaults() *ListStoredPaymentMethodsResponse`

NewListStoredPaymentMethodsResponseWithDefaults instantiates a new ListStoredPaymentMethodsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantAccount

`func (o *ListStoredPaymentMethodsResponse) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *ListStoredPaymentMethodsResponse) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *ListStoredPaymentMethodsResponse) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.

### HasMerchantAccount

`func (o *ListStoredPaymentMethodsResponse) HasMerchantAccount() bool`

HasMerchantAccount returns a boolean if a field has been set.

### GetShopperReference

`func (o *ListStoredPaymentMethodsResponse) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *ListStoredPaymentMethodsResponse) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *ListStoredPaymentMethodsResponse) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *ListStoredPaymentMethodsResponse) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.

### GetStoredPaymentMethods

`func (o *ListStoredPaymentMethodsResponse) GetStoredPaymentMethods() []StoredPaymentMethodResource`

GetStoredPaymentMethods returns the StoredPaymentMethods field if non-nil, zero value otherwise.

### GetStoredPaymentMethodsOk

`func (o *ListStoredPaymentMethodsResponse) GetStoredPaymentMethodsOk() (*[]StoredPaymentMethodResource, bool)`

GetStoredPaymentMethodsOk returns a tuple with the StoredPaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredPaymentMethods

`func (o *ListStoredPaymentMethodsResponse) SetStoredPaymentMethods(v []StoredPaymentMethodResource)`

SetStoredPaymentMethods sets StoredPaymentMethods field to given value.

### HasStoredPaymentMethods

`func (o *ListStoredPaymentMethodsResponse) HasStoredPaymentMethods() bool`

HasStoredPaymentMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


