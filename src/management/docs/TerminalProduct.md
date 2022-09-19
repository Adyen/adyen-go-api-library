# TerminalProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Information about items included and integration options. | [optional] 
**Id** | Pointer to **string** | The unique identifier of the product. | [optional] 
**ItemsIncluded** | Pointer to **[]string** | A list of parts included in the terminal package. | [optional] 
**Name** | Pointer to **string** | The descriptive name of the product. | [optional] 
**Price** | Pointer to [**Amount**](Amount.md) |  | [optional] 

## Methods

### NewTerminalProduct

`func NewTerminalProduct() *TerminalProduct`

NewTerminalProduct instantiates a new TerminalProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalProductWithDefaults

`func NewTerminalProductWithDefaults() *TerminalProduct`

NewTerminalProductWithDefaults instantiates a new TerminalProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TerminalProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TerminalProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TerminalProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TerminalProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *TerminalProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TerminalProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TerminalProduct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TerminalProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemsIncluded

`func (o *TerminalProduct) GetItemsIncluded() []string`

GetItemsIncluded returns the ItemsIncluded field if non-nil, zero value otherwise.

### GetItemsIncludedOk

`func (o *TerminalProduct) GetItemsIncludedOk() (*[]string, bool)`

GetItemsIncludedOk returns a tuple with the ItemsIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsIncluded

`func (o *TerminalProduct) SetItemsIncluded(v []string)`

SetItemsIncluded sets ItemsIncluded field to given value.

### HasItemsIncluded

`func (o *TerminalProduct) HasItemsIncluded() bool`

HasItemsIncluded returns a boolean if a field has been set.

### GetName

`func (o *TerminalProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TerminalProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TerminalProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TerminalProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *TerminalProduct) GetPrice() Amount`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TerminalProduct) GetPriceOk() (*Amount, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TerminalProduct) SetPrice(v Amount)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *TerminalProduct) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


