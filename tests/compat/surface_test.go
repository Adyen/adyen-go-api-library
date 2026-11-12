// Package compat freezes the public contract that v21 consumers depend on, so that
// the OpenAPI Generator v7 migration surfaces every source or wire change explicitly.
// A failure here is not automatically a bug: it is a change that must be recorded in
// the breaking-change ledger and the migration notes before the guard is updated.
package compat

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/adyen/adyen-go-api-library/v21/src/acswebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/adyen"
	"github.com/adyen/adyen-go-api-library/v21/src/balancecontrol"
	"github.com/adyen/adyen-go-api-library/v21/src/balanceplatform"
	"github.com/adyen/adyen-go-api-library/v21/src/balancewebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/binlookup"
	"github.com/adyen/adyen-go-api-library/v21/src/checkout"
	"github.com/adyen/adyen-go-api-library/v21/src/common"
	"github.com/adyen/adyen-go-api-library/v21/src/configurationwebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/dataprotection"
	"github.com/adyen/adyen-go-api-library/v21/src/disputes"
	"github.com/adyen/adyen-go-api-library/v21/src/hmacvalidator"
	"github.com/adyen/adyen-go-api-library/v21/src/legalentity"
	"github.com/adyen/adyen-go-api-library/v21/src/management"
	"github.com/adyen/adyen-go-api-library/v21/src/managementwebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/payments"
	"github.com/adyen/adyen-go-api-library/v21/src/paymentsapp"
	"github.com/adyen/adyen-go-api-library/v21/src/payout"
	"github.com/adyen/adyen-go-api-library/v21/src/posterminalmanagement"
	"github.com/adyen/adyen-go-api-library/v21/src/recurring"
	"github.com/adyen/adyen-go-api-library/v21/src/reportwebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/storedvalue"
	"github.com/adyen/adyen-go-api-library/v21/src/transactionwebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/transfers"
	"github.com/adyen/adyen-go-api-library/v21/src/transferwebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/webhook"
	"github.com/stretchr/testify/assert"
)

// Client entry points and service accessors.
var (
	_ func(*common.Config) *adyen.APIClient                         = adyen.NewClient
	_ func(*adyen.APIClient, common.Environment, string)            = (*adyen.APIClient).SetEnvironment
	_ func(*adyen.APIClient) *common.Config                         = (*adyen.APIClient).GetConfig
	_ func(*adyen.APIClient) *paymentsapp.PaymentsAppApi            = (*adyen.APIClient).PaymentsApp
	_ func(*adyen.APIClient) *checkout.APIClient                    = (*adyen.APIClient).Checkout
	_ func(*adyen.APIClient) *payments.APIClient                    = (*adyen.APIClient).Payments
	_ func(*adyen.APIClient) *payout.APIClient                      = (*adyen.APIClient).Payout
	_ func(*adyen.APIClient) *recurring.GeneralApi                  = (*adyen.APIClient).Recurring
	_ func(*adyen.APIClient) *binlookup.GeneralApi                  = (*adyen.APIClient).BinLookup
	_ func(*adyen.APIClient) *balancecontrol.GeneralApi             = (*adyen.APIClient).BalanceControl
	_ func(*adyen.APIClient) *storedvalue.GeneralApi                = (*adyen.APIClient).StoredValue
	_ func(*adyen.APIClient) *posterminalmanagement.GeneralApi      = (*adyen.APIClient).PosTerminalManagement
	_ func(*adyen.APIClient) *transfers.APIClient                   = (*adyen.APIClient).Transfers
	_ func(*adyen.APIClient) *balanceplatform.APIClient             = (*adyen.APIClient).BalancePlatform
	_ func(*adyen.APIClient) *management.APIClient                  = (*adyen.APIClient).Management
	_ func(*adyen.APIClient) *legalentity.APIClient                 = (*adyen.APIClient).LegalEntity
	_ func(*adyen.APIClient) *disputes.GeneralApi                   = (*adyen.APIClient).Disputes
	_ func(*adyen.APIClient) *dataprotection.GeneralApi             = (*adyen.APIClient).DataProtection
	_ func(*checkout.APIClient) *checkout.PaymentsApi               = func(c *checkout.APIClient) *checkout.PaymentsApi { return c.PaymentsApi }
	_ func(*transfers.APIClient) *transfers.TransfersApi            = func(c *transfers.APIClient) *transfers.TransfersApi { return c.TransfersApi }
	_ func(*checkout.PaymentsApi) func() string                     = func(a *checkout.PaymentsApi) func() string { return a.BasePath }
	_ func(*checkout.PaymentsApi) *common.Client                    = func(a *checkout.PaymentsApi) *common.Client { return a.Client }
	_ func(*recurring.GeneralApi) func() string                     = func(a *recurring.GeneralApi) func() string { return a.BasePath }
	_ func(*paymentsapp.PaymentsAppApi) func() string               = func(a *paymentsapp.PaymentsAppApi) func() string { return a.BasePath }
	_ func(*checkout.PaymentsApi) checkout.PaymentsApiPaymentsInput = (*checkout.PaymentsApi).PaymentsInput
)

// Request builders: value-receiver fluent setters and (response, *http.Response, error) results.
var (
	_ func(checkout.PaymentsApiPaymentsInput, string) checkout.PaymentsApiPaymentsInput                                                    = checkout.PaymentsApiPaymentsInput.IdempotencyKey
	_ func(checkout.PaymentsApiPaymentsInput, checkout.PaymentRequest) checkout.PaymentsApiPaymentsInput                                   = checkout.PaymentsApiPaymentsInput.PaymentRequest
	_ func(*checkout.PaymentsApi, context.Context, checkout.PaymentsApiPaymentsInput) (checkout.PaymentResponse, *http.Response, error)    = (*checkout.PaymentsApi).Payments
	_ func(transfers.TransfersApiTransferFundsInput, string) transfers.TransfersApiTransferFundsInput                                      = transfers.TransfersApiTransferFundsInput.IdempotencyKey
	_ func(transfers.TransfersApiTransferFundsInput, transfers.TransferInfo) transfers.TransfersApiTransferFundsInput                      = transfers.TransfersApiTransferFundsInput.TransferInfo
	_ func(*transfers.TransfersApi, context.Context, transfers.TransfersApiTransferFundsInput) (transfers.Transfer, *http.Response, error) = (*transfers.TransfersApi).TransferFunds
	_ func(*transfers.TransfersApi, context.Context, transfers.TransfersApiApproveInitiatedTransfersInput) (*http.Response, error)         = (*transfers.TransfersApi).ApproveInitiatedTransfers
	_ func(checkout.Amount, string, checkout.CheckoutPaymentMethod, string, string) *checkout.PaymentRequest                               = checkout.NewPaymentRequest
	_ func(*checkout.CardDetails) checkout.CheckoutPaymentMethod                                                                           = checkout.CardDetailsAsCheckoutPaymentMethod
)

// Shared helpers, errors and nullable wrappers.
var (
	_ error                                                                   = common.APIError{}
	_ error                                                                   = common.RestServiceError{}
	_ func([]byte, string) common.APIError                                    = common.NewAPIError
	_ func(context.Context, string) context.Context                           = common.WithIdempotencyKey
	_ func(string) *string                                                    = common.PtrString
	_ func(*string) *common.NullableString                                    = common.NewNullableString
	_ func([]byte) interface{ Decode(interface{}) error }                     = func(b []byte) interface{ Decode(interface{}) error } { return common.NewStrictDecoder(b) }
	_ func(common.APIError) ([]byte, string, float64, string, string, string) = func(e common.APIError) ([]byte, string, float64, string, string, string) {
		return e.RawBody, e.Err, e.Status, e.Message, e.Code, e.Type
	}
	_ func(common.RestServiceError) (string, string, *string, []common.InvalidField, *string, int32, string, string) = func(e common.RestServiceError) (string, string, *string, []common.InvalidField, *string, int32, string, string) {
		return e.Detail, e.ErrorCode, e.Instance, e.InvalidFields, e.RequestId, e.Status, e.Title, e.Type
	}
	_ func(checkout.PaymentResponse) (*string, *string, *map[string]string) = func(r checkout.PaymentResponse) (*string, *string, *map[string]string) {
		return r.PspReference, r.ResultCode, r.AdditionalData
	}
	_ func(management.PayAtTable) common.NullableString = func(p management.PayAtTable) common.NullableString { return p.PaymentInstrument }
)

// Configuration fields and their types.
var _ = common.Config{
	Username:                      "",
	Password:                      "",
	MerchantAccount:               "",
	Environment:                   common.TestEnv,
	ApplicationName:               "",
	Endpoint:                      "",
	ApiKey:                        "",
	ConnectionTimeoutMillis:       time.Duration(0),
	CertificatePath:               "",
	DisputesEndpoint:              "",
	BalancePlatformEndpoint:       "",
	ManagementEndpoint:            "",
	PaymentsAppEndpoint:           "",
	LegalEntityEndpoint:           "",
	TransfersEndpoint:             "",
	PosTerminalManagementEndpoint: "",
	DataProtectionEndpoint:        "",
	CheckoutEndpoint:              "",
	TerminalApiCloudEndpoint:      "",
	TerminalApiLocalEndpoint:      "",
	TerminalCertificatePath:       "",
	LiveEndpointURLPrefix:         "",
	DefaultHeader:                 map[string]string{},
	Debug:                         false,
	UserAgent:                     "",
	HTTPClient:                    (*http.Client)(nil),
	Log3XXError:                   false,
	Log4XXError:                   false,
	Log5XXError:                   false,
}

// Webhooks: the classic notification keeps a pointer-to-map AdditionalData
// (v21.2.1 was retracted for changing it), and typed handlers keep their names.
var (
	_ func(string) (*webhook.Webhook, error)                                               = webhook.HandleRequest
	_ func(*webhook.Webhook) []*webhook.NotificationRequestItem                            = (*webhook.Webhook).GetNotificationItems
	_ func(webhook.NotificationRequestItem) *map[string]interface{}                        = func(n webhook.NotificationRequestItem) *map[string]interface{} { return n.AdditionalData }
	_ func(webhook.Webhook) (string, *[]webhook.NotificationItem)                          = func(w webhook.Webhook) (string, *[]webhook.NotificationItem) { return w.Live, w.NotificationItems }
	_ func(webhook.NotificationRequestItem) (webhook.Amount, *time.Time, string, []string) = func(n webhook.NotificationRequestItem) (webhook.Amount, *time.Time, string, []string) {
		return n.Amount, n.EventDate, n.Success, n.Operations
	}
	_ func(webhook.NotificationRequestItem, string) bool = hmacvalidator.ValidateHmac
	_ func(interface{}, string) (string, error)          = hmacvalidator.CalculateHmac
	_ func(string, string, string) bool                  = hmacvalidator.ValidateHmacPayload
	_ func(interface{}) string                           = hmacvalidator.GetDataToSign

	_ func(string) (*acswebhook.AuthenticationNotificationRequest, error)                          = acswebhook.HandleAuthenticationNotificationRequest
	_ func(string) (*balancewebhook.BalanceAccountBalanceNotificationRequest, error)               = balancewebhook.HandleBalanceAccountBalanceNotificationRequest
	_ func(string) (*configurationwebhook.AccountHolderNotificationRequest, error)                 = configurationwebhook.HandleAccountHolderNotificationRequest
	_ func(string) (*configurationwebhook.BalanceAccountNotificationRequest, error)                = configurationwebhook.HandleBalanceAccountNotificationRequest
	_ func(string) (*configurationwebhook.CardOrderNotificationRequest, error)                     = configurationwebhook.HandleCardOrderNotificationRequest
	_ func(string) (*configurationwebhook.PaymentNotificationRequest, error)                       = configurationwebhook.HandlePaymentNotificationRequest
	_ func(string) (*configurationwebhook.SweepConfigurationNotificationRequest, error)            = configurationwebhook.HandleSweepConfigurationNotificationRequest
	_ func(string) (*managementwebhook.MerchantCreatedNotificationRequest, error)                  = managementwebhook.HandleMerchantCreatedNotificationRequest
	_ func(string) (*managementwebhook.MerchantUpdatedNotificationRequest, error)                  = managementwebhook.HandleMerchantUpdatedNotificationRequest
	_ func(string) (*managementwebhook.PaymentMethodCreatedNotificationRequest, error)             = managementwebhook.HandlePaymentMethodCreatedNotificationRequest
	_ func(string) (*managementwebhook.PaymentMethodRequestRemovedNotificationRequest, error)      = managementwebhook.HandlePaymentMethodRequestRemovedNotificationRequest
	_ func(string) (*managementwebhook.PaymentMethodScheduledForRemovalNotificationRequest, error) = managementwebhook.HandlePaymentMethodScheduledForRemovalNotificationRequest
	_ func(string) (*reportwebhook.ReportNotificationRequest, error)                               = reportwebhook.HandleReportNotificationRequest
	_ func(string) (*transactionwebhook.TransactionNotificationRequestV4, error)                   = transactionwebhook.HandleTransactionNotificationRequestV4
	_ func(string) (*transferwebhook.TransferNotificationRequest, error)                           = transferwebhook.HandleTransferNotificationRequest
)

// Changing any of these values changes the URL every request is sent to. Bumps are
// expected during the migration (e.g. Checkout v71 -> v72) but must be deliberate.
func TestAPIVersionConstants(t *testing.T) {
	assert.Equal(t, map[string]string{
		"Payment":               "v68",
		"PaymentsApp":           "v1",
		"Recurring":             "v68",
		"Checkout":              "v71",
		"BinLookup":             "v54",
		"BalanceControl":        "v1",
		"Disputes":              "v30",
		"StoredValue":           "v46",
		"BalancePlatform":       "v2",
		"Transfers":             "v4",
		"Management":            "v3",
		"LegalEntity":           "v4",
		"PosTerminalManagement": "v1",
		"DataProtection":        "v1",
	}, map[string]string{
		"Payment":               adyen.PaymentAPIVersion,
		"PaymentsApp":           adyen.PaymentsAppAPIVersion,
		"Recurring":             adyen.RecurringAPIVersion,
		"Checkout":              adyen.CheckoutAPIVersion,
		"BinLookup":             adyen.BinLookupAPIVersion,
		"BalanceControl":        adyen.BalanceControlAPIVersion,
		"Disputes":              adyen.DisputesAPIVersion,
		"StoredValue":           adyen.StoredValueAPIVersion,
		"BalancePlatform":       adyen.BalancePlatformAPIVersion,
		"Transfers":             adyen.TransfersAPIVersion,
		"Management":            adyen.ManagementAPIVersion,
		"LegalEntity":           adyen.LegalEntityAPIVersion,
		"PosTerminalManagement": adyen.PosTerminalManagementAPIVersion,
		"DataProtection":        adyen.DataProtectionAPIVersion,
	})
	assert.Equal(t, "adyen-go-api-library", common.LibName)
	assert.Equal(t, common.Environment("TEST"), common.TestEnv)
	assert.Equal(t, common.Environment("LIVE"), common.LiveEnv)
}
