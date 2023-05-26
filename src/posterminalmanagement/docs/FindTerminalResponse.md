# FindTerminalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyAccount** | **string** | The company account that the terminal is associated with. If this is the only account level shown in the response, the terminal is assigned to the inventory of the company account. | 
**MerchantAccount** | Pointer to **string** | The merchant account that the terminal is associated with. If the response doesn&#39;t contain a &#x60;store&#x60; the terminal is assigned to this merchant account. | [optional] 
**MerchantInventory** | Pointer to **bool** | Boolean that indicates if the terminal is assigned to the merchant inventory. This is returned when the terminal is assigned to a merchant account.  - If **true**, this indicates that the terminal is in the merchant inventory. This also means that the terminal cannot be boarded.  - If **false**, this indicates that the terminal is assigned to the merchant account as an in-store terminal. This means that the terminal is ready to be boarded, or is already boarded. | [optional] 
**Store** | Pointer to **string** | The store code of the store that the terminal is assigned to. | [optional] 
**Terminal** | **string** | The unique terminal ID. | 

## Methods

### NewFindTerminalResponse

`func NewFindTerminalResponse(companyAccount string, terminal string, ) *FindTerminalResponse`

NewFindTerminalResponse instantiates a new FindTerminalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindTerminalResponseWithDefaults

`func NewFindTerminalResponseWithDefaults() *FindTerminalResponse`

NewFindTerminalResponseWithDefaults instantiates a new FindTerminalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyAccount

`func (o *FindTerminalResponse) GetCompanyAccount() string`

GetCompanyAccount returns the CompanyAccount field if non-nil, zero value otherwise.

### GetCompanyAccountOk

`func (o *FindTerminalResponse) GetCompanyAccountOk() (*string, bool)`

GetCompanyAccountOk returns a tuple with the CompanyAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAccount

`func (o *FindTerminalResponse) SetCompanyAccount(v string)`

SetCompanyAccount sets CompanyAccount field to given value.


### GetMerchantAccount

`func (o *FindTerminalResponse) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *FindTerminalResponse) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *FindTerminalResponse) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.

### HasMerchantAccount

`func (o *FindTerminalResponse) HasMerchantAccount() bool`

HasMerchantAccount returns a boolean if a field has been set.

### GetMerchantInventory

`func (o *FindTerminalResponse) GetMerchantInventory() bool`

GetMerchantInventory returns the MerchantInventory field if non-nil, zero value otherwise.

### GetMerchantInventoryOk

`func (o *FindTerminalResponse) GetMerchantInventoryOk() (*bool, bool)`

GetMerchantInventoryOk returns a tuple with the MerchantInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantInventory

`func (o *FindTerminalResponse) SetMerchantInventory(v bool)`

SetMerchantInventory sets MerchantInventory field to given value.

### HasMerchantInventory

`func (o *FindTerminalResponse) HasMerchantInventory() bool`

HasMerchantInventory returns a boolean if a field has been set.

### GetStore

`func (o *FindTerminalResponse) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *FindTerminalResponse) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *FindTerminalResponse) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *FindTerminalResponse) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetTerminal

`func (o *FindTerminalResponse) GetTerminal() string`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *FindTerminalResponse) GetTerminalOk() (*string, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *FindTerminalResponse) SetTerminal(v string)`

SetTerminal sets Terminal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


