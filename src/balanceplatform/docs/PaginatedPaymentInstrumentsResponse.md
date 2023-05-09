# PaginatedPaymentInstrumentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNext** | **bool** | Indicates whether there are more items on the next page. | 
**HasPrevious** | **bool** | Indicates whether there are more items on the previous page. | 
**PaymentInstruments** | [**[]PaymentInstrument**](PaymentInstrument.md) | List of payment instruments associated with the balance account. | 

## Methods

### NewPaginatedPaymentInstrumentsResponse

`func NewPaginatedPaymentInstrumentsResponse(hasNext bool, hasPrevious bool, paymentInstruments []PaymentInstrument, ) *PaginatedPaymentInstrumentsResponse`

NewPaginatedPaymentInstrumentsResponse instantiates a new PaginatedPaymentInstrumentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedPaymentInstrumentsResponseWithDefaults

`func NewPaginatedPaymentInstrumentsResponseWithDefaults() *PaginatedPaymentInstrumentsResponse`

NewPaginatedPaymentInstrumentsResponseWithDefaults instantiates a new PaginatedPaymentInstrumentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNext

`func (o *PaginatedPaymentInstrumentsResponse) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *PaginatedPaymentInstrumentsResponse) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *PaginatedPaymentInstrumentsResponse) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.


### GetHasPrevious

`func (o *PaginatedPaymentInstrumentsResponse) GetHasPrevious() bool`

GetHasPrevious returns the HasPrevious field if non-nil, zero value otherwise.

### GetHasPreviousOk

`func (o *PaginatedPaymentInstrumentsResponse) GetHasPreviousOk() (*bool, bool)`

GetHasPreviousOk returns a tuple with the HasPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrevious

`func (o *PaginatedPaymentInstrumentsResponse) SetHasPrevious(v bool)`

SetHasPrevious sets HasPrevious field to given value.


### GetPaymentInstruments

`func (o *PaginatedPaymentInstrumentsResponse) GetPaymentInstruments() []PaymentInstrument`

GetPaymentInstruments returns the PaymentInstruments field if non-nil, zero value otherwise.

### GetPaymentInstrumentsOk

`func (o *PaginatedPaymentInstrumentsResponse) GetPaymentInstrumentsOk() (*[]PaymentInstrument, bool)`

GetPaymentInstrumentsOk returns a tuple with the PaymentInstruments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstruments

`func (o *PaginatedPaymentInstrumentsResponse) SetPaymentInstruments(v []PaymentInstrument)`

SetPaymentInstruments sets PaymentInstruments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


