# TestOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | Pointer to **string** | Unique identifier of the merchant account that the notification is about. | [optional] 
**Output** | Pointer to **string** | The response your server returned for the test webhook.  Your server must respond with **[accepted]** for the test webhook to be successful (&#x60;data.status&#x60;: **success**). Find out more about [accepting notifications](https://docs.adyen.com/development-resources/webhooks#accept-notifications)  You can use the value of this field together with the [&#x60;responseCode&#x60;](https://docs.adyen.com/api-explorer/#/ManagementService/v1/post/merchants/{merchantId}/webhooks/{id}/test__resParam_data-responseCode) value to troubleshoot unsuccessful test webhooks. | [optional] 
**RequestSent** | Pointer to **string** | The [body of the notification webhook](https://docs.adyen.com/development-resources/webhooks/understand-notifications#notification-structure) that was sent to your server. | [optional] 
**ResponseCode** | Pointer to **string** | The HTTP response code for your server&#39;s response to the test webhook.  You can use the value of this field together with the the [&#x60;output&#x60;](https://docs.adyen.com/api-explorer/#/ManagementService/v1/post/merchants/{merchantId}/webhooks/{id}/test__resParam_data-output) field value to troubleshoot failed test webhooks. | [optional] 
**ResponseTime** | Pointer to **string** | The time between sending the test webhook and receiving the response from your server. You can use it as an indication of how long your server takes to process a webhook notification. Measured in milliseconds, for example **304 ms**. | [optional] 
**Status** | **string** | The status of the test request. Possible values are: * **success**, if &#x60;data.output&#x60;: **[accepted]** and &#x60;data.responseCode&#x60;: **200**. * **failed**, in all other cases.  You can use the value of the [&#x60;output&#x60;](https://docs.adyen.com/api-explorer/#/ManagementService/v1/post/merchants/{merchantId}/webhooks/{id}/test__resParam_data-output) field together with the [&#x60;responseCode&#x60;](https://docs.adyen.com/api-explorer/#/ManagementService/v1/post/merchants/{merchantId}/webhooks/{id}/test__resParam_data-responseCode) value to troubleshoot failed test webhooks. | 

## Methods

### NewTestOutput

`func NewTestOutput(status string, ) *TestOutput`

NewTestOutput instantiates a new TestOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestOutputWithDefaults

`func NewTestOutputWithDefaults() *TestOutput`

NewTestOutputWithDefaults instantiates a new TestOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *TestOutput) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *TestOutput) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *TestOutput) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *TestOutput) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetOutput

`func (o *TestOutput) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *TestOutput) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *TestOutput) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *TestOutput) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetRequestSent

`func (o *TestOutput) GetRequestSent() string`

GetRequestSent returns the RequestSent field if non-nil, zero value otherwise.

### GetRequestSentOk

`func (o *TestOutput) GetRequestSentOk() (*string, bool)`

GetRequestSentOk returns a tuple with the RequestSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSent

`func (o *TestOutput) SetRequestSent(v string)`

SetRequestSent sets RequestSent field to given value.

### HasRequestSent

`func (o *TestOutput) HasRequestSent() bool`

HasRequestSent returns a boolean if a field has been set.

### GetResponseCode

`func (o *TestOutput) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *TestOutput) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *TestOutput) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *TestOutput) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### GetResponseTime

`func (o *TestOutput) GetResponseTime() string`

GetResponseTime returns the ResponseTime field if non-nil, zero value otherwise.

### GetResponseTimeOk

`func (o *TestOutput) GetResponseTimeOk() (*string, bool)`

GetResponseTimeOk returns a tuple with the ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTime

`func (o *TestOutput) SetResponseTime(v string)`

SetResponseTime sets ResponseTime field to given value.

### HasResponseTime

`func (o *TestOutput) HasResponseTime() bool`

HasResponseTime returns a boolean if a field has been set.

### GetStatus

`func (o *TestOutput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestOutput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestOutput) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


