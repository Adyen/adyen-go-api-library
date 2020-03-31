# Recurring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | **string** | The type of recurring contract to be used. Possible values: * &#x60;ONECLICK&#x60; – Payment details can be used to initiate a one-click payment, where the shopper enters the [card security code (CVC/CVV)](https://docs.adyen.com/payments-fundamentals/payment-glossary#card-security-code-cvc-cvv-cid). * &#x60;RECURRING&#x60; – Payment details can be used without the card security code to initiate [card-not-present transactions](https://docs.adyen.com/payments-fundamentals/payment-glossary#card-not-present-cnp). * &#x60;ONECLICK,RECURRING&#x60; – Payment details can be used regardless of whether the shopper is on your site or not. * &#x60;PAYOUT&#x60; – Payment details can be used to [make a payout](https://docs.adyen.com/checkout/online-payouts). | [optional] 
**RecurringDetailName** | **string** | A descriptive name for this detail. | [optional] 
**RecurringExpiry** | [**time.Time**](time.Time.md) | Date after which no further authorisations shall be performed. Only for 3D Secure 2. | [optional] 
**RecurringFrequency** | **string** | Minimum number of days between authorisations. Only for 3D Secure 2. | [optional] 
**TokenService** | **string** | The name of the token service. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


