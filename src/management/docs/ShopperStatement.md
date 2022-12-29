# ShopperStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DoingBusinessAsName** | **string** | The name you want to be shown on the shopper&#39;s bank or credit card statement. Maximum length: 22 characters; can&#39;t be all numbers. | 
**Type** | Pointer to **string** | The type of shopperstatement you want to use: fixed, append or dynamic | [optional] 

## Methods

### NewShopperStatement

`func NewShopperStatement(doingBusinessAsName string, ) *ShopperStatement`

NewShopperStatement instantiates a new ShopperStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopperStatementWithDefaults

`func NewShopperStatementWithDefaults() *ShopperStatement`

NewShopperStatementWithDefaults instantiates a new ShopperStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoingBusinessAsName

`func (o *ShopperStatement) GetDoingBusinessAsName() string`

GetDoingBusinessAsName returns the DoingBusinessAsName field if non-nil, zero value otherwise.

### GetDoingBusinessAsNameOk

`func (o *ShopperStatement) GetDoingBusinessAsNameOk() (*string, bool)`

GetDoingBusinessAsNameOk returns a tuple with the DoingBusinessAsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoingBusinessAsName

`func (o *ShopperStatement) SetDoingBusinessAsName(v string)`

SetDoingBusinessAsName sets DoingBusinessAsName field to given value.


### GetType

`func (o *ShopperStatement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShopperStatement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShopperStatement) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ShopperStatement) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


