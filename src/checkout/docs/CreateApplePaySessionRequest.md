# CreateApplePaySessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | This is the name that your shoppers will see in the Apple Pay interface.  The value returned as &#x60;configuration.merchantName&#x60; field from the [&#x60;/paymentMethods&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/paymentMethods) response. | 
**DomainName** | **string** | The domain name you provided when you added Apple Pay in your Customer Area.  This must match the &#x60;window.location.hostname&#x60; of the web shop. | 
**MerchantIdentifier** | **string** | Your merchant identifier registered with Apple Pay.  Use the value of the &#x60;configuration.merchantId&#x60; field from the [&#x60;/paymentMethods&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/paymentMethods) response. | 

## Methods

### NewCreateApplePaySessionRequest

`func NewCreateApplePaySessionRequest(displayName string, domainName string, merchantIdentifier string, ) *CreateApplePaySessionRequest`

NewCreateApplePaySessionRequest instantiates a new CreateApplePaySessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplePaySessionRequestWithDefaults

`func NewCreateApplePaySessionRequestWithDefaults() *CreateApplePaySessionRequest`

NewCreateApplePaySessionRequestWithDefaults instantiates a new CreateApplePaySessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CreateApplePaySessionRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateApplePaySessionRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateApplePaySessionRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDomainName

`func (o *CreateApplePaySessionRequest) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CreateApplePaySessionRequest) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CreateApplePaySessionRequest) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetMerchantIdentifier

`func (o *CreateApplePaySessionRequest) GetMerchantIdentifier() string`

GetMerchantIdentifier returns the MerchantIdentifier field if non-nil, zero value otherwise.

### GetMerchantIdentifierOk

`func (o *CreateApplePaySessionRequest) GetMerchantIdentifierOk() (*string, bool)`

GetMerchantIdentifierOk returns a tuple with the MerchantIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantIdentifier

`func (o *CreateApplePaySessionRequest) SetMerchantIdentifier(v string)`

SetMerchantIdentifier sets MerchantIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


