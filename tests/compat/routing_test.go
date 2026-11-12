package compat

import (
	"reflect"
	"sort"
	"testing"

	"github.com/adyen/adyen-go-api-library/v21/src/adyen"
	"github.com/adyen/adyen-go-api-library/v21/src/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// basePaths returns the distinct BasePath values exposed by a service accessor:
// a single-tag service carries BasePath itself, a multi-tag APIClient holds one
// *XxxApi per tag, and every tag must resolve to the same base path.
func basePaths(t *testing.T, service interface{}) []string {
	t.Helper()
	rv := reflect.ValueOf(service).Elem()
	if bp := rv.FieldByName("BasePath"); bp.IsValid() {
		return []string{bp.Call(nil)[0].String()}
	}
	seen := map[string]bool{}
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		if !rv.Type().Field(i).IsExported() || f.Kind() != reflect.Ptr || f.IsNil() {
			continue
		}
		if bp := f.Elem().FieldByName("BasePath"); bp.IsValid() {
			seen[bp.Call(nil)[0].String()] = true
		}
	}
	out := make([]string, 0, len(seen))
	for p := range seen {
		out = append(out, p)
	}
	sort.Strings(out)
	return out
}

func TestEndpointRouting(t *testing.T) {
	type route struct {
		test, live string
		// knownDefect documents a v21 routing bug that the migration must fix;
		// the expected value is the spec-correct URL, not the current one.
		knownDefect string
	}
	routes := map[string]route{
		"BalanceControl":        {"https://pal-test.adyen.com/pal/servlet/BalanceControl/v1", "https://abc123-pal-live.adyenpayments.com/pal/servlet/BalanceControl/v1", ""},
		"BalancePlatform":       {"https://balanceplatform-api-test.adyen.com/bcl/v2", "https://balanceplatform-api-live.adyen.com/bcl/v2", ""},
		"BinLookup":             {"https://pal-test.adyen.com/pal/servlet/BinLookup/v54", "https://abc123-pal-live.adyenpayments.com/pal/servlet/BinLookup/v54", ""},
		"Checkout":              {"https://checkout-test.adyen.com/checkout/v71", "https://abc123-checkout-live.adyenpayments.com/checkout/v71", ""},
		"DataProtection":        {"https://ca-test.adyen.com/ca/services/DataProtectionService/v1", "https://ca-live.adyen.com/ca/services/DataProtectionService/v1", ""},
		"Disputes":              {"https://ca-test.adyen.com/ca/services/DisputeService/v30", "https://ca-live.adyen.com/ca/services/DisputeService/v30", ""},
		"LegalEntity":           {"https://kyc-test.adyen.com/lem/v4", "https://kyc-live.adyen.com/lem/v4", ""},
		"Management":            {"https://management-test.adyen.com/v3", "https://management-live.adyen.com/v3", ""},
		"Payments":              {"https://pal-test.adyen.com/pal/servlet/Payment/v68", "https://abc123-pal-live.adyenpayments.com/pal/servlet/Payment/v68", ""},
		"PaymentsApp":           {"https://management-test.adyen.com/v1", "https://management-live.adyen.com/v1", ""},
		"Payout":                {"https://pal-test.adyen.com/pal/servlet/Payout/v68", "https://abc123-pal-live.adyenpayments.com/pal/servlet/Payout/v68", ""},
		"PosTerminalManagement": {"https://postfmapi-test.adyen.com/postfmapi/terminal/v1", "https://postfmapi-live.adyen.com/postfmapi/terminal/v1", ""},
		"Recurring":             {"https://pal-test.adyen.com/pal/servlet/Recurring/v68", "https://abc123-pal-live.adyenpayments.com/pal/servlet/Recurring/v68", ""},
		"StoredValue":           {"https://pal-test.adyen.com/pal/servlet/StoredValue/v46", "https://abc123-pal-live.adyenpayments.com/pal/servlet/StoredValue/v46", ""},
		"Transfers":             {"https://balanceplatform-api-test.adyen.com/btl/v4", "https://balanceplatform-api-live.adyen.com/btl/v4", ""},
	}

	accessors := map[string]bool{}
	clientType := reflect.TypeOf(&adyen.APIClient{})
	for i := 0; i < clientType.NumMethod(); i++ {
		m := clientType.Method(i)
		if m.Type.NumIn() == 1 && m.Type.NumOut() == 1 && m.Name != "GetConfig" {
			accessors[m.Name] = true
		}
	}
	for name := range routes {
		assert.True(t, accessors[name], "service accessor %s() was removed", name)
	}
	for name := range accessors {
		_, ok := routes[name]
		assert.True(t, ok, "new service accessor %s() has no routing expectation", name)
	}

	for name, want := range routes {
		name, want := name, want
		t.Run(name, func(t *testing.T) {
			if want.knownDefect != "" {
				t.Skip("known v21 defect: " + want.knownDefect)
			}
			for env, expected := range map[common.Environment]string{common.TestEnv: want.test, common.LiveEnv: want.live} {
				client := adyen.NewClient(&common.Config{ApiKey: "key", Environment: env, LiveEndpointURLPrefix: "abc123"})
				service := reflect.ValueOf(client).MethodByName(name).Call(nil)[0].Interface()
				require.Equal(t, []string{expected}, basePaths(t, service), "%s in %s", name, env)
			}
		})
	}
}

func TestSetEnvironmentSwitchesEndpoints(t *testing.T) {
	client := adyen.NewClient(&common.Config{ApiKey: "key", Environment: common.TestEnv})
	assert.Equal(t, "https://checkout-test.adyen.com/checkout/v71", client.Checkout().PaymentsApi.BasePath())

	client.SetEnvironment(common.LiveEnv, "abc123")
	assert.Equal(t, "https://abc123-checkout-live.adyenpayments.com/checkout/v71", client.Checkout().PaymentsApi.BasePath())
	assert.Equal(t, "https://abc123-pal-live.adyenpayments.com/pal/servlet/Payment/v68", client.Payments().PaymentsApi.BasePath())
}
