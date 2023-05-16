# TransactionSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Data** | Pointer to [**[]Transaction**](Transaction.md) | Contains the transactions that match the query parameters. | [optional] 

## Methods

### NewTransactionSearchResponse

`func NewTransactionSearchResponse() *TransactionSearchResponse`

NewTransactionSearchResponse instantiates a new TransactionSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSearchResponseWithDefaults

`func NewTransactionSearchResponseWithDefaults() *TransactionSearchResponse`

NewTransactionSearchResponseWithDefaults instantiates a new TransactionSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TransactionSearchResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TransactionSearchResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TransactionSearchResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TransactionSearchResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *TransactionSearchResponse) GetData() []Transaction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionSearchResponse) GetDataOk() (*[]Transaction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionSearchResponse) SetData(v []Transaction)`

SetData sets Data field to given value.

### HasData

`func (o *TransactionSearchResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


