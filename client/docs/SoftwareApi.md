# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api28SoftwareBundleGet**](SoftwareApi.md#Api28SoftwareBundleGet) | **Get** /api/2.8/software-bundle | List software-bundle
[**Api28SoftwareBundlePost**](SoftwareApi.md#Api28SoftwareBundlePost) | **Post** /api/2.8/software-bundle | Create software-bundle
[**Api28SoftwareGet**](SoftwareApi.md#Api28SoftwareGet) | **Get** /api/2.8/software | List software
[**Api28SoftwareInstallationStepsGet**](SoftwareApi.md#Api28SoftwareInstallationStepsGet) | **Get** /api/2.8/software-installation-steps | List software upgrade steps
[**Api28SoftwareInstallationsGet**](SoftwareApi.md#Api28SoftwareInstallationsGet) | **Get** /api/2.8/software-installations | List software upgrades
[**Api28SoftwareInstallationsPatch**](SoftwareApi.md#Api28SoftwareInstallationsPatch) | **Patch** /api/2.8/software-installations | Modify software upgrade
[**Api28SoftwareInstallationsPost**](SoftwareApi.md#Api28SoftwareInstallationsPost) | **Post** /api/2.8/software-installations | Create a software upgrade

# **Api28SoftwareBundleGet**
> InlineResponse200146 Api28SoftwareBundleGet(ctx, optional)
List software-bundle

Displays a list of software bundles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoftwareApiApi28SoftwareBundleGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwareApiApi28SoftwareBundleGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Narrows down the results to only the response objects that satisfy the filter criteria. | 
 **ids** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique resource IDs specified. Enter multiple resource IDs in comma-separated format. The &#x60;ids&#x60; and &#x60;names&#x60; parameters cannot be provided together. | 
 **limit** | **optional.Int32**| Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size. | 
 **offset** | **optional.Int32**| The starting position based on the results of the query in relation to the full set of response objects returned. | 
 **sort** | [**optional.Interface of []string**](string.md)| Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values. | 
 **totalItemCount** | **optional.Bool**| If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;. | 

### Return type

[**InlineResponse200146**](inline_response_200_146.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28SoftwareBundlePost**
> InlineResponse200147 Api28SoftwareBundlePost(ctx, body)
Create software-bundle

Creates and initiates a software bundle download.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Model28SoftwarebundleBody**](Model28SoftwarebundleBody.md)|  | 

### Return type

[**InlineResponse200147**](inline_response_200_147.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28SoftwareGet**
> InlineResponse200145 Api28SoftwareGet(ctx, optional)
List software

Displays a list of software versions available for update and the installation status of each software version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoftwareApiApi28SoftwareGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwareApiApi28SoftwareGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **filter** | **optional.String**| Narrows down the results to only the response objects that satisfy the filter criteria. | 
 **ids** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique resource IDs specified. Enter multiple resource IDs in comma-separated format. The &#x60;ids&#x60; and &#x60;names&#x60; parameters cannot be provided together. | 
 **limit** | **optional.Int32**| Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size. | 
 **names** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 
 **offset** | **optional.Int32**| The starting position based on the results of the query in relation to the full set of response objects returned. | 
 **sort** | [**optional.Interface of []string**](string.md)| Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values. | 
 **totalItemCount** | **optional.Bool**| If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;. | 
 **versions** | [**optional.Interface of []string**](string.md)| A comma-separated list of versions. | 

### Return type

[**InlineResponse200145**](inline_response_200_145.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28SoftwareInstallationStepsGet**
> InlineResponse200150 Api28SoftwareInstallationStepsGet(ctx, optional)
List software upgrade steps

Displays a list of currently running and completed software upgrade steps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoftwareApiApi28SoftwareInstallationStepsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwareApiApi28SoftwareInstallationStepsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **continuationToken** | **optional.String**| A token used to retrieve the next page of data with some consistency guaranteed. The token is a Base64 encoded value. Set &#x60;continuation_token&#x60; to the system-generated token taken from the &#x60;x-next-token&#x60; header field of the response. A query has reached its last page when the response does not include a token. Pagination requires the &#x60;limit&#x60; and &#x60;continuation_token&#x60; query parameters. | 
 **filter** | **optional.String**| Narrows down the results to only the response objects that satisfy the filter criteria. | 
 **ids** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique resource IDs specified. Enter multiple resource IDs in comma-separated format. The &#x60;ids&#x60; and &#x60;names&#x60; parameters cannot be provided together. | 
 **limit** | **optional.Int32**| Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size. | 
 **names** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 
 **offset** | **optional.Int32**| The starting position based on the results of the query in relation to the full set of response objects returned. | 
 **softwareInstallationIds** | [**optional.Interface of []string**](string.md)| A comma-separated list of software installation IDs. | 
 **sort** | [**optional.Interface of []string**](string.md)| Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values. | 
 **totalItemCount** | **optional.Bool**| If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;. | 

### Return type

[**InlineResponse200150**](inline_response_200_150.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28SoftwareInstallationsGet**
> InlineResponse200148 Api28SoftwareInstallationsGet(ctx, optional)
List software upgrades

Displays a list of software upgrades. This returns both currently running and past upgrades.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoftwareApiApi28SoftwareInstallationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwareApiApi28SoftwareInstallationsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.String**| Supplied by client during request or generated by server. | 
 **continuationToken** | **optional.String**| A token used to retrieve the next page of data with some consistency guaranteed. The token is a Base64 encoded value. Set &#x60;continuation_token&#x60; to the system-generated token taken from the &#x60;x-next-token&#x60; header field of the response. A query has reached its last page when the response does not include a token. Pagination requires the &#x60;limit&#x60; and &#x60;continuation_token&#x60; query parameters. | 
 **filter** | **optional.String**| Narrows down the results to only the response objects that satisfy the filter criteria. | 
 **ids** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique resource IDs specified. Enter multiple resource IDs in comma-separated format. The &#x60;ids&#x60; and &#x60;names&#x60; parameters cannot be provided together. | 
 **limit** | **optional.Int32**| Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size. | 
 **names** | [**optional.Interface of []string**](string.md)| Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;. | 
 **offset** | **optional.Int32**| The starting position based on the results of the query in relation to the full set of response objects returned. | 
 **softwareIds** | [**optional.Interface of []string**](string.md)| A comma-separated list of software IDs. | 
 **softwareNames** | [**optional.Interface of []string**](string.md)| A comma-separated list of software names. | 
 **sort** | [**optional.Interface of []string**](string.md)| Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values. | 
 **totalItemCount** | **optional.Bool**| If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;. | 

### Return type

[**InlineResponse200148**](inline_response_200_148.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28SoftwareInstallationsPatch**
> InlineResponse200149 Api28SoftwareInstallationsPatch(ctx, command, currentStepId, optional)
Modify software upgrade

Modifies a software upgrade by continuing, retrying, or aborting it. All `override_checks` are updated before the command being issued if `add_override_checks` is present. The `override_checks` parameter is valid when `command` is `continue` or `retry`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **command** | **string**| A user command that interacts with the upgrade. Commands may only be issued when the upgrade is paused. Valid values are &#x60;continue&#x60;, &#x60;retry&#x60;, and &#x60;abort&#x60;. The &#x60;continue&#x60; command continues a &#x60;paused&#x60; upgrade. The &#x60;retry&#x60; command retries the previous step. The &#x60;abort&#x60; command aborts the upgrade. | 
  **currentStepId** | **string**| The current step &#x60;id&#x60; of the installation. | 
 **optional** | ***SoftwareApiApi28SoftwareInstallationsPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwareApiApi28SoftwareInstallationsPatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of Model28SoftwareinstallationsBody1**](Model28SoftwareinstallationsBody1.md)|  | 
 **authorization** | **optional.**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.**| Supplied by client during request or generated by server. | 

### Return type

[**InlineResponse200149**](inline_response_200_149.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28SoftwareInstallationsPost**
> InlineResponse200149 Api28SoftwareInstallationsPost(ctx, body, softwareIds, optional)
Create a software upgrade

Creates and initiates a software upgrade.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Model28SoftwareinstallationsBody**](Model28SoftwareinstallationsBody.md)|  | 
  **softwareIds** | [**[]string**](string.md)| A comma-separated list of software IDs. | 
 **optional** | ***SoftwareApiApi28SoftwareInstallationsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwareApiApi28SoftwareInstallationsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **optional.**| Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;) | 
 **xRequestID** | **optional.**| Supplied by client during request or generated by server. | 

### Return type

[**InlineResponse200149**](inline_response_200_149.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

