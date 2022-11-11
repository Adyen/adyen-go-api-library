# StoreCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**Address2**](Address2.md) |  | 
**BusinessLineIds** | Pointer to **[]string** | The unique identifiers of the [business lines](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/businesslines__resParam_id) that the store is associated with. If not specified, the business line of the merchant account is used. Required when there are multiple business lines under the merchant account. | [optional] 
**Description** | **string** | Your description of the store. | 
**ExternalReferenceId** | Pointer to **string** | When using the Zip payment method: The location ID that Zip has assigned to your store. | [optional] 
**PhoneNumber** | **string** | The phone number of the store, including &#39;+&#39; and country code. | 
**Reference** | Pointer to **string** | Your reference to recognize the store by. Also known as the store code.  Allowed characters: Lowercase and uppercase letters without diacritics, numbers 0 through 9, hyphen (-), and underscore (_). | [optional] 
**ShopperStatement** | **string** | The store name to be shown on the shopper&#39;s bank or credit card statement and on the shopper receipt. Maximum length: 22 characters; can&#39;t be all numbers. | 
**SplitConfiguration** | Pointer to [**StoreSplitConfiguration**](StoreSplitConfiguration.md) |  | [optional] 

## Methods

### NewStoreCreationRequest

`func NewStoreCreationRequest(address Address2, description string, phoneNumber string, shopperStatement string, ) *StoreCreationRequest`

NewStoreCreationRequest instantiates a new StoreCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreCreationRequestWithDefaults

`func NewStoreCreationRequestWithDefaults() *StoreCreationRequest`

NewStoreCreationRequestWithDefaults instantiates a new StoreCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *StoreCreationRequest) GetAddress() Address2`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *StoreCreationRequest) GetAddressOk() (*Address2, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *StoreCreationRequest) SetAddress(v Address2)`

SetAddress sets Address field to given value.


### GetBusinessLineIds

`func (o *StoreCreationRequest) GetBusinessLineIds() []string`

GetBusinessLineIds returns the BusinessLineIds field if non-nil, zero value otherwise.

### GetBusinessLineIdsOk

`func (o *StoreCreationRequest) GetBusinessLineIdsOk() (*[]string, bool)`

GetBusinessLineIdsOk returns a tuple with the BusinessLineIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessLineIds

`func (o *StoreCreationRequest) SetBusinessLineIds(v []string)`

SetBusinessLineIds sets BusinessLineIds field to given value.

### HasBusinessLineIds

`func (o *StoreCreationRequest) HasBusinessLineIds() bool`

HasBusinessLineIds returns a boolean if a field has been set.

### GetDescription

`func (o *StoreCreationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StoreCreationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StoreCreationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExternalReferenceId

`func (o *StoreCreationRequest) GetExternalReferenceId() string`

GetExternalReferenceId returns the ExternalReferenceId field if non-nil, zero value otherwise.

### GetExternalReferenceIdOk

`func (o *StoreCreationRequest) GetExternalReferenceIdOk() (*string, bool)`

GetExternalReferenceIdOk returns a tuple with the ExternalReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReferenceId

`func (o *StoreCreationRequest) SetExternalReferenceId(v string)`

SetExternalReferenceId sets ExternalReferenceId field to given value.

### HasExternalReferenceId

`func (o *StoreCreationRequest) HasExternalReferenceId() bool`

HasExternalReferenceId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *StoreCreationRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *StoreCreationRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *StoreCreationRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetReference

`func (o *StoreCreationRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *StoreCreationRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *StoreCreationRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *StoreCreationRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetShopperStatement

`func (o *StoreCreationRequest) GetShopperStatement() string`

GetShopperStatement returns the ShopperStatement field if non-nil, zero value otherwise.

### GetShopperStatementOk

`func (o *StoreCreationRequest) GetShopperStatementOk() (*string, bool)`

GetShopperStatementOk returns a tuple with the ShopperStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperStatement

`func (o *StoreCreationRequest) SetShopperStatement(v string)`

SetShopperStatement sets ShopperStatement field to given value.


### GetSplitConfiguration

`func (o *StoreCreationRequest) GetSplitConfiguration() StoreSplitConfiguration`

GetSplitConfiguration returns the SplitConfiguration field if non-nil, zero value otherwise.

### GetSplitConfigurationOk

`func (o *StoreCreationRequest) GetSplitConfigurationOk() (*StoreSplitConfiguration, bool)`

GetSplitConfigurationOk returns a tuple with the SplitConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitConfiguration

`func (o *StoreCreationRequest) SetSplitConfiguration(v StoreSplitConfiguration)`

SetSplitConfiguration sets SplitConfiguration field to given value.

### HasSplitConfiguration

`func (o *StoreCreationRequest) HasSplitConfiguration() bool`

HasSplitConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


