package tests

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/adyen/adyen-go-api-library/src/api"
	"github.com/adyen/adyen-go-api-library/src/common"
	"github.com/adyen/adyen-go-api-library/src/payments"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Payment(t *testing.T) {
	godotenv.Load("./../.env")

	var (
		MerchantAccount = os.Getenv("ADYEN_MERCHANT")
		APIKey          = os.Getenv("ADYEN_API_KEY")
	)

	card := payments.Card{
		Cvc:         "737",
		ExpiryMonth: "03",
		ExpiryYear:  "2030",
		HolderName:  "John Smith",
		Number:      "4111111111111111",
	}

	amount := payments.Amount{
		Currency: "EUR",
		Value:    1500,
	}

	// Used for 3D tests (Skipped for now)
	//browserInfo := &payments.BrowserInfo{
	//    AcceptHeader:      "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	//    ColorDepth:        24,
	//    JavaEnabled:       true,
	//    JavaScriptEnabled: true,
	//    Language:          "en",
	//    ScreenHeight:      1080,
	//    ScreenWidth:       1920,
	//    TimeZoneOffset:    1,
	//    UserAgent:         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36",
	//}

	client := api.NewClient(&common.Config{
		ApiKey:      APIKey,
		Environment: "TEST",
	})
	// client.GetConfig().Debug = true

	assertForSuccessResponse := func(res interface{}, httpRes *http.Response, err error) {
		require.Nil(t, err)
		require.NotNil(t, httpRes)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "200 OK", httpRes.Status)
		require.NotNil(t, res)
	}

	assertForApplicationInfo := func(req *payments.ModificationRequest) {
		// check if req has ApplicationInfo added to it
		require.NotNil(t, req.ApplicationInfo)
		require.NotNil(t, req.ApplicationInfo.AdyenLibrary)
		require.Equal(t, common.LibName, req.ApplicationInfo.AdyenLibrary.Name)
		require.Equal(t, common.LibVersion, req.ApplicationInfo.AdyenLibrary.Version)
	}

	authorisePost := func() (payments.PaymentResult, *http.Response, error) {
		return client.Payments.Authorise(&payments.PaymentRequest{
			Card:            &card,
			Amount:          amount,
			Reference:       time.Now().String(),
			MerchantAccount: MerchantAccount,
		})
	}

	t.Run("General", func(t *testing.T) {
		t.Run("Authorise3d", func(t *testing.T) {
			t.Skip("skipping since 3d requires manual user authentication")
			res, httpRes, err := client.Payments.Authorise3d(&payments.PaymentRequest3d{})
			assertForSuccessResponse(res, httpRes, err)
		})

		t.Run("Authorise3ds2", func(t *testing.T) {
			t.Skip("skipping since 3d requires manual user authentication")
			res, httpRes, err := client.Payments.Authorise3ds2(&payments.PaymentRequest3ds2{})
			assertForSuccessResponse(res, httpRes, err)
		})

		t.Run("Authorise", func(t *testing.T) {
			res, httpRes, err := authorisePost()

			assertForSuccessResponse(res, httpRes, err)
			assert.NotNil(t, res.PspReference)
			assert.Equal(t, res.ResultCode, "Authorised")
		})

		t.Run("GetAuthenticationResult", func(t *testing.T) {
			t.Skip("skipping since this returns auth result after a 3d auth")
			res, httpRes, err := client.Payments.GetAuthenticationResult(&payments.AuthenticationResultRequest{})
			assertForSuccessResponse(res, httpRes, err)
		})

		t.Run("Retrieve3ds2Result", func(t *testing.T) {
			t.Skip("skipping since this returns auth result after a 3d auth")
			res, httpRes, err := client.Payments.Retrieve3ds2Result(&payments.ThreeDS2ResultRequest{})
			assertForSuccessResponse(res, httpRes, err)
		})
	})

	t.Run("Modifications", func(t *testing.T) {
		t.Run("AdjustAuthorisation", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			req := &payments.ModificationRequest{
				OriginalReference: authRes.PspReference,
				ModificationAmount: &payments.Amount{
					Currency: "EUR",
					Value:    1234,
				},
				Reference:       time.Now().String(),
				MerchantAccount: MerchantAccount,
			}
			require.Nil(t, req.ApplicationInfo)
			res, httpRes, err := client.Payments.AdjustAuthorisation(req)
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[adjustAuthorisation-received]", res.Response)
			assertForApplicationInfo(req)
		})

		t.Run("CancelOrRefund", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			req := &payments.ModificationRequest{
				OriginalReference: authRes.PspReference,
				Reference:         time.Now().String(),
				MerchantAccount:   MerchantAccount,
			}
			require.Nil(t, req.ApplicationInfo)
			res, httpRes, err := client.Payments.CancelOrRefund(req)
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[cancelOrRefund-received]", res.Response)
			assertForApplicationInfo(req)
		})

		t.Run("Cancel", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			req := &payments.ModificationRequest{
				OriginalReference: authRes.PspReference,
				Reference:         time.Now().String(),
				MerchantAccount:   MerchantAccount,
			}
			require.Nil(t, req.ApplicationInfo)
			res, httpRes, err := client.Payments.Cancel(req)
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[cancel-received]", res.Response)
			assertForApplicationInfo(req)
		})

		t.Run("Capture", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			req := &payments.ModificationRequest{
				OriginalReference: authRes.PspReference,
				ModificationAmount: &payments.Amount{
					Currency: "EUR",
					Value:    1234,
				},
				Reference:       time.Now().String(),
				MerchantAccount: MerchantAccount,
			}
			require.Nil(t, req.ApplicationInfo)
			res, httpRes, err := client.Payments.Capture(req)
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[capture-received]", res.Response)
			assertForApplicationInfo(req)
		})

		t.Run("Refund", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			req := &payments.ModificationRequest{
				OriginalReference: authRes.PspReference,
				ModificationAmount: &payments.Amount{
					Currency: "EUR",
					Value:    1234,
				},
				Reference:       time.Now().String(),
				MerchantAccount: MerchantAccount,
			}
			require.Nil(t, req.ApplicationInfo)
			res, httpRes, err := client.Payments.Refund(req)
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[refund-received]", res.Response)
			assertForApplicationInfo(req)
		})

		t.Run("TechnicalCancel", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			req := &payments.ModificationRequest{
				OriginalReference: authRes.PspReference,
				Reference:         time.Now().String(),
				MerchantAccount:   MerchantAccount,
			}
			require.Nil(t, req.ApplicationInfo)
			res, httpRes, err := client.Payments.TechnicalCancel(req)
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[technical-cancel-received]", res.Response)
			assertForApplicationInfo(req)
		})

		t.Run("VoidPendingRefund", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			req := &payments.ModificationRequest{
				OriginalReference: authRes.PspReference,
				Reference:         time.Now().String(),
				MerchantAccount:   MerchantAccount,
			}
			require.Nil(t, req.ApplicationInfo)
			res, httpRes, err := client.Payments.VoidPendingRefund(req)
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[voidPendingRefund-received]", res.Response)
			assertForApplicationInfo(req)
		})
	})
}
