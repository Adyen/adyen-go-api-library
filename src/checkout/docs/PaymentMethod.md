# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brands** | **[]string** | List of possible brands. For example: visa, mc. | [optional] 
**Configuration** |  Pointer to [**map[string]interface{}**](.md) | The configuration of the payment method. | [optional] 
**Details** |  Pointer to [**[]InputDetail**](InputDetail.md) | All input details to be provided to complete the payment with this payment method. | [optional] 
**Group** |  Pointer to [**PaymentMethodGroup**](PaymentMethodGroup.md) |  | [optional] 
**InputDetails** |  Pointer to [**[]InputDetail**](InputDetail.md) | All input details to be provided to complete the payment with this payment method. | [optional] 
**Name** | **string** | The displayable name of this payment method. | [optional] 
**PaymentMethodData** | **string** | Echo data required to send in next calls. | [optional] 
**SupportsRecurring** | **bool** | Indicates whether this payment method supports tokenization or not. | [optional] 
**Type** | **string** | The unique payment method code. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


