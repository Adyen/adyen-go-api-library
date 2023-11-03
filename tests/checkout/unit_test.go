package checkout

import (
	"context"
	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/checkout"
	"github.com/adyen/adyen-go-api-library/v8/src/common"
	"github.com/adyen/adyen-go-api-library/v8/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckout(t *testing.T) {
	client := adyen.NewClient(&common.Config{
		ApiKey:      "YOUR_ADYEN_API_KEY",
		Environment: "TEST",
	})
	merchantAccount := "YOUR_MERCHANT_ACCOUNT"

	mux := http.NewServeMux()

	mockResponse := tests.MockResponse(t, mux)
	mockOk := func(method, endpoint, fixture string) {
		mockResponse(http.StatusOK, method, endpoint, fixture)
	}
	mockCreated := func(method, endpoint, fixture string) {
		mockResponse(http.StatusCreated, method, endpoint, fixture)
	}

	// Success cases
	mockOk("POST", "/payments", "checkout/authorised_card.json")
	mockOk("POST", "/paymentMethods", "checkout/payment_methods.json")
	mockOk("POST", "/paymentMethods/balance", "checkout/gift_card_balance.json")
	mockOk("POST", "/orders", "checkout/order_created.json")
	mockOk("GET PATCH", "/paymentLinks/PL61C53A8B97E6915A", "checkout/payment_link.json")
	mockCreated("POST", "/sessions", "checkout/session_created.json")

	mux.HandleFunc("/orders/cancel", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{
		  "pspReference": "8816178914079738",
		  "resultCode": "Received"
		}`)
	})

	mockServer := httptest.NewServer(mux)
	defer mockServer.Close()
	client.Checkout().PaymentsApi.BasePath = func() string { return mockServer.URL }
	service := client.Checkout()

	t.Run("Configuration", func(t *testing.T) {
		liveClient := adyen.NewClient(&common.Config{
			ApiKey:                "LIVE_API_KEY",
			Environment:           common.LiveEnv,
			LiveEndpointURLPrefix: "abc123",
			Debug:                 false,
		})
		require.Equal(t, "https://abc123-checkout-live.adyenpayments.com/checkout/"+adyen.CheckoutAPIVersion, liveClient.Checkout().PaymentsApi.BasePath())
	})

	t.Run("PaymentMethods", func(t *testing.T) {
		t.Run("Create an API request that should pass", func(t *testing.T) {
			paymentMethodsRequest := *checkout.NewPaymentMethodsRequest(merchantAccount)
			req := service.PaymentsApi.PaymentMethodsInput().PaymentMethodsRequest(paymentMethodsRequest)
			res, httpRes, err := service.PaymentsApi.PaymentMethods(context.Background(), req)

			require.NotNil(t, res)
			require.NotNil(t, httpRes)
			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			resBody, readErr := ioutil.ReadAll(httpRes.Body)
			require.NoError(t, readErr)
			assert.Contains(t, string(resBody), "paymentMethods")
		})
	})

	t.Run("Payments", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := service.PaymentsApi.Payments(
				context.Background(),
				// missing payment method
				service.PaymentsApi.PaymentsInput().PaymentRequest(checkout.PaymentRequest{
					MerchantAccount: merchantAccount,
				}),
			)

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "error calling MarshalJSON for type checkout.CheckoutPaymentMethod")
			require.Nil(t, httpRes)
			require.NotNil(t, res)
		})

		t.Run("Credit card payment", func(t *testing.T) {
			card := checkout.NewCardDetails()
			card.SetEncryptedCardNumber("test_4111111111111111")
			card.SetEncryptedExpiryMonth("test_03")
			card.SetEncryptedExpiryYear("test_2030")
			card.SetEncryptedSecurityCode("test_737")
			paymentRequest := *checkout.NewPaymentRequest(
				*checkout.NewAmount("EUR", int64(1234)),
				merchantAccount,
				checkout.CardDetailsAsCheckoutPaymentMethod(card),
				"Reference_example",
				"ReturnUrl_example",
			)
			paymentRequest.SetCaptureDelayHours(0) // assert int zero value is sent

			req := service.PaymentsApi.PaymentsInput()
			req = req.PaymentRequest(paymentRequest)
			res, httpRes, err := service.PaymentsApi.Payments(context.Background(), req)

			require.NotNil(t, res)
			require.NotNil(t, httpRes)
			require.Nil(t, err)

			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "Authorised", res.GetResultCode())
			require.False(t, res.HasAction())
			assert.NotEmpty(t, res.GetPspReference())
			assert.NotEmpty(t, res.GetResultCode())
		})
	})

	t.Run("Sessions", func(t *testing.T) {
		t.Run("Create an API request that should pass", func(t *testing.T) {
			body := checkout.CreateCheckoutSessionRequest{
				Reference: "123456781235",
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				CountryCode:     common.PtrString("NL"),
				MerchantAccount: merchantAccount,
				Channel:         common.PtrString("Web"),
				ReturnUrl:       "http://localhost:3000/redirect",
			}
			req := service.PaymentsApi.SessionsInput().CreateCheckoutSessionRequest(body)

			res, httpRes, err := service.PaymentsApi.Sessions(context.Background(), req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 201, httpRes.StatusCode)
			require.Equal(t, "CS451F2AB1ED897A94", res.GetId())
		})
	})

	t.Run("Orders", func(t *testing.T) {
		t.Run("Get balance", func(t *testing.T) {
			req := service.OrdersApi.GetBalanceOfGiftCardInput()
			req = req.BalanceCheckRequest(checkout.BalanceCheckRequest{
				MerchantAccount: merchantAccount,
				PaymentMethod: map[string]string{
					"type":       "giftcard",
					"brand":      "givex",
					"number":     "603628672882001915092",
					"holderName": "balance EUR 100",
					"cvc":        "5754",
				},
			})
			res, httpRes, err := service.OrdersApi.GetBalanceOfGiftCard(context.Background(), req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, int64(5000), res.Balance.Value)
		})

		t.Run("Create order", func(t *testing.T) {
			req := service.OrdersApi.OrdersInput()
			req = req.CreateOrderRequest(checkout.CreateOrderRequest{
				Amount: checkout.Amount{
					Currency: "EUR",
					Value:    2500,
				},
				MerchantAccount: merchantAccount,
				Reference:       "CREATE_ORDER_REF",
			})
			res, httpRes, err := service.OrdersApi.Orders(context.Background(), req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, int64(2500), res.RemainingAmount.Value)
		})

		t.Run("Cancel order", func(t *testing.T) {
			cancelReq := service.OrdersApi.CancelOrderInput()
			cancelReq = cancelReq.CancelOrderRequest(checkout.CancelOrderRequest{
				MerchantAccount: merchantAccount,
				Order: checkout.EncryptedOrderData{
					OrderData:    "823fh892f8f18f4...148f13f9f3f",
					PspReference: "8815517812932012",
				},
			})
			res, httpRes, err := service.OrdersApi.CancelOrder(context.Background(), cancelReq)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "Received", res.ResultCode)
		})
	})

	t.Run("PaymentLinks", func(t *testing.T) {
		paymentLinkId := "PL61C53A8B97E6915A"
		t.Run("Get payment link", func(t *testing.T) {
			req := service.PaymentLinksApi.GetPaymentLinkInput(paymentLinkId)
			res, httpRes, err := service.PaymentLinksApi.GetPaymentLink(context.Background(), req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "shopper-reference-ekvL83", res.Reference)
			assert.Equal(t, "active", res.Status)
			assert.NotNil(t, res.Url)
		})

		t.Run("Update payment link", func(t *testing.T) {
			req := service.PaymentLinksApi.UpdatePaymentLinkInput(paymentLinkId)
			req = req.UpdatePaymentLinkRequest(checkout.UpdatePaymentLinkRequest{
				Status: "expired",
			})
			res, httpRes, err := service.PaymentLinksApi.UpdatePaymentLink(context.Background(), req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "shopper-reference-ekvL83", res.Reference)
			assert.NotNil(t, res.Url)
		})
	})
}
