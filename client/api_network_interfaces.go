
/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type NetworkInterfacesApiService service
/*
NetworkInterfacesApiService Delete network interface
Deletes a network interface on a controller.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *NetworkInterfacesApiApi28NetworkInterfacesDeleteOpts - Optional Parameters:
     * @param "Authorization" (optional.String) -  Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;)
     * @param "XRequestID" (optional.String) -  Supplied by client during request or generated by server.
     * @param "Names" (optional.Interface of []string) -  Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;.

*/

type NetworkInterfacesApiApi28NetworkInterfacesDeleteOpts struct {
    Authorization optional.String
    XRequestID optional.String
    Names optional.Interface
}

func (a *NetworkInterfacesApiService) Api28NetworkInterfacesDelete(ctx context.Context, localVarOptionals *NetworkInterfacesApiApi28NetworkInterfacesDeleteOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.8/network-interfaces"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Names.IsSet() {
		localVarQueryParams.Add("names", parameterToString(localVarOptionals.Names.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Authorization.IsSet() {
		localVarHeaderParams["Authorization"] = parameterToString(localVarOptionals.Authorization.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
NetworkInterfacesApiService List network interfaces
Displays all network interfaces for all controllers on the array.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *NetworkInterfacesApiApi28NetworkInterfacesGetOpts - Optional Parameters:
     * @param "Authorization" (optional.String) -  Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;)
     * @param "XRequestID" (optional.String) -  Supplied by client during request or generated by server.
     * @param "ContinuationToken" (optional.String) -  A token used to retrieve the next page of data with some consistency guaranteed. The token is a Base64 encoded value. Set &#x60;continuation_token&#x60; to the system-generated token taken from the &#x60;x-next-token&#x60; header field of the response. A query has reached its last page when the response does not include a token. Pagination requires the &#x60;limit&#x60; and &#x60;continuation_token&#x60; query parameters.
     * @param "Filter" (optional.String) -  Narrows down the results to only the response objects that satisfy the filter criteria.
     * @param "Limit" (optional.Int32) -  Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size.
     * @param "Names" (optional.Interface of []string) -  Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;.
     * @param "Offset" (optional.Int32) -  The starting position based on the results of the query in relation to the full set of response objects returned.
     * @param "Sort" (optional.Interface of []string) -  Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values.
     * @param "TotalItemCount" (optional.Bool) -  If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;.
@return InlineResponse20085
*/

type NetworkInterfacesApiApi28NetworkInterfacesGetOpts struct {
    Authorization optional.String
    XRequestID optional.String
    ContinuationToken optional.String
    Filter optional.String
    Limit optional.Int32
    Names optional.Interface
    Offset optional.Int32
    Sort optional.Interface
    TotalItemCount optional.Bool
}

func (a *NetworkInterfacesApiService) Api28NetworkInterfacesGet(ctx context.Context, localVarOptionals *NetworkInterfacesApiApi28NetworkInterfacesGetOpts) (InlineResponse20085, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20085
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.8/network-interfaces"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ContinuationToken.IsSet() {
		localVarQueryParams.Add("continuation_token", parameterToString(localVarOptionals.ContinuationToken.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Filter.IsSet() {
		localVarQueryParams.Add("filter", parameterToString(localVarOptionals.Filter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Names.IsSet() {
		localVarQueryParams.Add("names", parameterToString(localVarOptionals.Names.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.TotalItemCount.IsSet() {
		localVarQueryParams.Add("total_item_count", parameterToString(localVarOptionals.TotalItemCount.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Authorization.IsSet() {
		localVarHeaderParams["Authorization"] = parameterToString(localVarOptionals.Authorization.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20085
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
NetworkInterfacesApiService Modify network interface
Modifies a network interface on a controller.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param optional nil or *NetworkInterfacesApiApi28NetworkInterfacesPatchOpts - Optional Parameters:
     * @param "Authorization" (optional.String) -  Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;)
     * @param "XRequestID" (optional.String) -  Supplied by client during request or generated by server.
     * @param "Names" (optional.Interface of []string) -  Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;.
@return InlineResponse20086
*/

type NetworkInterfacesApiApi28NetworkInterfacesPatchOpts struct {
    Authorization optional.String
    XRequestID optional.String
    Names optional.Interface
}

func (a *NetworkInterfacesApiService) Api28NetworkInterfacesPatch(ctx context.Context, body Model28NetworkinterfacesPatchBody, localVarOptionals *NetworkInterfacesApiApi28NetworkInterfacesPatchOpts) (InlineResponse20086, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Patch")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20086
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.8/network-interfaces"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Names.IsSet() {
		localVarQueryParams.Add("names", parameterToString(localVarOptionals.Names.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Authorization.IsSet() {
		localVarHeaderParams["Authorization"] = parameterToString(localVarOptionals.Authorization.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20086
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
NetworkInterfacesApiService List network performance statistics
Displays network statistics, historical bandwidth, and error reporting for all specified network interfaces.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *NetworkInterfacesApiApi28NetworkInterfacesPerformanceGetOpts - Optional Parameters:
     * @param "Authorization" (optional.String) -  Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;)
     * @param "XRequestID" (optional.String) -  Supplied by client during request or generated by server.
     * @param "Filter" (optional.String) -  Narrows down the results to only the response objects that satisfy the filter criteria.
     * @param "EndTime" (optional.Int64) -  Displays historical performance data for the specified time window, where &#x60;start_time&#x60; is the beginning of the time window, and &#x60;end_time&#x60; is the end of the time window. The &#x60;start_time&#x60; and &#x60;end_time&#x60; parameters are specified in milliseconds since the UNIX epoch. If &#x60;start_time&#x60; is not specified, the start time will default to one resolution before the end time, meaning that the most recent sample of performance data will be displayed. If &#x60;end_time&#x60;is not specified, the end time will default to the current time. Include the &#x60;resolution&#x60; parameter to display the performance data at the specified resolution. If not specified, &#x60;resolution&#x60; defaults to the lowest valid resolution.
     * @param "Resolution" (optional.Int64) -  The number of milliseconds between samples of historical data. For array-wide performance metrics (&#x60;/arrays/performance&#x60; endpoint), valid values are &#x60;1000&#x60; (1 second), &#x60;30000&#x60; (30 seconds), &#x60;300000&#x60; (5 minutes), &#x60;1800000&#x60; (30 minutes), &#x60;7200000&#x60; (2 hours), &#x60;28800000&#x60; (8 hours), and &#x60;86400000&#x60; (24 hours). For performance metrics on storage objects (&#x60;&lt;object name&gt;/performance&#x60; endpoint), such as volumes, valid values are &#x60;30000&#x60; (30 seconds), &#x60;300000&#x60; (5 minutes), &#x60;1800000&#x60; (30 minutes), &#x60;7200000&#x60; (2 hours), &#x60;28800000&#x60; (8 hours), and &#x60;86400000&#x60; (24 hours). For space metrics, (&#x60;&lt;object name&gt;/space&#x60; endpoint), valid values are &#x60;300000&#x60; (5 minutes), &#x60;1800000&#x60; (30 minutes), &#x60;7200000&#x60; (2 hours), &#x60;28800000&#x60; (8 hours), and &#x60;86400000&#x60; (24 hours). Include the &#x60;start_time&#x60; parameter to display the performance data starting at the specified start time. If &#x60;start_time&#x60; is not specified, the start time will default to one resolution before the end time, meaning that the most recent sample of performance data will be displayed. Include the &#x60;end_time&#x60; parameter to display the performance data until the specified end time. If &#x60;end_time&#x60;is not specified, the end time will default to the current time. If the &#x60;resolution&#x60; parameter is not specified but either the &#x60;start_time&#x60; or &#x60;end_time&#x60; parameter is, then &#x60;resolution&#x60; will default to the lowest valid resolution.
     * @param "StartTime" (optional.Int64) -  Displays historical performance data for the specified time window, where &#x60;start_time&#x60; is the beginning of the time window, and &#x60;end_time&#x60; is the end of the time window. The &#x60;start_time&#x60; and &#x60;end_time&#x60; parameters are specified in milliseconds since the UNIX epoch. If &#x60;start_time&#x60; is not specified, the start time will default to one resolution before the end time, meaning that the most recent sample of performance data will be displayed. If &#x60;end_time&#x60;is not specified, the end time will default to the current time. Include the &#x60;resolution&#x60; parameter to display the performance data at the specified resolution. If not specified, &#x60;resolution&#x60; defaults to the lowest valid resolution.
     * @param "Limit" (optional.Int32) -  Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size.
     * @param "Names" (optional.Interface of []string) -  Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;.
     * @param "Offset" (optional.Int32) -  The starting position based on the results of the query in relation to the full set of response objects returned.
     * @param "Sort" (optional.Interface of []string) -  Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values.
     * @param "TotalItemCount" (optional.Bool) -  If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;.
     * @param "TotalOnly" (optional.Bool) -  If set to &#x60;true&#x60;, returns the aggregate value of all items after filtering. Where it makes more sense, the average value is displayed instead. The values are displayed for each name where meaningful. If &#x60;total_only&#x3D;true&#x60;, the &#x60;items&#x60; list will be empty.
@return InlineResponse20087
*/

type NetworkInterfacesApiApi28NetworkInterfacesPerformanceGetOpts struct {
    Authorization optional.String
    XRequestID optional.String
    Filter optional.String
    EndTime optional.Int64
    Resolution optional.Int64
    StartTime optional.Int64
    Limit optional.Int32
    Names optional.Interface
    Offset optional.Int32
    Sort optional.Interface
    TotalItemCount optional.Bool
    TotalOnly optional.Bool
}

func (a *NetworkInterfacesApiService) Api28NetworkInterfacesPerformanceGet(ctx context.Context, localVarOptionals *NetworkInterfacesApiApi28NetworkInterfacesPerformanceGetOpts) (InlineResponse20087, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20087
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.8/network-interfaces/performance"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Filter.IsSet() {
		localVarQueryParams.Add("filter", parameterToString(localVarOptionals.Filter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndTime.IsSet() {
		localVarQueryParams.Add("end_time", parameterToString(localVarOptionals.EndTime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Resolution.IsSet() {
		localVarQueryParams.Add("resolution", parameterToString(localVarOptionals.Resolution.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartTime.IsSet() {
		localVarQueryParams.Add("start_time", parameterToString(localVarOptionals.StartTime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Names.IsSet() {
		localVarQueryParams.Add("names", parameterToString(localVarOptionals.Names.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.TotalItemCount.IsSet() {
		localVarQueryParams.Add("total_item_count", parameterToString(localVarOptionals.TotalItemCount.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TotalOnly.IsSet() {
		localVarQueryParams.Add("total_only", parameterToString(localVarOptionals.TotalOnly.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Authorization.IsSet() {
		localVarHeaderParams["Authorization"] = parameterToString(localVarOptionals.Authorization.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20087
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
NetworkInterfacesApiService Create network interface
Creates a network interface on a controller on the array.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param optional nil or *NetworkInterfacesApiApi28NetworkInterfacesPostOpts - Optional Parameters:
     * @param "Authorization" (optional.String) -  Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;)
     * @param "XRequestID" (optional.String) -  Supplied by client during request or generated by server.
     * @param "Names" (optional.Interface of []string) -  Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;.
@return InlineResponse20086
*/

type NetworkInterfacesApiApi28NetworkInterfacesPostOpts struct {
    Authorization optional.String
    XRequestID optional.String
    Names optional.Interface
}

func (a *NetworkInterfacesApiService) Api28NetworkInterfacesPost(ctx context.Context, body Model28NetworkinterfacesBody, localVarOptionals *NetworkInterfacesApiApi28NetworkInterfacesPostOpts) (InlineResponse20086, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20086
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.8/network-interfaces"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Names.IsSet() {
		localVarQueryParams.Add("names", parameterToString(localVarOptionals.Names.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Authorization.IsSet() {
		localVarHeaderParams["Authorization"] = parameterToString(localVarOptionals.Authorization.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20086
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
