package tests

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adyen/adyen-go-api-library/v21/src/adyen"
	"github.com/adyen/adyen-go-api-library/v21/src/common"
	"github.com/adyen/adyen-go-api-library/v21/src/payments"
	"github.com/stretchr/testify/require"
)

func Test_UserAgent(t *testing.T) {
	cases := []struct {
		name              string
		appName           string
		expectedUserAgent string
	}{
		{
			name:              "User-Agent without ApplicationName",
			appName:           "",
			expectedUserAgent: fmt.Sprintf("%s/%s", common.LibName, common.LibVersion),
		},
		{
			name:              "User-Agent with ApplicationName",
			appName:           "MyTestApp",
			expectedUserAgent: fmt.Sprintf("%s %s/%s", "MyTestApp", common.LibName, common.LibVersion),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			handler := func(w http.ResponseWriter, r *http.Request) {
				// verify User-Agent header
				userAgent := r.Header.Get("User-Agent")
				require.Equal(t, tc.expectedUserAgent, userAgent)
				// verify adyen-library-name header
				libName := r.Header.Get("adyen-library-name")
				require.Equal(t, common.LibName, libName)
				// verify adyen-library-version header
				libVersion := r.Header.Get("adyen-library-version")
				require.Equal(t, common.LibVersion, libVersion)
				w.WriteHeader(http.StatusOK)
			}
			mockServer := httptest.NewServer(http.HandlerFunc(handler))
			defer mockServer.Close()

			config := &common.Config{
				Environment: common.TestEnv,
				HTTPClient:  mockServer.Client(),
			}

			if tc.appName != "" {
				config.ApplicationName = tc.appName
			}

			client := adyen.NewClient(config)
			client.Payments().PaymentsApi.BasePath = func() string { return mockServer.URL }

			req := client.Payments().PaymentsApi.AuthoriseInput().PaymentRequest(payments.PaymentRequest{})
			_, _, err := client.Payments().PaymentsApi.Authorise(context.Background(), req)
			require.NoError(t, err)
		})
	}
}
