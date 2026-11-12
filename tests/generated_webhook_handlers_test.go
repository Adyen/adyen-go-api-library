package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/adyen/adyen-go-api-library/v21/src/acswebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/balancewebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/configurationwebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/disputewebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/managementwebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/negativebalancewarningwebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/relayedauthorizationwebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/tokenizationwebhook"
)

func TestDisputeWebhookHandler(t *testing.T) {
	n, err := disputewebhook.HandleDisputeNotificationRequest(`{
		"type": "balancePlatform.dispute.created",
		"data": {
			"id": "DIS0001",
			"balancePlatform": "YOUR_BALANCE_PLATFORM",
			"disputedAmount": {"currency": "EUR", "value": 1000},
			"status": "open"
		}
	}`)
	require.NoError(t, err)
	assert.Equal(t, "balancePlatform.dispute.created", n.Type)
	assert.Equal(t, "DIS0001", n.Data.GetId())
	assert.Equal(t, int64(1000), n.Data.GetDisputedAmount().Value)
}

func TestNegativeBalanceWarningWebhookHandler(t *testing.T) {
	n, err := negativebalancewarningwebhook.HandleNegativeBalanceCompensationWarningNotificationRequest(`{
		"type": "balancePlatform.negativeBalanceCompensationWarning.scheduled",
		"environment": "test",
		"timestamp": "2024-01-01T10:00:00+01:00",
		"data": {
			"id": "NBC0001",
			"liableBalanceAccountId": "BA0001",
			"amount": {"currency": "EUR", "value": -500}
		}
	}`)
	require.NoError(t, err)
	assert.Equal(t, "test", n.Environment)
	assert.Equal(t, "BA0001", n.Data.GetLiableBalanceAccountId())
	assert.Equal(t, int64(-500), n.Data.GetAmount().Value)
	require.NotNil(t, n.Timestamp)
}

func TestRelayedAuthorizationWebhookHandler(t *testing.T) {
	n, err := relayedauthorizationwebhook.HandleRelayedAuthorisationRequest(`{
		"id": "RA0001",
		"type": "balancePlatform.authorisation.relayed",
		"environment": "test",
		"amount": {"currency": "EUR", "value": 4200},
		"paymentInstrument": {"id": "PI0001"}
	}`)
	require.NoError(t, err)
	assert.Equal(t, "RA0001", n.GetId())
	assert.Equal(t, "balancePlatform.authorisation.relayed", n.GetType())
	assert.Equal(t, int64(4200), n.GetAmount().Value)
	require.NotNil(t, n.PaymentInstrument)
	assert.Equal(t, "PI0001", n.PaymentInstrument.GetId())
}

func TestTokenizationWebhookHandler(t *testing.T) {
	n, err := tokenizationwebhook.HandleTokenizationCreatedDetailsNotificationRequest(`{
		"createdAt": "2024-01-01T10:00:00+01:00",
		"environment": "test",
		"eventId": "QBQQ9DLNRHHKGK38",
		"type": "recurring.token.created",
		"version": "1",
		"data": {
			"merchantAccount": "YOUR_MERCHANT_ACCOUNT",
			"operation": "created",
			"shopperReference": "YOUR_SHOPPER_REFERENCE",
			"storedPaymentMethodId": "M5N7TQ4TG5PFWR50",
			"type": "visastandarddebit"
		}
	}`)
	require.NoError(t, err)
	assert.Equal(t, "recurring.token.created", n.Type)
	assert.Equal(t, "M5N7TQ4TG5PFWR50", n.Data.StoredPaymentMethodId)
	assert.False(t, n.CreatedAt.IsZero())
}

// Handlers generated for webhook roots that had no handwritten handler in v21.
func TestGeneratedWebhookHandlersRejectMalformedJSON(t *testing.T) {
	malformed := `{"type": "x", "data": `
	handlers := map[string]func(string) (interface{}, error){
		"acs.relayedAuthentication": func(s string) (interface{}, error) { return acswebhook.HandleRelayedAuthenticationRequest(s) },
		"balance.releasedBlockedBalance": func(s string) (interface{}, error) {
			return balancewebhook.HandleReleasedBlockedBalanceNotificationRequest(s)
		},
		"configuration.accountPayoutAutoApplication": func(s string) (interface{}, error) {
			return configurationwebhook.HandleAccountPayoutAutoApplicationNotificationRequest(s)
		},
		"configuration.mandate": func(s string) (interface{}, error) { return configurationwebhook.HandleMandateNotificationRequest(s) },
		"configuration.networkToken": func(s string) (interface{}, error) {
			return configurationwebhook.HandleNetworkTokenNotificationRequest(s)
		},
		"configuration.payoutScheduleBA": func(s string) (interface{}, error) {
			return configurationwebhook.HandlePayoutScheduleBANotificationRequest(s)
		},
		"configuration.payoutScheduleBP": func(s string) (interface{}, error) {
			return configurationwebhook.HandlePayoutScheduleBPNotificationRequest(s)
		},
		"configuration.payoutScheduleState": func(s string) (interface{}, error) {
			return configurationwebhook.HandlePayoutScheduleStateNotificationRequest(s)
		},
		"configuration.score": func(s string) (interface{}, error) { return configurationwebhook.HandleScoreNotificationRequest(s) },
		"configuration.sweepExecution": func(s string) (interface{}, error) {
			return configurationwebhook.HandleSweepExecutionNotificationRequest(s)
		},
		"configuration.topUpConfiguration": func(s string) (interface{}, error) {
			return configurationwebhook.HandleTopUpConfigurationEventRequest(s)
		},
		"configuration.topUpConfigurationUpdated": func(s string) (interface{}, error) {
			return configurationwebhook.HandleTopUpConfigurationUpdatedEventRequest(s)
		},
		"dispute": func(s string) (interface{}, error) { return disputewebhook.HandleDisputeNotificationRequest(s) },
		"management.terminalBoarding": func(s string) (interface{}, error) {
			return managementwebhook.HandleTerminalBoardingNotificationRequest(s)
		},
		"management.terminalSettings": func(s string) (interface{}, error) {
			return managementwebhook.HandleTerminalSettingsNotificationRequest(s)
		},
		"negativeBalanceWarning": func(s string) (interface{}, error) {
			return negativebalancewarningwebhook.HandleNegativeBalanceCompensationWarningNotificationRequest(s)
		},
		"relayedAuthorization": func(s string) (interface{}, error) {
			return relayedauthorizationwebhook.HandleRelayedAuthorisationRequest(s)
		},
		"tokenization.alreadyExisting": func(s string) (interface{}, error) {
			return tokenizationwebhook.HandleTokenizationAlreadyExistingDetailsNotificationRequest(s)
		},
		"tokenization.created": func(s string) (interface{}, error) {
			return tokenizationwebhook.HandleTokenizationCreatedDetailsNotificationRequest(s)
		},
		"tokenization.disabled": func(s string) (interface{}, error) {
			return tokenizationwebhook.HandleTokenizationDisabledDetailsNotificationRequest(s)
		},
		"tokenization.updated": func(s string) (interface{}, error) {
			return tokenizationwebhook.HandleTokenizationUpdatedDetailsNotificationRequest(s)
		},
	}
	for name, handle := range handlers {
		t.Run(name, func(t *testing.T) {
			res, err := handle(malformed)
			assert.Error(t, err)
			assert.Nil(t, res)
		})
	}
}
