# StoreDetailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** |  Pointer to [**map[string]interface{}**](.md) | This field contains additional data, which may be required for a particular request. | [optional] 
**Bank** |  Pointer to [**BankAccount**](BankAccount.md) |  | [optional] 
**BillingAddress** |  Pointer to [**Address**](Address.md) |  | [optional] 
**Card** |  Pointer to [**Card**](Card.md) |  | [optional] 
**DateOfBirth** |  Pointer to [**time.Time**](time.Time.md) | The date of birth. Format: [ISO-8601](https://www.w3.org/TR/NOTE-datetime); example: YYYY-MM-DD For Paysafecard it must be the same as used when registering the Paysafecard account. &gt; This field is mandatory for natural persons. | 
**EntityType** | **string** | The type of the entity the payout is processed for. | 
**FraudOffset** | **int32** | An integer value that is added to the normal fraud score. The value can be either positive or negative. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**Nationality** | **string** | The shopper&#39;s nationality.  A valid value is an ISO 2-character country code (e.g. &#39;NL&#39;). | 
**Recurring** |  Pointer to [**Recurring**](Recurring.md) |  | 
**SelectedBrand** | **string** | The name of the brand to make a payout to.  For Paysafecard it must be set to &#x60;paysafecard&#x60;. | [optional] 
**ShopperEmail** | **string** | The shopper&#39;s email address. | 
**ShopperName** |  Pointer to [**Name**](Name.md) |  | [optional] 
**ShopperReference** | **string** | The shopper&#39;s reference for the payment transaction. | 
**SocialSecurityNumber** | **string** | The shopper&#39;s social security number. | [optional] 
**TelephoneNumber** | **string** | The shopper&#39;s phone number. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


