# PaymentSetupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentSession** | Pointer to **string** | The encoded payment session that you need to pass to the SDK. | [optional] 
**RecurringDetails** | Pointer to [**[]RecurringDetail**](RecurringDetail.md) | The detailed list of stored payment details required to generate payment forms. Will be empty if oneClick is set to false in the request. | [optional] 

## Methods

### NewPaymentSetupResponse

`func NewPaymentSetupResponse() *PaymentSetupResponse`

NewPaymentSetupResponse instantiates a new PaymentSetupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentSetupResponseWithDefaults

`func NewPaymentSetupResponseWithDefaults() *PaymentSetupResponse`

NewPaymentSetupResponseWithDefaults instantiates a new PaymentSetupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentSession

`func (o *PaymentSetupResponse) GetPaymentSession() string`

GetPaymentSession returns the PaymentSession field if non-nil, zero value otherwise.

### GetPaymentSessionOk

`func (o *PaymentSetupResponse) GetPaymentSessionOk() (*string, bool)`

GetPaymentSessionOk returns a tuple with the PaymentSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSession

`func (o *PaymentSetupResponse) SetPaymentSession(v string)`

SetPaymentSession sets PaymentSession field to given value.

### HasPaymentSession

`func (o *PaymentSetupResponse) HasPaymentSession() bool`

HasPaymentSession returns a boolean if a field has been set.

### GetRecurringDetails

`func (o *PaymentSetupResponse) GetRecurringDetails() []RecurringDetail`

GetRecurringDetails returns the RecurringDetails field if non-nil, zero value otherwise.

### GetRecurringDetailsOk

`func (o *PaymentSetupResponse) GetRecurringDetailsOk() (*[]RecurringDetail, bool)`

GetRecurringDetailsOk returns a tuple with the RecurringDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetails

`func (o *PaymentSetupResponse) SetRecurringDetails(v []RecurringDetail)`

SetRecurringDetails sets RecurringDetails field to given value.

### HasRecurringDetails

`func (o *PaymentSetupResponse) HasRecurringDetails() bool`

HasRecurringDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


