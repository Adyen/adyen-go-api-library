# AccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAgeIndicator** | **string** | Indicator for the length of time since this shopper account was created in the merchant&#39;s environment. Allowed values: * notApplicable * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days | [optional] 
**AccountChangeDate** |  Pointer to [**time.Time**](time.Time.md) | Date when the shopper&#39;s account was last changed. | [optional] 
**AccountChangeIndicator** | **string** | Indicator for the length of time since the shopper&#39;s account was last updated. Allowed values: * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days | [optional] 
**AccountCreationDate** |  Pointer to [**time.Time**](time.Time.md) | Date when the shopper&#39;s account was created. | [optional] 
**AccountType** | **string** | Indicates the type of account. For example, for a multi-account card product. Allowed values: * notApplicable * credit * debit | [optional] 
**AddCardAttemptsDay** | **int32** | Number of attempts the shopper tried to add a card to their account in the last day. | [optional] 
**DeliveryAddressUsageDate** |  Pointer to [**time.Time**](time.Time.md) | Date the selected delivery address was first used. | [optional] 
**DeliveryAddressUsageIndicator** | **string** | Indicator for the length of time since this delivery address was first used. Allowed values: * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days | [optional] 
**HomePhone** | **string** | Shopper&#39;s home phone number (including the country code). | [optional] 
**MobilePhone** | **string** | Shopper&#39;s mobile phone number (including the country code). | [optional] 
**PasswordChangeDate** |  Pointer to [**time.Time**](time.Time.md) | Date when the shopper last changed their password. | [optional] 
**PasswordChangeIndicator** | **string** | Indicator when the shopper has changed their password. Allowed values: * notApplicable * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days | [optional] 
**PastTransactionsDay** | **int32** | Number of all transactions (successful and abandoned) from this shopper in the past 24 hours. | [optional] 
**PastTransactionsYear** | **int32** | Number of all transactions (successful and abandoned) from this shopper in the past year. | [optional] 
**PaymentAccountAge** |  Pointer to [**time.Time**](time.Time.md) | Date this payment method was added to the shopper&#39;s account. | [optional] 
**PaymentAccountIndicator** | **string** | Indicator for the length of time since this payment method was added to this shopper&#39;s account. Allowed values: * notApplicable * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days | [optional] 
**PurchasesLast6Months** | **int32** | Number of successful purchases in the last six months. | [optional] 
**SuspiciousActivity** | **bool** | Whether suspicious activity was recorded on this account. | [optional] 
**WorkPhone** | **string** | Shopper&#39;s work phone number (including the country code). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


