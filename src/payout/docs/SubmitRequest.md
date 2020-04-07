# SubmitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** |  Pointer to [**map[string]interface{}**](.md) | This field contains additional data, which may be required for a particular request. | [optional] 
**Amount** |  Pointer to [**Amount**](Amount.md) |  | 
**DateOfBirth** |  Pointer to [**time.Time**](time.Time.md) | The date of birth. Format: ISO-8601; example: YYYY-MM-DD  For Paysafecard it must be the same as used when registering the Paysafecard account.  &gt; This field is mandatory for natural persons.  &gt; This field is required to update the existing &#x60;dateOfBirth&#x60; that is associated with this recurring contract. | [optional] 
**EntityType** | **string** | The type of the entity the payout is processed for.  Allowed values: * NaturalPerson * Company &gt; This field is required to update the existing &#x60;entityType&#x60; that is associated with this recurring contract. | [optional] 
**FraudOffset** | **int32** | An integer value that is added to the normal fraud score. The value can be either positive or negative. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier you want to process the transaction request with. | 
**Nationality** | **string** | The shopper&#39;s nationality.  A valid value is an ISO 2-character country code (e.g. &#39;NL&#39;).  &gt; This field is required to update the existing nationality that is associated with this recurring contract. | [optional] 
**Recurring** |  Pointer to [**Recurring**](Recurring.md) |  | 
**Reference** | **string** | The merchant reference for this payout. This reference will be used in all communication to the merchant about the status of the payout. Although it is a good idea to make sure it is unique, this is not a requirement. | 
**SelectedRecurringDetailReference** | **string** | This is the &#x60;recurringDetailReference&#x60; you want to use for this payout.  You can use the value LATEST to select the most recently used recurring detail. | 
**ShopperEmail** | **string** | The shopper&#39;s email address. | 
**ShopperName** |  Pointer to [**Name**](Name.md) |  | [optional] 
**ShopperReference** | **string** | The shopper&#39;s reference for the payout transaction. | 
**ShopperStatement** | **string** | The description of this payout. This description is shown on the bank statement of the shopper (if this is supported by the chosen payment method). | [optional] 
**SocialSecurityNumber** | **string** | The shopper&#39;s social security number. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


