/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// WebhooksCompanyLevelApi service
type WebhooksCompanyLevelApi common.Service

// All parameters accepted by WebhooksCompanyLevelApi.GenerateHmacKey
type WebhooksCompanyLevelApiGenerateHmacKeyInput struct {
	companyId string
	webhookId string
}

/*
Prepare a request for GenerateHmacKey
@param companyId The unique identifier of the company account.@param webhookId Unique identifier of the webhook configuration.
@return WebhooksCompanyLevelApiGenerateHmacKeyInput
*/
func (a *WebhooksCompanyLevelApi) GenerateHmacKeyInput(companyId string, webhookId string) WebhooksCompanyLevelApiGenerateHmacKeyInput {
	return WebhooksCompanyLevelApiGenerateHmacKeyInput{
		companyId: companyId,
		webhookId: webhookId,
	}
}

/*
GenerateHmacKey Generate an HMAC key

Returns an [HMAC key](https://en.wikipedia.org/wiki/HMAC) for the webhook identified in the path. This key allows you to check the integrity and the origin of the notifications you receive.By creating an HMAC key, you start receiving [HMAC-signed notifications](https://docs.adyen.com/development-resources/webhooks/verify-hmac-signatures#enable-hmac-signatures) from Adyen. Find out more about how to [verify HMAC signatures](https://docs.adyen.com/development-resources/webhooks/verify-hmac-signatures).

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Webhooks read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r WebhooksCompanyLevelApiGenerateHmacKeyInput - Request parameters, see GenerateHmacKeyInput
@return GenerateHmacKeyResponse, *http.Response, error
*/
func (a *WebhooksCompanyLevelApi) GenerateHmacKey(ctx context.Context, r WebhooksCompanyLevelApiGenerateHmacKeyInput) (GenerateHmacKeyResponse, *http.Response, error) {
	res := &GenerateHmacKeyResponse{}
	path := "/companies/{companyId}/webhooks/{webhookId}/generateHmac"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"webhookId"+"}", url.PathEscape(common.ParameterValueToString(r.webhookId, "webhookId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by WebhooksCompanyLevelApi.GetWebhook
type WebhooksCompanyLevelApiGetWebhookInput struct {
	companyId string
	webhookId string
}

/*
Prepare a request for GetWebhook
@param companyId Unique identifier of the [company account](https://docs.adyen.com/account/account-structure#company-account).@param webhookId Unique identifier of the webhook configuration.
@return WebhooksCompanyLevelApiGetWebhookInput
*/
func (a *WebhooksCompanyLevelApi) GetWebhookInput(companyId string, webhookId string) WebhooksCompanyLevelApiGetWebhookInput {
	return WebhooksCompanyLevelApiGetWebhookInput{
		companyId: companyId,
		webhookId: webhookId,
	}
}

/*
GetWebhook Get a webhook

Returns the configuration for the webhook identified in the path.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Webhooks read
* Management API—Webhooks read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r WebhooksCompanyLevelApiGetWebhookInput - Request parameters, see GetWebhookInput
@return Webhook, *http.Response, error
*/
func (a *WebhooksCompanyLevelApi) GetWebhook(ctx context.Context, r WebhooksCompanyLevelApiGetWebhookInput) (Webhook, *http.Response, error) {
	res := &Webhook{}
	path := "/companies/{companyId}/webhooks/{webhookId}"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"webhookId"+"}", url.PathEscape(common.ParameterValueToString(r.webhookId, "webhookId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by WebhooksCompanyLevelApi.ListAllWebhooks
type WebhooksCompanyLevelApiListAllWebhooksInput struct {
	companyId  string
	pageNumber *int32
	pageSize   *int32
}

// The number of the page to fetch.
func (r WebhooksCompanyLevelApiListAllWebhooksInput) PageNumber(pageNumber int32) WebhooksCompanyLevelApiListAllWebhooksInput {
	r.pageNumber = &pageNumber
	return r
}

// The number of items to have on a page, maximum 100. The default is 10 items on a page.
func (r WebhooksCompanyLevelApiListAllWebhooksInput) PageSize(pageSize int32) WebhooksCompanyLevelApiListAllWebhooksInput {
	r.pageSize = &pageSize
	return r
}

/*
Prepare a request for ListAllWebhooks
@param companyId Unique identifier of the [company account](https://docs.adyen.com/account/account-structure#company-account).
@return WebhooksCompanyLevelApiListAllWebhooksInput
*/
func (a *WebhooksCompanyLevelApi) ListAllWebhooksInput(companyId string) WebhooksCompanyLevelApiListAllWebhooksInput {
	return WebhooksCompanyLevelApiListAllWebhooksInput{
		companyId: companyId,
	}
}

/*
ListAllWebhooks List all webhooks

Lists all webhook configurations for the company account.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Webhooks read
* Management API—Webhooks read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r WebhooksCompanyLevelApiListAllWebhooksInput - Request parameters, see ListAllWebhooksInput
@return ListWebhooksResponse, *http.Response, error
*/
func (a *WebhooksCompanyLevelApi) ListAllWebhooks(ctx context.Context, r WebhooksCompanyLevelApiListAllWebhooksInput) (ListWebhooksResponse, *http.Response, error) {
	res := &ListWebhooksResponse{}
	path := "/companies/{companyId}/webhooks"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.pageNumber != nil {
		common.ParameterAddToQuery(queryParams, "pageNumber", r.pageNumber, "")
	}
	if r.pageSize != nil {
		common.ParameterAddToQuery(queryParams, "pageSize", r.pageSize, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by WebhooksCompanyLevelApi.RemoveWebhook
type WebhooksCompanyLevelApiRemoveWebhookInput struct {
	companyId string
	webhookId string
}

/*
Prepare a request for RemoveWebhook
@param companyId The unique identifier of the company account.@param webhookId Unique identifier of the webhook configuration.
@return WebhooksCompanyLevelApiRemoveWebhookInput
*/
func (a *WebhooksCompanyLevelApi) RemoveWebhookInput(companyId string, webhookId string) WebhooksCompanyLevelApiRemoveWebhookInput {
	return WebhooksCompanyLevelApiRemoveWebhookInput{
		companyId: companyId,
		webhookId: webhookId,
	}
}

/*
RemoveWebhook Remove a webhook

Remove the configuration for the webhook identified in the path.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Webhooks read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r WebhooksCompanyLevelApiRemoveWebhookInput - Request parameters, see RemoveWebhookInput
@return *http.Response, error
*/
func (a *WebhooksCompanyLevelApi) RemoveWebhook(ctx context.Context, r WebhooksCompanyLevelApiRemoveWebhookInput) (*http.Response, error) {
	var res interface{}
	path := "/companies/{companyId}/webhooks/{webhookId}"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"webhookId"+"}", url.PathEscape(common.ParameterValueToString(r.webhookId, "webhookId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodDelete,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}

	return httpRes, err
}

// All parameters accepted by WebhooksCompanyLevelApi.SetUpWebhook
type WebhooksCompanyLevelApiSetUpWebhookInput struct {
	companyId                   string
	createCompanyWebhookRequest *CreateCompanyWebhookRequest
}

func (r WebhooksCompanyLevelApiSetUpWebhookInput) CreateCompanyWebhookRequest(createCompanyWebhookRequest CreateCompanyWebhookRequest) WebhooksCompanyLevelApiSetUpWebhookInput {
	r.createCompanyWebhookRequest = &createCompanyWebhookRequest
	return r
}

/*
Prepare a request for SetUpWebhook
@param companyId Unique identifier of the [company account](https://docs.adyen.com/account/account-structure#company-account).
@return WebhooksCompanyLevelApiSetUpWebhookInput
*/
func (a *WebhooksCompanyLevelApi) SetUpWebhookInput(companyId string) WebhooksCompanyLevelApiSetUpWebhookInput {
	return WebhooksCompanyLevelApiSetUpWebhookInput{
		companyId: companyId,
	}
}

/*
SetUpWebhook Set up a webhook

Subscribe to receive webhook notifications about events related to your company account. You can add basic authentication to make sure the data is secure.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Webhooks read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r WebhooksCompanyLevelApiSetUpWebhookInput - Request parameters, see SetUpWebhookInput
@return Webhook, *http.Response, error
*/
func (a *WebhooksCompanyLevelApi) SetUpWebhook(ctx context.Context, r WebhooksCompanyLevelApiSetUpWebhookInput) (Webhook, *http.Response, error) {
	res := &Webhook{}
	path := "/companies/{companyId}/webhooks"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.createCompanyWebhookRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by WebhooksCompanyLevelApi.TestWebhook
type WebhooksCompanyLevelApiTestWebhookInput struct {
	companyId                 string
	webhookId                 string
	testCompanyWebhookRequest *TestCompanyWebhookRequest
}

func (r WebhooksCompanyLevelApiTestWebhookInput) TestCompanyWebhookRequest(testCompanyWebhookRequest TestCompanyWebhookRequest) WebhooksCompanyLevelApiTestWebhookInput {
	r.testCompanyWebhookRequest = &testCompanyWebhookRequest
	return r
}

/*
Prepare a request for TestWebhook
@param companyId The unique identifier of the company account.@param webhookId Unique identifier of the webhook configuration.
@return WebhooksCompanyLevelApiTestWebhookInput
*/
func (a *WebhooksCompanyLevelApi) TestWebhookInput(companyId string, webhookId string) WebhooksCompanyLevelApiTestWebhookInput {
	return WebhooksCompanyLevelApiTestWebhookInput{
		companyId: companyId,
		webhookId: webhookId,
	}
}

/*
TestWebhook Test a webhook

Sends sample notifications to test if the webhook is set up correctly.

We send sample notifications for maximum 20 of the merchant accounts that the webhook is configured for. If the webhook is configured for more than 20 merchant accounts, use the `merchantIds` array to specify a subset of the merchant accounts for which to send test notifications.

We send four test notifications for each event code you choose. They cover success and failure scenarios for the hard-coded currencies EUR and GBP, regardless of the currencies configured in the merchant accounts. For custom notifications, we only send the specified custom notification.

The response describes the result of the test. The `status` field tells you if the test was successful or not. You can use the other response fields to troubleshoot unsuccessful tests.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Webhooks read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r WebhooksCompanyLevelApiTestWebhookInput - Request parameters, see TestWebhookInput
@return TestWebhookResponse, *http.Response, error
*/
func (a *WebhooksCompanyLevelApi) TestWebhook(ctx context.Context, r WebhooksCompanyLevelApiTestWebhookInput) (TestWebhookResponse, *http.Response, error) {
	res := &TestWebhookResponse{}
	path := "/companies/{companyId}/webhooks/{webhookId}/test"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"webhookId"+"}", url.PathEscape(common.ParameterValueToString(r.webhookId, "webhookId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.testCompanyWebhookRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by WebhooksCompanyLevelApi.UpdateWebhook
type WebhooksCompanyLevelApiUpdateWebhookInput struct {
	companyId                   string
	webhookId                   string
	updateCompanyWebhookRequest *UpdateCompanyWebhookRequest
}

func (r WebhooksCompanyLevelApiUpdateWebhookInput) UpdateCompanyWebhookRequest(updateCompanyWebhookRequest UpdateCompanyWebhookRequest) WebhooksCompanyLevelApiUpdateWebhookInput {
	r.updateCompanyWebhookRequest = &updateCompanyWebhookRequest
	return r
}

/*
Prepare a request for UpdateWebhook
@param companyId The unique identifier of the company account.@param webhookId Unique identifier of the webhook configuration.
@return WebhooksCompanyLevelApiUpdateWebhookInput
*/
func (a *WebhooksCompanyLevelApi) UpdateWebhookInput(companyId string, webhookId string) WebhooksCompanyLevelApiUpdateWebhookInput {
	return WebhooksCompanyLevelApiUpdateWebhookInput{
		companyId: companyId,
		webhookId: webhookId,
	}
}

/*
UpdateWebhook Update a webhook

Make changes to the configuration of the webhook identified in the path. The request contains the new values you want to have in the webhook configuration. The response contains the full configuration for the webhook, which includes the new values from the request.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Webhooks read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r WebhooksCompanyLevelApiUpdateWebhookInput - Request parameters, see UpdateWebhookInput
@return Webhook, *http.Response, error
*/
func (a *WebhooksCompanyLevelApi) UpdateWebhook(ctx context.Context, r WebhooksCompanyLevelApiUpdateWebhookInput) (Webhook, *http.Response, error) {
	res := &Webhook{}
	path := "/companies/{companyId}/webhooks/{webhookId}"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"webhookId"+"}", url.PathEscape(common.ParameterValueToString(r.webhookId, "webhookId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.updateCompanyWebhookRequest,
		res,
		http.MethodPatch,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}
