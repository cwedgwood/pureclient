# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api28ProtectionGroupsDelete**](ProtectionGroupsApi.md#Api28ProtectionGroupsDelete) | **Delete** /api/2.8/protection-groups | Eradicate a protection group
[**Api28ProtectionGroupsGet**](ProtectionGroupsApi.md#Api28ProtectionGroupsGet) | **Get** /api/2.8/protection-groups | List protection groups
[**Api28ProtectionGroupsHostGroupsDelete**](ProtectionGroupsApi.md#Api28ProtectionGroupsHostGroupsDelete) | **Delete** /api/2.8/protection-groups/host-groups | Remove a host group from a protection group
[**Api28ProtectionGroupsHostGroupsGet**](ProtectionGroupsApi.md#Api28ProtectionGroupsHostGroupsGet) | **Get** /api/2.8/protection-groups/host-groups | List protection groups with host group members
[**Api28ProtectionGroupsHostGroupsPost**](ProtectionGroupsApi.md#Api28ProtectionGroupsHostGroupsPost) | **Post** /api/2.8/protection-groups/host-groups | Add a host group to a protection group
[**Api28ProtectionGroupsHostsDelete**](ProtectionGroupsApi.md#Api28ProtectionGroupsHostsDelete) | **Delete** /api/2.8/protection-groups/hosts | Remove a host from a protection group
[**Api28ProtectionGroupsHostsGet**](ProtectionGroupsApi.md#Api28ProtectionGroupsHostsGet) | **Get** /api/2.8/protection-groups/hosts | List protection groups with host members
[**Api28ProtectionGroupsHostsPost**](ProtectionGroupsApi.md#Api28ProtectionGroupsHostsPost) | **Post** /api/2.8/protection-groups/hosts | Add a host to a protection group
[**Api28ProtectionGroupsPatch**](ProtectionGroupsApi.md#Api28ProtectionGroupsPatch) | **Patch** /api/2.8/protection-groups | Manage a protection group
[**Api28ProtectionGroupsPerformanceReplicationByArrayGet**](ProtectionGroupsApi.md#Api28ProtectionGroupsPerformanceReplicationByArrayGet) | **Get** /api/2.8/protection-groups/performance/replication/by-array | List protection group replication performance data with array details
[**Api28ProtectionGroupsPerformanceReplicationGet**](ProtectionGroupsApi.md#Api28ProtectionGroupsPerformanceReplicationGet) | **Get** /api/2.8/protection-groups/performance/replication | List protection group replication performance data
[**Api28ProtectionGroupsPost**](ProtectionGroupsApi.md#Api28ProtectionGroupsPost) | **Post** /api/2.8/protection-groups | Create a protection group
[**Api28ProtectionGroupsSpaceGet**](ProtectionGroupsApi.md#Api28ProtectionGroupsSpaceGet) | **Get** /api/2.8/protection-groups/space | List protection group space information
[**Api28ProtectionGroupsTargetsDelete**](ProtectionGroupsApi.md#Api28ProtectionGroupsTargetsDelete) | **Delete** /api/2.8/protection-groups/targets | Removes a target from a protection group
[**Api28ProtectionGroupsTargetsGet**](ProtectionGroupsApi.md#Api28ProtectionGroupsTargetsGet) | **Get** /api/2.8/protection-groups/targets | List protection groups with targets
[**Api28ProtectionGroupsTargetsPatch**](ProtectionGroupsApi.md#Api28ProtectionGroupsTargetsPatch) | **Patch** /api/2.8/protection-groups/targets | Manage a protection group target
[**Api28ProtectionGroupsTargetsPost**](ProtectionGroupsApi.md#Api28ProtectionGroupsTargetsPost) | **Post** /api/2.8/protection-groups/targets | Add a target to a protection group
[**Api28ProtectionGroupsVolumesDelete**](ProtectionGroupsApi.md#Api28ProtectionGroupsVolumesDelete) | **Delete** /api/2.8/protection-groups/volumes | Remove a volume from a protection group
[**Api28ProtectionGroupsVolumesGet**](ProtectionGroupsApi.md#Api28ProtectionGroupsVolumesGet) | **Get** /api/2.8/protection-groups/volumes | List protection groups with volume members
[**Api28ProtectionGroupsVolumesPost**](ProtectionGroupsApi.md#Api28ProtectionGroupsVolumesPost) | **Post** /api/2.8/protection-groups/volumes | Add a volume to a protection group

# **Api28ProtectionGroupsDelete**
> Api28ProtectionGroupsDelete(ctx, optional)
Eradicate a protection group

Eradicates a protection group that has been destroyed and is pending eradication. Eradicated protection groups cannot be recovered. Protection groups are destroyed through the PATCH method. The`ids` or `names` parameter is required, but cannot be set together.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **names** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsGet**
> InlineResponse200117 Api28ProtectionGroupsGet(ctx, optional)
List protection groups

Returns a list of protection groups, including their associated source arrays, replication targets, hosts, host groups, and volumes. The list includes protection groups that were created on the local array to replicate snapshot data to other arrays or offload targets, created on a remote array and replicated asynchronously to this array, or created inside a pod on a remote array and stretched to the local array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **continuationToken** | **optional.String**| A token used to retrieve the next page of data with some consistency guaranteed. The token is a Base64 encoded value. Set &#x60;continuation_token&#x60; to the system-generated token taken from the &#x60;x-next-token&#x60; header field of the response. A query has reached its last page when the response does not include a token. Pagination requires the &#x60;limit&#x60; and &#x60;continuation_token&#x60; query parameters. | 
 **destroyed** | **optional.Bool**| If set to &#x60;true&#x60;, lists only destroyed objects that are in the eradication pending state. If set to &#x60;false&#x60;, lists only objects that are not destroyed. For destroyed objects, the time remaining is displayed in milliseconds. | 
 **filter** | **optional.String**| Narrows down the results to only the response objects that satisfy the filter criteria. | 
 **limit** | **optional.Int32**| Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size. | 
 **names** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 
 **offset** | **optional.Int32**| The starting position based on the results of the query in relation to the full set of response objects returned. | 
 **sort** | [**optional.Interface of []string**](string.md)| Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values. | 
 **totalItemCount** | **optional.Bool**| If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;. | 
 **totalOnly** | **optional.Bool**| If set to &#x60;true&#x60;, returns the aggregate value of all items after filtering. Where it makes more sense, the average value is displayed instead. The values are displayed for each name where meaningful. If &#x60;total_only&#x3D;true&#x60;, the &#x60;items&#x60; list will be empty. | 

### Return type

[**InlineResponse200117**](inline_response_200_117.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsHostGroupsDelete**
> Api28ProtectionGroupsHostGroupsDelete(ctx, optional)
Remove a host group from a protection group

Removes a host group member from a protection group. After the member has been removed, it is no longer protected by the group. Any protection group snapshots that were taken before the member was removed will not be affected. Removing a member from a protection group does not delete the member from the array, and the member can be added back to the protection group at any time. The `group_names` parameter represents the name of the protection group, and the `member_names` parameter represents the name of the host group. The `group_names` and `member_names` parameters are required and must be set together.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsHostGroupsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsHostGroupsDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **groupNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique group name specified. Examples of groups include host groups, pods, protection groups, and volume groups. Enter multiple names in comma-separated format. For example, &#x60;hgroup01,hgroup02&#x60;. | 
 **memberNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique member name specified. Examples of members include volumes, hosts, host groups, and directories. Enter multiple names in comma-separated format. For example, &#x60;vol01,vol02&#x60;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsHostGroupsGet**
> InlineResponse20072 Api28ProtectionGroupsHostGroupsGet(ctx, optional)
List protection groups with host group members

Returns a list of protection groups that have host group members.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsHostGroupsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsHostGroupsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **continuationToken** | **optional.String**| A token used to retrieve the next page of data with some consistency guaranteed. The token is a Base64 encoded value. Set &#x60;continuation_token&#x60; to the system-generated token taken from the &#x60;x-next-token&#x60; header field of the response. A query has reached its last page when the response does not include a token. Pagination requires the &#x60;limit&#x60; and &#x60;continuation_token&#x60; query parameters. | 
 **filter** | **optional.String**| Narrows down the results to only the response objects that satisfy the filter criteria. | 
 **groupNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique group name specified. Examples of groups include host groups, pods, protection groups, and volume groups. Enter multiple names in comma-separated format. For example, &#x60;hgroup01,hgroup02&#x60;. | 
 **limit** | **optional.Int32**| Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size. | 
 **memberNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique member name specified. Examples of members include volumes, hosts, host groups, and directories. Enter multiple names in comma-separated format. For example, &#x60;vol01,vol02&#x60;. | 
 **offset** | **optional.Int32**| The starting position based on the results of the query in relation to the full set of response objects returned. | 
 **sort** | [**optional.Interface of []string**](string.md)| Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values. | 
 **totalItemCount** | **optional.Bool**| If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;. | 

### Return type

[**InlineResponse20072**](inline_response_200_72.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsHostGroupsPost**
> InlineResponse20073 Api28ProtectionGroupsHostGroupsPost(ctx, optional)
Add a host group to a protection group

Adds a host group member to a protection group. Members that are already in the protection group are not affected. For asynchronous replication, only members of the same type can belong to a protection group. The `group_names` parameter represents the name of the protection group, and the `member_names` parameter represents the name of the host group. The `group_names` and `member_names` parameters are required and must be set together.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsHostGroupsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsHostGroupsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **groupNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique group name specified. Examples of groups include host groups, pods, protection groups, and volume groups. Enter multiple names in comma-separated format. For example, &#x60;hgroup01,hgroup02&#x60;. | 
 **memberNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique member name specified. Examples of members include volumes, hosts, host groups, and directories. Enter multiple names in comma-separated format. For example, &#x60;vol01,vol02&#x60;. | 

### Return type

[**InlineResponse20073**](inline_response_200_73.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsHostsDelete**
> Api28ProtectionGroupsHostsDelete(ctx, optional)
Remove a host from a protection group

Removes a host member from a protection group. After the member has been removed, it is no longer protected by the group. Any protection group snapshots that were taken before the member was removed will not be affected. Removing a member from a protection group does not delete the member from the array, and the member can be added back to the protection group at any time. The `group_names` parameter represents the name of the protection group, and the `member_names` parameter represents the name of the host. The `group_names` and `member_names` parameters are required and must be set together.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsHostsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsHostsDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **groupNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique group name specified. Examples of groups include host groups, pods, protection groups, and volume groups. Enter multiple names in comma-separated format. For example, &#x60;hgroup01,hgroup02&#x60;. | 
 **memberNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique member name specified. Examples of members include volumes, hosts, host groups, and directories. Enter multiple names in comma-separated format. For example, &#x60;vol01,vol02&#x60;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsHostsGet**
> InlineResponse20072 Api28ProtectionGroupsHostsGet(ctx, optional)
List protection groups with host members

Returns a list of protection groups that have host members.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsHostsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsHostsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **continuationToken** | **optional.String**| A token used to retrieve the next page of data with some consistency guaranteed. The token is a Base64 encoded value. Set &#x60;continuation_token&#x60; to the system-generated token taken from the &#x60;x-next-token&#x60; header field of the response. A query has reached its last page when the response does not include a token. Pagination requires the &#x60;limit&#x60; and &#x60;continuation_token&#x60; query parameters. | 
 **filter** | **optional.String**| Narrows down the results to only the response objects that satisfy the filter criteria. | 
 **groupNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique group name specified. Examples of groups include host groups, pods, protection groups, and volume groups. Enter multiple names in comma-separated format. For example, &#x60;hgroup01,hgroup02&#x60;. | 
 **limit** | **optional.Int32**| Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size. | 
 **memberNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique member name specified. Examples of members include volumes, hosts, host groups, and directories. Enter multiple names in comma-separated format. For example, &#x60;vol01,vol02&#x60;. | 
 **offset** | **optional.Int32**| The starting position based on the results of the query in relation to the full set of response objects returned. | 
 **sort** | [**optional.Interface of []string**](string.md)| Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values. | 
 **totalItemCount** | **optional.Bool**| If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;. | 

### Return type

[**InlineResponse20072**](inline_response_200_72.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsHostsPost**
> InlineResponse20073 Api28ProtectionGroupsHostsPost(ctx, optional)
Add a host to a protection group

Adds a host member to a protection group. Members that are already in the protection group are not affected. For asynchronous replication, only members of the same type can belong to a protection group. The `group_names` parameter represents the name of the protection group, and the `member_names` parameter represents the name of the host. The `group_names` and `member_names` parameters are required and must be set together.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsHostsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsHostsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **groupNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique group name specified. Examples of groups include host groups, pods, protection groups, and volume groups. Enter multiple names in comma-separated format. For example, &#x60;hgroup01,hgroup02&#x60;. | 
 **memberNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique member name specified. Examples of members include volumes, hosts, host groups, and directories. Enter multiple names in comma-separated format. For example, &#x60;vol01,vol02&#x60;. | 

### Return type

[**InlineResponse20073**](inline_response_200_73.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsPatch**
> InlineResponse200118 Api28ProtectionGroupsPatch(ctx, body, optional)
Manage a protection group

Configures the protection group schedules to generate and replicate snapshots to another array or to an external storage system. Also renames or destroys a protection group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProtectionGroup**](ProtectionGroup.md)|  | 
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsPatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.**| Supplied by client during request or generated by server. | 
 **names** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 

### Return type

[**InlineResponse200118**](inline_response_200_118.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsPerformanceReplicationByArrayGet**
> InlineResponse200120 Api28ProtectionGroupsPerformanceReplicationByArrayGet(ctx, optional)
List protection group replication performance data with array details

Returns the total number of bytes of replication data transmitted and received per second. The data is grouped by protection group and includes the names of the source array and targets for each protection group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsPerformanceReplicationByArrayGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsPerformanceReplicationByArrayGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **destroyed** | **optional.Bool**| If set to &#x60;true&#x60;, lists only destroyed objects that are in the eradication pending state. If set to &#x60;false&#x60;, lists only objects that are not destroyed. For destroyed objects, the time remaining is displayed in milliseconds. | 
 **filter** | **optional.String**| Narrows down the results to only the response objects that satisfy the filter criteria. | 
 **endTime** | **optional.Int64**| Displays historical performance data for the specified time window, where &#x60;start_time&#x60; is the beginning of the time window, and &#x60;end_time&#x60; is the end of the time window. The &#x60;start_time&#x60; and &#x60;end_time&#x60; parameters are specified in milliseconds since the UNIX epoch. If &#x60;start_time&#x60; is not specified, the start time will default to one resolution before the end time, meaning that the most recent sample of performance data will be displayed. If &#x60;end_time&#x60;is not specified, the end time will default to the current time. Include the &#x60;resolution&#x60; parameter to display the performance data at the specified resolution. If not specified, &#x60;resolution&#x60; defaults to the lowest valid resolution. | 
 **resolution** | **optional.Int64**| The number of milliseconds between samples of historical data. For array-wide performance metrics (&#x60;/arrays/performance&#x60; endpoint), valid values are &#x60;1000&#x60; (1 second), &#x60;30000&#x60; (30 seconds), &#x60;300000&#x60; (5 minutes), &#x60;1800000&#x60; (30 minutes), &#x60;7200000&#x60; (2 hours), &#x60;28800000&#x60; (8 hours), and &#x60;86400000&#x60; (24 hours). For performance metrics on storage objects (&#x60;&lt;object name&gt;/performance&#x60; endpoint), such as volumes, valid values are &#x60;30000&#x60; (30 seconds), &#x60;300000&#x60; (5 minutes), &#x60;1800000&#x60; (30 minutes), &#x60;7200000&#x60; (2 hours), &#x60;28800000&#x60; (8 hours), and &#x60;86400000&#x60; (24 hours). For space metrics, (&#x60;&lt;object name&gt;/space&#x60; endpoint), valid values are &#x60;300000&#x60; (5 minutes), &#x60;1800000&#x60; (30 minutes), &#x60;7200000&#x60; (2 hours), &#x60;28800000&#x60; (8 hours), and &#x60;86400000&#x60; (24 hours). Include the &#x60;start_time&#x60; parameter to display the performance data starting at the specified start time. If &#x60;start_time&#x60; is not specified, the start time will default to one resolution before the end time, meaning that the most recent sample of performance data will be displayed. Include the &#x60;end_time&#x60; parameter to display the performance data until the specified end time. If &#x60;end_time&#x60;is not specified, the end time will default to the current time. If the &#x60;resolution&#x60; parameter is not specified but either the &#x60;start_time&#x60; or &#x60;end_time&#x60; parameter is, then &#x60;resolution&#x60; will default to the lowest valid resolution. | 
 **startTime** | **optional.Int64**| Displays historical performance data for the specified time window, where &#x60;start_time&#x60; is the beginning of the time window, and &#x60;end_time&#x60; is the end of the time window. The &#x60;start_time&#x60; and &#x60;end_time&#x60; parameters are specified in milliseconds since the UNIX epoch. If &#x60;start_time&#x60; is not specified, the start time will default to one resolution before the end time, meaning that the most recent sample of performance data will be displayed. If &#x60;end_time&#x60;is not specified, the end time will default to the current time. Include the &#x60;resolution&#x60; parameter to display the performance data at the specified resolution. If not specified, &#x60;resolution&#x60; defaults to the lowest valid resolution. | 
 **limit** | **optional.Int32**| Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size. | 
 **offset** | **optional.Int32**| The starting position based on the results of the query in relation to the full set of response objects returned. | 
 **sort** | [**optional.Interface of []string**](string.md)| Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values. | 
 **totalItemCount** | **optional.Bool**| If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;. | 

### Return type

[**InlineResponse200120**](inline_response_200_120.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsPerformanceReplicationGet**
> InlineResponse200119 Api28ProtectionGroupsPerformanceReplicationGet(ctx, optional)
List protection group replication performance data

Returns the total number of bytes of replication data transmitted and received per second. The data is grouped by protection group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsPerformanceReplicationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsPerformanceReplicationGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **destroyed** | **optional.Bool**| If set to &#x60;true&#x60;, lists only destroyed objects that are in the eradication pending state. If set to &#x60;false&#x60;, lists only objects that are not destroyed. For destroyed objects, the time remaining is displayed in milliseconds. | 
 **filter** | **optional.String**| Narrows down the results to only the response objects that satisfy the filter criteria. | 
 **endTime** | **optional.Int64**| Displays historical performance data for the specified time window, where &#x60;start_time&#x60; is the beginning of the time window, and &#x60;end_time&#x60; is the end of the time window. The &#x60;start_time&#x60; and &#x60;end_time&#x60; parameters are specified in milliseconds since the UNIX epoch. If &#x60;start_time&#x60; is not specified, the start time will default to one resolution before the end time, meaning that the most recent sample of performance data will be displayed. If &#x60;end_time&#x60;is not specified, the end time will default to the current time. Include the &#x60;resolution&#x60; parameter to display the performance data at the specified resolution. If not specified, &#x60;resolution&#x60; defaults to the lowest valid resolution. | 
 **resolution** | **optional.Int64**| The number of milliseconds between samples of historical data. For array-wide performance metrics (&#x60;/arrays/performance&#x60; endpoint), valid values are &#x60;1000&#x60; (1 second), &#x60;30000&#x60; (30 seconds), &#x60;300000&#x60; (5 minutes), &#x60;1800000&#x60; (30 minutes), &#x60;7200000&#x60; (2 hours), &#x60;28800000&#x60; (8 hours), and &#x60;86400000&#x60; (24 hours). For performance metrics on storage objects (&#x60;&lt;object name&gt;/performance&#x60; endpoint), such as volumes, valid values are &#x60;30000&#x60; (30 seconds), &#x60;300000&#x60; (5 minutes), &#x60;1800000&#x60; (30 minutes), &#x60;7200000&#x60; (2 hours), &#x60;28800000&#x60; (8 hours), and &#x60;86400000&#x60; (24 hours). For space metrics, (&#x60;&lt;object name&gt;/space&#x60; endpoint), valid values are &#x60;300000&#x60; (5 minutes), &#x60;1800000&#x60; (30 minutes), &#x60;7200000&#x60; (2 hours), &#x60;28800000&#x60; (8 hours), and &#x60;86400000&#x60; (24 hours). Include the &#x60;start_time&#x60; parameter to display the performance data starting at the specified start time. If &#x60;start_time&#x60; is not specified, the start time will default to one resolution before the end time, meaning that the most recent sample of performance data will be displayed. Include the &#x60;end_time&#x60; parameter to display the performance data until the specified end time. If &#x60;end_time&#x60;is not specified, the end time will default to the current time. If the &#x60;resolution&#x60; parameter is not specified but either the &#x60;start_time&#x60; or &#x60;end_time&#x60; parameter is, then &#x60;resolution&#x60; will default to the lowest valid resolution. | 
 **startTime** | **optional.Int64**| Displays historical performance data for the specified time window, where &#x60;start_time&#x60; is the beginning of the time window, and &#x60;end_time&#x60; is the end of the time window. The &#x60;start_time&#x60; and &#x60;end_time&#x60; parameters are specified in milliseconds since the UNIX epoch. If &#x60;start_time&#x60; is not specified, the start time will default to one resolution before the end time, meaning that the most recent sample of performance data will be displayed. If &#x60;end_time&#x60;is not specified, the end time will default to the current time. Include the &#x60;resolution&#x60; parameter to display the performance data at the specified resolution. If not specified, &#x60;resolution&#x60; defaults to the lowest valid resolution. | 
 **limit** | **optional.Int32**| Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size. | 
 **offset** | **optional.Int32**| The starting position based on the results of the query in relation to the full set of response objects returned. | 
 **sort** | [**optional.Interface of []string**](string.md)| Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values. | 
 **totalItemCount** | **optional.Bool**| If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;. | 

### Return type

[**InlineResponse200119**](inline_response_200_119.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsPost**
> InlineResponse200118 Api28ProtectionGroupsPost(ctx, optional)
Create a protection group

Creates a protection group on the local array for asynchronous replication.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **names** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 
 **sourceNames** | [**optional.Interface of []string**](string.md)| The name of the protection group or protection group snapshot to be copied into a new or existing protection group. If the destination protection group and all of its volumes already exist, include the &#x60;overwrite&#x60; parameter to overwrite all of the existing volumes with the snapshot contents. If including the &#x60;overwrite&#x60; parameter, the names of the volumes that are being overwritten must match the names of the volumes that are being restored. If the source is a protection group, the latest snapshot of the protection group will be used as the source during the copy operation. | 
 **overwrite** | **optional.Bool**| If set to &#x60;true&#x60;, overwrites an existing volume during a volume copy operation. If set to &#x60;false&#x60; or not set at all and the target name is an existing volume, the volume copy operation fails. Required if the &#x60;source: id&#x60; or &#x60;source: name&#x60; body parameter is set and the source overwrites an existing volume during the volume copy operation. | 

### Return type

[**InlineResponse200118**](inline_response_200_118.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsSpaceGet**
> InlineResponse20076 Api28ProtectionGroupsSpaceGet(ctx, optional)
List protection group space information

Returns provisioned (virtual) size and physical storage consumption data for each protection group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsSpaceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsSpaceGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **destroyed** | **optional.Bool**| If set to &#x60;true&#x60;, lists only destroyed objects that are in the eradication pending state. If set to &#x60;false&#x60;, lists only objects that are not destroyed. For destroyed objects, the time remaining is displayed in milliseconds. | 
 **filter** | **optional.String**| Narrows down the results to only the response objects that satisfy the filter criteria. | 
 **limit** | **optional.Int32**| Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size. | 
 **names** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 
 **offset** | **optional.Int32**| The starting position based on the results of the query in relation to the full set of response objects returned. | 
 **sort** | [**optional.Interface of []string**](string.md)| Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values. | 
 **totalItemCount** | **optional.Bool**| If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;. | 

### Return type

[**InlineResponse20076**](inline_response_200_76.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsTargetsDelete**
> Api28ProtectionGroupsTargetsDelete(ctx, optional)
Removes a target from a protection group

Removes an array or offload target from a protection group. The `group_names` parameter represents the name of the protection group. The `member_names` parameter represents the name of the array or offload target that is being removed from the protection group. The `group_names` and `member_names` parameters are required and must be set together.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsTargetsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsTargetsDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **groupNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique group name specified. Examples of groups include host groups, pods, protection groups, and volume groups. Enter multiple names in comma-separated format. For example, &#x60;hgroup01,hgroup02&#x60;. | 
 **memberNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique member name specified. Examples of members include volumes, hosts, host groups, and directories. Enter multiple names in comma-separated format. For example, &#x60;vol01,vol02&#x60;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsTargetsGet**
> InlineResponse200121 Api28ProtectionGroupsTargetsGet(ctx, optional)
List protection groups with targets

Returns a list of protection groups that have target arrays or offload targets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsTargetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsTargetsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **continuationToken** | **optional.String**| A token used to retrieve the next page of data with some consistency guaranteed. The token is a Base64 encoded value. Set &#x60;continuation_token&#x60; to the system-generated token taken from the &#x60;x-next-token&#x60; header field of the response. A query has reached its last page when the response does not include a token. Pagination requires the &#x60;limit&#x60; and &#x60;continuation_token&#x60; query parameters. | 
 **filter** | **optional.String**| Narrows down the results to only the response objects that satisfy the filter criteria. | 
 **groupNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique group name specified. Examples of groups include host groups, pods, protection groups, and volume groups. Enter multiple names in comma-separated format. For example, &#x60;hgroup01,hgroup02&#x60;. | 
 **limit** | **optional.Int32**| Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size. | 
 **memberNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique member name specified. Examples of members include volumes, hosts, host groups, and directories. Enter multiple names in comma-separated format. For example, &#x60;vol01,vol02&#x60;. | 
 **offset** | **optional.Int32**| The starting position based on the results of the query in relation to the full set of response objects returned. | 
 **sort** | [**optional.Interface of []string**](string.md)| Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values. | 
 **totalItemCount** | **optional.Bool**| If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;. | 

### Return type

[**InlineResponse200121**](inline_response_200_121.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsTargetsPatch**
> InlineResponse200122 Api28ProtectionGroupsTargetsPatch(ctx, body, optional)
Manage a protection group target

Allows the source array to replicate protection group data to the target array, or disallows the source array from replicating protection group data to the target array. The `allowed` parameter must be set from the target array. The `group_names` parameter represents the name of the protection group. The `allowed` and `group_names` parameters are required and must be set together. Offload targets do not support the `allowed` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProtectiongroupsTargetsBody**](ProtectiongroupsTargetsBody.md)|  | 
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsTargetsPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsTargetsPatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.**| Supplied by client during request or generated by server. | 
 **groupNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique group name specified. Examples of groups include host groups, pods, protection groups, and volume groups. Enter multiple names in comma-separated format. For example, &#x60;hgroup01,hgroup02&#x60;. | 
 **memberNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique member name specified. Examples of members include volumes, hosts, host groups, and directories. Enter multiple names in comma-separated format. For example, &#x60;vol01,vol02&#x60;. | 

### Return type

[**InlineResponse200122**](inline_response_200_122.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsTargetsPost**
> InlineResponse200122 Api28ProtectionGroupsTargetsPost(ctx, optional)
Add a target to a protection group

Adds an array or offload target to a protection group. The `group_names` parameter represents the name of the protection group. The `member_names` parameter represents the name of the array or offload target that is being added to the protection group. The `group_names` and `member_names` parameters are required and must be set together.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsTargetsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsTargetsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **groupNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique group name specified. Examples of groups include host groups, pods, protection groups, and volume groups. Enter multiple names in comma-separated format. For example, &#x60;hgroup01,hgroup02&#x60;. | 
 **memberNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique member name specified. Examples of members include volumes, hosts, host groups, and directories. Enter multiple names in comma-separated format. For example, &#x60;vol01,vol02&#x60;. | 

### Return type

[**InlineResponse200122**](inline_response_200_122.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsVolumesDelete**
> Api28ProtectionGroupsVolumesDelete(ctx, optional)
Remove a volume from a protection group

Removes a volume member from a protection group. After the member has been removed, it is no longer protected by the group. Any protection group snapshots that were taken before the member was removed will not be affected. Removing a member from a protection group does not delete the member from the array, and the member can be added back to the protection group at any time. The `group_names` parameter represents the name of the protection group, and the `member_names` parameter represents the name of the volume. The `group_names` and `member_names` parameters are required and must be set together.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsVolumesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsVolumesDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **groupNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique group name specified. Examples of groups include host groups, pods, protection groups, and volume groups. Enter multiple names in comma-separated format. For example, &#x60;hgroup01,hgroup02&#x60;. | 
 **memberNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique member name specified. Examples of members include volumes, hosts, host groups, and directories. Enter multiple names in comma-separated format. For example, &#x60;vol01,vol02&#x60;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsVolumesGet**
> InlineResponse20072 Api28ProtectionGroupsVolumesGet(ctx, optional)
List protection groups with volume members

Returns a list of protection groups that have volume members.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsVolumesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsVolumesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **continuationToken** | **optional.String**| A token used to retrieve the next page of data with some consistency guaranteed. The token is a Base64 encoded value. Set &#x60;continuation_token&#x60; to the system-generated token taken from the &#x60;x-next-token&#x60; header field of the response. A query has reached its last page when the response does not include a token. Pagination requires the &#x60;limit&#x60; and &#x60;continuation_token&#x60; query parameters. | 
 **filter** | **optional.String**| Narrows down the results to only the response objects that satisfy the filter criteria. | 
 **groupNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique group name specified. Examples of groups include host groups, pods, protection groups, and volume groups. Enter multiple names in comma-separated format. For example, &#x60;hgroup01,hgroup02&#x60;. | 
 **limit** | **optional.Int32**| Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size. | 
 **memberNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique member name specified. Examples of members include volumes, hosts, host groups, and directories. Enter multiple names in comma-separated format. For example, &#x60;vol01,vol02&#x60;. | 
 **offset** | **optional.Int32**| The starting position based on the results of the query in relation to the full set of response objects returned. | 
 **sort** | [**optional.Interface of []string**](string.md)| Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values. | 
 **totalItemCount** | **optional.Bool**| If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;. | 

### Return type

[**InlineResponse20072**](inline_response_200_72.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28ProtectionGroupsVolumesPost**
> InlineResponse20073 Api28ProtectionGroupsVolumesPost(ctx, optional)
Add a volume to a protection group

Adds a volume member to a protection group. Members that are already in the protection group are not affected. For asynchronous replication, only members of the same type can belong to a protection group. The `group_names` parameter represents the name of the protection group, and the `member_names` parameter represents the name of the volume. The `group_names` and `member_names` parameters are required and must be set together.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionGroupsApiApi28ProtectionGroupsVolumesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionGroupsApiApi28ProtectionGroupsVolumesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **groupNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique group name specified. Examples of groups include host groups, pods, protection groups, and volume groups. Enter multiple names in comma-separated format. For example, &#x60;hgroup01,hgroup02&#x60;. | 
 **memberNames** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique member name specified. Examples of members include volumes, hosts, host groups, and directories. Enter multiple names in comma-separated format. For example, &#x60;vol01,vol02&#x60;. | 

### Return type

[**InlineResponse20073**](inline_response_200_73.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

