# BillingEntitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]BillingEntity**](BillingEntity.md) | List of legal entities that can be used for the billing of orders. | [optional] 

## Methods

### NewBillingEntitiesResponse

`func NewBillingEntitiesResponse() *BillingEntitiesResponse`

NewBillingEntitiesResponse instantiates a new BillingEntitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingEntitiesResponseWithDefaults

`func NewBillingEntitiesResponseWithDefaults() *BillingEntitiesResponse`

NewBillingEntitiesResponseWithDefaults instantiates a new BillingEntitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BillingEntitiesResponse) GetData() []BillingEntity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BillingEntitiesResponse) GetDataOk() (*[]BillingEntity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BillingEntitiesResponse) SetData(v []BillingEntity)`

SetData sets Data field to given value.

### HasData

`func (o *BillingEntitiesResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


