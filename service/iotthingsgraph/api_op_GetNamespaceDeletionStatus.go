// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetNamespaceDeletionStatusInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetNamespaceDeletionStatusInput) String() string {
	return awsutil.Prettify(s)
}

type GetNamespaceDeletionStatusOutput struct {
	_ struct{} `type:"structure"`

	// An error code returned by the namespace deletion task.
	ErrorCode NamespaceDeletionStatusErrorCodes `locationName:"errorCode" type:"string" enum:"true"`

	// An error code returned by the namespace deletion task.
	ErrorMessage *string `locationName:"errorMessage" type:"string"`

	// The ARN of the namespace that is being deleted.
	NamespaceArn *string `locationName:"namespaceArn" type:"string"`

	// The name of the namespace that is being deleted.
	NamespaceName *string `locationName:"namespaceName" type:"string"`

	// The status of the deletion request.
	Status NamespaceDeletionStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s GetNamespaceDeletionStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetNamespaceDeletionStatus = "GetNamespaceDeletionStatus"

// GetNamespaceDeletionStatusRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Gets the status of a namespace deletion task.
//
//    // Example sending a request using GetNamespaceDeletionStatusRequest.
//    req := client.GetNamespaceDeletionStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/GetNamespaceDeletionStatus
func (c *Client) GetNamespaceDeletionStatusRequest(input *GetNamespaceDeletionStatusInput) GetNamespaceDeletionStatusRequest {
	op := &aws.Operation{
		Name:       opGetNamespaceDeletionStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetNamespaceDeletionStatusInput{}
	}

	req := c.newRequest(op, input, &GetNamespaceDeletionStatusOutput{})

	return GetNamespaceDeletionStatusRequest{Request: req, Input: input, Copy: c.GetNamespaceDeletionStatusRequest}
}

// GetNamespaceDeletionStatusRequest is the request type for the
// GetNamespaceDeletionStatus API operation.
type GetNamespaceDeletionStatusRequest struct {
	*aws.Request
	Input *GetNamespaceDeletionStatusInput
	Copy  func(*GetNamespaceDeletionStatusInput) GetNamespaceDeletionStatusRequest
}

// Send marshals and sends the GetNamespaceDeletionStatus API request.
func (r GetNamespaceDeletionStatusRequest) Send(ctx context.Context) (*GetNamespaceDeletionStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetNamespaceDeletionStatusResponse{
		GetNamespaceDeletionStatusOutput: r.Request.Data.(*GetNamespaceDeletionStatusOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetNamespaceDeletionStatusResponse is the response type for the
// GetNamespaceDeletionStatus API operation.
type GetNamespaceDeletionStatusResponse struct {
	*GetNamespaceDeletionStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetNamespaceDeletionStatus request.
func (r *GetNamespaceDeletionStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
