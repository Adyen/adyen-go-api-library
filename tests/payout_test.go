package tests

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/common"
	"github.com/adyen/adyen-go-api-library/v8/src/payout"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Payout(t *testing.T) {
	client := adyen.NewClient(&common.Config{
		ApiKey:      "MyAPIKey",
		Environment: "TEST",
		Debug:       "true" == os.Getenv("DEBUG"),
	})

	merchantAccount := "Merch"
	dateOfBirth := "1990-01-01"
	cvc := "737"
	card := payout.Card{
		Cvc:         &cvc,
		ExpiryMonth: common.PtrString("03"),
		ExpiryYear:  common.PtrString("2030"),
		HolderName:  common.PtrString("John Smith"),
		Number:      common.PtrString("4111111111111111"),
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

	mux := http.NewServeMux()
	// Success cases
	mux.HandleFunc("/payout", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		w.Header().Set("Content-Type", "application/json")

		// Error case
		body, _ := ioutil.ReadAll(r.Body)
		if bytes.Contains(body, []byte("failCase")) {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, `{
				"status" : 500,
				"errorCode" : "903",
				"message" : "Internal error",
				"errorType" : "internal"
			}`)
			return
		}
		io.WriteString(w, `{
		  "additionalData": {
			"cvcResult": "0 Unknown",
			"authCode": "010523",
			"avsResult": "0 Unknown",
			"authorisationMid": "50",
			"acquirerAccountCode": "AcquirerAccountCode"
		  },
		  "pspReference": "ABC213",
		  "resultCode": "Authorised",
		  "authCode": "010523"
		}`)
	})
	mux.HandleFunc("/storeDetail", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{
		   "pspReference" : "ABC123",
		   "recurringDetailReference" : "XYZ123",
		   "resultCode" : "Success"
		}`)
	})
	mux.HandleFunc("/submitThirdParty", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{
		   "pspReference" : "ABC123",
		   "resultCode" : "[payout-submit-received]"
		}`)
	})
	mux.HandleFunc("/storeDetailAndSubmitThirdParty", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{
		   "pspReference" : "ABC123",
		   "resultCode" : "[payout-submit-received]"
		}`)
	})
	mux.HandleFunc("/confirmThirdParty", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{
		   "pspReference" : "ABC123",
		   "response" : "[payout-confirm-received]"
		}`)
	})
	mux.HandleFunc("/declineThirdParty", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{
		   "pspReference" : "ABC123",
		   "response" : "[payout-decline-received]"
		}`)
	})

	mockServer := httptest.NewServer(mux)
	defer mockServer.Close()
	client.Payout().InitializationApi.BasePath = func() string { return mockServer.URL }

	createStoreDetailAndSubmitThirdParty := func(ref string) (payout.StoreDetailAndSubmitResponse, *http.Response, error) {
		req := client.Payout().InitializationApi.StoreDetailAndSubmitThirdPartyInput()
		req = req.StoreDetailAndSubmitRequest(payout.StoreDetailAndSubmitRequest{
			Amount:           amount,
			Bank:             &bank,
			DateOfBirth:      dateOfBirth,
			EntityType:       entityType,
			MerchantAccount:  merchantAccount,
			Nationality:      nationality,
			Recurring:        recurring,
			Reference:        ref,
			ShopperEmail:     shopperEmail,
			ShopperName:      shopperName,
			ShopperReference: ref,
		})
		return client.Payout().InitializationApi.StoreDetailAndSubmitThirdParty(context.Background(), req)
	}

	t.Run("Configuration", func(t *testing.T) {
		testClient := adyen.NewClient(&common.Config{
			Environment: common.TestEnv,
		})
		assert.Equal(t, "https://pal-test.adyen.com/pal/servlet/Payout/v68", testClient.Payout().InitializationApi.BasePath())
		liveClient := adyen.NewClient(&common.Config{
			Environment:           common.LiveEnv,
			LiveEndpointURLPrefix: "go-live",
		})
		assert.Equal(t, "https://go-live-pal-live.adyenpayments.com/pal/servlet/Payout/v68", liveClient.Payout().ReviewingApi.BasePath())
	})

	t.Run("Instant Payouts", func(t *testing.T) {
		t.Run("Payout", func(t *testing.T) {
			t.Run("Create an API request that should pass", func(t *testing.T) {
				body := payout.PayoutRequest{
					Amount:          amount,
					MerchantAccount: merchantAccount,
					Reference:       "FOOBAR",
					Card:            &card,
					ShopperName:     shopperName,
				}
				payoutReq := client.Payout().InstantPayoutsApi.PayoutInput().PayoutRequest(body)
				res, httpRes, err := client.Payout().InstantPayoutsApi.Payout(context.Background(), payoutReq)

				require.Nil(t, err)
				require.NotNil(t, httpRes)
				assert.Equal(t, 200, httpRes.StatusCode)
				require.Equal(t, "ABC213", res.GetPspReference())
				assert.Equal(t, "Authorised", res.GetResultCode())
			})

			t.Run("Create an API request that should fail", func(t *testing.T) {
				body := payout.PayoutRequest{
					Amount:          amount,
					MerchantAccount: merchantAccount,
					Reference:       "failCase",
					Card:            &card,
					ShopperName:     shopperName,
				}
				payoutReq := client.Payout().InstantPayoutsApi.PayoutInput().PayoutRequest(body)
				_, httpRes, err := client.Payout().InstantPayoutsApi.Payout(context.Background(), payoutReq)

				require.NotNil(t, err)
				assert.Equal(t, 500, httpRes.StatusCode)
				require.Equal(t, "internal", err.(common.APIError).Type)
			})
		})
	})

	t.Run("Initialization", func(t *testing.T) {
		t.Run("StoreDetail", func(t *testing.T) {
			t.Run("Create an API request that should pass", func(t *testing.T) {
				req := client.Payout().InitializationApi.StoreDetailInput()
				req = req.StoreDetailRequest(payout.StoreDetailRequest{
					Bank:             &bank,
					DateOfBirth:      dateOfBirth,
					EntityType:       entityType,
					MerchantAccount:  merchantAccount,
					Nationality:      nationality,
					Recurring:        recurring,
					ShopperEmail:     shopperEmail,
					ShopperName:      shopperName,
					ShopperReference: "MyShopper",
				})
				res, httpRes, err := client.Payout().InitializationApi.StoreDetail(context.Background(), req)

				require.Nil(t, err)
				assert.Equal(t, 200, httpRes.StatusCode)
				assert.Equal(t, "ABC123", res.GetPspReference())
				assert.Equal(t, "Success", res.GetResultCode())
			})
		})

		t.Run("SubmitThirdParty", func(t *testing.T) {
			t.Run("Create an API request that should pass", func(t *testing.T) {
				req := client.Payout().InitializationApi.SubmitThirdPartyInput()
				req = req.SubmitRequest(payout.SubmitRequest{
					Amount:          amount,
					MerchantAccount: merchantAccount,
					Recurring: payout.Recurring{
						Contract: &contract,
					},
					EntityType:                       &entityType,
					Reference:                        "bar",
					SelectedRecurringDetailReference: "LATEST",
					ShopperEmail:                     shopperEmail,
					ShopperReference:                 "foo",
					ShopperName:                      shopperName,
					DateOfBirth:                      &dateOfBirth,
					Nationality:                      &nationality,
				})
				res, httpRes, err := client.Payout().InitializationApi.SubmitThirdParty(context.Background(), req)

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
				req := client.Payout().ReviewingApi.ConfirmThirdPartyInput()
				req = req.ModifyRequest(payout.ModifyRequest{
					MerchantAccount:   "merchantAccount",
					OriginalReference: "TheOG",
				})
				res, httpRes, err := client.Payout().ReviewingApi.ConfirmThirdParty(context.Background(), req)

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
				req := client.Payout().ReviewingApi.DeclineThirdPartyInput()
				req = req.ModifyRequest(payout.ModifyRequest{
					MerchantAccount:   merchantAccount,
					OriginalReference: "TheGO",
				})

				res, httpRes, err := client.Payout().ReviewingApi.DeclineThirdParty(context.Background(), req)

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
