/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// PlatformApi service
type PlatformApi common.Service

// All parameters accepted by PlatformApi.GetAllAccountHoldersUnderBalancePlatform
type PlatformApiGetAllAccountHoldersUnderBalancePlatformInput struct {
	id     string
	offset *int32
	limit  *int32
}

// The number of items that you want to skip.
func (r PlatformApiGetAllAccountHoldersUnderBalancePlatformInput) Offset(offset int32) PlatformApiGetAllAccountHoldersUnderBalancePlatformInput {
	r.offset = &offset
	return r
}

// The number of items returned per page, maximum 100 items. By default, the response returns 10 items per page.
func (r PlatformApiGetAllAccountHoldersUnderBalancePlatformInput) Limit(limit int32) PlatformApiGetAllAccountHoldersUnderBalancePlatformInput {
	r.limit = &limit
	return r
}

/*
Prepare a request for GetAllAccountHoldersUnderBalancePlatform
@param id The unique identifier of the balance platform.
@return PlatformApiGetAllAccountHoldersUnderBalancePlatformInput
*/
func (a *PlatformApi) GetAllAccountHoldersUnderBalancePlatformInput(id string) PlatformApiGetAllAccountHoldersUnderBalancePlatformInput {
	return PlatformApiGetAllAccountHoldersUnderBalancePlatformInput{
		id: id,
	}
}

/*
GetAllAccountHoldersUnderBalancePlatform Get all account holders under a balance platform

Returns a paginated list of all the account holders that belong to the balance platform. To fetch multiple pages, use the query parameters.

For example, to limit the page to 5 account holders and to skip the first 20, use `/balancePlatforms/{id}/accountHolders?limit=5&offset=20`.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PlatformApiGetAllAccountHoldersUnderBalancePlatformInput - Request parameters, see GetAllAccountHoldersUnderBalancePlatformInput
@return PaginatedAccountHoldersResponse, *http.Response, error
*/
func (a *PlatformApi) GetAllAccountHoldersUnderBalancePlatform(ctx context.Context, r PlatformApiGetAllAccountHoldersUnderBalancePlatformInput) (PaginatedAccountHoldersResponse, *http.Response, error) {
	res := &PaginatedAccountHoldersResponse{}
	path := "/balancePlatforms/{id}/accountHolders"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.offset != nil {
		common.ParameterAddToQuery(queryParams, "offset", r.offset, "")
	}
	if r.limit != nil {
		common.ParameterAddToQuery(queryParams, "limit", r.limit, "")
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

// All parameters accepted by PlatformApi.GetAllTransactionRulesForBalancePlatform
type PlatformApiGetAllTransactionRulesForBalancePlatformInput struct {
	id string
}

/*
Prepare a request for GetAllTransactionRulesForBalancePlatform
@param id The unique identifier of the balance platform.
@return PlatformApiGetAllTransactionRulesForBalancePlatformInput
*/
func (a *PlatformApi) GetAllTransactionRulesForBalancePlatformInput(id string) PlatformApiGetAllTransactionRulesForBalancePlatformInput {
	return PlatformApiGetAllTransactionRulesForBalancePlatformInput{
		id: id,
	}
}

/*
GetAllTransactionRulesForBalancePlatform Get all transaction rules for a balance platform

Returns a list of transaction rules associated with a balance platform.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PlatformApiGetAllTransactionRulesForBalancePlatformInput - Request parameters, see GetAllTransactionRulesForBalancePlatformInput
@return TransactionRulesResponse, *http.Response, error
*/
func (a *PlatformApi) GetAllTransactionRulesForBalancePlatform(ctx context.Context, r PlatformApiGetAllTransactionRulesForBalancePlatformInput) (TransactionRulesResponse, *http.Response, error) {
	res := &TransactionRulesResponse{}
	path := "/balancePlatforms/{id}/transactionRules"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
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

// All parameters accepted by PlatformApi.GetBalancePlatform
type PlatformApiGetBalancePlatformInput struct {
	id string
}

/*
Prepare a request for GetBalancePlatform
@param id The unique identifier of the balance platform.
@return PlatformApiGetBalancePlatformInput
*/
func (a *PlatformApi) GetBalancePlatformInput(id string) PlatformApiGetBalancePlatformInput {
	return PlatformApiGetBalancePlatformInput{
		id: id,
	}
}

/*
GetBalancePlatform Get a balance platform

Returns a balance platform.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PlatformApiGetBalancePlatformInput - Request parameters, see GetBalancePlatformInput
@return BalancePlatform, *http.Response, error
*/
func (a *PlatformApi) GetBalancePlatform(ctx context.Context, r PlatformApiGetBalancePlatformInput) (BalancePlatform, *http.Response, error) {
	res := &BalancePlatform{}
	path := "/balancePlatforms/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
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
