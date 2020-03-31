# AdditionalDataRetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetryChainAttemptNumber** | **int32** | The number of times the transaction (not order) has been retried between different payment service providers. For instance, the &#x60;chainAttemptNumber&#x60; set to 2 means that this transaction has been recently tried on another provider before being sent to Adyen. &gt; If you submit &#x60;retry.chainAttemptNumber&#x60;, &#x60;retry.orderAttemptNumber&#x60;, and &#x60;retry.skipRetry&#x60; values, we also recommend you provide the &#x60;merchantOrderReference&#x60; to facilitate linking payment attempts together. | [optional] 
**RetryOrderAttemptNumber** | **int32** | The index of the attempt to bill a particular order, which is identified by the &#x60;merchantOrderReference&#x60; field. For example, if a recurring transaction fails and is retried one day later, then the order number for these attempts would be 1 and 2, respectively. &gt; If you submit &#x60;retry.chainAttemptNumber&#x60;, &#x60;retry.orderAttemptNumber&#x60;, and &#x60;retry.skipRetry&#x60; values, we also recommend you provide the &#x60;merchantOrderReference&#x60; to facilitate linking payment attempts together. | [optional] 
**RetrySkipRetry** | **bool** | The Boolean value indicating whether Adyen should skip or retry this transaction, if possible. &gt; If you submit &#x60;retry.chainAttemptNumber&#x60;, &#x60;retry.orderAttemptNumber&#x60;, and &#x60;retry.skipRetry&#x60; values, we also recommend you provide the &#x60;merchantOrderReference&#x60; to facilitate linking payment attempts together. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


