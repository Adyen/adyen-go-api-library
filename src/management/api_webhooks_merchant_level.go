/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"context"
	_context "context"
	_nethttp "net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// WebhooksMerchantLevelApi WebhooksMerchantLevelApi service
type WebhooksMerchantLevelApi common.Service

type WebhooksMerchantLevelApiGenerateHmacKeyConfig struct {
	ctx        context.Context
	merchantId string
	webhookId  string
}

/*
GenerateHmacKey Generate an HMAC key

Returns an [HMAC key](https://en.wikipedia.org/wiki/HMAC) for the webhook identified in the path. This key allows you to check the integrity and the origin of the notifications you receive.By creating an HMAC key, you start receiving [HMAC-signed notifications](https://docs.adyen.com/development-resources/webhooks/verify-hmac-signatures#enable-hmac-signatures) from Adyen. Find out more about how to [verify HMAC signatures](https://docs.adyen.com/development-resources/webhooks/verify-hmac-signatures).

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Webhooks read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @param webhookId
 @return WebhooksMerchantLevelApiGenerateHmacKeyConfig
*/
func (a *WebhooksMerchantLevelApi) GenerateHmacKeyConfig(ctx context.Context, merchantId string, webhookId string) WebhooksMerchantLevelApiGenerateHmacKeyConfig {
	return WebhooksMerchantLevelApiGenerateHmacKeyConfig{
		ctx:        ctx,
		merchantId: merchantId,
		webhookId:  webhookId,
	}
}

/*
Generate an HMAC key
Returns an [HMAC key](https://en.wikipedia.org/wiki/HMAC) for the webhook identified in the path. This key allows you to check the integrity and the origin of the notifications you receive.By creating an HMAC key, you start receiving [HMAC-signed notifications](https://docs.adyen.com/development-resources/webhooks/verify-hmac-signatures#enable-hmac-signatures) from Adyen. Find out more about how to [verify HMAC signatures](https://docs.adyen.com/development-resources/webhooks/verify-hmac-signatures).  To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Webhooks read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param webhookId
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return GenerateHmacKeyResponse
*/

func (a *WebhooksMerchantLevelApi) GenerateHmacKey(r WebhooksMerchantLevelApiGenerateHmacKeyConfig) (GenerateHmacKeyResponse, *_nethttp.Response, error) {
	res := &GenerateHmacKeyResponse{}
	path := "/merchants/{merchantId}/webhooks/{webhookId}/generateHmac"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"webhookId"+"}", url.PathEscape(common.ParameterValueToString(r.webhookId, "webhookId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPost, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type WebhooksMerchantLevelApiGetWebhookConfig struct {
	ctx        context.Context
	merchantId string
	webhookId  string
}

/*
GetWebhook Get a webhook

Returns the configuration for the webhook identified in the path.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Webhooks read
* Management API—Webhooks read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @param webhookId Unique identifier of the webhook configuration.
 @return WebhooksMerchantLevelApiGetWebhookConfig
*/
func (a *WebhooksMerchantLevelApi) GetWebhookConfig(ctx context.Context, merchantId string, webhookId string) WebhooksMerchantLevelApiGetWebhookConfig {
	return WebhooksMerchantLevelApiGetWebhookConfig{
		ctx:        ctx,
		merchantId: merchantId,
		webhookId:  webhookId,
	}
}

/*
Get a webhook
Returns the configuration for the webhook identified in the path.  To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Webhooks read * Management API—Webhooks read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param webhookId Unique identifier of the webhook configuration.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return Webhook
*/

func (a *WebhooksMerchantLevelApi) GetWebhook(r WebhooksMerchantLevelApiGetWebhookConfig) (Webhook, *_nethttp.Response, error) {
	res := &Webhook{}
	path := "/merchants/{merchantId}/webhooks/{webhookId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"webhookId"+"}", url.PathEscape(common.ParameterValueToString(r.webhookId, "webhookId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type WebhooksMerchantLevelApiListAllWebhooksConfig struct {
	ctx        context.Context
	merchantId string
	pageNumber *int32
	pageSize   *int32
}

// The number of the page to fetch.
func (r WebhooksMerchantLevelApiListAllWebhooksConfig) PageNumber(pageNumber int32) WebhooksMerchantLevelApiListAllWebhooksConfig {
	r.pageNumber = &pageNumber
	return r
}

// The number of items to have on a page, maximum 100. The default is 10 items on a page.
func (r WebhooksMerchantLevelApiListAllWebhooksConfig) PageSize(pageSize int32) WebhooksMerchantLevelApiListAllWebhooksConfig {
	r.pageSize = &pageSize
	return r
}

/*
ListAllWebhooks List all webhooks

Lists all webhook configurations for the merchant account.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Webhooks read
* Management API—Webhooks read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @return WebhooksMerchantLevelApiListAllWebhooksConfig
*/
func (a *WebhooksMerchantLevelApi) ListAllWebhooksConfig(ctx context.Context, merchantId string) WebhooksMerchantLevelApiListAllWebhooksConfig {
	return WebhooksMerchantLevelApiListAllWebhooksConfig{
		ctx:        ctx,
		merchantId: merchantId,
	}
}

/*
List all webhooks
Lists all webhook configurations for the merchant account.  To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Webhooks read * Management API—Webhooks read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ListWebhooksResponse
*/

func (a *WebhooksMerchantLevelApi) ListAllWebhooks(r WebhooksMerchantLevelApiListAllWebhooksConfig) (ListWebhooksResponse, *_nethttp.Response, error) {
	res := &ListWebhooksResponse{}
	path := "/merchants/{merchantId}/webhooks"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryString := url.Values{}
	if r.pageNumber != nil {
		common.ParameterAddToQuery(queryString, "pageNumber", r.pageNumber, "")
	}
	if r.pageSize != nil {
		common.ParameterAddToQuery(queryString, "pageSize", r.pageSize, "")
	}
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path+"?"+queryString.Encode(), []_context.Context{r.ctx})
	return *res, httpRes, err
}

type WebhooksMerchantLevelApiRemoveWebhookConfig struct {
	ctx        context.Context
	merchantId string
	webhookId  string
}

/*
RemoveWebhook Remove a webhook

Remove the configuration for the webhook identified in the path.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Webhooks read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @param webhookId Unique identifier of the webhook configuration.
 @return WebhooksMerchantLevelApiRemoveWebhookConfig
*/
func (a *WebhooksMerchantLevelApi) RemoveWebhookConfig(ctx context.Context, merchantId string, webhookId string) WebhooksMerchantLevelApiRemoveWebhookConfig {
	return WebhooksMerchantLevelApiRemoveWebhookConfig{
		ctx:        ctx,
		merchantId: merchantId,
		webhookId:  webhookId,
	}
}

/*
Remove a webhook
Remove the configuration for the webhook identified in the path.  To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Webhooks read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param webhookId Unique identifier of the webhook configuration.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
*/

func (a *WebhooksMerchantLevelApi) RemoveWebhook(r WebhooksMerchantLevelApiRemoveWebhookConfig) (*_nethttp.Response, error) {
	var res interface{}
	path := "/merchants/{merchantId}/webhooks/{webhookId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"webhookId"+"}", url.PathEscape(common.ParameterValueToString(r.webhookId, "webhookId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodDelete, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return httpRes, err
}

type WebhooksMerchantLevelApiSetUpWebhookConfig struct {
	ctx                          context.Context
	merchantId                   string
	createMerchantWebhookRequest *CreateMerchantWebhookRequest
}

func (r WebhooksMerchantLevelApiSetUpWebhookConfig) CreateMerchantWebhookRequest(createMerchantWebhookRequest CreateMerchantWebhookRequest) WebhooksMerchantLevelApiSetUpWebhookConfig {
	r.createMerchantWebhookRequest = &createMerchantWebhookRequest
	return r
}

/*
SetUpWebhook Set up a webhook

Subscribe to receive webhook notifications about events related to your merchant account. You can add basic authentication to make sure the data is secure.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Webhooks read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @return WebhooksMerchantLevelApiSetUpWebhookConfig
*/
func (a *WebhooksMerchantLevelApi) SetUpWebhookConfig(ctx context.Context, merchantId string) WebhooksMerchantLevelApiSetUpWebhookConfig {
	return WebhooksMerchantLevelApiSetUpWebhookConfig{
		ctx:        ctx,
		merchantId: merchantId,
	}
}

/*
Set up a webhook
Subscribe to receive webhook notifications about events related to your merchant account. You can add basic authentication to make sure the data is secure.  To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Webhooks read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param req CreateMerchantWebhookRequest - reference of CreateMerchantWebhookRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return Webhook
*/

func (a *WebhooksMerchantLevelApi) SetUpWebhook(r WebhooksMerchantLevelApiSetUpWebhookConfig) (Webhook, *_nethttp.Response, error) {
	res := &Webhook{}
	path := "/merchants/{merchantId}/webhooks"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPost, r.createMerchantWebhookRequest, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type WebhooksMerchantLevelApiTestWebhookConfig struct {
	ctx                context.Context
	merchantId         string
	webhookId          string
	testWebhookRequest *TestWebhookRequest
}

func (r WebhooksMerchantLevelApiTestWebhookConfig) TestWebhookRequest(testWebhookRequest TestWebhookRequest) WebhooksMerchantLevelApiTestWebhookConfig {
	r.testWebhookRequest = &testWebhookRequest
	return r
}

/*
TestWebhook Test a webhook

Sends sample notifications to test if the webhook is set up correctly.

We send four test notifications for each event code you choose. They cover success and failure scenarios for the hard-coded currencies EUR and GBP, regardless of the currencies configured in the merchant accounts. For custom notifications, we only send the specified custom notification.

The response describes the result of the test. The `status` field tells you if the test was successful or not. You can use the other fields to troubleshoot unsuccessful tests.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Webhooks read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @param webhookId Unique identifier of the webhook configuration.
 @return WebhooksMerchantLevelApiTestWebhookConfig
*/
func (a *WebhooksMerchantLevelApi) TestWebhookConfig(ctx context.Context, merchantId string, webhookId string) WebhooksMerchantLevelApiTestWebhookConfig {
	return WebhooksMerchantLevelApiTestWebhookConfig{
		ctx:        ctx,
		merchantId: merchantId,
		webhookId:  webhookId,
	}
}

/*
Test a webhook
Sends sample notifications to test if the webhook is set up correctly.  We send four test notifications for each event code you choose. They cover success and failure scenarios for the hard-coded currencies EUR and GBP, regardless of the currencies configured in the merchant accounts. For custom notifications, we only send the specified custom notification.  The response describes the result of the test. The &#x60;status&#x60; field tells you if the test was successful or not. You can use the other fields to troubleshoot unsuccessful tests.  To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Webhooks read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param webhookId Unique identifier of the webhook configuration.
 * @param req TestWebhookRequest - reference of TestWebhookRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return TestWebhookResponse
*/

func (a *WebhooksMerchantLevelApi) TestWebhook(r WebhooksMerchantLevelApiTestWebhookConfig) (TestWebhookResponse, *_nethttp.Response, error) {
	res := &TestWebhookResponse{}
	path := "/merchants/{merchantId}/webhooks/{webhookId}/test"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"webhookId"+"}", url.PathEscape(common.ParameterValueToString(r.webhookId, "webhookId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPost, r.testWebhookRequest, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type WebhooksMerchantLevelApiUpdateWebhookConfig struct {
	ctx                          context.Context
	merchantId                   string
	webhookId                    string
	updateMerchantWebhookRequest *UpdateMerchantWebhookRequest
}

func (r WebhooksMerchantLevelApiUpdateWebhookConfig) UpdateMerchantWebhookRequest(updateMerchantWebhookRequest UpdateMerchantWebhookRequest) WebhooksMerchantLevelApiUpdateWebhookConfig {
	r.updateMerchantWebhookRequest = &updateMerchantWebhookRequest
	return r
}

/*
UpdateWebhook Update a webhook

Make changes to the configuration of the webhook identified in the path. The request contains the new values you want to have in the webhook configuration. The response contains the full configuration for the webhook, which includes the new values from the request.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Webhooks read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @param webhookId Unique identifier of the webhook configuration.
 @return WebhooksMerchantLevelApiUpdateWebhookConfig
*/
func (a *WebhooksMerchantLevelApi) UpdateWebhookConfig(ctx context.Context, merchantId string, webhookId string) WebhooksMerchantLevelApiUpdateWebhookConfig {
	return WebhooksMerchantLevelApiUpdateWebhookConfig{
		ctx:        ctx,
		merchantId: merchantId,
		webhookId:  webhookId,
	}
}

/*
Update a webhook
Make changes to the configuration of the webhook identified in the path. The request contains the new values you want to have in the webhook configuration. The response contains the full configuration for the webhook, which includes the new values from the request.  To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Webhooks read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param webhookId Unique identifier of the webhook configuration.
 * @param req UpdateMerchantWebhookRequest - reference of UpdateMerchantWebhookRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return Webhook
*/

func (a *WebhooksMerchantLevelApi) UpdateWebhook(r WebhooksMerchantLevelApiUpdateWebhookConfig) (Webhook, *_nethttp.Response, error) {
	res := &Webhook{}
	path := "/merchants/{merchantId}/webhooks/{webhookId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"webhookId"+"}", url.PathEscape(common.ParameterValueToString(r.webhookId, "webhookId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPatch, r.updateMerchantWebhookRequest, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}