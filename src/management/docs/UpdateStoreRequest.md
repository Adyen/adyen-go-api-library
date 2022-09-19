# UpdateStoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**UpdatableAddress**](UpdatableAddress.md) |  | [optional] 
**BusinessLineIds** | Pointer to **[]string** | The unique identifiers of the [business lines](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/businesslines__resParam_id) that the store is associated with. | [optional] 
**Description** | Pointer to **string** | The description of the store. | [optional] 
**ExternalReferenceId** | Pointer to **string** | When using the Zip payment method: The location ID that Zip has assigned to your store. | [optional] 
**SplitConfiguration** | Pointer to [**StoreSplitConfiguration**](StoreSplitConfiguration.md) |  | [optional] 
**Status** | Pointer to **string** | The status of the store. Possible values are:  - **active**: This value is assigned automatically when a store is created.  - **inactive**: The maximum [transaction limits and number of Store-and-Forward transactions](https://docs.adyen.com/point-of-sale/determine-account-structure/configure-features#payment-features) for the store are set to 0. This blocks new transactions, but captures are still possible. - **closed**: The terminals of the store are reassigned to the merchant inventory, so they can&#39;t process payments.  You can change the status from **active** to **inactive**, and from **inactive** to **active** or **closed**.  Once **closed**, a store can&#39;t be reopened. | [optional] 

## Methods

### NewUpdateStoreRequest

`func NewUpdateStoreRequest() *UpdateStoreRequest`

NewUpdateStoreRequest instantiates a new UpdateStoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStoreRequestWithDefaults

`func NewUpdateStoreRequestWithDefaults() *UpdateStoreRequest`

NewUpdateStoreRequestWithDefaults instantiates a new UpdateStoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UpdateStoreRequest) GetAddress() UpdatableAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateStoreRequest) GetAddressOk() (*UpdatableAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateStoreRequest) SetAddress(v UpdatableAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateStoreRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBusinessLineIds

`func (o *UpdateStoreRequest) GetBusinessLineIds() []string`

GetBusinessLineIds returns the BusinessLineIds field if non-nil, zero value otherwise.

### GetBusinessLineIdsOk

`func (o *UpdateStoreRequest) GetBusinessLineIdsOk() (*[]string, bool)`

GetBusinessLineIdsOk returns a tuple with the BusinessLineIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessLineIds

`func (o *UpdateStoreRequest) SetBusinessLineIds(v []string)`

SetBusinessLineIds sets BusinessLineIds field to given value.

### HasBusinessLineIds

`func (o *UpdateStoreRequest) HasBusinessLineIds() bool`

HasBusinessLineIds returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateStoreRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateStoreRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateStoreRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateStoreRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalReferenceId

`func (o *UpdateStoreRequest) GetExternalReferenceId() string`

GetExternalReferenceId returns the ExternalReferenceId field if non-nil, zero value otherwise.

### GetExternalReferenceIdOk

`func (o *UpdateStoreRequest) GetExternalReferenceIdOk() (*string, bool)`

GetExternalReferenceIdOk returns a tuple with the ExternalReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReferenceId

`func (o *UpdateStoreRequest) SetExternalReferenceId(v string)`

SetExternalReferenceId sets ExternalReferenceId field to given value.

### HasExternalReferenceId

`func (o *UpdateStoreRequest) HasExternalReferenceId() bool`

HasExternalReferenceId returns a boolean if a field has been set.

### GetSplitConfiguration

`func (o *UpdateStoreRequest) GetSplitConfiguration() StoreSplitConfiguration`

GetSplitConfiguration returns the SplitConfiguration field if non-nil, zero value otherwise.

### GetSplitConfigurationOk

`func (o *UpdateStoreRequest) GetSplitConfigurationOk() (*StoreSplitConfiguration, bool)`

GetSplitConfigurationOk returns a tuple with the SplitConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitConfiguration

`func (o *UpdateStoreRequest) SetSplitConfiguration(v StoreSplitConfiguration)`

SetSplitConfiguration sets SplitConfiguration field to given value.

### HasSplitConfiguration

`func (o *UpdateStoreRequest) HasSplitConfiguration() bool`

HasSplitConfiguration returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateStoreRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateStoreRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateStoreRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateStoreRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


