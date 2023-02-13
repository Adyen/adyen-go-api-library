# AcceptTermsOfServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedBy** | Pointer to **string** | The unique identifier of the user accepting the Terms of Service. | [optional] 
**IpAddress** | Pointer to **string** | The IP address of the user accepting the Terms of Service. | [optional] 

## Methods

### NewAcceptTermsOfServiceRequest

`func NewAcceptTermsOfServiceRequest() *AcceptTermsOfServiceRequest`

NewAcceptTermsOfServiceRequest instantiates a new AcceptTermsOfServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptTermsOfServiceRequestWithDefaults

`func NewAcceptTermsOfServiceRequestWithDefaults() *AcceptTermsOfServiceRequest`

NewAcceptTermsOfServiceRequestWithDefaults instantiates a new AcceptTermsOfServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedBy

`func (o *AcceptTermsOfServiceRequest) GetAcceptedBy() string`

GetAcceptedBy returns the AcceptedBy field if non-nil, zero value otherwise.

### GetAcceptedByOk

`func (o *AcceptTermsOfServiceRequest) GetAcceptedByOk() (*string, bool)`

GetAcceptedByOk returns a tuple with the AcceptedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedBy

`func (o *AcceptTermsOfServiceRequest) SetAcceptedBy(v string)`

SetAcceptedBy sets AcceptedBy field to given value.

### HasAcceptedBy

`func (o *AcceptTermsOfServiceRequest) HasAcceptedBy() bool`

HasAcceptedBy returns a boolean if a field has been set.

### GetIpAddress

`func (o *AcceptTermsOfServiceRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *AcceptTermsOfServiceRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *AcceptTermsOfServiceRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *AcceptTermsOfServiceRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


