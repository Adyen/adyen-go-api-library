package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com///testify/require"

	"github.com/adyen/adyen-go-api-library/v7/src/balanceplatformreportnotification"
)

func Test_BalancePlatform_Report_Notifications_HandleRequest(t *testing.T) {

	t.Run("should return report created success", func(t *testing.T) {
		notificationJson := `{
			"data": {
			  "balancePlatform": "YOUR_BALANCE_PLATFORM",
			  "accountHolder": {
				"id": "AH32272223222B59MTF7458DP"
			  },
			  "balanceAccount": {
				"id": "BA3227C223222B5B9SCR82TMV"
			  },
			  "creationDate": "2021-07-02T02:01:08+02:00",
			  "fileName": "balanceplatform_payments_accounting_report_2021_07_01.csv",
			  "reportType": "balanceplatform_payments_accounting_report",
			  "downloadUrl": "https://balanceplatform-test.adyen.com/balanceplatform/.../.../.../balanceplatform_payments_accounting_report_2021_07_01.csv"
			},
			"environment": "test",
			"type": "balancePlatform.report.created"
		  }
		`
		balancePlatformNotification, _ := balanceplatformreportnotification.HandleReportNotificationRequest(notificationJson)
		balancePlatformValue := "YOUR_BALANCE_PLATFORM"
		assert.Equal(t, &balancePlatformValue, balancePlatformNotification.Data.BalancePlatform)
		assert.Equal(t, "balanceplatform_payments_accounting_report", balancePlatformNotification.Data.ReportType)
	})
}
