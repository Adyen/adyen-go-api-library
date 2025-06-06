/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"context"
	"net/http"
	"net/url"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// ModificationsApi service
type ModificationsApi common.Service

// All parameters accepted by ModificationsApi.AdjustAuthorisation
type ModificationsApiAdjustAuthorisationInput struct {
	adjustAuthorisationRequest *AdjustAuthorisationRequest
}

func (r ModificationsApiAdjustAuthorisationInput) AdjustAuthorisationRequest(adjustAuthorisationRequest AdjustAuthorisationRequest) ModificationsApiAdjustAuthorisationInput {
	r.adjustAuthorisationRequest = &adjustAuthorisationRequest
	return r
}

/*
Prepare a request for AdjustAuthorisation

@return ModificationsApiAdjustAuthorisationInput
*/
func (a *ModificationsApi) AdjustAuthorisationInput() ModificationsApiAdjustAuthorisationInput {
	return ModificationsApiAdjustAuthorisationInput{}
}

/*
AdjustAuthorisation Change the authorised amount

Allows you to increase or decrease the authorised amount after the initial authorisation has taken place. This functionality enables for example tipping, improving the chances your authorisation will be valid, or charging the shopper when they have already left the merchant premises.

> This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce).
> If you have a [newer integration](https://docs.adyen.com/online-payments), and are doing:
> * [Asynchronous adjustments](https://docs.adyen.com/online-payments/adjust-authorisation#asynchronous-or-synchronous-adjustment), use the [`/payments/{paymentPspReference}/amountUpdates`](https://docs.adyen.com/api-explorer/#/CheckoutService/v67/post/payments/{paymentPspReference}/amountUpdates) endpoint on Checkout API.
> * [Synchronous adjustments](https://docs.adyen.com/online-payments/adjust-authorisation#asynchronous-or-synchronous-adjustment), use this endpoint.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ModificationsApiAdjustAuthorisationInput - Request parameters, see AdjustAuthorisationInput
@return ModificationResult, *http.Response, error
*/
func (a *ModificationsApi) AdjustAuthorisation(ctx context.Context, r ModificationsApiAdjustAuthorisationInput) (ModificationResult, *http.Response, error) {
	res := &ModificationResult{}
	path := "/adjustAuthorisation"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.adjustAuthorisationRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by ModificationsApi.Cancel
type ModificationsApiCancelInput struct {
	cancelRequest *CancelRequest
}

func (r ModificationsApiCancelInput) CancelRequest(cancelRequest CancelRequest) ModificationsApiCancelInput {
	r.cancelRequest = &cancelRequest
	return r
}

/*
Prepare a request for Cancel

@return ModificationsApiCancelInput
*/
func (a *ModificationsApi) CancelInput() ModificationsApiCancelInput {
	return ModificationsApiCancelInput{}
}

/*
Cancel Cancel an authorisation

Cancels the authorisation hold on a payment, returning a unique reference for this request. You can cancel payments after authorisation only for payment methods that support distinct authorisations and captures.

For more information, refer to [Cancel](https://docs.adyen.com/online-payments/classic-integrations/modify-payments/cancel).

> This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [`/payments/{paymentPspReference}/cancels`](https://docs.adyen.com/api-explorer/#/CheckoutService/payments/{paymentPspReference}/cancels) endpoint under Checkout API instead.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ModificationsApiCancelInput - Request parameters, see CancelInput
@return ModificationResult, *http.Response, error
*/
func (a *ModificationsApi) Cancel(ctx context.Context, r ModificationsApiCancelInput) (ModificationResult, *http.Response, error) {
	res := &ModificationResult{}
	path := "/cancel"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.cancelRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by ModificationsApi.CancelOrRefund
type ModificationsApiCancelOrRefundInput struct {
	cancelOrRefundRequest *CancelOrRefundRequest
}

func (r ModificationsApiCancelOrRefundInput) CancelOrRefundRequest(cancelOrRefundRequest CancelOrRefundRequest) ModificationsApiCancelOrRefundInput {
	r.cancelOrRefundRequest = &cancelOrRefundRequest
	return r
}

/*
Prepare a request for CancelOrRefund

@return ModificationsApiCancelOrRefundInput
*/
func (a *ModificationsApi) CancelOrRefundInput() ModificationsApiCancelOrRefundInput {
	return ModificationsApiCancelOrRefundInput{}
}

/*
CancelOrRefund Cancel or refund a payment

Cancels a payment if it has not been captured yet, or refunds it if it has already been captured. This is useful when it is not certain if the payment has been captured or not (for example, when using auto-capture).

Do not use this endpoint for payments that involve:

  - [Multiple partial captures](https://docs.adyen.com/online-payments/capture).

  - [Split data](https://docs.adyen.com/classic-platforms/processing-payments#providing-split-information) either at time of payment or capture for Adyen for Platforms.

    Instead, check if the payment has been captured and make a corresponding [`/refund`](https://docs.adyen.com/api-explorer/#/Payment/refund) or [`/cancel`](https://docs.adyen.com/api-explorer/#/Payment/cancel) call.

For more information, refer to [Cancel or refund](https://docs.adyen.com/online-payments/classic-integrations/modify-payments/cancel-or-refund).

> This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [`/payments/{paymentPspReference}/reversals`](https://docs.adyen.com/api-explorer/#/CheckoutService/payments/{paymentPspReference}/reversals) endpoint under Checkout API instead.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ModificationsApiCancelOrRefundInput - Request parameters, see CancelOrRefundInput
@return ModificationResult, *http.Response, error
*/
func (a *ModificationsApi) CancelOrRefund(ctx context.Context, r ModificationsApiCancelOrRefundInput) (ModificationResult, *http.Response, error) {
	res := &ModificationResult{}
	path := "/cancelOrRefund"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.cancelOrRefundRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by ModificationsApi.Capture
type ModificationsApiCaptureInput struct {
	captureRequest *CaptureRequest
}

func (r ModificationsApiCaptureInput) CaptureRequest(captureRequest CaptureRequest) ModificationsApiCaptureInput {
	r.captureRequest = &captureRequest
	return r
}

/*
Prepare a request for Capture

@return ModificationsApiCaptureInput
*/
func (a *ModificationsApi) CaptureInput() ModificationsApiCaptureInput {
	return ModificationsApiCaptureInput{}
}

/*
Capture Capture an authorisation

Captures the authorisation hold on a payment, returning a unique reference for this request. Usually the full authorisation amount is captured, however it's also possible to capture a smaller amount, which results in cancelling the remaining authorisation balance.

Payment methods that are captured automatically after authorisation don't need to be captured. However, submitting a capture request on these transactions will not result in double charges. If immediate or delayed auto-capture is enabled, calling the capture method is not necessary.

For more information refer to [Capture](https://docs.adyen.com/online-payments/classic-integrations/modify-payments/capture).

> This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [`/payments/{paymentPspReference}/captures`](https://docs.adyen.com/api-explorer/#/CheckoutService/v67/post/payments/{paymentPspReference}/captures) endpoint on Checkout API instead.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ModificationsApiCaptureInput - Request parameters, see CaptureInput
@return ModificationResult, *http.Response, error
*/
func (a *ModificationsApi) Capture(ctx context.Context, r ModificationsApiCaptureInput) (ModificationResult, *http.Response, error) {
	res := &ModificationResult{}
	path := "/capture"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.captureRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by ModificationsApi.Donate
type ModificationsApiDonateInput struct {
	donationRequest *DonationRequest
}

func (r ModificationsApiDonateInput) DonationRequest(donationRequest DonationRequest) ModificationsApiDonateInput {
	r.donationRequest = &donationRequest
	return r
}

/*
Prepare a request for Donate

@return ModificationsApiDonateInput

Deprecated
*/
func (a *ModificationsApi) DonateInput() ModificationsApiDonateInput {
	return ModificationsApiDonateInput{}
}

/*
Donate Create a donation

Schedules a new payment to be created (including a new authorisation request) for the specified donation using the payment details of the original payment.

> This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [`/donations`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/donations) endpoint under Checkout API instead.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ModificationsApiDonateInput - Request parameters, see DonateInput
@return ModificationResult, *http.Response, error

Deprecated
*/
func (a *ModificationsApi) Donate(ctx context.Context, r ModificationsApiDonateInput) (ModificationResult, *http.Response, error) {
	res := &ModificationResult{}
	path := "/donate"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.donationRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by ModificationsApi.Refund
type ModificationsApiRefundInput struct {
	refundRequest *RefundRequest
}

func (r ModificationsApiRefundInput) RefundRequest(refundRequest RefundRequest) ModificationsApiRefundInput {
	r.refundRequest = &refundRequest
	return r
}

/*
Prepare a request for Refund

@return ModificationsApiRefundInput
*/
func (a *ModificationsApi) RefundInput() ModificationsApiRefundInput {
	return ModificationsApiRefundInput{}
}

/*
Refund Refund a captured payment

Refunds a payment that has previously been captured, returning a unique reference for this request. Refunding can be done on the full captured amount or a partial amount. Multiple (partial) refunds will be accepted as long as their sum doesn't exceed the captured amount. Payments which have been authorised, but not captured, cannot be refunded, use the /cancel method instead.

Some payment methods/gateways do not support partial/multiple refunds.
A margin above the captured limit can be configured to cover shipping/handling costs.

For more information, refer to [Refund](https://docs.adyen.com/online-payments/classic-integrations/modify-payments/refund).

> This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [`/payments/{paymentPspReference}/refunds`](https://docs.adyen.com/api-explorer/#/CheckoutService/payments/{paymentPspReference}/refunds) endpoint under Checkout API instead.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ModificationsApiRefundInput - Request parameters, see RefundInput
@return ModificationResult, *http.Response, error
*/
func (a *ModificationsApi) Refund(ctx context.Context, r ModificationsApiRefundInput) (ModificationResult, *http.Response, error) {
	res := &ModificationResult{}
	path := "/refund"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.refundRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by ModificationsApi.TechnicalCancel
type ModificationsApiTechnicalCancelInput struct {
	technicalCancelRequest *TechnicalCancelRequest
}

func (r ModificationsApiTechnicalCancelInput) TechnicalCancelRequest(technicalCancelRequest TechnicalCancelRequest) ModificationsApiTechnicalCancelInput {
	r.technicalCancelRequest = &technicalCancelRequest
	return r
}

/*
Prepare a request for TechnicalCancel

@return ModificationsApiTechnicalCancelInput
*/
func (a *ModificationsApi) TechnicalCancelInput() ModificationsApiTechnicalCancelInput {
	return ModificationsApiTechnicalCancelInput{}
}

/*
TechnicalCancel Cancel an authorisation using your reference

This endpoint allows you to cancel a payment if you do not have the PSP reference of the original payment request available.

In your call, refer to the original payment by using the `reference` that you specified in your payment request.

For more information, see [Technical cancel](https://docs.adyen.com/online-payments/classic-integrations/modify-payments/cancel#technical-cancel).

> This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [`/cancels`](https://docs.adyen.com/api-explorer/#/CheckoutService/cancels) endpoint under Checkout API instead.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ModificationsApiTechnicalCancelInput - Request parameters, see TechnicalCancelInput
@return ModificationResult, *http.Response, error
*/
func (a *ModificationsApi) TechnicalCancel(ctx context.Context, r ModificationsApiTechnicalCancelInput) (ModificationResult, *http.Response, error) {
	res := &ModificationResult{}
	path := "/technicalCancel"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.technicalCancelRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by ModificationsApi.VoidPendingRefund
type ModificationsApiVoidPendingRefundInput struct {
	voidPendingRefundRequest *VoidPendingRefundRequest
}

func (r ModificationsApiVoidPendingRefundInput) VoidPendingRefundRequest(voidPendingRefundRequest VoidPendingRefundRequest) ModificationsApiVoidPendingRefundInput {
	r.voidPendingRefundRequest = &voidPendingRefundRequest
	return r
}

/*
Prepare a request for VoidPendingRefund

@return ModificationsApiVoidPendingRefundInput
*/
func (a *ModificationsApi) VoidPendingRefundInput() ModificationsApiVoidPendingRefundInput {
	return ModificationsApiVoidPendingRefundInput{}
}

/*
VoidPendingRefund Cancel an in-person refund

This endpoint allows you to cancel an unreferenced refund request before it has been completed.

In your call, you can refer to the original refund request either by using the `tenderReference`, or the `pspReference`. We recommend implementing based on the `tenderReference`, as this is generated for both offline and online transactions.

For more information, refer to [Cancel an unreferenced refund](https://docs.adyen.com/point-of-sale/refund-payment/cancel-unreferenced).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ModificationsApiVoidPendingRefundInput - Request parameters, see VoidPendingRefundInput
@return ModificationResult, *http.Response, error
*/
func (a *ModificationsApi) VoidPendingRefund(ctx context.Context, r ModificationsApiVoidPendingRefundInput) (ModificationResult, *http.Response, error) {
	res := &ModificationResult{}
	path := "/voidPendingRefund"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.voidPendingRefundRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
