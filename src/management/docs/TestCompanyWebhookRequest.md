# TestCompanyWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantIds** | Pointer to **[]string** | List of &#x60;merchantId&#x60; values for which test webhooks will be sent. The list can have a maximum of 20 &#x60;merchantId&#x60; values.  If not specified, we send sample notifications to all the merchant accounts that the webhook is configured for. If this is more than 20 merchant accounts, use this list to specify a subset of the merchant accounts for which to send test notifications. | [optional] 
**Notification** | Pointer to [**CustomNotification**](CustomNotification.md) |  | [optional] 
**Types** | Pointer to **[]string** | List of event codes for which to send test notifications. Only the webhook types below are supported.   Possible values if webhook &#x60;type&#x60;: **standard**:  * **AUTHORISATION** * **CHARGEBACK_REVERSED** * **ORDER_CLOSED** * **ORDER_OPENED** * **PAIDOUT_REVERSED** * **PAYOUT_THIRDPARTY** * **REFUNDED_REVERSED** * **REFUND_WITH_DATA** * **REPORT_AVAILABLE** * **CUSTOM** - set your custom notification fields in the [&#x60;notification&#x60;](https://docs.adyen.com/api-explorer/#/ManagementService/v1/post/companies/{companyId}/webhooks/{webhookId}/test__reqParam_notification) object.  Possible values if webhook &#x60;type&#x60;: **banktransfer-notification**:  * **PENDING**  Possible values if webhook &#x60;type&#x60;: **report-notification**:  * **REPORT_AVAILABLE**  Possible values if webhook &#x60;type&#x60;: **ideal-notification**:  * **AUTHORISATION**  Possible values if webhook &#x60;type&#x60;: **pending-notification**:  * **PENDING**  | [optional] 

## Methods

### NewTestCompanyWebhookRequest

`func NewTestCompanyWebhookRequest() *TestCompanyWebhookRequest`

NewTestCompanyWebhookRequest instantiates a new TestCompanyWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestCompanyWebhookRequestWithDefaults

`func NewTestCompanyWebhookRequestWithDefaults() *TestCompanyWebhookRequest`

NewTestCompanyWebhookRequestWithDefaults instantiates a new TestCompanyWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantIds

`func (o *TestCompanyWebhookRequest) GetMerchantIds() []string`

GetMerchantIds returns the MerchantIds field if non-nil, zero value otherwise.

### GetMerchantIdsOk

`func (o *TestCompanyWebhookRequest) GetMerchantIdsOk() (*[]string, bool)`

GetMerchantIdsOk returns a tuple with the MerchantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantIds

`func (o *TestCompanyWebhookRequest) SetMerchantIds(v []string)`

SetMerchantIds sets MerchantIds field to given value.

### HasMerchantIds

`func (o *TestCompanyWebhookRequest) HasMerchantIds() bool`

HasMerchantIds returns a boolean if a field has been set.

### GetNotification

`func (o *TestCompanyWebhookRequest) GetNotification() CustomNotification`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *TestCompanyWebhookRequest) GetNotificationOk() (*CustomNotification, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *TestCompanyWebhookRequest) SetNotification(v CustomNotification)`

SetNotification sets Notification field to given value.

### HasNotification

`func (o *TestCompanyWebhookRequest) HasNotification() bool`

HasNotification returns a boolean if a field has been set.

### GetTypes

`func (o *TestCompanyWebhookRequest) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *TestCompanyWebhookRequest) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *TestCompanyWebhookRequest) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *TestCompanyWebhookRequest) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


