/*
Adyen Payout API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payout

import (
	"context"
	_nethttp "net/http"
	"net/url"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// InitializationApi InitializationApi service
type InitializationApi common.Service

type InitializationApiStoreDetailConfig struct {
	ctx                context.Context
	storeDetailRequest *StoreDetailRequest
}

func (r InitializationApiStoreDetailConfig) StoreDetailRequest(storeDetailRequest StoreDetailRequest) InitializationApiStoreDetailConfig {
	r.storeDetailRequest = &storeDetailRequest
	return r
}

/*
StoreDetail Store payout details

Stores payment details under the `PAYOUT` recurring contract. These payment details can be used later to submit a payout via the `/submitThirdParty` call.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return InitializationApiStoreDetailConfig
*/
func (a *InitializationApi) StoreDetailConfig(ctx context.Context) InitializationApiStoreDetailConfig {
	return InitializationApiStoreDetailConfig{
		ctx: ctx,
	}
}

/*
Store payout details
Stores payment details under the &#x60;PAYOUT&#x60; recurring contract. These payment details can be used later to submit a payout via the &#x60;/submitThirdParty&#x60; call.
 * @param req StoreDetailRequest - reference of StoreDetailRequest).
 * @param ctxs ...context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return StoreDetailResponse
*/

func (a *InitializationApi) StoreDetail(r InitializationApiStoreDetailConfig) (StoreDetailResponse, *_nethttp.Response, error) {
	res := &StoreDetailResponse{}
	path := "/storeDetail"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		r.ctx,
		a.Client,
		r.storeDetailRequest,
		res,
		_nethttp.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

type InitializationApiStoreDetailAndSubmitThirdPartyConfig struct {
	ctx                         context.Context
	storeDetailAndSubmitRequest *StoreDetailAndSubmitRequest
}

func (r InitializationApiStoreDetailAndSubmitThirdPartyConfig) StoreDetailAndSubmitRequest(storeDetailAndSubmitRequest StoreDetailAndSubmitRequest) InitializationApiStoreDetailAndSubmitThirdPartyConfig {
	r.storeDetailAndSubmitRequest = &storeDetailAndSubmitRequest
	return r
}

/*
StoreDetailAndSubmitThirdParty Store details and submit a payout

Submits a payout and stores its details for subsequent payouts.

The submitted payout must be confirmed or declined either by a reviewer or via `/confirmThirdParty` or `/declineThirdParty` calls.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return InitializationApiStoreDetailAndSubmitThirdPartyConfig
*/
func (a *InitializationApi) StoreDetailAndSubmitThirdPartyConfig(ctx context.Context) InitializationApiStoreDetailAndSubmitThirdPartyConfig {
	return InitializationApiStoreDetailAndSubmitThirdPartyConfig{
		ctx: ctx,
	}
}

/*
Store details and submit a payout
Submits a payout and stores its details for subsequent payouts.  The submitted payout must be confirmed or declined either by a reviewer or via &#x60;/confirmThirdParty&#x60; or &#x60;/declineThirdParty&#x60; calls.
 * @param req StoreDetailAndSubmitRequest - reference of StoreDetailAndSubmitRequest).
 * @param ctxs ...context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return StoreDetailAndSubmitResponse
*/

func (a *InitializationApi) StoreDetailAndSubmitThirdParty(r InitializationApiStoreDetailAndSubmitThirdPartyConfig) (StoreDetailAndSubmitResponse, *_nethttp.Response, error) {
	res := &StoreDetailAndSubmitResponse{}
	path := "/storeDetailAndSubmitThirdParty"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		r.ctx,
		a.Client,
		r.storeDetailAndSubmitRequest,
		res,
		_nethttp.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

type InitializationApiSubmitThirdPartyConfig struct {
	ctx           context.Context
	submitRequest *SubmitRequest
}

func (r InitializationApiSubmitThirdPartyConfig) SubmitRequest(submitRequest SubmitRequest) InitializationApiSubmitThirdPartyConfig {
	r.submitRequest = &submitRequest
	return r
}

/*
SubmitThirdParty Submit a payout

Submits a payout using the previously stored payment details. To store payment details, use the `/storeDetail` API call.

The submitted payout must be confirmed or declined either by a reviewer or via `/confirmThirdParty` or `/declineThirdParty` calls.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return InitializationApiSubmitThirdPartyConfig
*/
func (a *InitializationApi) SubmitThirdPartyConfig(ctx context.Context) InitializationApiSubmitThirdPartyConfig {
	return InitializationApiSubmitThirdPartyConfig{
		ctx: ctx,
	}
}

/*
Submit a payout
Submits a payout using the previously stored payment details. To store payment details, use the &#x60;/storeDetail&#x60; API call.  The submitted payout must be confirmed or declined either by a reviewer or via &#x60;/confirmThirdParty&#x60; or &#x60;/declineThirdParty&#x60; calls.
 * @param req SubmitRequest - reference of SubmitRequest).
 * @param ctxs ...context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return SubmitResponse
*/

func (a *InitializationApi) SubmitThirdParty(r InitializationApiSubmitThirdPartyConfig) (SubmitResponse, *_nethttp.Response, error) {
	res := &SubmitResponse{}
	path := "/submitThirdParty"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		r.ctx,
		a.Client,
		r.submitRequest,
		res,
		_nethttp.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}