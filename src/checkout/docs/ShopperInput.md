# ShopperInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddress** | Pointer to **string** | Specifies visibility of billing address fields.  Permitted values: * editable * hidden * readOnly | [optional] 
**DeliveryAddress** | Pointer to **string** | Specifies visibility of delivery address fields.  Permitted values: * editable * hidden * readOnly | [optional] 
**PersonalDetails** | Pointer to **string** | Specifies visibility of personal details.  Permitted values: * editable * hidden * readOnly | [optional] 

## Methods

### NewShopperInput

`func NewShopperInput() *ShopperInput`

NewShopperInput instantiates a new ShopperInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopperInputWithDefaults

`func NewShopperInputWithDefaults() *ShopperInput`

NewShopperInputWithDefaults instantiates a new ShopperInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAddress

`func (o *ShopperInput) GetBillingAddress() string`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *ShopperInput) GetBillingAddressOk() (*string, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *ShopperInput) SetBillingAddress(v string)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *ShopperInput) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetDeliveryAddress

`func (o *ShopperInput) GetDeliveryAddress() string`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *ShopperInput) GetDeliveryAddressOk() (*string, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *ShopperInput) SetDeliveryAddress(v string)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *ShopperInput) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### GetPersonalDetails

`func (o *ShopperInput) GetPersonalDetails() string`

GetPersonalDetails returns the PersonalDetails field if non-nil, zero value otherwise.

### GetPersonalDetailsOk

`func (o *ShopperInput) GetPersonalDetailsOk() (*string, bool)`

GetPersonalDetailsOk returns a tuple with the PersonalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDetails

`func (o *ShopperInput) SetPersonalDetails(v string)`

SetPersonalDetails sets PersonalDetails field to given value.

### HasPersonalDetails

`func (o *ShopperInput) HasPersonalDetails() bool`

HasPersonalDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


