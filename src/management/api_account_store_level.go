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

// AccountStoreLevelApi AccountStoreLevelApi service
type AccountStoreLevelApi common.Service

type AccountStoreLevelApiCreateStoreConfig struct {
	ctx                                  context.Context
	storeCreationWithMerchantCodeRequest *StoreCreationWithMerchantCodeRequest
}

func (r AccountStoreLevelApiCreateStoreConfig) StoreCreationWithMerchantCodeRequest(storeCreationWithMerchantCodeRequest StoreCreationWithMerchantCodeRequest) AccountStoreLevelApiCreateStoreConfig {
	r.storeCreationWithMerchantCodeRequest = &storeCreationWithMerchantCodeRequest
	return r
}

/*
CreateStore Create a store

Creates a store for the merchant account specified in the request.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Stores read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AccountStoreLevelApiCreateStoreConfig
*/
func (a *AccountStoreLevelApi) CreateStoreConfig(ctx context.Context) AccountStoreLevelApiCreateStoreConfig {
	return AccountStoreLevelApiCreateStoreConfig{
		ctx: ctx,
	}
}

/*
Create a store
Creates a store for the merchant account specified in the request.  To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Stores read and write
 * @param req StoreCreationWithMerchantCodeRequest - reference of StoreCreationWithMerchantCodeRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return Store
*/

func (a *AccountStoreLevelApi) CreateStore(r AccountStoreLevelApiCreateStoreConfig) (Store, *_nethttp.Response, error) {
	res := &Store{}
	path := "/stores"
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPost, r.storeCreationWithMerchantCodeRequest, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type AccountStoreLevelApiCreateStoreByMerchantIdConfig struct {
	ctx                  context.Context
	merchantId           string
	storeCreationRequest *StoreCreationRequest
}

func (r AccountStoreLevelApiCreateStoreByMerchantIdConfig) StoreCreationRequest(storeCreationRequest StoreCreationRequest) AccountStoreLevelApiCreateStoreByMerchantIdConfig {
	r.storeCreationRequest = &storeCreationRequest
	return r
}

/*
CreateStoreByMerchantId Create a store

Creates a store for the merchant account identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Stores read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @return AccountStoreLevelApiCreateStoreByMerchantIdConfig
*/
func (a *AccountStoreLevelApi) CreateStoreByMerchantIdConfig(ctx context.Context, merchantId string) AccountStoreLevelApiCreateStoreByMerchantIdConfig {
	return AccountStoreLevelApiCreateStoreByMerchantIdConfig{
		ctx:        ctx,
		merchantId: merchantId,
	}
}

/*
Create a store
Creates a store for the merchant account identified in the path.  To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Stores read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param req StoreCreationRequest - reference of StoreCreationRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return Store
*/

func (a *AccountStoreLevelApi) CreateStoreByMerchantId(r AccountStoreLevelApiCreateStoreByMerchantIdConfig) (Store, *_nethttp.Response, error) {
	res := &Store{}
	path := "/merchants/{merchantId}/stores"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPost, r.storeCreationRequest, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type AccountStoreLevelApiGetStoreConfig struct {
	ctx        context.Context
	merchantId string
	storeId    string
}

/*
GetStore Get a store

Returns the details of the store identified in the path.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Stores read
* Management API—Stores read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @param storeId The unique identifier of the store.
 @return AccountStoreLevelApiGetStoreConfig
*/
func (a *AccountStoreLevelApi) GetStoreConfig(ctx context.Context, merchantId string, storeId string) AccountStoreLevelApiGetStoreConfig {
	return AccountStoreLevelApiGetStoreConfig{
		ctx:        ctx,
		merchantId: merchantId,
		storeId:    storeId,
	}
}

/*
Get a store
Returns the details of the store identified in the path.  To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Stores read * Management API—Stores read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param storeId The unique identifier of the store.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return Store
*/

func (a *AccountStoreLevelApi) GetStore(r AccountStoreLevelApiGetStoreConfig) (Store, *_nethttp.Response, error) {
	res := &Store{}
	path := "/merchants/{merchantId}/stores/{storeId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"storeId"+"}", url.PathEscape(common.ParameterValueToString(r.storeId, "storeId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type AccountStoreLevelApiGetStoreByIdConfig struct {
	ctx     context.Context
	storeId string
}

/*
GetStoreById Get a store

Returns the details of the store identified in the path.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Stores read
* Management API—Stores read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeId The unique identifier of the store.
 @return AccountStoreLevelApiGetStoreByIdConfig
*/
func (a *AccountStoreLevelApi) GetStoreByIdConfig(ctx context.Context, storeId string) AccountStoreLevelApiGetStoreByIdConfig {
	return AccountStoreLevelApiGetStoreByIdConfig{
		ctx:     ctx,
		storeId: storeId,
	}
}

/*
Get a store
Returns the details of the store identified in the path.  To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Stores read * Management API—Stores read and write
 * @param storeId The unique identifier of the store.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return Store
*/

func (a *AccountStoreLevelApi) GetStoreById(r AccountStoreLevelApiGetStoreByIdConfig) (Store, *_nethttp.Response, error) {
	res := &Store{}
	path := "/stores/{storeId}"
	path = strings.Replace(path, "{"+"storeId"+"}", url.PathEscape(common.ParameterValueToString(r.storeId, "storeId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type AccountStoreLevelApiListStoresConfig struct {
	ctx        context.Context
	pageNumber *int32
	pageSize   *int32
	reference  *string
	merchantId *string
}

// The number of the page to fetch.
func (r AccountStoreLevelApiListStoresConfig) PageNumber(pageNumber int32) AccountStoreLevelApiListStoresConfig {
	r.pageNumber = &pageNumber
	return r
}

// The number of items to have on a page, maximum 100. The default is 10 items on a page.
func (r AccountStoreLevelApiListStoresConfig) PageSize(pageSize int32) AccountStoreLevelApiListStoresConfig {
	r.pageSize = &pageSize
	return r
}

// The reference of the store.
func (r AccountStoreLevelApiListStoresConfig) Reference(reference string) AccountStoreLevelApiListStoresConfig {
	r.reference = &reference
	return r
}

// The unique identifier of the merchant account.
func (r AccountStoreLevelApiListStoresConfig) MerchantId(merchantId string) AccountStoreLevelApiListStoresConfig {
	r.merchantId = &merchantId
	return r
}

/*
ListStores Get a list of stores

Returns a list of stores. The list is grouped into pages as defined by the query parameters.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Stores read
* Management API—Stores read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AccountStoreLevelApiListStoresConfig
*/
func (a *AccountStoreLevelApi) ListStoresConfig(ctx context.Context) AccountStoreLevelApiListStoresConfig {
	return AccountStoreLevelApiListStoresConfig{
		ctx: ctx,
	}
}

/*
Get a list of stores
Returns a list of stores. The list is grouped into pages as defined by the query parameters.  To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Stores read * Management API—Stores read and write
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ListStoresResponse
*/

func (a *AccountStoreLevelApi) ListStores(r AccountStoreLevelApiListStoresConfig) (ListStoresResponse, *_nethttp.Response, error) {
	res := &ListStoresResponse{}
	path := "/stores"
	queryString := url.Values{}
	if r.pageNumber != nil {
		common.ParameterAddToQuery(queryString, "pageNumber", r.pageNumber, "")
	}
	if r.pageSize != nil {
		common.ParameterAddToQuery(queryString, "pageSize", r.pageSize, "")
	}
	if r.reference != nil {
		common.ParameterAddToQuery(queryString, "reference", r.reference, "")
	}
	if r.merchantId != nil {
		common.ParameterAddToQuery(queryString, "merchantId", r.merchantId, "")
	}
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path+"?"+queryString.Encode(), []_context.Context{r.ctx})
	return *res, httpRes, err
}

type AccountStoreLevelApiListStoresByMerchantIdConfig struct {
	ctx        context.Context
	merchantId string
	pageNumber *int32
	pageSize   *int32
	reference  *string
}

// The number of the page to fetch.
func (r AccountStoreLevelApiListStoresByMerchantIdConfig) PageNumber(pageNumber int32) AccountStoreLevelApiListStoresByMerchantIdConfig {
	r.pageNumber = &pageNumber
	return r
}

// The number of items to have on a page, maximum 100. The default is 10 items on a page.
func (r AccountStoreLevelApiListStoresByMerchantIdConfig) PageSize(pageSize int32) AccountStoreLevelApiListStoresByMerchantIdConfig {
	r.pageSize = &pageSize
	return r
}

// The reference of the store.
func (r AccountStoreLevelApiListStoresByMerchantIdConfig) Reference(reference string) AccountStoreLevelApiListStoresByMerchantIdConfig {
	r.reference = &reference
	return r
}

/*
ListStoresByMerchantId Get a list of stores

Returns a list of stores for the merchant account identified in the path. The list is grouped into pages as defined by the query parameters.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Stores read
* Management API—Stores read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @return AccountStoreLevelApiListStoresByMerchantIdConfig
*/
func (a *AccountStoreLevelApi) ListStoresByMerchantIdConfig(ctx context.Context, merchantId string) AccountStoreLevelApiListStoresByMerchantIdConfig {
	return AccountStoreLevelApiListStoresByMerchantIdConfig{
		ctx:        ctx,
		merchantId: merchantId,
	}
}

/*
Get a list of stores
Returns a list of stores for the merchant account identified in the path. The list is grouped into pages as defined by the query parameters.  To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Stores read * Management API—Stores read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ListStoresResponse
*/

func (a *AccountStoreLevelApi) ListStoresByMerchantId(r AccountStoreLevelApiListStoresByMerchantIdConfig) (ListStoresResponse, *_nethttp.Response, error) {
	res := &ListStoresResponse{}
	path := "/merchants/{merchantId}/stores"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryString := url.Values{}
	if r.pageNumber != nil {
		common.ParameterAddToQuery(queryString, "pageNumber", r.pageNumber, "")
	}
	if r.pageSize != nil {
		common.ParameterAddToQuery(queryString, "pageSize", r.pageSize, "")
	}
	if r.reference != nil {
		common.ParameterAddToQuery(queryString, "reference", r.reference, "")
	}
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path+"?"+queryString.Encode(), []_context.Context{r.ctx})
	return *res, httpRes, err
}

type AccountStoreLevelApiUpdateStoreConfig struct {
	ctx                context.Context
	merchantId         string
	storeId            string
	updateStoreRequest *UpdateStoreRequest
}

func (r AccountStoreLevelApiUpdateStoreConfig) UpdateStoreRequest(updateStoreRequest UpdateStoreRequest) AccountStoreLevelApiUpdateStoreConfig {
	r.updateStoreRequest = &updateStoreRequest
	return r
}

/*
UpdateStore Update a store

Updates the store identified in the path. You can only update some store parameters.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Stores read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @param storeId The unique identifier of the store.
 @return AccountStoreLevelApiUpdateStoreConfig
*/
func (a *AccountStoreLevelApi) UpdateStoreConfig(ctx context.Context, merchantId string, storeId string) AccountStoreLevelApiUpdateStoreConfig {
	return AccountStoreLevelApiUpdateStoreConfig{
		ctx:        ctx,
		merchantId: merchantId,
		storeId:    storeId,
	}
}

/*
Update a store
Updates the store identified in the path. You can only update some store parameters.  To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Stores read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param storeId The unique identifier of the store.
 * @param req UpdateStoreRequest - reference of UpdateStoreRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return Store
*/

func (a *AccountStoreLevelApi) UpdateStore(r AccountStoreLevelApiUpdateStoreConfig) (Store, *_nethttp.Response, error) {
	res := &Store{}
	path := "/merchants/{merchantId}/stores/{storeId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"storeId"+"}", url.PathEscape(common.ParameterValueToString(r.storeId, "storeId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPatch, r.updateStoreRequest, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type AccountStoreLevelApiUpdateStoreByIdConfig struct {
	ctx                context.Context
	storeId            string
	updateStoreRequest *UpdateStoreRequest
}

func (r AccountStoreLevelApiUpdateStoreByIdConfig) UpdateStoreRequest(updateStoreRequest UpdateStoreRequest) AccountStoreLevelApiUpdateStoreByIdConfig {
	r.updateStoreRequest = &updateStoreRequest
	return r
}

/*
UpdateStoreById Update a store

Updates the store identified in the path.
You can only update some store parameters.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Stores read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeId The unique identifier of the store.
 @return AccountStoreLevelApiUpdateStoreByIdConfig
*/
func (a *AccountStoreLevelApi) UpdateStoreByIdConfig(ctx context.Context, storeId string) AccountStoreLevelApiUpdateStoreByIdConfig {
	return AccountStoreLevelApiUpdateStoreByIdConfig{
		ctx:     ctx,
		storeId: storeId,
	}
}

/*
Update a store
Updates the store identified in the path. You can only update some store parameters.  To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Stores read and write
 * @param storeId The unique identifier of the store.
 * @param req UpdateStoreRequest - reference of UpdateStoreRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return Store
*/

func (a *AccountStoreLevelApi) UpdateStoreById(r AccountStoreLevelApiUpdateStoreByIdConfig) (Store, *_nethttp.Response, error) {
	res := &Store{}
	path := "/stores/{storeId}"
	path = strings.Replace(path, "{"+"storeId"+"}", url.PathEscape(common.ParameterValueToString(r.storeId, "storeId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPatch, r.updateStoreRequest, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}