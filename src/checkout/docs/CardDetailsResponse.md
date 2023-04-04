# CardDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brands** | Pointer to [**[]CardBrandDetails**](CardBrandDetails.md) | The list of brands identified for the card. | [optional] 

## Methods

### NewCardDetailsResponse

`func NewCardDetailsResponse() *CardDetailsResponse`

NewCardDetailsResponse instantiates a new CardDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardDetailsResponseWithDefaults

`func NewCardDetailsResponseWithDefaults() *CardDetailsResponse`

NewCardDetailsResponseWithDefaults instantiates a new CardDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrands

`func (o *CardDetailsResponse) GetBrands() []CardBrandDetails`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *CardDetailsResponse) GetBrandsOk() (*[]CardBrandDetails, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *CardDetailsResponse) SetBrands(v []CardBrandDetails)`

SetBrands sets Brands field to given value.

### HasBrands

`func (o *CardDetailsResponse) HasBrands() bool`

HasBrands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


