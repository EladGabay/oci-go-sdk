// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package rover

import (
	"github.com/oracle/oci-go-sdk/v39/common"
	"net/http"
)

// GetRoverNodeEncryptionKeyRequest wrapper for the GetRoverNodeEncryptionKey operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/GetRoverNodeEncryptionKey.go.html to see an example of how to use GetRoverNodeEncryptionKeyRequest.
type GetRoverNodeEncryptionKeyRequest struct {

	// Serial number of the rover node.
	RoverNodeId *string `mandatory:"true" contributesTo:"path" name:"roverNodeId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetRoverNodeEncryptionKeyRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetRoverNodeEncryptionKeyRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetRoverNodeEncryptionKeyRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetRoverNodeEncryptionKeyResponse wrapper for the GetRoverNodeEncryptionKey operation
type GetRoverNodeEncryptionKeyResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The RoverNodeEncryptionKey instance
	RoverNodeEncryptionKey `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetRoverNodeEncryptionKeyResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetRoverNodeEncryptionKeyResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
