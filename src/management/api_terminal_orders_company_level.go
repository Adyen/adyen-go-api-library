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

// TerminalOrdersCompanyLevelApi TerminalOrdersCompanyLevelApi service
type TerminalOrdersCompanyLevelApi common.Service

type TerminalOrdersCompanyLevelApiCancelOrderConfig struct {
	ctx       context.Context
	companyId string
	orderId   string
}

/*
CancelOrder Cancel an order

Cancels the terminal products order identified in the path.
Cancelling is only possible while the order has the status **Placed**.
To cancel an order, make a POST call without a request body. The response returns the full order details, but with the status changed to **Cancelled**.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The unique identifier of the company account.
 @param orderId The unique identifier of the order.
 @return TerminalOrdersCompanyLevelApiCancelOrderConfig
*/
func (a *TerminalOrdersCompanyLevelApi) CancelOrderConfig(ctx context.Context, companyId string, orderId string) TerminalOrdersCompanyLevelApiCancelOrderConfig {
	return TerminalOrdersCompanyLevelApiCancelOrderConfig{
		ctx:       ctx,
		companyId: companyId,
		orderId:   orderId,
	}
}

/*
Cancel an order
Cancels the terminal products order identified in the path. Cancelling is only possible while the order has the status **Placed**. To cancel an order, make a POST call without a request body. The response returns the full order details, but with the status changed to **Cancelled**.  To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal ordering read and write
 * @param companyId The unique identifier of the company account.
 * @param orderId The unique identifier of the order.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return TerminalOrder
*/

func (a *TerminalOrdersCompanyLevelApi) CancelOrder(r TerminalOrdersCompanyLevelApiCancelOrderConfig) (TerminalOrder, *_nethttp.Response, error) {
	res := &TerminalOrder{}
	path := "/companies/{companyId}/terminalOrders/{orderId}/cancel"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"orderId"+"}", url.PathEscape(common.ParameterValueToString(r.orderId, "orderId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPost, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalOrdersCompanyLevelApiCreateOrderConfig struct {
	ctx                  context.Context
	companyId            string
	terminalOrderRequest *TerminalOrderRequest
}

func (r TerminalOrdersCompanyLevelApiCreateOrderConfig) TerminalOrderRequest(terminalOrderRequest TerminalOrderRequest) TerminalOrdersCompanyLevelApiCreateOrderConfig {
	r.terminalOrderRequest = &terminalOrderRequest
	return r
}

/*
CreateOrder Create an order

Creates an order for payment terminal products for the company identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The unique identifier of the company account.
 @return TerminalOrdersCompanyLevelApiCreateOrderConfig
*/
func (a *TerminalOrdersCompanyLevelApi) CreateOrderConfig(ctx context.Context, companyId string) TerminalOrdersCompanyLevelApiCreateOrderConfig {
	return TerminalOrdersCompanyLevelApiCreateOrderConfig{
		ctx:       ctx,
		companyId: companyId,
	}
}

/*
Create an order
Creates an order for payment terminal products for the company identified in the path.  To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal ordering read and write
 * @param companyId The unique identifier of the company account.
 * @param req TerminalOrderRequest - reference of TerminalOrderRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return TerminalOrder
*/

func (a *TerminalOrdersCompanyLevelApi) CreateOrder(r TerminalOrdersCompanyLevelApiCreateOrderConfig) (TerminalOrder, *_nethttp.Response, error) {
	res := &TerminalOrder{}
	path := "/companies/{companyId}/terminalOrders"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPost, r.terminalOrderRequest, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalOrdersCompanyLevelApiCreateShippingLocationConfig struct {
	ctx              context.Context
	companyId        string
	shippingLocation *ShippingLocation
}

func (r TerminalOrdersCompanyLevelApiCreateShippingLocationConfig) ShippingLocation(shippingLocation ShippingLocation) TerminalOrdersCompanyLevelApiCreateShippingLocationConfig {
	r.shippingLocation = &shippingLocation
	return r
}

/*
CreateShippingLocation Create a shipping location

Creates a shipping location for the company identified in the path. A shipping location defines an address where orders can be delivered.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The unique identifier of the company account.
 @return TerminalOrdersCompanyLevelApiCreateShippingLocationConfig
*/
func (a *TerminalOrdersCompanyLevelApi) CreateShippingLocationConfig(ctx context.Context, companyId string) TerminalOrdersCompanyLevelApiCreateShippingLocationConfig {
	return TerminalOrdersCompanyLevelApiCreateShippingLocationConfig{
		ctx:       ctx,
		companyId: companyId,
	}
}

/*
Create a shipping location
Creates a shipping location for the company identified in the path. A shipping location defines an address where orders can be delivered.  To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal ordering read and write
 * @param companyId The unique identifier of the company account.
 * @param req ShippingLocation - reference of ShippingLocation).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ShippingLocation
*/

func (a *TerminalOrdersCompanyLevelApi) CreateShippingLocation(r TerminalOrdersCompanyLevelApiCreateShippingLocationConfig) (ShippingLocation, *_nethttp.Response, error) {
	res := &ShippingLocation{}
	path := "/companies/{companyId}/shippingLocations"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPost, r.shippingLocation, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalOrdersCompanyLevelApiGetOrderConfig struct {
	ctx       context.Context
	companyId string
	orderId   string
}

/*
GetOrder Get an order

Returns the details of the terminal products order identified in the path.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read
* Management API—Terminal ordering read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The unique identifier of the company account.
 @param orderId The unique identifier of the order.
 @return TerminalOrdersCompanyLevelApiGetOrderConfig
*/
func (a *TerminalOrdersCompanyLevelApi) GetOrderConfig(ctx context.Context, companyId string, orderId string) TerminalOrdersCompanyLevelApiGetOrderConfig {
	return TerminalOrdersCompanyLevelApiGetOrderConfig{
		ctx:       ctx,
		companyId: companyId,
		orderId:   orderId,
	}
}

/*
Get an order
Returns the details of the terminal products order identified in the path.  To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal ordering read * Management API—Terminal ordering read and write
 * @param companyId The unique identifier of the company account.
 * @param orderId The unique identifier of the order.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return TerminalOrder
*/

func (a *TerminalOrdersCompanyLevelApi) GetOrder(r TerminalOrdersCompanyLevelApiGetOrderConfig) (TerminalOrder, *_nethttp.Response, error) {
	res := &TerminalOrder{}
	path := "/companies/{companyId}/terminalOrders/{orderId}"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"orderId"+"}", url.PathEscape(common.ParameterValueToString(r.orderId, "orderId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalOrdersCompanyLevelApiListBillingEntitiesConfig struct {
	ctx       context.Context
	companyId string
	name      *string
}

// The name of the billing entity.
func (r TerminalOrdersCompanyLevelApiListBillingEntitiesConfig) Name(name string) TerminalOrdersCompanyLevelApiListBillingEntitiesConfig {
	r.name = &name
	return r
}

/*
ListBillingEntities Get a list of billing entities

Returns the billing entities of the company identified in the path and all merchant accounts belonging to the company.
A billing entity is a legal entity where we charge orders to. An order for terminal products must contain the ID of a billing entity.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read
* Management API—Terminal ordering read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The unique identifier of the company account.
 @return TerminalOrdersCompanyLevelApiListBillingEntitiesConfig
*/
func (a *TerminalOrdersCompanyLevelApi) ListBillingEntitiesConfig(ctx context.Context, companyId string) TerminalOrdersCompanyLevelApiListBillingEntitiesConfig {
	return TerminalOrdersCompanyLevelApiListBillingEntitiesConfig{
		ctx:       ctx,
		companyId: companyId,
	}
}

/*
Get a list of billing entities
Returns the billing entities of the company identified in the path and all merchant accounts belonging to the company. A billing entity is a legal entity where we charge orders to. An order for terminal products must contain the ID of a billing entity.  To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal ordering read * Management API—Terminal ordering read and write
 * @param companyId The unique identifier of the company account.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return BillingEntitiesResponse
*/

func (a *TerminalOrdersCompanyLevelApi) ListBillingEntities(r TerminalOrdersCompanyLevelApiListBillingEntitiesConfig) (BillingEntitiesResponse, *_nethttp.Response, error) {
	res := &BillingEntitiesResponse{}
	path := "/companies/{companyId}/billingEntities"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	queryString := url.Values{}
	if r.name != nil {
		common.ParameterAddToQuery(queryString, "name", r.name, "")
	}
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path+"?"+queryString.Encode(), []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalOrdersCompanyLevelApiListOrdersConfig struct {
	ctx                    context.Context
	companyId              string
	customerOrderReference *string
	status                 *string
	offset                 *int32
	limit                  *int32
}

// Your purchase order number.
func (r TerminalOrdersCompanyLevelApiListOrdersConfig) CustomerOrderReference(customerOrderReference string) TerminalOrdersCompanyLevelApiListOrdersConfig {
	r.customerOrderReference = &customerOrderReference
	return r
}

// The order status. Possible values (not case-sensitive): Placed, Confirmed, Cancelled, Shipped, Delivered.
func (r TerminalOrdersCompanyLevelApiListOrdersConfig) Status(status string) TerminalOrdersCompanyLevelApiListOrdersConfig {
	r.status = &status
	return r
}

// The number of orders to skip.
func (r TerminalOrdersCompanyLevelApiListOrdersConfig) Offset(offset int32) TerminalOrdersCompanyLevelApiListOrdersConfig {
	r.offset = &offset
	return r
}

// The number of orders to return.
func (r TerminalOrdersCompanyLevelApiListOrdersConfig) Limit(limit int32) TerminalOrdersCompanyLevelApiListOrdersConfig {
	r.limit = &limit
	return r
}

/*
ListOrders Get a list of orders

Returns a lists of terminal products orders for the company identified in the path.
To filter the list, use one or more of the query parameters.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read
* Management API—Terminal ordering read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The unique identifier of the company account.
 @return TerminalOrdersCompanyLevelApiListOrdersConfig
*/
func (a *TerminalOrdersCompanyLevelApi) ListOrdersConfig(ctx context.Context, companyId string) TerminalOrdersCompanyLevelApiListOrdersConfig {
	return TerminalOrdersCompanyLevelApiListOrdersConfig{
		ctx:       ctx,
		companyId: companyId,
	}
}

/*
Get a list of orders
Returns a lists of terminal products orders for the company identified in the path. To filter the list, use one or more of the query parameters.  To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal ordering read * Management API—Terminal ordering read and write
 * @param companyId The unique identifier of the company account.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return TerminalOrdersResponse
*/

func (a *TerminalOrdersCompanyLevelApi) ListOrders(r TerminalOrdersCompanyLevelApiListOrdersConfig) (TerminalOrdersResponse, *_nethttp.Response, error) {
	res := &TerminalOrdersResponse{}
	path := "/companies/{companyId}/terminalOrders"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	queryString := url.Values{}
	if r.customerOrderReference != nil {
		common.ParameterAddToQuery(queryString, "customerOrderReference", r.customerOrderReference, "")
	}
	if r.status != nil {
		common.ParameterAddToQuery(queryString, "status", r.status, "")
	}
	if r.offset != nil {
		common.ParameterAddToQuery(queryString, "offset", r.offset, "")
	}
	if r.limit != nil {
		common.ParameterAddToQuery(queryString, "limit", r.limit, "")
	}
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path+"?"+queryString.Encode(), []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalOrdersCompanyLevelApiListShippingLocationsConfig struct {
	ctx       context.Context
	companyId string
	name      *string
	offset    *int32
	limit     *int32
}

// The name of the shipping location.
func (r TerminalOrdersCompanyLevelApiListShippingLocationsConfig) Name(name string) TerminalOrdersCompanyLevelApiListShippingLocationsConfig {
	r.name = &name
	return r
}

// The number of locations to skip.
func (r TerminalOrdersCompanyLevelApiListShippingLocationsConfig) Offset(offset int32) TerminalOrdersCompanyLevelApiListShippingLocationsConfig {
	r.offset = &offset
	return r
}

// The number of locations to return.
func (r TerminalOrdersCompanyLevelApiListShippingLocationsConfig) Limit(limit int32) TerminalOrdersCompanyLevelApiListShippingLocationsConfig {
	r.limit = &limit
	return r
}

/*
ListShippingLocations Get a list of shipping locations

Returns the shipping locations for the company identified in the path and all merchant accounts belonging to the company.
A shipping location includes the address where orders can be delivered, and an ID which you need to specify when ordering terminal products.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read
* Management API—Terminal ordering read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The unique identifier of the company account.
 @return TerminalOrdersCompanyLevelApiListShippingLocationsConfig
*/
func (a *TerminalOrdersCompanyLevelApi) ListShippingLocationsConfig(ctx context.Context, companyId string) TerminalOrdersCompanyLevelApiListShippingLocationsConfig {
	return TerminalOrdersCompanyLevelApiListShippingLocationsConfig{
		ctx:       ctx,
		companyId: companyId,
	}
}

/*
Get a list of shipping locations
Returns the shipping locations for the company identified in the path and all merchant accounts belonging to the company. A shipping location includes the address where orders can be delivered, and an ID which you need to specify when ordering terminal products.  To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal ordering read * Management API—Terminal ordering read and write
 * @param companyId The unique identifier of the company account.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ShippingLocationsResponse
*/

func (a *TerminalOrdersCompanyLevelApi) ListShippingLocations(r TerminalOrdersCompanyLevelApiListShippingLocationsConfig) (ShippingLocationsResponse, *_nethttp.Response, error) {
	res := &ShippingLocationsResponse{}
	path := "/companies/{companyId}/shippingLocations"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	queryString := url.Values{}
	if r.name != nil {
		common.ParameterAddToQuery(queryString, "name", r.name, "")
	}
	if r.offset != nil {
		common.ParameterAddToQuery(queryString, "offset", r.offset, "")
	}
	if r.limit != nil {
		common.ParameterAddToQuery(queryString, "limit", r.limit, "")
	}
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path+"?"+queryString.Encode(), []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalOrdersCompanyLevelApiListTerminalModelsConfig struct {
	ctx       context.Context
	companyId string
}

/*
ListTerminalModels Get a list of terminal models

Returns a list of payment terminal models that the company identified in the path has access to.
The response includes the terminal model ID, which can be used as a query parameter when getting a list of terminals or a list of products for ordering.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read
* Management API—Terminal ordering read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The unique identifier of the company account.
 @return TerminalOrdersCompanyLevelApiListTerminalModelsConfig
*/
func (a *TerminalOrdersCompanyLevelApi) ListTerminalModelsConfig(ctx context.Context, companyId string) TerminalOrdersCompanyLevelApiListTerminalModelsConfig {
	return TerminalOrdersCompanyLevelApiListTerminalModelsConfig{
		ctx:       ctx,
		companyId: companyId,
	}
}

/*
Get a list of terminal models
Returns a list of payment terminal models that the company identified in the path has access to. The response includes the terminal model ID, which can be used as a query parameter when getting a list of terminals or a list of products for ordering.  To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal ordering read * Management API—Terminal ordering read and write
 * @param companyId The unique identifier of the company account.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return TerminalModelsResponse
*/

func (a *TerminalOrdersCompanyLevelApi) ListTerminalModels(r TerminalOrdersCompanyLevelApiListTerminalModelsConfig) (TerminalModelsResponse, *_nethttp.Response, error) {
	res := &TerminalModelsResponse{}
	path := "/companies/{companyId}/terminalModels"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalOrdersCompanyLevelApiListTerminalProductsConfig struct {
	ctx             context.Context
	companyId       string
	country         *string
	terminalModelId *string
	offset          *int32
	limit           *int32
}

// The country to return products for, in [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format. For example, **US**
func (r TerminalOrdersCompanyLevelApiListTerminalProductsConfig) Country(country string) TerminalOrdersCompanyLevelApiListTerminalProductsConfig {
	r.country = &country
	return r
}

// The terminal model to return products for. Use the ID returned in the [GET &#x60;/terminalModels&#x60;](https://docs.adyen.com/api-explorer/#/ManagementService/latest/get/companies/{companyId}/terminalModels) response. For example, **Verifone.M400**
func (r TerminalOrdersCompanyLevelApiListTerminalProductsConfig) TerminalModelId(terminalModelId string) TerminalOrdersCompanyLevelApiListTerminalProductsConfig {
	r.terminalModelId = &terminalModelId
	return r
}

// The number of products to skip.
func (r TerminalOrdersCompanyLevelApiListTerminalProductsConfig) Offset(offset int32) TerminalOrdersCompanyLevelApiListTerminalProductsConfig {
	r.offset = &offset
	return r
}

// The number of products to return.
func (r TerminalOrdersCompanyLevelApiListTerminalProductsConfig) Limit(limit int32) TerminalOrdersCompanyLevelApiListTerminalProductsConfig {
	r.limit = &limit
	return r
}

/*
ListTerminalProducts Get a list of terminal products

Returns a country-specific list of payment terminal packages and parts that the company identified in the path has access to.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read
* Management API—Terminal ordering read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The unique identifier of the company account.
 @return TerminalOrdersCompanyLevelApiListTerminalProductsConfig
*/
func (a *TerminalOrdersCompanyLevelApi) ListTerminalProductsConfig(ctx context.Context, companyId string) TerminalOrdersCompanyLevelApiListTerminalProductsConfig {
	return TerminalOrdersCompanyLevelApiListTerminalProductsConfig{
		ctx:       ctx,
		companyId: companyId,
	}
}

/*
Get a list of terminal products
Returns a country-specific list of payment terminal packages and parts that the company identified in the path has access to.   To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal ordering read * Management API—Terminal ordering read and write
 * @param companyId The unique identifier of the company account.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return TerminalProductsResponse
*/

func (a *TerminalOrdersCompanyLevelApi) ListTerminalProducts(r TerminalOrdersCompanyLevelApiListTerminalProductsConfig) (TerminalProductsResponse, *_nethttp.Response, error) {
	res := &TerminalProductsResponse{}
	path := "/companies/{companyId}/terminalProducts"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	queryString := url.Values{}
	if r.country != nil {
		common.ParameterAddToQuery(queryString, "country", r.country, "")
	}
	if r.terminalModelId != nil {
		common.ParameterAddToQuery(queryString, "terminalModelId", r.terminalModelId, "")
	}
	if r.offset != nil {
		common.ParameterAddToQuery(queryString, "offset", r.offset, "")
	}
	if r.limit != nil {
		common.ParameterAddToQuery(queryString, "limit", r.limit, "")
	}
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path+"?"+queryString.Encode(), []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalOrdersCompanyLevelApiUpdateOrderConfig struct {
	ctx                  context.Context
	companyId            string
	orderId              string
	terminalOrderRequest *TerminalOrderRequest
}

func (r TerminalOrdersCompanyLevelApiUpdateOrderConfig) TerminalOrderRequest(terminalOrderRequest TerminalOrderRequest) TerminalOrdersCompanyLevelApiUpdateOrderConfig {
	r.terminalOrderRequest = &terminalOrderRequest
	return r
}

/*
UpdateOrder Update an order

Updates the terminal products order identified in the path.
Updating is only possible while the order has the status **Placed**.

The request body only needs to contain what you want to change.
However, to update the products in the `items` array, you must provide the entire array. For example, if the array has three items:
 To remove one item, the array must include the remaining two items. Or to add one item, the array must include all four items.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The unique identifier of the company account.
 @param orderId The unique identifier of the order.
 @return TerminalOrdersCompanyLevelApiUpdateOrderConfig
*/
func (a *TerminalOrdersCompanyLevelApi) UpdateOrderConfig(ctx context.Context, companyId string, orderId string) TerminalOrdersCompanyLevelApiUpdateOrderConfig {
	return TerminalOrdersCompanyLevelApiUpdateOrderConfig{
		ctx:       ctx,
		companyId: companyId,
		orderId:   orderId,
	}
}

/*
Update an order
Updates the terminal products order identified in the path. Updating is only possible while the order has the status **Placed**.  The request body only needs to contain what you want to change.  However, to update the products in the &#x60;items&#x60; array, you must provide the entire array. For example, if the array has three items:  To remove one item, the array must include the remaining two items. Or to add one item, the array must include all four items.  To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal ordering read and write
 * @param companyId The unique identifier of the company account.
 * @param orderId The unique identifier of the order.
 * @param req TerminalOrderRequest - reference of TerminalOrderRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return TerminalOrder
*/

func (a *TerminalOrdersCompanyLevelApi) UpdateOrder(r TerminalOrdersCompanyLevelApiUpdateOrderConfig) (TerminalOrder, *_nethttp.Response, error) {
	res := &TerminalOrder{}
	path := "/companies/{companyId}/terminalOrders/{orderId}"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"orderId"+"}", url.PathEscape(common.ParameterValueToString(r.orderId, "orderId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPatch, r.terminalOrderRequest, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}