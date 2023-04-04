# AmazonPayDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmazonPayToken** | Pointer to **string** | This is the &#x60;amazonPayToken&#x60; that you obtained from the [Get Checkout Session](https://amazon-pay-acquirer-guide.s3-eu-west-1.amazonaws.com/v1/amazon-pay-api-v2/checkout-session.html#get-checkout-session) response. | [optional] 
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**Type** | Pointer to **string** | **amazonpay** | [optional] [default to "amazonpay"]

## Methods

### NewAmazonPayDetails

`func NewAmazonPayDetails() *AmazonPayDetails`

NewAmazonPayDetails instantiates a new AmazonPayDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonPayDetailsWithDefaults

`func NewAmazonPayDetailsWithDefaults() *AmazonPayDetails`

NewAmazonPayDetailsWithDefaults instantiates a new AmazonPayDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmazonPayToken

`func (o *AmazonPayDetails) GetAmazonPayToken() string`

GetAmazonPayToken returns the AmazonPayToken field if non-nil, zero value otherwise.

### GetAmazonPayTokenOk

`func (o *AmazonPayDetails) GetAmazonPayTokenOk() (*string, bool)`

GetAmazonPayTokenOk returns a tuple with the AmazonPayToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonPayToken

`func (o *AmazonPayDetails) SetAmazonPayToken(v string)`

SetAmazonPayToken sets AmazonPayToken field to given value.

### HasAmazonPayToken

`func (o *AmazonPayDetails) HasAmazonPayToken() bool`

HasAmazonPayToken returns a boolean if a field has been set.

### GetCheckoutAttemptId

`func (o *AmazonPayDetails) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *AmazonPayDetails) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *AmazonPayDetails) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *AmazonPayDetails) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetType

`func (o *AmazonPayDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AmazonPayDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AmazonPayDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AmazonPayDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


