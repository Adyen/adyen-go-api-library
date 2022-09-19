# PayPalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectCapture** | Pointer to **bool** | Indicates if direct (immediate) capture for PayPal is enabled. If set to **true**, this setting overrides the [capture](https://docs.adyen.com/online-payments/capture) settings of your merchant account. Default value: **true**. | [optional] 
**DirectSettlement** | **bool** | Must be set to **true** to confirm that the settlement to your bank account is performed directly by PayPal. Default value: **null**. | 
**PayerId** | **string** | PayPal Merchant ID. Character length and limitations: 13 single-byte alphanumeric characters. | 
**Subject** | **string** | Your business email address. | 

## Methods

### NewPayPalInfo

`func NewPayPalInfo(directSettlement bool, payerId string, subject string, ) *PayPalInfo`

NewPayPalInfo instantiates a new PayPalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayPalInfoWithDefaults

`func NewPayPalInfoWithDefaults() *PayPalInfo`

NewPayPalInfoWithDefaults instantiates a new PayPalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectCapture

`func (o *PayPalInfo) GetDirectCapture() bool`

GetDirectCapture returns the DirectCapture field if non-nil, zero value otherwise.

### GetDirectCaptureOk

`func (o *PayPalInfo) GetDirectCaptureOk() (*bool, bool)`

GetDirectCaptureOk returns a tuple with the DirectCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectCapture

`func (o *PayPalInfo) SetDirectCapture(v bool)`

SetDirectCapture sets DirectCapture field to given value.

### HasDirectCapture

`func (o *PayPalInfo) HasDirectCapture() bool`

HasDirectCapture returns a boolean if a field has been set.

### GetDirectSettlement

`func (o *PayPalInfo) GetDirectSettlement() bool`

GetDirectSettlement returns the DirectSettlement field if non-nil, zero value otherwise.

### GetDirectSettlementOk

`func (o *PayPalInfo) GetDirectSettlementOk() (*bool, bool)`

GetDirectSettlementOk returns a tuple with the DirectSettlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectSettlement

`func (o *PayPalInfo) SetDirectSettlement(v bool)`

SetDirectSettlement sets DirectSettlement field to given value.


### GetPayerId

`func (o *PayPalInfo) GetPayerId() string`

GetPayerId returns the PayerId field if non-nil, zero value otherwise.

### GetPayerIdOk

`func (o *PayPalInfo) GetPayerIdOk() (*string, bool)`

GetPayerIdOk returns a tuple with the PayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerId

`func (o *PayPalInfo) SetPayerId(v string)`

SetPayerId sets PayerId field to given value.


### GetSubject

`func (o *PayPalInfo) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PayPalInfo) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PayPalInfo) SetSubject(v string)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


