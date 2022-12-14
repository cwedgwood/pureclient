# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api28SmiSGet**](SMISApi.md#Api28SmiSGet) | **Get** /api/2.8/smi-s | List SMI-S settings
[**Api28SmiSPatch**](SMISApi.md#Api28SmiSPatch) | **Patch** /api/2.8/smi-s | Modify SLP and SMI-S

# **Api28SmiSGet**
> InlineResponse200136 Api28SmiSGet(ctx, optional)
List SMI-S settings

Displays the SMI-S settings, including the name of the array and whether SLP and WBEM-HTTPS are enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SMISApiApi28SmiSGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SMISApiApi28SmiSGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **filter** | **optional.String**| Narrows down the results to only the response objects that satisfy the filter criteria. | 
 **limit** | **optional.Int32**| Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size. | 
 **offset** | **optional.Int32**| The starting position based on the results of the query in relation to the full set of response objects returned. | 
 **sort** | [**optional.Interface of []string**](string.md)| Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values. | 
 **totalItemCount** | **optional.Bool**| If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;. | 

### Return type

[**InlineResponse200136**](inline_response_200_136.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28SmiSPatch**
> InlineResponse200137 Api28SmiSPatch(ctx, body, optional)
Modify SLP and SMI-S

Modifies the Service Location Protocol (SLP) and the SMI-S provider, enabling or disabling them.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmiS**](SmiS.md)|  | 
 **optional** | ***SMISApiApi28SmiSPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SMISApiApi28SmiSPatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.**| Supplied by client during request or generated by server. | 

### Return type

[**InlineResponse200137**](inline_response_200_137.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

