
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

type DirectoryQuotasApiService service
/*
DirectoryQuotasApiService List directories with attached quota policies
Displays a list of directories and the quota policies attached to them. Directories with multiple policies are listed repeatedly (once per policy). The directories without a policy attached are not listed.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *DirectoryQuotasApiApi28DirectoryQuotasGetOpts - Optional Parameters:
     * @param "Authorization" (optional.String) -  Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;)
     * @param "XRequestID" (optional.String) -  Supplied by client during request or generated by server.
     * @param "ContinuationToken" (optional.String) -  A token used to retrieve the next page of data with some consistency guaranteed. The token is a Base64 encoded value. Set &#x60;continuation_token&#x60; to the system-generated token taken from the &#x60;x-next-token&#x60; header field of the response. A query has reached its last page when the response does not include a token. Pagination requires the &#x60;limit&#x60; and &#x60;continuation_token&#x60; query parameters.
     * @param "DirectoryIds" (optional.Interface of []string) -  Performs the operation on the unique managed directory IDs specified. Enter multiple managed directory IDs in comma-separated format. The &#x60;directory_ids&#x60; and &#x60;directory_names&#x60; parameters cannot be provided together.
     * @param "DirectoryNames" (optional.Interface of []string) -  Performs the operation on the managed directory names specified. Enter multiple full managed directory names in comma-separated format. For example, &#x60;fs:dir01,fs:dir02&#x60;.
     * @param "Filter" (optional.String) -  Narrows down the results to only the response objects that satisfy the filter criteria.
     * @param "Limit" (optional.Int32) -  Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size.
     * @param "Offset" (optional.Int32) -  The starting position based on the results of the query in relation to the full set of response objects returned.
     * @param "PolicyIds" (optional.Interface of []string) -  Performs the operation on the unique policy IDs specified. Enter multiple policy IDs in comma-separated format. The &#x60;policy_ids&#x60; and &#x60;policy_names&#x60; parameters cannot be provided together.
     * @param "PolicyNames" (optional.Interface of []string) -  Performs the operation on the policy names specified. Enter multiple policy names in comma-separated format. For example, &#x60;name01,name02&#x60;.
     * @param "Sort" (optional.Interface of []string) -  Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values.
     * @param "TotalItemCount" (optional.Bool) -  If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;.
@return InlineResponse20054
*/

type DirectoryQuotasApiApi28DirectoryQuotasGetOpts struct {
    Authorization optional.String
    XRequestID optional.String
    ContinuationToken optional.String
    DirectoryIds optional.Interface
    DirectoryNames optional.Interface
    Filter optional.String
    Limit optional.Int32
    Offset optional.Int32
    PolicyIds optional.Interface
    PolicyNames optional.Interface
    Sort optional.Interface
    TotalItemCount optional.Bool
}

func (a *DirectoryQuotasApiService) Api28DirectoryQuotasGet(ctx context.Context, localVarOptionals *DirectoryQuotasApiApi28DirectoryQuotasGetOpts) (InlineResponse20054, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20054
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.8/directory-quotas"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ContinuationToken.IsSet() {
		localVarQueryParams.Add("continuation_token", parameterToString(localVarOptionals.ContinuationToken.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DirectoryIds.IsSet() {
		localVarQueryParams.Add("directory_ids", parameterToString(localVarOptionals.DirectoryIds.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.DirectoryNames.IsSet() {
		localVarQueryParams.Add("directory_names", parameterToString(localVarOptionals.DirectoryNames.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Filter.IsSet() {
		localVarQueryParams.Add("filter", parameterToString(localVarOptionals.Filter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PolicyIds.IsSet() {
		localVarQueryParams.Add("policy_ids", parameterToString(localVarOptionals.PolicyIds.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.PolicyNames.IsSet() {
		localVarQueryParams.Add("policy_names", parameterToString(localVarOptionals.PolicyNames.Value(), "csv"))
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
			var v InlineResponse20054
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
