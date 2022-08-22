# InlineResponse20014

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoreItemsRemaining** | **bool** | Returns a value of &#x60;true&#x60; if subsequent items can be retrieved. | [optional] [default to null]
**TotalItemCount** | **int32** | The total number of records after applying all filter query parameters. The &#x60;total_item_count&#x60; will be calculated if and only if the corresponding query parameter &#x60;total_item_count&#x60; is set to &#x60;true&#x60;. If this query parameter is not set or set to &#x60;false&#x60;, a value of &#x60;null&#x60; will be returned. | [optional] [default to null]
**ContinuationToken** | **string** | Continuation token that can be provided in the &#x60;continuation_token&#x60; query param to get the next page of data. If you use the continuation token to page through data you are guaranteed to get all items exactly once regardless of how items are modified. If an item is added or deleted during the pagination then it may or may not be returned. The continuation token is generated if the limit is less than the remaining number of items, and the default sort is used (no sort is specified). | [optional] [default to null]
**Items** | [**[]AlertWatcher**](AlertWatcher.md) | A list of alert watcher objects. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

