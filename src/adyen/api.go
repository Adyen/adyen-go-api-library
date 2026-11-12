/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package adyen

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adyen/adyen-go-api-library/v21/src/paymentsapp"

	"github.com/adyen/adyen-go-api-library/v21/src/dataprotection"

	"github.com/adyen/adyen-go-api-library/v21/src/balancecontrol"
	"github.com/adyen/adyen-go-api-library/v21/src/balanceplatform"
	"github.com/adyen/adyen-go-api-library/v21/src/binlookup"
	"github.com/adyen/adyen-go-api-library/v21/src/capital"
	"github.com/adyen/adyen-go-api-library/v21/src/checkout"
	"github.com/adyen/adyen-go-api-library/v21/src/common"
	"github.com/adyen/adyen-go-api-library/v21/src/disputes"
	"github.com/adyen/adyen-go-api-library/v21/src/documentcollector"
	"github.com/adyen/adyen-go-api-library/v21/src/foreignexchange"
	"github.com/adyen/adyen-go-api-library/v21/src/legalentity"
	"github.com/adyen/adyen-go-api-library/v21/src/management"
	"github.com/adyen/adyen-go-api-library/v21/src/payments"
	"github.com/adyen/adyen-go-api-library/v21/src/payout"
	"github.com/adyen/adyen-go-api-library/v21/src/posmobile"
	"github.com/adyen/adyen-go-api-library/v21/src/posterminalmanagement"
	"github.com/adyen/adyen-go-api-library/v21/src/recurring"
	"github.com/adyen/adyen-go-api-library/v21/src/sessionauthentication"
	"github.com/adyen/adyen-go-api-library/v21/src/softposconfiguration"
	"github.com/adyen/adyen-go-api-library/v21/src/storedvalue"
	"github.com/adyen/adyen-go-api-library/v21/src/transfers"
)

// Constants used for the client API
const (
	EndpointTest                      = "https://pal-test.adyen.com"
	EndpointLive                      = "https://pal-live.adyen.com"
	EndpointLiveSuffix                = "-pal-live.adyenpayments.com"
	CheckoutEndpointTest              = "https://checkout-test.adyen.com/checkout"
	CheckoutEndpointLiveSuffix        = "-checkout-live.adyenpayments.com/checkout"
	TerminalAPIEndpointTest           = "https://terminal-api-test.adyen.com"
	TerminalAPIEndpointLive           = "https://terminal-api-live.adyen.com"
	DisputesEndpointTest              = "https://ca-test.adyen.com/ca/services/DisputeService"
	DisputesEndpointLive              = "https://ca-live.adyen.com/ca/services/DisputeService"
	BalancePlatformEndpointTest       = "https://balanceplatform-api-test.adyen.com/bcl"
	BalancePlatformEndpointLive       = "https://balanceplatform-api-live.adyen.com/bcl"
	TransfersEndpointTest             = "https://balanceplatform-api-test.adyen.com/btl"
	TransfersEndpointLive             = "https://balanceplatform-api-live.adyen.com/btl"
	ManagementEndpointTest            = "https://management-test.adyen.com"
	ManagementEndpointLive            = "https://management-live.adyen.com"
	LegalEntityEntityTest             = "https://kyc-test.adyen.com/lem"
	LegalEntityEntityLive             = "https://kyc-live.adyen.com/lem"
	PosTerminalManagementEndpointTest = "https://postfmapi-test.adyen.com/postfmapi/terminal"
	PosTerminalManagementEndpointLive = "https://postfmapi-live.adyen.com/postfmapi/terminal"
	DataProtectionEndpointTest        = "https://ca-test.adyen.com/ca/services/DataProtectionService"
	DataProtectionEndpointLive        = "https://ca-live.adyen.com/ca/services/DataProtectionService"
	PosMobileEndpointTest             = "https://checkout-test.adyen.com/checkout/possdk"
	PosMobileEndpointLiveSuffix       = "-checkout-live.adyenpayments.com/checkout/possdk"
	CapitalEndpointTest               = "https://balanceplatform-api-test.adyen.com/capital"
	CapitalEndpointLive               = "https://balanceplatform-api-live.adyen.com/capital"
	ForeignExchangeEndpointTest       = "https://balanceplatform-api-test.adyen.com/fx/api"
	ForeignExchangeEndpointLive       = "https://balanceplatform-api-live.adyen.com/fx/api"
	SessionAuthenticationEndpointTest = "https://test.adyen.com/authe/api"
	SessionAuthenticationEndpointLive = "https://authe-live.adyen.com/authe/api"
	DocumentCollectorEndpointTest     = "https://document-collector-test.adyen.com"
	DocumentCollectorEndpointLive     = "https://document-collector-live.adyen.com"
	// SoftPOS Configuration has one LIVE endpoint per region; SetEnvironment uses Europe.
	SoftPOSConfigurationEndpointTest     = "https://softposconfig-test.adyen.com/softposconfig"
	SoftPOSConfigurationEndpointLive     = "https://softposconfig-live.adyen.com/softposconfig"
	SoftPOSConfigurationEndpointLiveAU   = "https://softposconfig-live-au.adyen.com/softposconfig"
	SoftPOSConfigurationEndpointLiveAPSE = "https://softposconfig-live-apse.adyen.com/softposconfig"
	SoftPOSConfigurationEndpointLiveNEA  = "https://softposconfig-live-nea.adyen.com/softposconfig"
	SoftPOSConfigurationEndpointLiveUS   = "https://softposconfig-live-us.adyen.com/softposconfig"
)

// also update LibVersion in src/common/configuration.go when a version is updated and a major lib version is released
const (
	PaymentAPIVersion               = "v68"
	PaymentsAppAPIVersion           = "v1"
	RecurringAPIVersion             = "v68"
	CheckoutAPIVersion              = "v72"
	BinLookupAPIVersion             = "v54"
	BalanceControlAPIVersion        = "v1"
	EndpointProtocol                = "https://"
	DisputesAPIVersion              = "v30"
	StoredValueAPIVersion           = "v46"
	BalancePlatformAPIVersion       = "v2"
	TransfersAPIVersion             = "v4"
	ManagementAPIVersion            = "v3"
	LegalEntityAPIVersion           = "v4"
	PosTerminalManagementAPIVersion = "v1"
	DataProtectionAPIVersion        = "v1"
	PosMobileAPIVersion             = "v68"
	CapitalAPIVersion               = "v1"
	ForeignExchangeAPIVersion       = "v1"
	SessionAuthenticationAPIVersion = "v1"
	DocumentCollectorAPIVersion     = "v1"
	SoftPOSConfigurationAPIVersion  = "v3"
)

// APIClient Manages access to Adyen API services.
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	client *common.Client
	// API Services
	checkout              *checkout.APIClient
	payments              *payments.APIClient
	paymentsApp           *paymentsapp.PaymentsAppApi
	payout                *payout.APIClient
	recurring             *recurring.GeneralApi
	binLookup             *binlookup.GeneralApi
	balancecontrol        *balancecontrol.GeneralApi
	posTerminalManagement *posterminalmanagement.GeneralApi
	disputes              *disputes.GeneralApi
	storedValue           *storedvalue.GeneralApi
	balancePlatform       *balanceplatform.APIClient
	transfers             *transfers.APIClient
	management            *management.APIClient
	legalEntity           *legalentity.APIClient
	dataProtection        *dataprotection.GeneralApi
	posMobile             *posmobile.GeneralApi
	capital               *capital.APIClient
	foreignExchange       *foreignexchange.APIClient
	sessionAuthentication *sessionauthentication.APIClient
	documentCollector     *documentcollector.APIClient
	softPOSConfiguration  *softposconfiguration.APIClient
}

// NewClient optionally a custom http.Client can be passed via the Config allow for advanced features such as caching.
func NewClient(cfg *common.Config) *APIClient {

	if cfg.HTTPClient == nil {
		// init http.Client with default settings
		cfg.HTTPClient = http.DefaultClient
		if cfg.ConnectionTimeoutMillis != 0 {
			// set timeout as defined in config
			cfg.HTTPClient.Timeout = cfg.ConnectionTimeoutMillis
		} else {
			// set default timeout to 60 seconds otherwise
			cfg.HTTPClient.Timeout = 60 * time.Second
		}
	}

	if cfg.DefaultHeader == nil {
		cfg.DefaultHeader = make(map[string]string)
	}

	c := &APIClient{}
	c.client = &common.Client{}
	c.client.Cfg = cfg

	c.SetEnvironment(cfg.Environment, cfg.LiveEndpointURLPrefix)

	return c
}

// API Services
func (c *APIClient) PaymentsApp() *paymentsapp.PaymentsAppApi {
	if c.paymentsApp == nil {
		c.paymentsApp = &paymentsapp.PaymentsAppApi{
			Client: c.client,
			BasePath: func() string {
				// SetEnvironment already appends PaymentsAppAPIVersion to this endpoint
				return c.client.Cfg.PaymentsAppEndpoint
			},
		}
	}
	return c.paymentsApp
}

func (c *APIClient) Checkout() *checkout.APIClient {
	if c.checkout == nil {
		c.checkout = checkout.NewAPIClient(c.client)
	}
	return c.checkout
}

func (c *APIClient) Payments() *payments.APIClient {
	if c.payments == nil {
		c.payments = payments.NewAPIClient(c.client)
		c.payments.PaymentsApi.BasePath = func() string {
			return fmt.Sprintf("%s/pal/servlet/Payment/%s", c.client.Cfg.Endpoint, PaymentAPIVersion)
		}
	}
	return c.payments
}

func (c *APIClient) Payout() *payout.APIClient {
	if c.payout == nil {
		c.payout = payout.NewAPIClient(c.client)
		c.payout.InitializationApi.BasePath = func() string {
			return fmt.Sprintf("%s/pal/servlet/Payout/%s", c.client.Cfg.Endpoint, PaymentAPIVersion)
		}
	}
	return c.payout
}

func (c *APIClient) Recurring() *recurring.GeneralApi {
	if c.recurring == nil {
		c.recurring = &recurring.GeneralApi{
			Client: c.client,
			BasePath: func() string {
				return fmt.Sprintf("%s/pal/servlet/Recurring/%s", c.client.Cfg.Endpoint, RecurringAPIVersion)
			},
		}
	}
	return c.recurring
}

func (c *APIClient) BinLookup() *binlookup.GeneralApi {
	if c.binLookup == nil {
		c.binLookup = &binlookup.GeneralApi{
			Client: c.client,
			BasePath: func() string {
				return fmt.Sprintf("%s/pal/servlet/BinLookup/%s", c.client.Cfg.Endpoint, BinLookupAPIVersion)
			},
		}
	}
	return c.binLookup
}

func (c *APIClient) BalanceControl() *balancecontrol.GeneralApi {
	if c.balancecontrol == nil {
		c.balancecontrol = &balancecontrol.GeneralApi{
			Client: c.client,
			BasePath: func() string {
				return fmt.Sprintf("%s/pal/servlet/BalanceControl/%s", c.client.Cfg.Endpoint, BalanceControlAPIVersion)
			},
		}
	}
	return c.balancecontrol
}

func (c *APIClient) StoredValue() *storedvalue.GeneralApi {
	if c.storedValue == nil {
		c.storedValue = &storedvalue.GeneralApi{
			Client: c.client,
			BasePath: func() string {
				return fmt.Sprintf("%s/pal/servlet/StoredValue/%s", c.client.Cfg.Endpoint, StoredValueAPIVersion)
			},
		}
	}
	return c.storedValue
}

func (c *APIClient) PosTerminalManagement() *posterminalmanagement.GeneralApi {
	if c.posTerminalManagement == nil {
		c.posTerminalManagement = &posterminalmanagement.GeneralApi{
			Client: c.client,
			BasePath: func() string {
				return fmt.Sprintf("%s/%s", c.client.Cfg.PosTerminalManagementEndpoint, PosTerminalManagementAPIVersion)
			},
		}
	}
	return c.posTerminalManagement
}

func (c *APIClient) Transfers() *transfers.APIClient {
	if c.transfers == nil {
		c.transfers = transfers.NewAPIClient(c.client)
		c.transfers.TransfersApi.BasePath = func() string {
			return fmt.Sprintf("%s/%s", c.client.Cfg.TransfersEndpoint, TransfersAPIVersion)
		}
	}
	return c.transfers
}

func (c *APIClient) BalancePlatform() *balanceplatform.APIClient {
	if c.balancePlatform == nil {
		c.balancePlatform = balanceplatform.NewAPIClient(c.client)
	}
	return c.balancePlatform
}

func (c *APIClient) Management() *management.APIClient {
	if c.management == nil {
		c.management = management.NewAPIClient(c.client)
	}
	return c.management
}

func (c *APIClient) LegalEntity() *legalentity.APIClient {
	if c.legalEntity == nil {
		c.legalEntity = legalentity.NewAPIClient(c.client)
	}
	return c.legalEntity
}

func (c *APIClient) Disputes() *disputes.GeneralApi {
	if c.disputes == nil {
		c.disputes = &disputes.GeneralApi{
			Client: c.client,
			BasePath: func() string {
				return fmt.Sprintf("%s/%s", c.client.Cfg.DisputesEndpoint, DisputesAPIVersion)
			},
		}
	}
	return c.disputes
}

func (c *APIClient) DataProtection() *dataprotection.GeneralApi {
	if c.dataProtection == nil {
		c.dataProtection = &dataprotection.GeneralApi{
			Client: c.client,
			BasePath: func() string {
				return fmt.Sprintf("%s/%s", c.client.Cfg.DataProtectionEndpoint, DataProtectionAPIVersion)
			},
		}
	}
	return c.dataProtection
}

func (c *APIClient) PosMobile() *posmobile.GeneralApi {
	if c.posMobile == nil {
		c.posMobile = &posmobile.GeneralApi{
			Client: c.client,
			BasePath: func() string {
				// SetEnvironment already appends PosMobileAPIVersion to this endpoint
				return c.client.Cfg.PosMobileEndpoint
			},
		}
	}
	return c.posMobile
}

func (c *APIClient) Capital() *capital.APIClient {
	if c.capital == nil {
		c.capital = capital.NewAPIClient(c.client)
	}
	return c.capital
}

func (c *APIClient) ForeignExchange() *foreignexchange.APIClient {
	if c.foreignExchange == nil {
		c.foreignExchange = foreignexchange.NewAPIClient(c.client)
	}
	return c.foreignExchange
}

func (c *APIClient) SessionAuthentication() *sessionauthentication.APIClient {
	if c.sessionAuthentication == nil {
		c.sessionAuthentication = sessionauthentication.NewAPIClient(c.client)
	}
	return c.sessionAuthentication
}

func (c *APIClient) DocumentCollector() *documentcollector.APIClient {
	if c.documentCollector == nil {
		c.documentCollector = documentcollector.NewAPIClient(c.client)
	}
	return c.documentCollector
}

func (c *APIClient) SoftPOSConfiguration() *softposconfiguration.APIClient {
	if c.softPOSConfiguration == nil {
		c.softPOSConfiguration = softposconfiguration.NewAPIClient(c.client)
	}
	return c.softPOSConfiguration
}

/*
SetEnvironment This defines the payment environment for live or test

  - @param environment           This defines the payment environment live or test
  - @param liveEndpointUrlPrefix Provide the unique live url prefix from the "API URLs and Response" menu in the Adyen Customer Area
*/
func (c *APIClient) SetEnvironment(env common.Environment, liveEndpointURLPrefix string) {
	if env == common.LiveEnv {
		c.client.Cfg.Environment = env
		c.client.Cfg.DisputesEndpoint = DisputesEndpointLive
		if liveEndpointURLPrefix != "" {
			c.client.Cfg.Endpoint = EndpointProtocol + liveEndpointURLPrefix + EndpointLiveSuffix
			c.client.Cfg.CheckoutEndpoint = EndpointProtocol + liveEndpointURLPrefix + CheckoutEndpointLiveSuffix
			c.client.Cfg.PosMobileEndpoint = EndpointProtocol + liveEndpointURLPrefix + PosMobileEndpointLiveSuffix + "/" + PosMobileAPIVersion
		} else {
			c.client.Cfg.Endpoint = EndpointLive
			c.client.Cfg.CheckoutEndpoint = ""
			// POS Mobile has no LIVE endpoint without the live URL prefix.
			c.client.Cfg.PosMobileEndpoint = ""
		}
		c.client.Cfg.TerminalApiCloudEndpoint = TerminalAPIEndpointLive
		c.client.Cfg.BalancePlatformEndpoint = BalancePlatformEndpointLive
		c.client.Cfg.TransfersEndpoint = TransfersEndpointLive
		c.client.Cfg.ManagementEndpoint = ManagementEndpointLive
		c.client.Cfg.PaymentsAppEndpoint = ManagementEndpointLive
		c.client.Cfg.LegalEntityEndpoint = LegalEntityEntityLive
		c.client.Cfg.PosTerminalManagementEndpoint = PosTerminalManagementEndpointLive
		c.client.Cfg.DataProtectionEndpoint = DataProtectionEndpointLive
		c.client.Cfg.CapitalEndpoint = CapitalEndpointLive
		c.client.Cfg.ForeignExchangeEndpoint = ForeignExchangeEndpointLive
		c.client.Cfg.SessionAuthenticationEndpoint = SessionAuthenticationEndpointLive
		c.client.Cfg.DocumentCollectorEndpoint = DocumentCollectorEndpointLive
		c.client.Cfg.SoftPOSConfigurationEndpoint = SoftPOSConfigurationEndpointLive
	} else {
		c.client.Cfg.Environment = env
		c.client.Cfg.Endpoint = EndpointTest
		c.client.Cfg.CheckoutEndpoint = CheckoutEndpointTest
		c.client.Cfg.TerminalApiCloudEndpoint = TerminalAPIEndpointTest
		c.client.Cfg.DisputesEndpoint = DisputesEndpointTest
		c.client.Cfg.BalancePlatformEndpoint = BalancePlatformEndpointTest
		c.client.Cfg.TransfersEndpoint = TransfersEndpointTest
		c.client.Cfg.ManagementEndpoint = ManagementEndpointTest
		c.client.Cfg.PaymentsAppEndpoint = ManagementEndpointTest
		c.client.Cfg.LegalEntityEndpoint = LegalEntityEntityTest
		c.client.Cfg.PosTerminalManagementEndpoint = PosTerminalManagementEndpointTest
		c.client.Cfg.DataProtectionEndpoint = DataProtectionEndpointTest
		c.client.Cfg.PosMobileEndpoint = PosMobileEndpointTest + "/" + PosMobileAPIVersion
		c.client.Cfg.CapitalEndpoint = CapitalEndpointTest
		c.client.Cfg.ForeignExchangeEndpoint = ForeignExchangeEndpointTest
		c.client.Cfg.SessionAuthenticationEndpoint = SessionAuthenticationEndpointTest
		c.client.Cfg.DocumentCollectorEndpoint = DocumentCollectorEndpointTest
		c.client.Cfg.SoftPOSConfigurationEndpoint = SoftPOSConfigurationEndpointTest
	}

	c.client.Cfg.CheckoutEndpoint += "/" + CheckoutAPIVersion
	c.client.Cfg.BalancePlatformEndpoint += "/" + BalancePlatformAPIVersion
	c.client.Cfg.ManagementEndpoint += "/" + ManagementAPIVersion
	c.client.Cfg.PaymentsAppEndpoint += "/" + PaymentsAppAPIVersion
	c.client.Cfg.LegalEntityEndpoint += "/" + LegalEntityAPIVersion
	c.client.Cfg.CapitalEndpoint += "/" + CapitalAPIVersion
	c.client.Cfg.ForeignExchangeEndpoint += "/" + ForeignExchangeAPIVersion
	c.client.Cfg.SessionAuthenticationEndpoint += "/" + SessionAuthenticationAPIVersion
	c.client.Cfg.DocumentCollectorEndpoint += "/" + DocumentCollectorAPIVersion
	c.client.Cfg.SoftPOSConfigurationEndpoint += "/" + SoftPOSConfigurationAPIVersion
}

// GetConfig Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *common.Config {
	return c.client.Cfg
}
