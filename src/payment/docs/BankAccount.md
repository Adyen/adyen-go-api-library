# BankAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccountNumber** | **string** | The bank account number (without separators). | [optional] 
**BankCity** | **string** | The bank city. | [optional] 
**BankLocationId** | **string** | The location id of the bank. The field value is &#x60;nil&#x60; in most cases. | [optional] 
**BankName** | **string** | The name of the bank. | [optional] 
**Bic** | **string** | The [Business Identifier Code](https://en.wikipedia.org/wiki/ISO_9362) (BIC) is the SWIFT address assigned to a bank. The field value is &#x60;nil&#x60; in most cases. | [optional] 
**CountryCode** | **string** | Country code where the bank is located.  A valid value is an ISO two-character country code (e.g. &#39;NL&#39;). | [optional] 
**Iban** | **string** | The [International Bank Account Number](https://en.wikipedia.org/wiki/International_Bank_Account_Number) (IBAN). | [optional] 
**OwnerName** | **string** | The name of the bank account holder. If you submit a name with non-Latin characters, we automatically replace some of them with corresponding Latin characters to meet the FATF recommendations. For example: * χ12 is converted to ch12. * üA is converted to euA. * Peter Møller is converted to Peter Mller, because banks don&#39;t accept &#39;ø&#39;. After replacement, the ownerName must have at least three alphanumeric characters (A-Z, a-z, 0-9), and at least one of them must be a valid Latin character (A-Z, a-z). For example: * John17 - allowed. * J17 - allowed. * 171 - not allowed. * John-7 - allowed. &gt; If provided details don&#39;t match the required format, the response returns the error message: 203 &#39;Invalid bank account holder name&#39;. | [optional] 
**TaxId** | **string** | The bank account holder&#39;s tax ID. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


