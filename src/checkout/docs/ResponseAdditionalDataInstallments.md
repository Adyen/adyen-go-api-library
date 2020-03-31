# ResponseAdditionalDataInstallments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallmentsValue** | **string** | The number of installments that the payment amount should be charged with.  Example: 5 &gt; Only relevant for card payments in countries that support installments. | [optional] 
**InstallmentPaymentDataInstallmentType** | **string** | Type of installment. The value of &#x60;installmentType&#x60; should be **IssuerFinanced**. | [optional] 
**InstallmentPaymentDataPaymentOptions** | **string** | Possible values: * PayInInstallmentsOnly * PayInFullOnly * PayInFullOrInstallments | [optional] 
**InstallmentPaymentDataOptionItemNrNumberOfInstallments** | **string** | Total number of installments possible for this payment. | [optional] 
**InstallmentPaymentDataOptionItemNrInterestRate** | **string** | Interest rate for the installment period. | [optional] 
**InstallmentPaymentDataOptionItemNrInstallmentFee** | **string** | Installment fee amount in minor units. | [optional] 
**InstallmentPaymentDataOptionItemNrAnnualPercentageRate** | **string** | Annual interest rate. | [optional] 
**InstallmentPaymentDataOptionItemNrFirstInstallmentAmount** | **string** | First Installment Amount in minor units. | [optional] 
**InstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount** | **string** | Subsequent Installment Amount in minor units. | [optional] 
**InstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments** | **string** | Minimum number of installments possible for this payment. | [optional] 
**InstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments** | **string** | Maximum number of installments possible for this payment. | [optional] 
**InstallmentPaymentDataOptionItemNrTotalAmountDue** | **string** | Total amount in minor units. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


