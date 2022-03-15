// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v62/common"
	"net/http"
	"strings"
)

// ListRunsRequest wrapper for the ListRuns operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataflow/ListRuns.go.html to see an example of how to use ListRunsRequest.
type ListRunsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique identifier for the request. If provided, the returned request ID will include this value.
	// Otherwise, a random request ID will be generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The ID of the application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// The OCID of the user who created the resource.
	OwnerPrincipalId *string `mandatory:"false" contributesTo:"query" name:"ownerPrincipalId"`

	// The displayName prefix.
	DisplayNameStartsWith *string `mandatory:"false" contributesTo:"query" name:"displayNameStartsWith"`

	// The LifecycleState of the run.
	LifecycleState ListRunsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The epoch time that the resource was created.
	TimeCreatedGreaterThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThan"`

	// The maximum number of results to return in a paginated `List` call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` or `opc-prev-page` response header from the last `List` call
	// to sent back to server for getting the next page of results.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field used to sort the results. Multiple fields are not supported.
	SortBy ListRunsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The ordering of results in ascending or descending order.
	SortOrder ListRunsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The query parameter for the Spark application name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRunsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRunsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRunsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRunsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRunsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRunsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListRunsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRunsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRunsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRunsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRunsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRunsResponse wrapper for the ListRuns operation
type ListRunsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []RunSummary instances
	Items []RunSummary `presentIn:"body"`

	// Retrieves the previous page of results.
	// When this header appears in the response, previous pages of results exist.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Retrieves the next page of results. When this header appears in the response,
	// additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListRunsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRunsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRunsLifecycleStateEnum Enum with underlying type: string
type ListRunsLifecycleStateEnum string

// Set of constants representing the allowable values for ListRunsLifecycleStateEnum
const (
	ListRunsLifecycleStateAccepted   ListRunsLifecycleStateEnum = "ACCEPTED"
	ListRunsLifecycleStateInProgress ListRunsLifecycleStateEnum = "IN_PROGRESS"
	ListRunsLifecycleStateCanceling  ListRunsLifecycleStateEnum = "CANCELING"
	ListRunsLifecycleStateCanceled   ListRunsLifecycleStateEnum = "CANCELED"
	ListRunsLifecycleStateFailed     ListRunsLifecycleStateEnum = "FAILED"
	ListRunsLifecycleStateSucceeded  ListRunsLifecycleStateEnum = "SUCCEEDED"
	ListRunsLifecycleStateStopping   ListRunsLifecycleStateEnum = "STOPPING"
	ListRunsLifecycleStateStopped    ListRunsLifecycleStateEnum = "STOPPED"
)

var mappingListRunsLifecycleStateEnum = map[string]ListRunsLifecycleStateEnum{
	"ACCEPTED":    ListRunsLifecycleStateAccepted,
	"IN_PROGRESS": ListRunsLifecycleStateInProgress,
	"CANCELING":   ListRunsLifecycleStateCanceling,
	"CANCELED":    ListRunsLifecycleStateCanceled,
	"FAILED":      ListRunsLifecycleStateFailed,
	"SUCCEEDED":   ListRunsLifecycleStateSucceeded,
	"STOPPING":    ListRunsLifecycleStateStopping,
	"STOPPED":     ListRunsLifecycleStateStopped,
}

var mappingListRunsLifecycleStateEnumLowerCase = map[string]ListRunsLifecycleStateEnum{
	"accepted":    ListRunsLifecycleStateAccepted,
	"in_progress": ListRunsLifecycleStateInProgress,
	"canceling":   ListRunsLifecycleStateCanceling,
	"canceled":    ListRunsLifecycleStateCanceled,
	"failed":      ListRunsLifecycleStateFailed,
	"succeeded":   ListRunsLifecycleStateSucceeded,
	"stopping":    ListRunsLifecycleStateStopping,
	"stopped":     ListRunsLifecycleStateStopped,
}

// GetListRunsLifecycleStateEnumValues Enumerates the set of values for ListRunsLifecycleStateEnum
func GetListRunsLifecycleStateEnumValues() []ListRunsLifecycleStateEnum {
	values := make([]ListRunsLifecycleStateEnum, 0)
	for _, v := range mappingListRunsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListRunsLifecycleStateEnumStringValues Enumerates the set of values in String for ListRunsLifecycleStateEnum
func GetListRunsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"CANCELING",
		"CANCELED",
		"FAILED",
		"SUCCEEDED",
		"STOPPING",
		"STOPPED",
	}
}

// GetMappingListRunsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRunsLifecycleStateEnum(val string) (ListRunsLifecycleStateEnum, bool) {
	enum, ok := mappingListRunsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRunsSortByEnum Enum with underlying type: string
type ListRunsSortByEnum string

// Set of constants representing the allowable values for ListRunsSortByEnum
const (
	ListRunsSortByTimecreated               ListRunsSortByEnum = "timeCreated"
	ListRunsSortByDisplayname               ListRunsSortByEnum = "displayName"
	ListRunsSortByLanguage                  ListRunsSortByEnum = "language"
	ListRunsSortByRundurationinmilliseconds ListRunsSortByEnum = "runDurationInMilliseconds"
	ListRunsSortByLifecyclestate            ListRunsSortByEnum = "lifecycleState"
	ListRunsSortByTotalocpu                 ListRunsSortByEnum = "totalOCpu"
	ListRunsSortByDatareadinbytes           ListRunsSortByEnum = "dataReadInBytes"
	ListRunsSortByDatawritteninbytes        ListRunsSortByEnum = "dataWrittenInBytes"
)

var mappingListRunsSortByEnum = map[string]ListRunsSortByEnum{
	"timeCreated":               ListRunsSortByTimecreated,
	"displayName":               ListRunsSortByDisplayname,
	"language":                  ListRunsSortByLanguage,
	"runDurationInMilliseconds": ListRunsSortByRundurationinmilliseconds,
	"lifecycleState":            ListRunsSortByLifecyclestate,
	"totalOCpu":                 ListRunsSortByTotalocpu,
	"dataReadInBytes":           ListRunsSortByDatareadinbytes,
	"dataWrittenInBytes":        ListRunsSortByDatawritteninbytes,
}

var mappingListRunsSortByEnumLowerCase = map[string]ListRunsSortByEnum{
	"timecreated":               ListRunsSortByTimecreated,
	"displayname":               ListRunsSortByDisplayname,
	"language":                  ListRunsSortByLanguage,
	"rundurationinmilliseconds": ListRunsSortByRundurationinmilliseconds,
	"lifecyclestate":            ListRunsSortByLifecyclestate,
	"totalocpu":                 ListRunsSortByTotalocpu,
	"datareadinbytes":           ListRunsSortByDatareadinbytes,
	"datawritteninbytes":        ListRunsSortByDatawritteninbytes,
}

// GetListRunsSortByEnumValues Enumerates the set of values for ListRunsSortByEnum
func GetListRunsSortByEnumValues() []ListRunsSortByEnum {
	values := make([]ListRunsSortByEnum, 0)
	for _, v := range mappingListRunsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRunsSortByEnumStringValues Enumerates the set of values in String for ListRunsSortByEnum
func GetListRunsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"language",
		"runDurationInMilliseconds",
		"lifecycleState",
		"totalOCpu",
		"dataReadInBytes",
		"dataWrittenInBytes",
	}
}

// GetMappingListRunsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRunsSortByEnum(val string) (ListRunsSortByEnum, bool) {
	enum, ok := mappingListRunsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRunsSortOrderEnum Enum with underlying type: string
type ListRunsSortOrderEnum string

// Set of constants representing the allowable values for ListRunsSortOrderEnum
const (
	ListRunsSortOrderAsc  ListRunsSortOrderEnum = "ASC"
	ListRunsSortOrderDesc ListRunsSortOrderEnum = "DESC"
)

var mappingListRunsSortOrderEnum = map[string]ListRunsSortOrderEnum{
	"ASC":  ListRunsSortOrderAsc,
	"DESC": ListRunsSortOrderDesc,
}

var mappingListRunsSortOrderEnumLowerCase = map[string]ListRunsSortOrderEnum{
	"asc":  ListRunsSortOrderAsc,
	"desc": ListRunsSortOrderDesc,
}

// GetListRunsSortOrderEnumValues Enumerates the set of values for ListRunsSortOrderEnum
func GetListRunsSortOrderEnumValues() []ListRunsSortOrderEnum {
	values := make([]ListRunsSortOrderEnum, 0)
	for _, v := range mappingListRunsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRunsSortOrderEnumStringValues Enumerates the set of values in String for ListRunsSortOrderEnum
func GetListRunsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRunsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRunsSortOrderEnum(val string) (ListRunsSortOrderEnum, bool) {
	enum, ok := mappingListRunsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
