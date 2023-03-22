# CheckoutUtilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginKeys** | Pointer to **map[string]string** | The list of origin keys for all requested domains. For each list item, the key is the domain and the value is the origin key. | [optional] 

## Methods

### NewCheckoutUtilityResponse

`func NewCheckoutUtilityResponse() *CheckoutUtilityResponse`

NewCheckoutUtilityResponse instantiates a new CheckoutUtilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutUtilityResponseWithDefaults

`func NewCheckoutUtilityResponseWithDefaults() *CheckoutUtilityResponse`

NewCheckoutUtilityResponseWithDefaults instantiates a new CheckoutUtilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginKeys

`func (o *CheckoutUtilityResponse) GetOriginKeys() map[string]string`

GetOriginKeys returns the OriginKeys field if non-nil, zero value otherwise.

### GetOriginKeysOk

`func (o *CheckoutUtilityResponse) GetOriginKeysOk() (*map[string]string, bool)`

GetOriginKeysOk returns a tuple with the OriginKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginKeys

`func (o *CheckoutUtilityResponse) SetOriginKeys(v map[string]string)`

SetOriginKeys sets OriginKeys field to given value.

### HasOriginKeys

`func (o *CheckoutUtilityResponse) HasOriginKeys() bool`

HasOriginKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


