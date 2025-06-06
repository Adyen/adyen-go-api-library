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

// APICredentialsMerchantLevelApi service
type APICredentialsMerchantLevelApi common.Service

// All parameters accepted by APICredentialsMerchantLevelApi.CreateApiCredential
type APICredentialsMerchantLevelApiCreateApiCredentialInput struct {
	merchantId                         string
	createMerchantApiCredentialRequest *CreateMerchantApiCredentialRequest
}

func (r APICredentialsMerchantLevelApiCreateApiCredentialInput) CreateMerchantApiCredentialRequest(createMerchantApiCredentialRequest CreateMerchantApiCredentialRequest) APICredentialsMerchantLevelApiCreateApiCredentialInput {
	r.createMerchantApiCredentialRequest = &createMerchantApiCredentialRequest
	return r
}

/*
Prepare a request for CreateApiCredential
@param merchantId The unique identifier of the merchant account.
@return APICredentialsMerchantLevelApiCreateApiCredentialInput
*/
func (a *APICredentialsMerchantLevelApi) CreateApiCredentialInput(merchantId string) APICredentialsMerchantLevelApiCreateApiCredentialInput {
	return APICredentialsMerchantLevelApiCreateApiCredentialInput{
		merchantId: merchantId,
	}
}

/*
CreateApiCredential Create an API credential

Creates an [API credential](https://docs.adyen.com/development-resources/api-credentials) for the company account identified in the path. In the request, you can specify the roles and allowed origins for the new API credential.

The response includes the:
* [API key](https://docs.adyen.com/development-resources/api-authentication#api-key-authentication): used for API request authentication.
* [Client key](https://docs.adyen.com/development-resources/client-side-authentication#how-it-works): public key used for client-side authentication.
* [Username and password](https://docs.adyen.com/development-resources/api-authentication#using-basic-authentication): used for basic authentication.

> Make sure you store the API key securely in your system. You won't be able to retrieve it later.

If your API key is lost or compromised, you need to [generate a new API key](https://docs.adyen.com/api-explorer/#/ManagementService/v1/post/merchants/{merchantId}/apiCredentials/{apiCredentialId}/generateApiKey).

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r APICredentialsMerchantLevelApiCreateApiCredentialInput - Request parameters, see CreateApiCredentialInput
@return CreateApiCredentialResponse, *http.Response, error
*/
func (a *APICredentialsMerchantLevelApi) CreateApiCredential(ctx context.Context, r APICredentialsMerchantLevelApiCreateApiCredentialInput) (CreateApiCredentialResponse, *http.Response, error) {
	res := &CreateApiCredentialResponse{}
	path := "/merchants/{merchantId}/apiCredentials"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.createMerchantApiCredentialRequest,
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

// All parameters accepted by APICredentialsMerchantLevelApi.GetApiCredential
type APICredentialsMerchantLevelApiGetApiCredentialInput struct {
	merchantId      string
	apiCredentialId string
}

/*
Prepare a request for GetApiCredential
@param merchantId The unique identifier of the merchant account.@param apiCredentialId Unique identifier of the API credential.
@return APICredentialsMerchantLevelApiGetApiCredentialInput
*/
func (a *APICredentialsMerchantLevelApi) GetApiCredentialInput(merchantId string, apiCredentialId string) APICredentialsMerchantLevelApiGetApiCredentialInput {
	return APICredentialsMerchantLevelApiGetApiCredentialInput{
		merchantId:      merchantId,
		apiCredentialId: apiCredentialId,
	}
}

/*
GetApiCredential Get an API credential

Returns the [API credential](https://docs.adyen.com/development-resources/api-credentials) identified in the path.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r APICredentialsMerchantLevelApiGetApiCredentialInput - Request parameters, see GetApiCredentialInput
@return ApiCredential, *http.Response, error
*/
func (a *APICredentialsMerchantLevelApi) GetApiCredential(ctx context.Context, r APICredentialsMerchantLevelApiGetApiCredentialInput) (ApiCredential, *http.Response, error) {
	res := &ApiCredential{}
	path := "/merchants/{merchantId}/apiCredentials/{apiCredentialId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"apiCredentialId"+"}", url.PathEscape(common.ParameterValueToString(r.apiCredentialId, "apiCredentialId")), -1)
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

// All parameters accepted by APICredentialsMerchantLevelApi.ListApiCredentials
type APICredentialsMerchantLevelApiListApiCredentialsInput struct {
	merchantId string
	pageNumber *int32
	pageSize   *int32
}

// The number of the page to fetch.
func (r APICredentialsMerchantLevelApiListApiCredentialsInput) PageNumber(pageNumber int32) APICredentialsMerchantLevelApiListApiCredentialsInput {
	r.pageNumber = &pageNumber
	return r
}

// The number of items to have on a page, maximum 100. The default is 10 items on a page.
func (r APICredentialsMerchantLevelApiListApiCredentialsInput) PageSize(pageSize int32) APICredentialsMerchantLevelApiListApiCredentialsInput {
	r.pageSize = &pageSize
	return r
}

/*
Prepare a request for ListApiCredentials
@param merchantId The unique identifier of the merchant account.
@return APICredentialsMerchantLevelApiListApiCredentialsInput
*/
func (a *APICredentialsMerchantLevelApi) ListApiCredentialsInput(merchantId string) APICredentialsMerchantLevelApiListApiCredentialsInput {
	return APICredentialsMerchantLevelApiListApiCredentialsInput{
		merchantId: merchantId,
	}
}

/*
ListApiCredentials Get a list of API credentials

Returns the list of [API credentials](https://docs.adyen.com/development-resources/api-credentials) for the merchant account. The list is grouped into pages as defined by the query parameters.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r APICredentialsMerchantLevelApiListApiCredentialsInput - Request parameters, see ListApiCredentialsInput
@return ListMerchantApiCredentialsResponse, *http.Response, error
*/
func (a *APICredentialsMerchantLevelApi) ListApiCredentials(ctx context.Context, r APICredentialsMerchantLevelApiListApiCredentialsInput) (ListMerchantApiCredentialsResponse, *http.Response, error) {
	res := &ListMerchantApiCredentialsResponse{}
	path := "/merchants/{merchantId}/apiCredentials"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
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

// All parameters accepted by APICredentialsMerchantLevelApi.UpdateApiCredential
type APICredentialsMerchantLevelApiUpdateApiCredentialInput struct {
	merchantId                         string
	apiCredentialId                    string
	updateMerchantApiCredentialRequest *UpdateMerchantApiCredentialRequest
}

func (r APICredentialsMerchantLevelApiUpdateApiCredentialInput) UpdateMerchantApiCredentialRequest(updateMerchantApiCredentialRequest UpdateMerchantApiCredentialRequest) APICredentialsMerchantLevelApiUpdateApiCredentialInput {
	r.updateMerchantApiCredentialRequest = &updateMerchantApiCredentialRequest
	return r
}

/*
Prepare a request for UpdateApiCredential
@param merchantId The unique identifier of the merchant account.@param apiCredentialId Unique identifier of the API credential.
@return APICredentialsMerchantLevelApiUpdateApiCredentialInput
*/
func (a *APICredentialsMerchantLevelApi) UpdateApiCredentialInput(merchantId string, apiCredentialId string) APICredentialsMerchantLevelApiUpdateApiCredentialInput {
	return APICredentialsMerchantLevelApiUpdateApiCredentialInput{
		merchantId:      merchantId,
		apiCredentialId: apiCredentialId,
	}
}

/*
UpdateApiCredential Update an API credential

Changes the API credential's roles, or allowed origins. The request has the new values for the fields you want to change. The response contains the full updated API credential, including the new values from the request.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r APICredentialsMerchantLevelApiUpdateApiCredentialInput - Request parameters, see UpdateApiCredentialInput
@return ApiCredential, *http.Response, error
*/
func (a *APICredentialsMerchantLevelApi) UpdateApiCredential(ctx context.Context, r APICredentialsMerchantLevelApiUpdateApiCredentialInput) (ApiCredential, *http.Response, error) {
	res := &ApiCredential{}
	path := "/merchants/{merchantId}/apiCredentials/{apiCredentialId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"apiCredentialId"+"}", url.PathEscape(common.ParameterValueToString(r.apiCredentialId, "apiCredentialId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.updateMerchantApiCredentialRequest,
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
