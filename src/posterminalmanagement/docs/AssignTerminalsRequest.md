# AssignTerminalsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyAccount** | **string** | Your company account. To return terminals to the company inventory, specify only this parameter and the &#x60;terminals&#x60;. | 
**MerchantAccount** | Pointer to **string** | Name of the merchant account. Specify this parameter to assign terminals to this merchant account or to a store under this merchant account. | [optional] 
**MerchantInventory** | Pointer to **bool** | Boolean that indicates if you are assigning the terminals to the merchant inventory. Do not use when assigning terminals to a store. Required when assigning the terminal to a merchant account.  - Set this to **true** to assign the terminals to the merchant inventory. This also means that the terminals cannot be boarded.  - Set this to **false** to assign the terminals to the merchant account as in-store terminals. This makes the terminals ready to be boarded and to process payments through the specified merchant account. | [optional] 
**Store** | Pointer to **string** | The store code of the store that you want to assign the terminals to. | [optional] 
**Terminals** | **[]string** | Array containing a list of terminal IDs that you want to assign or reassign to the merchant account or store, or that you want to return to the company inventory.  For example, &#x60;[\&quot;V400m-324689776\&quot;,\&quot;P400Plus-329127412\&quot;]&#x60;. | 

## Methods

### NewAssignTerminalsRequest

`func NewAssignTerminalsRequest(companyAccount string, terminals []string, ) *AssignTerminalsRequest`

NewAssignTerminalsRequest instantiates a new AssignTerminalsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignTerminalsRequestWithDefaults

`func NewAssignTerminalsRequestWithDefaults() *AssignTerminalsRequest`

NewAssignTerminalsRequestWithDefaults instantiates a new AssignTerminalsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyAccount

`func (o *AssignTerminalsRequest) GetCompanyAccount() string`

GetCompanyAccount returns the CompanyAccount field if non-nil, zero value otherwise.

### GetCompanyAccountOk

`func (o *AssignTerminalsRequest) GetCompanyAccountOk() (*string, bool)`

GetCompanyAccountOk returns a tuple with the CompanyAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAccount

`func (o *AssignTerminalsRequest) SetCompanyAccount(v string)`

SetCompanyAccount sets CompanyAccount field to given value.


### GetMerchantAccount

`func (o *AssignTerminalsRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *AssignTerminalsRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *AssignTerminalsRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.

### HasMerchantAccount

`func (o *AssignTerminalsRequest) HasMerchantAccount() bool`

HasMerchantAccount returns a boolean if a field has been set.

### GetMerchantInventory

`func (o *AssignTerminalsRequest) GetMerchantInventory() bool`

GetMerchantInventory returns the MerchantInventory field if non-nil, zero value otherwise.

### GetMerchantInventoryOk

`func (o *AssignTerminalsRequest) GetMerchantInventoryOk() (*bool, bool)`

GetMerchantInventoryOk returns a tuple with the MerchantInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantInventory

`func (o *AssignTerminalsRequest) SetMerchantInventory(v bool)`

SetMerchantInventory sets MerchantInventory field to given value.

### HasMerchantInventory

`func (o *AssignTerminalsRequest) HasMerchantInventory() bool`

HasMerchantInventory returns a boolean if a field has been set.

### GetStore

`func (o *AssignTerminalsRequest) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *AssignTerminalsRequest) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *AssignTerminalsRequest) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *AssignTerminalsRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetTerminals

`func (o *AssignTerminalsRequest) GetTerminals() []string`

GetTerminals returns the Terminals field if non-nil, zero value otherwise.

### GetTerminalsOk

`func (o *AssignTerminalsRequest) GetTerminalsOk() (*[]string, bool)`

GetTerminalsOk returns a tuple with the Terminals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminals

`func (o *AssignTerminalsRequest) SetTerminals(v []string)`

SetTerminals sets Terminals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


