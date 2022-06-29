# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api28LoginPost**](AuthorizationApi.md#Api28LoginPost) | **Post** /api/2.8/login | POST login (placeholder)
[**Api28LogoutPost**](AuthorizationApi.md#Api28LogoutPost) | **Post** /api/2.8/logout | POST logout (placeholder)
[**ApiApiVersionGet**](AuthorizationApi.md#ApiApiVersionGet) | **Get** /api/api_version | List available API versions
[**Oauth210TokenPost**](AuthorizationApi.md#Oauth210TokenPost) | **Post** /oauth2/1.0/token | Get access token

# **Api28LoginPost**
> InlineResponse2001 Api28LoginPost(ctx, optional)
POST login (placeholder)

Exchange an API token for a session token. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthorizationApiApi28LoginPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiApi28LoginPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiToken** | **optional.String**| API token for a user. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api28LogoutPost**
> Api28LogoutPost(ctx, optional)
POST logout (placeholder)

Invalidate a session token. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthorizationApiApi28LogoutPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiApi28LogoutPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **optional.String**| Session token for a user. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiApiVersionGet**
> InlineResponse200 ApiApiVersionGet(ctx, )
List available API versions

Returns a list of available API versions. No authentication is required to access this endpoint. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Oauth210TokenPost**
> OauthTokenResponse Oauth210TokenPost(ctx, grantType, subjectToken, subjectTokenType)
Get access token

Exchanges an ID Token for an OAuth 2.0 access token. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **grantType** | **string**|  | 
  **subjectToken** | **string**|  | 
  **subjectTokenType** | **string**|  | 

### Return type

[**OauthTokenResponse**](oauth_token_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

