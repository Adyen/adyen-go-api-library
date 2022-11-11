# Opi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnablePayAtTable** | Pointer to **bool** | Indicates if Pay at Table is enabled. | [optional] 
**PayAtTableStoreNumber** | Pointer to **string** | The store number to use for Pay at Table. | [optional] 
**PayAtTableURL** | Pointer to **string** | The URL and port number used for Pay at Table communication. | [optional] 

## Methods

### NewOpi

`func NewOpi() *Opi`

NewOpi instantiates a new Opi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpiWithDefaults

`func NewOpiWithDefaults() *Opi`

NewOpiWithDefaults instantiates a new Opi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnablePayAtTable

`func (o *Opi) GetEnablePayAtTable() bool`

GetEnablePayAtTable returns the EnablePayAtTable field if non-nil, zero value otherwise.

### GetEnablePayAtTableOk

`func (o *Opi) GetEnablePayAtTableOk() (*bool, bool)`

GetEnablePayAtTableOk returns a tuple with the EnablePayAtTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePayAtTable

`func (o *Opi) SetEnablePayAtTable(v bool)`

SetEnablePayAtTable sets EnablePayAtTable field to given value.

### HasEnablePayAtTable

`func (o *Opi) HasEnablePayAtTable() bool`

HasEnablePayAtTable returns a boolean if a field has been set.

### GetPayAtTableStoreNumber

`func (o *Opi) GetPayAtTableStoreNumber() string`

GetPayAtTableStoreNumber returns the PayAtTableStoreNumber field if non-nil, zero value otherwise.

### GetPayAtTableStoreNumberOk

`func (o *Opi) GetPayAtTableStoreNumberOk() (*string, bool)`

GetPayAtTableStoreNumberOk returns a tuple with the PayAtTableStoreNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayAtTableStoreNumber

`func (o *Opi) SetPayAtTableStoreNumber(v string)`

SetPayAtTableStoreNumber sets PayAtTableStoreNumber field to given value.

### HasPayAtTableStoreNumber

`func (o *Opi) HasPayAtTableStoreNumber() bool`

HasPayAtTableStoreNumber returns a boolean if a field has been set.

### GetPayAtTableURL

`func (o *Opi) GetPayAtTableURL() string`

GetPayAtTableURL returns the PayAtTableURL field if non-nil, zero value otherwise.

### GetPayAtTableURLOk

`func (o *Opi) GetPayAtTableURLOk() (*string, bool)`

GetPayAtTableURLOk returns a tuple with the PayAtTableURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayAtTableURL

`func (o *Opi) SetPayAtTableURL(v string)`

SetPayAtTableURL sets PayAtTableURL field to given value.

### HasPayAtTableURL

`func (o *Opi) HasPayAtTableURL() bool`

HasPayAtTableURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


