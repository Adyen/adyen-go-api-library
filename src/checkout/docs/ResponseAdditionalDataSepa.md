# ResponseAdditionalDataSepa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SepadirectdebitDateOfSignature** | **string** | The transaction signature date.  Format: yyyy-MM-dd | [optional] 
**SepadirectdebitMandateId** | **string** | Its value corresponds to the  pspReference value of the transaction. | [optional] 
**SepadirectdebitSequenceType** | **string** | This field can take one of the following values: * OneOff: (OOFF) Direct debit instruction to initiate exactly one direct debit transaction. * First: (FRST) Initial/first collection in a series of direct debit instructions. * Recurring: (RCUR) Direct debit instruction to carry out regular direct debit transactions initiated by the creditor. * Final: (FNAL) Last/final collection in a series of direct debit instructions.  Example: OOFF | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


