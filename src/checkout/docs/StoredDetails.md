# StoredDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bank** | Pointer to [**BankAccount**](BankAccount.md) |  | [optional] 
**Card** | Pointer to [**Card**](Card.md) |  | [optional] 
**EmailAddress** | Pointer to **string** | The email associated with stored payment details. | [optional] 

## Methods

### NewStoredDetails

`func NewStoredDetails() *StoredDetails`

NewStoredDetails instantiates a new StoredDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoredDetailsWithDefaults

`func NewStoredDetailsWithDefaults() *StoredDetails`

NewStoredDetailsWithDefaults instantiates a new StoredDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBank

`func (o *StoredDetails) GetBank() BankAccount`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *StoredDetails) GetBankOk() (*BankAccount, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *StoredDetails) SetBank(v BankAccount)`

SetBank sets Bank field to given value.

### HasBank

`func (o *StoredDetails) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetCard

`func (o *StoredDetails) GetCard() Card`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *StoredDetails) GetCardOk() (*Card, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *StoredDetails) SetCard(v Card)`

SetCard sets Card field to given value.

### HasCard

`func (o *StoredDetails) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetEmailAddress

`func (o *StoredDetails) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *StoredDetails) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *StoredDetails) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *StoredDetails) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


