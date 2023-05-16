# BillingEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Email** | Pointer to **string** | The email address of the billing entity. | [optional] 
**Id** | Pointer to **string** | The unique identifier of the billing entity, for use as &#x60;billingEntityId&#x60; when creating an order. | [optional] 
**Name** | Pointer to **string** | The unique name of the billing entity. | [optional] 
**TaxId** | Pointer to **string** | The tax number of the billing entity. | [optional] 

## Methods

### NewBillingEntity

`func NewBillingEntity() *BillingEntity`

NewBillingEntity instantiates a new BillingEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingEntityWithDefaults

`func NewBillingEntityWithDefaults() *BillingEntity`

NewBillingEntityWithDefaults instantiates a new BillingEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *BillingEntity) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BillingEntity) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BillingEntity) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BillingEntity) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetEmail

`func (o *BillingEntity) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BillingEntity) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BillingEntity) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BillingEntity) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *BillingEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BillingEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingEntity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTaxId

`func (o *BillingEntity) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *BillingEntity) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *BillingEntity) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *BillingEntity) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


