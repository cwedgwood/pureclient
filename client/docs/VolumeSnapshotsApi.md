# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api28VolumeSnapshotsDelete**](VolumeSnapshotsApi.md#Api28VolumeSnapshotsDelete) | **Delete** /api/2.8/volume-snapshots | Eradicate a volume snapshot
[**Api28VolumeSnapshotsGet**](VolumeSnapshotsApi.md#Api28VolumeSnapshotsGet) | **Get** /api/2.8/volume-snapshots | List volume snapshots
[**Api28VolumeSnapshotsPatch**](VolumeSnapshotsApi.md#Api28VolumeSnapshotsPatch) | **Patch** /api/2.8/volume-snapshots | Manage a volume snapshot
[**Api28VolumeSnapshotsPost**](VolumeSnapshotsApi.md#Api28VolumeSnapshotsPost) | **Post** /api/2.8/volume-snapshots | Generate a volume snapshot
[**Api28VolumeSnapshotsTagsBatchPut**](VolumeSnapshotsApi.md#Api28VolumeSnapshotsTagsBatchPut) | **Put** /api/2.8/volume-snapshots/tags/batch | Update tags
[**Api28VolumeSnapshotsTagsDelete**](VolumeSnapshotsApi.md#Api28VolumeSnapshotsTagsDelete) | **Delete** /api/2.8/volume-snapshots/tags | Delete tags
[**Api28VolumeSnapshotsTagsGet**](VolumeSnapshotsApi.md#Api28VolumeSnapshotsTagsGet) | **Get** /api/2.8/volume-snapshots/tags | List tags
[**Api28VolumeSnapshotsTransferGet**](VolumeSnapshotsApi.md#Api28VolumeSnapshotsTransferGet) | **Get** /api/2.8/volume-snapshots/transfer | List volume snapshots with transfer statistics

# **Api28VolumeSnapshotsDelete**
> Api28VolumeSnapshotsDelete(ctx, optional)
Eradicate a volume snapshot

Eradicate a volume snapshot that has been destroyed and is pending eradication. Eradicated volumes snapshots cannot be recovered. Volume snapshots are destroyed through the `PATCH` method. The `ids` or `names` parameter is required, but cannot be set together.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VolumeSnapshotsApiApi28VolumeSnapshotsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeSnapshotsApiApi28VolumeSnapshotsDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **ids** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique resource IDs specified. Enter multiple resource IDs in comma-separated format. The &#x60;ids&#x60; and &#x60;names&#x60; parameters cannot be provided together. | 
 **names** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 
 **replicationSnapshot** | **optional.Bool**| If set to &#x60;true&#x60;, allow destruction/eradication of snapshots in use by replication. If set to &#x60;false&#x60;, allow destruction/eradication of snapshots not in use by replication. If not specified, defaults to &#x60;false&#x60;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28VolumeSnapshotsGet**
> InlineResponse200165 Api28VolumeSnapshotsGet(ctx, optional)
List volume snapshots

Return a list of volume snapshots, including those pending eradication.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VolumeSnapshotsApiApi28VolumeSnapshotsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeSnapshotsApiApi28VolumeSnapshotsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **continuationToken** | **optional.String**| A token used to retrieve the next page of data with some consistency guaranteed. The token is a Base64 encoded value. Set &#x60;continuation_token&#x60; to the system-generated token taken from the &#x60;x-next-token&#x60; header field of the response. A query has reached its last page when the response does not include a token. Pagination requires the &#x60;limit&#x60; and &#x60;continuation_token&#x60; query parameters. | 
 **destroyed** | **optional.Bool**| If set to &#x60;true&#x60;, lists only destroyed objects that are in the eradication pending state. If set to &#x60;false&#x60;, lists only objects that are not destroyed. For destroyed objects, the time remaining is displayed in milliseconds. | 
 **filter** | **optional.String**| Narrows down the results to only the response objects that satisfy the filter criteria. | 
 **ids** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique resource IDs specified. Enter multiple resource IDs in comma-separated format. The &#x60;ids&#x60; and &#x60;names&#x60; parameters cannot be provided together. | 
 **limit** | **optional.Int32**| Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size. | 
 **names** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 
 **offset** | **optional.Int32**| The starting position based on the results of the query in relation to the full set of response objects returned. | 
 **sort** | [**optional.Interface of []string**](string.md)| Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values. | 
 **sourceIds** | [**optional.Interface of []string**](string.md)| Performs the operation on the source ID specified. Enter multiple source IDs in comma-separated format. | 
 **sourceNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the source name specified. Enter multiple source names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 
 **totalItemCount** | **optional.Bool**| If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;. | 
 **totalOnly** | **optional.Bool**| If set to &#x60;true&#x60;, returns the aggregate value of all items after filtering. Where it makes more sense, the average value is displayed instead. The values are displayed for each name where meaningful. If &#x60;total_only&#x3D;true&#x60;, the &#x60;items&#x60; list will be empty. | 

### Return type

[**InlineResponse200165**](inline_response_200_165.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28VolumeSnapshotsPatch**
> InlineResponse200166 Api28VolumeSnapshotsPatch(ctx, body, optional)
Manage a volume snapshot

Rename, destroy, or recover a volume snapshot. To rename the suffix of a volume snapshot, set `name` to the new suffix name. To recover a volume snapshot that has been destroyed and is pending eradication, set `destroyed=true`. The `ids` or `names` parameter is required, but cannot be set together.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Object**](.md)|  | 
 **optional** | ***VolumeSnapshotsApiApi28VolumeSnapshotsPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeSnapshotsApiApi28VolumeSnapshotsPatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.**| Supplied by client during request or generated by server. | 
 **ids** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique resource IDs specified. Enter multiple resource IDs in comma-separated format. The &#x60;ids&#x60; and &#x60;names&#x60; parameters cannot be provided together. | 
 **names** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 
 **replicationSnapshot** | **optional.**| If set to &#x60;true&#x60;, allow destruction/eradication of snapshots in use by replication. If set to &#x60;false&#x60;, allow destruction/eradication of snapshots not in use by replication. If not specified, defaults to &#x60;false&#x60;. | 

### Return type

[**InlineResponse200166**](inline_response_200_166.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28VolumeSnapshotsPost**
> InlineResponse200166 Api28VolumeSnapshotsPost(ctx, body, optional)
Generate a volume snapshot

Create a point-in-time snapshot of the contents of a volume. The `source_ids` or `source_names` parameter is required, but cannot be set together.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Model28VolumesnapshotsBody**](Model28VolumesnapshotsBody.md)|  | 
 **optional** | ***VolumeSnapshotsApiApi28VolumeSnapshotsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeSnapshotsApiApi28VolumeSnapshotsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.**| Supplied by client during request or generated by server. | 
 **on** | **optional.**| Performs the operation on the target name specified. For example, &#x60;targetName01&#x60;. | 
 **sourceIds** | [**optional.Interface of []string**](string.md)| Performs the operation on the source ID specified. Enter multiple source IDs in comma-separated format. | 
 **sourceNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the source name specified. Enter multiple source names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 

### Return type

[**InlineResponse200166**](inline_response_200_166.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28VolumeSnapshotsTagsBatchPut**
> InlineResponse200162 Api28VolumeSnapshotsTagsBatchPut(ctx, body, optional)
Update tags

Updates tags.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]Tag**](Tag.md)| A list of tags to be created or modified. | 
 **optional** | ***VolumeSnapshotsApiApi28VolumeSnapshotsTagsBatchPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeSnapshotsApiApi28VolumeSnapshotsTagsBatchPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.**| Supplied by client during request or generated by server. | 
 **resourceIds** | [**optional.Interface of []string**](string.md)| A comma-separated list of resource IDs. The &#x60;resource_ids&#x60; and &#x60;resource_names&#x60; parameters cannot be provided together. | 
 **resourceNames** | [**optional.Interface of []string**](string.md)| A comma-separated list of resource names. The &#x60;resource_ids&#x60; and &#x60;resource_names&#x60; parameters cannot be provided together. | 

### Return type

[**InlineResponse200162**](inline_response_200_162.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28VolumeSnapshotsTagsDelete**
> Api28VolumeSnapshotsTagsDelete(ctx, optional)
Delete tags

Deletes specified tags on volume snapshot objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VolumeSnapshotsApiApi28VolumeSnapshotsTagsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeSnapshotsApiApi28VolumeSnapshotsTagsDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **keys** | [**optional.Interface of []string**](string.md)| A comma-separated list of tag keys. | 
 **namespaces** | [**optional.Interface of []string**](string.md)| A comma-separated list of namespaces. | 
 **resourceIds** | [**optional.Interface of []string**](string.md)| A comma-separated list of resource IDs. The &#x60;resource_ids&#x60; and &#x60;resource_names&#x60; parameters cannot be provided together. | 
 **resourceNames** | [**optional.Interface of []string**](string.md)| A comma-separated list of resource names. The &#x60;resource_ids&#x60; and &#x60;resource_names&#x60; parameters cannot be provided together. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28VolumeSnapshotsTagsGet**
> InlineResponse200161 Api28VolumeSnapshotsTagsGet(ctx, optional)
List tags

Displays the list of tags on volume snapshot objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VolumeSnapshotsApiApi28VolumeSnapshotsTagsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeSnapshotsApiApi28VolumeSnapshotsTagsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **continuationToken** | **optional.String**| A token used to retrieve the next page of data with some consistency guaranteed. The token is a Base64 encoded value. Set &#x60;continuation_token&#x60; to the system-generated token taken from the &#x60;x-next-token&#x60; header field of the response. A query has reached its last page when the response does not include a token. Pagination requires the &#x60;limit&#x60; and &#x60;continuation_token&#x60; query parameters. | 
 **filter** | **optional.String**| Narrows down the results to only the response objects that satisfy the filter criteria. | 
 **limit** | **optional.Int32**| Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size. | 
 **namespaces** | [**optional.Interface of []string**](string.md)| A comma-separated list of namespaces. | 
 **offset** | **optional.Int32**| The starting position based on the results of the query in relation to the full set of response objects returned. | 
 **resourceDestroyed** | **optional.Bool**| If set to &#x60;true&#x60;, returns only objects from destroyed resources. Returns an error if the name of a live resource is specified in the &#x60;resource_names&#x60; query parameter. If set to &#x60;false&#x60;, returns only objects from live resources. Returns an error if the name of a destroyed resource is specified in the &#x60;resource_names&#x60; query parameter. | 
 **resourceIds** | [**optional.Interface of []string**](string.md)| A comma-separated list of resource IDs. The &#x60;resource_ids&#x60; and &#x60;resource_names&#x60; parameters cannot be provided together. | 
 **resourceNames** | [**optional.Interface of []string**](string.md)| A comma-separated list of resource names. The &#x60;resource_ids&#x60; and &#x60;resource_names&#x60; parameters cannot be provided together. | 
 **sort** | [**optional.Interface of []string**](string.md)| Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values. | 
 **totalItemCount** | **optional.Bool**| If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;. | 

### Return type

[**InlineResponse200161**](inline_response_200_161.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28VolumeSnapshotsTransferGet**
> InlineResponse200167 Api28VolumeSnapshotsTransferGet(ctx, optional)
List volume snapshots with transfer statistics

Returns a list of volume snapshots and their transfer statistics.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VolumeSnapshotsApiApi28VolumeSnapshotsTransferGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeSnapshotsApiApi28VolumeSnapshotsTransferGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **destroyed** | **optional.Bool**| If set to &#x60;true&#x60;, lists only destroyed objects that are in the eradication pending state. If set to &#x60;false&#x60;, lists only objects that are not destroyed. For destroyed objects, the time remaining is displayed in milliseconds. | 
 **filter** | **optional.String**| Narrows down the results to only the response objects that satisfy the filter criteria. | 
 **ids** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique resource IDs specified. Enter multiple resource IDs in comma-separated format. The &#x60;ids&#x60; and &#x60;names&#x60; parameters cannot be provided together. | 
 **limit** | **optional.Int32**| Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size. | 
 **offset** | **optional.Int32**| The starting position based on the results of the query in relation to the full set of response objects returned. | 
 **sort** | [**optional.Interface of []string**](string.md)| Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values. | 
 **sourceIds** | [**optional.Interface of []string**](string.md)| Performs the operation on the source ID specified. Enter multiple source IDs in comma-separated format. | 
 **sourceNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the source name specified. Enter multiple source names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 
 **totalItemCount** | **optional.Bool**| If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;. | 
 **totalOnly** | **optional.Bool**| If set to &#x60;true&#x60;, returns the aggregate value of all items after filtering. Where it makes more sense, the average value is displayed instead. The values are displayed for each name where meaningful. If &#x60;total_only&#x3D;true&#x60;, the &#x60;items&#x60; list will be empty. | 

### Return type

[**InlineResponse200167**](inline_response_200_167.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

