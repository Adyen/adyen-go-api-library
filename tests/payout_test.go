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

	"github.com/adyen/adyen-go-api-library/v6/src/adyen"
	"github.com/adyen/adyen-go-api-library/v6/src/checkout"
	"github.com/adyen/adyen-go-api-library/v6/src/common"
	"github.com/adyen/adyen-go-api-library/v6/src/payout"
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
	cvc := "737"
	card := &payout.Card{
		Cvc:         &cvc,
		ExpiryMonth: "03",
		ExpiryYear:  "2030",
		HolderName:  "John Smith",
		Number:      "4111111111111111",
	}
	shopperName := &payout.Name{
		FirstName: "John",
		LastName:  "Smith",
	}
	amount := payout.Amount{
		Currency: "EUR",
		Value:    1,
	}
	countryCode := "NL"
	iban := "NL13TEST0123456789"
	ownerName := "S. Hopper"
	contract := "PAYOUT"
	bank := payout.BankAccount{
		CountryCode: &countryCode,
		Iban:        &iban,
		OwnerName:   &ownerName,
	}
	entityType := "NaturalPerson"
	recurring := payout.Recurring{
		Contract: &contract,
	}
	shopperEmail := "test@adyen.com"
	nationality := "NL"

	createStoreDetail := func(ref string) (payout.StoreDetailResponse, *_nethttp.Response, error) {
		return clientStore.Payout.StoreDetail(&payout.StoreDetailRequest{
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

	createStoreDetailAndSubmitThirdParty := func(ref string) (payout.StoreDetailAndSubmitResponse, *_nethttp.Response, error) {
		return clientStore.Payout.StoreDetailAndSubmitThirdParty(&payout.StoreDetailAndSubmitRequest{
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
				cc := checkout.NewCardDetails()
				cc.SetEncryptedCardNumber("test_4111111111111111")
				cc.SetEncryptedExpiryMonth("test_03")
				cc.SetEncryptedExpiryYear("test_2030")
				cc.SetEncryptedSecurityCode("test_737")
				cc.SetHolderName("John Smith")
				paymentRes, _, _ := client.Checkout.Payments(&checkout.PaymentRequest{
					Reference: "123456781235",
					Amount: checkout.Amount{
						Value:    12500,
						Currency: "EUR",
					},
					MerchantAccount: MerchantAccount,
					PaymentMethod:   checkout.CardDetailsAsCheckoutPaymentMethod(cc),
				})

				res, httpRes, err := client.Payout.Payout(&payout.PayoutRequest{
					Amount:          amount,
					MerchantAccount: MerchantAccount,
					Reference:       paymentRes.GetPspReference(),
					Card:            card,
					ShopperName:     shopperName,
				})

				authorised := common.Authorised.String()
				require.Nil(t, err)
				require.NotNil(t, httpRes)
				assert.Equal(t, 200, httpRes.StatusCode)
				require.NotNil(t, res)
				assert.NotEmpty(t, res.PspReference)
				assert.Equal(t, &authorised, res.ResultCode)
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
				require.NotNil(t, res)
				assert.NotEmpty(t, res.PspReference)
				assert.NotEmpty(t, res.ResultCode)
			})
		})

		t.Run("SubmitThirdParty", func(t *testing.T) {
			t.Run("Create an API request that should pass", func(t *testing.T) {
				ref := time.Now().String()
				storeDetail, _, _ := createStoreDetail(ref)
				contract := "PAYOUT"
				res, httpRes, err := clientStore.Payout.SubmitThirdParty(&payout.SubmitRequest{
					Amount:          amount,
					MerchantAccount: MerchantAccount,
					Recurring: payout.Recurring{
						Contract: &contract,
					},
					EntityType:                       &entityType,
					Reference:                        storeDetail.PspReference,
					SelectedRecurringDetailReference: "LATEST",
					ShopperEmail:                     shopperEmail,
					ShopperReference:                 ref,
					ShopperName:                      shopperName,
					DateOfBirth:                      &dateOfBirth,
					Nationality:                      &nationality,
				})

				require.Nil(t, err)
				require.NotNil(t, httpRes)
				assert.Equal(t, 200, httpRes.StatusCode)
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
				res, httpRes, err := clientReview.Payout.ConfirmThirdParty(&payout.ModifyRequest{
					MerchantAccount:   MerchantAccount,
					OriginalReference: payoutDetail.PspReference,
				})

				require.Nil(t, err)
				require.NotNil(t, httpRes)
				assert.Equal(t, 200, httpRes.StatusCode)
				require.NotNil(t, res)
				assert.NotEmpty(t, res.PspReference)
				assert.Equal(t, "[payout-confirm-received]", res.Response)
			})
		})
		t.Run("DeclineThirdParty", func(t *testing.T) {
			t.Run("Create an API request that should pass", func(t *testing.T) {
				ref := time.Now().String()
				payoutDetail, _, _ := createStoreDetailAndSubmitThirdParty(ref)
				res, httpRes, err := clientReview.Payout.DeclineThirdParty(&payout.ModifyRequest{
					MerchantAccount:   MerchantAccount,
					OriginalReference: payoutDetail.PspReference,
				})

				require.Nil(t, err)
				require.NotNil(t, httpRes)
				assert.Equal(t, 200, httpRes.StatusCode)
				require.NotNil(t, res)
				assert.NotEmpty(t, res.PspReference)
				assert.Equal(t, "[payout-decline-received]", res.Response)
			})
		})
	})
}
