# RecurringDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** |  Pointer to [**map[string]interface{}**](.md) | This field contains additional data, which may be returned in a particular response.  The additionalData object consists of entries, each of which includes the key and value. | [optional] 
**Alias** | **string** | The alias of the credit card number.  Applies only to recurring contracts storing credit card details | [optional] 
**AliasType** | **string** | The alias type of the credit card number.  Applies only to recurring contracts storing credit card details. | [optional] 
**Bank** |  Pointer to [**BankAccount**](BankAccount.md) |  | [optional] 
**BillingAddress** |  Pointer to [**Address**](Address.md) |  | [optional] 
**Card** |  Pointer to [**Card**](Card.md) |  | [optional] 
**ContractTypes** | **[]string** | Types of recurring contracts. | [optional] 
**CreationDate** |  Pointer to [**time.Time**](time.Time.md) | The date when the recurring details were created. | [optional] 
**FirstPspReference** | **string** | The &#x60;pspReference&#x60; of the first recurring payment that created the recurring detail. | [optional] 
**Name** | **string** | An optional descriptive name for this recurring detail. | [optional] 
**PaymentMethodVariant** | **string** | The  type or sub-brand of a payment method used, e.g. Visa Debit, Visa Corporate, etc. For more information, refer to [PaymentMethodVariant](https://docs.adyen.com/development-resources/paymentmethodvariant). | [optional] 
**RecurringDetailReference** | **string** | The reference that uniquely identifies the recurring detail. | 
**ShopperName** |  Pointer to [**Name**](Name.md) |  | [optional] 
**SocialSecurityNumber** | **string** | A shopper&#39;s social security number (only in countries where it is legal to collect). | [optional] 
**Variant** | **string** | The payment method, such as “mc\&quot;, \&quot;visa\&quot;, \&quot;ideal\&quot;, \&quot;paypal\&quot;. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


