package tests

import (
    "context"
    "github.com/adyen/adyen-go-api-library/v6/src/adyen"
    "github.com/adyen/adyen-go-api-library/v6/src/checkout"
    "github.com/adyen/adyen-go-api-library/v6/src/common"
    "os"
    "testing"

    "github.com/joho/godotenv"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func Test_Checkout_Next(t *testing.T) {
    godotenv.Load("./../.env")

    MerchantAccount := os.Getenv("ADYEN_MERCHANT")

    client := adyen.NewClient(&common.Config{
        ApiKey:      os.Getenv("ADYEN_API_KEY"),
        Environment: "TEST",
    })
    //client.GetConfig().Debug = true

    t.Run("PaymentMethods", func(t *testing.T) {
        t.Run("Create an API request that should fail", func(t *testing.T) {
            paymentMethodsRequest := checkout.PaymentMethodsRequest{}

            res, httpRes, err := client.Checkout.PaymentMethods(&paymentMethodsRequest)

            require.NotNil(t, res)
            require.NotNil(t, httpRes)
            require.NotNil(t, err)
            assert.Equal(t, 403, httpRes.StatusCode)
            assert.Equal(t, "403 : Invalid Merchant Account (security: 901)", err.Error())
        })

        t.Run("Create an API request that should pass", func(t *testing.T) {
            paymentMethodsRequest := *checkout.NewPaymentMethodsRequest(MerchantAccount)

            res, httpRes, err := client.Checkout.PaymentMethods(&paymentMethodsRequest)

            require.NotNil(t, res)
            require.NotNil(t, httpRes)
            require.Nil(t, err)
            assert.Equal(t, 200, httpRes.StatusCode)
        })
    })

    t.Run("Payments", func(t *testing.T) {
        t.Run("Credit card payment", func(t *testing.T) {
            card := checkout.NewCardDetails()
            card.SetEncryptedCardNumber("test_4111111111111111")
            card.SetEncryptedExpiryMonth("test_03")
            card.SetEncryptedExpiryYear("test_2030")
            card.SetEncryptedSecurityCode("test_737")
            paymentRequest := *checkout.NewPaymentRequest(
                *checkout.NewAmount("EUR", int64(1234)),
                MerchantAccount,
                checkout.CardDetailsAsCheckoutPaymentMethod(card),
                "Reference_example",
                "ReturnUrl_example",
            )                                      // PaymentRequest |  (optional)
            paymentRequest.SetCaptureDelayHours(0) // assert int zero value is sent

            res, httpRes, err := client.Checkout.Payments(&paymentRequest)

            require.NotNil(t, res)
            require.NotNil(t, httpRes)
            require.Nil(t, err)

            assert.Equal(t, 200, httpRes.StatusCode)
            assert.Equal(t, "Authorised", res.GetResultCode())
            require.False(t, res.HasAction())
            assert.NotEmpty(t, res.GetPspReference())
            assert.NotEmpty(t, res.GetResultCode())
        })

        t.Run("iDEAL payment", func(t *testing.T) {
            idempotencyKey := "b9c3947f-b282-4059-a645-56ddbbd2fef3"
            ctx := common.WithIdempotencyKey(context.Background(), idempotencyKey)
            ideal := checkout.NewIdealDetails("1121")
            paymentRequest := *checkout.NewPaymentRequest(
                *checkout.NewAmount("EUR", int64(1234)),
                MerchantAccount,
                checkout.IdealDetailsAsCheckoutPaymentMethod(ideal),
                "Reference_example",
                "ReturnUrl_example",
            ) // PaymentRequest |  (optional)

            res, httpRes, err := client.Checkout.Payments(&paymentRequest, ctx)

            require.NotNil(t, res)
            require.NotNil(t, httpRes)
            require.Nil(t, err)

            assert.Equal(t, 200, httpRes.StatusCode)
            assert.Equal(t, "RedirectShopper", res.GetResultCode())
            require.True(t, res.HasAction())
            assert.Equal(t, "ideal", res.GetAction().CheckoutRedirectAction.GetPaymentMethodType())
            assert.NotEmpty(t, res.GetResultCode())
        })
    })
}
