/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type InlineResponse20055 struct {
	// Returns a value of `true` if subsequent items can be retrieved.
	MoreItemsRemaining bool `json:"more_items_remaining,omitempty"`
	// The total number of records after applying all filter query parameters. The `total_item_count` will be calculated if and only if the corresponding query parameter `total_item_count` is set to `true`. If this query parameter is not set or set to `false`, a value of `null` will be returned.
	TotalItemCount int32 `json:"total_item_count,omitempty"`
	// Continuation token that can be provided in the `continuation_token` query param to get the next page of data. If you use the continuation token to page through data you are guaranteed to get all items exactly once regardless of how items are modified. If an item is added or deleted during the pagination then it may or may not be returned. The continuation token is generated if the limit is less than the remaining number of items, and the default sort is used (no sort is specified).
	ContinuationToken string `json:"continuation_token,omitempty"`
	Items []DirectoryService `json:"items,omitempty"`
}
