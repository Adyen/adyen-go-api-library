/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package tests

import (
	"fmt"
	_nethttp "net/http"
	"os"
	"testing"
	"time"

	"github.com/adyen/adyen-go-api-library/v2/src/adyen"
	"github.com/adyen/adyen-go-api-library/v2/src/checkout"
	"github.com/adyen/adyen-go-api-library/v2/src/common"
	"github.com/adyen/adyen-go-api-library/v2/src/payouts"
	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Payout(t *testing.T) {
	godotenv.Load("./../.env")

	var (
		APIKey          = os.Getenv("ADYEN_API_KEY")
		ReviewAPIKey    = os.Getenv("ADYEN_REVIEWPAYOUT_APIKEY")
		StoreAPIKey     = os.Getenv("ADYEN_STOREPAYOUT_APIKEY")
		MerchantAccount = os.Getenv("ADYEN_MERCHANT")
	)

	client := adyen.NewClient(&common.Config{
		ApiKey:      APIKey,
		Environment: "TEST",
	})
	clientStore := adyen.NewClient(&common.Config{
		ApiKey:      StoreAPIKey,
		Environment: "TEST",
	})
	clientReview := adyen.NewClient(&common.Config{
		ApiKey:      ReviewAPIKey,
		Environment: "TEST",
	})

	dateOfBirth := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	card := &payouts.Card{
		Cvc:         "737",
		ExpiryMonth: "03",
		ExpiryYear:  "2030",
		HolderName:  "John Smith",
		Number:      "4111111111111111",
	}
	shopperName := &payouts.Name{
		FirstName: "John",
		LastName:  "Smith",
	}
	amount := payouts.Amount{
		Currency: "EUR",
		Value:    1,
	}
	bank := payouts.BankAccount{
		CountryCode: "NL",
		Iban:        "NL13TEST0123456789",
		OwnerName:   "S. Hopper",
	}
	entityType := "NaturalPerson"
	recurring := payouts.Recurring{
		Contract: "PAYOUT",
	}
	shopperEmail := "test@adyen.com"
	nationality := "NL"

	createStoreDetail := func(ref string) (payouts.StoreDetailResponse, *_nethttp.Response, error) {
		return clientStore.Payouts.StoreDetail(&payouts.StoreDetailRequest{
			Bank:             &bank,
			DateOfBirth:      dateOfBirth,
			EntityType:       entityType,
			MerchantAccount:  MerchantAccount,
			Nationality:      nationality,
			Recurring:        recurring,
			ShopperEmail:     shopperEmail,
			ShopperName:      shopperName,
			ShopperReference: ref,
		})
	}

	createStoreDetailAndSubmitThirdParty := func(ref string) (payouts.StoreDetailAndSubmitResponse, *_nethttp.Response, error) {
		return clientStore.Payouts.StoreDetailAndSubmitThirdParty(&payouts.StoreDetailAndSubmitRequest{
			Amount:           amount,
			Bank:             &bank,
			DateOfBirth:      dateOfBirth,
			EntityType:       entityType,
			MerchantAccount:  MerchantAccount,
			Nationality:      nationality,
			Recurring:        recurring,
			Reference:        ref,
			ShopperEmail:     shopperEmail,
			ShopperName:      shopperName,
			ShopperReference: ref,
		})
	}

	t.Run("Instant Payouts", func(t *testing.T) {
		t.Run("Payout", func(t *testing.T) {
			t.Run("Create an API request that should pass", func(t *testing.T) {
				paymentRes, _, _ := client.Checkout.Payments(&checkout.PaymentRequest{
					Reference: "123456781235",
					Amount: checkout.Amount{
						Value:    12500,
						Currency: "EUR",
					},
					MerchantAccount: MerchantAccount,
					PaymentMethod: map[string]interface{}{
						"type":        "scheme",
						"number":      "4111111111111111",
						"expiryMonth": "03",
						"expiryYear":  "2030",
						"holderName":  "John Smith",
						"cvc":         "737",
					},
				})

				res, httpRes, err := client.Payouts.Payout(&payouts.PayoutRequest{
					Amount:          amount,
					MerchantAccount: MerchantAccount,
					Reference:       paymentRes.PspReference,
					Card:            card,
					ShopperName:     shopperName,
					DateOfBirth:     &dateOfBirth,
					Nationality:     nationality,
				})

				require.Nil(t, err)
				require.NotNil(t, httpRes)
				assert.Equal(t, 200, httpRes.StatusCode)
				assert.Equal(t, "200 OK", httpRes.Status)
				require.NotNil(t, res)
				assert.NotEmpty(t, res.PspReference)
				assert.Equal(t, common.Received, res.ResultCode)
			})
		})
	})

	t.Run("Initialization", func(t *testing.T) {
		t.Run("StoreDetail", func(t *testing.T) {
			t.Run("Create an API request that should pass", func(t *testing.T) {
				ref := time.Now().String()
				res, httpRes, err := createStoreDetail(ref)

				fmt.Println(err)
				require.Nil(t, err)
				require.NotNil(t, httpRes)
				assert.Equal(t, 200, httpRes.StatusCode)
				assert.Equal(t, "200 OK", httpRes.Status)
				require.NotNil(t, res)
				assert.NotEmpty(t, res.PspReference)
				assert.NotEmpty(t, res.ResultCode)
			})
		})

		t.Run("SubmitThirdParty", func(t *testing.T) {
			t.Run("Create an API request that should pass", func(t *testing.T) {
				ref := time.Now().String()
				storeDetail, _, _ := createStoreDetail(ref)
				res, httpRes, err := clientStore.Payouts.SubmitThirdParty(&payouts.SubmitRequest{
					Amount:          amount,
					MerchantAccount: MerchantAccount,
					Recurring: payouts.Recurring{
						Contract: "PAYOUT",
					},
					EntityType:                       entityType,
					Reference:                        storeDetail.PspReference,
					SelectedRecurringDetailReference: "LATEST",
					ShopperEmail:                     shopperEmail,
					ShopperReference:                 ref,
					ShopperName:                      shopperName,
					DateOfBirth:                      &dateOfBirth,
					Nationality:                      nationality,
				})

				require.Nil(t, err)
				require.NotNil(t, httpRes)
				assert.Equal(t, 200, httpRes.StatusCode)
				assert.Equal(t, "200 OK", httpRes.Status)
				require.NotNil(t, res)
				assert.NotEmpty(t, res.PspReference)
				assert.NotEmpty(t, res.ResultCode)
				assert.Equal(t, "[payout-submit-received]", res.ResultCode)
			})
		})

		t.Run("StoreDetailAndSubmitThirdParty", func(t *testing.T) {
			t.Run("Create an API request that should pass", func(t *testing.T) {
				ref := time.Now().String()
				res, httpRes, err := createStoreDetailAndSubmitThirdParty(ref)

				require.Nil(t, err)
				require.NotNil(t, httpRes)
				assert.Equal(t, 200, httpRes.StatusCode)
				assert.Equal(t, "200 OK", httpRes.Status)
				require.NotNil(t, res)
				assert.NotEmpty(t, res.PspReference)
				assert.NotEmpty(t, res.ResultCode)
				assert.Equal(t, "[payout-submit-received]", res.ResultCode)
			})
		})
	})

	t.Run("Reviewing", func(t *testing.T) {
		t.Run("ConfirmThirdParty", func(t *testing.T) {
			t.Run("Create an API request that should pass", func(t *testing.T) {
				ref := time.Now().String()
				payoutDetail, _, _ := createStoreDetailAndSubmitThirdParty(ref)
				res, httpRes, err := clientReview.Payouts.ConfirmThirdParty(&payouts.ModifyRequest{
					MerchantAccount:   MerchantAccount,
					OriginalReference: payoutDetail.PspReference,
				})

				require.Nil(t, err)
				require.NotNil(t, httpRes)
				assert.Equal(t, 200, httpRes.StatusCode)
				assert.Equal(t, "200 OK", httpRes.Status)
				require.NotNil(t, res)
				assert.NotEmpty(t, res.PspReference)
				assert.Equal(t, "[payout-confirm-received]", res.Response)
			})
		})
		t.Run("DeclineThirdParty", func(t *testing.T) {
			t.Run("Create an API request that should pass", func(t *testing.T) {
				ref := time.Now().String()
				payoutDetail, _, _ := createStoreDetailAndSubmitThirdParty(ref)
				res, httpRes, err := clientReview.Payouts.DeclineThirdParty(&payouts.ModifyRequest{
					MerchantAccount:   MerchantAccount,
					OriginalReference: payoutDetail.PspReference,
				})

				require.Nil(t, err)
				require.NotNil(t, httpRes)
				assert.Equal(t, 200, httpRes.StatusCode)
				assert.Equal(t, "200 OK", httpRes.Status)
				require.NotNil(t, res)
				assert.NotEmpty(t, res.PspReference)
				assert.Equal(t, "[payout-decline-received]", res.Response)
			})
		})
	})
}
