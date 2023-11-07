//go:build integration
// +build integration

package tests

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"net/http"
	"os"
	"testing"

	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/common"
	"github.com/adyen/adyen-go-api-library/v8/src/platformsaccount"
	"github.com/adyen/adyen-go-api-library/v8/src/platformsfund"
	"github.com/adyen/adyen-go-api-library/v8/src/platformsnotificationconfiguration"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Platforms(t *testing.T) {
	godotenv.Load("./../.env")

	var (
		Username = os.Getenv("ADYEN_MARKETPLACE_USER")
		Password = os.Getenv("ADYEN_MARKETPLACE_PASSWORD")
	)

	client := adyen.NewClient(&common.Config{
		MarketPayEndpoint: adyen.MarketpayEndpointTest,
		Username:          Username,
		Password:          Password,
	})
	// client.GetConfig().Debug = true

	createAccountHolder := func(id string) (platformsaccount.CreateAccountHolderResponse, *http.Response, error) {
		return client.PlatformsAccount().CreateAccountHolder(&platformsaccount.CreateAccountHolderRequest{
			AccountHolderCode: id,
			AccountHolderDetails: platformsaccount.AccountHolderDetails{
				Email:           "go_library@test.com",
				FullPhoneNumber: "+312030291928",
				WebAddress:      "http://example.com",
				IndividualDetails: &platformsaccount.IndividualDetails{
					Name: &platformsaccount.ViasName{
						FirstName: "John",
						Gender:    "MALE",
						LastName:  "Doe",
					},
				},
				Address: &platformsaccount.ViasAddress{
					Country: "NL",
				},
			},
			LegalEntity: "Individual",
		})

	}

	createAccount := func(id string) (platformsaccount.CreateAccountResponse, *http.Response, error) {
		return client.PlatformsAccount().CreateAccount(&platformsaccount.CreateAccountRequest{
			AccountHolderCode: uuid.New().String(),
			Description:       "Create Account GoLang E2E",
			Metadata:          map[string]string{"meta": "data"},
			PayoutSchedule:    "WEEKLY",
		})

	}

	uploadDocument := func(id string) (platformsaccount.UpdateAccountHolderResponse, *http.Response, error) {
		myImage := image.NewRGBA(image.Rect(0, 0, 5000, 5000))
		var buff bytes.Buffer
		png.Encode(&buff, myImage)
		encodedString := base64.StdEncoding.EncodeToString(buff.Bytes())
		return client.PlatformsAccount().UploadDocument(&platformsaccount.UploadDocumentRequest{
			DocumentContent: encodedString,
			DocumentDetail: platformsaccount.DocumentDetail{
				AccountHolderCode: id,
				DocumentType:      "ID_CARD_FRONT",
				Description:       "Go Lib Test Document",
				Filename:          "id_card_front.png",
			},
		})

	}

	createConfiguration := func() (platformsnotificationconfiguration.GetNotificationConfigurationResponse, *http.Response, error) {
		return client.PlatformsNotificationConfiguration().CreateNotificationConfiguration(&platformsnotificationconfiguration.CreateNotificationConfigurationRequest{
			ConfigurationDetails: platformsnotificationconfiguration.NotificationConfigurationDetails{
				Active:    true,
				NotifyURL: "https://www.adyen.com/notification-handler",
				EventConfigs: []platformsnotificationconfiguration.NotificationEventConfiguration{
					{EventType: "ACCOUNT_HOLDER_VERIFICATION", IncludeMode: "INCLUDE"},
				},
				SslProtocol: "SSL",
				Description: "Go Lang Lib Test" + uuid.New().String(),
			},
		})

	}

	t.Run("Account", func(t *testing.T) {
		t.Run("Create account holder", func(t *testing.T) {
			res, httpRes, err := createAccountHolder(uuid.New().String())

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.NotNil(t, res.PspReference)
		})
		t.Run("Create account", func(t *testing.T) {
			res, httpRes, err := createAccount(uuid.New().String())

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.NotNil(t, res.PspReference)
		})
		t.Run("Get account holder", func(t *testing.T) {
			// to manually test, comment next line...
			t.Skip("Async 'GET' request. Skipping...")

			accountHolderCode := uuid.New().String()
			createAccountHolder(accountHolderCode)

			// ... and uncomment the line below
			// time.Sleep(20 * time.Second)

			res, httpRes, err := client.PlatformsAccount().GetAccountHolder(&platformsaccount.GetAccountHolderRequest{
				AccountHolderCode: accountHolderCode,
			})

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, res.AccountHolderDetails.Email, "go_library@test.com")
		})
		t.Run("Update account holder", func(t *testing.T) {
			accountHolderCode := uuid.New().String()
			createAccountHolder(accountHolderCode)
			res, httpRes, err := client.PlatformsAccount().UpdateAccountHolder(&platformsaccount.UpdateAccountHolderRequest{
				AccountHolderCode:    accountHolderCode,
				AccountHolderDetails: &platformsaccount.AccountHolderDetails{Address: &platformsaccount.ViasAddress{Country: "BE"}},
			})

			require.Nil(t, err)
			assert.Equal(t, 202, httpRes.StatusCode)
			assert.Equal(t, res.AccountHolderDetails.Address.Country, "BE")
		})
		t.Run("Check account holder", func(t *testing.T) {
			accountHolderCode := uuid.New().String()
			createAccountHolder(accountHolderCode)
			res, httpRes, err := client.PlatformsAccount().CheckAccountHolder(&platformsaccount.PerformVerificationRequest{
				AccountHolderCode: accountHolderCode,
				AccountStateType:  "Processing",
				Tier:              2,
			})

			require.Nil(t, err)
			assert.Equal(t, 202, httpRes.StatusCode)
			assert.Equal(t, res.ResultCode, "Success")
		})

		t.Run("Upload Document", func(t *testing.T) {
			accountHolderCode := uuid.New().String()
			createAccountHolder(accountHolderCode)
			res, httpRes, err := uploadDocument(accountHolderCode)

			require.Nil(t, err)
			assert.Contains(t, [2]int{200, 202}, httpRes.StatusCode)
			assert.NotNil(t, res.PspReference)
		})

		t.Run("Get uploaded document", func(t *testing.T) {
			// to manually test, comment next line...
			t.Skip("Async 'GET' request. Skipping...")

			accountHolderCode := uuid.New().String()
			createAccountHolder(accountHolderCode)
			uploadDocument(accountHolderCode)

			// ... and uncomment the line below
			// time.Sleep(5 * time.Second)

			res, httpRes, err := client.PlatformsAccount().GetUploadedDocuments(&platformsaccount.GetUploadedDocumentsRequest{
				AccountHolderCode: accountHolderCode,
			})

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.NotNil(t, res.PspReference)
		})
		t.Run("Close account", func(t *testing.T) {
			// to manually test, comment next line...
			t.Skip("Async request. Skipping...")

			r, _, _ := createAccount(uuid.New().String())

			// ... and uncomment the line below
			// time.Sleep(5 * time.Second)

			res, httpRes, err := client.PlatformsAccount().CloseAccount(&platformsaccount.CloseAccountRequest{
				AccountCode: r.AccountCode,
			})

			require.Nil(t, err)
			assert.Equal(t, 202, httpRes.StatusCode)
			assert.Equal(t, res.Status, "Closed")
		})
		t.Run("Suspend account holder", func(t *testing.T) {
			accountHolderCode := uuid.New().String()
			createAccountHolder(accountHolderCode)

			res, httpRes, err := client.PlatformsAccount().SuspendAccountHolder(&platformsaccount.SuspendAccountHolderRequest{
				AccountHolderCode: accountHolderCode,
			})

			require.Nil(t, err)
			assert.Equal(t, 202, httpRes.StatusCode)
			assert.NotNil(t, res.PspReference)
		})
		t.Run("Unsuspend account holder", func(t *testing.T) {
			t.Skip("Integration test")

			accountHolderCode := uuid.New().String()
			createAccountHolder(accountHolderCode)

			client.PlatformsAccount().SuspendAccountHolder(&platformsaccount.SuspendAccountHolderRequest{
				AccountHolderCode: accountHolderCode,
			})

			res, httpRes, err := client.PlatformsAccount().UnSuspendAccountHolder(&platformsaccount.UnSuspendAccountHolderRequest{
				AccountHolderCode: accountHolderCode,
			})

			require.Nil(t, err)
			assert.Equal(t, 202, httpRes.StatusCode)
			assert.NotNil(t, res.PspReference)
		})
		t.Run("Update account holder state", func(t *testing.T) {
			accountHolderCode := uuid.New().String()
			createAccountHolder(accountHolderCode)

			res, httpRes, err := client.PlatformsAccount().UpdateAccountHolderState(&platformsaccount.UpdateAccountHolderStateRequest{
				AccountHolderCode: accountHolderCode,
				Disable:           false,
				StateType:         "Payout",
			})

			require.Nil(t, err)
			assert.Equal(t, 202, httpRes.StatusCode)
			assert.NotNil(t, res.PspReference)
		})
		t.Run("Close account holder", func(t *testing.T) {
			accountHolderCode := uuid.New().String()
			createAccountHolder(accountHolderCode)

			res, httpRes, err := client.PlatformsAccount().CloseAccountHolder(&platformsaccount.CloseAccountHolderRequest{
				AccountHolderCode: accountHolderCode,
			})

			require.Nil(t, err)
			assert.Equal(t, 202, httpRes.StatusCode)
			assert.NotNil(t, res.PspReference)
		})
	})
	t.Run("Fund", func(t *testing.T) {
		t.Run("Retrieve account holder balance", func(t *testing.T) {
			t.Skip("Async request. Skipping...")

			accountHolderCode := uuid.New().String()
			createAccountHolder(accountHolderCode)

			res, httpRes, err := client.PlatformsFund().AccountHolderBalance(&platformsfund.AccountHolderBalanceRequest{
				AccountHolderCode: accountHolderCode,
			})

			require.Nil(t, err)
			assert.Equal(t, 202, httpRes.StatusCode)
			assert.Equal(t, res.BalancePerAccount, "")
		})
		t.Run("Retrieve list of transactions", func(t *testing.T) {
			t.Skip("Async request. Skipping...")

			accountHolderCode := uuid.New().String()
			createAccountHolder(accountHolderCode)

			res, httpRes, err := client.PlatformsFund().AccountHolderTransactionList(&platformsfund.AccountHolderTransactionListRequest{
				AccountHolderCode: accountHolderCode,
			})

			require.Nil(t, err)
			assert.Equal(t, 202, httpRes.StatusCode)
			assert.Equal(t, res.AccountTransactionLists, "")
		})
		t.Run("Transfer between accounts", func(t *testing.T) {
			// t.Skip("Async request. Skipping...")

			accountHolderCode := uuid.New().String()
			createAccountHolder(accountHolderCode)

			res, httpRes, err := client.PlatformsFund().TransferFunds(&platformsfund.TransferFundsRequest{
				SourceAccountCode:      "8515883280985939",
				DestinationAccountCode: "8815883278206345",
				Amount:                 platformsfund.Amount{Currency: "EUR", Value: 1},
				TransferCode:           "SUBSCRIPTION",
			})

			require.Nil(t, err)
			assert.Equal(t, 202, httpRes.StatusCode)
			assert.NotNil(t, res.PspReference)
		})
	})
	t.Run("Notification Configuration", func(t *testing.T) {
		t.Run("Create configuration", func(t *testing.T) {
			res, httpRes, err := createConfiguration()
			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.True(t, res.ConfigurationDetails.Active)
		})
		t.Run("Retrieve configurations", func(t *testing.T) {
			t.Skip("Fix me")
			res, httpRes, err := client.PlatformsNotificationConfiguration().GetNotificationConfigurationList(nil)

			require.Nil(t, err)
			assert.Equal(t, 202, httpRes.StatusCode)
			assert.NotNil(t, res.PspReference)
		})
		t.Run("Retrieve configuration", func(t *testing.T) {
			t.Skip("Async. Skipping...")
			r, _, _ := createConfiguration()
			res, httpRes, err := client.PlatformsNotificationConfiguration().GetNotificationConfiguration(&platformsnotificationconfiguration.GetNotificationConfigurationRequest{
				NotificationId: r.ConfigurationDetails.NotificationId,
			})

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.True(t, res.ConfigurationDetails.Active)
		})
		t.Run("Update configuration", func(t *testing.T) {
			t.Skip("Async request. Skipping...")
			r, _, _ := createConfiguration()
			res, httpRes, err := client.PlatformsNotificationConfiguration().UpdateNotificationConfiguration(&platformsnotificationconfiguration.UpdateNotificationConfigurationRequest{
				ConfigurationDetails: platformsnotificationconfiguration.NotificationConfigurationDetails{
					Active:    true,
					NotifyURL: "https://www.adyen.com/notification-handler",
					EventConfigs: []platformsnotificationconfiguration.NotificationEventConfiguration{
						{EventType: "ACCOUNT_HOLDER_VERIFICATION", IncludeMode: "EXCLUDE"},
					},
					SslProtocol:    "SSL",
					Description:    "Go Lang Lib Test" + uuid.New().String(),
					NotificationId: r.ConfigurationDetails.NotificationId,
				},
			})

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, res.ResultCode, "")
		})
		t.Run("Delete configuration", func(t *testing.T) {
			t.Skip("Async request. Skipping...")
			r, _, _ := createConfiguration()
			res, httpRes, err := client.PlatformsNotificationConfiguration().DeleteNotificationConfigurations(&platformsnotificationconfiguration.DeleteNotificationConfigurationRequest{
				NotificationIds: []int64{r.ConfigurationDetails.NotificationId},
			})

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, res.ResultCode, "")
		})
	})
}
