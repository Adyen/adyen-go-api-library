/*
Adyen Balance Control API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balancecontrol

import (
	"context"
	"net/http"
	"net/url"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// GeneralApi service
type GeneralApi common.Service

// All parameters accepted by GeneralApi.BalanceTransfer
type GeneralApiBalanceTransferInput struct {
	balanceTransferRequest *BalanceTransferRequest
}

func (r GeneralApiBalanceTransferInput) BalanceTransferRequest(balanceTransferRequest BalanceTransferRequest) GeneralApiBalanceTransferInput {
	r.balanceTransferRequest = &balanceTransferRequest
	return r
}

/*
Prepare a request for BalanceTransfer

@return GeneralApiBalanceTransferInput

Deprecated since Adyen Balance Control API v1
*/
func (a *GeneralApi) BalanceTransferInput() GeneralApiBalanceTransferInput {
	return GeneralApiBalanceTransferInput{}
}

/*
BalanceTransfer Start a balance transfer

Starts a balance transfer request between merchant accounts. The following conditions must be met before you can successfully transfer balances:

* The source and destination merchant accounts must be under the same company account and legal entity.

* The source merchant account must have sufficient funds.

* The source and destination merchant accounts must have at least one common processing currency.

When sending multiple API requests with the same source and destination merchant accounts, send the requests sequentially and *not* in parallel. Some requests may not be processed if the requests are sent in parallel.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiBalanceTransferInput - Request parameters, see BalanceTransferInput
@return BalanceTransferResponse, *http.Response, error

Deprecated since Adyen Balance Control API v1
*/
func (a *GeneralApi) BalanceTransfer(ctx context.Context, r GeneralApiBalanceTransferInput) (BalanceTransferResponse, *http.Response, error) {
	res := &BalanceTransferResponse{}
	path := "/balanceTransfer"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.balanceTransferRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
