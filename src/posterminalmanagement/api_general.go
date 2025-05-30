/*
POS Terminal Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package posterminalmanagement

import (
	"context"
	"net/http"
	"net/url"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// GeneralApi service
type GeneralApi common.Service

// All parameters accepted by GeneralApi.AssignTerminals
type GeneralApiAssignTerminalsInput struct {
	assignTerminalsRequest *AssignTerminalsRequest
}

func (r GeneralApiAssignTerminalsInput) AssignTerminalsRequest(assignTerminalsRequest AssignTerminalsRequest) GeneralApiAssignTerminalsInput {
	r.assignTerminalsRequest = &assignTerminalsRequest
	return r
}

/*
Prepare a request for AssignTerminals

@return GeneralApiAssignTerminalsInput

Deprecated since POS Terminal Management API v1
Use [Management API](https://docs.adyen.com/api-explorer/Management/latest/overview).
*/
func (a *GeneralApi) AssignTerminalsInput() GeneralApiAssignTerminalsInput {
	return GeneralApiAssignTerminalsInput{}
}

/*
AssignTerminals Assign terminals

Assigns one or more payment terminals to a merchant account or a store. You can also use this endpoint to reassign terminals between merchant accounts or stores, and to unassign a terminal and return it to company inventory.

>From January 1, 2025 POS Terminal Management API is deprecated and support stops on April 1, 2025. To automate the management of your terminal fleet, use our [Management API](https://docs.adyen.com/api-explorer/Management/latest/overview).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiAssignTerminalsInput - Request parameters, see AssignTerminalsInput
@return AssignTerminalsResponse, *http.Response, error

Deprecated since POS Terminal Management API v1
Use [Management API](https://docs.adyen.com/api-explorer/Management/latest/overview).
*/
func (a *GeneralApi) AssignTerminals(ctx context.Context, r GeneralApiAssignTerminalsInput) (AssignTerminalsResponse, *http.Response, error) {
	res := &AssignTerminalsResponse{}
	path := "/assignTerminals"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.assignTerminalsRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.FindTerminal
type GeneralApiFindTerminalInput struct {
	findTerminalRequest *FindTerminalRequest
}

func (r GeneralApiFindTerminalInput) FindTerminalRequest(findTerminalRequest FindTerminalRequest) GeneralApiFindTerminalInput {
	r.findTerminalRequest = &findTerminalRequest
	return r
}

/*
Prepare a request for FindTerminal

@return GeneralApiFindTerminalInput

Deprecated since POS Terminal Management API v1
Use [Management API](https://docs.adyen.com/api-explorer/Management/latest/overview).
*/
func (a *GeneralApi) FindTerminalInput() GeneralApiFindTerminalInput {
	return GeneralApiFindTerminalInput{}
}

/*
FindTerminal Get the account or store of a terminal

Returns the company account, merchant account, or store that a payment terminal is assigned to.

>From January 1, 2025 POS Terminal Management API is deprecated and support stops on April 1, 2025. To automate the management of your terminal fleet, use our [Management API](https://docs.adyen.com/api-explorer/Management/latest/overview).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiFindTerminalInput - Request parameters, see FindTerminalInput
@return FindTerminalResponse, *http.Response, error

Deprecated since POS Terminal Management API v1
Use [Management API](https://docs.adyen.com/api-explorer/Management/latest/overview).
*/
func (a *GeneralApi) FindTerminal(ctx context.Context, r GeneralApiFindTerminalInput) (FindTerminalResponse, *http.Response, error) {
	res := &FindTerminalResponse{}
	path := "/findTerminal"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.findTerminalRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.GetStoresUnderAccount
type GeneralApiGetStoresUnderAccountInput struct {
	getStoresUnderAccountRequest *GetStoresUnderAccountRequest
}

func (r GeneralApiGetStoresUnderAccountInput) GetStoresUnderAccountRequest(getStoresUnderAccountRequest GetStoresUnderAccountRequest) GeneralApiGetStoresUnderAccountInput {
	r.getStoresUnderAccountRequest = &getStoresUnderAccountRequest
	return r
}

/*
Prepare a request for GetStoresUnderAccount

@return GeneralApiGetStoresUnderAccountInput

Deprecated since POS Terminal Management API v1
Use [Management API](https://docs.adyen.com/api-explorer/Management/latest/overview).
*/
func (a *GeneralApi) GetStoresUnderAccountInput() GeneralApiGetStoresUnderAccountInput {
	return GeneralApiGetStoresUnderAccountInput{}
}

/*
GetStoresUnderAccount Get the stores of an account

Returns a list of stores associated with a company account or a merchant account, including the status of each store.

>From January 1, 2025 POS Terminal Management API is deprecated and support stops on April 1, 2025. To automate the management of your terminal fleet, use our [Management API](https://docs.adyen.com/api-explorer/Management/latest/overview).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiGetStoresUnderAccountInput - Request parameters, see GetStoresUnderAccountInput
@return GetStoresUnderAccountResponse, *http.Response, error

Deprecated since POS Terminal Management API v1
Use [Management API](https://docs.adyen.com/api-explorer/Management/latest/overview).
*/
func (a *GeneralApi) GetStoresUnderAccount(ctx context.Context, r GeneralApiGetStoresUnderAccountInput) (GetStoresUnderAccountResponse, *http.Response, error) {
	res := &GetStoresUnderAccountResponse{}
	path := "/getStoresUnderAccount"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.getStoresUnderAccountRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.GetTerminalDetails
type GeneralApiGetTerminalDetailsInput struct {
	getTerminalDetailsRequest *GetTerminalDetailsRequest
}

func (r GeneralApiGetTerminalDetailsInput) GetTerminalDetailsRequest(getTerminalDetailsRequest GetTerminalDetailsRequest) GeneralApiGetTerminalDetailsInput {
	r.getTerminalDetailsRequest = &getTerminalDetailsRequest
	return r
}

/*
Prepare a request for GetTerminalDetails

@return GeneralApiGetTerminalDetailsInput

Deprecated since POS Terminal Management API v1
Use [Management API](https://docs.adyen.com/api-explorer/Management/latest/overview).
*/
func (a *GeneralApi) GetTerminalDetailsInput() GeneralApiGetTerminalDetailsInput {
	return GeneralApiGetTerminalDetailsInput{}
}

/*
GetTerminalDetails Get the details of a terminal

Returns the details of a payment terminal, including where the terminal is assigned to. The response returns the same details that are provided in the terminal list in your Customer Area and in the Terminal Fleet report.

>From January 1, 2025 POS Terminal Management API is deprecated and support stops on April 1, 2025. To automate the management of your terminal fleet, use our [Management API](https://docs.adyen.com/api-explorer/Management/latest/overview).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiGetTerminalDetailsInput - Request parameters, see GetTerminalDetailsInput
@return GetTerminalDetailsResponse, *http.Response, error

Deprecated since POS Terminal Management API v1
Use [Management API](https://docs.adyen.com/api-explorer/Management/latest/overview).
*/
func (a *GeneralApi) GetTerminalDetails(ctx context.Context, r GeneralApiGetTerminalDetailsInput) (GetTerminalDetailsResponse, *http.Response, error) {
	res := &GetTerminalDetailsResponse{}
	path := "/getTerminalDetails"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.getTerminalDetailsRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.GetTerminalsUnderAccount
type GeneralApiGetTerminalsUnderAccountInput struct {
	getTerminalsUnderAccountRequest *GetTerminalsUnderAccountRequest
}

func (r GeneralApiGetTerminalsUnderAccountInput) GetTerminalsUnderAccountRequest(getTerminalsUnderAccountRequest GetTerminalsUnderAccountRequest) GeneralApiGetTerminalsUnderAccountInput {
	r.getTerminalsUnderAccountRequest = &getTerminalsUnderAccountRequest
	return r
}

/*
Prepare a request for GetTerminalsUnderAccount

@return GeneralApiGetTerminalsUnderAccountInput

Deprecated since POS Terminal Management API v1
Use [Management API](https://docs.adyen.com/api-explorer/Management/latest/overview).
*/
func (a *GeneralApi) GetTerminalsUnderAccountInput() GeneralApiGetTerminalsUnderAccountInput {
	return GeneralApiGetTerminalsUnderAccountInput{}
}

/*
GetTerminalsUnderAccount Get the list of terminals

Returns a list of payment terminals associated with a company account, merchant account, or store. The response shows whether the terminals are in the inventory, or in-store (ready for boarding or already boarded).

>From January 1, 2025 POS Terminal Management API is deprecated and support stops on April 1, 2025. To automate the management of your terminal fleet, use our [Management API](https://docs.adyen.com/api-explorer/Management/latest/overview).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiGetTerminalsUnderAccountInput - Request parameters, see GetTerminalsUnderAccountInput
@return GetTerminalsUnderAccountResponse, *http.Response, error

Deprecated since POS Terminal Management API v1
Use [Management API](https://docs.adyen.com/api-explorer/Management/latest/overview).
*/
func (a *GeneralApi) GetTerminalsUnderAccount(ctx context.Context, r GeneralApiGetTerminalsUnderAccountInput) (GetTerminalsUnderAccountResponse, *http.Response, error) {
	res := &GetTerminalsUnderAccountResponse{}
	path := "/getTerminalsUnderAccount"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.getTerminalsUnderAccountRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
