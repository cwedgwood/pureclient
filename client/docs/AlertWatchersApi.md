# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api28AlertWatchersDelete**](AlertWatchersApi.md#Api28AlertWatchersDelete) | **Delete** /api/2.8/alert-watchers | Delete alert watcher
[**Api28AlertWatchersGet**](AlertWatchersApi.md#Api28AlertWatchersGet) | **Get** /api/2.8/alert-watchers | List alert watchers
[**Api28AlertWatchersPatch**](AlertWatchersApi.md#Api28AlertWatchersPatch) | **Patch** /api/2.8/alert-watchers | Modify alert watcher
[**Api28AlertWatchersPost**](AlertWatchersApi.md#Api28AlertWatchersPost) | **Post** /api/2.8/alert-watchers | Create alert watcher
[**Api28AlertWatchersTestGet**](AlertWatchersApi.md#Api28AlertWatchersTestGet) | **Get** /api/2.8/alert-watchers/test | List alert watcher test

# **Api28AlertWatchersDelete**
> Api28AlertWatchersDelete(ctx, names, optional)
Delete alert watcher

Delete alert watcher email address from the list of alert watchers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **names** | [**[]string**](string.md)| Performs the operation on the unique name specified. For example, &#x60;name01&#x60;. Enter multiple names in comma-separated format. | 
 **optional** | ***AlertWatchersApiApi28AlertWatchersDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertWatchersApiApi28AlertWatchersDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28AlertWatchersGet**
> InlineResponse20014 Api28AlertWatchersGet(ctx, optional)
List alert watchers

Displays alert watcher email addressess and indicates whether they are `enabled`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertWatchersApiApi28AlertWatchersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertWatchersApiApi28AlertWatchersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **continuationToken** | **optional.String**| A token used to retrieve the next page of data with some consistency guaranteed. The token is a Base64 encoded value. Set &#x60;continuation_token&#x60; to the system-generated token taken from the &#x60;x-next-token&#x60; header field of the response. A query has reached its last page when the response does not include a token. Pagination requires the &#x60;limit&#x60; and &#x60;continuation_token&#x60; query parameters. | 
 **filter** | **optional.String**| Narrows down the results to only the response objects that satisfy the filter criteria. | 
 **limit** | **optional.Int32**| Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size. | 
 **names** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 
 **offset** | **optional.Int32**| The starting position based on the results of the query in relation to the full set of response objects returned. | 
 **sort** | [**optional.Interface of []string**](string.md)| Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values. | 
 **totalItemCount** | **optional.Bool**| If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;. | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28AlertWatchersPatch**
> InlineResponse20015 Api28AlertWatchersPatch(ctx, body, names, optional)
Modify alert watcher

Modify alert watcher email address by enabling or disabling it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Model28AlertwatchersBody1**](Model28AlertwatchersBody1.md)|  | 
  **names** | [**[]string**](string.md)| Performs the operation on the unique name specified. For example, &#x60;name01&#x60;. Enter multiple names in comma-separated format. | 
 **optional** | ***AlertWatchersApiApi28AlertWatchersPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertWatchersApiApi28AlertWatchersPatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **optional.**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.**| Supplied by client during request or generated by server. | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28AlertWatchersPost**
> InlineResponse20015 Api28AlertWatchersPost(ctx, names, optional)
Create alert watcher

Creates one or more alert watcher email addresses, adding them to the list of alert watchers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **names** | [**[]string**](string.md)| Performs the operation on the unique name specified. For example, &#x60;name01&#x60;. Enter multiple names in comma-separated format. | 
 **optional** | ***AlertWatchersApiApi28AlertWatchersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertWatchersApiApi28AlertWatchersPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Model28AlertwatchersBody**](Model28AlertwatchersBody.md)|  | 
 **authorization** | **optional.**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.**| Supplied by client during request or generated by server. | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28AlertWatchersTestGet**
> InlineResponse20016 Api28AlertWatchersTestGet(ctx, optional)
List alert watcher test

Displays alert watcher email test results.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertWatchersApiApi28AlertWatchersTestGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertWatchersApiApi28AlertWatchersTestGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **filter** | **optional.String**| Narrows down the results to only the response objects that satisfy the filter criteria. | 
 **limit** | **optional.Int32**| Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size. | 
 **names** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 
 **offset** | **optional.Int32**| The starting position based on the results of the query in relation to the full set of response objects returned. | 
 **sort** | [**optional.Interface of []string**](string.md)| Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values. | 
 **totalItemCount** | **optional.Bool**| If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;. | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

