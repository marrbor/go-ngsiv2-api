/*
 * FIWARE-NGSI v2 Specification
 *
 * TODO: Add a description
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type AttributeValueApiService service

/* 
AttributeValueApiService Get Attribute Value
This operation returns the &#x60;value&#x60; property with the value of the attribute. * If attribute value is JSON Array or Object:   * If &#x60;Accept&#x60; header can be expanded to &#x60;application/json&#x60; or &#x60;text/plain&#x60; return the value as a JSON with a     response type of application/json or text/plain (whichever is the first in &#x60;Accept&#x60; header or     &#x60;application/json&#x60; in case of &#x60;Accept: *_/_*&#x60;).   * Else return a HTTP error \&quot;406 Not Acceptable: accepted MIME types: application/json, text/plain\&quot; * If attribute value is a string, number, null or boolean:   * If &#x60;Accept&#x60; header can be expanded to text/plain return the value as text. In case of a string, citation     marks are used at the begining and end.   * Else return a HTTP error \&quot;406 Not Acceptable: accepted MIME types: text/plain\&quot; Response: * Successful operation uses 200 OK. * Errors use a non-2xx and (optionally) an error payload. See subsection on \&quot;Error Responses\&quot; for   more details.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param entityId Id of the entity in question
 * @param attrName Name of the attribute to be retrieved.
 * @param optional nil or *GetAttributeValueOpts - Optional Parameters:
     * @param "Type_" (optional.String) -  Entity type, to avoid ambiguity in case there are several entities with the same entity id.

@return GetAttributeValueResponse
*/

type GetAttributeValueOpts struct { 
	Type_ optional.String
}

func (a *AttributeValueApiService) GetAttributeValue(ctx context.Context, entityId string, attrName string, localVarOptionals *GetAttributeValueOpts) (GetAttributeValueResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue GetAttributeValueResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/entities/{entityId}/attrs/{attrName}/value"
	localVarPath = strings.Replace(localVarPath, "{"+"entityId"+"}", fmt.Sprintf("%v", entityId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attrName"+"}", fmt.Sprintf("%v", attrName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
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
			var v GetAttributeValueResponse
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
AttributeValueApiService Update Attribute Value
The request payload is the new attribute value. * If the request payload MIME type is &#x60;application/json&#x60;, then the value of the attribute is set to   the JSON object or array coded in the payload (if the payload is not a valid JSON document,   then an error is returned). * If the request payload MIME type is &#x60;text/plain&#x60;, then the following algorithm is applied to the   payload:   * If the payload starts and ends with citation-marks (&#x60;\&quot;&#x60;), the value is taken as a string     (the citation marks themselves are not considered part of the string)   * If &#x60;true&#x60; or &#x60;false&#x60;, the value is taken as a boolean.   * If &#x60;null&#x60;, the value is taken as null.   * If these first three tests &#39;fail&#39;, the text is interpreted as a number.   * If not a valid number, then an error is returned and the attribute&#39;s value is unchanged. The payload MIME type in the request is specified in the &#x60;Content-Type&#x60; HTTP header. Response: * Successful operation uses 204 No Content * Errors use a non-2xx and (optionally) an error payload. See subsection on \&quot;Error Responses\&quot; for   more details.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param entityId Id of the entity to be updated.
 * @param attrName Attribute name.
 * @param contentType 
 * @param body 
 * @param optional nil or *UpdateAttributeValueOpts - Optional Parameters:
     * @param "Type_" (optional.String) -  Entity type, to avoid ambiguity in case there are several entities with the same entity id.


*/

type UpdateAttributeValueOpts struct { 
	Type_ optional.String
}

func (a *AttributeValueApiService) UpdateAttributeValue(ctx context.Context, entityId string, attrName string, contentType string, body UpdateAttributeValueRequest, localVarOptionals *UpdateAttributeValueOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/entities/{entityId}/attrs/{attrName}/value"
	localVarPath = strings.Replace(localVarPath, "{"+"entityId"+"}", fmt.Sprintf("%v", entityId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attrName"+"}", fmt.Sprintf("%v", attrName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
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
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	// body params
	localVarPostBody = &body
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