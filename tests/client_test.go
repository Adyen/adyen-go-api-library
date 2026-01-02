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
	t.Run("User-Agent without ApplicationName", func(t *testing.T) {
		expectedUserAgent := fmt.Sprintf("%s/%s", common.LibName, common.LibVersion)

		handler := func(w http.ResponseWriter, r *http.Request) {
			userAgent := r.Header.Get("User-Agent")
			require.Equal(t, expectedUserAgent, userAgent)
			w.WriteHeader(http.StatusOK)
		}

		mockServer := httptest.NewServer(http.HandlerFunc(handler))
		defer mockServer.Close()

		client := adyen.NewClient(&common.Config{
			Environment: common.TestEnv,
			HTTPClient:  mockServer.Client(),
		})
		client.Payments().PaymentsApi.BasePath = func() string { return mockServer.URL }

		req := client.Payments().PaymentsApi.AuthoriseInput().PaymentRequest(payments.PaymentRequest{})
		_, _, err := client.Payments().PaymentsApi.Authorise(context.Background(), req)
		require.NoError(t, err)
	})

	t.Run("User-Agent with ApplicationName", func(t *testing.T) {
		appName := "MyTestApp"
		expectedUserAgent := fmt.Sprintf("%s %s/%s", appName, common.LibName, common.LibVersion)

		handler := func(w http.ResponseWriter, r *http.Request) {
			userAgent := r.Header.Get("User-Agent")
			require.Equal(t, expectedUserAgent, userAgent)
			w.WriteHeader(http.StatusOK)
		}

		mockServer := httptest.NewServer(http.HandlerFunc(handler))
		defer mockServer.Close()

		client := adyen.NewClient(&common.Config{
			ApplicationName: appName,
			Environment:     common.TestEnv,
			HTTPClient:      mockServer.Client(),
		})
		client.Payments().PaymentsApi.BasePath = func() string { return mockServer.URL }

		req := client.Payments().PaymentsApi.AuthoriseInput().PaymentRequest(payments.PaymentRequest{})
		_, _, err := client.Payments().PaymentsApi.Authorise(context.Background(), req)
		require.NoError(t, err)
	})
}
